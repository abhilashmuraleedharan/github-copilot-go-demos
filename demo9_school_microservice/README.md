# School Management Microservice

A Go-based microservice for managing school operations including students, teachers, classes, academic enrollments, exams, and achievements. Built using a Data-Centric CRUD Architecture with Couchbase as the data store.

## Architecture

This microservice follows **Approach 2: Data-Centric CRUD Architecture with Service Layer**:

- **Presentation Layer**: REST API handlers using Gorilla Mux
- **Service Layer**: Business logic, validation, and orchestration
- **Repository Layer**: Data access abstraction over Couchbase
- **Data Store**: Couchbase Community Server

### Key Features

- ✅ RESTful CRUD APIs for all entities
- ✅ Configurable Couchbase credentials via environment variables
- ✅ Automatic grade calculation for exam results
- ✅ Referential integrity validation
- ✅ Health check endpoint
- ✅ Request logging middleware
- ✅ Docker Compose deployment
- ✅ Designed for ~200 TPS peak load

## Project Structure

```
demo9_school_microservice/
├── config/              # Configuration management
│   └── config.go
├── models/              # Domain entity models
│   └── models.go
├── repository/          # Data access layer
│   └── repository.go
├── service/             # Business logic layer
│   └── service.go
├── handlers/            # HTTP request handlers
│   ├── handlers.go
│   └── handlers_test.go # Unit tests for handlers
├── scripts/             # Automation scripts
│   ├── test-coverage.ps1  # Windows coverage script
│   └── test-coverage.sh   # Linux/Mac coverage script
├── .github/
│   └── workflows/
│       └── test.yml     # CI/CD testing workflow
├── helm/                # Kubernetes Helm charts
├── main.go              # Application entry point
├── go.mod               # Go module definition
├── Makefile             # Build automation
├── Dockerfile           # Container image
├── docker-compose.yml   # Multi-container setup
├── README.md            # This file
├── CHANGELOG.md         # Version history
└── DESIGN.md            # Architecture documentation
```

## API Endpoints

### Health Check
- `GET /health` - Service health status

### Students
- `POST /api/students` - Create a new student
- `GET /api/students/{id}` - Get student by ID
- `PUT /api/students/{id}` - Update student
- `DELETE /api/students/{id}` - Delete student

### Teachers
- `POST /api/teachers` - Create a new teacher
- `GET /api/teachers/{id}` - Get teacher by ID
- `PUT /api/teachers/{id}` - Update teacher
- `DELETE /api/teachers/{id}` - Delete teacher

### Classes
- `POST /api/classes` - Create a new class
- `GET /api/classes/{id}` - Get class by ID
- `PUT /api/classes/{id}` - Update class
- `DELETE /api/classes/{id}` - Delete class

### Academic Enrollments
- `POST /api/academics` - Create enrollment
- `GET /api/academics/{id}` - Get enrollment by ID
- `PUT /api/academics/{id}` - Update enrollment
- `DELETE /api/academics/{id}` - Delete enrollment

### Exams
- `POST /api/exams` - Create a new exam
- `GET /api/exams/{id}` - Get exam by ID
- `PUT /api/exams/{id}` - Update exam
- `DELETE /api/exams/{id}` - Delete exam

### Exam Results
- `POST /api/exam-results` - Create exam result
- `GET /api/exam-results/{id}` - Get exam result by ID
- `PUT /api/exam-results/{id}` - Update exam result
- `DELETE /api/exam-results/{id}` - Delete exam result

### Achievements
- `POST /api/achievements` - Create achievement
- `GET /api/achievements/{id}` - Get achievement by ID
- `PUT /api/achievements/{id}` - Update achievement
- `DELETE /api/achievements/{id}` - Delete achievement

## Configuration

The service is configured via environment variables:

| Variable | Default | Description |
|----------|---------|-------------|
| `SERVER_PORT` | 8080 | HTTP server port |
| `COUCHBASE_CONNECTION_STRING` | couchbase://localhost | Couchbase cluster URL |
| `COUCHBASE_USERNAME` | Administrator | Couchbase username |
| `COUCHBASE_PASSWORD` | password | Couchbase password |
| `COUCHBASE_BUCKET` | school | Couchbase bucket name |

## Getting Started with Docker Compose

### Prerequisites

- Docker (version 20.10 or later)
- Docker Compose (version 2.0 or later)
- 4GB RAM minimum (for Couchbase)

### Step-by-Step Launch Instructions

#### Step 1: Navigate to Project Directory

```powershell
cd d:\DEMO_15_DEC_25\school_mgmt\github-copilot-go-demos\demo9_school_microservice
```

#### Step 2: Start Services with Docker Compose

```powershell
docker-compose up -d
```

This command will:
- Pull the Couchbase Community Server image (if not cached)
- Build the Go microservice image
- Start Couchbase server
- Initialize Couchbase cluster with default credentials
- Create the `school` bucket
- Create a primary index
- Start the school management service

#### Step 3: Monitor Service Startup

Watch the logs to ensure services start successfully:

```powershell
docker-compose logs -f
```

Wait for these messages:
- ✅ `school-couchbase-init_1` shows "Couchbase initialization complete!"
- ✅ `school-service_1` shows "Server listening on :8080"

Press `Ctrl+C` to exit log viewing (services continue running).

#### Step 4: Verify Services are Running

Check service health:

```powershell
# Check all containers are running
docker-compose ps

# Test health endpoint
curl http://localhost:8080/health
```

Expected response:
```json
{"status":"healthy"}
```

#### Step 5: Access Services

- **School Management API**: http://localhost:8080
- **Couchbase Web Console**: http://localhost:8091
  - Username: `Administrator`
  - Password: `password`

## Testing the API

### Example 1: Create a Teacher

```powershell
curl -X POST http://localhost:8080/api/teachers `
  -H "Content-Type: application/json" `
  -d '{
    "id": "teacher001",
    "firstName": "John",
    "lastName": "Smith",
    "email": "john.smith@school.edu",
    "subject": "Mathematics",
    "hireDate": "2024-01-15T00:00:00Z"
  }'
```

### Example 2: Create a Class

```powershell
curl -X POST http://localhost:8080/api/classes `
  -H "Content-Type: application/json" `
  -d '{
    "id": "class001",
    "name": "Algebra I",
    "subject": "Mathematics",
    "teacherId": "teacher001",
    "schedule": "MWF 10:00-11:00",
    "maxStudents": 30
  }'
```

### Example 3: Create a Student

```powershell
curl -X POST http://localhost:8080/api/students `
  -H "Content-Type: application/json" `
  -d '{
    "id": "student001",
    "firstName": "Alice",
    "lastName": "Johnson",
    "email": "alice.j@school.edu",
    "dateOfBirth": "2008-05-20T00:00:00Z",
    "enrollmentDate": "2024-09-01T00:00:00Z",
    "grade": "10"
  }'
```

### Example 4: Enroll Student in Class

```powershell
curl -X POST http://localhost:8080/api/academics `
  -H "Content-Type: application/json" `
  -d '{
    "id": "academic001",
    "studentId": "student001",
    "classId": "class001",
    "enrollmentDate": "2024-09-01T00:00:00Z",
    "status": "active"
  }'
```

### Example 5: Create an Exam

```powershell
curl -X POST http://localhost:8080/api/exams `
  -H "Content-Type: application/json" `
  -d '{
    "id": "exam001",
    "classId": "class001",
    "title": "Midterm Exam",
    "description": "Algebra I Midterm covering chapters 1-5",
    "examDate": "2024-11-15T09:00:00Z",
    "maxScore": 100
  }'
```

### Example 6: Record Exam Result

```powershell
curl -X POST http://localhost:8080/api/exam-results `
  -H "Content-Type: application/json" `
  -d '{
    "id": "result001",
    "examId": "exam001",
    "studentId": "student001",
    "score": 87,
    "takenDate": "2024-11-15T09:00:00Z"
  }'
```

Note: The grade will be automatically calculated (87/100 = B).

### Example 7: Create an Achievement

```powershell
curl -X POST http://localhost:8080/api/achievements `
  -H "Content-Type: application/json" `
  -d '{
    "id": "achievement001",
    "studentId": "student001",
    "title": "Honor Roll",
    "description": "Achieved honor roll for Fall 2024 semester",
    "awardDate": "2024-12-15T00:00:00Z",
    "category": "academic"
  }'
```

### Example 8: Retrieve Data

```powershell
# Get student details
curl http://localhost:8080/api/students/student001

# Get teacher details
curl http://localhost:8080/api/teachers/teacher001

# Get exam result
curl http://localhost:8080/api/exam-results/result001
```

## Managing the Service

### View Logs

```powershell
# All services
docker-compose logs -f

# Specific service
docker-compose logs -f school-service
docker-compose logs -f couchbase
```

### Stop Services

```powershell
docker-compose stop
```

### Start Stopped Services

```powershell
docker-compose start
```

### Restart Services

```powershell
docker-compose restart
```

### Stop and Remove All Containers

```powershell
docker-compose down
```

### Stop and Remove Everything (including volumes)

```powershell
docker-compose down -v
```

**Warning**: This deletes all data in Couchbase!

## Troubleshooting

### Issue: Service won't start

**Solution**: Check if ports are already in use:
```powershell
netstat -ano | findstr :8080
netstat -ano | findstr :8091
```

### Issue: Couchbase initialization fails

**Solution**: Ensure at least 4GB RAM is available to Docker. Check Docker Desktop settings.

### Issue: Can't connect to Couchbase

**Solution**: Wait 30-60 seconds after `docker-compose up` for full initialization. Check logs:
```powershell
docker-compose logs couchbase-init
```

### Issue: API returns 500 errors

**Solution**: Verify Couchbase bucket exists:
1. Open http://localhost:8091
2. Login (Administrator/password)
3. Check "Buckets" section for "school" bucket

### Issue: Need to rebuild after code changes

**Solution**: Rebuild and restart the service:
```powershell
docker-compose up -d --build school-service
```

## Running Tests

### Quick Test Run

```bash
# Run all tests
go test ./... -v

# Run specific package tests
go test ./handlers -v
```

### Test Coverage

#### Using Makefile

```bash
# Run tests with coverage
make test-coverage

# Generate HTML coverage report
make test-coverage-html

# View coverage report in browser
make coverage-report
```

#### Using Scripts

**Windows (PowerShell):**
```powershell
.\scripts\test-coverage.ps1
```

**Linux/Mac (Bash):**
```bash
chmod +x scripts/test-coverage.sh
./scripts/test-coverage.sh
```

#### Manual Coverage Commands

```bash
# Generate coverage profile
go test ./... -coverprofile=coverage.out -covermode=atomic

# View coverage summary
go tool cover -func=coverage.out

# Generate HTML report
go tool cover -html=coverage.out -o coverage.html

# Open in browser (Windows)
start coverage.html

# Open in browser (Mac)
open coverage.html

# Open in browser (Linux)
xdg-open coverage.html
```

### Test with Race Detection

```bash
# Run tests with race detector
go test ./... -v -race

# With coverage
go test ./... -v -race -coverprofile=coverage.out
```

### CI/CD Testing

Tests run automatically on:
- Push to main or develop branches
- Pull requests to main or develop

GitHub Actions workflow:
- Runs all tests with race detection
- Generates coverage reports
- Uploads to Codecov
- Archives coverage artifacts

View workflow: [.github/workflows/test.yml](.github/workflows/test.yml)

```powershell
docker-compose up -d --build
```

## Development

### Local Development (without Docker)

1. Install Go 1.21 or later
2. Install Couchbase locally or use Docker:
   ```powershell
   docker run -d --name couchbase -p 8091-8096:8091-8096 -p 11210:11210 couchbase:community-7.2.4
   ```
3. Set up Couchbase manually via Web Console
4. Install dependencies:
   ```powershell
   go mod download
   ```
5. Run the service:
   ```powershell
   go run main.go
   ```

### Environment Variables for Local Development

Create a `.env` file or set environment variables:

```powershell
$env:SERVER_PORT="8080"
$env:COUCHBASE_CONNECTION_STRING="couchbase://localhost"
$env:COUCHBASE_USERNAME="Administrator"
$env:COUCHBASE_PASSWORD="password"
$env:COUCHBASE_BUCKET="school"
```

## Performance

- Designed to handle **~200 TPS** during peak hours (10 AM - 5 PM)
- HTTP timeouts: 15s read/write, 60s idle
- Couchbase connection pooling enabled
- Health checks for load balancer integration

## Dependencies

- `github.com/couchbase/gocb/v2 v2.7.2` - Couchbase Go SDK
- `github.com/gorilla/mux v1.8.1` - HTTP router

## License

See LICENSE file in the repository root.

## Contributing

This is a demo project. For production use, consider adding:
- Authentication and authorization
- Input sanitization and validation enhancements
- Pagination for list endpoints
- Rate limiting
- Metrics and monitoring
- Comprehensive unit and integration tests
- CI/CD pipelines
