# School Microservice API Documentation

## Overview

The School Microservice is a domain-oriented microservice architecture that provides comprehensive APIs for managing school operations. The system consists of five independent services, each handling a specific domain within the school ecosystem.

## Service Architecture

### Service Endpoints

| Service | Port | Base URL | Description |
|---------|------|----------|-------------|
| Students | 8081 | `/students` | Student enrollment and management |
| Teachers | 8082 | `/teachers` | Faculty and staff management |
| Classes | 8083 | `/classes` | Course and classroom management |
| Academics | 8084 | `/academics` | Grade and academic record management |
| Achievements | 8085 | `/achievements` | Student achievement and award tracking |

### Health Endpoints

All services provide health check endpoints:
- `GET /{service}/health` - Returns service health status

## API Reference

### 1. Students Service (`localhost:8081`)

#### Student Model
```go
// Student represents a student in the school system
type Student struct {
    ID          string    `json:"id"`          // Unique student identifier (auto-generated: STU20250821...)
    FirstName   string    `json:"firstName"`   // Student's first name
    LastName    string    `json:"lastName"`    // Student's last name
    Email       string    `json:"email"`       // Student's email address
    DateOfBirth time.Time `json:"dateOfBirth"` // Student's date of birth
    Grade       string    `json:"grade"`       // Current grade level (e.g., "9", "10", "11", "12")
    Address     string    `json:"address"`     // Home address
    Phone       string    `json:"phone"`       // Contact phone number
    EnrollDate  time.Time `json:"enrollDate"`  // Enrollment date (auto-set if not provided)
    Status      string    `json:"status"`      // active, inactive, graduated (default: "active")
    CreatedAt   time.Time `json:"createdAt"`   // Record creation timestamp
    UpdatedAt   time.Time `json:"updatedAt"`   // Last update timestamp
}
```

#### Endpoints

##### `GET /students`
**Description**: Retrieve all students with concurrency-safe implementation
**Response**: Array of Student objects
**Features**:
- Context-aware timeout and cancellation support
- Response caching (60 seconds)
- Streaming JSON encoder for large datasets
- Error handling without exposing internal details

```json
[
  {
    "id": "STU20250821140530",
    "firstName": "John",
    "lastName": "Doe",
    "email": "john.doe@school.edu",
    "dateOfBirth": "2008-05-15T00:00:00Z",
    "grade": "10",
    "address": "123 Main St, Anytown, USA",
    "phone": "+1-555-0123",
    "enrollDate": "2023-08-15T00:00:00Z",
    "status": "active",
    "createdAt": "2025-08-21T14:05:30Z",
    "updatedAt": "2025-08-21T14:05:30Z"
  }
]
```

##### `GET /students/{id}`
**Description**: Retrieve a specific student by ID
**Parameters**: 
- `id` (path) - Student ID
**Response**: Student object or 404 if not found

##### `POST /students`
**Description**: Create a new student record
**Request Body**: Student object (ID will be auto-generated if not provided)
**Response**: Created student object with generated ID and timestamps
**Status Code**: 201 Created

##### `PUT /students/{id}`
**Description**: Update an existing student record
**Parameters**: 
- `id` (path) - Student ID
**Request Body**: Student object with updated fields
**Response**: Updated student object

##### `DELETE /students/{id}`
**Description**: Delete a student record
**Parameters**: 
- `id` (path) - Student ID
**Response**: 204 No Content on success

---

### 2. Teachers Service (`localhost:8082`)

#### Teacher Model
```go
// Teacher represents a teacher in the school system
type Teacher struct {
    ID          string    `json:"id"`          // Unique teacher identifier
    FirstName   string    `json:"firstName"`   // Teacher's first name
    LastName    string    `json:"lastName"`    // Teacher's last name
    Email       string    `json:"email"`       // Teacher's email address
    Phone       string    `json:"phone"`       // Contact phone number
    Department  string    `json:"department"`  // Academic department (e.g., "Mathematics", "Science")
    Subject     string    `json:"subject"`     // Primary subject taught
    HireDate    time.Time `json:"hireDate"`    // Date of hire
    Salary      float64   `json:"salary"`      // Annual salary
    Address     string    `json:"address"`     // Home address
    Status      string    `json:"status"`      // active, inactive, retired
    CreatedAt   time.Time `json:"createdAt"`   // Record creation timestamp
    UpdatedAt   time.Time `json:"updatedAt"`   // Last update timestamp
}
```

#### Endpoints

##### `GET /teachers`
**Description**: Retrieve all teachers
**Response**: Array of Teacher objects

##### `GET /teachers/{id}`
**Description**: Retrieve a specific teacher by ID
**Parameters**: 
- `id` (path) - Teacher ID
**Response**: Teacher object or 404 if not found

##### `POST /teachers`
**Description**: Create a new teacher record
**Request Body**: Teacher object
**Response**: Created teacher object
**Status Code**: 201 Created

##### `PUT /teachers/{id}`
**Description**: Update an existing teacher record
**Parameters**: 
- `id` (path) - Teacher ID
**Request Body**: Teacher object with updated fields
**Response**: Updated teacher object

##### `DELETE /teachers/{id}`
**Description**: Delete a teacher record
**Parameters**: 
- `id` (path) - Teacher ID
**Response**: 204 No Content on success

---

### 3. Classes Service (`localhost:8083`)

#### Class Model
```go
// Class represents a class in the school system
type Class struct {
    ID          string    `json:"id"`          // Unique class identifier
    Name        string    `json:"name"`        // Class name (e.g., "Advanced Mathematics")
    Subject     string    `json:"subject"`     // Subject area
    TeacherID   string    `json:"teacherId"`   // Assigned teacher ID
    Grade       string    `json:"grade"`       // Grade level
    Room        string    `json:"room"`        // Classroom number/location
    Schedule    string    `json:"schedule"`    // Class schedule (e.g., "Mon,Wed,Fri 9:00-10:00")
    MaxStudents int       `json:"maxStudents"` // Maximum enrollment capacity
    StudentIDs  []string  `json:"studentIds"`  // Array of enrolled student IDs
    Semester    string    `json:"semester"`    // Academic semester
    Year        int       `json:"year"`        // Academic year
    Status      string    `json:"status"`      // active, inactive, completed
    CreatedAt   time.Time `json:"createdAt"`   // Record creation timestamp
    UpdatedAt   time.Time `json:"updatedAt"`   // Last update timestamp
}
```

#### Endpoints

##### `GET /classes`
**Description**: Retrieve all classes
**Response**: Array of Class objects

##### `GET /classes/{id}`
**Description**: Retrieve a specific class by ID
**Parameters**: 
- `id` (path) - Class ID
**Response**: Class object or 404 if not found

##### `POST /classes`
**Description**: Create a new class
**Request Body**: Class object
**Response**: Created class object
**Status Code**: 201 Created

##### `PUT /classes/{id}`
**Description**: Update an existing class
**Parameters**: 
- `id` (path) - Class ID
**Request Body**: Class object with updated fields
**Response**: Updated class object

##### `DELETE /classes/{id}`
**Description**: Delete a class
**Parameters**: 
- `id` (path) - Class ID
**Response**: 204 No Content on success

---

### 4. Academics Service (`localhost:8084`)

#### Academic Model
```go
// Academic represents an academic record in the school system
type Academic struct {
    ID         string    `json:"id"`         // Unique academic record identifier
    StudentID  string    `json:"studentId"`  // Reference to student
    ClassID    string    `json:"classId"`    // Reference to class
    Subject    string    `json:"subject"`    // Subject name
    ExamType   string    `json:"examType"`   // midterm, final, quiz, assignment
    Score      float64   `json:"score"`      // Points scored
    MaxScore   float64   `json:"maxScore"`   // Maximum possible points
    Grade      string    `json:"grade"`      // Letter grade (A, B, C, D, F) - auto-calculated
    ExamDate   time.Time `json:"examDate"`   // Date of exam/assignment
    Semester   string    `json:"semester"`   // Academic semester
    Year       int       `json:"year"`       // Academic year
    TeacherID  string    `json:"teacherId"`  // Grading teacher
    Comments   string    `json:"comments"`   // Additional comments
    CreatedAt  time.Time `json:"createdAt"`  // Record creation timestamp
    UpdatedAt  time.Time `json:"updatedAt"`  // Last update timestamp
}
```

#### Grade Calculation Logic
The system automatically calculates letter grades based on percentage:
- A: 90-100%
- B: 80-89%
- C: 70-79%
- D: 60-69%
- F: Below 60%

#### Endpoints

##### `GET /academics`
**Description**: Retrieve all academic records
**Response**: Array of Academic objects

##### `GET /academics/{id}`
**Description**: Retrieve a specific academic record by ID
**Parameters**: 
- `id` (path) - Academic record ID
**Response**: Academic object or 404 if not found

##### `POST /academics`
**Description**: Create a new academic record
**Request Body**: Academic object (grade will be auto-calculated from score/maxScore)
**Response**: Created academic object with calculated grade
**Status Code**: 201 Created

##### `PUT /academics/{id}`
**Description**: Update an existing academic record
**Parameters**: 
- `id` (path) - Academic record ID
**Request Body**: Academic object with updated fields
**Response**: Updated academic object with recalculated grade

##### `DELETE /academics/{id}`
**Description**: Delete an academic record
**Parameters**: 
- `id` (path) - Academic record ID
**Response**: 204 No Content on success

---

### 5. Achievements Service (`localhost:8085`)

#### Achievement Model
```go
// Achievement represents an achievement in the school system
type Achievement struct {
    ID          string    `json:"id"`          // Unique achievement identifier
    StudentID   string    `json:"studentId"`   // Reference to student
    Title       string    `json:"title"`       // Achievement title
    Description string    `json:"description"` // Detailed description
    Category    string    `json:"category"`    // academic, sports, arts, community, leadership
    Level       string    `json:"level"`       // school, district, state, national, international
    AwardedBy   string    `json:"awardedBy"`   // Awarding organization
    AwardDate   time.Time `json:"awardDate"`   // Date achievement was earned
    Points      int       `json:"points"`      // Achievement points value
    Certificate string    `json:"certificate"` // URL to certificate document
    Status      string    `json:"status"`      // pending, approved, revoked
    TeacherID   string    `json:"teacherId"`   // Nominating teacher
    Comments    string    `json:"comments"`    // Additional comments
    CreatedAt   time.Time `json:"createdAt"`   // Record creation timestamp
    UpdatedAt   time.Time `json:"updatedAt"`   // Last update timestamp
}
```

#### Achievement Categories
- **academic**: Academic excellence, honor roll, etc.
- **sports**: Athletic achievements and competitions
- **arts**: Creative and artistic accomplishments
- **community**: Community service and volunteer work
- **leadership**: Leadership roles and responsibilities

#### Achievement Levels
- **school**: School-level recognition
- **district**: District-wide achievements
- **state**: State-level competitions and honors
- **national**: National recognition
- **international**: International achievements

#### Endpoints

##### `GET /achievements`
**Description**: Retrieve all achievements
**Response**: Array of Achievement objects

##### `GET /achievements/{id}`
**Description**: Retrieve a specific achievement by ID
**Parameters**: 
- `id` (path) - Achievement ID
**Response**: Achievement object or 404 if not found

##### `POST /achievements`
**Description**: Create a new achievement record
**Request Body**: Achievement object
**Response**: Created achievement object
**Status Code**: 201 Created

##### `PUT /achievements/{id}`
**Description**: Update an existing achievement record
**Parameters**: 
- `id` (path) - Achievement ID
**Request Body**: Achievement object with updated fields
**Response**: Updated achievement object

##### `DELETE /achievements/{id}`
**Description**: Delete an achievement record
**Parameters**: 
- `id` (path) - Achievement ID
**Response**: 204 No Content on success

---

## Common Response Formats

### Success Response
All successful requests return JSON with appropriate HTTP status codes:
- `200 OK` - Successful GET, PUT operations
- `201 Created` - Successful POST operations
- `204 No Content` - Successful DELETE operations

### Error Response
Error responses follow a consistent format:
```json
{
  "error": "Error message describing what went wrong",
  "status": 400
}
```

### Common HTTP Status Codes
- `400 Bad Request` - Invalid request format or missing required fields
- `404 Not Found` - Resource not found
- `500 Internal Server Error` - Server-side error

---

## Authentication & Security

### Current Implementation
- No authentication required (development/demo purposes)
- CORS enabled for cross-origin requests
- Input validation on all endpoints
- Error messages sanitized to prevent information disclosure

### Production Considerations
For production deployment, consider implementing:
- JWT token-based authentication
- Role-based access control (RBAC)
- Rate limiting
- Input sanitization and validation
- Audit logging
- HTTPS enforcement

---

## Configuration

### Environment Variables
Each service supports the following configuration via environment variables:

```env
# Database Configuration
COUCHBASE_HOST=localhost           # Couchbase server host
COUCHBASE_USERNAME=admin          # Database username
COUCHBASE_PASSWORD=password       # Database password
COUCHBASE_BUCKET=school          # Database bucket name

# Service Configuration
PORT=8081                        # Service port (varies by service)
LOG_LEVEL=info                   # Logging level
```

### Docker Compose
Services can be started together using Docker Compose:
```bash
docker-compose up -d
```

### Kubernetes Deployment
Use Helm for Kubernetes deployment:
```bash
helm install school-microservice ./helm/school-microservice
```

---

## Data Relationships

### Entity Relationships
- **Students** ↔ **Classes**: Many-to-many (students can enroll in multiple classes)
- **Teachers** ↔ **Classes**: One-to-many (one teacher per class)
- **Students** ↔ **Academics**: One-to-many (multiple academic records per student)
- **Students** ↔ **Achievements**: One-to-many (multiple achievements per student)
- **Classes** ↔ **Academics**: One-to-many (multiple academic records per class)

### Data Consistency
- Foreign key references are maintained at the application level
- Each service manages its own data domain
- Cross-service data consistency is eventual consistency model

---

## Performance Considerations

### Concurrency Features
- Context-aware operations with timeout support
- Streaming JSON encoding for large responses
- Connection pooling for database operations
- Response caching where appropriate

### Optimization Recommendations
- Use pagination for large result sets
- Implement result caching for frequently accessed data
- Consider database indexing for search operations
- Monitor and optimize slow queries

---

## Development & Testing

### Local Development
1. Install Go 1.21 or later
2. Install and configure Couchbase
3. Set environment variables
4. Run individual services: `go run main.go`

### API Testing
Use tools like curl, Postman, or automated testing frameworks to test the APIs:

```bash
# Test student creation
curl -X POST http://localhost:8081/students \
  -H "Content-Type: application/json" \
  -d '{"firstName":"John","lastName":"Doe","email":"john@school.edu","grade":"10"}'

# Test health endpoint
curl http://localhost:8081/health
```

### Health Monitoring
All services provide health endpoints that return service status and can be used for:
- Load balancer health checks
- Kubernetes liveness/readiness probes
- Service monitoring and alerting

---

## Support & Maintenance

### Logging
Services log important events and errors. In production:
- Configure structured logging (JSON format)
- Set appropriate log levels
- Implement log aggregation and monitoring

### Monitoring
Recommended monitoring metrics:
- Response times and latency
- Error rates and status codes
- Database connection health
- Resource utilization (CPU, memory)

### Troubleshooting
Common issues and solutions:
1. **Service won't start**: Check environment variables and database connectivity
2. **404 errors**: Verify service is running on expected port
3. **Database errors**: Confirm Couchbase is running and accessible
4. **Performance issues**: Check database indexes and query optimization
