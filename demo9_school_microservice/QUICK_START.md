# Quick Start Guide - School Management Microservice

## Service Status
✅ **RUNNING** - All services are operational

## Access URLs
- **API**: http://localhost:8080
- **Couchbase Console**: http://localhost:8091
  - Username: `Administrator`
  - Password: `password`

## Quick Commands

### Check Service Status
```powershell
docker ps --filter "name=school"
```

### View Logs
```powershell
# All services
docker-compose logs -f

# Specific service
docker logs school-service -f
docker logs school-couchbase
```

### Test Health
```powershell
curl http://localhost:8080/health
# OR
Invoke-WebRequest http://localhost:8080/health -UseBasicParsing
```

### Stop Services
```powershell
docker-compose stop
```

### Start Services
```powershell
docker-compose start
```

### Restart Everything
```powershell
docker-compose restart
```

### Stop and Remove (keeps data)
```powershell
docker-compose down
```

### Stop and Remove Everything (deletes data)
```powershell
docker-compose down -v
```

## Example API Calls

### Create a Teacher
```powershell
$teacher = @{
    id = "teacher002"
    firstName = "Jane"
    lastName = "Doe"
    email = "jane.doe@school.edu"
    subject = "Science"
    hireDate = "2024-01-20T00:00:00Z"
} | ConvertTo-Json

Invoke-WebRequest -Uri http://localhost:8080/api/teachers `
    -Method POST `
    -Body $teacher `
    -ContentType "application/json" `
    -UseBasicParsing
```

### Get a Teacher
```powershell
Invoke-WebRequest -Uri http://localhost:8080/api/teachers/teacher001 `
    -UseBasicParsing | Select-Object -ExpandProperty Content
```

### Create a Student
```powershell
$student = @{
    id = "student002"
    firstName = "Bob"
    lastName = "Wilson"
    email = "bob.w@school.edu"
    dateOfBirth = "2009-03-15T00:00:00Z"
    enrollmentDate = "2024-09-01T00:00:00Z"
    grade = "9"
} | ConvertTo-Json

Invoke-WebRequest -Uri http://localhost:8080/api/students `
    -Method POST `
    -Body $student `
    -ContentType "application/json" `
    -UseBasicParsing
```

## Files Reference

- **README.md** - Complete documentation and setup instructions
- **CHANGELOG.md** - Detailed change history and technical specifications
- **DEMO_SUMMARY.md** - Demo results and verification
- **docker-compose.yml** - Service orchestration configuration
- **Dockerfile** - Container build instructions

## Architecture

```
Client → HTTP (8080) → School Service (Go) → Couchbase (11210)
                                           ↓
                            Couchbase Web UI (8091)
```

## Troubleshooting

### Service won't start
```powershell
# Check if ports are in use
netstat -ano | findstr :8080
netstat -ano | findstr :8091

# View error logs
docker logs school-service
```

### Can't connect to API
```powershell
# Verify service is healthy
docker ps --filter "name=school-service"

# Wait for startup (can take 30-60 seconds)
Start-Sleep -Seconds 30
curl http://localhost:8080/health
```

### Database errors
```powershell
# Check Couchbase is ready
docker logs school-couchbase-init

# Verify bucket exists (login to http://localhost:8091)
# Check for "school" bucket in Buckets section
```

## Environment Variables

Customize by creating a `.env` file:
```env
SERVER_PORT=8080
COUCHBASE_CONNECTION_STRING=couchbase://couchbase
COUCHBASE_USERNAME=Administrator
COUCHBASE_PASSWORD=password
COUCHBASE_BUCKET=school
```

## Current Test Data

The following test data has been created:
- **Teacher**: teacher001 (John Smith, Mathematics)
- **Student**: student001 (Alice Johnson, Grade 10)
- **Class**: class001 (Algebra I, MWF 10:00-11:00)
- **Exam**: exam001 (Midterm Exam, 100 points)
- **Exam Result**: result001 (Student: student001, Score: 87, Grade: B)
- **Achievement**: achievement001 (Honor Roll)

## Next Steps

1. Review [README.md](README.md) for complete API documentation
2. Check [CHANGELOG.md](CHANGELOG.md) for implementation details
3. Read [DEMO_SUMMARY.md](DEMO_SUMMARY.md) for verification results
4. Explore Couchbase Console at http://localhost:8091
5. Try creating your own test data using the API endpoints

---

**Service Running Since:** December 15, 2025  
**Status:** Healthy ✅  
**Response Time:** <1ms (reads), 40-65ms (writes)
