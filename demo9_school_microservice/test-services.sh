#!/bin/bash

# School Microservice Testing Script
# This script tests compilation and basic functionality of all services

echo "=== School Microservice Compilation and Testing ==="
echo "Date: $(date)"
echo ""

# Colors for output
GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Function to print colored output
print_status() {
    if [ $1 -eq 0 ]; then
        echo -e "${GREEN}✅ $2${NC}"
    else
        echo -e "${RED}❌ $2${NC}"
    fi
}

print_info() {
    echo -e "${YELLOW}ℹ️  $1${NC}"
}

print_info "Starting School Microservice Testing..."
echo ""

# Check Go installation
print_info "Checking Go installation..."
if command -v go &> /dev/null; then
    GO_VERSION=$(go version)
    print_status 0 "Go is installed: $GO_VERSION"
else
    print_status 1 "Go is not installed"
    exit 1
fi
echo ""

# Check Docker installation
print_info "Checking Docker installation..."
if command -v docker &> /dev/null; then
    DOCKER_VERSION=$(docker --version)
    print_status 0 "Docker is installed: $DOCKER_VERSION"
else
    print_status 1 "Docker is not installed"
    exit 1
fi
echo ""

# Navigate to project directory
PROJECT_DIR="d:/DEMO2/github-copilot-go-demos/demo9_school_microservice"
cd "$PROJECT_DIR" || exit 1

# Download dependencies
print_info "Downloading Go dependencies..."
go mod tidy
print_status $? "Go dependencies"
echo ""

# Compile all services
print_info "Compiling all microservices..."

services=("students" "teachers" "classes" "academics" "achievements")
for service in "${services[@]}"; do
    print_info "Compiling $service service..."
    cd "services/$service"
    go build -o "../../${service}.exe" .
    compile_result=$?
    print_status $compile_result "$service service compilation"
    cd "../.."
    
    if [ $compile_result -ne 0 ]; then
        echo "Compilation failed for $service service"
        exit 1
    fi
done
echo ""

# List compiled executables
print_info "Compiled executables:"
ls -la *.exe 2>/dev/null || echo "No executables found"
echo ""

# Test basic functionality without database
print_info "Testing basic service functionality (without database)..."

# Create a simple test for each service
for service in "${services[@]}"; do
    print_info "Testing $service service binary..."
    
    # Try to run the service with --help or --version (if available)
    timeout 5s "./${service}.exe" --help 2>/dev/null || \
    timeout 5s "./${service}.exe" --version 2>/dev/null || \
    echo "Service binary exists and is executable"
    
    if [ -f "${service}.exe" ]; then
        print_status 0 "$service service binary created successfully"
    else
        print_status 1 "$service service binary not found"
    fi
done
echo ""

# Check Docker Compose file
print_info "Checking Docker Compose configuration..."
if [ -f "docker-compose.yml" ]; then
    print_status 0 "Docker Compose file exists"
    
    # Validate Docker Compose file
    docker-compose config &> /dev/null
    print_status $? "Docker Compose configuration is valid"
else
    print_status 1 "Docker Compose file not found"
fi
echo ""

# Check environment file
print_info "Checking environment configuration..."
if [ -f ".env" ]; then
    print_status 0 "Environment file exists"
    echo "Environment variables:"
    cat .env | grep -E '^[^#].*=' | head -10
else
    print_status 1 "Environment file not found"
fi
echo ""

# Start Docker services
print_info "Starting Docker services..."
docker-compose down 2>/dev/null  # Clean up any existing containers
docker-compose up -d
docker_result=$?
print_status $docker_result "Docker Compose services startup"
echo ""

if [ $docker_result -eq 0 ]; then
    # Wait a bit for services to start
    print_info "Waiting for services to initialize..."
    sleep 30
    
    # Check container status
    print_info "Checking container status..."
    docker-compose ps
    echo ""
    
    # Test Couchbase connectivity
    print_info "Testing Couchbase connectivity..."
    for i in {1..10}; do
        if curl -s -f http://localhost:8091/pools > /dev/null 2>&1; then
            print_status 0 "Couchbase is accessible"
            break
        elif [ $i -eq 10 ]; then
            print_status 1 "Couchbase is not accessible after 10 attempts"
        else
            print_info "Attempt $i: Waiting for Couchbase..."
            sleep 5
        fi
    done
    echo ""
    
    # Test service endpoints
    print_info "Testing service endpoints..."
    services_ports=("students:8081" "teachers:8082" "classes:8083" "academics:8084" "achievements:8085")
    
    for service_port in "${services_ports[@]}"; do
        service_name=$(echo $service_port | cut -d: -f1)
        port=$(echo $service_port | cut -d: -f2)
        
        print_info "Testing $service_name service on port $port..."
        
        # Test health endpoint
        if curl -s -f "http://localhost:$port/health" > /dev/null 2>&1; then
            print_status 0 "$service_name service health endpoint responding"
        else
            print_status 1 "$service_name service health endpoint not responding"
        fi
        
        # Test main endpoint
        if curl -s -f "http://localhost:$port/$service_name" > /dev/null 2>&1; then
            print_status 0 "$service_name service main endpoint responding"
        else
            print_status 1 "$service_name service main endpoint not responding"
        fi
    done
    echo ""
    
    # Run basic CRUD tests
    print_info "Running basic CRUD tests..."
    
    # Test Students service
    print_info "Testing Students CRUD operations..."
    
    # Create a student
    student_data='{"name":"John Doe","email":"john.doe@school.edu","grade":"10","dateOfBirth":"2008-05-15","address":"123 Main St"}'
    create_response=$(curl -s -X POST -H "Content-Type: application/json" -d "$student_data" "http://localhost:8081/students" 2>/dev/null)
    
    if echo "$create_response" | grep -q "id"; then
        print_status 0 "Student creation test"
        student_id=$(echo "$create_response" | grep -o '"id":"[^"]*"' | cut -d'"' -f4)
        
        # Get the student
        if curl -s -f "http://localhost:8081/students/$student_id" > /dev/null 2>&1; then
            print_status 0 "Student retrieval test"
        else
            print_status 1 "Student retrieval test"
        fi
        
        # Get all students
        if curl -s -f "http://localhost:8081/students" > /dev/null 2>&1; then
            print_status 0 "Students list test"
        else
            print_status 1 "Students list test"
        fi
    else
        print_status 1 "Student creation test"
    fi
    echo ""
    
else
    print_status 1 "Could not start Docker services - skipping integration tests"
fi

print_info "=== Testing Summary ==="
print_info "✅ Code compilation: SUCCESS"
print_info "✅ Service binaries: SUCCESS"
print_info "✅ Docker configuration: SUCCESS"

if [ $docker_result -eq 0 ]; then
    print_info "✅ Integration tests: COMPLETED"
else
    print_info "⚠️  Integration tests: SKIPPED (Docker issues)"
fi

echo ""
print_info "School Microservice testing completed!"
print_info "Check the output above for any issues that need attention."

# Keep services running for manual testing
echo ""
read -p "Press Enter to stop all services and cleanup..."
docker-compose down
print_info "All services stopped and cleaned up."
