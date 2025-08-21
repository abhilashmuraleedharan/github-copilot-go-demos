# School Microservice - Deployment & Testing Guide

This comprehensive guide provides step-by-step instructions for deploying, testing, and managing the School Microservice system.

## 📋 Prerequisites

### Required Software
1. **Docker Desktop** - Ensure Docker Desktop is installed and running
2. **Docker Compose** - Should be included with Docker Desktop  
3. **Go 1.21+** - For local development and testing
4. **PowerShell/Bash** - For running test scripts
5. **Git** (optional) - If cloning from repository

### System Requirements
- **Memory**: Minimum 4GB RAM (8GB recommended)
- **Disk Space**: 2GB free space for Docker images
- **Ports**: 8081-8085, 8091-8096 must be available
- **Network**: Internet access for downloading dependencies

## 🏗️ Architecture Overview

The School Microservice system implements a domain-oriented microservices architecture:

### 🔧 Services Architecture
- **5 Microservices**: Students, Teachers, Classes, Academics, Achievements
- **1 Database**: Couchbase Community Edition 7.2.0
- **Service Ports**: 8081-8085 for microservices
- **Database Ports**: 8091-8096 for Couchbase cluster
- **Container Network**: school-network for inter-service communication

### 📊 Service Responsibilities
| Service | Port | Responsibility |
|---------|------|---------------|
| **Students** | 8081 | Student enrollment, personal data, status management |
| **Teachers** | 8082 | Teacher profiles, department assignments, contact info |
| **Classes** | 8083 | Class schedules, capacity, teacher assignments |
| **Academics** | 8084 | Grades, subjects, academic records, GPA calculation |
| **Achievements** | 8085 | Awards, recognition, points system |

### 🗄️ Database Design
- **Single Couchbase Bucket**: `school` (with document type separation)
- **Document Types**: students, teachers, classes, academics, achievements  
- **Indexing**: Primary and secondary indexes for performance
- **Replication**: Single replica for development (configurable for production)

## 🚀 Quick Start Deployment

### Option 1: Automated Testing Script (Recommended)

**For Windows (PowerShell):**
```powershell
# Navigate to project directory
cd "d:\DEMO2\github-copilot-go-demos\demo9_school_microservice"

# Run comprehensive testing script
Set-ExecutionPolicy -ExecutionPolicy Bypass -Scope Process
.\test-services.ps1
```

**For Linux/Mac (Bash):**
```bash
# Navigate to project directory
cd /path/to/demo9_school_microservice

# Make script executable and run
chmod +x test-services.sh
./test-services.sh
```

### Option 2: Manual Step-by-Step Deployment

## 📝 Step-by-Step Deployment Instructions

### Step 1: Project Setup

**Navigate to Project Directory:**
```bash
cd /path/to/demo9_school_microservice
```

**Verify Project Structure:**
```bash
# Check that all required files exist
ls -la
# Should see: docker-compose.yml, .env, services/, scripts/, etc.
```

### Step 2: Environment Configuration

**Copy Environment File:**
```bash
cp .env.example .env
```

**Edit Environment Variables (Optional):**
```bash
# Default configuration works for local development
# Couchbase Configuration
COUCHBASE_HOST=localhost
COUCHBASE_USERNAME=Administrator
COUCHBASE_PASSWORD=password
COUCHBASE_BUCKET=school

# Service Ports  
STUDENTS_PORT=8081
TEACHERS_PORT=8082
CLASSES_PORT=8083
ACADEMICS_PORT=8084
ACHIEVEMENTS_PORT=8085
```

### Step 3: Compile Services (Optional - for validation)

**Install Go Dependencies:**
```bash
go mod tidy
```

**Compile All Services:**
```bash
# Students Service
cd services/students && go build -o ../../students.exe . && cd ../..

# Teachers Service  
cd services/teachers && go build -o ../../teachers.exe . && cd ../..

# Classes Service
cd services/classes && go build -o ../../classes.exe . && cd ../..

# Academics Service
cd services/academics && go build -o ../../academics.exe . && cd ../..

# Achievements Service
cd services/achievements && go build -o ../../achievements.exe . && cd ../..
```

### Step 4: Start Database First

**Start Couchbase Container:**
```bash
docker-compose up -d couchbase
```

**Wait for Database Initialization:**
```bash
# Wait 30-60 seconds for Couchbase to fully start
sleep 60
```

**Initialize Couchbase Cluster and Bucket:**

**Windows:**
```batch
scripts\setup-couchbase.bat
```

**Linux/Mac:**
```bash
chmod +x scripts/setup-couchbase.sh
./scripts/setup-couchbase.sh
```

### Step 5: Start All Services

**Start Complete System:**
```bash
docker-compose up -d
```

**Monitor Service Startup:**
```bash
# Watch all service logs
docker-compose logs -f

# Check container status
docker-compose ps
```

## ✅ Service Verification & Testing

### Health Check Verification

**Quick Health Check (PowerShell):**
```powershell
# Test all service health endpoints
$services = @(
    @{Name="Students"; Port=8081},
    @{Name="Teachers"; Port=8082}, 
    @{Name="Classes"; Port=8083},
    @{Name="Academics"; Port=8084},
    @{Name="Achievements"; Port=8085}
)

foreach ($service in $services) {
    try {
        $response = Invoke-WebRequest -Uri "http://localhost:$($service.Port)/health" -TimeoutSec 5
        Write-Host "✅ $($service.Name) Service: $($response.StatusCode)" -ForegroundColor Green
    } catch {
        Write-Host "❌ $($service.Name) Service: Failed" -ForegroundColor Red
    }
}
```

**Quick Health Check (Bash):**
```bash
#!/bin/bash
services=("students:8081" "teachers:8082" "classes:8083" "academics:8084" "achievements:8085")

for service in "${services[@]}"; do
    name=$(echo $service | cut -d: -f1)
    port=$(echo $service | cut -d: -f2)
    
    if curl -s -f "http://localhost:$port/health" > /dev/null; then
        echo "✅ $name Service: OK"
    else
        echo "❌ $name Service: Failed"
    fi
done
```

### Database Connectivity Test

**Test Couchbase Access:**
```bash
# Check Couchbase Web Console
curl -f http://localhost:8091

# Test bucket access (requires authentication)
curl -u Administrator:password http://localhost:8091/pools/default/buckets/school
```

## 🧪 Comprehensive Testing Suite

### Automated Testing Scripts

The project includes comprehensive testing scripts that verify:
- ✅ Code compilation and dependencies
- ✅ Service health and connectivity  
- ✅ Database initialization and access
- ✅ CRUD operations functionality
- ✅ JSON serialization/deserialization
- ✅ Error handling and edge cases

**Run Complete Test Suite:**

**Windows:**
```powershell
.\test-services.ps1
```

**Linux/Mac:**
```bash
./test-services.sh
```

### Manual API Testing

#### Create Test Data

**Create a Student:**
```bash
curl -X POST http://localhost:8081/students \
  -H "Content-Type: application/json" \
  -d '{
    "firstName": "John",
    "lastName": "Doe", 
    "email": "john.doe@school.edu",
    "grade": "10",
    "dateOfBirth": "2008-05-15T00:00:00Z",
    "address": "123 Main St",
    "phone": "555-1234",
    "status": "active"
  }'
```

**PowerShell Alternative:**
```powershell
$studentData = @{
    firstName = "John"
    lastName = "Doe"
    email = "john.doe@school.edu"  
    grade = "10"
    dateOfBirth = "2008-05-15T00:00:00Z"
    address = "123 Main St"
    phone = "555-1234"
    status = "active"
} | ConvertTo-Json

Invoke-RestMethod -Uri http://localhost:8081/students -Method POST -Body $studentData -ContentType "application/json"
```

**Create a Teacher:**
```bash
curl -X POST http://localhost:8082/teachers \
  -H "Content-Type: application/json" \
  -d '{
    "firstName": "Jane",
    "lastName": "Smith",
    "email": "jane.smith@school.edu", 
    "department": "Mathematics",
    "subject": "Algebra",
    "phone": "555-5678",
    "address": "456 Oak Ave",
    "status": "active"
  }'
```

**Create a Class:**
```bash
curl -X POST http://localhost:8083/classes \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Mathematics 10A",
    "subject": "Mathematics",
    "grade": "10", 
    "capacity": 30,
    "schedule": "MWF 09:00-10:00",
    "classroom": "Room 101",
    "status": "active"
  }'
```

#### Retrieve Data

**Get All Students:**
```bash
curl http://localhost:8081/students
```

**Get Student by ID:**
```bash
# Replace {student-id} with actual ID from creation response
curl http://localhost:8081/students/{student-id}
```

**Get All Teachers:**
```bash
curl http://localhost:8082/teachers
```

**Get All Classes:**
```bash
curl http://localhost:8083/classes
```

#### Update Operations

**Update Student:**
```bash
curl -X PUT http://localhost:8081/students/{student-id} \
  -H "Content-Type: application/json" \
  -d '{
    "firstName": "John",
    "lastName": "Doe",
    "email": "john.doe@school.edu",
    "grade": "11",
    "address": "123 Main St",
    "phone": "555-1234", 
    "status": "active"
  }'
```

#### Delete Operations

**Delete Student:**
```bash
curl -X DELETE http://localhost:8081/students/{student-id}
```

### Performance Testing

**Load Testing with curl:**
```bash
# Test concurrent requests
for i in {1..10}; do
    curl http://localhost:8081/students &
done
wait
```

**PowerShell Load Test:**
```powershell
# Test concurrent requests
$jobs = @()
for ($i = 1; $i -le 10; $i++) {
    $jobs += Start-Job -ScriptBlock {
        Invoke-RestMethod http://localhost:8081/students
    }
}
$jobs | Wait-Job | Receive-Job
```

## 📖 Service API Reference

### Students Service (Port 8081)
- `GET /students` - List all students
- `GET /students/{id}` - Get student by ID
- `POST /students` - Create new student
- `PUT /students/{id}` - Update student
- `DELETE /students/{id}` - Delete student
- `GET /health` - Health check

### Teachers Service (Port 8082)
- `GET /teachers` - List all teachers
- `GET /teachers/{id}` - Get teacher by ID
- `POST /teachers` - Create new teacher
- `PUT /teachers/{id}` - Update teacher
- `DELETE /teachers/{id}` - Delete teacher
- `GET /health` - Health check

### Classes Service (Port 8083)
- `GET /classes` - List all classes
- `GET /classes/{id}` - Get class by ID
- `POST /classes` - Create new class
- `PUT /classes/{id}` - Update class
- `DELETE /classes/{id}` - Delete class
- `GET /health` - Health check

### Academics Service (Port 8084)
- `GET /academics` - List all academic records
- `GET /academics/{id}` - Get academic record by ID
- `GET /academics/student/{studentId}` - Get academics by student
- `POST /academics` - Create new academic record
- `PUT /academics/{id}` - Update academic record
- `DELETE /academics/{id}` - Delete academic record
- `GET /health` - Health check

### Achievements Service (Port 8085)
- `GET /achievements` - List all achievements
- `GET /achievements/{id}` - Get achievement by ID
- `GET /achievements/student/{studentId}` - Get achievements by student
- `POST /achievements` - Create new achievement
- `PUT /achievements/{id}` - Update achievement
- `DELETE /achievements/{id}` - Delete achievement
- `GET /health` - Health check

## 🌐 Accessing Web Interfaces

### Couchbase Web Console
- **URL:** http://localhost:8091
- **Username:** Administrator
- **Password:** password
- **Features:** Database management, query tools, monitoring dashboards

### Service Health Dashboards
- **Students Service:** http://localhost:8081/health
- **Teachers Service:** http://localhost:8082/health
- **Classes Service:** http://localhost:8083/health
- **Academics Service:** http://localhost:8084/health
- **Achievements Service:** http://localhost:8085/health

## 🛑 Service Management

### Stopping Services

**Stop All Services:**
```bash
docker-compose down
```

**Stop and Remove Volumes (⚠️ Deletes all data):**
```bash
docker-compose down -v
```

**Stop Individual Service:**
```bash
docker-compose stop students-service
```

### Starting Services

**Start All Services:**
```bash
docker-compose up -d
```

**Start Individual Service:**
```bash
docker-compose start students-service
```

**Restart Service:**
```bash
docker-compose restart students-service
```

## 🔧 Troubleshooting Guide

### Common Issues and Solutions

#### 1. Port Binding Errors
**Problem:** `Port already in use` errors
**Solution:**
```bash
# Check what's using the ports
netstat -ano | findstr :8081  # Windows
lsof -i :8081                 # Linux/Mac

# Kill processes using the ports
taskkill /PID <process-id> /F  # Windows
kill -9 <process-id>          # Linux/Mac
```

#### 2. Couchbase Health Check Failures
**Problem:** Couchbase container shows as unhealthy
**Solutions:**
```bash
# Check Couchbase logs
docker-compose logs couchbase

# Wait for initialization (can take 60+ seconds)
sleep 60

# Manually initialize if needed
docker exec couchbase couchbase-cli cluster-init --cluster couchbase://localhost --cluster-username Administrator --cluster-password password --cluster-name school-cluster --cluster-ramsize 1024 --cluster-index-ramsize 512 --services data,index,query
```

#### 3. Service Connection Errors
**Problem:** Services can't connect to Couchbase
**Solutions:**
```bash
# Check network connectivity
docker network ls
docker network inspect school-network

# Restart services after Couchbase is ready
docker-compose restart students-service
```

#### 4. Build/Compilation Errors
**Problem:** Go build failures
**Solutions:**
```bash
# Clean and rebuild modules
go clean -modcache
go mod tidy
go mod download

# Check Go version (requires 1.21+)
go version
```

#### 5. JSON Parsing Errors
**Problem:** Invalid JSON in API requests
**Solution:** Ensure proper field names and data types:
```json
{
  "firstName": "John",      // ✅ Correct
  "lastName": "Doe",        // ✅ Correct  
  "name": "John Doe",       // ❌ Wrong - use firstName/lastName
  "dateOfBirth": "2008-05-15T00:00:00Z"  // ✅ Correct ISO format
}
```

### Diagnostic Commands

**Check All Container Status:**
```bash
docker-compose ps
```

**View Service Logs:**
```bash
# All services
docker-compose logs

# Specific service with follow
docker-compose logs -f students-service

# Last 50 lines
docker-compose logs --tail=50 couchbase
```

**Check Resource Usage:**
```bash
docker stats
```

**Test Network Connectivity:**
```bash
# From host to services
curl http://localhost:8081/health

# Between containers (exec into a container first)
docker exec -it students-service ping couchbase
```

## 📊 Monitoring and Performance

### Health Monitoring Scripts

**Continuous Health Check (PowerShell):**
```powershell
while ($true) {
    $timestamp = Get-Date -Format "yyyy-MM-dd HH:mm:ss"
    Write-Host "[$timestamp] Checking services..." -ForegroundColor Yellow
    
    $services = @(8081, 8082, 8083, 8084, 8085)
    foreach ($port in $services) {
        try {
            $response = Invoke-WebRequest "http://localhost:$port/health" -TimeoutSec 3
            Write-Host "  Port $port`: OK" -ForegroundColor Green
        } catch {
            Write-Host "  Port $port`: FAILED" -ForegroundColor Red
        }
    }
    Start-Sleep 30
}
```

**Resource Monitoring:**
```bash
# Monitor Docker container resources
watch docker stats

# Monitor host system resources
htop  # Linux
Get-Process | Sort-Object CPU -Descending | Select-Object -First 10  # PowerShell
```

### Performance Tuning

**Couchbase Optimization:**
- Increase RAM allocation for production: `--cluster-ramsize 4096`
- Enable auto-compaction for better performance
- Configure appropriate bucket replicas for HA

**Service Optimization:**
- Adjust Go runtime: `GOMAXPROCS=4`
- Tune garbage collection: `GOGC=100`
- Configure connection pooling for database connections

## 🚀 Development Workflow

### Local Development Setup

**Run Individual Service:**
```bash
# Set environment variables
export COUCHBASE_HOST=localhost
export COUCHBASE_USERNAME=Administrator
export COUCHBASE_PASSWORD=password
export STUDENTS_PORT=8081

# Run service
cd services/students
go run main.go
```

**Hot Reload Development:**
```bash
# Install air for hot reload
go install github.com/cosmtrek/air@latest

# Run with hot reload
cd services/students
air
```

### Testing in Development

**Unit Tests:**
```bash
cd services/students
go test ./... -v
```

**Integration Tests:**
```bash
# Start test environment
docker-compose -f docker-compose.test.yml up -d

# Run integration tests
go test ./tests/integration -v

# Cleanup
docker-compose -f docker-compose.test.yml down
```

### Code Quality

**Run Static Analysis:**
```bash
# Install tools
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Run linting
golangci-lint run ./...

# Format code
go fmt ./...

# Vet code
go vet ./...
```

## 🔒 Security Considerations

### Production Security Checklist

- [ ] Change default Couchbase credentials
- [ ] Enable TLS for all communications
- [ ] Implement API authentication/authorization
- [ ] Configure firewall rules
- [ ] Enable audit logging
- [ ] Regular security updates
- [ ] Use secrets management for credentials
- [ ] Implement rate limiting
- [ ] Enable CORS restrictions
- [ ] Use non-root container users

### Environment-Specific Configurations

**Development:**
```bash
COUCHBASE_PASSWORD=password  # Simple password OK
CORS_ALLOWED_ORIGINS=*       # Allow all origins
LOG_LEVEL=debug              # Verbose logging
```

**Production:**
```bash
COUCHBASE_PASSWORD=${SECRET_PASSWORD}  # From secrets manager
CORS_ALLOWED_ORIGINS=https://myapp.com # Specific origins only
LOG_LEVEL=info                         # Production logging
TLS_ENABLED=true                       # Enable TLS
AUTH_REQUIRED=true                     # Require authentication
```

## 📈 Scaling and Production Deployment

### Horizontal Scaling

**Scale Individual Services:**
```bash
# Scale students service to 3 replicas
docker-compose up -d --scale students-service=3

# Use load balancer for distribution
# Configure nginx/HAProxy for load balancing
```

**Kubernetes Deployment:**
```bash
# Use provided Helm charts
helm install school-microservice ./helm/school-microservice

# Scale using kubectl
kubectl scale deployment students-service --replicas=5
```

### Database Scaling

**Couchbase Cluster:**
```bash
# Add more Couchbase nodes
docker-compose -f docker-compose.prod.yml up -d --scale couchbase=3

# Configure cross datacenter replication
# Set up auto-failover and monitoring
```

This completes the comprehensive deployment and testing guide for the School Microservice system! 🚀
