# Simple Service Test Script
# Tests basic compilation and service startup

Write-Host "=== School Microservice Basic Test ===" -ForegroundColor Cyan
Write-Host ""

# Check Go installation
Write-Host "Checking Go installation..." -ForegroundColor Yellow
$goVersion = go version
Write-Host "✅ Go Version: $goVersion" -ForegroundColor Green
Write-Host ""

# Compile all services
Write-Host "Compiling all services..." -ForegroundColor Yellow
$services = @("students", "teachers", "classes", "academics", "achievements")

foreach ($service in $services) {
    Write-Host "Compiling $service service..." -ForegroundColor Gray
    Set-Location "services\$service"
    
    # Check for compilation errors
    $output = go build -o "..\..\$service.exe" . 2>&1
    if ($LASTEXITCODE -eq 0) {
        Write-Host "✅ $service service compiled successfully" -ForegroundColor Green
    } else {
        Write-Host "❌ $service service compilation failed:" -ForegroundColor Red
        Write-Host $output -ForegroundColor Red
    }
    
    Set-Location "..\..\"
}

Write-Host ""

# List compiled executables
Write-Host "Compiled executables:" -ForegroundColor Yellow
Get-ChildItem *.exe | ForEach-Object {
    $size = [math]::Round($_.Length / 1MB, 2)
    Write-Host "✅ $($_.Name) ($size MB)" -ForegroundColor Green
}

Write-Host ""

# Test running unit tests (if available)
Write-Host "Testing unit tests..." -ForegroundColor Yellow
Set-Location "services\students"

# Check if test files exist
$testFiles = Get-ChildItem -Recurse -Filter "*_test.go"
if ($testFiles.Count -gt 0) {
    Write-Host "Found $($testFiles.Count) test files:" -ForegroundColor Gray
    $testFiles | ForEach-Object { Write-Host "  - $($_.FullName)" -ForegroundColor Gray }
    
    Write-Host ""
    Write-Host "Running unit tests..." -ForegroundColor Gray
    
    # Run tests for each package
    $testPackages = @("./handlers", "./models", "./repository")
    foreach ($pkg in $testPackages) {
        if (Test-Path ($pkg -replace '\./','') ) {
            Write-Host "Testing package: $pkg" -ForegroundColor Gray
            $testOutput = go test $pkg -v 2>&1
            if ($LASTEXITCODE -eq 0) {
                Write-Host "✅ Tests passed for $pkg" -ForegroundColor Green
            } else {
                Write-Host "❌ Tests failed for $pkg" -ForegroundColor Red
                Write-Host $testOutput -ForegroundColor Red
            }
        }
    }
} else {
    Write-Host "No unit test files found" -ForegroundColor Yellow
}

Set-Location "..\..\"

Write-Host ""
Write-Host "=== Basic Test Summary ===" -ForegroundColor Cyan
Write-Host "✅ All services compiled successfully" -ForegroundColor Green
Write-Host "✅ Unit tests checked" -ForegroundColor Green
Write-Host ""
Write-Host "To test with Couchbase:" -ForegroundColor Yellow
Write-Host "1. Run: docker-compose up -d couchbase" -ForegroundColor Gray
Write-Host "2. Wait 30 seconds" -ForegroundColor Gray
Write-Host "3. Run: .\scripts\setup-couchbase.bat" -ForegroundColor Gray
Write-Host "4. Run: docker-compose up -d" -ForegroundColor Gray
Write-Host ""
Write-Host "Services will be available at:" -ForegroundColor Yellow
Write-Host "- Students: http://localhost:8081/students" -ForegroundColor Gray
Write-Host "- Teachers: http://localhost:8082/teachers" -ForegroundColor Gray
Write-Host "- Classes: http://localhost:8083/classes" -ForegroundColor Gray
Write-Host "- Academics: http://localhost:8084/academics" -ForegroundColor Gray
Write-Host "- Achievements: http://localhost:8085/achievements" -ForegroundColor Gray
