#!/bin/bash

# Service Health Check and Startup Script
# This script waits for all services to become healthy before proceeding

set -e

# Configuration
TIMEOUT=300
CHECK_INTERVAL=5
MAX_RETRIES=60

# Service URLs
GATEWAY_URL=${GATEWAY_URL:-"http://localhost:8080"}
STUDENT_SERVICE_URL=${STUDENT_SERVICE_URL:-"http://localhost:8081"}
TEACHER_SERVICE_URL=${TEACHER_SERVICE_URL:-"http://localhost:8082"}
ACADEMIC_SERVICE_URL=${ACADEMIC_SERVICE_URL:-"http://localhost:8083"}
ACHIEVEMENT_SERVICE_URL=${ACHIEVEMENT_SERVICE_URL:-"http://localhost:8084"}
COUCHBASE_URL=${COUCHBASE_URL:-"http://localhost:8091"}

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Logging functions
log_info() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

log_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

log_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

log_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Check if a service is healthy
check_service_health() {
    local service_name="$1"
    local service_url="$2"
    local endpoint="${3:-/health}"
    
    if curl -s -f "$service_url$endpoint" > /dev/null 2>&1; then
        return 0
    else
        return 1
    fi
}

# Wait for a service to become healthy
wait_for_service() {
    local service_name="$1"
    local service_url="$2"
    local endpoint="${3:-/health}"
    local retries=0
    
    log_info "Waiting for $service_name to become healthy..."
    
    while [ $retries -lt $MAX_RETRIES ]; do
        if check_service_health "$service_name" "$service_url" "$endpoint"; then
            log_success "$service_name is healthy!"
            return 0
        fi
        
        retries=$((retries + 1))
        log_info "Attempt $retries/$MAX_RETRIES: $service_name not ready yet..."
        sleep $CHECK_INTERVAL
    done
    
    log_error "$service_name failed to become healthy within $TIMEOUT seconds"
    return 1
}

# Check if Docker Compose services are running
check_docker_compose() {
    if ! command -v docker-compose &> /dev/null && ! command -v docker &> /dev/null; then
        log_error "Docker or Docker Compose not found"
        return 1
    fi
    
    log_info "Checking Docker Compose services status..."
    
    if command -v docker-compose &> /dev/null; then
        docker-compose ps
    else
        docker compose ps
    fi
}

# Wait for all services to be healthy
wait_for_all_services() {
    log_info "Starting health checks for all services..."
    
    # Array of services with their URLs and health endpoints
    declare -a services=(
        "Couchbase:$COUCHBASE_URL:/pools"
        "API Gateway:$GATEWAY_URL:/health"
        "Student Service:$STUDENT_SERVICE_URL:/health"
        "Teacher Service:$TEACHER_SERVICE_URL:/health"
        "Academic Service:$ACADEMIC_SERVICE_URL:/health"
        "Achievement Service:$ACHIEVEMENT_SERVICE_URL:/health"
    )
    
    local failed_services=()
    
    for service_info in "${services[@]}"; do
        IFS=':' read -r service_name service_url endpoint <<< "$service_info"
        
        if ! wait_for_service "$service_name" "$service_url" "$endpoint"; then
            failed_services+=("$service_name")
        fi
    done
    
    if [ ${#failed_services[@]} -eq 0 ]; then
        log_success "All services are healthy and ready!"
        return 0
    else
        log_error "The following services failed health checks: ${failed_services[*]}"
        return 1
    fi
}

# Perform comprehensive health checks
comprehensive_health_check() {
    log_info "Performing comprehensive health checks..."
    
    # Check individual service endpoints
    log_info "Testing API Gateway endpoints..."
    if curl -s "$GATEWAY_URL/api/students" > /dev/null; then
        log_success "Students endpoint accessible"
    else
        log_warning "Students endpoint not accessible"
    fi
    
    if curl -s "$GATEWAY_URL/api/teachers" > /dev/null; then
        log_success "Teachers endpoint accessible"
    else
        log_warning "Teachers endpoint not accessible"
    fi
    
    if curl -s "$GATEWAY_URL/api/courses" > /dev/null; then
        log_success "Courses endpoint accessible"
    else
        log_warning "Courses endpoint not accessible"
    fi
    
    if curl -s "$GATEWAY_URL/api/achievements" > /dev/null; then
        log_success "Achievements endpoint accessible"
    else
        log_warning "Achievements endpoint not accessible"
    fi
    
    # Test database connectivity
    log_info "Testing Couchbase connectivity..."
    if curl -s "$COUCHBASE_URL/pools" > /dev/null; then
        log_success "Couchbase is accessible"
    else
        log_warning "Couchbase not accessible"
    fi
}

# Run basic smoke tests
run_smoke_tests() {
    log_info "Running basic smoke tests..."
    
    # Test creating a student
    log_info "Testing student creation..."
    response=$(curl -s -X POST "$GATEWAY_URL/api/students" \
        -H "Content-Type: application/json" \
        -d '{
            "firstName": "Test",
            "lastName": "Student",
            "email": "test.student@school.edu",
            "grade": "10",
            "status": "active"
        }' 2>/dev/null || echo "FAILED")
    
    if [[ "$response" != "FAILED" && "$response" != *"error"* ]]; then
        log_success "Student creation test passed"
    else
        log_warning "Student creation test failed"
    fi
    
    # Test getting all students
    log_info "Testing student retrieval..."
    if curl -s "$GATEWAY_URL/api/students" > /dev/null; then
        log_success "Student retrieval test passed"
    else
        log_warning "Student retrieval test failed"
    fi
}

# Show service information
show_service_info() {
    log_info "Service Information:"
    echo "===================="
    echo "API Gateway:      $GATEWAY_URL"
    echo "Student Service:  $STUDENT_SERVICE_URL"
    echo "Teacher Service:  $TEACHER_SERVICE_URL"
    echo "Academic Service: $ACADEMIC_SERVICE_URL"
    echo "Achievement Service: $ACHIEVEMENT_SERVICE_URL"
    echo "Couchbase:        $COUCHBASE_URL"
    echo
    echo "Web Interfaces:"
    echo "- Couchbase Console: $COUCHBASE_URL"
    echo "- API Documentation: $GATEWAY_URL/docs (if available)"
    echo
}

# Show next steps
show_next_steps() {
    echo
    log_success "All services are ready! Here are some next steps:"
    echo
    echo "1. Load sample data:"
    echo "   ./scripts/load-comprehensive-data.sh"
    echo
    echo "2. Test the API with sample commands:"
    echo "   curl $GATEWAY_URL/api/students"
    echo "   curl $GATEWAY_URL/api/teachers"
    echo
    echo "3. Access the Couchbase Console:"
    echo "   Open $COUCHBASE_URL in your browser"
    echo "   Login: Administrator / password"
    echo
    echo "4. Run comprehensive tests:"
    echo "   ./scripts/run-tests.sh"
    echo
    echo "5. Check the documentation:"
    echo "   - API Commands: ./scripts/sample-curl-commands.md"
    echo "   - Docker Commands: ./scripts/docker-compose-commands.md"
    echo "   - Quick Start Guide: ./QUICK_START_GUIDE.md"
    echo
}

# Main execution function
main() {
    log_info "School Management System - Service Health Checker"
    echo "================================================="
    
    # Show service information
    show_service_info
    
    # Check Docker Compose status
    check_docker_compose
    
    # Wait for all services
    if wait_for_all_services; then
        # Run comprehensive health checks
        comprehensive_health_check
        
        # Run smoke tests
        run_smoke_tests
        
        # Show next steps
        show_next_steps
        
        log_success "Health check completed successfully!"
        exit 0
    else
        log_error "Health check failed. Please check the service logs:"
        echo "  docker-compose logs"
        echo "  docker-compose logs [service-name]"
        exit 1
    fi
}

# Command line options
case "${1:-main}" in
    "check")
        wait_for_all_services
        ;;
    "smoke")
        run_smoke_tests
        ;;
    "info")
        show_service_info
        ;;
    "help"|"-h"|"--help")
        echo "Usage: $0 [command]"
        echo
        echo "Commands:"
        echo "  main    - Run full health check and setup (default)"
        echo "  check   - Only check service health"
        echo "  smoke   - Run smoke tests only"
        echo "  info    - Show service information"
        echo "  help    - Show this help message"
        ;;
    *)
        main
        ;;
esac
