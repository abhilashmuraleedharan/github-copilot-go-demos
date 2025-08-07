# School Management System - API Gateway Documentation

## Package Overview

The `handlers` package provides HTTP handlers for the School Management System API Gateway. The gateway acts as a central entry point for all microservices in the system, providing request routing, service discovery, and unified API access.

## Architecture

### Gateway Pattern
The API Gateway follows a proxy pattern where incoming requests are routed to appropriate backend microservices based on URL patterns. It supports full HTTP method proxying (GET, POST, PUT, DELETE) and maintains request/response integrity while providing centralized logging and monitoring.

### Service Discovery
Services are discovered through environment variables or default to localhost for development. In production Kubernetes deployments, service URLs are automatically configured through Helm templates.

### Supported Services
- **Student Service** (Port 8081): Student data management
- **Teacher Service** (Port 8082): Teacher data management  
- **Academic Service** (Port 8083): Academic records and classes
- **Achievement Service** (Port 8084): Achievements and badges

## URL Routing

| Gateway Route | Target Service | Description |
|---------------|---------------|-------------|
| `/api/v1/students/*` | Student Service | Student management operations |
| `/api/v1/teachers/*` | Teacher Service | Teacher management operations |
| `/api/v1/academics/*` | Academic Service | Academic records management |
| `/api/v1/classes/*` | Academic Service | Class management operations |
| `/api/v1/achievements/*` | Achievement Service | Achievement management |
| `/api/v1/badges/*` | Achievement Service | Badge management operations |
| `/health` | Gateway | Gateway health check |

## API Reference

### GatewayHandler

The main handler struct that manages HTTP requests for the API Gateway service.

```go
type GatewayHandler struct {
    studentServiceURL     string
    teacherServiceURL     string
    academicServiceURL    string
    achievementServiceURL string
}
```

#### Configuration

Service URLs are configured through environment variables:

| Environment Variable | Default Value | Description |
|---------------------|---------------|-------------|
| `STUDENT_SERVICE_URL` | `http://localhost:8081` | Student service endpoint |
| `TEACHER_SERVICE_URL` | `http://localhost:8082` | Teacher service endpoint |
| `ACADEMIC_SERVICE_URL` | `http://localhost:8083` | Academic service endpoint |
| `ACHIEVEMENT_SERVICE_URL` | `http://localhost:8084` | Achievement service endpoint |

### Constructor

#### NewGatewayHandler()

Creates a new GatewayHandler instance with service URLs configured from environment variables or sensible defaults.

**Returns:**
- `*GatewayHandler`: Configured gateway handler instance

**Example:**
```go
handler := NewGatewayHandler()
router.GET("/health", handler.HealthCheck)
```

## Endpoints

### Health Check

#### GET /health

Provides a health check endpoint for the API Gateway.

**Purpose:**
- Monitor gateway operational status
- Load balancer health checks
- Kubernetes readiness/liveness probes

**Response Format:**
```json
{
  "status": "healthy",
  "service": "api-gateway", 
  "services": {
    "student": "http://localhost:8081",
    "teacher": "http://localhost:8082",
    "academic": "http://localhost:8083",
    "achievement": "http://localhost:8084"
  }
}
```

**Status Codes:**
- `200 OK`: Gateway is healthy and operational

**Example:**
```bash
curl http://localhost:8080/health
```

### Student Service Proxy

#### * /api/v1/students/*

Proxies HTTP requests to the Student Service.

**Supported Methods:** GET, POST, PUT, DELETE  
**Target Service:** Student Service (Port 8081)

**URL Mapping:**
- `/api/v1/students` → `{STUDENT_SERVICE_URL}/api/v1/students`
- `/api/v1/students/{id}` → `{STUDENT_SERVICE_URL}/api/v1/students/{id}`

**Request Flow:**
1. Extract request path, query parameters, headers, and body
2. Construct target URL with Student Service base URL
3. Forward request to Student Service
4. Return response with original status code and headers

**Error Responses:**
- `502 Bad Gateway`: Student Service is unavailable
- `500 Internal Server Error`: Failed to create proxy request

**Examples:**
```bash
# Get all students
curl http://localhost:8080/api/v1/students

# Create new student
curl -X POST http://localhost:8080/api/v1/students \
  -H "Content-Type: application/json" \
  -d '{"first_name":"John","last_name":"Doe","email":"john@school.edu"}'

# Get specific student
curl http://localhost:8080/api/v1/students/student-001
```

### Teacher Service Proxy

#### * /api/v1/teachers/*

Proxies HTTP requests to the Teacher Service.

**Supported Methods:** GET, POST, PUT, DELETE  
**Target Service:** Teacher Service (Port 8082)

**URL Mapping:**
- `/api/v1/teachers` → `{TEACHER_SERVICE_URL}/api/v1/teachers`
- `/api/v1/teachers/{id}` → `{TEACHER_SERVICE_URL}/api/v1/teachers/{id}`

**Examples:**
```bash
# Get all teachers
curl http://localhost:8080/api/v1/teachers

# Create new teacher
curl -X POST http://localhost:8080/api/v1/teachers \
  -H "Content-Type: application/json" \
  -d '{"first_name":"Dr. Jane","last_name":"Smith","department":"Mathematics"}'
```

### Academic Service Proxy

#### * /api/v1/academics/* and /api/v1/classes/*

Proxies HTTP requests to the Academic Service with intelligent routing.

**Supported Methods:** GET, POST, PUT, DELETE  
**Target Service:** Academic Service (Port 8083)

**URL Mapping:**
- `/api/v1/academics/*` → `{ACADEMIC_SERVICE_URL}/api/v1/academics`
- `/api/v1/classes/*` → `{ACADEMIC_SERVICE_URL}/api/v1/classes`

**Routing Logic:**
The Academic Service handles both academic records and class management through different endpoints on the same service instance.

**Examples:**
```bash
# Academic records
curl http://localhost:8080/api/v1/academics

# Class management
curl http://localhost:8080/api/v1/classes
```

### Achievement Service Proxy

#### * /api/v1/achievements/* and /api/v1/badges/*

Proxies HTTP requests to the Achievement Service with intelligent routing.

**Supported Methods:** GET, POST, PUT, DELETE  
**Target Service:** Achievement Service (Port 8084)

**URL Mapping:**
- `/api/v1/achievements/*` → `{ACHIEVEMENT_SERVICE_URL}/api/v1/achievements`
- `/api/v1/badges/*` → `{ACHIEVEMENT_SERVICE_URL}/api/v1/badges`

**Routing Logic:**
The Achievement Service manages both student achievements and badge systems through different endpoints on the same service instance.

**Examples:**
```bash
# Student achievements
curl http://localhost:8080/api/v1/achievements

# Badge management
curl http://localhost:8080/api/v1/badges
```

## Internal Implementation

### Request Proxying

The `proxyRequest` method handles the low-level details of HTTP request proxying:

#### Process Flow
1. **URL Construction**: Replaces gateway paths with service-specific paths
2. **Query Preservation**: Maintains query parameters from original request
3. **Body Forwarding**: Reads and forwards request body for POST/PUT requests
4. **Header Copying**: Copies all request headers to target request
5. **Request Execution**: Executes HTTP request to target service
6. **Response Streaming**: Streams response back to client with original status and headers

#### Error Handling
- Logs all proxy failures for monitoring and debugging
- Returns `500 Internal Server Error` for request creation failures
- Returns `502 Bad Gateway` for service communication failures

#### Performance Considerations
- Uses `io.Copy` for efficient response streaming
- Reuses HTTP client instance
- Preserves original request/response headers for compatibility

#### Security Features
- Forwards all headers including authentication tokens
- Does not modify or inspect request/response bodies
- Maintains end-to-end encryption for HTTPS services

## Error Responses

### Standard Error Format

All error responses follow a consistent format:

```json
{
  "error": "Error description",
  "status": "error"
}
```

### Common Error Scenarios

| Status Code | Scenario | Description |
|-------------|----------|-------------|
| `500 Internal Server Error` | Request creation failure | Failed to create proxy request |
| `502 Bad Gateway` | Service unavailable | Target service is not reachable |
| `404 Not Found` | Invalid route | No matching route for the request |

## Monitoring and Logging

### Health Monitoring
- Built-in health check endpoint at `/health`
- Service URL validation and connectivity status
- Integration with Kubernetes health probes

### Logging Strategy
- Structured logging using Logrus
- Error-level logging for all proxy failures
- Request/response tracking for debugging

### Metrics (Future Enhancement)
Ready for integration with:
- Prometheus metrics collection
- Request count and latency tracking
- Error rate monitoring
- Service dependency mapping

## Development and Testing

### Local Development
```bash
# Set environment variables for local services
export STUDENT_SERVICE_URL=http://localhost:8081
export TEACHER_SERVICE_URL=http://localhost:8082
export ACADEMIC_SERVICE_URL=http://localhost:8083
export ACHIEVEMENT_SERVICE_URL=http://localhost:8084

# Run the gateway
go run main.go
```

### Testing
```bash
# Test health endpoint
curl http://localhost:8080/health

# Test service proxying
curl http://localhost:8080/api/v1/students
```

### Kubernetes Deployment
In Kubernetes environments, service URLs are automatically configured:
```yaml
env:
  - name: STUDENT_SERVICE_URL
    value: "http://school-mgmt-student-service:8081"
  - name: TEACHER_SERVICE_URL
    value: "http://school-mgmt-teacher-service:8082"
```

## Security Considerations

### Container Security
- Runs as non-root user
- Read-only filesystem capability
- Minimal Alpine Linux base image

### Network Security
- Kubernetes network policies for service isolation
- TLS termination at ingress level
- Secure secret management for credentials

### Future Security Enhancements
- JWT token validation
- Rate limiting and throttling
- Request/response logging for audit trails
- Circuit breaker pattern for fault tolerance
