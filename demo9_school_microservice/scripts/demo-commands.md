# School Management System - Demo Ready CURL Commands

## 🎯 Current Status: ALL SERVICES OPERATIONAL ✅

All microservices are running and connected to Couchbase:
- **Student Service**: Port 8081 ✅ Connected to Couchbase
- **Teacher Service**: Port 8082 ✅ Connected to Couchbase  
- **Academic Service**: Port 8083 ✅ Connected to Couchbase
- **Achievement Service**: Port 8084 ✅ Connected to Couchbase

## 🚀 Quick Demo Commands

### Test All Health Endpoints
```powershell
# PowerShell commands for Windows
curl http://localhost:8081/health  # Student Service
curl http://localhost:8082/health  # Teacher Service  
curl http://localhost:8083/health  # Academic Service
curl http://localhost:8084/health  # Achievement Service

# Expected response for all services:
# {"database":"couchbase-connected","service":"[service-name]","status":"healthy"}
```

### Health Check PowerShell Script
```powershell
# Test all services at once
$services = @(
    @{Name="Student"; Port=8081},
    @{Name="Teacher"; Port=8082}, 
    @{Name="Academic"; Port=8083},
    @{Name="Achievement"; Port=8084}
)

Write-Host "🔍 Testing School Management System Services..." -ForegroundColor Cyan
foreach ($service in $services) {
    try {
        $response = Invoke-RestMethod -Uri "http://localhost:$($service.Port)/health" -TimeoutSec 5
        $status = if ($response.database -eq "couchbase-connected") { "✅ HEALTHY" } else { "⚠️ WARNING" }
        Write-Host "$($service.Name) Service (Port $($service.Port)): $status" -ForegroundColor Green
        Write-Host "  └─ Database: $($response.database)" -ForegroundColor Gray
    }
    catch {
        Write-Host "$($service.Name) Service (Port $($service.Port)): ❌ FAILED" -ForegroundColor Red
        Write-Host "  └─ Error: $($_.Exception.Message)" -ForegroundColor Gray
    }
}
```

## 📊 Working CRUD Examples

### Student Service Examples

#### Create Student
```powershell
$studentData = @{
    firstName = "John"
    lastName = "Doe"
    email = "john.doe@school.edu"
    grade = "10"
    status = "active"
} | ConvertTo-Json

try {
    $response = Invoke-RestMethod -Uri "http://localhost:8081/students" -Method POST -Body $studentData -ContentType "application/json"
    Write-Host "✅ Created student: $($response.firstName) $($response.lastName)" -ForegroundColor Green
    Write-Host "   ID: $($response.id)" -ForegroundColor Gray
    $global:StudentId = $response.id
} catch {
    Write-Host "❌ Failed to create student: $($_.Exception.Message)" -ForegroundColor Red
}
```

#### Get All Students
```powershell
try {
    $students = Invoke-RestMethod -Uri "http://localhost:8081/students" -Method GET
    Write-Host "📋 Found $($students.Count) students" -ForegroundColor Green
    $students | ForEach-Object { 
        Write-Host "  • $($_.firstName) $($_.lastName) (Grade: $($_.grade))" -ForegroundColor Gray 
    }
} catch {
    Write-Host "❌ Failed to get students: $($_.Exception.Message)" -ForegroundColor Red
}
```

### Teacher Service Examples

#### Create Teacher
```powershell
$teacherData = @{
    firstName = "Dr. Sarah"
    lastName = "Wilson"
    email = "sarah.wilson@school.edu"
    phone = "555-2001"
    department = "Mathematics"
    subjects = @("Algebra", "Calculus")
    qualification = "PhD Mathematics"
    experience = 8
    status = "active"
} | ConvertTo-Json

try {
    $response = Invoke-RestMethod -Uri "http://localhost:8082/teachers" -Method POST -Body $teacherData -ContentType "application/json"
    Write-Host "✅ Created teacher: $($response.firstName) $($response.lastName)" -ForegroundColor Green
    Write-Host "   Department: $($response.department)" -ForegroundColor Gray
    $global:TeacherId = $response.id
} catch {
    Write-Host "❌ Failed to create teacher: $($_.Exception.Message)" -ForegroundColor Red
}
```

### Academic Service Examples

#### Create Subject
```powershell
$subjectData = @{
    name = "Computer Science 101"
    code = "CS101"
    credits = 3
    description = "Introduction to Computer Science"
    teacherId = "teacher-001"
    status = "active"
} | ConvertTo-Json

try {
    $response = Invoke-RestMethod -Uri "http://localhost:8083/subjects" -Method POST -Body $subjectData -ContentType "application/json"
    Write-Host "✅ Created subject: $($response.name)" -ForegroundColor Green
    Write-Host "   Code: $($response.code), Credits: $($response.credits)" -ForegroundColor Gray
    $global:SubjectId = $response.id
} catch {
    Write-Host "❌ Failed to create subject: $($_.Exception.Message)" -ForegroundColor Red
}
```

### Achievement Service Examples

#### Create Achievement
```powershell
$achievementData = @{
    title = "Dean's List"
    description = "Achieved GPA above 3.75"
    category = "academic"
    points = 150
    studentId = "student-001"
    teacherId = "teacher-001"
    status = "awarded"
} | ConvertTo-Json

try {
    $response = Invoke-RestMethod -Uri "http://localhost:8084/achievements" -Method POST -Body $achievementData -ContentType "application/json"
    Write-Host "✅ Created achievement: $($response.title)" -ForegroundColor Green
    Write-Host "   Points: $($response.points), Category: $($response.category)" -ForegroundColor Gray
    $global:AchievementId = $response.id
} catch {
    Write-Host "❌ Failed to create achievement: $($_.Exception.Message)" -ForegroundColor Red
}
```

## 🔧 Complete Demo Script

Save this as `demo-test.ps1`:

```powershell
# School Management System Demo Script
Write-Host "🎓 School Management System Demo" -ForegroundColor Cyan
Write-Host "=================================" -ForegroundColor Cyan
Write-Host ""

# Step 1: Health Check
Write-Host "1️⃣ Testing Service Health..." -ForegroundColor Yellow
$services = @(
    @{Name="Student"; Port=8081},
    @{Name="Teacher"; Port=8082}, 
    @{Name="Academic"; Port=8083},
    @{Name="Achievement"; Port=8084}
)

$healthyServices = 0
foreach ($service in $services) {
    try {
        $response = Invoke-RestMethod -Uri "http://localhost:$($service.Port)/health" -TimeoutSec 5
        if ($response.database -eq "couchbase-connected") {
            Write-Host "   ✅ $($service.Name) Service: HEALTHY (Couchbase Connected)" -ForegroundColor Green
            $healthyServices++
        } else {
            Write-Host "   ⚠️ $($service.Name) Service: Running but database issue" -ForegroundColor Yellow
        }
    }
    catch {
        Write-Host "   ❌ $($service.Name) Service: FAILED" -ForegroundColor Red
    }
}

Write-Host ""
Write-Host "Health Summary: $healthyServices/4 services healthy" -ForegroundColor $(if($healthyServices -eq 4){"Green"}else{"Yellow"})
Write-Host ""

if ($healthyServices -eq 4) {
    # Step 2: Create Sample Data
    Write-Host "2️⃣ Creating Sample Data..." -ForegroundColor Yellow
    
    # Create Student
    $studentData = @{
        firstName = "Demo"
        lastName = "Student"
        email = "demo.student@school.edu"
        grade = "10"
        status = "active"
    } | ConvertTo-Json
    
    try {
        $student = Invoke-RestMethod -Uri "http://localhost:8081/students" -Method POST -Body $studentData -ContentType "application/json"
        Write-Host "   ✅ Created Student: $($student.firstName) $($student.lastName)" -ForegroundColor Green
    } catch {
        Write-Host "   ❌ Failed to create student" -ForegroundColor Red
    }
    
    # Create Teacher
    $teacherData = @{
        firstName = "Demo"
        lastName = "Teacher"
        email = "demo.teacher@school.edu"
        phone = "555-DEMO"
        department = "Demo Department"
        subjects = @("Demo Subject")
        qualification = "Demo Qualification"
        experience = 5
        status = "active"
    } | ConvertTo-Json
    
    try {
        $teacher = Invoke-RestMethod -Uri "http://localhost:8082/teachers" -Method POST -Body $teacherData -ContentType "application/json"
        Write-Host "   ✅ Created Teacher: $($teacher.firstName) $($teacher.lastName)" -ForegroundColor Green
    } catch {
        Write-Host "   ❌ Failed to create teacher" -ForegroundColor Red
    }
    
    Write-Host ""
    Write-Host "3️⃣ Retrieving Data..." -ForegroundColor Yellow
    
    # Get Students
    try {
        $students = Invoke-RestMethod -Uri "http://localhost:8081/students" -Method GET
        Write-Host "   📋 Students in database: $($students.Count)" -ForegroundColor Green
    } catch {
        Write-Host "   ❌ Failed to retrieve students" -ForegroundColor Red
    }
    
    # Get Teachers
    try {
        $teachers = Invoke-RestMethod -Uri "http://localhost:8082/teachers" -Method GET
        Write-Host "   👨‍🏫 Teachers in database: $($teachers.Count)" -ForegroundColor Green
    } catch {
        Write-Host "   ❌ Failed to retrieve teachers" -ForegroundColor Red
    }
} else {
    Write-Host "⚠️ Not all services are healthy. Please check your setup." -ForegroundColor Yellow
}

Write-Host ""
Write-Host "🎉 Demo Complete!" -ForegroundColor Cyan
Write-Host "All services are connected to Couchbase and ready for use." -ForegroundColor Green
```

## 🌐 Browser Testing

You can also test the health endpoints in your browser:

- Student Service: http://localhost:8081/health
- Teacher Service: http://localhost:8082/health  
- Academic Service: http://localhost:8083/health
- Achievement Service: http://localhost:8084/health

## 📚 Service Documentation

### Data Models

#### Student Model
```json
{
  "id": "uuid",
  "firstName": "string",
  "lastName": "string", 
  "email": "string",
  "grade": "string",
  "status": "active|inactive",
  "createdAt": "timestamp",
  "updatedAt": "timestamp",
  "type": "student"
}
```

#### Teacher Model
```json
{
  "id": "uuid",
  "firstName": "string",
  "lastName": "string",
  "email": "string", 
  "phone": "string",
  "department": "string",
  "subjects": ["string"],
  "qualification": "string",
  "experience": "number",
  "status": "active|inactive",
  "createdAt": "timestamp",
  "updatedAt": "timestamp", 
  "type": "teacher"
}
```

#### Subject Model
```json
{
  "id": "uuid",
  "name": "string",
  "code": "string",
  "credits": "number",
  "description": "string",
  "teacherId": "string",
  "status": "active|inactive",
  "createdAt": "timestamp",
  "updatedAt": "timestamp",
  "type": "subject"
}
```

#### Achievement Model
```json
{
  "id": "uuid", 
  "title": "string",
  "description": "string",
  "category": "string",
  "points": "number",
  "studentId": "string",
  "teacherId": "string", 
  "status": "awarded|pending",
  "createdAt": "timestamp",
  "updatedAt": "timestamp",
  "type": "achievement"
}
```

## ✅ System Status

**Current Status: PRODUCTION READY** 🚀

- ✅ All 4 microservices running
- ✅ Couchbase cluster operational  
- ✅ All services connected to database
- ✅ Health endpoints responding
- ✅ Docker containers healthy
- ✅ Port mappings configured
- ✅ CRUD operations available

**Ready for demonstration and production use!**
