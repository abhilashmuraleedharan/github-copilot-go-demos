# 🚀 School Management System - Quick Start Guide

## 🐳 Docker Compose Commands

### Start All Services
```bash
# Start all services in detached mode
docker-compose up -d

# Start with logs visible
docker-compose up

# Start specific services
docker-compose up -d couchbase api-gateway

# Rebuild and start (if code changes)
docker-compose up -d --build
```

### Stop Services
```bash
# Stop all services
docker-compose down

# Stop and remove volumes (⚠️ This will delete all data)
docker-compose down -v

# Stop specific service
docker-compose stop api-gateway
```

### Service Management
```bash
# View running services
docker-compose ps

# View logs
docker-compose logs -f api-gateway
docker-compose logs -f couchbase

# Restart specific service
docker-compose restart student-service

# Scale services (if needed)
docker-compose up -d --scale student-service=2
```

### Cleanup Commands
```bash
# Remove stopped containers
docker-compose rm

# Remove all containers, networks, and volumes
docker-compose down -v --rmi all

# Prune unused Docker resources
docker system prune -a
```

## 📊 Data Loading Scripts

### 1. Automatic Setup Script (Recommended)

**Run the comprehensive demo script:**

#### Linux/macOS:
```bash
# Make executable
chmod +x scripts/couchbase-demo.sh

# Run full setup and demo
./scripts/couchbase-demo.sh

# Or specific operations
./scripts/couchbase-demo.sh setup   # Initialize only
./scripts/couchbase-demo.sh demo    # CRUD demos only
./scripts/couchbase-demo.sh test    # Test connection only
```

#### Windows:
```cmd
# Run full setup and demo
scripts\couchbase-demo.ps1

# Or use the batch launcher
scripts\run-demo.bat

# PowerShell execution policy (if needed)
Set-ExecutionPolicy -ExecutionPolicy RemoteSigned -Scope CurrentUser
```

### 2. Manual Data Loading

#### Initialize Couchbase Cluster:
```bash
# Wait for Couchbase to be ready
docker-compose logs -f couchbase

# Initialize cluster
curl -X POST http://localhost:8091/pools/default \
  -d 'memoryQuota=512' \
  -d 'indexMemoryQuota=256'

# Setup administrator
curl -X POST http://localhost:8091/settings/web \
  -d 'username=Administrator' \
  -d 'password=password123' \
  -d 'port=SAME'

# Create bucket
curl -X POST http://localhost:8091/pools/default/buckets \
  -u Administrator:password123 \
  -d 'name=schoolmgmt' \
  -d 'ramQuotaMB=512' \
  -d 'bucketType=membase'
```

#### Create Collections:
```bash
# Create scope
curl -X POST http://localhost:8091/pools/default/buckets/schoolmgmt/scopes \
  -u Administrator:password123 \
  -d 'name=school'

# Create collections
curl -X POST http://localhost:8091/pools/default/buckets/schoolmgmt/scopes/school/collections \
  -u Administrator:password123 \
  -d 'name=students'

curl -X POST http://localhost:8091/pools/default/buckets/schoolmgmt/scopes/school/collections \
  -u Administrator:password123 \
  -d 'name=teachers'

curl -X POST http://localhost:8091/pools/default/buckets/schoolmgmt/scopes/school/collections \
  -u Administrator:password123 \
  -d 'name=academics'

curl -X POST http://localhost:8091/pools/default/buckets/schoolmgmt/scopes/school/collections \
  -u Administrator:password123 \
  -d 'name=achievements'
```

## 🔗 Sample CRUD Operations

### Base URLs
- **API Gateway**: http://localhost:8080
- **Direct Services**: 
  - Student: http://localhost:8081
  - Teacher: http://localhost:8082
  - Academic: http://localhost:8083
  - Achievement: http://localhost:8084

### Important: API Endpoint Paths
The API Gateway routes requests to microservices using these paths:
- **Students**: `http://localhost:8080/api/v1/students`
- **Teachers**: `http://localhost:8080/api/v1/teachers`
- **Academics**: `http://localhost:8080/api/v1/academics`
- **Classes**: `http://localhost:8080/api/v1/classes`
- **Achievements**: `http://localhost:8080/api/v1/achievements`
- **Badges**: `http://localhost:8080/api/v1/badges`
- **Health Check**: `http://localhost:8080/health`

### PowerShell Testing Commands
```powershell
# Test API Gateway health
Invoke-RestMethod -Uri "http://localhost:8080/health"

# Test students endpoint (may return empty array initially)
Invoke-RestMethod -Uri "http://localhost:8080/api/v1/students"

# Test individual services directly
Invoke-RestMethod -Uri "http://localhost:8081/health"  # Student Service
Invoke-RestMethod -Uri "http://localhost:8082/health"  # Teacher Service
```

### 👨‍🎓 Student Operations

#### Create Student
```bash
curl -X POST http://localhost:8080/api/v1/students \
  -H "Content-Type: application/json" \
  -d '{
    "id": "student-001",
    "first_name": "John",
    "last_name": "Doe",
    "email": "john.doe@school.edu",
    "grade": "10",
    "age": 16,
    "enrollment_date": "2024-08-01",
    "status": "active"
  }'
```

**PowerShell equivalent:**
```powershell
$studentData = @{
    id = "student-001"
    first_name = "John"
    last_name = "Doe"
    email = "john.doe@school.edu"
    grade = "10"
    age = 16
    enrollment_date = "2024-08-01"
    status = "active"
} | ConvertTo-Json

Invoke-RestMethod -Uri "http://localhost:8080/api/v1/students" -Method Post -Body $studentData -ContentType "application/json"
```

#### Get All Students
```bash
curl http://localhost:8080/api/v1/students
```

**PowerShell equivalent:**
```powershell
Invoke-RestMethod -Uri "http://localhost:8080/api/v1/students"
```

#### Get Specific Student
```bash
curl http://localhost:8080/api/v1/students/student-001
```

**PowerShell equivalent:**
```powershell
Invoke-RestMethod -Uri "http://localhost:8080/api/v1/students/student-001"
```

#### Update Student
```bash
curl -X PUT http://localhost:8080/api/v1/students/student-001 \
  -H "Content-Type: application/json" \
  -d '{
    "id": "student-001",
    "first_name": "John",
    "last_name": "Doe",
    "email": "john.doe@school.edu",
    "grade": "11",
    "age": 17,
    "status": "active"
  }'
```

**PowerShell equivalent:**
```powershell
$updateData = @{
    id = "student-001"
    first_name = "John"
    last_name = "Doe" 
    email = "john.doe@school.edu"
    grade = "11"
    age = 17
    status = "active"
} | ConvertTo-Json

Invoke-RestMethod -Uri "http://localhost:8080/api/v1/students/student-001" -Method Put -Body $updateData -ContentType "application/json"
```

#### Delete Student
```bash
curl -X DELETE http://localhost:8080/api/v1/students/student-001
```

**PowerShell equivalent:**
```powershell
Invoke-RestMethod -Uri "http://localhost:8080/api/v1/students/student-001" -Method Delete
```

### 👩‍🏫 Teacher Operations

#### Create Teacher
```bash
curl -X POST http://localhost:8080/api/v1/teachers \
  -H "Content-Type: application/json" \
  -d '{
    "id": "teacher-001",
    "first_name": "Dr. Emma",
    "last_name": "Wilson",
    "email": "emma.wilson@school.edu",
    "department": "Mathematics",
    "subjects": ["Algebra", "Geometry"],
    "experience": 8,
    "hire_date": "2020-01-15",
    "status": "active"
  }'
```

#### Get All Teachers
```bash
curl http://localhost:8080/api/v1/teachers
```

#### Get Teachers by Department
```bash
curl "http://localhost:8080/api/v1/teachers?department=Mathematics"
```

### 📚 Academic Operations

#### Create Academic Record
```bash
curl -X POST http://localhost:8080/api/v1/academics \
  -H "Content-Type: application/json" \
  -d '{
    "id": "academic-001",
    "student_id": "student-001",
    "teacher_id": "teacher-001",
    "subject": "Algebra",
    "grade": "A",
    "semester": "Spring 2024",
    "max_marks": 100,
    "obtained_marks": 95,
    "percentage": 95.0,
    "status": "pass"
  }'
```

#### Get Student Performance
```bash
curl "http://localhost:8080/api/v1/academics?student_id=student-001"
```

#### Create Class
```bash
curl -X POST http://localhost:8080/api/v1/classes \
  -H "Content-Type: application/json" \
  -d '{
    "id": "class-001",
    "name": "Algebra I",
    "teacher_id": "teacher-001",
    "schedule": "Mon, Wed, Fri 9:00 AM",
    "room": "101",
    "capacity": 30,
    "enrolled": 25
  }'
```

### 🏆 Achievement Operations

#### Create Achievement
```bash
curl -X POST http://localhost:8080/api/v1/achievements \
  -H "Content-Type: application/json" \
  -d '{
    "id": "achievement-001",
    "student_id": "student-001",
    "title": "Math Excellence Award",
    "description": "Outstanding performance in mathematics",
    "category": "academic",
    "points": 100,
    "date": "2024-04-20",
    "status": "active"
  }'
```

#### Get Student Achievements
```bash
curl "http://localhost:8080/api/v1/achievements?student_id=student-001"
```

#### Get Leaderboard
```bash
curl http://localhost:8080/api/v1/achievements/leaderboard
```

#### Create Badge
```bash
curl -X POST http://localhost:8080/api/v1/badges \
  -H "Content-Type: application/json" \
  -d '{
    "id": "badge-001",
    "name": "Honor Roll",
    "description": "Achieved honor roll status",
    "icon": "🎓",
    "criteria": "GPA >= 3.5",
    "active": true
  }'
```

## 🔍 Health Check Commands

### Service Health Checks
```bash
# API Gateway health
curl http://localhost:8080/health

# Individual service health
curl http://localhost:8081/health  # Student Service
curl http://localhost:8082/health  # Teacher Service
curl http://localhost:8083/health  # Academic Service
curl http://localhost:8084/health  # Achievement Service

# Couchbase health
curl http://localhost:8091/pools
```

### System Status
```bash
# Check all containers
docker-compose ps

# Check container logs
docker-compose logs api-gateway
docker-compose logs couchbase

# Check resource usage
docker stats
```

## 📋 Quick Start Workflow

### 1. Complete Setup (Recommended)
```bash
# 1. Start all services
docker-compose up -d

# 2. Wait for services to be ready (check logs)
docker-compose logs -f couchbase

# 3. Run automated setup and demo
./scripts/couchbase-demo.sh

# 4. Test the system
curl http://localhost:8080/health
curl http://localhost:8080/api/v1/students
```

### 2. Manual Setup
```bash
# 1. Start services
docker-compose up -d

# 2. Wait for Couchbase (about 60 seconds)
sleep 60

# 3. Initialize Couchbase manually (see Manual Data Loading above)

# 4. Test individual operations (see CRUD examples above)
```

### 3. Development Workflow
```bash
# 1. Make code changes

# 2. Rebuild and restart specific service
docker-compose up -d --build student-service

# 3. Test changes
curl http://localhost:8080/api/v1/students

# 4. View logs
docker-compose logs -f student-service
```

## 🎯 Recommended Setup Workflows

### Windows Users (PowerShell) - Complete Setup
```powershell
# Run our comprehensive setup script (handles all common issues)
.\scripts\complete-setup-windows.ps1 -Verbose

# This script will:
# 1. Fix PowerShell execution policy
# 2. Check Docker installation
# 3. Start services if needed
# 4. Wait for all services to be healthy
# 5. Initialize Couchbase
# 6. Test all API endpoints
# 7. Create a test student
```

### Linux/macOS Users - Complete Setup
```bash
# Make scripts executable
chmod +x scripts/*.sh

# Run comprehensive setup
./scripts/wait-for-services.sh && ./scripts/load-comprehensive-data.sh

# Or run step by step:
docker-compose up -d
./scripts/wait-for-services.sh
./scripts/load-comprehensive-data.sh
```

### Quick Health Check (Any Platform)
```bash
# Docker Compose
docker-compose ps

# API Gateway
curl http://localhost:8080/health  # Linux/macOS
# OR
Invoke-RestMethod -Uri "http://localhost:8080/health"  # PowerShell
```

## 🎉 Success Verification

After running the setup, you should be able to:
1. ✅ Access API Gateway health: http://localhost:8080/health
2. ✅ View Couchbase UI: http://localhost:8091 (Administrator/password123)
3. ✅ Create and retrieve students via API
4. ✅ See data in Couchbase collections
5. ✅ Run CRUD operations on all services

**Your School Management System is now ready for development and testing!** 🚀

## 🛠️ Troubleshooting

### Common Issues:

#### Services Not Starting:
```bash
# Check Docker status
docker --version
docker-compose --version

# Check port conflicts
netstat -an | findstr "8080"  # Windows
lsof -i :8080                 # Linux/macOS

# Restart Docker Desktop
```

#### Couchbase Connection Issues:
```bash
# Check Couchbase container logs
docker-compose logs couchbase

# Verify Couchbase is accessible
curl http://localhost:8091/pools

# Restart Couchbase
docker-compose restart couchbase
```

#### Script Permission Issues (Linux/macOS):
```bash
# Fix permissions
chmod +x scripts/couchbase-demo.sh
chmod +x k8s/deploy.sh

# Run with bash explicitly
bash scripts/couchbase-demo.sh
```

### Performance Tips:
1. **Allocate Sufficient Memory**: Docker Desktop → Settings → Resources → Memory (4GB+)
2. **Use SSD Storage**: Better performance for database operations
3. **Close Unnecessary Applications**: Free up system resources
4. **Monitor Resource Usage**: `docker stats` to check container resource usage

## 🔧 Windows PowerShell Setup (Important!)

### PowerShell Execution Policy Fix
If you encounter "execution policy" errors, run this first:
```powershell
# Run as Administrator or use -Scope CurrentUser
Set-ExecutionPolicy -ExecutionPolicy RemoteSigned -Scope CurrentUser -Force

# Verify the policy was set
Get-ExecutionPolicy -List
```

### PowerShell curl Alias Issue
PowerShell has a `curl` alias that conflicts with regular curl commands. Use these alternatives:
```powershell
# Instead of: curl http://localhost:8080/health
# Use:
Invoke-RestMethod -Uri "http://localhost:8080/health"

# Or for web requests:
Invoke-WebRequest -Uri "http://localhost:8091/pools"

# To use real curl in PowerShell:
curl.exe http://localhost:8080/health
```

### Enhanced Service Health Check (Windows)
Use our improved PowerShell health checker:
```powershell
# Run the enhanced health checker (handles Couchbase startup delays)
.\scripts\wait-for-services.ps1

# Or run with custom timeout (default is 300 seconds)
.\scripts\wait-for-services.ps1 -TimeoutSeconds 600
```

### Couchbase Startup Time
Couchbase typically takes **2-3 minutes** to fully initialize. Be patient!
```powershell
# Monitor Couchbase logs while waiting
docker-compose logs -f couchbase

# Test Couchbase readiness (401 Unauthorized means it's ready for setup)
try { 
    Invoke-WebRequest -Uri "http://localhost:8091/pools" -TimeoutSec 10 
} catch { 
    Write-Host "Status: $($_.Exception.Response.StatusCode)" 
}
```
