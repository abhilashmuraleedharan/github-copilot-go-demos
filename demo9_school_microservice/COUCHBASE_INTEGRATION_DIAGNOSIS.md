# 🔧 Couchbase Integration Diagnosis & Fix Guide

## 🚨 Identified Issues & Solutions

### Issue 1: Service Architecture Inconsistency
**Problem:** The services have two different handler patterns:
- **Old pattern**: Direct handlers in `main.go` (currently active)
- **New pattern**: Repository-Service-Handler architecture (exists but not used)

**Current Active Code:** The services are using direct handlers in `main.go` which **DO** have Couchbase integration.

### Issue 2: Service Startup Dependencies
**Problem:** Services may start before Couchbase is fully initialized.

**Solution:** Use proper startup sequence:

```powershell
# 1. Start ONLY Couchbase first
docker-compose up -d couchbase

# 2. Wait for Couchbase to be ready (5-10 minutes on first start)
# Watch logs until you see "Couchbase Server has started"
docker-compose logs -f couchbase

# 3. Initialize Couchbase
.\scripts\init-couchbase.ps1

# 4. Start remaining services
docker-compose up -d
```

### Issue 3: Bucket/Collection Configuration
**Problem:** Services expect `schoolmgmt` bucket but may not exist.

**Fix:**
```powershell
# Check if bucket exists
curl -u Administrator:password123 http://localhost:8091/pools/default/buckets

# If bucket doesn't exist, create it
curl -X POST http://localhost:8091/pools/default/buckets \
  -u Administrator:password123 \
  -d 'name=schoolmgmt&ramQuotaMB=256&bucketType=membase'

# Create primary index for N1QL queries
curl -u Administrator:password123 \
  -X POST http://localhost:8093/query/service \
  -d 'statement=CREATE PRIMARY INDEX ON schoolmgmt'
```

## 🔍 Step-by-Step Diagnosis Process

### Step 1: Check Service Status
```powershell
# Check what's running
docker ps

# Check service logs for errors
docker-compose logs student-service | Select-String -Pattern "error|fail|panic"
docker-compose logs teacher-service | Select-String -Pattern "error|fail|panic"
docker-compose logs academic-service | Select-String -Pattern "error|fail|panic"
```

### Step 2: Test Couchbase Connectivity
```powershell
# Check if Couchbase is accessible
curl http://localhost:8091/pools

# Check if authenticated access works
curl -u Administrator:password123 http://localhost:8091/pools/default
```

### Step 3: Test Service Endpoints
```powershell
# Test health endpoints
curl http://localhost:8081/health
curl http://localhost:8082/health
curl http://localhost:8083/health
curl http://localhost:8084/health
```

### Step 4: Test Database Operations
```powershell
# Try creating a student to test database insertion
curl -X POST http://localhost:8081/api/v1/students \
  -H "Content-Type: application/json" \
  -d '{
    "firstName": "Test",
    "lastName": "Student", 
    "email": "test@school.edu",
    "grade": "10"
  }'

# Check if the data was inserted in Couchbase
curl -u Administrator:password123 \
  -X POST http://localhost:8093/query/service \
  -d 'statement=SELECT * FROM schoolmgmt WHERE type = "student"'
```

## 🛠️ Complete Fix Procedure

### 1. Clean Start
```powershell
# Stop all services
docker-compose down

# Remove volumes to start fresh (CAUTION: This deletes all data)
docker-compose down -v

# Start only Couchbase
docker-compose up -d couchbase
```

### 2. Wait for Couchbase Initialization
```powershell
# Monitor Couchbase logs until ready
docker-compose logs -f couchbase

# Look for this message: "Couchbase Server has started"
# This may take 5-10 minutes on first startup
```

### 3. Initialize Couchbase
```powershell
# Run initialization script
.\scripts\init-couchbase.ps1

# OR manual initialization:
# 1. Open http://localhost:8091
# 2. Setup new cluster
# 3. Username: Administrator, Password: password123
# 4. Create bucket named "schoolmgmt"
```

### 4. Start All Services
```powershell
# Start remaining services
docker-compose up -d

# Wait for all services to be healthy
Start-Sleep 30
```

### 5. Verify Integration
```powershell
# Check all services are healthy
curl http://localhost:8081/health
curl http://localhost:8082/health
curl http://localhost:8083/health
curl http://localhost:8084/health

# Test data insertion
curl -X POST http://localhost:8081/api/v1/students \
  -H "Content-Type: application/json" \
  -d '{
    "firstName": "Integration",
    "lastName": "Test",
    "email": "integration.test@school.edu",
    "grade": "12"
  }'

# Verify data in Couchbase
curl -u Administrator:password123 \
  -X POST http://localhost:8093/query/service \
  -d 'statement=SELECT COUNT(*) FROM schoolmgmt WHERE type = "student"'
```

## 🚨 Common Error Messages & Solutions

### Error: "connection refused"
**Cause:** Couchbase not running or not ready
**Solution:** 
```powershell
docker-compose up -d couchbase
# Wait 5-10 minutes for startup
```

### Error: "bucket not found"
**Cause:** schoolmgmt bucket doesn't exist
**Solution:**
```powershell
curl -X POST http://localhost:8091/pools/default/buckets \
  -u Administrator:password123 \
  -d 'name=schoolmgmt&ramQuotaMB=256&bucketType=membase'
```

### Error: "401 Unauthorized"
**Cause:** Couchbase not initialized with credentials
**Solution:**
```powershell
.\scripts\init-couchbase.ps1
```

### Error: "cluster not ready"
**Cause:** Services starting before Couchbase is fully operational
**Solution:** Use staged startup process above

## 📊 Monitoring & Logging

### Enhanced Logging
Add this to check Couchbase operations:
```powershell
# Follow service logs with timestamps
docker-compose logs -f --timestamps student-service

# Filter for Couchbase-related logs
docker-compose logs student-service | Select-String -Pattern "couchbase|bucket|cluster"
```

### Couchbase Query Interface
1. Open http://localhost:8091
2. Go to "Query" tab
3. Run diagnostic queries:
```sql
-- Check all documents
SELECT META().id, type FROM schoolmgmt;

-- Count documents by type
SELECT type, COUNT(*) as count FROM schoolmgmt GROUP BY type;

-- Check recent students
SELECT * FROM schoolmgmt WHERE type = "student" ORDER BY created_at DESC LIMIT 5;
```

## ✅ Success Indicators

After following the fix procedure, you should see:

1. **Service Health Checks Pass:**
```json
{
  "status": "healthy",
  "service": "student-service",
  "database": "couchbase-connected"
}
```

2. **Data Insertion Works:**
```json
{
  "success": true,
  "data": {
    "id": "uuid-here",
    "firstName": "Test",
    "lastName": "Student"
  },
  "message": "Student created successfully"
}
```

3. **Data Retrieval Works:**
```json
{
  "success": true,
  "data": {
    "students": [...],
    "count": 1
  }
}
```

4. **Couchbase Query Shows Data:**
```json
{
  "results": [
    {
      "count": 1
    }
  ]
}
```

## 🔧 Emergency Debugging Commands

If services still can't connect to Couchbase:

```powershell
# Check network connectivity
docker network ls
docker network inspect schoolmgmt-network

# Check Couchbase container details
docker inspect schoolmgmt-couchbase

# Test connection from within service container
docker exec -it schoolmgmt-student-service curl http://couchbase:8091/pools

# Check environment variables
docker exec -it schoolmgmt-student-service env | Select-String -Pattern "COUCHBASE"
```

---

**The main issue is likely that Couchbase needs proper initialization before the services start. Follow the "Complete Fix Procedure" above for the best results.**
