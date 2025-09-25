# School Management Microservice - Testing Guide

<!-- [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-09-24 -->

This document provides comprehensive instructions for running unit tests and generating test coverage reports for the School Management Microservice.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Running Tests](#running-tests)
- [Test Coverage](#test-coverage)
- [Test Structure](#test-structure)
- [Writing Tests](#writing-tests)
- [CI/CD Integration](#cicd-integration)
- [Troubleshooting](#troubleshooting)

## Prerequisites

Before running tests, ensure you have:

1. **Go 1.20+** installed
2. **Project dependencies** downloaded
3. **Testing framework** (testify) available

### Install Dependencies

```bash
# Download all dependencies including test dependencies
go mod download
go mod tidy
```

## Running Tests

### Basic Test Execution

#### Option 1: Using Go Command

```bash
# Run all tests
go test ./...

# Run tests with verbose output
go test -v ./...

# Run tests for specific package
go test ./internal/service

# Run specific test function
go test -v ./internal/service -run TestStudentService_CreateStudent_Success
```

#### Option 2: Using Makefile

```bash
# Run all tests (quiet)
make test

# Run all tests (verbose)
make test-verbose

# Run tests with coverage
make coverage

# Generate HTML coverage report
make coverage-html
```

#### Option 3: Using PowerShell Script (Windows)

```powershell
# Run the PowerShell coverage script
.\test-coverage.ps1
```

#### Option 4: Using Shell Script (Linux/macOS)

```bash
# Make script executable (first time only)
chmod +x test-coverage.sh

# Run the shell coverage script
./test-coverage.sh
```

### Test Output Examples

**Successful Test Run:**
```
=== RUN   TestStudentService_CreateStudent_Success
--- PASS: TestStudentService_CreateStudent_Success (0.00s)
=== RUN   TestStudentService_CreateStudent_ValidationError_MissingFirstName
--- PASS: TestStudentService_CreateStudent_ValidationError_MissingFirstName (0.00s)
=== RUN   TestStudentService_GetStudent_Success
--- PASS: TestStudentService_GetStudent_Success (0.00s)
PASS
ok      school-microservice/internal/service    0.123s
```

**Failed Test Run:**
```
=== RUN   TestStudentService_CreateStudent_Success
--- FAIL: TestStudentService_CreateStudent_Success (0.00s)
    student_service_test.go:89: Expected no error, but got: validation failed
FAIL
exit status 1
FAIL    school-microservice/internal/service    0.045s
```

## Test Coverage

### Generating Coverage Reports

#### Method 1: Using Go Tools

```bash
# Generate coverage profile
go test -coverprofile=coverage.out ./...

# View coverage summary in terminal
go tool cover -func=coverage.out

# Generate HTML coverage report
go tool cover -html=coverage.out -o coverage.html

# Open HTML report in browser (Windows)
start coverage.html

# Open HTML report in browser (macOS)
open coverage.html

# Open HTML report in browser (Linux)
xdg-open coverage.html
```

#### Method 2: Using Makefile

```bash
# Generate coverage with summary
make coverage

# Generate HTML coverage report
make coverage-html
```

#### Method 3: Using Scripts

**PowerShell (Windows):**
```powershell
.\test-coverage.ps1
```

**Shell (Linux/macOS):**
```bash
./test-coverage.sh
```

### Coverage Report Output

**Terminal Coverage Summary:**
```
school-microservice/internal/service/service.go:28:        CreateStudent           85.7%
school-microservice/internal/service/service.go:45:        GetStudent             100.0%
school-microservice/internal/service/service.go:53:        GetAllStudents          90.0%
school-microservice/internal/service/service.go:67:        UpdateStudent           88.2%
school-microservice/internal/service/service.go:98:        DeleteStudent           92.3%
school-microservice/internal/service/service.go:113:       GetStudentsByGrade     100.0%
school-microservice/internal/service/service.go:121:       GetStudentByEmail      100.0%
school-microservice/internal/service/service.go:129:       validateStudentRequest  95.5%
total:                                                      (statements)            91.2%
```

**HTML Coverage Report Features:**
- **Line-by-line coverage** with color coding
- **Green:** Covered lines
- **Red:** Uncovered lines
- **Gray:** Non-executable lines
- **Interactive navigation** through packages and files
- **Coverage percentage** for each function and file

### Coverage Files Generated

After running coverage tests, you'll find:

- **`coverage.out`** - Raw coverage data (binary format)
- **`coverage.html`** - Interactive HTML report
- **`coverage-summary.txt`** - Text summary of coverage percentages

## Test Structure

### Current Test Coverage

The project includes comprehensive unit tests for:

#### StudentService Tests
- ✅ **CreateStudent** - 12 test cases
  - Success scenarios
  - Validation errors (missing fields, invalid formats)
  - Business rule violations (duplicate email, invalid age)
  - Repository errors
- ✅ **GetStudent** - 3 test cases
  - Success retrieval
  - Empty ID validation
  - Not found scenarios
- ✅ **GetAllStudents** - 3 test cases
  - Successful pagination
  - Default pagination handling
  - Page size limiting
- ✅ **UpdateStudent** - 4 test cases
  - Success updates
  - Validation errors
  - Student not found
  - Email conflict handling
- ✅ **DeleteStudent** - 3 test cases
  - Success deletion
  - Empty ID validation
  - Student not found
- ✅ **GetStudentsByGrade** - 2 test cases
  - Success retrieval
  - Empty grade validation
- ✅ **GetStudentByEmail** - 2 test cases
  - Success retrieval
  - Empty email validation
- ✅ **Validation Tests** - 3 additional test cases
  - Complete validation error scenarios
  - Valid grade testing
  - Optional field handling

**Total Test Cases: 32**

### Test File Organization

```
internal/
└── service/
    ├── service.go              # Implementation
    └── student_service_test.go # Unit tests
```

## Writing Tests

### Test Naming Conventions

Follow the pattern: `TestServiceName_MethodName_Scenario`

```go
func TestStudentService_CreateStudent_Success(t *testing.T)
func TestStudentService_CreateStudent_ValidationError_MissingFirstName(t *testing.T)
func TestStudentService_GetStudent_NotFound(t *testing.T)
```

### Test Structure Pattern

```go
func TestStudentService_MethodName_Scenario(t *testing.T) {
    // Arrange - Set up test data and mocks
    service, mockRepo := setupStudentService()
    ctx := context.Background()
    // ... setup test data
    
    // Setup mock expectations
    mockRepo.On("Method", ctx, params).Return(result, error)
    
    // Act - Execute the method under test
    result, err := service.Method(ctx, params)
    
    // Assert - Verify the results
    assert.NoError(t, err)
    assert.Equal(t, expected, result)
    mockRepo.AssertExpectations(t)
}
```

### Mock Setup

The tests use `testify/mock` for mocking repository dependencies:

```go
// MockStudentRepository implements the repository interface
type MockStudentRepository struct {
    mock.Mock
}

// Setup service with mocked dependencies
func setupStudentService() (*studentService, *MockStudentRepository) {
    mockStudentRepo := &MockStudentRepository{}
    mockRepo := &repository.Repository{
        Student: mockStudentRepo,
    }
    service := &studentService{repo: mockRepo}
    return service, mockStudentRepo
}
```

### Test Data Helpers

Use helper functions for consistent test data:

```go
// Helper to create valid student request
func createValidStudentRequest() *models.CreateStudentRequest {
    dateOfBirth := time.Date(2010, time.May, 15, 0, 0, 0, 0, time.UTC)
    return &models.CreateStudentRequest{
        FirstName:   "John",
        LastName:    "Doe",
        Email:       "john.doe@example.com",
        DateOfBirth: dateOfBirth,
        Grade:       "8",
        // ... other fields
    }
}
```

## CI/CD Integration

### GitHub Actions Example

```yaml
name: Test and Coverage

on:
  push:
    branches: [ main, develop ]
  pull_request:
    branches: [ main ]

jobs:
  test:
    runs-on: ubuntu-latest
    
    steps:
    - uses: actions/checkout@v3
    
    - name: Set up Go
      uses: actions/setup-go@v3
      with:
        go-version: '1.20'
        
    - name: Download dependencies
      run: go mod download
      
    - name: Run tests with coverage
      run: go test -v -coverprofile=coverage.out ./...
      
    - name: Generate coverage report
      run: go tool cover -html=coverage.out -o coverage.html
      
    - name: Upload coverage reports
      uses: actions/upload-artifact@v3
      with:
        name: coverage-report
        path: |
          coverage.out
          coverage.html
```

### Coverage Thresholds

Set minimum coverage requirements:

```bash
# Extract coverage percentage
COVERAGE=$(go tool cover -func=coverage.out | grep total | awk '{print $3}' | sed 's/%//')

# Check if coverage meets threshold (e.g., 80%)
if (( $(echo "$COVERAGE < 80" | bc -l) )); then
    echo "Coverage $COVERAGE% is below threshold of 80%"
    exit 1
fi
```

## Troubleshooting

### Common Issues

#### 1. Import Path Issues

**Error:**
```
package school-microservice/internal/service: cannot find package
```

**Solution:**
- Ensure you're running tests from the project root directory
- Verify `go.mod` module name matches import paths
- Run `go mod tidy` to clean up dependencies

#### 2. Mock Expectations Not Met

**Error:**
```
mock: Unexpected Method Call
```

**Solution:**
- Verify all mock expectations are properly set up
- Check method signatures match exactly
- Ensure `AssertExpectations(t)` is called

#### 3. Coverage Files Not Generated

**Error:**
```
coverage.out: no such file or directory
```

**Solution:**
- Ensure tests pass before coverage generation
- Check write permissions in the directory
- Run `go test -coverprofile=coverage.out ./...` explicitly

#### 4. Testify Import Issues

**Error:**
```
cannot find package "github.com/stretchr/testify"
```

**Solution:**
```bash
# Add testify dependency
go get github.com/stretchr/testify@latest
go mod tidy
```

### Debugging Tests

#### Running Individual Tests

```bash
# Run specific test
go test -v ./internal/service -run TestStudentService_CreateStudent_Success

# Run tests matching pattern
go test -v ./internal/service -run "TestStudentService_.*Success"
```

#### Verbose Output

```bash
# Maximum verbosity
go test -v -count=1 ./...
```

#### Race Condition Detection

```bash
# Run with race detector
go test -race ./...
```

### Performance Testing

```bash
# Run benchmarks (if any)
go test -bench=. ./...

# Memory allocation profiling
go test -bench=. -benchmem ./...
```

## Best Practices

### Test Organization

1. **One test file per source file** - `service.go` → `service_test.go`
2. **Group related tests** - Use test suites for complex scenarios
3. **Clear test names** - Describe the scenario being tested
4. **Independent tests** - Each test should be able to run in isolation

### Test Data Management

1. **Use helper functions** - Create reusable test data generators
2. **Avoid hardcoded values** - Use constants or configuration
3. **Test edge cases** - Include boundary conditions and error scenarios
4. **Mock external dependencies** - Keep tests focused on the unit under test

### Coverage Goals

1. **Aim for >90% coverage** - Balance between coverage and maintainability
2. **Focus on critical paths** - Prioritize business logic coverage
3. **Test error conditions** - Don't just test the happy path
4. **Document uncovered lines** - Justify why certain lines aren't covered

---

**Testing Guide Information:**
- **Version:** 1.0.0
- **Last Updated:** September 24, 2025
- **Authors:** GitHub Copilot
- **Next Review:** October 24, 2025