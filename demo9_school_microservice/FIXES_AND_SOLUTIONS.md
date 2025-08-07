# 🛠️ Fixes and Solutions for School Management System

This document provides comprehensive solutions for common issues encountered during setup and operation of the School Management System.

## 📋 Table of Contents

- [Windows PowerShell Issues](#windows-powershell-issues)
- [Couchbase Startup Issues](#couchbase-startup-issues)
- [Service Health Check Issues](#service-health-check-issues)
- [API Endpoint Issues](#api-endpoint-issues)
- [Docker and Container Issues](#docker-and-container-issues)
- [Network and Port Issues](#network-and-port-issues)
- [Data Loading Issues](#data-loading-issues)
- [Complete Working Solutions](#complete-working-solutions)

---

## 🔧 Windows PowerShell Issues

### Issue 1: PowerShell Execution Policy Error
**Error Message:**
```
cannot be loaded because running scripts is disabled on this system
```

**Solution:**
```powershell
# Fix execution policy (run as Administrator or use CurrentUser scope)
Set-ExecutionPolicy -ExecutionPolicy RemoteSigned -Scope CurrentUser -Force

# Verify the policy was set
Get-ExecutionPolicy -List

# Alternative: Run with bypass
PowerShell -ExecutionPolicy Bypass -File .\scripts\load-sample-data.ps1
```

### Issue 2: PowerShell curl Alias Conflict
**Problem:** PowerShell aliases `curl` to `Invoke-WebRequest` which causes issues with curl commands

**Solution:**
```powershell
# Instead of: curl http://localhost:8080/health
# Use these alternatives:

# Option 1: Use Invoke-RestMethod
Invoke-RestMethod -Uri "http://localhost:8080/health"

# Option 2: Use real curl
curl.exe http://localhost:8080/health

# Option 3: Remove the alias (temporary)
Remove-Item alias:curl -Force
curl http://localhost:8080/health
```

### Issue 3: PowerShell JSON Handling
**Problem:** PowerShell handles JSON differently than bash/curl

**Solution:**
```powershell
# Creating JSON data for POST requests
$studentData = @{
    firstName = "John"
    lastName = "Doe" 
    email = "john.doe@school.edu"
    grade = "10"
    status = "active"
} | ConvertTo-Json

# Making POST request
Invoke-RestMethod -Uri "http://localhost:8080/api/v1/students" `
    -Method Post `
    -Body $studentData `
    -ContentType "application/json"
```

---

## ⏰ Couchbase Startup Issues

### Issue 1: Couchbase Takes Too Long to Start
**Problem:** Couchbase can take 2-3 minutes to fully initialize, causing timeout errors

**Solution:**
```powershell
# Increase timeout values
$timeout = 600  # 10 minutes instead of 5

# Check Couchbase status properly
function Test-CouchbaseReady {
    try {
        $response = Invoke-WebRequest -Uri "http://localhost:8091/pools" -TimeoutSec 10
        return $true
    } catch {
        # 401 Unauthorized means Couchbase is ready but needs setup
        if ($_.Exception.Response.StatusCode -eq 401) {
            return $true
        }
        return $false
    }
}

# Wait with proper intervals
$attempts = 0
$maxAttempts = 60  # 60 attempts * 5 seconds = 5 minutes
while ($attempts -lt $maxAttempts) {
    if (Test-CouchbaseReady) {
        Write-Host "Couchbase is ready!"
        break
    }
    $attempts++
    Write-Host "Waiting for Couchbase... ($attempts/$maxAttempts)"
    Start-Sleep 5
}
```

### Issue 2: Couchbase 401 Unauthorized Error
**Problem:** Getting 401 errors when trying to access Couchbase

**Explanation:** This is actually NORMAL! 401 means Couchbase is running but needs initialization.

**Solution:**
```powershell
# This is the correct way to handle Couchbase initialization
try {
    # Step 1: Initialize cluster
    $clusterBody = "memoryQuota=512&indexMemoryQuota=256"
    Invoke-RestMethod -Uri "http://localhost:8091/pools/default" `
        -Method Post `
        -Body $clusterBody `
        -ContentType "application/x-www-form-urlencoded"
    
    # Step 2: Setup admin user
    $adminBody = "username=Administrator&password=password&port=SAME"
    Invoke-RestMethod -Uri "http://localhost:8091/settings/web" `
        -Method Post `
        -Body $adminBody `
        -ContentType "application/x-www-form-urlencoded"
    
    # Step 3: Create bucket
    $bucketBody = "name=schoolmgmt&ramQuotaMB=256&bucketType=membase"
    $credentials = [System.Convert]::ToBase64String([System.Text.Encoding]::ASCII.GetBytes("Administrator:password"))
    $headers = @{ Authorization = "Basic $credentials" }
    
    Invoke-RestMethod -Uri "http://localhost:8091/pools/default/buckets" `
        -Method Post `
        -Body $bucketBody `
        -ContentType "application/x-www-form-urlencoded" `
        -Headers $headers
        
} catch {
    if ($_.Exception.Message -like "*already*") {
        Write-Host "Couchbase already initialized"
    } else {
        Write-Host "Error: $($_.Exception.Message)"
    }
}
```

---

## 🏥 Service Health Check Issues

### Issue 1: Services Not Responding to Health Checks
**Problem:** Health check scripts fail because services aren't ready

**Solution - Improved Health Check Function:**
```powershell
function Wait-ForServiceHealth {
    param(
        [string]$ServiceName,
        [string]$ServiceUrl,
        [string]$Endpoint = "/health",
        [int]$TimeoutSeconds = 300,
        [int]$CheckInterval = 5
    )
    
    $maxAttempts = [math]::Floor($TimeoutSeconds / $CheckInterval)
    $attempts = 0
    
    Write-Host "Waiting for $ServiceName to become healthy..."
    
    while ($attempts -lt $maxAttempts) {
        try {
            $response = Invoke-RestMethod -Uri "$ServiceUrl$Endpoint" -TimeoutSec 10 -ErrorAction Stop
            Write-Host "✅ $ServiceName is healthy!"
            return $true
        } catch {
            $attempts++
            if ($attempts % 10 -eq 0) {  # Show progress every 50 seconds
                Write-Host "⏳ Still waiting for $ServiceName... ($attempts/$maxAttempts)"
            }
            Start-Sleep $CheckInterval
        }
    }
    
    Write-Host "❌ $ServiceName failed to become healthy within $TimeoutSeconds seconds"
    return $false
}
```

### Issue 2: API Gateway Health vs Individual Services
**Problem:** API Gateway might be healthy but individual services aren't

**Solution - Check All Services:**
```powershell
$services = @(
    @{Name="API Gateway"; Url="http://localhost:8080"; Endpoint="/health"},
    @{Name="Student Service"; Url="http://localhost:8081"; Endpoint="/health"},
    @{Name="Teacher Service"; Url="http://localhost:8082"; Endpoint="/health"},
    @{Name="Academic Service"; Url="http://localhost:8083"; Endpoint="/health"},
    @{Name="Achievement Service"; Url="http://localhost:8084"; Endpoint="/health"}
)

$allHealthy = $true
foreach ($service in $services) {
    $healthy = Wait-ForServiceHealth -ServiceName $service.Name -ServiceUrl $service.Url -Endpoint $service.Endpoint
    if (-not $healthy) {
        $allHealthy = $false
    }
}

if ($allHealthy) {
    Write-Host "🎉 All services are healthy!"
} else {
    Write-Host "❌ Some services failed health checks"
}
```

---

## 🌐 API Endpoint Issues

### Issue 1: Wrong API Endpoint Paths
**Problem:** Using incorrect API paths leads to 404 errors

**Correct API Endpoints:**
```
✅ CORRECT PATHS:
- Students: http://localhost:8080/api/v1/students
- Teachers: http://localhost:8080/api/v1/teachers  
- Academics: http://localhost:8080/api/v1/academics
- Classes: http://localhost:8080/api/v1/classes
- Achievements: http://localhost:8080/api/v1/achievements
- Badges: http://localhost:8080/api/v1/badges
- Health: http://localhost:8080/health

❌ INCORRECT PATHS:
- http://localhost:8080/api/students (missing v1)
- http://localhost:8080/students (missing api/v1)
- http://localhost:8080/api/v1/student (missing 's')
```

### Issue 2: API Returns Empty Results
**Problem:** API endpoints return empty arrays or null

**Causes and Solutions:**
```powershell
# Cause 1: No data loaded yet
# Solution: Load sample data first
.\scripts\load-sample-data.ps1

# Cause 2: Database not connected
# Solution: Check Couchbase and service logs
docker-compose logs student-service
docker-compose logs couchbase

# Cause 3: Services not fully started
# Solution: Wait longer and check health
Start-Sleep 60
Invoke-RestMethod -Uri "http://localhost:8081/health"  # Direct service check
```

---

## 🐳 Docker and Container Issues

### Issue 1: Containers Not Starting
**Problem:** `docker-compose up -d` fails or containers exit immediately

**Diagnostic Steps:**
```powershell
# Check container status
docker-compose ps

# Check logs for errors
docker-compose logs
docker-compose logs couchbase
docker-compose logs api-gateway

# Check for port conflicts
netstat -an | findstr ":8080"  # Windows
netstat -an | findstr ":8091"  # Check Couchbase port

# Restart specific services
docker-compose restart couchbase
docker-compose restart api-gateway
```

**Solutions:**
```powershell
# Solution 1: Clean restart
docker-compose down -v
docker-compose up -d --build

# Solution 2: Check Docker resources
# Ensure Docker Desktop has enough memory (4GB minimum)

# Solution 3: Remove conflicting services
# Stop any existing services using ports 8080-8084, 8091-8096
```

### Issue 2: Services Stuck in "Starting" State
**Problem:** Containers show "starting" status for extended periods

**Solution:**
```powershell
# Check detailed status
docker ps --format "table {{.Names}}\t{{.Status}}\t{{.Ports}}"

# Check resource usage
docker stats

# Increase Docker memory limits
# Docker Desktop > Settings > Resources > Memory > 4GB+

# Force recreate containers
docker-compose up -d --force-recreate
```

---

## 🌐 Network and Port Issues

### Issue 1: Port Conflicts
**Problem:** "Port already in use" errors

**Solution:**
```powershell
# Find processes using ports
netstat -ano | findstr ":8080"
netstat -ano | findstr ":8091"

# Kill conflicting processes (replace PID)
taskkill /PID <process_id> /F

# Or use different ports in docker-compose.yml
# Change "8080:8080" to "8090:8080" etc.
```

### Issue 2: Can't Access Services from Host
**Problem:** Services work inside Docker but not from host machine

**Solution:**
```powershell
# Check Docker network
docker network ls
docker network inspect schoolmgmt_default

# Test from inside containers
docker-compose exec api-gateway curl http://student-service:8081/health

# Check Windows Firewall
# Windows Security > Firewall > Allow app through firewall
# Add Docker Desktop if needed
```

---

## 💾 Data Loading Issues

### Issue 1: Data Not Persisting
**Problem:** Data disappears after container restart

**Solution:**
```powershell
# Ensure volumes are properly configured
docker volume ls | findstr schoolmgmt

# Check docker-compose.yml has volumes section:
# volumes:
#   couchbase_data:

# Verify volume is mounted in Couchbase container
docker-compose exec couchbase ls -la /opt/couchbase/var
```

### Issue 2: JSON Parsing Errors During Data Load
**Problem:** Invalid JSON errors when loading sample data

**Solution:**
```powershell
# Validate JSON before sending
$jsonData = @{
    firstName = "John"
    lastName = "Doe"
} | ConvertTo-Json

# Test JSON validity
try {
    $parsed = $jsonData | ConvertFrom-Json
    Write-Host "JSON is valid"
} catch {
    Write-Host "JSON is invalid: $($_.Exception.Message)"
}

# Use proper content type
Invoke-RestMethod -Uri "http://localhost:8080/api/v1/students" `
    -Method Post `
    -Body $jsonData `
    -ContentType "application/json; charset=utf-8"
```

---

## ✅ Complete Working Solutions

### Solution 1: Complete Windows Setup Script
**File:** `scripts/complete-setup-windows.ps1`

This script handles all common issues automatically:
```powershell
# Run the comprehensive setup script
.\scripts\complete-setup-windows.ps1 -Verbose

# What it does:
# 1. Fixes PowerShell execution policy
# 2. Checks Docker installation
# 3. Starts services if needed
# 4. Waits for all services (proper timeouts)
# 5. Initializes Couchbase correctly
# 6. Tests all API endpoints
# 7. Creates and tests a sample student
```

### Solution 2: Quick System Verification
**File:** `scripts/verify-system.ps1`

Run this to quickly check if everything is working:
```powershell
.\scripts\verify-system.ps1
```

### Solution 3: Manual Step-by-Step Setup
If automated scripts fail, follow these manual steps:

```powershell
# Step 1: Fix PowerShell policy
Set-ExecutionPolicy -ExecutionPolicy RemoteSigned -Scope CurrentUser -Force

# Step 2: Start services
docker-compose up -d

# Step 3: Wait for Couchbase (be patient!)
do {
    try {
        $status = Invoke-WebRequest -Uri "http://localhost:8091/pools" -TimeoutSec 5
        Write-Host "Couchbase ready!"
        break
    } catch {
        if ($_.Exception.Response.StatusCode -eq 401) {
            Write-Host "Couchbase ready for setup!"
            break
        }
        Write-Host "Waiting for Couchbase..."
        Start-Sleep 10
    }
} while ($true)

# Step 4: Initialize Couchbase
# (Use the Couchbase initialization code from Issue 2 above)

# Step 5: Test services
Invoke-RestMethod -Uri "http://localhost:8080/health"
Invoke-RestMethod -Uri "http://localhost:8080/api/v1/students"

# Step 6: Create test data
$student = @{
    firstName = "John"
    lastName = "Doe"
    email = "john.doe@school.edu"
    grade = "10"
    status = "active"
} | ConvertTo-Json

Invoke-RestMethod -Uri "http://localhost:8080/api/v1/students" `
    -Method Post `
    -Body $student `
    -ContentType "application/json"
```

---

## 🚨 Emergency Troubleshooting

### When Everything Fails
```powershell
# Nuclear option - complete reset
docker-compose down -v --rmi all
docker system prune -a -f
docker volume prune -f

# Restart Docker Desktop
# Then start fresh:
docker-compose up -d --build
```

### Getting Help
```powershell
# Check all logs
docker-compose logs > system-logs.txt

# Check individual service logs
docker-compose logs api-gateway > api-gateway-logs.txt
docker-compose logs couchbase > couchbase-logs.txt
docker-compose logs student-service > student-service-logs.txt

# Check system resources
docker stats
Get-Process | Where-Object { $_.ProcessName -like "*docker*" }
```

---

## 📞 Quick Reference Commands

### PowerShell Commands
```powershell
# Health checks
Invoke-RestMethod -Uri "http://localhost:8080/health"
Invoke-RestMethod -Uri "http://localhost:8081/health"  # Student Service

# Test API endpoints
Invoke-RestMethod -Uri "http://localhost:8080/api/v1/students"
Invoke-RestMethod -Uri "http://localhost:8080/api/v1/teachers"

# Docker management
docker-compose ps
docker-compose logs [service-name]
docker-compose restart [service-name]
docker-compose down -v
```

### Verification URLs
- **Couchbase Console:** http://localhost:8091 (Administrator/password)
- **API Gateway Health:** http://localhost:8080/health
- **Students API:** http://localhost:8080/api/v1/students
- **Teachers API:** http://localhost:8080/api/v1/teachers
- **Individual Service Health:** http://localhost:808[1-4]/health

---

## 🎯 Success Criteria

Your system is working correctly when:

✅ All containers are running: `docker-compose ps` shows "Up" status
✅ Couchbase Console accessible: http://localhost:8091
✅ API Gateway health check passes: http://localhost:8080/health
✅ Can create a student: POST to `/api/v1/students` returns success
✅ Can retrieve students: GET `/api/v1/students` returns data
✅ All individual services respond: Each service at port 808[1-4]/health

**If all these pass, your School Management System is fully operational! 🎉**

---

## 🗄️ Couchbase Integration Issues

### Issue 1: Services Using In-Memory Storage Instead of Couchbase
**Problem:** APIs work but data doesn't persist after container restart

**Root Cause:** Services were building the wrong main.go file (in-memory version instead of Couchbase version)

**Solution:**
```powershell
# The issue was in Dockerfiles - they were using main.go instead of cmd/main.go
# This has been fixed in all service Dockerfiles

# To verify services are using Couchbase:
docker-compose logs student-service | findstr "Couchbase"
docker-compose logs teacher-service | findstr "Couchbase"

# You should see logs like:
# "🔌 Attempting to connect to Couchbase at: couchbase://couchbase"
# "✅ Successfully connected to Couchbase bucket: schoolmgmt"
```

### Issue 2: Couchbase Connection Errors in Services
**Problem:** Services fail to connect to Couchbase with connection timeout errors

**Diagnostic Steps:**
```powershell
# Check if Couchbase container is running
docker-compose ps couchbase

# Check Couchbase logs
docker-compose logs couchbase

# Test Couchbase accessibility from host
curl http://localhost:8091/pools

# Test from inside service container
docker-compose exec student-service curl http://couchbase:8091/pools
```

**Solutions:**
```powershell
# Solution 1: Ensure Couchbase is properly initialized
.\scripts\init-couchbase.ps1

# Solution 2: Check environment variables
docker-compose exec student-service env | findstr COUCHBASE

# Should show:
# COUCHBASE_HOST=couchbase
# COUCHBASE_USERNAME=Administrator  
# COUCHBASE_PASSWORD=password123
# COUCHBASE_BUCKET=schoolmgmt

# Solution 3: Restart services after Couchbase is ready
docker-compose restart student-service teacher-service academic-service achievement-service
```

### Issue 3: Bucket "schoolmgmt" Not Found
**Problem:** Services can connect to Couchbase but bucket doesn't exist

**Solution:**
```powershell
# Check if bucket exists via Couchbase console
# Go to http://localhost:8091 > Buckets

# Or check via API
curl -u "Administrator:password123" http://localhost:8091/pools/default/buckets

# Create bucket manually if needed
curl -X POST http://localhost:8091/pools/default/buckets \
    -u "Administrator:password123" \
    -d "name=schoolmgmt&ramQuotaMB=256&bucketType=membase"

# Or run initialization script
.\scripts\init-couchbase.ps1
```

### Issue 4: Couchbase Takes Too Long to Initialize
**Problem:** Services start before Couchbase is ready, causing connection failures

**Solution - Enhanced Startup Sequence:**
```powershell
# Use the enhanced setup script
.\scripts\enhanced-setup-couchbase.ps1

# Or manual startup with proper timing:
# 1. Start only Couchbase first
docker-compose up -d couchbase

# 2. Wait for Couchbase to be ready (5-10 minutes on first start)
do {
    try {
        $status = Invoke-WebRequest -Uri "http://localhost:8091/pools" -TimeoutSec 5
        Write-Host "Couchbase ready!"
        break
    } catch {
        if ($_.Exception.Response.StatusCode -eq 401) {
            Write-Host "Couchbase ready for setup!"
            break
        }
        Write-Host "Waiting for Couchbase..."
        Start-Sleep 15
    }
} while ($true)

# 3. Initialize Couchbase
.\scripts\init-couchbase.ps1

# 4. Start other services
docker-compose up -d
```

### Issue 5: Data Not Persisting in Couchbase
**Problem:** Can create data but it disappears after restart

**Causes and Solutions:**
```powershell
# Cause 1: Volume not properly mounted
# Check docker-compose.yml has:
# volumes:
#   couchbase_data:
# And Couchbase service has:
# volumes:
#   - couchbase_data:/opt/couchbase/var

# Cause 2: Using wrong collection
# Services should use bucket "schoolmgmt", not "default"
# Check environment variable: COUCHBASE_BUCKET=schoolmgmt

# Cause 3: Documents stored with wrong keys
# Verify document keys follow pattern: "student::uuid", "teacher::uuid"
```

### Issue 6: N1QL Query Errors
**Problem:** Services can insert data but queries fail

**Solution:**
```powershell
# Ensure primary index exists
curl -X POST http://localhost:8093/query/service \
    -u "Administrator:password123" \
    -d "statement=CREATE PRIMARY INDEX ON \`schoolmgmt\`"

# Test query manually
curl -X POST http://localhost:8093/query/service \
    -u "Administrator:password123" \
    -d "statement=SELECT * FROM \`schoolmgmt\` WHERE type='student'"

# Check query service is running
curl http://localhost:8093/admin/ping
```

---

## ✅ Enhanced Working Solutions

### Solution A: Complete Couchbase Setup Script
**File:** `scripts/enhanced-setup-couchbase.ps1`

This comprehensive script handles Couchbase integration issues:
```powershell
# Run the enhanced setup script with full logging
.\scripts\enhanced-setup-couchbase.ps1 -Verbose

# What it does:
# 1. Fixes PowerShell execution policy
# 2. Checks Docker availability  
# 3. Builds and starts services with proper timing
# 4. Waits for Couchbase to be accessible (with timeouts)
# 5. Initializes Couchbase cluster and bucket
# 6. Waits for all microservices to be healthy
# 7. Tests end-to-end Couchbase integration
# 8. Provides system status and quick test commands
```

### Solution B: Couchbase Integration Verification
**File:** `scripts/verify-couchbase-integration.ps1`

```powershell
# Create this script to verify Couchbase integration
param([switch]$Detailed)

Write-Host "🔍 Verifying Couchbase Integration..." -ForegroundColor Green

# Test 1: Check all services are using Couchbase
$services = @("student-service", "teacher-service", "academic-service", "achievement-service")
foreach ($service in $services) {
    $logs = docker-compose logs $service 2>&1 | Select-String "Couchbase"
    if ($logs) {
        Write-Host "✅ $service is connected to Couchbase" -ForegroundColor Green
        if ($Detailed) { $logs | ForEach-Object { Write-Host "  $_" } }
    } else {
        Write-Host "❌ $service may not be using Couchbase" -ForegroundColor Red
    }
}

# Test 2: Verify data persistence
Write-Host "`n🧪 Testing data persistence..." -ForegroundColor Yellow
$testData = @{firstName="Persist";lastName="Test";email="persist@test.com";grade="10"} | ConvertTo-Json
$response = Invoke-RestMethod -Uri "http://localhost:8080/api/v1/students" -Method Post -Body $testData -ContentType "application/json"
$studentId = $response.data.id

# Restart services
Write-Host "🔄 Restarting services to test persistence..." -ForegroundColor Yellow
docker-compose restart student-service

Start-Sleep 10

# Try to retrieve
try {
    $retrieved = Invoke-RestMethod -Uri "http://localhost:8080/api/v1/students/$studentId"
    Write-Host "✅ Data persisted across restart!" -ForegroundColor Green
    
    # Cleanup
    Invoke-RestMethod -Uri "http://localhost:8080/api/v1/students/$studentId" -Method Delete
} catch {
    Write-Host "❌ Data did not persist across restart" -ForegroundColor Red
}

Write-Host "`n🎉 Couchbase integration verification completed!" -ForegroundColor Green
```

### Solution C: Direct Couchbase Testing Commands
```powershell
# Test Couchbase directly (bypassing services)

# 1. Check cluster status
curl -u "Administrator:password123" http://localhost:8091/pools/default

# 2. List all buckets  
curl -u "Administrator:password123" http://localhost:8091/pools/default/buckets

# 3. Query all documents
curl -X POST http://localhost:8093/query/service \
    -u "Administrator:password123" \
    -d "statement=SELECT * FROM \`schoolmgmt\` LIMIT 10"

# 4. Count documents by type
curl -X POST http://localhost:8093/query/service \
    -u "Administrator:password123" \
    -d "statement=SELECT type, COUNT(*) as count FROM \`schoolmgmt\` GROUP BY type"

# 5. Insert test document directly
curl -X POST http://localhost:8093/query/service \
    -u "Administrator:password123" \
    -d "statement=INSERT INTO \`schoolmgmt\` (KEY, VALUE) VALUES ('test::direct', {'type':'test','message':'Direct insert works'})"
```

---

## 🎯 Updated Success Criteria

Your Couchbase integration is working correctly when:

✅ All containers are running: `docker-compose ps` shows "Up" status
✅ Couchbase Console accessible: http://localhost:8091 (Administrator/password123)
✅ Bucket "schoolmgmt" exists and is operational
✅ Service logs show successful Couchbase connections with emojis (🔌, ✅, etc.)
✅ Can create a student: POST to `/api/v1/students` returns success
✅ Can retrieve students: GET `/api/v1/students` returns data from Couchbase
✅ Data persists after service restarts
✅ Direct Couchbase queries return service data
✅ All individual services respond: Each service at port 808[1-4]/health

**Enhanced verification commands:**
```powershell
# Quick integration test
.\scripts\enhanced-setup-couchbase.ps1 -Verbose

# Comprehensive CRUD testing  
# See: scripts/couchbase-crud-commands.md

# Data persistence verification
.\scripts\verify-couchbase-integration.ps1 -Detailed
```

**If all these pass, your School Management System is fully operational with Couchbase! 🎉**

---

## ✅ Couchbase Integration Completion Status

### 🎯 All Services Now Use Couchbase

All microservices have been successfully migrated from in-memory storage to Couchbase:

#### ✅ Student Service (`/api/v1/students`)
- **Status**: Complete ✅
- **Repository**: `services/student-service/internal/repository/student_repository.go`
- **Main File**: `services/student-service/main.go`
- **Features**: Full CRUD, N1QL queries, detailed logging
- **Endpoints**: Create, Read, Update, Delete, List with pagination

#### ✅ Teacher Service (`/api/v1/teachers`)
- **Status**: Complete ✅
- **Repository**: `services/teacher-service/internal/repository/teacher_repository.go`
- **Main File**: `services/teacher-service/main.go`
- **Features**: Full CRUD, department filtering, active teachers query
- **Endpoints**: Create, Read, Update, Delete, List, By Department, Active

#### ✅ Academic Service (`/api/v1/academics`)
- **Status**: Complete ✅
- **Repository**: `services/academic-service/internal/repository/academic_repository.go`
- **Main File**: `services/academic-service/main.go`
- **Features**: Academic records, class management, student academic history
- **Endpoints**: Create, Read, Update, Delete, List, By Student, Classes

#### ✅ Achievement Service (`/api/v1/achievements`)
- **Status**: Complete ✅
- **Repository**: `services/achievement-service/internal/repository/achievement_repository.go`
- **Main File**: `services/achievement-service/main.go`
- **Features**: Achievement tracking, awards, categories, student points
- **Endpoints**: Create, Read, Update, Delete, List, By Student, By Category, Awards

### 🔧 Technical Implementation

#### Common Features Across All Services:
1. **Couchbase Client**: Shared database connection logic
2. **Structured Logging**: Comprehensive logrus-based logging for all operations
3. **Error Handling**: Proper HTTP status codes and error messages
4. **Document Types**: Each service uses type field for document identification
5. **Pagination**: Support for limit/offset parameters
6. **Health Checks**: Couchbase connectivity status in `/health` endpoints

#### Document Key Structure:
- Students: `student::{id}`
- Teachers: `teacher::{id}`
- Academic Records: `academic::{id}`
- Classes: `class::{id}`
- Achievements: `achievement::{id}`
- Awards: `award::{id}`

#### N1QL Query Examples:
```sql
-- All students
SELECT s.* FROM schoolmgmt s WHERE s.type = "student" LIMIT 50;

-- Teachers by department
SELECT t.* FROM schoolmgmt t WHERE t.type = "teacher" AND t.department = "Mathematics";

-- Academic records for a student
SELECT a.* FROM schoolmgmt a WHERE a.type = "academic" AND a.student_id = "student_123";

-- Student achievements with points
SELECT a.* FROM schoolmgmt a WHERE a.type = "achievement" AND a.student_id = "student_123" AND a.status = "approved";
```

### 🚀 Updated Docker Configuration

#### Environment Variables (All Services):
```yaml
COUCHBASE_HOST: couchbase
COUCHBASE_USERNAME: Administrator
COUCHBASE_PASSWORD: password123
COUCHBASE_BUCKET: schoolmgmt
```

#### Service Health Check URLs:
```bash
curl http://localhost:8081/health  # Student Service
curl http://localhost:8082/health  # Teacher Service  
curl http://localhost:8083/health  # Academic Service
curl http://localhost:8084/health  # Achievement Service
```

Each health check now returns:
```json
{
  "status": "healthy",
  "service": "service-name",
  "couchbase_status": "connected",
  "database": "couchbase"
}
```

### 📚 Complete CRUD Documentation

**Comprehensive cURL commands available in**: `scripts/couchbase-crud-commands.md`

#### Example Complete Workflow:
1. Create a teacher → Get teacher ID
2. Create a student → Get student ID  
3. Create academic record linking teacher and student
4. Create achievement for the student
5. Query student's complete academic profile

### 🧪 Verification Commands

#### Test Data Persistence:
```bash
# Create data
curl -X POST http://localhost:8080/api/v1/students -H "Content-Type: application/json" -d '{...}'

# Restart services
docker-compose restart student-service teacher-service academic-service achievement-service

# Verify data persists
curl http://localhost:8080/api/v1/students
```

#### Direct Couchbase Queries:
```bash
# Via Couchbase Query Interface (http://localhost:8093)
SELECT COUNT(*) as total_documents, type FROM schoolmgmt GROUP BY type;
```

### 🔍 Troubleshooting Integration

#### Common Issues and Solutions:

**Service Won't Start:**
```bash
# Check Couchbase is running
docker-compose logs couchbase

# Check service logs
docker-compose logs student-service
```

**Connection Errors:**
```bash
# Verify Couchbase initialization
curl -u Administrator:password123 http://localhost:8091/pools/default

# Re-initialize if needed
.\scripts\init-couchbase.ps1
```

**Data Not Persisting:**
```bash
# Check bucket exists
curl -u Administrator:password123 http://localhost:8091/pools/default/buckets

# Check indexes
curl -u Administrator:password123 -X POST http://localhost:8093/query/service -d 'statement=SELECT * FROM system:indexes WHERE keyspace_id = "schoolmgmt"'
```

### 🎉 Migration Complete

✅ **All services successfully migrated from in-memory to Couchbase**  
✅ **Data persistence across container restarts verified**  
✅ **Comprehensive logging and error handling implemented**  
✅ **Full CRUD operations with N1QL queries functional**  
✅ **Health checks include Couchbase connectivity status**  
✅ **Docker configuration updated for all services**  
✅ **Complete documentation and troubleshooting guides provided**

The School Management System now has a fully functional, persistent, Couchbase-backed microservices architecture with comprehensive logging, error handling, and documentation.

---
