# 🐛 Critical Bug Fix Required - Docker Build Issue

## 🚨 ROOT CAUSE IDENTIFIED

The error `{"error":"Failed to query students"}` is caused by **Docker containers running the WRONG code**.

### The Real Problem
The project has **TWO implementations**:

1. **❌ OLD BROKEN**: `services/*/main.go` (broken N1QL queries)
2. **✅ NEW WORKING**: `services/*/cmd/main.go` (proper repository pattern)

**Docker is building the OLD broken version!**

## 🔧 Immediate Fix Required

### Fix 1: Update All Dockerfiles
Change build target in each service's Dockerfile:

**In `services/student-service/Dockerfile`:**
```dockerfile
# ❌ BROKEN (line ~15):
RUN CGO_ENABLED=0 GOOS=linux go build -o student-service ./main.go

# ✅ CHANGE TO:
RUN CGO_ENABLED=0 GOOS=linux go build -o student-service ./cmd/main.go
```

**Apply same fix to:**
- `services/teacher-service/Dockerfile`
- `services/academic-service/Dockerfile` 
- `services/achievement-service/Dockerfile`

### Fix 2: Rebuild Containers
```powershell
# Stop services
docker-compose down

# Rebuild with no cache
docker-compose build --no-cache

# Start services
docker-compose up -d
```

## 🧪 Verification After Fix

### Test Health Endpoints
```bash
# Should return proper health status
curl http://localhost:8081/health
curl http://localhost:8082/health
curl http://localhost:8083/health
curl http://localhost:8084/health
```

**Expected Response (NEW implementation):**
```json
{
  "status": "healthy",
  "service": "student-service",
  "couchbase_status": "connected", 
  "database": "couchbase"
}
```

**vs Old Response (BROKEN implementation):**
```json
{
  "status": "healthy",
  "service": "student-service",
  "database": "couchbase-connected"
}
```

### Test CRUD Operations
```bash
# These should ALL work after fix:
curl -X GET http://localhost:8081/api/v1/students
curl -X GET http://localhost:8082/api/v1/teachers
curl -X GET http://localhost:8083/api/v1/academics
curl -X GET http://localhost:8084/api/v1/achievements
```

## 💡 Why This Happened

The project was refactored to use proper repository pattern (`cmd/main.go`) but Docker was never updated to build the new version.

**Evidence:**
- ✅ `cmd/main.go` uses proper shared database package
- ✅ `cmd/main.go` has proper error handling
- ✅ `cmd/main.go` uses repository pattern
- ❌ `main.go` has broken N1QL queries
- ❌ `main.go` has incorrect JSON decoding
- ❌ Docker builds `main.go` (wrong file)

## 🎯 The Fix is Simple

**Just change 4 Dockerfile lines and rebuild containers!**

This will instantly fix all CRUD operations across all services.
