# [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-09-24
# PowerShell script for running tests with coverage

# Clean up previous coverage files
if (Test-Path "coverage.out") {
    Remove-Item "coverage.out"
}
if (Test-Path "coverage.html") {
    Remove-Item "coverage.html"
}
if (Test-Path "coverage-summary.txt") {
    Remove-Item "coverage-summary.txt"
}

Write-Host "Running tests with coverage..." -ForegroundColor Green
go test -v -coverprofile=coverage.out ./...

if ($LASTEXITCODE -eq 0) {
    Write-Host "Generating HTML coverage report..." -ForegroundColor Green
    go tool cover -html=coverage.out -o coverage.html
    
    Write-Host "Generating coverage summary..." -ForegroundColor Green
    go tool cover -func=coverage.out | Tee-Object -FilePath coverage-summary.txt
    
    Write-Host "`nCoverage report generated successfully!" -ForegroundColor Green
    Write-Host "HTML Report: coverage.html" -ForegroundColor Yellow
    Write-Host "Summary: coverage-summary.txt" -ForegroundColor Yellow
    
    # Display coverage summary
    Write-Host "`nCoverage Summary:" -ForegroundColor Cyan
    Get-Content coverage-summary.txt | Select-Object -Last 1
} else {
    Write-Host "Tests failed!" -ForegroundColor Red
    exit $LASTEXITCODE
}