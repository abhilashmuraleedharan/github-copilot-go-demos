# School Management System - Demo Guide

This guide will walk you through setting up and testing the School Management System.

## Prerequisites

- Docker and Docker Compose installed
- PowerShell or Bash terminal
- Internet connection for downloading Docker images

## Quick Start

1. **Navigate to the project directory:**
```powershell
cd d:\demo\schoolmgmt
```

2. **Start all services:**
```powershell
docker-compose up -d
```

3. **Verify services are running:**
```powershell
docker-compose ps
```

All services should show as "healthy". Wait a few moments for all health checks to pass.

## Service Endpoints

### API Gateway (Port 8080) - **Main Entry Point**
- Health: `GET http://localhost:8080/health`
- System Health: `GET http://localhost:8080/health/system`
- Dashboard Stats: `GET http://localhost:8080/dashboard/stats`
- Dashboard Summary: `GET http://localhost:8080/dashboard/summary`
- Service Discovery: `GET http://localhost:8080/services`
- All APIs are accessible through: `http://localhost:8080/api/v1/`

### Individual Services (for direct access if needed)
- Student Service: `http://localhost:8081`
- Teacher Service: `http://localhost:8082`
- Academic Service: `http://localhost:8083`
- Achievement Service: `http://localhost:8084`
- Couchbase UI: `http://localhost:8091`

## API Testing Examples (PowerShell)

### Gateway Health Check
```powershell
Invoke-WebRequest -Uri http://localhost:8080/health -Method GET
```

### Students API
```powershell
# Create a student
Invoke-WebRequest -Uri http://localhost:8080/api/v1/students -Method POST -Headers @{"Content-Type"="application/json"} -Body '{"name":"John Doe","email":"john.doe@school.edu","grade":"10","age":16}'

# Get all students
Invoke-WebRequest -Uri http://localhost:8080/api/v1/students -Method GET

# Get student by ID
Invoke-WebRequest -Uri http://localhost:8080/api/v1/students/student-1 -Method GET

# Update student
Invoke-WebRequest -Uri http://localhost:8080/api/v1/students/student-1 -Method PUT -Headers @{"Content-Type"="application/json"} -Body '{"name":"John Smith","email":"john.smith@school.edu","grade":"11","age":17}'

# Delete student
Invoke-WebRequest -Uri http://localhost:8080/api/v1/students/student-1 -Method DELETE
```

### Teachers API
```powershell
# Create a teacher
Invoke-WebRequest -Uri http://localhost:8080/api/v1/teachers -Method POST -Headers @{"Content-Type"="application/json"} -Body '{"name":"Dr. Smith","email":"smith@school.edu","department":"Mathematics","hire_date":"2020-01-15"}'

# Get all teachers
Invoke-WebRequest -Uri http://localhost:8080/api/v1/teachers -Method GET

# Get teacher by ID
Invoke-WebRequest -Uri http://localhost:8080/api/v1/teachers/teacher-1 -Method GET
```

### Academics API
```powershell
# Create an academic record
Invoke-WebRequest -Uri http://localhost:8080/api/v1/academics -Method POST -Headers @{"Content-Type"="application/json"} -Body '{"student_id":"student-1","teacher_id":"teacher-1","subject":"Mathematics","grade":"A","semester":"Spring 2024"}'

# Get academic records
Invoke-WebRequest -Uri http://localhost:8080/api/v1/academics -Method GET

# Get academic records by student
Invoke-WebRequest -Uri http://localhost:8080/api/v1/academics/student/student-1 -Method GET

# Get classes
Invoke-WebRequest -Uri http://localhost:8080/api/v1/classes -Method GET
```

### Achievements API
```powershell
# Create an achievement
Invoke-WebRequest -Uri http://localhost:8080/api/v1/achievements -Method POST -Headers @{"Content-Type"="application/json"} -Body '{"student_id":"student-1","title":"Math Excellence","description":"Outstanding performance in mathematics","category":"academic","date":"2024-03-15","points":100}'

# Get achievements
Invoke-WebRequest -Uri http://localhost:8080/api/v1/achievements -Method GET

# Get leaderboard
Invoke-WebRequest -Uri http://localhost:8080/api/v1/leaderboard -Method GET

# Get achievement statistics
Invoke-WebRequest -Uri http://localhost:8080/api/v1/achievements/stats -Method GET

# Create an award
Invoke-WebRequest -Uri http://localhost:8080/api/v1/awards -Method POST -Headers @{"Content-Type"="application/json"} -Body '{"title":"Student of the Month","description":"Outstanding overall performance","category":"excellence","criteria":"Academic and extracurricular excellence"}'
```

## Complete Demo Workflow

Follow this complete workflow to test all features:

### 1. System Health Check
```powershell
# Check gateway health
Invoke-WebRequest -Uri http://localhost:8080/health -Method GET

# Check all services health
Invoke-WebRequest -Uri http://localhost:8080/health/system -Method GET
```

### 2. Create Sample Data
```powershell
# Create students
Invoke-WebRequest -Uri http://localhost:8080/api/v1/students -Method POST -Headers @{"Content-Type"="application/json"} -Body '{"name":"Alice Johnson","email":"alice@school.edu","grade":"10","age":16}'

Invoke-WebRequest -Uri http://localhost:8080/api/v1/students -Method POST -Headers @{"Content-Type"="application/json"} -Body '{"name":"Bob Smith","email":"bob@school.edu","grade":"11","age":17}'

# Create teachers
Invoke-WebRequest -Uri http://localhost:8080/api/v1/teachers -Method POST -Headers @{"Content-Type"="application/json"} -Body '{"name":"Dr. Emma Wilson","email":"emma@school.edu","department":"Mathematics","hire_date":"2020-01-15"}'

Invoke-WebRequest -Uri http://localhost:8080/api/v1/teachers -Method POST -Headers @{"Content-Type"="application/json"} -Body '{"name":"Prof. James Brown","email":"james@school.edu","department":"Science","hire_date":"2019-08-20"}'

# Create academic records
Invoke-WebRequest -Uri http://localhost:8080/api/v1/academics -Method POST -Headers @{"Content-Type"="application/json"} -Body '{"student_id":"student-1","teacher_id":"teacher-1","subject":"Algebra","grade":"A","semester":"Spring 2024"}'

Invoke-WebRequest -Uri http://localhost:8080/api/v1/academics -Method POST -Headers @{"Content-Type"="application/json"} -Body '{"student_id":"student-2","teacher_id":"teacher-2","subject":"Physics","grade":"B+","semester":"Spring 2024"}'

# Create achievements
Invoke-WebRequest -Uri http://localhost:8080/api/v1/achievements -Method POST -Headers @{"Content-Type"="application/json"} -Body '{"student_id":"student-1","title":"Math Excellence","description":"Perfect score in algebra","category":"academic","date":"2024-03-15","points":100}'

Invoke-WebRequest -Uri http://localhost:8080/api/v1/achievements -Method POST -Headers @{"Content-Type"="application/json"} -Body '{"student_id":"student-2","title":"Science Fair Winner","description":"First place in physics project","category":"competition","date":"2024-04-10","points":150}'
```

### 3. View Dashboard
```powershell
# Get comprehensive dashboard stats
Invoke-WebRequest -Uri http://localhost:8080/dashboard/stats -Method GET

# Get dashboard summary
Invoke-WebRequest -Uri http://localhost:8080/dashboard/summary -Method GET

# View leaderboard
Invoke-WebRequest -Uri http://localhost:8080/api/v1/leaderboard -Method GET
```

## Available Features

### Student Management
- ✅ Create, Read, Update, Delete students
- ✅ Search students by various criteria
- ✅ Student profile management

### Teacher Management
- ✅ Create, Read, Update, Delete teachers
- ✅ Department-wise teacher organization
- ✅ Teacher profile management

### Academic Records
- ✅ Grade management
- ✅ Subject-wise performance tracking
- ✅ Student-teacher academic relationships
- ✅ Class and semester management

### Achievement System
- ✅ Student achievements and awards
- ✅ Points-based recognition system
- ✅ Leaderboards and rankings
- ✅ Category-wise achievement tracking

### API Gateway Features
- ✅ Centralized API routing
- ✅ Service discovery
- ✅ Health monitoring
- ✅ CORS support
- ✅ Comprehensive dashboard

### Architecture Features
- ✅ Microservices architecture
- ✅ Domain-driven design
- ✅ RESTful APIs
- ✅ Docker containerization
- ✅ Service mesh with API Gateway
- ✅ Health checks and monitoring

## Stopping the System

```powershell
docker-compose down
```

To remove all data and volumes:
```powershell
docker-compose down -v
```

## Troubleshooting

1. **Services not starting:** 
   - Check if ports 8080-8084 and 8091-8096 are available
   - Run `docker-compose logs [service-name]` for detailed logs

2. **API calls failing:** 
   - Ensure all services show "healthy" status
   - Wait a few minutes after startup for all services to be ready

3. **Database connection issues:** 
   - Restart Couchbase: `docker-compose restart couchbase`
   - Check Couchbase UI at http://localhost:8091

4. **API Gateway not responding:**
   - Check gateway logs: `docker-compose logs api-gateway`
   - Restart gateway: `docker-compose restart api-gateway`

For detailed service logs:
```powershell
# View all service logs
docker-compose logs

# View specific service logs
docker-compose logs api-gateway
docker-compose logs student-service
docker-compose logs teacher-service
docker-compose logs academic-service
docker-compose logs achievement-service
docker-compose logs couchbase
```

## Next Steps

1. **Persistent Storage**: The system currently uses in-memory storage. For production, integrate with Couchbase for persistent data storage.

2. **Authentication**: Add JWT-based authentication and authorization.

3. **Advanced Features**: 
   - File uploads for student/teacher photos
   - Email notifications
   - Report generation
   - Advanced analytics

4. **Monitoring**: Add metrics and logging aggregation.

5. **Frontend**: Develop a web frontend using React, Vue, or Angular.
