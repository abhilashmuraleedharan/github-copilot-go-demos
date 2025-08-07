# 📡 Complete CRUD cURL Commands for Couchbase-backed Microservices

This comprehensive guide provides detailed cURL commands for testing all CRUD operations on the School Management System's Couchbase-integrated microservices.

## 🔧 Prerequisites

1. **Start Services:**
   ```bash
   docker-compose up -d
   ```

2. **Initialize Couchbase:**
   ```bash
   # Linux/macOS
   ./scripts/init-couchbase.sh
   
   # Windows
   .\scripts\init-couchbase.ps1
   ```

3. **Verify Services:**
   ```bash
   curl http://localhost:8080/health  # API Gateway
   curl http://localhost:8081/health  # Student Service
   curl http://localhost:8082/health  # Teacher Service
   curl http://localhost:8083/health  # Academic Service
   curl http://localhost:8084/health  # Achievement Service
   ```

---

## 🧑‍🎓 Student Service - Full CRUD with All Fields

### ✅ Create Student (Complete)
```bash
curl -X POST http://localhost:8080/api/v1/students \
  -H "Content-Type: application/json" \
  -d '{
    "firstName": "John",
    "lastName": "Doe",
    "email": "john.doe@school.edu",
    "dateOfBirth": "2005-03-15",
    "grade": "10",
    "address": "123 Main Street, Springfield, IL 62701",
    "phone": "555-1234",
    "parentName": "Jane Doe",
    "parentPhone": "555-5678"
  }'
```

### ✅ Create Student (Minimal)
```bash
curl -X POST http://localhost:8080/api/v1/students \
  -H "Content-Type: application/json" \
  -d '{
    "firstName": "Alice", 
    "lastName": "Smith",
    "email": "alice.smith@school.edu",
    "grade": "11"
  }'
```

### 📋 Get All Students
```bash
curl http://localhost:8080/api/v1/students
```

### 🔍 Get Student by ID
```bash
# Replace {student_id} with actual ID from create response
curl http://localhost:8080/api/v1/students/{student_id}
```

### ✏️ Update Student (Complete)
```bash
curl -X PUT http://localhost:8080/api/v1/students/{student_id} \
  -H "Content-Type: application/json" \
  -d '{
    "firstName": "Johnny",
    "lastName": "Doe-Updated",
    "email": "johnny.doe.updated@school.edu",
    "dateOfBirth": "2005-03-15",
    "grade": "11",
    "address": "456 Oak Avenue, Springfield, IL 62701",
    "phone": "555-9999",
    "parentName": "Jane Doe-Updated",
    "parentPhone": "555-8888"
  }'
```

### ✏️ Update Student (Partial)
```bash
curl -X PATCH http://localhost:8080/api/v1/students/{student_id} \
  -H "Content-Type: application/json" \
  -d '{
    "grade": "12",
    "email": "promoted.student@school.edu"
  }'
```

### 🗑️ Delete Student
```bash
curl -X DELETE http://localhost:8080/api/v1/students/{student_id}
```

---

## 👩‍🏫 Teacher Service - Full CRUD with Couchbase Backend

### ✅ Create Teacher
```bash
curl -X POST http://localhost:8080/api/v1/teachers \
  -H "Content-Type: application/json" \
  -d '{
    "firstName": "Dr. Sarah",
    "lastName": "Johnson",
    "email": "sarah.johnson@school.edu",
    "phone": "555-9876",
    "department": "Mathematics",
    "subjects": ["Algebra", "Calculus", "Statistics"],
    "qualification": "Ph.D. in Mathematics",
    "experience": 15,
    "status": "active"
  }'
```

### 📄 Get Teacher by ID
```bash
# Replace {teacher_id} with actual ID from create response
curl -X GET http://localhost:8080/api/v1/teachers/{teacher_id}
```

### 📋 List All Teachers
```bash
# List with default pagination
curl -X GET http://localhost:8080/api/v1/teachers

# List with custom pagination
curl -X GET "http://localhost:8080/api/v1/teachers?limit=10&offset=0"
```

### 🔄 Update Teacher
```bash
curl -X PUT http://localhost:8080/api/v1/teachers/{teacher_id} \
  -H "Content-Type: application/json" \
  -d '{
    "phone": "555-9999",
    "experience": 16,
    "subjects": ["Advanced Calculus", "Linear Algebra", "Statistics"]
  }'
```

### ❌ Delete Teacher
```bash
curl -X DELETE http://localhost:8080/api/v1/teachers/{teacher_id}
```

### 📊 Get Teachers by Department
```bash
curl -X GET http://localhost:8080/api/v1/teachers/department/Mathematics
```

### 🟢 Get Active Teachers
```bash
curl -X GET http://localhost:8080/api/v1/teachers/active
```

---

## 📚 Academic Service - Full CRUD with Couchbase Backend

### ✅ Create Academic Record
```bash
curl -X POST http://localhost:8080/api/v1/academics \
  -H "Content-Type: application/json" \
  -d '{
    "studentId": "student_1641234567890",
    "teacherId": "teacher_1641234567890",
    "subject": "Mathematics",
    "grade": "10",
    "semester": "Fall 2024",
    "academicYear": "2024-2025",
    "examType": "Midterm",
    "examDate": "2024-03-15T10:00:00Z",
    "maxMarks": 100,
    "obtainedMarks": 85,
    "remarks": "Excellent performance"
  }'
```

### 📄 Get Academic Record by ID
```bash
# Replace {academic_id} with actual ID from create response
curl -X GET http://localhost:8080/api/v1/academics/{academic_id}
```

### 📋 List All Academic Records
```bash
# List with default pagination
curl -X GET http://localhost:8080/api/v1/academics

# List with custom pagination
curl -X GET "http://localhost:8080/api/v1/academics?limit=20&offset=0"
```

### 🔄 Update Academic Record
```bash
curl -X PUT http://localhost:8080/api/v1/academics/{academic_id} \
  -H "Content-Type: application/json" \
  -d '{
    "obtainedMarks": 90,
    "remarks": "Outstanding improvement"
  }'
```

### ❌ Delete Academic Record
```bash
curl -X DELETE http://localhost:8080/api/v1/academics/{academic_id}
```

### 🧑‍🎓 Get Academic Records by Student
```bash
curl -X GET http://localhost:8080/api/v1/academics/student/{student_id}
```

### ✅ Create Class
```bash
curl -X POST http://localhost:8080/api/v1/classes \
  -H "Content-Type: application/json" \
  -d '{
    "className": "Advanced Mathematics",
    "grade": "11",
    "section": "A",
    "teacherId": "teacher_1641234567890",
    "subject": "Calculus",
    "academicYear": "2024-2025",
    "semester": "Fall",
    "studentIds": ["student_1641234567890", "student_1641234567891"],
    "maxCapacity": 30,
    "status": "active"
  }'
```

### 📄 Get Class by ID
```bash
curl -X GET http://localhost:8080/api/v1/classes/{class_id}
```

### 📋 List All Classes
```bash
curl -X GET http://localhost:8080/api/v1/classes
```

---

## 🏆 Achievement Service - Couchbase Integration (To Be Updated)

**Note:** Achievement service still needs to be updated with Couchbase integration following the same pattern as Student, Teacher, and Academic services.

### Current Endpoints (In-Memory):
```bash
# Create Achievement
curl -X POST http://localhost:8080/api/v1/achievements \
  -H "Content-Type: application/json" \
  -d '{
    "studentId": "student-1",
    "teacherId": "teacher-1",
    "title": "Science Fair Winner",
    "description": "First place in school science fair",
    "category": "academic",
    "type": "award",
    "level": "school",
    "date": "2024-03-15T00:00:00Z",
    "points": 50
  }'

# Get Achievement
curl -X GET http://localhost:8080/api/v1/achievements/{achievement_id}

# List Achievements
curl -X GET http://localhost:8080/api/v1/achievements

# Get Student Achievements
curl -X GET http://localhost:8080/api/v1/achievements/student/{student_id}
```

---

## 🔍 Advanced Couchbase N1QL Queries

### Direct Couchbase Query Interface

Access the Couchbase Query Interface at: http://localhost:8093

#### Query All Students
```sql
SELECT s.* FROM schoolmgmt s WHERE s.type = "student" LIMIT 10;
```

#### Query All Teachers by Department
```sql
SELECT t.* FROM schoolmgmt t 
WHERE t.type = "teacher" AND t.department = "Mathematics";
```

#### Query Academic Records with Percentage > 80
```sql
SELECT a.* FROM schoolmgmt a 
WHERE a.type = "academic" AND a.percentage > 80;
```

#### Cross-Service Query: Students with Their Academic Records
```sql
SELECT s.firstName, s.lastName, a.subject, a.percentage
FROM schoolmgmt s
JOIN schoolmgmt a ON s.id = a.student_id
WHERE s.type = "student" AND a.type = "academic"
ORDER BY a.percentage DESC
LIMIT 20;
```

#### Performance Analytics Query
```sql
SELECT 
  a.subject,
  COUNT(*) as total_records,
  AVG(a.percentage) as avg_percentage,
  MAX(a.percentage) as max_percentage,
  MIN(a.percentage) as min_percentage
FROM schoolmgmt a
WHERE a.type = "academic"
GROUP BY a.subject
ORDER BY avg_percentage DESC;
```

---

## 🧪 Complete Test Workflow

### 1. Create Full Academic Ecosystem
```bash
# 1. Create a teacher
TEACHER_RESPONSE=$(curl -s -X POST http://localhost:8080/api/v1/teachers \
  -H "Content-Type: application/json" \
  -d '{
    "firstName": "Dr. Emily",
    "lastName": "Rodriguez",
    "email": "emily.rodriguez@school.edu",
    "phone": "555-2468",
    "department": "Science",
    "subjects": ["Biology", "Chemistry"],
    "qualification": "Ph.D. in Biology",
    "experience": 12,
    "status": "active"
  }')

echo "Teacher created: $TEACHER_RESPONSE"
TEACHER_ID=$(echo $TEACHER_RESPONSE | jq -r '.data.id')

# 2. Create a student
STUDENT_RESPONSE=$(curl -s -X POST http://localhost:8080/api/v1/students \
  -H "Content-Type: application/json" \
  -d '{
    "firstName": "Alice",
    "lastName": "Smith",
    "email": "alice.smith@school.edu",
    "dateOfBirth": "2006-05-20",
    "grade": "11",
    "address": "456 Oak Avenue, Springfield, IL 62702",
    "phone": "555-3579",
    "parentName": "Robert Smith",
    "parentPhone": "555-4680"
  }')

echo "Student created: $STUDENT_RESPONSE"
STUDENT_ID=$(echo $STUDENT_RESPONSE | jq -r '.data.id')

# 3. Create academic record
curl -X POST http://localhost:8080/api/v1/academics \
  -H "Content-Type: application/json" \
  -d "{
    \"studentId\": \"$STUDENT_ID\",
    \"teacherId\": \"$TEACHER_ID\",
    \"subject\": \"Biology\",
    \"grade\": \"11\",
    \"semester\": \"Fall 2024\",
    \"academicYear\": \"2024-2025\",
    \"examType\": \"Final\",
    \"maxMarks\": 100,
    \"obtainedMarks\": 92,
    \"remarks\": \"Excellent understanding of molecular biology\"
  }"

# 4. Create a class
curl -X POST http://localhost:8080/api/v1/classes \
  -H "Content-Type: application/json" \
  -d "{
    \"className\": \"Advanced Biology\",
    \"grade\": \"11\",
    \"section\": \"A\",
    \"teacherId\": \"$TEACHER_ID\",
    \"subject\": \"Biology\",
    \"academicYear\": \"2024-2025\",
    \"semester\": \"Fall\",
    \"studentIds\": [\"$STUDENT_ID\"],
    \"maxCapacity\": 25,
    \"status\": \"active\"
  }"
```

### 2. Verify Data Persistence
```bash
# Restart services to test persistence
docker-compose restart student-service teacher-service academic-service

# Wait for services to start
sleep 30

# Verify data still exists
curl http://localhost:8080/api/v1/students
curl http://localhost:8080/api/v1/teachers
curl http://localhost:8080/api/v1/academics
```

---

## 🛠️ Troubleshooting Commands

### Check Service Health
```bash
curl http://localhost:8080/health
curl http://localhost:8081/health
curl http://localhost:8082/health
curl http://localhost:8083/health
curl http://localhost:8084/health
```

### Check Couchbase Status
```bash
curl http://localhost:8091/pools/default
```

### View Service Logs
```bash
docker-compose logs student-service
docker-compose logs teacher-service
docker-compose logs academic-service
docker-compose logs couchbase
```

### Test Couchbase Connection
```bash
curl -u Administrator:password123 \
  -X POST http://localhost:8093/query/service \
  -d 'statement=SELECT COUNT(*) FROM schoolmgmt'
```

---

## 📝 Notes

1. **Service Status:** ✅ Student, ✅ Teacher, ✅ Academic services now use Couchbase
2. **Achievement Service:** 🔄 Still needs Couchbase integration
3. **Data Persistence:** All data is stored in Couchbase and survives container restarts
4. **Error Handling:** Services return appropriate HTTP status codes and error messages
5. **Logging:** Detailed logging is available in service containers for debugging

For more detailed troubleshooting, see `FIXES_AND_SOLUTIONS.md`.
