# 🧪 School Microservice Testing Guide

This guide provides comprehensive instructions for testing the School Microservice system, from basic compilation verification to full integration testing.

## 📚 Table of Contents

- [Quick Start Testing](#quick-start-testing)
- [Testing Scripts Overview](#testing-scripts-overview)
- [Manual Testing Procedures](#manual-testing-procedures)
- [API Testing Examples](#api-testing-examples)
- [Troubleshooting Testing Issues](#troubleshooting-testing-issues)
- [Continuous Integration Testing](#continuous-integration-testing)

## 🚀 Quick Start Testing

### Option 1: Automated Full Test (Recommended)

**Windows (PowerShell):**
```powershell
# Navigate to project directory
cd "d:\DEMO2\github-copilot-go-demos\demo9_school_microservice"

# Run comprehensive test suite
Set-ExecutionPolicy -ExecutionPolicy Bypass -Scope Process
.\test-services.ps1
```

**Linux/Mac (Bash):**
```bash
# Navigate to project directory  
cd /path/to/demo9_school_microservice

# Make executable and run
chmod +x test-services.sh
./test-services.sh
```

### Option 2: Quick Compilation Test Only

**Windows:**
```powershell
.\test-services.ps1 -SkipDocker -QuickTest
```

**Linux/Mac:**
```bash
./test-services.sh --skip-docker --quick
```

## 🛠️ Testing Scripts Overview

### Main Testing Scripts

| Script | Purpose | Platform | Features |
|--------|---------|----------|----------|
| `test-services.ps1` | Comprehensive testing | Windows | Full integration testing with Docker |
| `test-services.sh` | Comprehensive testing | Linux/Mac | Full integration testing with Docker |
| `basic-test.ps1` | Simple compilation test | Windows | Quick verification without Docker |

### Script Parameters

**Windows PowerShell (`test-services.ps1`):**
```powershell
# Skip Docker and integration tests
.\test-services.ps1 -SkipDocker

# Quick test mode (skip detailed API tests)
.\test-services.ps1 -QuickTest

# Test specific service only
.\test-services.ps1 -ServiceName "students"

# Don't cleanup Docker containers after testing
.\test-services.ps1 -NoCleanup
```

**Linux/Mac Bash (`test-services.sh`):**
```bash
# Skip Docker and integration tests
./test-services.sh --skip-docker

# Quick test mode
./test-services.sh --quick

# Test specific service
./test-services.sh --service students

# Don't cleanup after testing
./test-services.sh --no-cleanup
```

## 📋 Manual Testing Procedures

### 1. Prerequisites Verification

**Check Required Software:**
```bash
# Verify Go installation
go version
# Expected: go version go1.21+ 

# Verify Docker installation
docker --version
docker-compose --version

# Check Docker daemon status
docker ps
```

**Verify Project Structure:**
```bash
# Required files should exist
ls -la docker-compose.yml
ls -la .env
ls -la go.mod
ls -la services/
```

### 2. Compilation Testing

**Download Dependencies:**
```bash
go mod tidy
go mod download
```

**Compile Individual Services:**
```bash
# Students Service
cd services/students
go build -v .
cd ../..

# Repeat for all services: teachers, classes, academics, achievements
```

**Build All Services (Script):**
```bash
# Windows
for %s in (students teachers classes academics achievements) do (
    cd services\%s && go build -o ..\..\%s.exe . && cd ..\..
)

# Linux/Mac
for service in students teachers classes academics achievements; do
    cd services/$service && go build -o ../../$service . && cd ../..
done
```

### 3. Database Setup Testing

**Start Couchbase Container:**
```bash
docker-compose up -d couchbase
```

**Wait for Initialization:**
```bash
# Wait 60 seconds for Couchbase to be ready
sleep 60  # Linux/Mac
Start-Sleep 60  # PowerShell
```

**Test Couchbase Connectivity:**
```bash
# Check web interface
curl http://localhost:8091

# PowerShell alternative
Invoke-WebRequest http://localhost:8091
```

**Initialize Database:**
```bash
# Windows
scripts\setup-couchbase.bat

# Linux/Mac
chmod +x scripts/setup-couchbase.sh
./scripts/setup-couchbase.sh
```

### 4. Service Testing

**Start All Services:**
```bash
docker-compose up -d
```

**Check Service Status:**
```bash
docker-compose ps
docker-compose logs
```

**Test Health Endpoints:**
```bash
# Test all health endpoints
curl http://localhost:8081/health  # Students
curl http://localhost:8082/health  # Teachers
curl http://localhost:8083/health  # Classes
curl http://localhost:8084/health  # Academics
curl http://localhost:8085/health  # Achievements
```

## 🔧 API Testing Examples

### Students Service Testing

#### Create Student
**Bash/curl:**
```bash
curl -X POST http://localhost:8081/students \
  -H "Content-Type: application/json" \
  -d '{
    "firstName": "John",
    "lastName": "Doe",
    "email": "john.doe@test.edu",
    "grade": "10",
    "dateOfBirth": "2008-05-15T00:00:00Z",
    "address": "123 Test St",
    "phone": "555-1234",
    "status": "active"
  }'
```

**PowerShell:**
```powershell
$studentData = @{
    firstName = "John"
    lastName = "Doe"
    email = "john.doe@test.edu"
    grade = "10"
    dateOfBirth = "2008-05-15T00:00:00Z"
    address = "123 Test St"
    phone = "555-1234"
    status = "active"
} | ConvertTo-Json

$response = Invoke-RestMethod -Uri http://localhost:8081/students -Method POST -Body $studentData -ContentType "application/json"
Write-Host "Created student with ID: $($response.id)"
```

#### Retrieve Students
```bash
# Get all students
curl http://localhost:8081/students

# Get specific student (replace {id} with actual ID)
curl http://localhost:8081/students/{id}
```

#### Update Student
```bash
curl -X PUT http://localhost:8081/students/{id} \
  -H "Content-Type: application/json" \
  -d '{
    "firstName": "John",
    "lastName": "Doe",
    "email": "john.doe@test.edu",
    "grade": "11",
    "address": "123 Test St",
    "phone": "555-1234",
    "status": "active"
  }'
```

#### Delete Student
```bash
curl -X DELETE http://localhost:8081/students/{id}
```

### Teachers Service Testing

#### Create Teacher
```bash
curl -X POST http://localhost:8082/teachers \
  -H "Content-Type: application/json" \
  -d '{
    "firstName": "Jane",
    "lastName": "Smith",
    "email": "jane.smith@test.edu",
    "department": "Mathematics",
    "subject": "Algebra",
    "phone": "555-5678",
    "address": "456 Teacher Ave",
    "status": "active"
  }'
```

### Classes Service Testing

#### Create Class
```bash
curl -X POST http://localhost:8083/classes \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Math 101",
    "subject": "Mathematics",
    "grade": "10",
    "capacity": 30,
    "schedule": "MWF 09:00-10:00",
    "classroom": "Room 101",
    "status": "active"
  }'
```

### Load Testing

#### Simple Load Test (Bash)
```bash
# Test 10 concurrent requests
for i in {1..10}; do
    curl http://localhost:8081/students &
done
wait
echo "Load test completed"
```

#### PowerShell Load Test
```powershell
# Test 10 concurrent requests
$jobs = @()
for ($i = 1; $i -le 10; $i++) {
    $jobs += Start-Job -ScriptBlock {
        Invoke-RestMethod http://localhost:8081/students
    }
}
$jobs | Wait-Job | Receive-Job
$jobs | Remove-Job
Write-Host "Load test completed"
```

## 🐛 Troubleshooting Testing Issues

### Common Issues and Solutions

#### 1. Compilation Errors

**Issue: Missing dependencies**
```bash
# Solution: Clean and rebuild
go clean -modcache
go mod tidy
go mod download
```

**Issue: Go version too old**
```bash
# Check version (need 1.21+)
go version

# Update Go if needed
# Download from https://golang.org/dl/
```

#### 2. Docker Issues

**Issue: Port conflicts**
```bash
# Check what's using ports
netstat -ano | findstr :8081  # Windows
lsof -i :8081                 # Linux/Mac

# Kill conflicting processes
taskkill /PID <pid> /F        # Windows  
kill -9 <pid>                 # Linux/Mac
```

**Issue: Docker daemon not running**
```bash
# Start Docker Desktop
# Or restart Docker service (Linux)
sudo systemctl restart docker
```

#### 3. Database Connection Issues

**Issue: Couchbase not responding**
```bash
# Check Couchbase logs
docker-compose logs couchbase

# Restart Couchbase
docker-compose restart couchbase

# Wait longer for initialization
sleep 120
```

**Issue: Authentication failures**
```bash
# Verify credentials in .env file
cat .env | grep COUCHBASE

# Re-initialize if needed
docker exec couchbase couchbase-cli cluster-init \
  --cluster couchbase://localhost \
  --cluster-username Administrator \
  --cluster-password password \
  --cluster-name school-cluster \
  --cluster-ramsize 1024 \
  --services data,index,query
```

#### 4. API Testing Issues

**Issue: JSON parsing errors**
- Ensure proper field names (firstName/lastName, not name)
- Use correct date format: "2008-05-15T00:00:00Z"
- Check Content-Type header: "application/json"

**Issue: Service not responding**
```bash
# Check service logs
docker-compose logs students-service

# Restart specific service
docker-compose restart students-service

# Check service health
curl http://localhost:8081/health
```

### Diagnostic Commands

**Full System Status:**
```bash
# Check all containers
docker-compose ps

# Check resource usage
docker stats

# Check logs for all services
docker-compose logs --tail=50

# Check network connectivity
docker network ls
docker network inspect school-network
```

**Service-Specific Diagnostics:**
```bash
# Check specific service
docker-compose logs students-service

# Execute commands in container
docker exec -it students-service /bin/sh

# Test internal connectivity
docker exec students-service ping couchbase
```

## 🔄 Continuous Integration Testing

### GitHub Actions Example

Create `.github/workflows/test.yml`:
```yaml
name: Test School Microservice

on: [push, pull_request]

jobs:
  test:
    runs-on: ubuntu-latest
    
    steps:
    - uses: actions/checkout@v3
    
    - name: Set up Go
      uses: actions/setup-go@v3
      with:
        go-version: 1.21
    
    - name: Test compilation
      run: |
        go mod tidy
        ./test-services.sh --skip-docker --quick
    
    - name: Test with Docker
      run: |
        ./test-services.sh --no-cleanup
    
    - name: Upload test results
      uses: actions/upload-artifact@v3
      if: always()
      with:
        name: test-results
        path: test-results/
```

### Local CI Testing

**Run CI-style tests locally:**
```bash
# Quick validation (no Docker)
./test-services.sh --skip-docker --quick

# Full integration test
./test-services.sh

# Generate test report
./test-services.sh --generate-report
```

## 📊 Test Coverage and Metrics

### Unit Test Coverage
```bash
# Run with coverage
cd services/students
go test -v -cover ./...

# Generate coverage report
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html
```

### Performance Benchmarks
```bash
# Run benchmark tests
cd services/students
go test -bench=. -benchmem ./...
```

### API Response Time Testing
```bash
# Test response times
time curl http://localhost:8081/students

# Detailed timing with curl
curl -w "@curl-format.txt" -o /dev/null -s http://localhost:8081/students
```

## 🎯 Testing Best Practices

### Test Categories

1. **Unit Tests** - Test individual components in isolation
2. **Integration Tests** - Test service interactions
3. **Contract Tests** - Test API contracts and data formats
4. **Load Tests** - Test performance under load
5. **End-to-End Tests** - Test complete user workflows

### Test Data Management

- Use consistent test data across environments
- Clean up test data after tests complete
- Use separate test databases when possible
- Document test data requirements

### Environment Isolation

- Use Docker for consistent test environments
- Test in environments that match production
- Isolate tests from each other
- Use environment variables for configuration

This comprehensive testing guide ensures reliable deployment and operation of the School Microservice system! 🚀
