# Enhanced PowerShell Wait Script for School Management System
param(
    [int]$TimeoutSeconds = 300,
    [int]$CheckInterval = 10
)

# Function to write colored output
function Write-ColorOutput {
    param(
        [string]$Message,
        [string]$Color = "White"
    )
    
    $colorMap = @{
        "Red" = "Red"
        "Green" = "Green"
        "Yellow" = "Yellow"
        "Blue" = "Cyan"
        "White" = "White"
    }
    
    Write-Host $Message -ForegroundColor $colorMap[$Color]
}

function Write-Info { param([string]$Message) Write-ColorOutput "[INFO] $Message" "Blue" }
function Write-Success { param([string]$Message) Write-ColorOutput "[SUCCESS] $Message" "Green" }
function Write-Warning { param([string]$Message) Write-ColorOutput "[WARNING] $Message" "Yellow" }
function Write-Error { param([string]$Message) Write-ColorOutput "[ERROR] $Message" "Red" }

# Test if a service is healthy
function Test-ServiceHealth {
    param(
        [string]$ServiceName,
        [string]$ServiceUrl,
        [string]$Endpoint = "/health"
    )
    
    try {
        $response = Invoke-RestMethod -Uri "$ServiceUrl$Endpoint" -TimeoutSec 10 -ErrorAction Stop
        return $true
    }
    catch {
        return $false
    }
}

# Test Couchbase specifically (it returns 401 when ready but not configured)
function Test-CouchbaseHealth {
    param([string]$CouchbaseUrl = "http://localhost:8091")
    
    try {
        $response = Invoke-WebRequest -Uri "$CouchbaseUrl/pools" -TimeoutSec 10 -ErrorAction Stop
        return $true
    }
    catch {
        # 401 Unauthorized means Couchbase is running but needs setup
        if ($_.Exception.Response.StatusCode -eq 401) {
            return $true
        }
        return $false
    }
}

# Wait for a service to become healthy
function Wait-ForService {
    param(
        [string]$ServiceName,
        [string]$ServiceUrl,
        [string]$Endpoint = "/health",
        [bool]$IsCouchbase = $false
    )
    
    Write-Info "Waiting for $ServiceName to become healthy..."
    $attempts = 0
    $maxAttempts = [math]::Floor($TimeoutSeconds / $CheckInterval)
    
    while ($attempts -lt $maxAttempts) {
        $isHealthy = if ($IsCouchbase) { 
            Test-CouchbaseHealth -CouchbaseUrl $ServiceUrl 
        } else { 
            Test-ServiceHealth -ServiceName $ServiceName -ServiceUrl $ServiceUrl -Endpoint $Endpoint 
        }
        
        if ($isHealthy) {
            Write-Success "$ServiceName is healthy!"
            return $true
        }
        
        $attempts++
        Write-Info "Attempt $attempts/$maxAttempts`: $ServiceName not ready yet..."
        Start-Sleep -Seconds $CheckInterval
    }
    
    Write-Error "$ServiceName failed to become healthy within $TimeoutSeconds seconds"
    return $false
}

# Main script
Write-Info "School Management System - Service Health Checker (PowerShell)"
Write-Info "=============================================================="

# Service definitions
$services = @(
    @{ Name = "Couchbase"; Url = "http://localhost:8091"; Endpoint = "/pools"; IsCouchbase = $true },
    @{ Name = "API Gateway"; Url = "http://localhost:8080"; Endpoint = "/health"; IsCouchbase = $false },
    @{ Name = "Student Service"; Url = "http://localhost:8081"; Endpoint = "/health"; IsCouchbase = $false },
    @{ Name = "Teacher Service"; Url = "http://localhost:8082"; Endpoint = "/health"; IsCouchbase = $false },
    @{ Name = "Academic Service"; Url = "http://localhost:8083"; Endpoint = "/health"; IsCouchbase = $false },
    @{ Name = "Achievement Service"; Url = "http://localhost:8084"; Endpoint = "/health"; IsCouchbase = $false }
)

$failedServices = @()

# Check each service
foreach ($service in $services) {
    $success = Wait-ForService -ServiceName $service.Name -ServiceUrl $service.Url -Endpoint $service.Endpoint -IsCouchbase $service.IsCouchbase
    if (-not $success) {
        $failedServices += $service.Name
    }
}

# Report results
if ($failedServices.Count -eq 0) {
    Write-Success "All services are healthy and ready!"
    
    # Test basic API endpoints
    Write-Info "Testing API Gateway endpoints..."
    
    try {
        $students = Invoke-RestMethod -Uri "http://localhost:8080/api/students" -TimeoutSec 10 -ErrorAction SilentlyContinue
        Write-Success "Students endpoint accessible"
    } catch {
        Write-Warning "Students endpoint not accessible (this is normal if no data loaded yet)"
    }
    
    Write-Info ""
    Write-Success "🎉 All services are ready! Next steps:"
    Write-Host "1. Load sample data: .\scripts\load-sample-data.ps1"
    Write-Host "2. Access Couchbase Console: http://localhost:8091"
    Write-Host "3. Test API: Invoke-RestMethod -Uri 'http://localhost:8080/api/students'"
    Write-Host "4. View documentation: Get-Content .\QUICK_START_GUIDE.md"
} else {
    Write-Error "The following services failed health checks: $($failedServices -join ', ')"
    Write-Info "Troubleshooting steps:"
    Write-Host "1. Check Docker containers: docker ps"
    Write-Host "2. Check service logs: docker-compose logs [service-name]"
    Write-Host "3. Restart services: docker-compose restart"
    Write-Host "4. Check port conflicts: netstat -an | findstr ':8080'"
    exit 1
}
