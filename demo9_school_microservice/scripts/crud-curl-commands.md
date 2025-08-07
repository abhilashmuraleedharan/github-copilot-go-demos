# CRUD cURL Commands for School Management System

This document provides working cURL commands for all microservices in the School Management System, based on the current Couchbase-integrated implementation.

## Service Endpoints

```bash
# Direct Service Access
STUDENT_SERVICE="http://localhost:8081"
TEACHER_SERVICE="http://localhost:8082"
ACADEMIC_SERVICE="http://localhost:8083"
ACHIEVEMENT_SERVICE="http://localhost:8084"
API_GATEWAY="http://localhost:8080"
```

## Health Checks

### Test All Service Health
```bash
# Check individual service health
curl http://localhost:8081/health
curl http://localhost:8082/health
curl http://localhost:8083/health
curl http://localhost:8084/health

# Expected response for all:
# {"database":"couchbase-connected","service":"[service-name]","status":"healthy"}
```

## Student Service (Port 8081)

### Student Data Structure
```json
{
  "id": "generated-uuid",
  "firstName": "John",
  "lastName": "Doe",
  "email": "john.doe@school.edu",
  "grade": "10",
  "status": "active",
  "createdAt": "2025-08-05T18:00:00Z",
  "updatedAt": "2025-08-05T18:00:00Z",
  "type": "student"
}
```

### Create Student
```bash
# Create a new student
curl -X POST http://localhost:8081/students \
  -H "Content-Type: application/json" \
  -d '{
    "firstName": "Jane",
    "lastName": "Smith",
    "email": "jane.smith@school.edu",
    "grade": "11",
    "status": "active"
  }'

# Create another student
curl -X POST http://localhost:8081/students \
  -H "Content-Type: application/json" \
  -d '{
    "firstName": "Mike",
    "lastName": "Johnson",
    "email": "mike.johnson@school.edu",
    "grade": "12",
    "status": "active"
  }'
```

### Read Students
```bash
# Get all students
curl http://localhost:8081/students

# Get specific student by ID (replace with actual ID from create response)
curl http://localhost:8081/students/{student-id}
```

### Update Student
```bash
# Update student (replace {student-id} with actual ID)
curl -X PUT http://localhost:8081/students/{student-id} \
  -H "Content-Type: application/json" \
  -d '{
    "firstName": "Jane",
    "lastName": "Smith-Updated",
    "email": "jane.smith.updated@school.edu",
    "grade": "12",
    "status": "active"
  }'
```

### Delete Student
```bash
# Delete student (replace {student-id} with actual ID)
curl -X DELETE http://localhost:8081/students/{student-id}
```

## Teacher Service (Port 8082)

### Teacher Data Structure
```json
{
  "id": "generated-uuid",
  "firstName": "Dr. Sarah",
  "lastName": "Wilson",
  "email": "sarah.wilson@school.edu",
  "phone": "555-1234",
  "department": "Mathematics",
  "subjects": ["Algebra", "Calculus"],
  "qualification": "PhD Mathematics",
  "experience": 8,
  "status": "active",
  "createdAt": "2025-08-05T18:00:00Z",
  "updatedAt": "2025-08-05T18:00:00Z",
  "type": "teacher"
}
```

### Create Teacher
```bash
# Create a new teacher
curl -X POST http://localhost:8082/teachers \
  -H "Content-Type: application/json" \
  -d '{
    "firstName": "Dr. Sarah",
    "lastName": "Wilson",
    "email": "sarah.wilson@school.edu",
    "phone": "555-2001",
    "department": "Mathematics",
    "subjects": ["Algebra", "Calculus"],
    "qualification": "PhD Mathematics",
    "experience": 8,
    "status": "active"
  }'

# Create another teacher
curl -X POST http://localhost:8082/teachers \
  -H "Content-Type: application/json" \
  -d '{
    "firstName": "John",
    "lastName": "Parker",
    "email": "john.parker@school.edu",
    "phone": "555-2002",
    "department": "English",
    "subjects": ["Literature", "Creative Writing"],
    "qualification": "MA English Literature",
    "experience": 5,
    "status": "active"
  }'
```

### Read Teachers
```bash
# Get all teachers
curl http://localhost:8082/teachers

# Get specific teacher by ID
curl http://localhost:8082/teachers/{teacher-id}
```

### Update Teacher
```bash
# Update teacher
curl -X PUT http://localhost:8082/teachers/{teacher-id} \
  -H "Content-Type: application/json" \
  -d '{
    "firstName": "Dr. Sarah",
    "lastName": "Wilson-Updated",
    "email": "sarah.wilson.updated@school.edu",
    "phone": "555-2001",
    "department": "Mathematics",
    "subjects": ["Advanced Algebra", "Calculus", "Statistics"],
    "qualification": "PhD Mathematics",
    "experience": 9,
    "status": "active"
  }'
```

### Delete Teacher
```bash
# Delete teacher
curl -X DELETE http://localhost:8082/teachers/{teacher-id}
```

## Academic Service (Port 8083)

### Subject Data Structure
```json
{
  "id": "generated-uuid",
  "name": "Computer Science 101",
  "code": "CS101",
  "credits": 3,
  "description": "Introduction to Computer Science",
  "teacherId": "teacher-uuid",
  "status": "active",
  "createdAt": "2025-08-05T18:00:00Z",
  "updatedAt": "2025-08-05T18:00:00Z",
  "type": "subject"
}
```

### Create Subject
```bash
# Create a new subject
curl -X POST http://localhost:8083/subjects \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Computer Science 101",
    "code": "CS101",
    "credits": 3,
    "description": "Introduction to Computer Science",
    "teacherId": "teacher-001",
    "status": "active"
  }'

# Create another subject
curl -X POST http://localhost:8083/subjects \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Advanced Mathematics",
    "code": "MATH201",
    "credits": 4,
    "description": "Advanced calculus and linear algebra",
    "teacherId": "teacher-002",
    "status": "active"
  }'
```

### Read Subjects
```bash
# Get all subjects
curl http://localhost:8083/subjects

# Get specific subject by ID
curl http://localhost:8083/subjects/{subject-id}
```

### Update Subject
```bash
# Update subject
curl -X PUT http://localhost:8083/subjects/{subject-id} \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Computer Science 101 - Updated",
    "code": "CS101",
    "credits": 4,
    "description": "Introduction to Computer Science with practical labs",
    "teacherId": "teacher-001",
    "status": "active"
  }'
```

### Delete Subject
```bash
# Delete subject
curl -X DELETE http://localhost:8083/subjects/{subject-id}
```

## Achievement Service (Port 8084)

### Achievement Data Structure
```json
{
  "id": "generated-uuid",
  "title": "Dean's List",
  "description": "Achieved GPA above 3.75",
  "category": "academic",
  "points": 150,
  "studentId": "student-uuid",
  "teacherId": "teacher-uuid",
  "status": "awarded",
  "createdAt": "2025-08-05T18:00:00Z",
  "updatedAt": "2025-08-05T18:00:00Z",
  "type": "achievement"
}
```

### Create Achievement
```bash
# Create a new achievement
curl -X POST http://localhost:8084/achievements \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Dean'\''s List",
    "description": "Achieved GPA above 3.75 for Fall 2024 semester",
    "category": "academic",
    "points": 150,
    "studentId": "student-001",
    "teacherId": "teacher-001",
    "status": "awarded"
  }'

# Create sports achievement
curl -X POST http://localhost:8084/achievements \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Basketball Championship",
    "description": "Won regional basketball championship",
    "category": "sports",
    "points": 200,
    "studentId": "student-002",
    "teacherId": "teacher-002",
    "status": "awarded"
  }'
```

### Read Achievements
```bash
# Get all achievements
curl http://localhost:8084/achievements

# Get specific achievement by ID
curl http://localhost:8084/achievements/{achievement-id}
```

### Update Achievement
```bash
# Update achievement
curl -X PUT http://localhost:8084/achievements/{achievement-id} \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Dean'\''s List - Fall 2024",
    "description": "Achieved GPA above 3.75 for Fall 2024 semester with honors",
    "category": "academic",
    "points": 175,
    "studentId": "student-001",
    "teacherId": "teacher-001",
    "status": "awarded"
  }'
```

### Delete Achievement
```bash
# Delete achievement
curl -X DELETE http://localhost:8084/achievements/{achievement-id}
```

## PowerShell Examples (Windows)

### Health Check All Services
```powershell
# Check all service health
$services = @(8081, 8082, 8083, 8084)
foreach ($port in $services) {
    $response = Invoke-RestMethod -Uri "http://localhost:$port/health"
    Write-Output "Port $port - Service: $($response.service) - Status: $($response.status) - Database: $($response.database)"
}
```

### Create Student (PowerShell)
```powershell
$studentData = @{
    firstName = "Alice"
    lastName = "Brown"
    email = "alice.brown@school.edu"
    grade = "10"
    status = "active"
} | ConvertTo-Json

$response = Invoke-RestMethod -Uri "http://localhost:8081/students" -Method POST -Body $studentData -ContentType "application/json"
Write-Output "Created student with ID: $($response.id)"
```

### Create Teacher (PowerShell)
```powershell
$teacherData = @{
    firstName = "Dr. Emily"
    lastName = "Davis"
    email = "emily.davis@school.edu"
    phone = "555-3001"
    department = "Science"
    subjects = @("Biology", "Chemistry")
    qualification = "PhD Biology"
    experience = 10
    status = "active"
} | ConvertTo-Json

$response = Invoke-RestMethod -Uri "http://localhost:8082/teachers" -Method POST -Body $teacherData -ContentType "application/json"
Write-Output "Created teacher with ID: $($response.id)"
```

### Create Subject (PowerShell)
```powershell
$subjectData = @{
    name = "Biology Advanced"
    code = "BIO301"
    credits = 4
    description = "Advanced biology with lab work"
    teacherId = "teacher-001"
    status = "active"
} | ConvertTo-Json

$response = Invoke-RestMethod -Uri "http://localhost:8083/subjects" -Method POST -Body $subjectData -ContentType "application/json"
Write-Output "Created subject with ID: $($response.id)"
```

### Create Achievement (PowerShell)
```powershell
$achievementData = @{
    title = "Science Fair Winner"
    description = "First place in regional science fair"
    category = "academic"
    points = 300
    studentId = "student-001"
    teacherId = "teacher-001"
    status = "awarded"
} | ConvertTo-Json

$response = Invoke-RestMethod -Uri "http://localhost:8084/achievements" -Method POST -Body $achievementData -ContentType "application/json"
Write-Output "Created achievement with ID: $($response.id)"
```

## Complete Test Workflow

### Step 1: Test All Health Endpoints
```bash
echo "Testing health endpoints..."
curl -s http://localhost:8081/health | jq .
curl -s http://localhost:8082/health | jq .
curl -s http://localhost:8083/health | jq .
curl -s http://localhost:8084/health | jq .
```

### Step 2: Create Test Data
```bash
# Create a student
STUDENT_RESPONSE=$(curl -s -X POST http://localhost:8081/students \
  -H "Content-Type: application/json" \
  -d '{
    "firstName": "Test",
    "lastName": "Student",
    "email": "test.student@school.edu",
    "grade": "10",
    "status": "active"
  }')

STUDENT_ID=$(echo $STUDENT_RESPONSE | jq -r '.id')
echo "Created student with ID: $STUDENT_ID"

# Create a teacher
TEACHER_RESPONSE=$(curl -s -X POST http://localhost:8082/teachers \
  -H "Content-Type: application/json" \
  -d '{
    "firstName": "Test",
    "lastName": "Teacher",
    "email": "test.teacher@school.edu",
    "phone": "555-0000",
    "department": "Test Department",
    "subjects": ["Test Subject"],
    "qualification": "Test Qualification",
    "experience": 1,
    "status": "active"
  }')

TEACHER_ID=$(echo $TEACHER_RESPONSE | jq -r '.id')
echo "Created teacher with ID: $TEACHER_ID"

# Create a subject
SUBJECT_RESPONSE=$(curl -s -X POST http://localhost:8083/subjects \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Test Subject",
    "code": "TEST101",
    "credits": 3,
    "description": "Test subject for demonstration",
    "teacherId": "'$TEACHER_ID'",
    "status": "active"
  }')

SUBJECT_ID=$(echo $SUBJECT_RESPONSE | jq -r '.id')
echo "Created subject with ID: $SUBJECT_ID"

# Create an achievement
ACHIEVEMENT_RESPONSE=$(curl -s -X POST http://localhost:8084/achievements \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Test Achievement",
    "description": "Test achievement for demonstration",
    "category": "test",
    "points": 100,
    "studentId": "'$STUDENT_ID'",
    "teacherId": "'$TEACHER_ID'",
    "status": "awarded"
  }')

ACHIEVEMENT_ID=$(echo $ACHIEVEMENT_RESPONSE | jq -r '.id')
echo "Created achievement with ID: $ACHIEVEMENT_ID"
```

### Step 3: Test Read Operations
```bash
echo "Testing read operations..."
curl -s http://localhost:8081/students | jq .
curl -s http://localhost:8082/teachers | jq .
curl -s http://localhost:8083/subjects | jq .
curl -s http://localhost:8084/achievements | jq .
```

### Step 4: Test Update Operations
```bash
# Update student
curl -X PUT http://localhost:8081/students/$STUDENT_ID \
  -H "Content-Type: application/json" \
  -d '{
    "firstName": "Updated",
    "lastName": "Student",
    "email": "updated.student@school.edu",
    "grade": "11",
    "status": "active"
  }'

# Update teacher
curl -X PUT http://localhost:8082/teachers/$TEACHER_ID \
  -H "Content-Type: application/json" \
  -d '{
    "firstName": "Updated",
    "lastName": "Teacher",
    "email": "updated.teacher@school.edu",
    "phone": "555-0001",
    "department": "Updated Department",
    "subjects": ["Updated Subject"],
    "qualification": "Updated Qualification",
    "experience": 2,
    "status": "active"
  }'
```

### Step 5: Test Delete Operations
```bash
# Delete in reverse order to maintain referential integrity
curl -X DELETE http://localhost:8084/achievements/$ACHIEVEMENT_ID
curl -X DELETE http://localhost:8083/subjects/$SUBJECT_ID
curl -X DELETE http://localhost:8082/teachers/$TEACHER_ID
curl -X DELETE http://localhost:8081/students/$STUDENT_ID
```

## Expected Responses

### Successful Create Response
```json
{
  "id": "generated-uuid",
  "firstName": "John",
  "lastName": "Doe",
  "email": "john.doe@school.edu",
  "grade": "10",
  "status": "active",
  "createdAt": "2025-08-05T18:00:00Z",
  "updatedAt": "2025-08-05T18:00:00Z",
  "type": "student"
}
```

### Health Check Response
```json
{
  "database": "couchbase-connected",
  "service": "student-service",
  "status": "healthy"
}
```

### Error Responses
```json
// 400 Bad Request
{
  "error": "validation error message"
}

// 404 Not Found
{
  "error": "Resource not found"
}

// 500 Internal Server Error
{
  "error": "Internal server error message"
}
```

## Troubleshooting

### Common Issues
1. **404 Not Found**: Endpoint may not be registered - check service logs
2. **500 Internal Server Error**: Database connection issue - verify Couchbase is running
3. **Connection Refused**: Service may not be running - check `docker-compose ps`

### Debug Commands
```bash
# Check service logs
docker logs schoolmgmt-student-service
docker logs schoolmgmt-teacher-service
docker logs schoolmgmt-academic-service
docker logs schoolmgmt-achievement-service

# Check service status
docker-compose ps

# Check Couchbase status
curl http://localhost:8091
```

All services are now connected to Couchbase and ready for CRUD operations!