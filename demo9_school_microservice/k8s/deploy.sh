#!/bin/bash

# School Management System - Kubernetes Deployment Script
# This script helps deploy the School Management System to a Kubernetes cluster

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Configuration
NAMESPACE="school-demo"
CHART_NAME="school-management"
RELEASE_NAME="school-mgmt"
CHART_PATH="./k8s/helm/school-management"

# Function to print colored output
print_status() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

print_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Function to check if command exists
check_command() {
    if ! command -v $1 &> /dev/null; then
        print_error "$1 is not installed. Please install it and try again."
        exit 1
    fi
}

# Function to check cluster connectivity
check_cluster() {
    print_status "Checking Kubernetes cluster connectivity..."
    if ! kubectl cluster-info &> /dev/null; then
        print_error "Cannot connect to Kubernetes cluster. Please check your kubeconfig."
        exit 1
    fi
    print_success "Connected to Kubernetes cluster"
}

# Function to build Docker images
build_images() {
    print_status "Building Docker images..."
    
    # Build API Gateway
    print_status "Building API Gateway image..."
    docker build -f k8s/dockerfiles/api-gateway.Dockerfile -t school-mgmt/api-gateway:latest .
    
    # Build Student Service
    print_status "Building Student Service image..."
    docker build -f k8s/dockerfiles/student-service.Dockerfile -t school-mgmt/student-service:latest .
    
    # Build Teacher Service
    print_status "Building Teacher Service image..."
    docker build -f k8s/dockerfiles/teacher-service.Dockerfile -t school-mgmt/teacher-service:latest .
    
    # Build Academic Service
    print_status "Building Academic Service image..."
    docker build -f k8s/dockerfiles/academic-service.Dockerfile -t school-mgmt/academic-service:latest .
    
    # Build Achievement Service
    print_status "Building Achievement Service image..."
    docker build -f k8s/dockerfiles/achievement-service.Dockerfile -t school-mgmt/achievement-service:latest .
    
    print_success "All Docker images built successfully"
}

# Function to load images to kind cluster (if using kind)
load_images_to_kind() {
    if kubectl config current-context | grep -q "kind"; then
        print_status "Detected kind cluster. Loading images..."
        kind load docker-image school-mgmt/api-gateway:latest
        kind load docker-image school-mgmt/student-service:latest
        kind load docker-image school-mgmt/teacher-service:latest
        kind load docker-image school-mgmt/academic-service:latest
        kind load docker-image school-mgmt/achievement-service:latest
        print_success "Images loaded to kind cluster"
    fi
}

# Function to deploy with Helm
deploy_helm() {
    print_status "Deploying School Management System with Helm..."
    
    # Create namespace if it doesn't exist
    kubectl create namespace $NAMESPACE --dry-run=client -o yaml | kubectl apply -f -
    
    # Deploy or upgrade the Helm release
    helm upgrade --install $RELEASE_NAME $CHART_PATH \
        --namespace $NAMESPACE \
        --set-string global.imageTag=latest \
        --wait --timeout=10m
    
    print_success "Helm deployment completed"
}

# Function to check deployment status
check_deployment() {
    print_status "Checking deployment status..."
    
    # Wait for all deployments to be ready
    kubectl wait --for=condition=available --timeout=300s deployment --all -n $NAMESPACE
    
    # Get pod status
    echo
    print_status "Pod status:"
    kubectl get pods -n $NAMESPACE
    
    # Get service status
    echo
    print_status "Service status:"
    kubectl get services -n $NAMESPACE
    
    # Get ingress status if enabled
    if kubectl get ingress -n $NAMESPACE &> /dev/null; then
        echo
        print_status "Ingress status:"
        kubectl get ingress -n $NAMESPACE
    fi
}

# Function to display access information
display_access_info() {
    print_success "Deployment completed successfully!"
    echo
    print_status "Access Information:"
    
    # Get API Gateway service
    GATEWAY_SERVICE=$(kubectl get service -n $NAMESPACE -l app.kubernetes.io/component=api-gateway -o jsonpath='{.items[0].metadata.name}')
    
    if [[ -n "$GATEWAY_SERVICE" ]]; then
        # Check if it's a LoadBalancer service
        SERVICE_TYPE=$(kubectl get service $GATEWAY_SERVICE -n $NAMESPACE -o jsonpath='{.spec.type}')
        
        if [[ "$SERVICE_TYPE" == "LoadBalancer" ]]; then
            EXTERNAL_IP=$(kubectl get service $GATEWAY_SERVICE -n $NAMESPACE -o jsonpath='{.status.loadBalancer.ingress[0].ip}')
            if [[ -n "$EXTERNAL_IP" ]]; then
                echo "  API Gateway: http://$EXTERNAL_IP:8080"
            else
                print_warning "LoadBalancer external IP is pending. Use port-forward for now."
                echo "  Run: kubectl port-forward -n $NAMESPACE svc/$GATEWAY_SERVICE 8080:8080"
            fi
        elif [[ "$SERVICE_TYPE" == "NodePort" ]]; then
            NODE_PORT=$(kubectl get service $GATEWAY_SERVICE -n $NAMESPACE -o jsonpath='{.spec.ports[0].nodePort}')
            NODE_IP=$(kubectl get nodes -o jsonpath='{.items[0].status.addresses[?(@.type=="ExternalIP")].address}')
            if [[ -z "$NODE_IP" ]]; then
                NODE_IP=$(kubectl get nodes -o jsonpath='{.items[0].status.addresses[?(@.type=="InternalIP")].address}')
            fi
            echo "  API Gateway: http://$NODE_IP:$NODE_PORT"
        else
            echo "  Run: kubectl port-forward -n $NAMESPACE svc/$GATEWAY_SERVICE 8080:8080"
            echo "  Then access: http://localhost:8080"
        fi
    fi
    
    # Check for ingress
    INGRESS_HOST=$(kubectl get ingress -n $NAMESPACE -o jsonpath='{.items[0].spec.rules[0].host}' 2>/dev/null || echo "")
    if [[ -n "$INGRESS_HOST" ]]; then
        echo "  Ingress: http://$INGRESS_HOST"
    fi
    
    echo
    print_status "Useful commands:"
    echo "  View logs: kubectl logs -f deployment/$RELEASE_NAME-api-gateway -n $NAMESPACE"
    echo "  Scale services: kubectl scale deployment/$RELEASE_NAME-api-gateway --replicas=3 -n $NAMESPACE"
    echo "  Delete deployment: helm uninstall $RELEASE_NAME -n $NAMESPACE"
}

# Function to clean up deployment
cleanup() {
    print_status "Cleaning up deployment..."
    helm uninstall $RELEASE_NAME -n $NAMESPACE 2>/dev/null || true
    kubectl delete namespace $NAMESPACE --ignore-not-found=true
    print_success "Cleanup completed"
}

# Main script
main() {
    echo "School Management System - Kubernetes Deployment"
    echo "================================================"
    
    # Parse command line arguments
    case "${1:-deploy}" in
        "build")
            check_command docker
            build_images
            ;;
        "deploy")
            check_command kubectl
            check_command helm
            check_command docker
            check_cluster
            build_images
            load_images_to_kind
            deploy_helm
            check_deployment
            display_access_info
            ;;
        "status")
            check_command kubectl
            check_cluster
            check_deployment
            display_access_info
            ;;
        "cleanup")
            check_command kubectl
            check_command helm
            check_cluster
            cleanup
            ;;
        "help"|"-h"|"--help")
            echo "Usage: $0 [command]"
            echo
            echo "Commands:"
            echo "  build    - Build Docker images only"
            echo "  deploy   - Build images and deploy to Kubernetes (default)"
            echo "  status   - Check deployment status and display access info"
            echo "  cleanup  - Remove the deployment and namespace"
            echo "  help     - Show this help message"
            echo
            echo "Examples:"
            echo "  $0                 # Deploy the application"
            echo "  $0 build           # Build Docker images only"
            echo "  $0 status          # Check deployment status"
            echo "  $0 cleanup         # Clean up deployment"
            ;;
        *)
            print_error "Unknown command: $1"
            echo "Use '$0 help' for usage information."
            exit 1
            ;;
    esac
}

# Run main function with all arguments
main "$@"
