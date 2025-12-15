# [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-12-15
# PowerShell script for running tests with coverage on Windows

Write-Host "School Management Microservice - Test Coverage Script" -ForegroundColor Cyan
Write-Host "=======================================================" -ForegroundColor Cyan
Write-Host ""

# Navigate to project root
$projectRoot = Split-Path -Parent $PSScriptRoot
Set-Location $projectRoot

Write-Host "Running tests with coverage..." -ForegroundColor Yellow
go test ./... -coverprofile=coverage.out -covermode=atomic

if ($LASTEXITCODE -eq 0) {
    Write-Host ""
    Write-Host "Tests passed successfully!" -ForegroundColor Green
    Write-Host ""
    
    Write-Host "Coverage Summary:" -ForegroundColor Cyan
    go tool cover -func=coverage.out | Select-String "total"
    
    Write-Host ""
    Write-Host "Generating HTML coverage report..." -ForegroundColor Yellow
    go tool cover -html=coverage.out -o coverage.html
    
    Write-Host ""
    Write-Host "Coverage report generated successfully!" -ForegroundColor Green
    Write-Host "  - Text report: coverage.out" -ForegroundColor White
    Write-Host "  - HTML report: coverage.html" -ForegroundColor White
    
    Write-Host ""
    $openBrowser = Read-Host "Open HTML coverage report in browser? (Y/N)"
    if ($openBrowser -eq "Y" -or $openBrowser -eq "y") {
        Start-Process "coverage.html"
    }
} else {
    Write-Host ""
    Write-Host "Tests failed!" -ForegroundColor Red
    Write-Host "Please fix the failing tests and try again." -ForegroundColor Red
    exit 1
}
