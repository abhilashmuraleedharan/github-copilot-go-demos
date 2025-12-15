#!/bin/bash
# [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-12-15
# Bash script for running tests with coverage on Linux/Mac

set -e

echo "School Management Microservice - Test Coverage Script"
echo "======================================================="
echo ""

# Navigate to project root
cd "$(dirname "$0")/.."

echo "Running tests with coverage..."
go test ./... -coverprofile=coverage.out -covermode=atomic

if [ $? -eq 0 ]; then
    echo ""
    echo "Tests passed successfully!"
    echo ""
    
    echo "Coverage Summary:"
    go tool cover -func=coverage.out | grep total
    
    echo ""
    echo "Generating HTML coverage report..."
    go tool cover -html=coverage.out -o coverage.html
    
    echo ""
    echo "Coverage report generated successfully!"
    echo "  - Text report: coverage.out"
    echo "  - HTML report: coverage.html"
    
    echo ""
    read -p "Open HTML coverage report in browser? (y/n) " -n 1 -r
    echo
    if [[ $REPLY =~ ^[Yy]$ ]]; then
        if command -v open > /dev/null 2>&1; then
            open coverage.html
        elif command -v xdg-open > /dev/null 2>&1; then
            xdg-open coverage.html
        else
            echo "Please open coverage.html manually in your browser"
        fi
    fi
else
    echo ""
    echo "Tests failed!"
    echo "Please fix the failing tests and try again."
    exit 1
fi
