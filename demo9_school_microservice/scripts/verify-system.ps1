# Quick verification script - tests if the system is working properly
Write-Host "🔍 School Management System - Quick Verification" -ForegroundColor Green

# Test 1: Check if containers are running
Write-Host "`n📦 Checking containers..." -ForegroundColor Yellow
$containers = docker ps --format "{{.Names}}" 2>$null
if ($containers -match "schoolmgmt") {
    Write-Host "✅ Found School Management containers running" -ForegroundColor Green
    $containers | Where-Object { $_ -match "schoolmgmt" } | ForEach-Object { Write-Host "  - $_" -ForegroundColor Cyan }
} else {
    Write-Host "❌ No School Management containers found" -ForegroundColor Red
    Write-Host "Run: docker-compose up -d" -ForegroundColor Yellow
}

# Test 2: Check API Gateway
Write-Host "`n🌐 Testing API Gateway..." -ForegroundColor Yellow
try {
    $health = Invoke-RestMethod -Uri "http://localhost:8080/health" -TimeoutSec 5
    Write-Host "✅ API Gateway is healthy" -ForegroundColor Green
    Write-Host "  Status: $($health.status)" -ForegroundColor Cyan
} catch {
    Write-Host "❌ API Gateway health check failed" -ForegroundColor Red
    Write-Host "  Error: $($_.Exception.Message)" -ForegroundColor Yellow
}

# Test 3: Check Couchbase
Write-Host "`n💾 Testing Couchbase..." -ForegroundColor Yellow
try {
    $couchbase = Invoke-WebRequest -Uri "http://localhost:8091/pools" -TimeoutSec 5
    Write-Host "✅ Couchbase is accessible (Status: $($couchbase.StatusCode))" -ForegroundColor Green
} catch {
    if ($_.Exception.Response.StatusCode -eq 401) {
        Write-Host "✅ Couchbase is running (needs authentication)" -ForegroundColor Green
    } else {
        Write-Host "❌ Couchbase test failed" -ForegroundColor Red
        Write-Host "  Error: $($_.Exception.Message)" -ForegroundColor Yellow
    }
}

# Test 4: Test students endpoint
Write-Host "`n👥 Testing Students API..." -ForegroundColor Yellow
try {
    $students = Invoke-RestMethod -Uri "http://localhost:8080/api/v1/students" -TimeoutSec 5
    if ($students -is [Array]) {
        Write-Host "✅ Students API working - Found $($students.Count) students" -ForegroundColor Green
    } else {
        Write-Host "✅ Students API accessible" -ForegroundColor Green
    }
} catch {
    Write-Host "⚠️  Students API test failed" -ForegroundColor Yellow
    Write-Host "  Error: $($_.Exception.Message)" -ForegroundColor Yellow
    Write-Host "  This may be normal if services are still starting up" -ForegroundColor Cyan
}

# Test 5: Quick CRUD test
Write-Host "`n🧪 Running quick CRUD test..." -ForegroundColor Yellow
try {
    # Create a test student
    $testStudent = @{
        id = "verification-test-001"
        firstName = "Verification"
        lastName = "Test"
        email = "verification.test@school.edu"
        grade = "12"
        status = "active"
    } | ConvertTo-Json

    $created = Invoke-RestMethod -Uri "http://localhost:8080/api/v1/students" -Method Post -Body $testStudent -ContentType "application/json" -TimeoutSec 10
    Write-Host "✅ CREATE test passed" -ForegroundColor Green
    
    # Read the student back
    Start-Sleep 1
    $retrieved = Invoke-RestMethod -Uri "http://localhost:8080/api/v1/students/verification-test-001" -TimeoutSec 10
    Write-Host "✅ READ test passed" -ForegroundColor Green
    
    # Clean up - delete the test student
    Invoke-RestMethod -Uri "http://localhost:8080/api/v1/students/verification-test-001" -Method Delete -TimeoutSec 10
    Write-Host "✅ DELETE test passed" -ForegroundColor Green
    Write-Host "✅ Full CRUD test successful!" -ForegroundColor Green
    
} catch {
    Write-Host "⚠️  CRUD test failed" -ForegroundColor Yellow
    Write-Host "  Error: $($_.Exception.Message)" -ForegroundColor Yellow
    Write-Host "  The API may not be fully connected to the database yet" -ForegroundColor Cyan
}

# Summary
Write-Host "`n📋 Verification Summary:" -ForegroundColor Yellow
Write-Host "If all tests passed ✅, your system is working correctly!" -ForegroundColor Green
Write-Host "If some tests failed ⚠️/❌, try:" -ForegroundColor Yellow
Write-Host "1. Wait a few more minutes for services to fully start" -ForegroundColor Cyan
Write-Host "2. Run: docker-compose restart" -ForegroundColor Cyan
Write-Host "3. Check logs: docker-compose logs [service-name]" -ForegroundColor Cyan
Write-Host "4. Run the complete setup: .\scripts\complete-setup-windows.ps1" -ForegroundColor Cyan

Write-Host "`n🌐 Quick Access Links:" -ForegroundColor Yellow
Write-Host "- Couchbase Console: http://localhost:8091 (Administrator/password)" -ForegroundColor Cyan
Write-Host "- API Gateway Health: http://localhost:8080/health" -ForegroundColor Cyan
Write-Host "- Students API: http://localhost:8080/api/v1/students" -ForegroundColor Cyan
