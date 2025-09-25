# [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-09-24
# Makefile for School Management Microservice

.PHONY: help build test test-verbose coverage coverage-html clean deps lint fmt run docker-build docker-run

# Default target
help:
	@echo "Available targets:"
	@echo "  build          - Build the application binary"
	@echo "  test           - Run all unit tests"
	@echo "  test-verbose   - Run all unit tests with verbose output"
	@echo "  coverage       - Run tests with coverage report"
	@echo "  coverage-html  - Run tests and generate HTML coverage report"
	@echo "  clean          - Clean build artifacts and coverage reports"
	@echo "  deps           - Download and tidy dependencies"
	@echo "  lint           - Run golangci-lint (requires golangci-lint installed)"
	@echo "  fmt            - Format Go code"
	@echo "  run            - Run the application locally"
	@echo "  docker-build   - Build Docker image"
	@echo "  docker-run     - Run Docker container"

# Build the application
build:
	@echo "Building application..."
	@go build -o bin/school-service ./cmd/main.go

# Run all tests
test:
	@echo "Running unit tests..."
	@go test ./...

# Run all tests with verbose output
test-verbose:
	@echo "Running unit tests (verbose)..."
	@go test -v ./...

# Run tests with coverage
coverage:
	@echo "Running tests with coverage..."
	@go test -coverprofile=coverage.out ./...
	@go tool cover -func=coverage.out
	@echo "Coverage profile saved to coverage.out"

# Run tests and generate HTML coverage report
coverage-html:
	@echo "Running tests with coverage and generating HTML report..."
	@go test -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@go tool cover -func=coverage.out | tee coverage-summary.txt
	@echo "HTML coverage report generated: coverage.html"
	@echo "Coverage summary saved: coverage-summary.txt"

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	@rm -f bin/school-service
	@rm -f coverage.out coverage.html coverage-summary.txt
	@echo "Clean complete"

# Download and tidy dependencies
deps:
	@echo "Downloading dependencies..."
	@go mod download
	@go mod tidy

# Run golangci-lint (requires installation)
lint:
	@echo "Running golangci-lint..."
	@golangci-lint run

# Format Go code
fmt:
	@echo "Formatting Go code..."
	@go fmt ./...

# Run the application locally
run:
	@echo "Starting school microservice..."
	@go run ./cmd/main.go

# Build Docker image
docker-build:
	@echo "Building Docker image..."
	@docker build -t school-microservice:latest .

# Run Docker container
docker-run:
	@echo "Running Docker container..."
	@docker run -p 8080:8080 --env-file .env school-microservice:latest

# Development targets
dev-setup: deps
	@echo "Setting up development environment..."
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Test targets for CI/CD
test-ci: test coverage
	@echo "CI tests completed"

# Quick development cycle
dev: fmt lint test
	@echo "Development checks passed"