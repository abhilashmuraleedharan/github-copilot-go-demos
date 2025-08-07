# 🔧 Quick Fix Script - Update Dockerfiles

# This script fixes the critical bug where Docker builds the wrong main.go file

Write-Host "🔧 Fixing Dockerfiles to build cmd/main.go instead of main.go..." -ForegroundColor Yellow

$services = @("student-service", "teacher-service", "academic-service", "achievement-service")

foreach ($service in $services) {
    $dockerfilePath = "services\$service\Dockerfile"
    
    if (Test-Path $dockerfilePath) {
        Write-Host "📝 Updating $dockerfilePath..." -ForegroundColor Cyan
        
        # Read the file
        $content = Get-Content $dockerfilePath
        
        # Replace the build line
        $newContent = $content -replace "go build -o $service \./main\.go", "go build -o $service ./cmd/main.go"
        
        # Write back to file
        $newContent | Set-Content $dockerfilePath
        
        Write-Host "✅ Updated $dockerfilePath" -ForegroundColor Green
    } else {
        Write-Host "⚠️ $dockerfilePath not found" -ForegroundColor Yellow
    }
}

Write-Host ""
Write-Host "🐳 Now rebuilding Docker containers..." -ForegroundColor Yellow
Write-Host "Running: docker-compose build --no-cache" -ForegroundColor Cyan

# Stop services
docker-compose down

# Rebuild containers
docker-compose build --no-cache

Write-Host ""
Write-Host "🚀 Starting services..." -ForegroundColor Yellow
docker-compose up -d

Write-Host ""
Write-Host "⏳ Waiting for services to start..." -ForegroundColor Yellow
Start-Sleep 30

Write-Host ""
Write-Host "🧪 Testing health endpoints..." -ForegroundColor Yellow

$endpoints = @(
    "http://localhost:8081/health",
    "http://localhost:8082/health", 
    "http://localhost:8083/health",
    "http://localhost:8084/health"
)

foreach ($endpoint in $endpoints) {
    try {
        $response = Invoke-WebRequest -Uri $endpoint -TimeoutSec 10
        $status = if ($response.StatusCode -eq 200) { "✅ HEALTHY" } else { "❌ UNHEALTHY" }
        Write-Host "$status - $endpoint" -ForegroundColor $(if ($response.StatusCode -eq 200) { "Green" } else { "Red" })
    } catch {
        Write-Host "❌ FAILED - $endpoint : $($_.Exception.Message)" -ForegroundColor Red
    }
}

Write-Host ""
Write-Host "🧪 Testing CRUD endpoints..." -ForegroundColor Yellow

$crudEndpoints = @(
    "http://localhost:8081/api/v1/students",
    "http://localhost:8082/api/v1/teachers",
    "http://localhost:8083/api/v1/academics", 
    "http://localhost:8084/api/v1/achievements"
)

foreach ($endpoint in $crudEndpoints) {
    try {
        $response = Invoke-WebRequest -Uri $endpoint -TimeoutSec 10
        $content = $response.Content | ConvertFrom-Json
        if ($content.success -eq $true) {
            Write-Host "✅ WORKING - $endpoint" -ForegroundColor Green
        } else {
            Write-Host "⚠️ PARTIAL - $endpoint" -ForegroundColor Yellow
        }
    } catch {
        Write-Host "❌ FAILED - $endpoint : $($_.Exception.Message)" -ForegroundColor Red
    }
}

Write-Host ""
Write-Host "🎉 Fix complete! All services should now work properly." -ForegroundColor Green
Write-Host "📋 Use the COUCHBASE_INTEGRATION_TEST.md file to test CRUD operations." -ForegroundColor Cyan
