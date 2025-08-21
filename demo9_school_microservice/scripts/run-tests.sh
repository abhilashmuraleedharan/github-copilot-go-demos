# [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-08-21
# Test Configuration and Coverage Script

# This script provides comprehensive test execution and coverage reporting
# for the School Microservice Student Service

# Test Execution Configuration
TEST_TIMEOUT=30s
TEST_VERBOSE=true
TEST_RACE_DETECTION=true
TEST_COVERAGE=true
COVERAGE_OUTPUT_FILE=coverage.out
COVERAGE_HTML_FILE=coverage.html

# Coverage thresholds
COVERAGE_THRESHOLD_TOTAL=80
COVERAGE_THRESHOLD_PACKAGE=75

# Test directories
STUDENT_SERVICE_DIR=./services/students
TEST_DIRS=("./services/students/handlers" "./services/students/repository" "./services/students/models")

echo "🧪 School Microservice - Student Service Test Suite"
echo "=================================================="
echo ""

# Function to run tests with coverage for a specific package
run_package_tests() {
    local package_dir=$1
    local package_name=$(basename $package_dir)
    
    echo "📦 Testing package: $package_name"
    echo "Directory: $package_dir"
    echo ""
    
    # Run tests with coverage
    cd $package_dir
    
    # Execute tests with race detection and coverage
    go test -v -race -timeout=$TEST_TIMEOUT -coverprofile="${package_name}_coverage.out" -covermode=atomic ./...
    
    if [ $? -eq 0 ]; then
        echo "✅ $package_name tests PASSED"
        
        # Generate coverage report
        if [ -f "${package_name}_coverage.out" ]; then
            coverage_percent=$(go tool cover -func="${package_name}_coverage.out" | grep total | awk '{print $3}' | sed 's/%//')
            echo "📊 Coverage: ${coverage_percent}%"
            
            # Check coverage threshold
            if (( $(echo "$coverage_percent >= $COVERAGE_THRESHOLD_PACKAGE" | bc -l) )); then
                echo "✅ Coverage meets threshold (${COVERAGE_THRESHOLD_PACKAGE}%)"
            else
                echo "⚠️  Coverage below threshold (${COVERAGE_THRESHOLD_PACKAGE}%)"
            fi
            
            # Generate HTML coverage report
            go tool cover -html="${package_name}_coverage.out" -o="${package_name}_coverage.html"
            echo "📄 HTML coverage report: ${package_name}_coverage.html"
        fi
    else
        echo "❌ $package_name tests FAILED"
    fi
    
    echo ""
    cd - > /dev/null
}

# Function to run all tests and generate combined coverage
run_all_tests() {
    echo "🔄 Running all Student Service tests..."
    echo ""
    
    cd $STUDENT_SERVICE_DIR
    
    # Run all tests with combined coverage
    go test -v -race -timeout=$TEST_TIMEOUT -coverprofile="$COVERAGE_OUTPUT_FILE" -covermode=atomic ./...
    
    if [ $? -eq 0 ]; then
        echo "✅ All Student Service tests PASSED"
        
        # Generate combined coverage report
        if [ -f "$COVERAGE_OUTPUT_FILE" ]; then
            echo ""
            echo "📊 Combined Coverage Report:"
            go tool cover -func="$COVERAGE_OUTPUT_FILE"
            
            # Extract total coverage
            total_coverage=$(go tool cover -func="$COVERAGE_OUTPUT_FILE" | grep total | awk '{print $3}' | sed 's/%//')
            echo ""
            echo "📈 Total Coverage: ${total_coverage}%"
            
            # Check total coverage threshold
            if (( $(echo "$total_coverage >= $COVERAGE_THRESHOLD_TOTAL" | bc -l) )); then
                echo "✅ Total coverage meets threshold (${COVERAGE_THRESHOLD_TOTAL}%)"
            else
                echo "⚠️  Total coverage below threshold (${COVERAGE_THRESHOLD_TOTAL}%)"
            fi
            
            # Generate HTML coverage report
            go tool cover -html="$COVERAGE_OUTPUT_FILE" -o="$COVERAGE_HTML_FILE"
            echo "📄 HTML coverage report: $COVERAGE_HTML_FILE"
            echo ""
            echo "🌐 Open coverage report in browser:"
            echo "   file://$(pwd)/$COVERAGE_HTML_FILE"
        fi
    else
        echo "❌ Some Student Service tests FAILED"
        exit 1
    fi
    
    cd - > /dev/null
}

# Function to run benchmarks
run_benchmarks() {
    echo "⚡ Running benchmarks..."
    echo ""
    
    cd $STUDENT_SERVICE_DIR
    
    # Run benchmarks for all packages
    go test -bench=. -benchmem ./...
    
    cd - > /dev/null
}

# Function to clean up test artifacts
cleanup() {
    echo "🧹 Cleaning up test artifacts..."
    
    cd $STUDENT_SERVICE_DIR
    
    # Remove coverage files
    find . -name "*.out" -type f -delete
    find . -name "*coverage.html" -type f -delete
    
    cd - > /dev/null
    
    echo "✅ Cleanup complete"
}

# Function to check test dependencies
check_dependencies() {
    echo "🔍 Checking test dependencies..."
    
    # Check if Go is installed
    if ! command -v go &> /dev/null; then
        echo "❌ Go is not installed or not in PATH"
        exit 1
    fi
    
    # Check Go version (require 1.19+)
    go_version=$(go version | awk '{print $3}' | sed 's/go//')
    echo "📋 Go version: $go_version"
    
    # Check if bc is installed for float comparison
    if ! command -v bc &> /dev/null; then
        echo "⚠️  bc not found - coverage threshold checking disabled"
        COVERAGE_THRESHOLD_TOTAL=0
        COVERAGE_THRESHOLD_PACKAGE=0
    fi
    
    echo "✅ Dependencies check complete"
    echo ""
}

# Function to display test statistics
show_test_stats() {
    echo ""
    echo "📈 Test Statistics Summary"
    echo "=========================="
    
    cd $STUDENT_SERVICE_DIR
    
    # Count test files
    test_files=$(find . -name "*_test.go" | wc -l)
    echo "📁 Test files: $test_files"
    
    # Count test functions
    test_functions=$(grep -r "^func Test" . --include="*_test.go" | wc -l)
    echo "🧪 Test functions: $test_functions"
    
    # Count benchmark functions
    benchmark_functions=$(grep -r "^func Benchmark" . --include="*_test.go" | wc -l)
    echo "⚡ Benchmark functions: $benchmark_functions"
    
    # Count example functions
    example_functions=$(grep -r "^func Example" . --include="*_test.go" | wc -l)
    echo "📖 Example functions: $example_functions"
    
    echo ""
    
    cd - > /dev/null
}

# Main execution based on command line arguments
case "${1:-all}" in
    "handlers")
        check_dependencies
        run_package_tests "./services/students/handlers"
        ;;
    "repository")
        check_dependencies
        run_package_tests "./services/students/repository"
        ;;
    "models")
        check_dependencies
        run_package_tests "./services/students/models"
        ;;
    "benchmarks")
        check_dependencies
        run_benchmarks
        ;;
    "stats")
        show_test_stats
        ;;
    "clean")
        cleanup
        ;;
    "all"|*)
        check_dependencies
        run_all_tests
        run_benchmarks
        show_test_stats
        echo ""
        echo "🎉 Test suite execution complete!"
        echo ""
        echo "📋 Next steps:"
        echo "   • Review coverage report: $COVERAGE_HTML_FILE"
        echo "   • Check for any failing tests"
        echo "   • Optimize slow benchmarks if needed"
        echo "   • Add more tests for uncovered code paths"
        ;;
esac
