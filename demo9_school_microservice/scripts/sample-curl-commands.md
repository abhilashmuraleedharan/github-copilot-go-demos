# Sample cURL Commands for School Management System

This document provides comprehensive cURL commands for testing all CRUD operations across all microservices.

## Base URLs

```bash
# API Gateway (Main Entry Point)
GATEWAY_URL="http://localhost:8080"

# Individual Services (Direct Access)
STUDENT_SERVICE_URL="http://localhost:8081"
TEACHER_SERVICE_URL="http://localhost:8082"
ACADEMIC_SERVICE_URL="http://localhost:8083"
ACHIEVEMENT_SERVICE_URL="http://localhost:8084"
```

## Health Checks

### Service Health Status
```bash
# Check API Gateway health
curl -X GET "$GATEWAY_URL/health"

# Check all services health
curl -X GET "$GATEWAY_URL/api/health"
curl -X GET "$STUDENT_SERVICE_URL/health"
curl -X GET "$TEACHER_SERVICE_URL/health"
curl -X GET "$ACADEMIC_SERVICE_URL/health"
curl -X GET "$ACHIEVEMENT_SERVICE_URL/health"
```

## Student Service CRUD Operations

### Create Student
```bash
# Create a new student
curl -X POST "$GATEWAY_URL/api/students" \
  -H "Content-Type: application/json" \
  -d '{
    "firstName": "Jane",
    "lastName": "Smith",
    "email": "jane.smith@school.edu",
    "dateOfBirth": "2005-08-20",
    "grade": "10",
    "enrollmentDate": "2024-09-01",
    "status": "active",
    "address": {
      "street": "456 Oak Avenue",
      "city": "Springfield",
      "state": "IL",
      "zipCode": "62702"
    },
    "parentContact": {
      "name": "Robert Smith",
      "phone": "555-0102",
      "email": "robert.smith@email.com"
    }
  }'

# Create student with minimal data
curl -X POST "$GATEWAY_URL/api/students" \
  -H "Content-Type: application/json" \
  -d '{
    "firstName": "Mike",
    "lastName": "Johnson",
    "email": "mike.johnson@school.edu",
    "grade": "11",
    "status": "active"
  }'
```

### Read Students
```bash
# Get all students
curl -X GET "$GATEWAY_URL/api/students"

# Get student by ID
curl -X GET "$GATEWAY_URL/api/students/student-001"

# Search students by grade
curl -X GET "$GATEWAY_URL/api/students?grade=10"

# Search students by status
curl -X GET "$GATEWAY_URL/api/students?status=active"

# Get students with pagination
curl -X GET "$GATEWAY_URL/api/students?page=1&limit=10"
```

### Update Student
```bash
# Update student information
curl -X PUT "$GATEWAY_URL/api/students/student-001" \
  -H "Content-Type: application/json" \
  -d '{
    "firstName": "John",
    "lastName": "Doe",
    "email": "john.doe.updated@school.edu",
    "dateOfBirth": "2005-03-15",
    "grade": "11",
    "enrollmentDate": "2023-09-01",
    "status": "active",
    "address": {
      "street": "123 Main Street Updated",
      "city": "Springfield",
      "state": "IL",
      "zipCode": "62701"
    },
    "parentContact": {
      "name": "Jane Doe",
      "phone": "555-0101",
      "email": "jane.doe.updated@email.com"
    }
  }'

# Partial update (PATCH)
curl -X PATCH "$GATEWAY_URL/api/students/student-001" \
  -H "Content-Type: application/json" \
  -d '{
    "grade": "12",
    "status": "active"
  }'
```

### Delete Student
```bash
# Delete student by ID
curl -X DELETE "$GATEWAY_URL/api/students/student-001"

# Soft delete (change status to inactive)
curl -X PATCH "$GATEWAY_URL/api/students/student-001" \
  -H "Content-Type: application/json" \
  -d '{
    "status": "inactive"
  }'
```

## Teacher Service CRUD Operations

### Create Teacher
```bash
# Create a new teacher
curl -X POST "$GATEWAY_URL/api/teachers" \
  -H "Content-Type: application/json" \
  -d '{
    "firstName": "Dr. Sarah",
    "lastName": "Wilson",
    "email": "sarah.wilson@school.edu",
    "department": "Mathematics",
    "subject": "Calculus",
    "hireDate": "2020-08-15",
    "status": "active",
    "qualifications": [
      "PhD Mathematics",
      "MS Applied Mathematics",
      "BA Mathematics"
    ],
    "officeHours": "Mon-Wed-Fri 2:00-4:00 PM",
    "phone": "555-2001"
  }'

# Create teacher with minimal data
curl -X POST "$GATEWAY_URL/api/teachers" \
  -H "Content-Type: application/json" \
  -d '{
    "firstName": "John",
    "lastName": "Parker",
    "email": "john.parker@school.edu",
    "department": "English",
    "subject": "Creative Writing",
    "status": "active"
  }'
```

### Read Teachers
```bash
# Get all teachers
curl -X GET "$GATEWAY_URL/api/teachers"

# Get teacher by ID
curl -X GET "$GATEWAY_URL/api/teachers/teacher-001"

# Search teachers by department
curl -X GET "$GATEWAY_URL/api/teachers?department=Mathematics"

# Search teachers by subject
curl -X GET "$GATEWAY_URL/api/teachers?subject=Biology"

# Get active teachers only
curl -X GET "$GATEWAY_URL/api/teachers?status=active"
```

### Update Teacher
```bash
# Update teacher information
curl -X PUT "$GATEWAY_URL/api/teachers/teacher-001" \
  -H "Content-Type: application/json" \
  -d '{
    "firstName": "Dr. Robert",
    "lastName": "Anderson",
    "email": "robert.anderson.updated@school.edu",
    "department": "Mathematics",
    "subject": "Advanced Algebra",
    "hireDate": "2015-08-15",
    "status": "active",
    "qualifications": [
      "PhD Mathematics",
      "MEd Education",
      "MS Statistics"
    ],
    "officeHours": "Mon-Wed-Fri 1:00-3:00 PM",
    "phone": "555-1001"
  }'

# Update office hours only
curl -X PATCH "$GATEWAY_URL/api/teachers/teacher-001" \
  -H "Content-Type: application/json" \
  -d '{
    "officeHours": "Tue-Thu 3:00-5:00 PM"
  }'
```

### Delete Teacher
```bash
# Delete teacher by ID
curl -X DELETE "$GATEWAY_URL/api/teachers/teacher-001"

# Change teacher status to inactive
curl -X PATCH "$GATEWAY_URL/api/teachers/teacher-001" \
  -H "Content-Type: application/json" \
  -d '{
    "status": "inactive"
  }'
```

## Academic Service CRUD Operations

### Create Course
```bash
# Create a new course
curl -X POST "$GATEWAY_URL/api/courses" \
  -H "Content-Type: application/json" \
  -d '{
    "courseCode": "PHYS101",
    "courseName": "Introduction to Physics",
    "department": "Science",
    "teacherId": "teacher-003",
    "credits": 1.0,
    "semester": "Spring 2025",
    "capacity": 20,
    "enrolledStudents": [],
    "schedule": {
      "days": ["Monday", "Wednesday", "Friday"],
      "time": "10:00-11:00",
      "room": "Lab-102"
    }
  }'

# Create course with enrolled students
curl -X POST "$GATEWAY_URL/api/courses" \
  -H "Content-Type: application/json" \
  -d '{
    "courseCode": "ART201",
    "courseName": "Digital Art",
    "department": "Arts",
    "teacherId": "teacher-002",
    "credits": 0.5,
    "semester": "Spring 2025",
    "capacity": 15,
    "enrolledStudents": ["student-001", "student-002"],
    "schedule": {
      "days": ["Tuesday", "Thursday"],
      "time": "14:00-15:30",
      "room": "Art-201"
    }
  }'
```

### Read Courses
```bash
# Get all courses
curl -X GET "$GATEWAY_URL/api/courses"

# Get course by ID
curl -X GET "$GATEWAY_URL/api/courses/course-001"

# Search courses by department
curl -X GET "$GATEWAY_URL/api/courses?department=Mathematics"

# Search courses by semester
curl -X GET "$GATEWAY_URL/api/courses?semester=Fall 2024"

# Search courses by teacher
curl -X GET "$GATEWAY_URL/api/courses?teacherId=teacher-001"

# Get courses with available spots
curl -X GET "$GATEWAY_URL/api/courses?hasAvailableSpots=true"
```

### Update Course
```bash
# Update course information
curl -X PUT "$GATEWAY_URL/api/courses/course-001" \
  -H "Content-Type: application/json" \
  -d '{
    "courseCode": "MATH101",
    "courseName": "Algebra I - Advanced",
    "department": "Mathematics",
    "teacherId": "teacher-001",
    "credits": 1.0,
    "semester": "Fall 2024",
    "capacity": 30,
    "enrolledStudents": ["student-001", "student-002", "student-005"],
    "schedule": {
      "days": ["Monday", "Wednesday", "Friday"],
      "time": "09:00-10:00",
      "room": "Math-101"
    }
  }'

# Add student to course
curl -X PATCH "$GATEWAY_URL/api/courses/course-001/enroll" \
  -H "Content-Type: application/json" \
  -d '{
    "studentId": "student-003"
  }'

# Remove student from course
curl -X PATCH "$GATEWAY_URL/api/courses/course-001/unenroll" \
  -H "Content-Type: application/json" \
  -d '{
    "studentId": "student-001"
  }'
```

### Delete Course
```bash
# Delete course by ID
curl -X DELETE "$GATEWAY_URL/api/courses/course-001"
```

### Grades Management
```bash
# Create a grade entry
curl -X POST "$GATEWAY_URL/api/grades" \
  -H "Content-Type: application/json" \
  -d '{
    "studentId": "student-001",
    "courseId": "course-001",
    "teacherId": "teacher-001",
    "assignmentName": "Final Exam",
    "grade": "A",
    "points": 95,
    "maxPoints": 100,
    "dateGraded": "2024-12-15",
    "semester": "Fall 2024",
    "category": "exam"
  }'

# Get all grades
curl -X GET "$GATEWAY_URL/api/grades"

# Get grades by student
curl -X GET "$GATEWAY_URL/api/grades?studentId=student-001"

# Get grades by course
curl -X GET "$GATEWAY_URL/api/grades?courseId=course-001"

# Get grades by semester
curl -X GET "$GATEWAY_URL/api/grades?semester=Fall 2024"

# Update grade
curl -X PUT "$GATEWAY_URL/api/grades/grade-001" \
  -H "Content-Type: application/json" \
  -d '{
    "studentId": "student-001",
    "courseId": "course-001",
    "teacherId": "teacher-001",
    "assignmentName": "Midterm Exam - Corrected",
    "grade": "B+",
    "points": 88,
    "maxPoints": 100,
    "dateGraded": "2024-10-15",
    "semester": "Fall 2024",
    "category": "exam"
  }'

# Delete grade
curl -X DELETE "$GATEWAY_URL/api/grades/grade-001"
```

## Achievement Service CRUD Operations

### Create Achievement
```bash
# Create a new achievement
curl -X POST "$GATEWAY_URL/api/achievements" \
  -H "Content-Type: application/json" \
  -d '{
    "studentId": "student-001",
    "title": "Dean'\''s List",
    "description": "Achieved GPA above 3.75 for Fall 2024 semester",
    "category": "academic",
    "dateAwarded": "2024-12-01",
    "points": 150,
    "level": "semester",
    "metadata": {
      "gpa": 3.85,
      "semester": "Fall 2024",
      "coursesCompleted": 6
    }
  }'

# Create competition achievement
curl -X POST "$GATEWAY_URL/api/achievements" \
  -H "Content-Type: application/json" \
  -d '{
    "studentId": "student-002",
    "title": "Robotics Competition Winner",
    "description": "First place in regional robotics competition",
    "category": "competition",
    "dateAwarded": "2024-11-15",
    "points": 250,
    "level": "regional",
    "metadata": {
      "event": "Regional Robotics Championship 2024",
      "placement": 1,
      "teamSize": 4,
      "competitorCount": 45
    }
  }'
```

### Read Achievements
```bash
# Get all achievements
curl -X GET "$GATEWAY_URL/api/achievements"

# Get achievement by ID
curl -X GET "$GATEWAY_URL/api/achievements/achievement-001"

# Get achievements by student
curl -X GET "$GATEWAY_URL/api/achievements?studentId=student-001"

# Get achievements by category
curl -X GET "$GATEWAY_URL/api/achievements?category=academic"

# Get achievements by level
curl -X GET "$GATEWAY_URL/api/achievements?level=semester"

# Get achievements by date range
curl -X GET "$GATEWAY_URL/api/achievements?fromDate=2024-10-01&toDate=2024-11-30"

# Get top achievements (highest points)
curl -X GET "$GATEWAY_URL/api/achievements?sortBy=points&order=desc&limit=10"
```

### Update Achievement
```bash
# Update achievement information
curl -X PUT "$GATEWAY_URL/api/achievements/achievement-001" \
  -H "Content-Type: application/json" \
  -d '{
    "studentId": "student-002",
    "title": "Honor Roll - Fall 2024",
    "description": "Achieved GPA above 3.5 for Fall 2024 semester with exceptional performance",
    "category": "academic",
    "dateAwarded": "2024-11-01",
    "points": 125,
    "level": "semester",
    "metadata": {
      "gpa": 3.82,
      "semester": "Fall 2024",
      "coursesCompleted": 5,
      "perfectAttendance": true
    }
  }'

# Update achievement points only
curl -X PATCH "$GATEWAY_URL/api/achievements/achievement-001" \
  -H "Content-Type: application/json" \
  -d '{
    "points": 150
  }'
```

### Delete Achievement
```bash
# Delete achievement by ID
curl -X DELETE "$GATEWAY_URL/api/achievements/achievement-001"
```

## Advanced Queries and Analytics

### Student Analytics
```bash
# Get student's complete academic profile
curl -X GET "$GATEWAY_URL/api/students/student-001/profile"

# Get student's GPA and grades summary
curl -X GET "$GATEWAY_URL/api/students/student-001/grades/summary"

# Get student's achievements summary
curl -X GET "$GATEWAY_URL/api/students/student-001/achievements/summary"

# Get student's enrolled courses
curl -X GET "$GATEWAY_URL/api/students/student-001/courses"
```

### Teacher Analytics
```bash
# Get teacher's courses and students
curl -X GET "$GATEWAY_URL/api/teachers/teacher-001/courses"

# Get teacher's grading statistics
curl -X GET "$GATEWAY_URL/api/teachers/teacher-001/grades/statistics"

# Get students taught by teacher
curl -X GET "$GATEWAY_URL/api/teachers/teacher-001/students"
```

### Course Analytics
```bash
# Get course enrollment statistics
curl -X GET "$GATEWAY_URL/api/courses/course-001/enrollment/stats"

# Get course grade distribution
curl -X GET "$GATEWAY_URL/api/courses/course-001/grades/distribution"

# Get course performance metrics
curl -X GET "$GATEWAY_URL/api/courses/course-001/performance"
```

### System-wide Analytics
```bash
# Get school-wide statistics
curl -X GET "$GATEWAY_URL/api/analytics/overview"

# Get enrollment trends
curl -X GET "$GATEWAY_URL/api/analytics/enrollment/trends"

# Get achievement statistics
curl -X GET "$GATEWAY_URL/api/analytics/achievements/stats"

# Get grade distribution across all courses
curl -X GET "$GATEWAY_URL/api/analytics/grades/distribution"
```

## Batch Operations

### Bulk Student Creation
```bash
# Create multiple students at once
curl -X POST "$GATEWAY_URL/api/students/batch" \
  -H "Content-Type: application/json" \
  -d '{
    "students": [
      {
        "firstName": "Alice",
        "lastName": "Johnson",
        "email": "alice.johnson@school.edu",
        "grade": "9",
        "status": "active"
      },
      {
        "firstName": "Bob",
        "lastName": "Williams",
        "email": "bob.williams@school.edu",
        "grade": "10",
        "status": "active"
      }
    ]
  }'
```

### Bulk Grade Entry
```bash
# Enter grades for multiple students
curl -X POST "$GATEWAY_URL/api/grades/batch" \
  -H "Content-Type: application/json" \
  -d '{
    "grades": [
      {
        "studentId": "student-001",
        "courseId": "course-001",
        "teacherId": "teacher-001",
        "assignmentName": "Quiz 1",
        "grade": "A",
        "points": 18,
        "maxPoints": 20,
        "category": "quiz"
      },
      {
        "studentId": "student-002",
        "courseId": "course-001",
        "teacherId": "teacher-001",
        "assignmentName": "Quiz 1",
        "grade": "B+",
        "points": 17,
        "maxPoints": 20,
        "category": "quiz"
      }
    ]
  }'
```

## Error Handling Examples

### Invalid Data
```bash
# Try to create student with invalid email
curl -X POST "$GATEWAY_URL/api/students" \
  -H "Content-Type: application/json" \
  -d '{
    "firstName": "Test",
    "lastName": "Student",
    "email": "invalid-email",
    "grade": "10"
  }'

# Try to access non-existent resource
curl -X GET "$GATEWAY_URL/api/students/non-existent-id"

# Try to update with invalid data
curl -X PUT "$GATEWAY_URL/api/students/student-001" \
  -H "Content-Type: application/json" \
  -d '{
    "grade": "invalid-grade"
  }'
```

## Testing Scripts

### Quick Health Check
```bash
#!/bin/bash
# Save as test-health.sh

services=("$GATEWAY_URL" "$STUDENT_SERVICE_URL" "$TEACHER_SERVICE_URL" "$ACADEMIC_SERVICE_URL" "$ACHIEVEMENT_SERVICE_URL")

for service in "${services[@]}"; do
    echo "Testing $service/health"
    curl -s -o /dev/null -w "%{http_code}" "$service/health" && echo " - OK" || echo " - FAILED"
done
```

### Load Test Sample
```bash
#!/bin/bash
# Save as load-test.sh

# Create 10 students rapidly
for i in {1..10}; do
    curl -X POST "$GATEWAY_URL/api/students" \
      -H "Content-Type: application/json" \
      -d "{
        \"firstName\": \"Student$i\",
        \"lastName\": \"Test$i\",
        \"email\": \"student$i.test$i@school.edu\",
        \"grade\": \"10\",
        \"status\": \"active\"
      }" &
done

wait
echo "Created 10 test students"
```

## PowerShell Examples (Windows)

### Windows PowerShell Equivalents
```powershell
# Set variables
$GATEWAY_URL = "http://localhost:8080"

# Get all students
Invoke-RestMethod -Uri "$GATEWAY_URL/api/students" -Method GET

# Create new student
$studentData = @{
    firstName = "Jane"
    lastName = "Doe"
    email = "jane.doe@school.edu"
    grade = "10"
    status = "active"
} | ConvertTo-Json

Invoke-RestMethod -Uri "$GATEWAY_URL/api/students" -Method POST -Body $studentData -ContentType "application/json"

# Health check
Invoke-RestMethod -Uri "$GATEWAY_URL/health" -Method GET
```

## Environment Variables for Testing

```bash
# Set these in your environment for easier testing
export GATEWAY_URL="http://localhost:8080"
export STUDENT_SERVICE_URL="http://localhost:8081"
export TEACHER_SERVICE_URL="http://localhost:8082"
export ACADEMIC_SERVICE_URL="http://localhost:8083"
export ACHIEVEMENT_SERVICE_URL="http://localhost:8084"

# Now you can use variables in curl commands
curl -X GET "$GATEWAY_URL/api/students"
```

## Testing Checklist

- [ ] All services are running and healthy
- [ ] Can create, read, update, and delete students
- [ ] Can create, read, update, and delete teachers
- [ ] Can create, read, update, and delete courses
- [ ] Can create, read, update, and delete grades
- [ ] Can create, read, update, and delete achievements
- [ ] Search and filtering work correctly
- [ ] Pagination works for large datasets
- [ ] Error handling returns appropriate status codes
- [ ] Batch operations work correctly
- [ ] Analytics endpoints return valid data

Use these cURL commands to thoroughly test your School Management System API!
