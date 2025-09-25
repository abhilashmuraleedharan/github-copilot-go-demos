# [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-09-24
# Test Configuration and Scripts

## Test Coverage Configuration

### Coverage Scripts

# PowerShell script for running tests with coverage
# File: test-coverage.ps1
if (Test-Path "coverage.out") {
    Remove-Item "coverage.out"
}
if (Test-Path "coverage.html") {
    Remove-Item "coverage.html"
}

Write-Host "Running tests with coverage..." -ForegroundColor Green
go test -v -coverprofile=coverage.out ./...

if ($LASTEXITCODE -eq 0) {
    Write-Host "Generating HTML coverage report..." -ForegroundColor Green
    go tool cover -html=coverage.out -o coverage.html
    
    Write-Host "Generating coverage summary..." -ForegroundColor Green
    go tool cover -func=coverage.out | Tee-Object -FilePath coverage-summary.txt
    
    Write-Host "Coverage report generated successfully!" -ForegroundColor Green
    Write-Host "HTML Report: coverage.html" -ForegroundColor Yellow
    Write-Host "Summary: coverage-summary.txt" -ForegroundColor Yellow
} else {
    Write-Host "Tests failed!" -ForegroundColor Red
    exit $LASTEXITCODE
}

# Linux/macOS script for running tests with coverage  
# File: test-coverage.sh
#!/bin/bash
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
    
    echo "Coverage report generated successfully!"
    echo "HTML Report: coverage.html"
    echo "Summary: coverage-summary.txt"
else
    echo "Tests failed!"
    exit 1
fi