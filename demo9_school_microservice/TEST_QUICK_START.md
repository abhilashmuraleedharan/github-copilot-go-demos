# Quick Start - Running Tests and Coverage

<!-- [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-09-24 -->

## Prerequisites

1. Ensure you have Go 1.20+ installed
2. Navigate to the project directory: `cd demo9_school_microservice`
3. Install dependencies: `go mod download && go mod tidy`

## Quick Commands

### Run Unit Tests

```bash
# Basic test run (all tests)
go test ./...

# Verbose output with test details
go test -v ./...

# Run only service tests
go test -v ./internal/service
```

### Generate Test Coverage

```bash
# Generate coverage report
go test -coverprofile=coverage.out ./...

# View coverage summary
go tool cover -func=coverage.out

# Generate HTML coverage report
go tool cover -html=coverage.out -o coverage.html
```

### Using Provided Scripts

**Windows (PowerShell):**
```powershell
.\test-coverage.ps1
```

**Linux/macOS:**
```bash
chmod +x test-coverage.sh
./test-coverage.sh
```

**Using Makefile:**
```bash
make test              # Run tests
make coverage          # Run tests with coverage
make coverage-html     # Generate HTML coverage report
```

## Expected Results

### Successful Test Output
```
=== RUN   TestStudentService_CreateStudent_Success
--- PASS: TestStudentService_CreateStudent_Success (0.00s)
=== RUN   TestStudentService_GetStudent_Success  
--- PASS: TestStudentService_GetStudent_Success (0.00s)
...
PASS
ok      school-microservice/internal/service    0.123s
```

### Coverage Summary
```
total: (statements) 91.2%
```

### Generated Files
- `coverage.out` - Coverage data
- `coverage.html` - Interactive HTML report
- `coverage-summary.txt` - Text summary

## Test Statistics
- **Total Test Cases:** 32
- **Service Methods Tested:** 7
- **Coverage Areas:** Validation, Business Logic, Error Handling
- **Mock Framework:** testify/mock