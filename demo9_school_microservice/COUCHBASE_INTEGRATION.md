# Couchbase Database Integration Guide

This guide provides details for accessing and integrating Couchbase database with the School Management System.

## 🗄️ Couchbase Database Access

### Database URLs and Endpoints

#### Administrative Interface
- **Couchbase Admin Console**: `http://localhost:8091`
- **Default Credentials**: 
  - Username: `Administrator`
  - Password: `password123` (configurable in docker-compose.yml)

#### API Endpoints
- **REST API Base**: `http://localhost:8091`
- **Query Service**: `http://localhost:8093`
- **Search Service**: `http://localhost:8094`
- **Analytics Service**: `http://localhost:8095`
- **Eventing Service**: `http://localhost:8096`

#### SDK Connection
- **Connection String**: `couchbase://localhost`
- **Bucket Name**: `schoolmgmt`
- **Username**: `Administrator`
- **Password**: `password123`

### Port Configuration
```yaml
# From docker-compose.yml
ports:
  - "8091-8096:8091-8096"  # Web Console & REST API
  - "11210:11210"          # Data Service
```

## 🔧 Initial Setup Commands

### 1. Initialize Cluster (First Time Setup)
```bash
# Initialize the cluster
curl -v -X POST http://localhost:8091/pools/default \
  -d 'memoryQuota=512' \
  -d 'indexMemoryQuota=256'

# Setup Administrator
curl -v -X POST http://localhost:8091/settings/web \
  -d 'username=Administrator' \
  -d 'password=password123' \
  -d 'port=SAME'

# Create bucket for school management
curl -v -X POST http://localhost:8091/pools/default/buckets \
  -u Administrator:password123 \
  -d 'name=schoolmgmt' \
  -d 'ramQuotaMB=512' \
  -d 'bucketType=membase'
```

### 2. Create Collections and Scopes
```bash
# Create scope for school management
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

## 📚 CRUD Operations with cURL

### Students Collection

#### Create Student (POST)
```bash
# Create a new student document
curl -X POST http://localhost:8093/query/service \
  -u Administrator:password123 \
  -H "Content-Type: application/json" \
  -d '{
    "statement": "INSERT INTO `schoolmgmt`.`school`.`students` (KEY, VALUE) VALUES (\"student-001\", { \"id\": \"student-001\", \"first_name\": \"John\", \"last_name\": \"Doe\", \"email\": \"john.doe@school.edu\", \"grade\": \"10\", \"age\": 16, \"date_of_birth\": \"2008-05-15\", \"address\": { \"street\": \"123 Main St\", \"city\": \"Springfield\", \"state\": \"IL\", \"zip_code\": \"62701\" }, \"parent_contact\": { \"name\": \"Jane Doe\", \"phone\": \"+1-555-0123\", \"email\": \"jane.doe@email.com\" }, \"enrollment_date\": \"2024-08-01\", \"status\": \"active\", \"created_at\": \"2024-08-05T10:30:00Z\", \"updated_at\": \"2024-08-05T10:30:00Z\" })"
  }'
  

# Alternative: Direct document insertion via Key-Value API
curl -X PUT http://localhost:8091/pools/default/buckets/schoolmgmt/docs/student-002 \
  -u Administrator:password123 \
  -H "Content-Type: application/json" \
  -d '{
    "id": "student-002",
    "first_name": "Alice",
    "last_name": "Johnson",
    "email": "alice.johnson@school.edu",
    "grade": "11",
    "age": 17,
    "date_of_birth": "2007-03-22",
    "address": {
      "street": "456 Oak Ave",
      "city": "Springfield",
      "state": "IL",
      "zip_code": "62702"
    },
    "parent_contact": {
      "name": "Bob Johnson",
      "phone": "+1-555-0456",
      "email": "bob.johnson@email.com"
    },
    "enrollment_date": "2024-08-01",
    "status": "active",
    "created_at": "2024-08-05T10:35:00Z",
    "updated_at": "2024-08-05T10:35:00Z"
  }'
```

#### Read Student (GET)
```bash
# Get student by document key
curl -X POST http://localhost:8093/query/service \
  -u Administrator:password123 \
  -H "Content-Type: application/json" \
  -d '{
    "statement": "SELECT * FROM `schoolmgmt`.`school`.`students` USE KEYS [\"student-001\"]"
  }'
```
# Get student by document key
curl -X GET http://localhost:8091/pools/default/buckets/schoolmgmt/docs/student-001 \
  -u Administrator:password123

# Query students using N1QL
curl -X POST http://localhost:8093/query/service \
  -u Administrator:password123 \
  -H "Content-Type: application/json" \
  -d '{
    "statement": "SELECT * FROM schoolmgmt.school.students WHERE grade = \"10\""
  }'

# Get all students
curl -X POST http://localhost:8093/query/service \
  -u Administrator:password123 \
  -H "Content-Type: application/json" \
  -d '{
    "statement": "SELECT META().id, * FROM schoolmgmt.school.students ORDER BY last_name, first_name"
  }'

# Search students by name
curl -X POST http://localhost:8093/query/service \
  -u Administrator:password123 \
  -H "Content-Type: application/json" \
  -d '{
    "statement": "SELECT * FROM schoolmgmt.school.students WHERE LOWER(first_name) LIKE \"%john%\" OR LOWER(last_name) LIKE \"%john%\""
  }'
```

#### Update Student (PUT)
```bash
# Update specific fields using N1QL
curl -X POST http://localhost:8093/query/service \
  -u Administrator:password123 \
  -H "Content-Type: application/json" \
  -d '{
    "statement": "UPDATE schoolmgmt.school.students SET grade = \"11\", age = 17, updated_at = \"2024-08-05T11:00:00Z\" WHERE META().id = \"student-001\""
  }'

# Update entire document via Key-Value API
curl -X PUT http://localhost:8091/pools/default/buckets/schoolmgmt/docs/student-001 \
  -u Administrator:password123 \
  -H "Content-Type: application/json" \
  -d '{
    "id": "student-001",
    "first_name": "John",
    "last_name": "Doe",
    "email": "john.doe@school.edu",
    "grade": "11",
    "age": 17,
    "date_of_birth": "2008-05-15",
    "address": {
      "street": "123 Main St",
      "city": "Springfield",
      "state": "IL",
      "zip_code": "62701"
    },
    "parent_contact": {
      "name": "Jane Doe",
      "phone": "+1-555-0123",
      "email": "jane.doe@email.com"
    },
    "enrollment_date": "2024-08-01",
    "status": "active",
    "created_at": "2024-08-05T10:30:00Z",
    "updated_at": "2024-08-05T11:00:00Z"
  }'
```

#### Delete Student (DELETE)
```bash
# Delete student using N1QL
curl -X POST http://localhost:8093/query/service \
  -u Administrator:password123 \
  -H "Content-Type: application/json" \
  -d '{
    "statement": "DELETE FROM schoolmgmt.school.students WHERE META().id = \"student-001\""
  }'

# Delete via Key-Value API
curl -X DELETE http://localhost:8091/pools/default/buckets/schoolmgmt/docs/student-001 \
  -u Administrator:password123
```

### Teachers Collection

#### Create Teacher
```bash
curl -X POST http://localhost:8093/query/service \
  -u Administrator:password123 \
  -H "Content-Type: application/json" \
  -d '{
    "statement": "INSERT INTO schoolmgmt.school.teachers (KEY, VALUE) VALUES (\"teacher-001\", {
      \"id\": \"teacher-001\",
      \"first_name\": \"Dr. Emma\",
      \"last_name\": \"Wilson\",
      \"email\": \"emma.wilson@school.edu\",
      \"phone\": \"+1-555-0789\",
      \"department\": \"Mathematics\",
      \"subjects\": [\"Algebra\", \"Geometry\", \"Calculus\"],
      \"qualification\": \"PhD in Mathematics\",
      \"experience\": 8,
      \"hire_date\": \"2020-01-15\",
      \"employee_id\": \"EMP-2020-001\",
      \"salary\": 75000,
      \"address\": {
        \"street\": \"789 Elm St\",
        \"city\": \"Springfield\",
        \"state\": \"IL\",
        \"zip_code\": \"62703\"
      },
      \"emergency_contact\": {
        \"name\": \"James Wilson\",
        \"relationship\": \"Spouse\",
        \"phone\": \"+1-555-0790\"
      },
      \"status\": \"active\",
      \"created_at\": \"2024-08-05T10:45:00Z\",
      \"updated_at\": \"2024-08-05T10:45:00Z\"
    })"
  }'
```

#### Read Teachers
```bash
# Get all teachers in Mathematics department
curl -X POST http://localhost:8093/query/service \
  -u Administrator:password123 \
  -H "Content-Type: application/json" \
  -d '{
    "statement": "SELECT * FROM schoolmgmt.school.teachers WHERE department = \"Mathematics\""
  }'

# Get teacher by ID
curl -X GET http://localhost:8091/pools/default/buckets/schoolmgmt/docs/teacher-001 \
  -u Administrator:password123
```

#### Update Teacher
```bash
# Update teacher salary and experience
curl -X POST http://localhost:8093/query/service \
  -u Administrator:password123 \
  -H "Content-Type: application/json" \
  -d '{
    "statement": "UPDATE schoolmgmt.school.teachers SET salary = 80000, experience = 9, updated_at = \"2024-08-05T12:00:00Z\" WHERE META().id = \"teacher-001\""
  }'
```

#### Delete Teacher
```bash
# Soft delete (change status)
curl -X POST http://localhost:8093/query/service \
  -u Administrator:password123 \
  -H "Content-Type: application/json" \
  -d '{
    "statement": "UPDATE schoolmgmt.school.teachers SET status = \"inactive\", updated_at = \"2024-08-05T12:30:00Z\" WHERE META().id = \"teacher-001\""
  }'

# Hard delete
curl -X DELETE http://localhost:8091/pools/default/buckets/schoolmgmt/docs/teacher-001 \
  -u Administrator:password123
```

### Academic Records Collection

#### Create Academic Record
```bash
curl -X POST http://localhost:8093/query/service \
  -u Administrator:password123 \
  -H "Content-Type: application/json" \
  -d '{
    "statement": "INSERT INTO schoolmgmt.school.academics (KEY, VALUE) VALUES (\"academic-001\", {
      \"id\": \"academic-001\",
      \"student_id\": \"student-001\",
      \"teacher_id\": \"teacher-001\",
      \"subject\": \"Algebra\",
      \"grade\": \"A\",
      \"semester\": \"Spring 2024\",
      \"academic_year\": 2024,
      \"exam_type\": \"Final\",
      \"max_marks\": 100,
      \"obtained_marks\": 95,
      \"percentage\": 95.0,
      \"exam_date\": \"2024-04-15\",
      \"status\": \"pass\",
      \"comments\": \"Excellent performance\",
      \"created_at\": \"2024-08-05T11:00:00Z\",
      \"updated_at\": \"2024-08-05T11:00:00Z\"
    })"
  }'
```

#### Query Academic Records
```bash
# Get all academic records for a student
curl -X POST http://localhost:8093/query/service \
  -u Administrator:password123 \
  -H "Content-Type: application/json" \
  -d '{
    "statement": "SELECT * FROM schoolmgmt.school.academics WHERE student_id = \"student-001\" ORDER BY academic_year DESC, semester"
  }'

# Get academic records by teacher
curl -X POST http://localhost:8093/query/service \
  -u Administrator:password123 \
  -H "Content-Type: application/json" \
  -d '{
    "statement": "SELECT * FROM schoolmgmt.school.academics WHERE teacher_id = \"teacher-001\" ORDER BY exam_date DESC"
  }'

# Get students with high performance (>90%)
curl -X POST http://localhost:8093/query/service \
  -u Administrator:password123 \
  -H "Content-Type: application/json" \
  -d '{
    "statement": "SELECT student_id, subject, percentage FROM schoolmgmt.school.academics WHERE percentage > 90 ORDER BY percentage DESC"
  }'
```

### Achievements Collection

#### Create Achievement
```bash
curl -X POST http://localhost:8093/query/service \
  -u Administrator:password123 \
  -H "Content-Type: application/json" \
  -d '{
    "statement": "INSERT INTO schoolmgmt.school.achievements (KEY, VALUE) VALUES (\"achievement-001\", {
      \"id\": \"achievement-001\",
      \"student_id\": \"student-001\",
      \"title\": \"Math Excellence Award\",
      \"description\": \"Outstanding performance in mathematics\",
      \"category\": \"academic\",
      \"points\": 100,
      \"date\": \"2024-04-20\",
      \"awarded_by\": \"teacher-001\",
      \"certificate_number\": \"CERT-2024-001\",
      \"status\": \"active\",
      \"created_at\": \"2024-08-05T11:30:00Z\",
      \"updated_at\": \"2024-08-05T11:30:00Z\"
    })"
  }'
```

#### Query Achievements
```bash
# Get student achievements
curl -X POST http://localhost:8093/query/service \
  -u Administrator:password123 \
  -H "Content-Type: application/json" \
  -d '{
    "statement": "SELECT * FROM schoolmgmt.school.achievements WHERE student_id = \"student-001\" ORDER BY date DESC"
  }'

# Get leaderboard (top students by points)
curl -X POST http://localhost:8093/query/service \
  -u Administrator:password123 \
  -H "Content-Type: application/json" \
  -d '{
    "statement": "SELECT student_id, SUM(points) as total_points FROM schoolmgmt.school.achievements GROUP BY student_id ORDER BY total_points DESC LIMIT 10"
  }'
```

## 🔍 Advanced Queries

### Complex Reporting Queries
```bash
# Student performance report
curl -X POST http://localhost:8093/query/service \
  -u Administrator:password123 \
  -H "Content-Type: application/json" \
  -d '{
    "statement": "SELECT s.first_name, s.last_name, s.grade, AVG(a.percentage) as avg_score FROM schoolmgmt.school.students s JOIN schoolmgmt.school.academics a ON META(s).id = a.student_id GROUP BY s.first_name, s.last_name, s.grade ORDER BY avg_score DESC"
  }'

# Teacher workload analysis
curl -X POST http://localhost:8093/query/service \
  -u Administrator:password123 \
  -H "Content-Type: application/json" \
  -d '{
    "statement": "SELECT t.first_name, t.last_name, t.department, COUNT(a.id) as student_count FROM schoolmgmt.school.teachers t LEFT JOIN schoolmgmt.school.academics a ON META(t).id = a.teacher_id GROUP BY t.first_name, t.last_name, t.department ORDER BY student_count DESC"
  }'

# Grade distribution
curl -X POST http://localhost:8093/query/service \
  -u Administrator:password123 \
  -H "Content-Type: application/json" \
  -d '{
    "statement": "SELECT grade, COUNT(*) as count FROM schoolmgmt.school.academics GROUP BY grade ORDER BY grade"
  }'
```

## 🛠️ Database Administration

### Backup Commands
```bash
# Create backup
curl -X POST http://localhost:8091/controller/startBackup \
  -u Administrator:password123 \
  -d 'bucket=schoolmgmt' \
  -d 'backup_dir=/backup/schoolmgmt'

# List backups
curl -X GET http://localhost:8091/controller/listBackups \
  -u Administrator:password123
```

### Index Management
```bash
# Create primary index
curl -X POST http://localhost:8093/query/service \
  -u Administrator:password123 \
  -H "Content-Type: application/json" \
  -d '{
    "statement": "CREATE PRIMARY INDEX ON schoolmgmt.school.students"
  }'

# Create secondary indexes for better performance
curl -X POST http://localhost:8093/query/service \
  -u Administrator:password123 \
  -H "Content-Type: application/json" \
  -d '{
    "statement": "CREATE INDEX idx_student_grade ON schoolmgmt.school.students(grade)"
  }'

curl -X POST http://localhost:8093/query/service \
  -u Administrator:password123 \
  -H "Content-Type: application/json" \
  -d '{
    "statement": "CREATE INDEX idx_teacher_department ON schoolmgmt.school.teachers(department)"
  }'

curl -X POST http://localhost:8093/query/service \
  -u Administrator:password123 \
  -H "Content-Type: application/json" \
  -d '{
    "statement": "CREATE INDEX idx_academic_student ON schoolmgmt.school.academics(student_id, academic_year)"
  }'
```

## 🔧 Integration with Go Services

### Sample Go SDK Connection Code
```go
package main

import (
    "github.com/couchbase/gocb/v2"
    "log"
    "time"
)

func connectToCouchbase() (*gocb.Cluster, error) {
    // Connect to cluster
    cluster, err := gocb.Connect(
        "couchbase://localhost",
        gocb.ClusterOptions{
            Username: "Administrator",
            Password: "password123",
        },
    )
    if err != nil {
        return nil, err
    }

    // Wait for cluster to be ready
    err = cluster.WaitUntilReady(5*time.Second, nil)
    if err != nil {
        return nil, err
    }

    return cluster, nil
}

func main() {
    cluster, err := connectToCouchbase()
    if err != nil {
        log.Fatal(err)
    }
    defer cluster.Close(nil)

    // Get bucket and collection
    bucket := cluster.Bucket("schoolmgmt")
    collection := bucket.Scope("school").Collection("students")

    // Example usage
    student := map[string]interface{}{
        "first_name": "John",
        "last_name":  "Doe",
        "email":      "john.doe@school.edu",
        "grade":      "10",
    }

    _, err = collection.Insert("student-123", student, nil)
    if err != nil {
        log.Fatal(err)
    }

    log.Println("Student created successfully")
}
```

## 📊 Monitoring and Health Checks

### Database Health Checks
```bash
# Check cluster health
curl -X GET http://localhost:8091/pools/default \
  -u Administrator:password123

# Check bucket health
curl -X GET http://localhost:8091/pools/default/buckets/schoolmgmt \
  -u Administrator:password123

# Check node status
curl -X GET http://localhost:8091/pools/nodes \
  -u Administrator:password123

# Query service health
curl -X GET http://localhost:8093/admin/ping \
  -u Administrator:password123
```

### Performance Monitoring
```bash
# Get bucket statistics
curl -X GET http://localhost:8091/pools/default/buckets/schoolmgmt/stats \
  -u Administrator:password123

# Get query statistics
curl -X GET http://localhost:8093/admin/stats \
  -u Administrator:password123
```

## 🚀 Production Considerations

1. **Security**: Change default passwords and enable SSL/TLS
2. **Backup**: Set up automated backup schedules
3. **Monitoring**: Integrate with monitoring tools like Prometheus
4. **Scaling**: Configure cluster with multiple nodes for high availability
5. **Indexing**: Create appropriate indexes for query performance
6. **Memory**: Allocate sufficient memory quotas based on data size

## 🚀 Quick Start with Automated Scripts

For easy setup and testing, use the provided automation scripts:

### Windows Users
```cmd
# Simple batch launcher
cd scripts
run-demo.bat

# Or run PowerShell directly
.\scripts\couchbase-demo.ps1
```

### Linux/macOS Users
```bash
# Make executable and run
chmod +x scripts/couchbase-demo.sh
./scripts/couchbase-demo.sh
```

### Script Features
- 🔄 **Interactive menus** for easy navigation
- ✅ **Automated setup** of cluster, collections, and indexes
- 🎯 **Complete CRUD demos** for all collections
- 📊 **Real-time statistics** and progress feedback
- 🛠️ **Error handling** and connection testing
- 📋 **Multiple execution modes** (setup, demo, test, interactive)

See [`scripts/README.md`](./scripts/README.md) for detailed documentation.

---

## 📝 Notes

- All examples use basic authentication for simplicity
- In production, use SSL/TLS and proper authentication mechanisms
- Adjust memory quotas based on your data requirements
- Create indexes for frequently queried fields
- Use prepared statements for better performance in applications
- Monitor query performance and optimize as needed

---

This guide provides comprehensive examples for integrating Couchbase with the School Management System. All URLs and credentials are configured for the current Docker Compose setup.
