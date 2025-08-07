# Simple test script to verify services and load data
Write-Host "=== School Management System Setup Test ===" -ForegroundColor Green

# Test 1: Check if all containers are running
Write-Host "`n1. Checking Docker containers..." -ForegroundColor Yellow
try {
    $containers = docker ps --format "{{.Names}},{{.Status}}"
    Write-Host "Running containers:"
    $containers | ForEach-Object { Write-Host "  $_" }
} catch {
    Write-Host "Error checking containers: $($_.Exception.Message)" -ForegroundColor Red
}

# Test 2: Test API Gateway
Write-Host "`n2. Testing API Gateway..." -ForegroundColor Yellow
try {
    $healthResponse = Invoke-RestMethod -Uri "http://localhost:8080/health" -TimeoutSec 10
    Write-Host "API Gateway health check: SUCCESS" -ForegroundColor Green
    Write-Host "Response: $($healthResponse | ConvertTo-Json -Compress)"
} catch {
    Write-Host "API Gateway health check: FAILED - $($_.Exception.Message)" -ForegroundColor Red
}

# Test 3: Test Couchbase
Write-Host "`n3. Testing Couchbase..." -ForegroundColor Yellow
try {
    $couchbaseResponse = Invoke-WebRequest -Uri "http://localhost:8091/pools" -TimeoutSec 10
    Write-Host "Couchbase status: $($couchbaseResponse.StatusCode)" -ForegroundColor Green
} catch {
    if ($_.Exception.Response.StatusCode -eq 401) {
        Write-Host "Couchbase is running but needs initialization (401 Unauthorized)" -ForegroundColor Yellow
    } else {
        Write-Host "Couchbase test: FAILED - $($_.Exception.Message)" -ForegroundColor Red
    }
}

# Test 4: Try to access individual services
Write-Host "`n4. Testing individual microservices..." -ForegroundColor Yellow
$services = @(
    @{Name="Student Service"; Port=8081},
    @{Name="Teacher Service"; Port=8082},
    @{Name="Academic Service"; Port=8083},
    @{Name="Achievement Service"; Port=8084}
)

foreach ($service in $services) {
    try {
        $response = Invoke-RestMethod -Uri "http://localhost:$($service.Port)/health" -TimeoutSec 5
        Write-Host "  $($service.Name): SUCCESS" -ForegroundColor Green
    } catch {
        Write-Host "  $($service.Name): FAILED - $($_.Exception.Message)" -ForegroundColor Red
    }
}

# Test 5: Initialize Couchbase if needed
Write-Host "`n5. Initializing Couchbase if needed..." -ForegroundColor Yellow
try {
    # Try to set up cluster
    $clusterBody = "memoryQuota=512&indexMemoryQuota=256"
    $clusterResponse = Invoke-RestMethod -Uri "http://localhost:8091/pools/default" -Method Post -Body $clusterBody -ContentType "application/x-www-form-urlencoded" -TimeoutSec 10
    Write-Host "Cluster setup: SUCCESS" -ForegroundColor Green
} catch {
    Write-Host "Cluster setup: $($_.Exception.Message)" -ForegroundColor Yellow
}

try {
    # Try to set up admin user
    $adminBody = "username=Administrator&password=password&port=SAME"
    $adminResponse = Invoke-RestMethod -Uri "http://localhost:8091/settings/web" -Method Post -Body $adminBody -ContentType "application/x-www-form-urlencoded" -TimeoutSec 10
    Write-Host "Admin user setup: SUCCESS" -ForegroundColor Green
} catch {
    Write-Host "Admin user setup: $($_.Exception.Message)" -ForegroundColor Yellow
}

# Test 6: Create bucket
Write-Host "`n6. Creating bucket..." -ForegroundColor Yellow
try {
    $bucketBody = "name=schoolmgmt&ramQuotaMB=256&bucketType=membase"
    $credentials = [System.Convert]::ToBase64String([System.Text.Encoding]::ASCII.GetBytes("Administrator:password"))
    $headers = @{ Authorization = "Basic $credentials" }
    
    $bucketResponse = Invoke-RestMethod -Uri "http://localhost:8091/pools/default/buckets" -Method Post -Body $bucketBody -ContentType "application/x-www-form-urlencoded" -Headers $headers -TimeoutSec 10
    Write-Host "Bucket creation: SUCCESS" -ForegroundColor Green
} catch {
    Write-Host "Bucket creation: $($_.Exception.Message)" -ForegroundColor Yellow
}

Write-Host "`n=== Setup Test Complete ===" -ForegroundColor Green
Write-Host "Next steps:" -ForegroundColor Cyan
Write-Host "1. Access Couchbase Console: http://localhost:8091 (Administrator/password)"
Write-Host "2. Try API endpoints with proper paths"
Write-Host "3. Check the service logs if there are issues: docker-compose logs [service-name]"
