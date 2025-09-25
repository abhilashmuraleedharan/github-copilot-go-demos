#!/bin/bash
# [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-09-24
# Shell script for running tests with coverage

set -e

# Clean up previous coverage files
rm -f coverage.out coverage.html coverage-summary.txt

echo "Running tests with coverage..."
go test -v -coverprofile=coverage.out ./...

if [ $? -eq 0 ]; then
    echo "Generating HTML coverage report..."
    go tool cover -html=coverage.out -o coverage.html
    
    echo "Generating coverage summary..."
    go tool cover -func=coverage.out | tee coverage-summary.txt
    
    echo ""
    echo "Coverage report generated successfully!"
    echo "HTML Report: coverage.html"
    echo "Summary: coverage-summary.txt"
    
    # Display coverage summary
    echo ""
    echo "Coverage Summary:"
    tail -1 coverage-summary.txt
else
    echo "Tests failed!"
    exit 1
fi