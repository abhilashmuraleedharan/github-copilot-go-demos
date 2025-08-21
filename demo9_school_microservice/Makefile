# [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-08-21
# Makefile for School Microservice Student Service Testing

# Variables
STUDENT_SERVICE_DIR = ./services/students
COVERAGE_FILE = coverage.out
COVERAGE_HTML = coverage.html
TEST_TIMEOUT = 30s
GO_TEST_FLAGS = -v -race -timeout=$(TEST_TIMEOUT)
COVERAGE_FLAGS = -coverprofile=$(COVERAGE_FILE) -covermode=atomic
BENCH_FLAGS = -bench=. -benchmem

# Colors for output
GREEN = \033[0;32m
YELLOW = \033[1;33m
RED = \033[0;31m
NC = \033[0m # No Color

.PHONY: help test test-handlers test-repository test-models test-coverage test-benchmarks clean deps-check

# Default target
all: deps-check test test-benchmarks coverage-html

# Display help information
help:
	@echo "$(GREEN)School Microservice Student Service - Test Commands$(NC)"
	@echo "======================================================"
	@echo ""
	@echo "$(YELLOW)Available targets:$(NC)"
	@echo "  help              - Show this help message"
	@echo "  test              - Run all tests with coverage"
	@echo "  test-handlers     - Run handler tests only"
	@echo "  test-repository   - Run repository tests only"
	@echo "  test-models       - Run model tests only"
	@echo "  test-coverage     - Run tests and generate coverage report"
	@echo "  test-benchmarks   - Run benchmark tests"
	@echo "  coverage-html     - Generate HTML coverage report"
	@echo "  coverage-func     - Show function coverage summary"
	@echo "  deps-check        - Check test dependencies"
	@echo "  clean             - Clean test artifacts"
	@echo "  lint              - Run code linting"
	@echo "  vet               - Run go vet analysis"
	@echo "  fmt               - Format Go code"
	@echo ""
	@echo "$(YELLOW)Examples:$(NC)"
	@echo "  make test                    # Run all tests"
	@echo "  make test-handlers           # Test handlers only"
	@echo "  make test-coverage           # Tests with coverage"
	@echo "  make test-benchmarks         # Run benchmarks"
	@echo "  make clean                   # Clean artifacts"

# Check dependencies
deps-check:
	@echo "$(YELLOW)🔍 Checking dependencies...$(NC)"
	@command -v go >/dev/null 2>&1 || { echo "$(RED)❌ Go is not installed$(NC)"; exit 1; }
	@echo "$(GREEN)📋 Go version: $$(go version | awk '{print $$3}')$(NC)"
	@echo "$(GREEN)✅ Dependencies check passed$(NC)"
	@echo ""

# Run all tests
test: deps-check
	@echo "$(YELLOW)🧪 Running all Student Service tests...$(NC)"
	@cd $(STUDENT_SERVICE_DIR) && \
		go test $(GO_TEST_FLAGS) $(COVERAGE_FLAGS) ./... && \
		echo "$(GREEN)✅ All tests passed$(NC)" || \
		{ echo "$(RED)❌ Some tests failed$(NC)"; exit 1; }

# Run handler tests only
test-handlers: deps-check
	@echo "$(YELLOW)📦 Testing handlers package...$(NC)"
	@cd $(STUDENT_SERVICE_DIR)/handlers && \
		go test $(GO_TEST_FLAGS) -coverprofile=handlers_coverage.out -covermode=atomic . && \
		echo "$(GREEN)✅ Handler tests passed$(NC)" || \
		{ echo "$(RED)❌ Handler tests failed$(NC)"; exit 1; }

# Run repository tests only
test-repository: deps-check
	@echo "$(YELLOW)📦 Testing repository package...$(NC)"
	@cd $(STUDENT_SERVICE_DIR)/repository && \
		go test $(GO_TEST_FLAGS) -coverprofile=repository_coverage.out -covermode=atomic . && \
		echo "$(GREEN)✅ Repository tests passed$(NC)" || \
		{ echo "$(RED)❌ Repository tests failed$(NC)"; exit 1; }

# Run model tests only
test-models: deps-check
	@echo "$(YELLOW)📦 Testing models package...$(NC)"
	@cd $(STUDENT_SERVICE_DIR)/models && \
		go test $(GO_TEST_FLAGS) -coverprofile=models_coverage.out -covermode=atomic . && \
		echo "$(GREEN)✅ Model tests passed$(NC)" || \
		{ echo "$(RED)❌ Model tests failed$(NC)"; exit 1; }

# Run tests with coverage analysis
test-coverage: test coverage-func coverage-html
	@echo "$(GREEN)📊 Coverage analysis complete$(NC)"

# Generate HTML coverage report
coverage-html:
	@if [ -f "$(STUDENT_SERVICE_DIR)/$(COVERAGE_FILE)" ]; then \
		echo "$(YELLOW)📄 Generating HTML coverage report...$(NC)"; \
		cd $(STUDENT_SERVICE_DIR) && \
		go tool cover -html=$(COVERAGE_FILE) -o=$(COVERAGE_HTML) && \
		echo "$(GREEN)📄 HTML report: $(STUDENT_SERVICE_DIR)/$(COVERAGE_HTML)$(NC)"; \
		echo "$(GREEN)🌐 Open in browser: file://$$(pwd)/$(STUDENT_SERVICE_DIR)/$(COVERAGE_HTML)$(NC)"; \
	else \
		echo "$(RED)❌ Coverage file not found. Run 'make test' first.$(NC)"; \
	fi

# Show function coverage summary
coverage-func:
	@if [ -f "$(STUDENT_SERVICE_DIR)/$(COVERAGE_FILE)" ]; then \
		echo "$(YELLOW)📊 Function coverage summary:$(NC)"; \
		cd $(STUDENT_SERVICE_DIR) && go tool cover -func=$(COVERAGE_FILE); \
	else \
		echo "$(RED)❌ Coverage file not found. Run 'make test' first.$(NC)"; \
	fi

# Run benchmark tests
test-benchmarks: deps-check
	@echo "$(YELLOW)⚡ Running benchmark tests...$(NC)"
	@cd $(STUDENT_SERVICE_DIR) && \
		go test $(BENCH_FLAGS) ./... && \
		echo "$(GREEN)⚡ Benchmarks completed$(NC)"

# Run code linting
lint: deps-check
	@echo "$(YELLOW)🔍 Running code linting...$(NC)"
	@if command -v golangci-lint >/dev/null 2>&1; then \
		cd $(STUDENT_SERVICE_DIR) && golangci-lint run ./...; \
		echo "$(GREEN)✅ Linting completed$(NC)"; \
	else \
		echo "$(YELLOW)⚠️  golangci-lint not installed. Install with: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest$(NC)"; \
	fi

# Run go vet analysis
vet: deps-check
	@echo "$(YELLOW)🔍 Running go vet analysis...$(NC)"
	@cd $(STUDENT_SERVICE_DIR) && \
		go vet ./... && \
		echo "$(GREEN)✅ Go vet completed$(NC)"

# Format Go code
fmt: deps-check
	@echo "$(YELLOW)🎨 Formatting Go code...$(NC)"
	@cd $(STUDENT_SERVICE_DIR) && \
		go fmt ./... && \
		echo "$(GREEN)✅ Code formatting completed$(NC)"

# Run all quality checks
quality: fmt vet lint test
	@echo "$(GREEN)🎉 All quality checks completed$(NC)"

# Clean test artifacts
clean:
	@echo "$(YELLOW)🧹 Cleaning test artifacts...$(NC)"
	@cd $(STUDENT_SERVICE_DIR) && \
		find . -name "*.out" -type f -delete && \
		find . -name "*coverage.html" -type f -delete && \
		echo "$(GREEN)✅ Cleanup completed$(NC)"

# Run tests in CI mode (optimized for CI/CD)
test-ci: deps-check
	@echo "$(YELLOW)🔄 Running tests in CI mode...$(NC)"
	@cd $(STUDENT_SERVICE_DIR) && \
		go test -v -race -timeout=$(TEST_TIMEOUT) -coverprofile=$(COVERAGE_FILE) -covermode=atomic ./... \
		-json > test-results.json && \
		echo "$(GREEN)✅ CI tests completed$(NC)"

# Generate test report for CI
test-report: test-ci
	@echo "$(YELLOW)📋 Generating test report...$(NC)"
	@cd $(STUDENT_SERVICE_DIR) && \
		if [ -f "$(COVERAGE_FILE)" ]; then \
			echo "Coverage: $$(go tool cover -func=$(COVERAGE_FILE) | grep total | awk '{print $$3}')" > test-summary.txt; \
			echo "$(GREEN)📄 Test summary: $(STUDENT_SERVICE_DIR)/test-summary.txt$(NC)"; \
		fi

# Watch tests (requires entr or inotify-tools)
test-watch:
	@echo "$(YELLOW)👁️  Watching for changes and running tests...$(NC)"
	@if command -v entr >/dev/null 2>&1; then \
		find $(STUDENT_SERVICE_DIR) -name "*.go" | entr -c make test; \
	elif command -v inotifywait >/dev/null 2>&1; then \
		while inotifywait -r -e modify $(STUDENT_SERVICE_DIR); do make test; done; \
	else \
		echo "$(RED)❌ Watch requires 'entr' or 'inotify-tools'. Install with your package manager.$(NC)"; \
	fi

# Show test statistics
test-stats: deps-check
	@echo "$(YELLOW)📈 Test Statistics$(NC)"
	@echo "=================="
	@cd $(STUDENT_SERVICE_DIR) && \
		echo "📁 Test files: $$(find . -name "*_test.go" | wc -l)" && \
		echo "🧪 Test functions: $$(grep -r "^func Test" . --include="*_test.go" | wc -l)" && \
		echo "⚡ Benchmark functions: $$(grep -r "^func Benchmark" . --include="*_test.go" | wc -l)" && \
		echo "📖 Example functions: $$(grep -r "^func Example" . --include="*_test.go" | wc -l)"

# Profile memory usage during tests
test-memprofile: deps-check
	@echo "$(YELLOW)🧠 Running tests with memory profiling...$(NC)"
	@cd $(STUDENT_SERVICE_DIR) && \
		go test -memprofile=mem.prof -bench=. ./... && \
		echo "$(GREEN)📊 Memory profile: $(STUDENT_SERVICE_DIR)/mem.prof$(NC)" && \
		echo "$(YELLOW)Analyze with: go tool pprof mem.prof$(NC)"

# Profile CPU usage during tests
test-cpuprofile: deps-check
	@echo "$(YELLOW)⚡ Running tests with CPU profiling...$(NC)"
	@cd $(STUDENT_SERVICE_DIR) && \
		go test -cpuprofile=cpu.prof -bench=. ./... && \
		echo "$(GREEN)📊 CPU profile: $(STUDENT_SERVICE_DIR)/cpu.prof$(NC)" && \
		echo "$(YELLOW)Analyze with: go tool pprof cpu.prof$(NC)"

# Create test coverage badge (requires coverage-badge tool)
coverage-badge:
	@if [ -f "$(STUDENT_SERVICE_DIR)/$(COVERAGE_FILE)" ]; then \
		coverage=$$(cd $(STUDENT_SERVICE_DIR) && go tool cover -func=$(COVERAGE_FILE) | grep total | awk '{print $$3}' | sed 's/%//'); \
		echo "$(GREEN)📊 Coverage: $$coverage%$(NC)"; \
		if [ $$coverage -ge 90 ]; then \
			echo "$(GREEN)🏆 Excellent coverage!$(NC)"; \
		elif [ $$coverage -ge 80 ]; then \
			echo "$(YELLOW)👍 Good coverage$(NC)"; \
		else \
			echo "$(RED)⚠️  Coverage below 80%$(NC)"; \
		fi; \
	else \
		echo "$(RED)❌ Run tests first to generate coverage$(NC)"; \
	fi
