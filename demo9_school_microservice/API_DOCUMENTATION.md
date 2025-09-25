# School Management Microservice API Documentation

<!-- [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-09-24 -->

This document provides comprehensive API documentation for the School Management Microservice, following Go documentation conventions and best practices.

## Table of Contents

- [Overview](#overview)
- [API Endpoints](#api-endpoints)
- [Data Models](#data-models)
- [Request/Response Formats](#requestresponse-formats)
- [Error Handling](#error-handling)
- [Authentication](#authentication)
- [Package Documentation](#package-documentation)

## Overview

The School Management Microservice provides a RESTful API for managing school data including students, teachers, classes, academic records, and achievements. The service is built with Go using the Gin framework and Couchbase as the database backend.

**Base URL:** `http://localhost:8080`  
**API Version:** v1  
**Content-Type:** `application/json`

## API Endpoints

### Health Check

#### GET /health
Returns the health status of the service.

**Response:**
```json
{
  "status": "healthy",
  "timestamp": "2025-09-24T10:00:00Z",
  "version": "1.0.0"
}
```

#### GET /health/ready
Returns readiness status for Kubernetes probes.

**Response:**
```json
{
  "status": "ready",
  "database": "connected",
  "timestamp": "2025-09-24T10:00:00Z"
}
```

### Students API

#### POST /api/v1/students
Creates a new student record.

**Request Body:**
```json
{
  "first_name": "John",
  "last_name": "Doe",
  "email": "john.doe@example.com",
  "date_of_birth": "2010-05-15T00:00:00Z",
  "grade": "8",
  "address": "123 Main St, City, State",
  "phone": "1234567890",
  "parent_name": "Jane Doe",
  "parent_phone": "0987654321"
}
```

**Response (201 Created):**
```json
{
  "success": true,
  "message": "Student created successfully",
  "data": {
    "id": "student_123",
    "first_name": "John",
    "last_name": "Doe",
    "email": "john.doe@example.com",
    "date_of_birth": "2010-05-15T00:00:00Z",
    "grade": "8",
    "address": "123 Main St, City, State",
    "phone": "1234567890",
    "parent_name": "Jane Doe",
    "parent_phone": "0987654321",
    "created_at": "2025-09-24T10:00:00Z",
    "updated_at": "2025-09-24T10:00:00Z"
  }
}
```

#### GET /api/v1/students/{id}
Retrieves a student by ID.

**Parameters:**
- `id` (path) - The student ID

**Response (200 OK):**
```json
{
  "success": true,
  "message": "Student retrieved successfully",
  "data": {
    "id": "student_123",
    "first_name": "John",
    "last_name": "Doe",
    "email": "john.doe@example.com",
    "date_of_birth": "2010-05-15T00:00:00Z",
    "grade": "8",
    "address": "123 Main St, City, State",
    "phone": "1234567890",
    "parent_name": "Jane Doe",
    "parent_phone": "0987654321",
    "created_at": "2025-09-24T10:00:00Z",
    "updated_at": "2025-09-24T10:00:00Z"
  }
}
```

#### GET /api/v1/students
Retrieves all students with pagination.

**Query Parameters:**
- `page` (optional) - Page number (default: 1)
- `page_size` (optional) - Number of items per page (default: 10, max: 100)

**Response (200 OK):**
```json
{
  "success": true,
  "message": "Students retrieved successfully",
  "data": [
    {
      "id": "student_123",
      "first_name": "John",
      "last_name": "Doe",
      "email": "john.doe@example.com",
      "date_of_birth": "2010-05-15T00:00:00Z",
      "grade": "8",
      "address": "123 Main St, City, State",
      "phone": "1234567890",
      "parent_name": "Jane Doe",
      "parent_phone": "0987654321",
      "created_at": "2025-09-24T10:00:00Z",
      "updated_at": "2025-09-24T10:00:00Z"
    }
  ],
  "total": 25,
  "page": 1,
  "page_size": 10,
  "total_pages": 3
}
```

#### PUT /api/v1/students/{id}
Updates an existing student record.

**Parameters:**
- `id` (path) - The student ID

**Request Body:** Same as POST /api/v1/students

**Response (200 OK):** Same structure as POST response

#### DELETE /api/v1/students/{id}
Deletes a student record.

**Parameters:**
- `id` (path) - The student ID

**Response (200 OK):**
```json
{
  "success": true,
  "message": "Student deleted successfully"
}
```

#### GET /api/v1/students/grade/{grade}
Retrieves all students in a specific grade.

**Parameters:**
- `grade` (path) - The grade (KG, 1-12)

**Response (200 OK):**
```json
{
  "success": true,
  "message": "Students retrieved successfully",
  "data": [
    {
      "id": "student_123",
      "first_name": "John",
      "last_name": "Doe",
      "email": "john.doe@example.com",
      "date_of_birth": "2010-05-15T00:00:00Z",
      "grade": "8",
      "address": "123 Main St, City, State",
      "phone": "1234567890",
      "parent_name": "Jane Doe",
      "parent_phone": "0987654321",
      "created_at": "2025-09-24T10:00:00Z",
      "updated_at": "2025-09-24T10:00:00Z"
    }
  ]
}
```

### Teachers API (Placeholder)

#### POST /api/v1/teachers
Creates a new teacher record.

**Status:** Not fully implemented

#### GET /api/v1/teachers/{id}
Retrieves a teacher by ID.

**Status:** Not fully implemented

### Classes API (Placeholder)

#### POST /api/v1/classes
Creates a new class record.

**Status:** Not fully implemented

#### GET /api/v1/classes/{id}
Retrieves a class by ID.

**Status:** Not fully implemented

### Academic Records API (Placeholder)

#### POST /api/v1/academics
Creates a new academic record.

**Status:** Not fully implemented

#### GET /api/v1/academics/{id}
Retrieves an academic record by ID.

**Status:** Not fully implemented

### Achievements API (Placeholder)

#### POST /api/v1/achievements
Creates a new achievement record.

**Status:** Not fully implemented

#### GET /api/v1/achievements/{id}
Retrieves an achievement by ID.

**Status:** Not fully implemented

## Data Models

### Student
Represents a student in the school system.

```go
type Student struct {
    ID          string    `json:"id" db:"id"`                         // Unique identifier
    FirstName   string    `json:"first_name" db:"first_name"`         // Student's first name
    LastName    string    `json:"last_name" db:"last_name"`           // Student's last name
    Email       string    `json:"email" db:"email"`                   // Student's email address (unique)
    DateOfBirth time.Time `json:"date_of_birth" db:"date_of_birth"`   // Student's date of birth
    Grade       string    `json:"grade" db:"grade"`                   // Current grade (KG, 1-12)
    Address     string    `json:"address" db:"address"`               // Home address
    Phone       string    `json:"phone" db:"phone"`                   // Student's phone number
    ParentName  string    `json:"parent_name" db:"parent_name"`       // Parent/guardian name
    ParentPhone string    `json:"parent_phone" db:"parent_phone"`     // Parent/guardian phone
    CreatedAt   time.Time `json:"created_at" db:"created_at"`         // Record creation timestamp
    UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`         // Record update timestamp
}
```

### Teacher
Represents a teacher in the school system.

```go
type Teacher struct {
    ID          string    `json:"id" db:"id"`                         // Unique identifier
    FirstName   string    `json:"first_name" db:"first_name"`         // Teacher's first name
    LastName    string    `json:"last_name" db:"last_name"`           // Teacher's last name
    Email       string    `json:"email" db:"email"`                   // Teacher's email address
    Phone       string    `json:"phone" db:"phone"`                   // Teacher's phone number
    Department  string    `json:"department" db:"department"`         // Department/subject area
    Subject     string    `json:"subject" db:"subject"`               // Primary subject taught
    Experience  int       `json:"experience" db:"experience"`         // Years of experience
    Salary      float64   `json:"salary" db:"salary"`                 // Monthly salary
    HireDate    time.Time `json:"hire_date" db:"hire_date"`           // Date of hiring
    CreatedAt   time.Time `json:"created_at" db:"created_at"`         // Record creation timestamp
    UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`         // Record update timestamp
}
```

### Class
Represents a class in the school system.

```go
type Class struct {
    ID        string    `json:"id" db:"id"`                         // Unique identifier
    Name      string    `json:"name" db:"name"`                     // Class name
    Subject   string    `json:"subject" db:"subject"`               // Subject taught
    TeacherID string    `json:"teacher_id" db:"teacher_id"`         // Assigned teacher ID
    Grade     string    `json:"grade" db:"grade"`                   // Grade level
    Room      string    `json:"room" db:"room"`                     // Classroom number/name
    Schedule  string    `json:"schedule" db:"schedule"`             // Class schedule
    Capacity  int       `json:"capacity" db:"capacity"`             // Maximum students
    Enrolled  int       `json:"enrolled" db:"enrolled"`             // Currently enrolled
    CreatedAt time.Time `json:"created_at" db:"created_at"`         // Record creation timestamp
    UpdatedAt time.Time `json:"updated_at" db:"updated_at"`         // Record update timestamp
}
```

### Academic
Represents academic records and exam results.

```go
type Academic struct {
    ID        string    `json:"id" db:"id"`                         // Unique identifier
    StudentID string    `json:"student_id" db:"student_id"`         // Student reference
    ClassID   string    `json:"class_id" db:"class_id"`             // Class reference
    ExamType  string    `json:"exam_type" db:"exam_type"`           // Type: midterm, final, quiz, assignment
    Subject   string    `json:"subject" db:"subject"`               // Subject name
    MaxMarks  float64   `json:"max_marks" db:"max_marks"`           // Maximum possible marks
    ObtMarks  float64   `json:"obt_marks" db:"obt_marks"`           // Marks obtained
    Grade     string    `json:"grade" db:"grade"`                   // Letter grade (A, B, C, D, F)
    ExamDate  time.Time `json:"exam_date" db:"exam_date"`           // Date of examination
    Remarks   string    `json:"remarks" db:"remarks"`               // Additional comments
    CreatedAt time.Time `json:"created_at" db:"created_at"`         // Record creation timestamp
    UpdatedAt time.Time `json:"updated_at" db:"updated_at"`         // Record update timestamp
}
```

### Achievement
Represents student achievements and awards.

```go
type Achievement struct {
    ID          string    `json:"id" db:"id"`                         // Unique identifier
    StudentID   string    `json:"student_id" db:"student_id"`         // Student reference
    Title       string    `json:"title" db:"title"`                   // Achievement title
    Description string    `json:"description" db:"description"`       // Detailed description
    Category    string    `json:"category" db:"category"`             // Category: academic, sports, cultural, behavior
    Level       string    `json:"level" db:"level"`                   // Level: school, district, state, national
    Date        time.Time `json:"date" db:"date"`                     // Achievement date
    AwardedBy   string    `json:"awarded_by" db:"awarded_by"`         // Awarding authority
    CreatedAt   time.Time `json:"created_at" db:"created_at"`         // Record creation timestamp
    UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`         // Record update timestamp
}
```

## Request/Response Formats

### Standard API Response
All API responses follow a consistent format:

```go
type APIResponse struct {
    Success bool        `json:"success"`         // Indicates if request was successful
    Message string      `json:"message"`         // Human-readable message
    Data    interface{} `json:"data,omitempty"`  // Response data (omitted on error)
    Error   string      `json:"error,omitempty"` // Error message (omitted on success)
}
```

### Paginated Response
For endpoints that return multiple items:

```go
type PaginatedResponse struct {
    Success    bool        `json:"success"`         // Indicates if request was successful
    Message    string      `json:"message"`         // Human-readable message
    Data       interface{} `json:"data,omitempty"`  // Array of items
    Total      int         `json:"total"`           // Total number of items
    Page       int         `json:"page"`            // Current page number
    PageSize   int         `json:"page_size"`       // Items per page
    TotalPages int         `json:"total_pages"`     // Total number of pages
    Error      string      `json:"error,omitempty"` // Error message (omitted on success)
}
```

### Create Student Request
Request format for creating/updating students:

```go
type CreateStudentRequest struct {
    FirstName   string    `json:"first_name" binding:"required"`     // Required: Student's first name
    LastName    string    `json:"last_name" binding:"required"`      // Required: Student's last name
    Email       string    `json:"email" binding:"required,email"`    // Required: Valid email address
    DateOfBirth time.Time `json:"date_of_birth" binding:"required"`  // Required: Date of birth
    Grade       string    `json:"grade" binding:"required"`          // Required: Grade (KG, 1-12)
    Address     string    `json:"address"`                           // Optional: Home address
    Phone       string    `json:"phone"`                             // Optional: Phone number
    ParentName  string    `json:"parent_name"`                       // Optional: Parent/guardian name
    ParentPhone string    `json:"parent_phone"`                      // Optional: Parent/guardian phone
}
```

## Error Handling

The API uses standard HTTP status codes and returns detailed error information:

### HTTP Status Codes
- `200 OK` - Successful GET, PUT, DELETE operations
- `201 Created` - Successful POST operations
- `400 Bad Request` - Invalid request data or parameters
- `404 Not Found` - Requested resource not found
- `500 Internal Server Error` - Server-side error

### Error Response Format
```json
{
  "success": false,
  "message": "Brief error description",
  "error": "Detailed error message"
}
```

### Common Validation Errors
- Missing required fields
- Invalid email format
- Invalid grade (must be KG or 1-12)
- Invalid age (must be between 3 and 25 years)
- Duplicate email addresses
- Phone numbers must be at least 10 digits

## Authentication

**Current Status:** No authentication implemented  
**Future Enhancement:** JWT-based authentication planned

## Package Documentation

### package models
The models package defines all data structures used throughout the application, including entity models, request/response DTOs, and API response formats.

Key types:
- `Student`, `Teacher`, `Class`, `Academic`, `Achievement` - Core domain entities
- `APIResponse`, `PaginatedResponse` - Standard response formats  
- `CreateStudentRequest`, `CreateTeacherRequest`, etc. - Request DTOs

### package handler
The handler package implements HTTP request handlers using the Gin framework. It processes HTTP requests, validates input, calls service layer methods, and returns appropriate HTTP responses.

Key types:
- `StudentHandler` - Handles student-related HTTP endpoints
- `Handler` - Aggregates all entity handlers

Key methods:
- `CreateStudent(c *gin.Context)` - Handles POST /api/v1/students
- `GetStudent(c *gin.Context)` - Handles GET /api/v1/students/:id  
- `GetAllStudents(c *gin.Context)` - Handles GET /api/v1/students with pagination
- `UpdateStudent(c *gin.Context)` - Handles PUT /api/v1/students/:id
- `DeleteStudent(c *gin.Context)` - Handles DELETE /api/v1/students/:id
- `GetStudentsByGrade(c *gin.Context)` - Handles GET /api/v1/students/grade/:grade

### package service
The service package implements business logic and validation rules. It acts as an intermediary between HTTP handlers and the repository layer.

Key interfaces:
- `StudentService` - Defines student business logic operations

Key methods:
- `CreateStudent(ctx context.Context, req *CreateStudentRequest) (*Student, error)` - Creates student with validation
- `GetStudent(ctx context.Context, id string) (*Student, error)` - Retrieves student by ID
- `GetAllStudents(ctx context.Context, page, pageSize int) ([]*Student, int, error)` - Gets paginated student list
- `UpdateStudent(ctx context.Context, id string, req *CreateStudentRequest) (*Student, error)` - Updates student with validation
- `DeleteStudent(ctx context.Context, id string) error` - Deletes student
- `GetStudentsByGrade(ctx context.Context, grade string) ([]*Student, error)` - Gets students by grade
- `GetStudentByEmail(ctx context.Context, email string) (*Student, error)` - Finds student by email

Business rules enforced:
- Email uniqueness validation
- Age validation (3-25 years)
- Grade validation (KG, 1-12)
- Phone number validation (minimum 10 digits)
- Required field validation

### package repository
The repository package provides data access abstraction with Couchbase database integration.

Key interfaces:
- `StudentRepository` - Defines data access operations for students

Key methods:
- `Create(ctx context.Context, student *Student) error` - Persists new student
- `GetByID(ctx context.Context, id string) (*Student, error)` - Retrieves by ID
- `GetAll(ctx context.Context, page, pageSize int) ([]*Student, int, error)` - Paginated retrieval
- `Update(ctx context.Context, id string, student *Student) error` - Updates existing student
- `Delete(ctx context.Context, id string) error` - Removes student
- `GetByGrade(ctx context.Context, grade string) ([]*Student, error)` - Query by grade
- `GetByEmail(ctx context.Context, email string) (*Student, error)` - Query by email

### package config
The config package manages application configuration, including environment variables and database connection settings.

Key functions:
- Configuration loading from environment variables
- Database connection string management
- Server configuration (port, timeouts, etc.)

---

*Generated: September 24, 2025*  
*Version: 1.0.0*  
*Last Updated: September 24, 2025*