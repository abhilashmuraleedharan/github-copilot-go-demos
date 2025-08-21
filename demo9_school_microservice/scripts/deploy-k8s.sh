#!/bin/bash

# [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-08-21
# School Microservice Deployment Script for Kubernetes

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Configuration
NAMESPACE="school-demo"
RELEASE_NAME="school-microservice"
CHART_PATH="./helm/school-microservice"
VALUES_FILE=""
ENVIRONMENT="development"

# Functions
log_info() {
    echo -e "${GREEN}[INFO]${NC} $1"
}

log_warn() {
    echo -e "${YELLOW}[WARN]${NC} $1"
}

log_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

show_usage() {
    echo "Usage: $0 [OPTIONS]"
    echo ""
    echo "Options:"
    echo "  -e, --environment ENV    Environment (development|staging|production) [default: development]"
    echo "  -n, --namespace NAME     Kubernetes namespace [default: school-demo]"
    echo "  -r, --release NAME       Helm release name [default: school-microservice]"
    echo "  -v, --values FILE        Custom values file"
    echo "  -h, --help              Show this help message"
    echo ""
    echo "Examples:"
    echo "  $0 --environment production"
    echo "  $0 --namespace school-prod --environment production"
    echo "  $0 --values custom-values.yaml"
}

check_prerequisites() {
    log_info "Checking prerequisites..."
    
    # Check kubectl
    if ! command -v kubectl &> /dev/null; then
        log_error "kubectl is not installed or not in PATH"
        exit 1
    fi
    
    # Check helm
    if ! command -v helm &> /dev/null; then
        log_error "helm is not installed or not in PATH"
        exit 1
    fi
    
    # Check cluster connectivity
    if ! kubectl cluster-info &> /dev/null; then
        log_error "Cannot connect to Kubernetes cluster"
        exit 1
    fi
    
    log_info "Prerequisites check passed"
}

create_namespace() {
    log_info "Creating namespace $NAMESPACE if it doesn't exist..."
    kubectl create namespace "$NAMESPACE" --dry-run=client -o yaml | kubectl apply -f -
}

add_helm_repos() {
    log_info "Adding required Helm repositories..."
    helm repo add couchbase https://couchbase-partners.github.io/helm-charts/ || true
    helm repo update
}

deploy_application() {
    log_info "Deploying School Microservice to $ENVIRONMENT environment..."
    
    local helm_args=(
        "upgrade" "--install" "$RELEASE_NAME"
        "$CHART_PATH"
        "--namespace" "$NAMESPACE"
        "--create-namespace"
        "--wait"
        "--timeout" "600s"
    )
    
    # Add environment-specific values file
    if [[ -f "$CHART_PATH/values-$ENVIRONMENT.yaml" ]]; then
        helm_args+=("--values" "$CHART_PATH/values-$ENVIRONMENT.yaml")
        log_info "Using environment-specific values: values-$ENVIRONMENT.yaml"
    fi
    
    # Add custom values file if specified
    if [[ -n "$VALUES_FILE" ]]; then
        if [[ -f "$VALUES_FILE" ]]; then
            helm_args+=("--values" "$VALUES_FILE")
            log_info "Using custom values file: $VALUES_FILE"
        else
            log_error "Custom values file not found: $VALUES_FILE"
            exit 1
        fi
    fi
    
    # Execute helm deployment
    helm "${helm_args[@]}"
}

verify_deployment() {
    log_info "Verifying deployment..."
    
    # Wait for deployments to be ready
    local services=("students" "teachers" "classes" "academics" "achievements")
    
    for service in "${services[@]}"; do
        local deployment_name="$RELEASE_NAME-$service"
        log_info "Waiting for deployment $deployment_name to be ready..."
        kubectl wait --for=condition=available --timeout=300s deployment/"$deployment_name" -n "$NAMESPACE"
    done
    
    # Check pod status
    log_info "Checking pod status..."
    kubectl get pods -n "$NAMESPACE" -l app.kubernetes.io/name=school-microservice
    
    # Check service status
    log_info "Checking service status..."
    kubectl get services -n "$NAMESPACE" -l app.kubernetes.io/name=school-microservice
}

test_health_endpoints() {
    log_info "Testing health endpoints..."
    
    local services=("students:8081" "teachers:8082" "classes:8083" "academics:8084" "achievements:8085")
    
    for service_port in "${services[@]}"; do
        local service="${service_port%:*}"
        local port="${service_port#*:}"
        local service_name="$RELEASE_NAME-$service"
        
        log_info "Testing $service service health endpoint..."
        
        # Port forward in background
        kubectl port-forward "svc/$service_name" "$port:$port" -n "$NAMESPACE" &
        local pf_pid=$!
        
        # Wait a moment for port forward to establish
        sleep 2
        
        # Test health endpoint
        if curl -s -f "http://localhost:$port/health" > /dev/null; then
            log_info "$service service is healthy"
        else
            log_warn "$service service health check failed"
        fi
        
        # Kill port forward
        kill $pf_pid 2>/dev/null || true
        wait $pf_pid 2>/dev/null || true
    done
}

show_access_info() {
    log_info "Deployment completed successfully!"
    echo ""
    echo "Access Information:"
    echo "==================="
    echo "Namespace: $NAMESPACE"
    echo "Release: $RELEASE_NAME"
    echo ""
    
    # Show ingress information
    local ingress_ip=$(kubectl get ingress "$RELEASE_NAME" -n "$NAMESPACE" -o jsonpath='{.status.loadBalancer.ingress[0].ip}' 2>/dev/null || echo "pending")
    local ingress_host=$(kubectl get ingress "$RELEASE_NAME" -n "$NAMESPACE" -o jsonpath='{.spec.rules[0].host}' 2>/dev/null || echo "not configured")
    
    if [[ "$ingress_host" != "not configured" ]]; then
        echo "Ingress Host: $ingress_host"
        echo "Ingress IP: $ingress_ip"
        echo ""
        echo "API Endpoints:"
        echo "  Students:     http://$ingress_host/students"
        echo "  Teachers:     http://$ingress_host/teachers"
        echo "  Classes:      http://$ingress_host/classes"
        echo "  Academics:    http://$ingress_host/academics"
        echo "  Achievements: http://$ingress_host/achievements"
    else
        echo "To access services via port forwarding:"
        echo "  kubectl port-forward svc/$RELEASE_NAME-students 8081:8081 -n $NAMESPACE"
        echo "  kubectl port-forward svc/$RELEASE_NAME-teachers 8082:8082 -n $NAMESPACE"
        echo "  kubectl port-forward svc/$RELEASE_NAME-classes 8083:8083 -n $NAMESPACE"
        echo "  kubectl port-forward svc/$RELEASE_NAME-academics 8084:8084 -n $NAMESPACE"
        echo "  kubectl port-forward svc/$RELEASE_NAME-achievements 8085:8085 -n $NAMESPACE"
    fi
    
    echo ""
    echo "Useful Commands:"
    echo "  View pods:     kubectl get pods -n $NAMESPACE"
    echo "  View logs:     kubectl logs -f deployment/$RELEASE_NAME-students -n $NAMESPACE"
    echo "  View services: kubectl get svc -n $NAMESPACE"
    echo "  Uninstall:     helm uninstall $RELEASE_NAME -n $NAMESPACE"
}

# Parse command line arguments
while [[ $# -gt 0 ]]; do
    case $1 in
        -e|--environment)
            ENVIRONMENT="$2"
            shift 2
            ;;
        -n|--namespace)
            NAMESPACE="$2"
            shift 2
            ;;
        -r|--release)
            RELEASE_NAME="$2"
            shift 2
            ;;
        -v|--values)
            VALUES_FILE="$2"
            shift 2
            ;;
        -h|--help)
            show_usage
            exit 0
            ;;
        *)
            log_error "Unknown option $1"
            show_usage
            exit 1
            ;;
    esac
done

# Validate environment
if [[ ! "$ENVIRONMENT" =~ ^(development|staging|production)$ ]]; then
    log_error "Invalid environment: $ENVIRONMENT. Must be one of: development, staging, production"
    exit 1
fi

# Main execution
main() {
    log_info "Starting School Microservice deployment..."
    log_info "Environment: $ENVIRONMENT"
    log_info "Namespace: $NAMESPACE"
    log_info "Release: $RELEASE_NAME"
    
    check_prerequisites
    create_namespace
    add_helm_repos
    deploy_application
    verify_deployment
    test_health_endpoints
    show_access_info
}

# Execute main function
main
