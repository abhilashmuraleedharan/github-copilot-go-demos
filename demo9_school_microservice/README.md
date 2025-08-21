# School Microservice

A distributed microservice architecture for managing school operations including students, teachers, classes, academics, and achievements.

## Architecture

This project implements a domain-oriented microservices architecture with the following services:

- **Students Service** (Port 8081) - Manages student data and operations
- **Teachers Service** (Port 8082) - Manages teacher data and operations  
- **Classes Service** (Port 8083) - Manages class data and operations
- **Academics Service** (Port 8084) - Manages academic records and grades
- **Achievements Service** (Port 8085) - Manages student achievements and awards

## Technology Stack

- **Language**: Go 1.21
- **Database**: Couchbase
- **Containerization**: Docker & Docker Compose
- **HTTP Router**: Gorilla Mux
- **CORS**: rs/cors

## Quick Start with Docker Compose

1. Clone the repository
2. Navigate to the demo9_school_microservice directory
3. Copy `.env.example` to `.env` and configure your Couchbase settings
4. Run: `docker-compose up --build`

## API Endpoints

Each service exposes RESTful endpoints on their respective ports:

### Students Service (8081)
- GET /students - List all students
- GET /students/{id} - Get student by ID
- POST /students - Create new student
- PUT /students/{id} - Update student
- DELETE /students/{id} - Delete student

### Teachers Service (8082)
- GET /teachers - List all teachers
- GET /teachers/{id} - Get teacher by ID
- POST /teachers - Create new teacher
- PUT /teachers/{id} - Update teacher
- DELETE /teachers/{id} - Delete teacher

### Classes Service (8083)
- GET /classes - List all classes
- GET /classes/{id} - Get class by ID
- POST /classes - Create new class
- PUT /classes/{id} - Update class
- DELETE /classes/{id} - Delete class

### Academics Service (8084)
- GET /academics - List all academic records
- GET /academics/{id} - Get academic record by ID
- GET /academics/student/{studentId} - Get academics by student
- POST /academics - Create new academic record
- PUT /academics/{id} - Update academic record
- DELETE /academics/{id} - Delete academic record

### Achievements Service (8085)
- GET /achievements - List all achievements
- GET /achievements/{id} - Get achievement by ID
- GET /achievements/student/{studentId} - Get achievements by student
- POST /achievements - Create new achievement
- PUT /achievements/{id} - Update achievement
- DELETE /achievements/{id} - Delete achievement

## Configuration

All services use environment variables for configuration. See `.env.example` for required variables.

## Health Checks

Each service provides a health check endpoint at `/health`.
