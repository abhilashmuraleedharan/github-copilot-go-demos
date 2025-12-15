# [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-12-15
# Makefile for School Management Microservice

.PHONY: help test test-verbose test-coverage test-coverage-html coverage-report clean build run docker-build docker-run

# Default target
help:
	@echo "School Management Microservice - Available targets:"
	@echo "  make test                - Run all tests"
	@echo "  make test-verbose        - Run tests with verbose output"
	@echo "  make test-coverage       - Run tests with coverage report"
	@echo "  make test-coverage-html  - Generate HTML coverage report"
	@echo "  make coverage-report     - View coverage report in browser"
	@echo "  make clean               - Remove build artifacts and coverage files"
	@echo "  make build               - Build the application"
	@echo "  make run                 - Run the application"
	@echo "  make docker-build        - Build Docker image"
	@echo "  make docker-run          - Run Docker Compose"

# Run all tests
test:
	@echo "Running tests..."
	go test ./... -v

# Run tests with verbose output
test-verbose:
	@echo "Running tests with verbose output..."
	go test ./... -v -race

# Run tests with coverage
test-coverage:
	@echo "Running tests with coverage..."
	go test ./... -coverprofile=coverage.out -covermode=atomic
	@echo ""
	@echo "Coverage Summary:"
	go tool cover -func=coverage.out | grep total
	@echo ""
	@echo "Detailed coverage report saved to coverage.out"
	@echo "Run 'make test-coverage-html' to generate HTML report"

# Generate HTML coverage report
test-coverage-html:
	@echo "Generating HTML coverage report..."
	go test ./... -coverprofile=coverage.out -covermode=atomic
	go tool cover -html=coverage.out -o coverage.html
	@echo "HTML coverage report generated: coverage.html"
	@echo "Open coverage.html in your browser to view detailed coverage"

# View coverage report in browser (cross-platform)
coverage-report: test-coverage-html
	@echo "Opening coverage report in browser..."
	@if command -v open > /dev/null 2>&1; then \
		open coverage.html; \
	elif command -v xdg-open > /dev/null 2>&1; then \
		xdg-open coverage.html; \
	elif command -v start > /dev/null 2>&1; then \
		start coverage.html; \
	else \
		echo "Please open coverage.html manually in your browser"; \
	fi

# Clean build artifacts and coverage files
clean:
	@echo "Cleaning build artifacts and coverage files..."
	rm -f school-microservice coverage.out coverage.html
	go clean

# Build the application
build:
	@echo "Building application..."
	go build -o school-microservice .

# Run the application
run:
	@echo "Running application..."
	go run .

# Build Docker image
docker-build:
	@echo "Building Docker image..."
	docker build -t school-microservice:latest .

# Run Docker Compose
docker-run:
	@echo "Starting services with Docker Compose..."
	docker-compose up -d

# Stop Docker Compose
docker-stop:
	@echo "Stopping Docker Compose services..."
	docker-compose down

# Run tests in Docker
docker-test:
	@echo "Running tests in Docker..."
	docker run --rm school-microservice:latest go test ./... -v
