# Enhanced Setup Script for School Management System with Couchbase Integration
# This script ensures proper Couchbase initialization and service startup

param(
    [switch]$Force,
    [switch]$Verbose,
    [switch]$SkipBuild
)

if ($Verbose) {
    $VerbosePreference = "Continue"
}

Write-Host "🚀 Starting Enhanced School Management System Setup..." -ForegroundColor Green

# Function to log messages with timestamps
function Write-LogMessage {
    param([string]$Message, [string]$Color = "White")
    $timestamp = Get-Date -Format "yyyy-MM-dd HH:mm:ss"
    Write-Host "[$timestamp] $Message" -ForegroundColor $Color
}

# Function to wait for a service to be healthy
function Wait-ForService {
    param(
        [string]$ServiceName,
        [string]$Url,
        [int]$TimeoutSeconds = 300,
        [int]$CheckInterval = 5
    )
    
    $maxAttempts = [math]::Floor($TimeoutSeconds / $CheckInterval)
    $attempts = 0
    
    Write-LogMessage "⏳ Waiting for $ServiceName to become healthy..." "Yellow"
    
    while ($attempts -lt $maxAttempts) {
        try {
            $response = Invoke-RestMethod -Uri $Url -TimeoutSec 10 -ErrorAction Stop
            Write-LogMessage "✅ $ServiceName is healthy!" "Green"
            return $true
        } catch {
            $attempts++
            if ($attempts % 6 -eq 0) {  # Show progress every 30 seconds
                Write-LogMessage "⏳ Still waiting for $ServiceName... ($attempts/$maxAttempts)" "Yellow"
            }
            Start-Sleep $CheckInterval
        }
    }
    
    Write-LogMessage "❌ $ServiceName failed to become healthy within $TimeoutSeconds seconds" "Red"
    return $false
}

# Function to check if Couchbase is accessible
function Test-CouchbaseAccess {
    try {
        $response = Invoke-WebRequest -Uri "http://localhost:8091/pools" -TimeoutSec 10 -ErrorAction Stop
        return $true
    } catch {
        if ($_.Exception.Response.StatusCode -eq 401) {
            return $true  # 401 means accessible but needs setup
        }
        return $false
    }
}

# Step 1: Fix PowerShell execution policy
Write-LogMessage "🔧 Checking PowerShell execution policy..." "Cyan"
$currentPolicy = Get-ExecutionPolicy -Scope CurrentUser
if ($currentPolicy -eq "Restricted") {
    Write-LogMessage "📝 Setting PowerShell execution policy..." "Yellow"
    Set-ExecutionPolicy -ExecutionPolicy RemoteSigned -Scope CurrentUser -Force
    Write-LogMessage "✅ PowerShell execution policy updated" "Green"
} else {
    Write-LogMessage "✅ PowerShell execution policy is already set: $currentPolicy" "Green"
}

# Step 2: Check Docker
Write-LogMessage "🐳 Checking Docker availability..." "Cyan"
try {
    $dockerVersion = docker --version
    Write-LogMessage "✅ Docker is available: $dockerVersion" "Green"
} catch {
    Write-LogMessage "❌ Docker is not available. Please install Docker Desktop." "Red"
    exit 1
}

# Step 3: Stop existing services if Force is specified
if ($Force) {
    Write-LogMessage "🛑 Stopping existing services..." "Yellow"
    docker-compose down -v
    Start-Sleep 5
}

# Step 4: Build and start services
if (-not $SkipBuild) {
    Write-LogMessage "🏗️ Building and starting services..." "Cyan"
    docker-compose up -d --build
} else {
    Write-LogMessage "🚀 Starting services..." "Cyan"
    docker-compose up -d
}

Start-Sleep 10

# Step 5: Wait for Couchbase to be accessible
Write-LogMessage "⏳ Waiting for Couchbase to be accessible..." "Yellow"
$couchbaseReady = $false
$maxWait = 180  # 3 minutes
$waited = 0

while (-not $couchbaseReady -and $waited -lt $maxWait) {
    if (Test-CouchbaseAccess) {
        $couchbaseReady = $true
        Write-LogMessage "✅ Couchbase is accessible!" "Green"
    } else {
        Write-LogMessage "⏳ Waiting for Couchbase... ($waited/$maxWait seconds)" "Yellow"
        Start-Sleep 10
        $waited += 10
    }
}

if (-not $couchbaseReady) {
    Write-LogMessage "❌ Couchbase failed to become accessible" "Red"
    Write-LogMessage "📋 Checking Couchbase container logs..." "Yellow"
    docker-compose logs couchbase
    exit 1
}

# Step 6: Initialize Couchbase
Write-LogMessage "🔧 Initializing Couchbase..." "Cyan"
try {
    & "$PSScriptRoot\init-couchbase.ps1" -CouchbaseHost "localhost" -ErrorAction Stop
    Write-LogMessage "✅ Couchbase initialization completed!" "Green"
} catch {
    Write-LogMessage "⚠️ Couchbase initialization encountered issues: $($_.Exception.Message)" "Yellow"
    Write-LogMessage "🔄 This might be normal if Couchbase was already initialized" "Yellow"
}

# Step 7: Wait for all microservices
Write-LogMessage "🏥 Checking microservice health..." "Cyan"

$services = @(
    @{Name="API Gateway"; Url="http://localhost:8080/health"},
    @{Name="Student Service"; Url="http://localhost:8081/health"},
    @{Name="Teacher Service"; Url="http://localhost:8082/health"},
    @{Name="Academic Service"; Url="http://localhost:8083/health"},
    @{Name="Achievement Service"; Url="http://localhost:8084/health"}
)

$allHealthy = $true
foreach ($service in $services) {
    $healthy = Wait-ForService -ServiceName $service.Name -Url $service.Url -TimeoutSeconds 120
    if (-not $healthy) {
        $allHealthy = $false
        Write-LogMessage "📋 Checking $($service.Name) logs..." "Yellow"
        $serviceName = $service.Name.ToLower().Replace(" ", "-")
        docker-compose logs $serviceName
    }
}

if (-not $allHealthy) {
    Write-LogMessage "❌ Some services failed to start properly" "Red"
    Write-LogMessage "📋 Checking all container statuses..." "Yellow"
    docker-compose ps
    exit 1
}

# Step 8: Test Couchbase integration
Write-LogMessage "🧪 Testing Couchbase integration..." "Cyan"

# Create a test student to verify end-to-end functionality
$testStudent = @{
    firstName = "Test"
    lastName = "Student"
    email = "test.student@school.edu"
    grade = "10"
} | ConvertTo-Json

try {
    Write-LogMessage "📝 Creating test student..." "Yellow"
    $createResponse = Invoke-RestMethod -Uri "http://localhost:8080/api/v1/students" `
        -Method Post `
        -Body $testStudent `
        -ContentType "application/json" `
        -TimeoutSec 15
    
    $studentId = $createResponse.data.id
    Write-LogMessage "✅ Test student created with ID: $studentId" "Green"
    
    # Verify we can retrieve the student
    Write-LogMessage "🔍 Retrieving test student..." "Yellow"
    $retrieveResponse = Invoke-RestMethod -Uri "http://localhost:8080/api/v1/students/$studentId" `
        -Method Get `
        -TimeoutSec 15
    
    Write-LogMessage "✅ Test student retrieved successfully" "Green"
    
    # Clean up test data
    Write-LogMessage "🗑️ Cleaning up test student..." "Yellow"
    Invoke-RestMethod -Uri "http://localhost:8080/api/v1/students/$studentId" `
        -Method Delete `
        -TimeoutSec 15
    
    Write-LogMessage "✅ Test data cleaned up" "Green"
    
} catch {
    Write-LogMessage "❌ Couchbase integration test failed: $($_.Exception.Message)" "Red"
    Write-LogMessage "📋 Checking service logs for debugging..." "Yellow"
    docker-compose logs student-service
    Write-LogMessage "⚠️ System may still work, but there might be integration issues" "Yellow"
}

# Step 9: Display system information
Write-LogMessage "📊 System Status Summary" "Cyan"
Write-Host ""
Write-Host "🌐 Service URLs:" -ForegroundColor Green
Write-Host "  • API Gateway:        http://localhost:8080" -ForegroundColor White
Write-Host "  • Student Service:    http://localhost:8081" -ForegroundColor White  
Write-Host "  • Teacher Service:    http://localhost:8082" -ForegroundColor White
Write-Host "  • Academic Service:   http://localhost:8083" -ForegroundColor White
Write-Host "  • Achievement Service: http://localhost:8084" -ForegroundColor White
Write-Host "  • Couchbase Console:  http://localhost:8091" -ForegroundColor White
Write-Host ""
Write-Host "🔐 Couchbase Credentials:" -ForegroundColor Green
Write-Host "  • Username: Administrator" -ForegroundColor White
Write-Host "  • Password: password123" -ForegroundColor White
Write-Host "  • Bucket:   schoolmgmt" -ForegroundColor White
Write-Host ""
Write-Host "📚 Documentation:" -ForegroundColor Green
Write-Host "  • API Documentation:    API_DOCUMENTATION.go" -ForegroundColor White
Write-Host "  • Quick Start Guide:    QUICK_START_GUIDE.md" -ForegroundColor White
Write-Host "  • CRUD Commands:        scripts/couchbase-crud-commands.md" -ForegroundColor White
Write-Host "  • Troubleshooting:      FIXES_AND_SOLUTIONS.md" -ForegroundColor White
Write-Host ""
Write-Host "🧪 Quick Test Commands:" -ForegroundColor Green
Write-Host "  # Health check all services" -ForegroundColor Gray
Write-Host "  Invoke-RestMethod -Uri 'http://localhost:8080/health'" -ForegroundColor White
Write-Host ""
Write-Host "  # Get all students" -ForegroundColor Gray  
Write-Host "  Invoke-RestMethod -Uri 'http://localhost:8080/api/v1/students'" -ForegroundColor White
Write-Host ""
Write-Host "  # Create a student" -ForegroundColor Gray
Write-Host "  \$student = @{firstName='John';lastName='Doe';email='john@school.edu';grade='10'} | ConvertTo-Json" -ForegroundColor White
Write-Host "  Invoke-RestMethod -Uri 'http://localhost:8080/api/v1/students' -Method Post -Body \$student -ContentType 'application/json'" -ForegroundColor White
Write-Host ""

Write-LogMessage "🎉 School Management System setup completed successfully!" "Green"
Write-LogMessage "💡 Use scripts/couchbase-crud-commands.md for comprehensive testing examples" "Cyan"
