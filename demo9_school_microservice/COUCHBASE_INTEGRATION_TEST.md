# 🧪 Couchbase Integration Test Guide

## 🚀 Quick Start - Bringing Up Services

### IMPORTANT: Fix Required First!
**Before running any tests, you MUST fix the Docker build issue:**

```powershell
# Run the automatic fix script
.\fix-dockerfiles.ps1

# OR manual fix - update each Dockerfile to build cmd/main.go:
# In services/student-service/Dockerfile (line ~15):
# Change: RUN CGO_ENABLED=0 GOOS=linux go build -o student-service ./main.go
# To:     RUN CGO_ENABLED=0 GOOS=linux go build -o student-service ./cmd/main.go
# Repeat for teacher-service, academic-service, achievement-service

# Then rebuild:
docker-compose down
docker-compose build --no-cache  
docker-compose up -d
```

### 1. Start Services with Docker Compose
```powershell
# Navigate to project root
cd d:\demo\schoolmgmt

# Start all services (Couchbase + microservices)
docker-compose up -d

# Wait for Couchbase to initialize (5-10 minutes on first run)
# Monitor Couchbase startup
docker-compose logs -f couchbase
```

### 2. Initialize Couchbase Database
```powershell
# Wait for Couchbase to be accessible (watch for "Couchbase Server has started" in logs)
# Then initialize the cluster and bucket
.\scripts\init-couchbase.ps1

# Alternative: Use enhanced setup script
.\scripts\enhanced-setup-couchbase.ps1
```

### 3. Verify All Services Are Running
```powershell
# Check service health
curl http://localhost:8080/health  # API Gateway
curl http://localhost:8081/health  # Student Service
curl http://localhost:8082/health  # Teacher Service  
curl http://localhost:8083/health  # Academic Service
curl http://localhost:8084/health  # Achievement Service

# Check Couchbase Web Console
# Open: http://localhost:8091
# Login: Administrator / password123
```

## 📊 CRUD Operations Test Suite

### 🎓 Student Service (Port 8081)

#### Create Student
```bash
curl -X POST http://localhost:8081/api/v1/students \
  -H "Content-Type: application/json" \
  -d '{
    "firstName": "John",
    "lastName": "Doe", 
    "email": "john.doe@school.edu",
    "grade": "10",
    "status": "active"
  }'
```

#### Get Student by ID
```bash
# Replace {student_id} with actual ID from create response
curl -X GET http://localhost:8081/api/v1/students/{student_id}
```

#### List All Students
```bash
curl -X GET http://localhost:8081/api/v1/students
```

#### Update Student
```bash
curl -X PUT http://localhost:8081/api/v1/students/{student_id} \
  -H "Content-Type: application/json" \
  -d '{
    "firstName": "John Updated",
    "grade": "11"
  }'
```

#### Delete Student
```bash
curl -X DELETE http://localhost:8081/api/v1/students/{student_id}
```

### 👨‍🏫 Teacher Service (Port 8082)

#### Create Teacher
```bash
curl -X POST http://localhost:8082/api/v1/teachers \
  -H "Content-Type: application/json" \
  -d '{
    "firstName": "Jane",
    "lastName": "Smith",
    "email": "jane.smith@school.edu",
    "department": "Mathematics",
    "subject": "Algebra",
    "status": "active"
  }'
```

#### Get Teacher by ID
```bash
curl -X GET http://localhost:8082/api/v1/teachers/{teacher_id}
```

#### List All Teachers
```bash
curl -X GET http://localhost:8082/api/v1/teachers
```

#### Update Teacher
```bash
curl -X PUT http://localhost:8082/api/v1/teachers/{teacher_id} \
  -H "Content-Type: application/json" \
  -d '{
    "department": "Advanced Mathematics",
    "subject": "Calculus"
  }'
```

#### Delete Teacher
```bash
curl -X DELETE http://localhost:8082/api/v1/teachers/{teacher_id}
```

### 📚 Academic Service (Port 8083)

#### Create Academic Record
```bash
curl -X POST http://localhost:8083/api/v1/academics \
  -H "Content-Type: application/json" \
  -d '{
    "studentId": "{student_id}",
    "teacherId": "{teacher_id}",
    "subject": "Mathematics",
    "academicYear": "2024-2025",
    "semester": "Fall",
    "totalMarks": 100,
    "obtainedMarks": 85,
    "grade": "A",
    "remarks": "Excellent performance"
  }'
```

#### Get Academic Record by ID
```bash
curl -X GET http://localhost:8083/api/v1/academics/{academic_id}
```

#### List All Academic Records
```bash
curl -X GET http://localhost:8083/api/v1/academics
```

#### Update Academic Record
```bash
curl -X PUT http://localhost:8083/api/v1/academics/{academic_id} \
  -H "Content-Type: application/json" \
  -d '{
    "obtainedMarks": 90,
    "grade": "A+",
    "remarks": "Outstanding improvement"
  }'
```

#### Delete Academic Record
```bash
curl -X DELETE http://localhost:8083/api/v1/academics/{academic_id}
```

### 🏆 Achievement Service (Port 8084)

#### Create Achievement
```bash
curl -X POST http://localhost:8084/api/v1/achievements \
  -H "Content-Type: application/json" \
  -d '{
    "studentId": "{student_id}",
    "title": "Academic Excellence Award",
    "description": "Top performer in Mathematics",
    "category": "Academic",
    "awardDate": "2024-12-01",
    "level": "School"
  }'
```

#### Get Achievement by ID
```bash
curl -X GET http://localhost:8084/api/v1/achievements/{achievement_id}
```

#### List All Achievements
```bash
curl -X GET http://localhost:8084/api/v1/achievements
```

#### Update Achievement
```bash
curl -X PUT http://localhost:8084/api/v1/achievements/{achievement_id} \
  -H "Content-Type: application/json" \
  -d '{
    "level": "District",
    "description": "Top performer in Regional Mathematics Competition"
  }'
```

#### Delete Achievement
```bash
curl -X DELETE http://localhost:8084/api/v1/achievements/{achievement_id}
```

## 🔄 Via API Gateway (Port 8080)

All the above endpoints are also available through the API Gateway:

```bash
# Replace service port (8081-8084) with 8080 in any curl command
# Example:
curl -X GET http://localhost:8080/api/v1/students
curl -X GET http://localhost:8080/api/v1/teachers
curl -X GET http://localhost:8080/api/v1/academics
curl -X GET http://localhost:8080/api/v1/achievements
```

## 🧪 Step-by-Step Debugging Workflow

### Step 1: Verify Couchbase is Working
```bash
# 1. Check Couchbase is accessible
curl http://localhost:8091/pools

# 2. Check bucket exists  
curl -u Administrator:password123 http://localhost:8091/pools/default/buckets

# 3. Test N1QL service
curl -u Administrator:password123 \
  -X POST http://localhost:8093/query/service \
  -d 'statement=SELECT "Couchbase is working" as message'
```

### Step 2: Test Document Creation
```bash
# 1. Create a test student via service
STUDENT_RESPONSE=$(curl -s -X POST http://localhost:8081/api/v1/students \
  -H "Content-Type: application/json" \
  -d '{
    "firstName": "Test",
    "lastName": "Student",
    "email": "test.student@school.edu",
    "grade": "10"
  }')

echo $STUDENT_RESPONSE

# 2. Extract student ID if creation succeeded
STUDENT_ID=$(echo $STUDENT_RESPONSE | jq -r '.data.id // empty')
echo "Student ID: $STUDENT_ID"

# 3. Check if document was actually created in Couchbase
if [ ! -z "$STUDENT_ID" ]; then
  curl -u Administrator:password123 \
    -X GET http://localhost:8091/pools/default/buckets/schoolmgmt/docs/student::$STUDENT_ID
fi
```

### Step 3: Test Document Retrieval
```bash
# 1. Try to get the student back via service API
if [ ! -z "$STUDENT_ID" ]; then
  curl -X GET http://localhost:8081/api/v1/students/$STUDENT_ID
fi

# 2. Try to list all students via service API  
curl -X GET http://localhost:8081/api/v1/students

# 3. If service fails, test direct N1QL query
curl -u Administrator:password123 \
  -X POST http://localhost:8093/query/service \
  -d 'statement=SELECT * FROM schoolmgmt WHERE type = "student" LIMIT 10'
```

### Step 4: Diagnose Service Logs
```bash
# Check service logs for errors
docker-compose logs student-service --tail=50

# Look for Couchbase connection errors
docker-compose logs student-service | grep -i "couchbase\|error\|fail"

# Check if service is actually connecting to Couchbase
docker-compose logs student-service | grep -i "connected\|cluster\|bucket"
```

## 🧪 Complete Integration Test Workflow

### 1. Create Full Academic Ecosystem
```bash
# Step 1: Create a teacher
TEACHER_RESPONSE=$(curl -s -X POST http://localhost:8081/api/v1/teachers \
  -H "Content-Type: application/json" \
  -d '{
    "firstName": "Dr. Alice",
    "lastName": "Johnson",
    "email": "alice.johnson@school.edu",
    "department": "Science",
    "subject": "Physics"
  }')

TEACHER_ID=$(echo $TEACHER_RESPONSE | jq -r '.data.id')

# Step 2: Create a student
STUDENT_RESPONSE=$(curl -s -X POST http://localhost:8081/api/v1/students \
  -H "Content-Type: application/json" \
  -d '{
    "firstName": "Bob",
    "lastName": "Wilson",
    "email": "bob.wilson@school.edu",
    "grade": "12"
  }')

STUDENT_ID=$(echo $STUDENT_RESPONSE | jq -r '.data.id')

# Step 3: Create academic record
curl -X POST http://localhost:8083/api/v1/academics \
  -H "Content-Type: application/json" \
  -d "{
    \"studentId\": \"$STUDENT_ID\",
    \"teacherId\": \"$TEACHER_ID\",
    \"subject\": \"Physics\",
    \"academicYear\": \"2024-2025\",
    \"semester\": \"Fall\",
    \"totalMarks\": 100,
    \"obtainedMarks\": 92,
    \"grade\": \"A+\"
  }"

# Step 4: Create achievement
curl -X POST http://localhost:8084/api/v1/achievements \
  -H "Content-Type: application/json" \
  -d "{
    \"studentId\": \"$STUDENT_ID\",
    \"title\": \"Physics Excellence Award\",
    \"description\": \"Outstanding performance in Physics\",
    \"category\": \"Academic\",
    \"level\": \"School\"
  }"
```

### 2. Verify Data Persistence
```bash
# Restart all services to test persistence
docker-compose restart student-service teacher-service academic-service achievement-service

# Wait for services to start
Start-Sleep 30

# Verify data still exists
curl http://localhost:8081/api/v1/students
curl http://localhost:8082/api/v1/teachers  
curl http://localhost:8083/api/v1/academics
curl http://localhost:8084/api/v1/achievements
```

## 🛠️ Troubleshooting Commands

### Check Service Health
```bash
curl http://localhost:8080/health  # API Gateway
curl http://localhost:8081/health  # Student Service
curl http://localhost:8082/health  # Teacher Service
curl http://localhost:8083/health  # Academic Service
curl http://localhost:8084/health  # Achievement Service
```

### Check Couchbase Status
```bash
# Check cluster status
curl -u Administrator:password123 http://localhost:8091/pools/default

# Check bucket exists
curl -u Administrator:password123 http://localhost:8091/pools/default/buckets

# Test N1QL query
curl -u Administrator:password123 \
  -X POST http://localhost:8093/query/service \
  -d 'statement=SELECT COUNT(*) FROM schoolmgmt'
```

### View Service Logs
```bash
docker-compose logs student-service
docker-compose logs teacher-service
docker-compose logs academic-service
docker-compose logs achievement-service
docker-compose logs couchbase
```

### Direct Couchbase Data Access & Debugging
```bash
# Test if any documents exist in bucket
curl -u Administrator:password123 \
  -X POST http://localhost:8093/query/service \
  -d 'statement=SELECT COUNT(*) FROM schoolmgmt'

# List all document types
curl -u Administrator:password123 \
  -X POST http://localhost:8093/query/service \
  -d 'statement=SELECT type, COUNT(*) as count FROM schoolmgmt GROUP BY type'

# List all documents in bucket
curl -u Administrator:password123 \
  -X POST http://localhost:8093/query/service \
  -d 'statement=SELECT META().id, type FROM schoolmgmt'

# Get specific document types with correct structure
curl -u Administrator:password123 \
  -X POST http://localhost:8093/query/service \
  -d 'statement=SELECT * FROM schoolmgmt WHERE type = "student" LIMIT 5'

# Test document key access directly
curl -u Administrator:password123 \
  -X GET http://localhost:8091/pools/default/buckets/schoolmgmt/docs/student::your-student-id-here
```

## 🔍 Common Issues & Solutions

### Issue: "Failed to query students" Error - CRITICAL BUG IDENTIFIED
**Root Cause:** The Docker containers are running the OLD broken `main.go` instead of the NEW working `cmd/main.go`.

**The Problem:**
```dockerfile
# ❌ BROKEN: Dockerfile builds old main.go
RUN CGO_ENABLED=0 GOOS=linux go build -o student-service ./main.go

# ✅ SHOULD BE: Build new cmd/main.go  
RUN CGO_ENABLED=0 GOOS=linux go build -o student-service ./cmd/main.go
```

**Immediate Fix Required:**
1. **Update Dockerfile** in each service to build `./cmd/main.go`
2. **Rebuild containers**: `docker-compose build --no-cache`
3. **Restart services**: `docker-compose up -d`

**Files to Fix:**
- `services/student-service/Dockerfile`
- `services/teacher-service/Dockerfile` 
- `services/academic-service/Dockerfile`
- `services/achievement-service/Dockerfile`

**Quick Test - Services Using Repository Pattern:**
```bash
# These should work if using cmd/main.go:
curl http://localhost:8081/health
curl http://localhost:8082/health
curl http://localhost:8083/health
curl http://localhost:8084/health
```

**Expected Response from cmd/main.go services:**
```json
{
  "status": "healthy",
  "service": "student-service", 
  "couchbase_status": "connected",
  "database": "couchbase"
}
```

### Issue: Services can't connect to Couchbase
**Solution:**
```bash
# Check if Couchbase is initialized
curl http://localhost:8091/pools

# Re-initialize if needed
.\scripts\init-couchbase.ps1
```

### Issue: "Bucket not found" errors
**Solution:**
```bash
# Check bucket exists
curl -u Administrator:password123 http://localhost:8091/pools/default/buckets

# Create bucket manually if missing
curl -X POST http://localhost:8091/pools/default/buckets \
  -u Administrator:password123 \
  -d 'name=schoolmgmt&ramQuotaMB=256&bucketType=membase'
```

### Issue: Data not persisting
**Solution:**
```bash
# Check if indexes exist
curl -u Administrator:password123 \
  -X POST http://localhost:8093/query/service \
  -d 'statement=SELECT * FROM system:indexes WHERE keyspace_id = "schoolmgmt"'

# Create primary index if missing
curl -u Administrator:password123 \
  -X POST http://localhost:8093/query/service \
  -d 'statement=CREATE PRIMARY INDEX ON schoolmgmt'
```

## 📋 Expected Response Formats

### Successful Create Response
```json
{
  "success": true,
  "data": {
    "id": "uuid-generated-id",
    "firstName": "John",
    "lastName": "Doe",
    "email": "john.doe@school.edu",
    "createdAt": "2024-08-05T10:30:00Z",
    "updatedAt": "2024-08-05T10:30:00Z"
  },
  "message": "Student created successfully"
}
```

### Successful List Response
```json
{
  "success": true,
  "data": {
    "students": [...],
    "count": 10
  }
}
```

### Error Response
```json
{
  "error": "Student not found"
}
```

---

**Note:** Replace `{student_id}`, `{teacher_id}`, etc. with actual IDs returned from create operations. All endpoints use JSON for both request and response data.
