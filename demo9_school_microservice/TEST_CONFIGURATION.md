# [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-08-21
# Test Configuration for School Microservice Student Service

# This file contains configuration and documentation for running tests
# and generating coverage reports for the Student Service

## Test Structure

The Student Service test suite is organized into three main packages:

### 1. Handlers Package (`./handlers/`)
- **Purpose**: Tests HTTP request handlers and API endpoints
- **Test Files**: `student_handler_test.go`
- **Coverage**: HTTP request/response handling, JSON serialization, error handling
- **Key Test Types**:
  - Unit tests for each HTTP endpoint (GET, POST, PUT, DELETE)
  - Integration tests for complete CRUD workflows
  - Error case testing (invalid JSON, missing resources, database errors)
  - Context cancellation and timeout testing
  - Benchmark tests for performance validation

### 2. Repository Package (`./repository/`)
- **Purpose**: Tests data access layer and database operations
- **Test Files**: `student_repository_test.go`
- **Coverage**: Database CRUD operations, error handling, context support
- **Key Test Types**:
  - Mock database testing with various scenarios
  - Context-aware operations testing
  - Error condition testing (connection failures, document not found)
  - Concurrent operation safety testing
  - Performance benchmarks for database operations

### 3. Models Package (`./models/`)
- **Purpose**: Tests data structures and validation logic
- **Test Files**: `student_test.go`
- **Coverage**: JSON serialization, data validation, field constraints
- **Key Test Types**:
  - JSON marshaling/unmarshaling tests
  - Data validation testing
  - Field constraint verification
  - Edge case handling (nil values, invalid formats)
  - Format validation (email, ID, status values)

## Test Execution

### Running All Tests
```bash
# Linux/Mac
./scripts/run-tests.sh

# Windows
scripts\run-tests.bat
```

### Running Specific Package Tests
```bash
# Handlers only
./scripts/run-tests.sh handlers

# Repository only  
./scripts/run-tests.sh repository

# Models only
./scripts/run-tests.sh models
```

### Running Benchmarks
```bash
./scripts/run-tests.sh benchmarks
```

### Viewing Test Statistics
```bash
./scripts/run-tests.sh stats
```

### Cleaning Test Artifacts
```bash
./scripts/run-tests.sh clean
```

## Coverage Requirements

### Coverage Thresholds
- **Total Coverage**: 80% minimum
- **Package Coverage**: 75% minimum per package
- **Critical Paths**: 90% minimum for core business logic

### Coverage Reporting
- **Text Report**: Displayed in console during test execution
- **HTML Report**: Generated as `coverage.html` for detailed analysis
- **Package Reports**: Individual coverage files per package

### Coverage Analysis
The coverage report includes:
- **Line Coverage**: Percentage of code lines executed during tests
- **Function Coverage**: Percentage of functions called during tests
- **Branch Coverage**: Percentage of decision branches tested

## Test Configuration Options

### Environment Variables
```bash
# Test execution timeout
TEST_TIMEOUT=30s

# Enable verbose output
TEST_VERBOSE=true

# Enable race condition detection
TEST_RACE_DETECTION=true

# Enable coverage collection
TEST_COVERAGE=true

# Coverage output files
COVERAGE_OUTPUT_FILE=coverage.out
COVERAGE_HTML_FILE=coverage.html
```

### Go Test Flags
The test scripts use the following Go test flags:
- `-v`: Verbose output showing individual test results
- `-race`: Race condition detection for concurrent code
- `-timeout=30s`: Test execution timeout
- `-coverprofile`: Coverage data output file
- `-covermode=atomic`: Atomic coverage mode for accurate concurrent testing
- `-bench=.`: Run all benchmark tests
- `-benchmem`: Include memory allocation statistics in benchmarks

## Test Data and Mocking

### Mock Implementations
- **MockStudentRepository**: Simulates database operations without real database
- **MockCouchbaseClient**: Simulates Couchbase database client
- **Test Data**: Predefined student records for consistent testing

### Test Scenarios
Each test package includes comprehensive scenarios:
- **Happy Path**: Normal operation with valid data
- **Error Cases**: Invalid input, missing resources, database failures
- **Edge Cases**: Boundary conditions, nil values, empty data
- **Concurrency**: Multiple simultaneous operations
- **Performance**: Load testing and benchmark validation

## Integration with CI/CD

### Automated Testing
The test suite is designed for integration with CI/CD pipelines:
- **Exit Codes**: Proper exit codes for build pipeline integration
- **Coverage Reports**: Machine-readable coverage data
- **Performance Metrics**: Benchmark results for performance regression detection
- **Test Artifacts**: Structured output for test result analysis

### Quality Gates
Recommended quality gates for CI/CD:
- All tests must pass
- Coverage must meet minimum thresholds
- No race conditions detected
- Benchmark performance within acceptable ranges
- No security vulnerabilities in dependencies

## Test Maintenance

### Adding New Tests
When adding new functionality:
1. **Unit Tests**: Test individual functions and methods
2. **Integration Tests**: Test component interactions
3. **Error Cases**: Test failure scenarios
4. **Performance Tests**: Add benchmarks for critical paths
5. **Documentation**: Update test documentation

### Test Review Checklist
- [ ] All public functions have corresponding tests
- [ ] Error cases are covered
- [ ] Edge cases are tested
- [ ] Mock objects are properly configured
- [ ] Test data is realistic and comprehensive
- [ ] Performance tests include realistic load scenarios
- [ ] Tests are deterministic and repeatable

## Troubleshooting

### Common Issues
1. **Test Timeouts**: Increase TEST_TIMEOUT for slow operations
2. **Race Conditions**: Review concurrent code and add proper synchronization
3. **Coverage Gaps**: Identify untested code paths and add targeted tests
4. **Mock Failures**: Verify mock configurations match expected behavior
5. **Performance Regression**: Compare benchmark results with baseline metrics

### Debug Options
```bash
# Run tests with debug output
go test -v -x ./...

# Run specific test with debug
go test -v -run TestSpecificFunction ./...

# Generate CPU profile during tests
go test -cpuprofile=cpu.prof ./...

# Generate memory profile during tests
go test -memprofile=mem.prof ./...
```

## Best Practices

### Test Design
- **Single Responsibility**: Each test should verify one specific behavior
- **Descriptive Names**: Test names should clearly describe what is being tested
- **Arrange-Act-Assert**: Structure tests with clear setup, execution, and verification
- **Independent Tests**: Tests should not depend on execution order
- **Realistic Data**: Use realistic test data that represents actual usage

### Performance Testing
- **Baseline Metrics**: Establish performance baselines for comparison
- **Load Testing**: Test with realistic data volumes
- **Memory Efficiency**: Monitor memory usage in benchmarks
- **Concurrent Testing**: Verify performance under concurrent load
- **Resource Cleanup**: Ensure proper resource cleanup in tests

### Maintenance
- **Regular Updates**: Keep test data and scenarios current
- **Refactoring**: Refactor tests when code changes
- **Documentation**: Maintain test documentation
- **Review**: Regular test review and improvement
- **Automation**: Automate test execution and reporting
