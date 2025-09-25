# [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-09-24
# School Management Microservice

A comprehensive school management microservice built with Go and Couchbase, designed to handle operations for students, teachers, classes, academics, and achievements.

## Features

- **Complete Student Management**: CRUD operations with validation and pagination
- **Scalable Architecture**: Repository pattern with service layer abstraction
- **Database Integration**: Couchbase NoSQL database with optimized indexing
- **Configurable Deployment**: Environment-based configuration management
- **Production Ready**: Docker containerization with health checks
- **High Performance**: Designed for 200 TPS peak load handling
- **API Documentation**: RESTful API with standardized response formats

## Architecture

```
├── cmd/                    # Application entry point
├── internal/
│   ├── config/            # Configuration management
│   ├── models/            # Data models and DTOs
│   ├── repository/        # Data access layer
│   ├── service/           # Business logic layer
│   └── handler/           # HTTP handlers
├── scripts/               # Deployment scripts
├── docker-compose.yml     # Full stack deployment
├── Dockerfile            # Container definition
└── README.md             # This file
```

## Prerequisites

Before you begin, ensure you have the following installed:

- **Docker**: Version 20.10 or higher
- **Docker Compose**: Version 2.0 or higher
- **Git**: For cloning the repository

### System Requirements

- **Memory**: Minimum 4GB RAM (8GB recommended)
- **Storage**: At least 2GB free disk space
- **Network**: Ports 8080, 8091-8097 should be available

## Quick Start

Follow these step-by-step instructions to launch the service using Docker Compose.

### Step 1: Clone the Repository

```bash
git clone <repository-url>
cd demo9_school_microservice
```

### Step 2: Review Environment Configuration

Check the environment configuration file:

```bash
cat .env.example
```

The default configuration includes:
- Server runs on port 8080
- Couchbase runs on default ports (8091-8097)
- Default credentials: Administrator/password

### Step 3: Make Scripts Executable (Linux/macOS)

```bash
chmod +x scripts/couchbase-init.sh
```

### Step 4: Launch the Service Stack

Start all services using Docker Compose:

```bash
docker-compose up -d
```

This command will:
1. Pull required Docker images (Couchbase, Nginx, Alpine)
2. Build the Go microservice image
3. Start Couchbase database server
4. Initialize the database with required buckets
5. Start the school management service
6. Start Nginx reverse proxy (optional)

### Step 5: Monitor Service Startup

Watch the logs to ensure all services start correctly:

```bash
# View all logs
docker-compose logs -f

# View specific service logs
docker-compose logs -f school-service
docker-compose logs -f couchbase
```

### Step 6: Wait for Services to be Ready

The services need time to initialize:
- **Couchbase**: ~2-3 minutes for complete startup
- **School Service**: ~30-60 seconds after Couchbase is ready

Check service status:

```bash
# Check container status
docker-compose ps

# Check individual health
curl http://localhost:8080/health
```

## Service Verification

Once the services are running, verify they're working correctly:

### 1. Health Check

```bash
curl http://localhost:8080/health
```

Expected response:
```json
{
  "status": "healthy",
  "service": "School Management Microservice",
  "version": "v1.0.0",
  "timestamp": "2025-09-24T10:30:00Z"
}
```

### 2. Service Information

```bash
curl http://localhost:8080/
```

### 3. API Endpoint Test

Create a test student:

```bash
curl -X POST http://localhost:8080/api/v1/students \
  -H "Content-Type: application/json" \
  -d '{
    "first_name": "John",
    "last_name": "Doe",
    "email": "john.doe@school.edu",
    "date_of_birth": "2010-05-15T00:00:00Z",
    "grade": "8",
    "address": "123 Main St",
    "phone": "1234567890",
    "parent_name": "Jane Doe",
    "parent_phone": "0987654321"
  }'
```

Retrieve all students:

```bash
curl http://localhost:8080/api/v1/students
```

## API Documentation

### Student Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/health` | Service health check |
| GET | `/api/v1/students` | List all students (paginated) |
| POST | `/api/v1/students` | Create new student |
| GET | `/api/v1/students/{id}` | Get student by ID |
| PUT | `/api/v1/students/{id}` | Update student |
| DELETE | `/api/v1/students/{id}` | Delete student |
| GET | `/api/v1/students/grade/{grade}` | Get students by grade |

### Query Parameters

- `page`: Page number for pagination (default: 1)
- `page_size`: Number of items per page (default: 10, max: 100)

### Response Format

All API responses follow this standard format:

```json
{
  "success": true,
  "message": "Operation completed successfully",
  "data": { ... },
  "error": null
}
```

Paginated responses include additional fields:
```json
{
  "success": true,
  "message": "Students retrieved successfully",
  "data": [...],
  "total": 100,
  "page": 1,
  "page_size": 10,
  "total_pages": 10
}
```

## Configuration

The service supports configuration through environment variables:

### Server Configuration
- `SERVER_HOST`: Server bind address (default: 0.0.0.0)
- `SERVER_PORT`: Server port (default: 8080)
- `SERVER_READ_TIMEOUT`: Request read timeout (default: 15s)
- `SERVER_WRITE_TIMEOUT`: Response write timeout (default: 15s)

### Database Configuration
- `COUCHBASE_HOST`: Couchbase server host (default: couchbase)
- `COUCHBASE_PORT`: Couchbase server port (default: 8091)
- `COUCHBASE_USERNAME`: Database username (default: Administrator)
- `COUCHBASE_PASSWORD`: Database password (default: password)
- `COUCHBASE_BUCKET_NAME`: Database bucket name (default: school)

## Troubleshooting

### Common Issues

1. **Port Conflicts**
   ```bash
   # Check if ports are in use
   netstat -tulpn | grep :8080
   netstat -tulpn | grep :8091
   ```

2. **Couchbase Not Ready**
   ```bash
   # Wait longer for Couchbase initialization
   docker-compose logs couchbase
   ```

3. **Service Won't Start**
   ```bash
   # Check container logs
   docker-compose logs school-service
   ```

4. **Database Connection Issues**
   ```bash
   # Verify Couchbase is accessible
   curl http://localhost:8091/ui/index.html
   ```

### Cleanup

To stop and remove all services:

```bash
# Stop services
docker-compose down

# Remove volumes (WARNING: This deletes all data)
docker-compose down -v

# Remove images
docker-compose down --rmi all
```

## Performance Considerations

- **Memory**: Couchbase requires at least 1GB RAM
- **Connections**: Service supports concurrent requests up to 200 TPS
- **Indexing**: Database indexes are automatically created for optimal query performance
- **Caching**: Consider adding Redis for high-traffic scenarios

## Development

For local development without Docker:

### Prerequisites
- Go 1.21 or higher
- Couchbase Server 7.2.0 or higher

### Setup
```bash
# Install dependencies
go mod download

# Copy environment file
cp .env.example .env

# Run the service
go run cmd/main.go
```

## Security Notes

- Default credentials are for development only
- Change passwords for production deployment
- Consider adding authentication middleware
- Use HTTPS in production environments

## Support

For issues and questions:
1. Check the logs: `docker-compose logs`
2. Verify configuration: Review environment variables
3. Check connectivity: Ensure all ports are accessible
4. Review documentation: API endpoints and expected formats

## License

This project is part of the GitHub Copilot Go Demos repository.