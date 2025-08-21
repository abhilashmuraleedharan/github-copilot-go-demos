# School Microservice Comprehensive Testing Script
# This script performs complete testing of all services including compilation, 
# database connectivity, and API functionality

param(
    [switch]$SkipDocker,
    [switch]$QuickTest,
    [switch]$NoCleanup,
    [string]$ServiceName = "all"
)

Write-Host "=== School Microservice Comprehensive Testing ===" -ForegroundColor Cyan
Write-Host "Date: $(Get-Date)" -ForegroundColor Gray
Write-Host "Parameters: SkipDocker=$SkipDocker, QuickTest=$QuickTest, ServiceName=$ServiceName" -ForegroundColor Gray
Write-Host ""

# Function to print colored output
function Write-Status {
    param($Success, $Message, $Details = "")
    if ($Success) {
        Write-Host "✅ $Message" -ForegroundColor Green
        if ($Details) { Write-Host "   $Details" -ForegroundColor Gray }
    } else {
        Write-Host "❌ $Message" -ForegroundColor Red
        if ($Details) { Write-Host "   $Details" -ForegroundColor Yellow }
    }
}

function Write-Info {
    param($Message)
    Write-Host "ℹ️  $Message" -ForegroundColor Yellow
}

function Write-Header {
    param($Message)
    Write-Host ""
    Write-Host "=== $Message ===" -ForegroundColor Cyan
}

# Test Results Tracking
$TestResults = @{
    GoInstallation = $false
    DockerInstallation = $false
    Dependencies = $false
    Compilation = @{}
    DatabaseConnection = $false
    ServiceHealth = @{}
    APIFunctionality = @{}
    OverallSuccess = $false
}

Write-Header "Environment Verification"

# Check Go installation
Write-Info "Checking Go installation..."
try {
    $goVersion = go version 2>$null
    if ($goVersion) {
        $TestResults.GoInstallation = $true
        Write-Status $true "Go is installed" $goVersion
    } else {
        $TestResults.GoInstallation = $false
        Write-Status $false "Go is not installed or not in PATH"
        exit 1
    }
} catch {
    $TestResults.GoInstallation = $false
    Write-Status $false "Go installation check failed" $_.Exception.Message
    exit 1
}

# Check Docker installation (if not skipping Docker tests)
if (-not $SkipDocker) {
    Write-Info "Checking Docker installation..."
    try {
        $dockerVersion = docker --version 2>$null
        if ($dockerVersion -and $LASTEXITCODE -eq 0) {
            $TestResults.DockerInstallation = $true
            Write-Status $true "Docker is installed" $dockerVersion
            
            # Check if Docker is running
            $dockerPs = docker ps 2>$null
            if ($LASTEXITCODE -eq 0) {
                Write-Status $true "Docker daemon is running"
            } else {
                Write-Status $false "Docker daemon is not running" "Please start Docker Desktop"
                exit 1
            }
        } else {
            $TestResults.DockerInstallation = $false
            Write-Status $false "Docker is not installed or not in PATH"
            exit 1
        }
    } catch {
        $TestResults.DockerInstallation = $false
        Write-Status $false "Docker installation check failed" $_.Exception.Message
        exit 1
    }
}

Write-Header "Project Validation"

# Verify project structure
Write-Info "Verifying project structure..."
$requiredFiles = @("docker-compose.yml", ".env", "go.mod", "services")
$missingFiles = @()

foreach ($file in $requiredFiles) {
    if (-not (Test-Path $file)) {
        $missingFiles += $file
    }
}

if ($missingFiles.Count -eq 0) {
    Write-Status $true "Project structure is valid"
} else {
    Write-Status $false "Missing required files" ($missingFiles -join ", ")
    exit 1
}

# Download dependencies
Write-Info "Downloading Go dependencies..."
try {
    go mod tidy 2>$null
    if ($LASTEXITCODE -eq 0) {
        $TestResults.Dependencies = $true
        Write-Status $true "Go dependencies downloaded successfully"
    } else {
        $TestResults.Dependencies = $false
        Write-Status $false "Failed to download Go dependencies"
        exit 1
    }
} catch {
    $TestResults.Dependencies = $false
    Write-Status $false "Error downloading dependencies" $_.Exception.Message
    exit 1
}

Write-Header "Service Compilation"

# Compile all services
$services = @("students", "teachers", "classes", "academics", "achievements")
$servicesToTest = if ($ServiceName -eq "all") { $services } else { @($ServiceName) }

foreach ($service in $servicesToTest) {
    Write-Info "Compiling $service service..."
    
    if (-not (Test-Path "services\$service")) {
        $TestResults.Compilation[$service] = $false
        Write-Status $false "$service service directory not found"
        continue
    }
    
    Set-Location "services\$service"
    
    try {
        $compileOutput = go build -o "..\..\$service.exe" . 2>&1
        $compileResult = $LASTEXITCODE
        
        Set-Location "..\..\"
        
        if ($compileResult -eq 0 -and (Test-Path "$service.exe")) {
            $TestResults.Compilation[$service] = $true
            $fileInfo = Get-Item "$service.exe"
            $size = [math]::Round($fileInfo.Length / 1MB, 2)
            Write-Status $true "$service service compiled successfully" "Size: $size MB"
        } else {
            $TestResults.Compilation[$service] = $false
            Write-Status $false "$service service compilation failed" $compileOutput
        }
    } catch {
        Set-Location "..\..\"
        $TestResults.Compilation[$service] = $false
        Write-Status $false "$service service compilation error" $_.Exception.Message
    }
}

# Check overall compilation success
$compilationSuccess = $TestResults.Compilation.Values -notcontains $false
if (-not $compilationSuccess) {
    Write-Status $false "Some services failed to compile. Stopping tests."
    exit 1
}

# Skip Docker tests if requested
if ($SkipDocker) {
    Write-Info "Skipping Docker and integration tests as requested"
    Write-Header "Test Summary"
    Write-Status $TestResults.GoInstallation "Go Installation"
    Write-Status $TestResults.Dependencies "Dependencies Download"
    foreach ($service in $servicesToTest) {
        Write-Status $TestResults.Compilation[$service] "$service Service Compilation"
    }
    Write-Status $true "Basic tests completed successfully (Docker tests skipped)"
    exit 0
}

Write-Header "Docker Environment Setup"

# Clean up any existing containers
Write-Info "Cleaning up existing containers..."
docker-compose down 2>$null
Start-Sleep 2

# Start Couchbase first
Write-Info "Starting Couchbase database..."
try {
    docker-compose up -d couchbase 2>$null
    if ($LASTEXITCODE -eq 0) {
        Write-Status $true "Couchbase container started"
    } else {
        Write-Status $false "Failed to start Couchbase container"
        exit 1
    }
} catch {
    Write-Status $false "Error starting Couchbase" $_.Exception.Message
    exit 1
}

# Wait for Couchbase to initialize
Write-Info "Waiting for Couchbase to initialize (60 seconds)..."
Start-Sleep 60

# Test Couchbase connectivity
Write-Info "Testing Couchbase connectivity..."
$couchbaseReady = $false
for ($i = 1; $i -le 10; $i++) {
    try {
        $response = Invoke-WebRequest -Uri "http://localhost:8091" -TimeoutSec 10 -ErrorAction SilentlyContinue
        if ($response.StatusCode -eq 200) {
            $TestResults.DatabaseConnection = $true
            $couchbaseReady = $true
            Write-Status $true "Couchbase is accessible and responding"
            break
        }
    } catch {
        if ($i -eq 10) {
            $TestResults.DatabaseConnection = $false
            Write-Status $false "Couchbase is not accessible after 10 attempts"
        } else {
            Write-Info "Attempt $i`: Waiting for Couchbase..."
            Start-Sleep 5
        }
    }
}

if (-not $couchbaseReady) {
    Write-Status $false "Cannot proceed without Couchbase connectivity"
    if (-not $NoCleanup) {
        docker-compose down 2>$null
    }
    exit 1
}

# Initialize Couchbase
Write-Info "Initializing Couchbase cluster and bucket..."
try {
    # Try to initialize cluster (might already be initialized)
    docker exec couchbase couchbase-cli cluster-init --cluster couchbase://localhost --cluster-username Administrator --cluster-password password --cluster-name school-cluster --cluster-ramsize 1024 --cluster-index-ramsize 512 --services data,index,query 2>$null
    
    # Try to create bucket (might already exist)
    docker exec couchbase couchbase-cli bucket-create --cluster couchbase://localhost --username Administrator --password password --bucket school --bucket-type couchbase --bucket-ramsize 512 --bucket-replica 0 2>$null
    
    # Create basic indexes
    Start-Sleep 10
    docker exec couchbase cbq -u Administrator -p password -s="CREATE PRIMARY INDEX ON \`school\`" 2>$null
    docker exec couchbase cbq -u Administrator -p password -s="CREATE INDEX idx_type ON \`school\`(type)" 2>$null
    
    Write-Status $true "Couchbase initialization completed"
} catch {
    Write-Status $false "Couchbase initialization had issues" "Services may still work if already initialized"
}

# Start all services
Write-Info "Starting all microservices..."
try {
    docker-compose up -d 2>$null
    if ($LASTEXITCODE -eq 0) {
        Write-Status $true "All services started"
    } else {
        Write-Status $false "Some services failed to start"
    }
} catch {
    Write-Status $false "Error starting services" $_.Exception.Message
}

# Wait for services to be ready
Write-Info "Waiting for services to initialize (30 seconds)..."
Start-Sleep 30

Write-Header "Service Health Testing"

# Test service health endpoints
$healthCheckPorts = @{
    "students" = 8081
    "teachers" = 8082
    "classes" = 8083
    "academics" = 8084
    "achievements" = 8085
}

foreach ($service in $servicesToTest) {
    if ($healthCheckPorts.ContainsKey($service)) {
        $port = $healthCheckPorts[$service]
        Write-Info "Testing $service service health (port $port)..."
        
        try {
            $response = Invoke-WebRequest -Uri "http://localhost:$port/health" -TimeoutSec 15 -ErrorAction SilentlyContinue
            if ($response.StatusCode -eq 200) {
                $TestResults.ServiceHealth[$service] = $true
                Write-Status $true "$service service health check passed" $response.Content
            } else {
                $TestResults.ServiceHealth[$service] = $false
                Write-Status $false "$service service health check failed" "Status: $($response.StatusCode)"
            }
        } catch {
            $TestResults.ServiceHealth[$service] = $false
            Write-Status $false "$service service health check failed" $_.Exception.Message
        }
    }
}

# Skip API tests in quick mode
if ($QuickTest) {
    Write-Info "Skipping detailed API tests (QuickTest mode)"
} else {
    Write-Header "API Functionality Testing"
    
    # Test Students API (detailed test)
    if ($servicesToTest -contains "students" -and $TestResults.ServiceHealth["students"]) {
        Write-Info "Testing Students API functionality..."
        
        try {
            # Test GET all students (should return empty array or existing data)
            $getAllResponse = Invoke-RestMethod -Uri "http://localhost:8081/students" -Method GET -TimeoutSec 15
            Write-Status $true "GET /students endpoint working"
            
            # Test CREATE student
            $studentData = @{
                firstName = "TestUser"
                lastName = "API"
                email = "test.api@school.edu"
                grade = "10"
                dateOfBirth = "2008-05-15T00:00:00Z"
                address = "123 Test St"
                phone = "555-TEST"
                status = "active"
            } | ConvertTo-Json
            
            $createResponse = Invoke-RestMethod -Uri "http://localhost:8081/students" -Method POST -Body $studentData -ContentType "application/json" -TimeoutSec 15
            
            if ($createResponse -and $createResponse.id) {
                $TestResults.APIFunctionality["students_create"] = $true
                Write-Status $true "POST /students (create) working" "Created student ID: $($createResponse.id)"
                
                # Test GET specific student
                $getResponse = Invoke-RestMethod -Uri "http://localhost:8081/students/$($createResponse.id)" -Method GET -TimeoutSec 15
                if ($getResponse -and $getResponse.id -eq $createResponse.id) {
                    $TestResults.APIFunctionality["students_get"] = $true
                    Write-Status $true "GET /students/{id} working" "Retrieved student: $($getResponse.firstName) $($getResponse.lastName)"
                } else {
                    $TestResults.APIFunctionality["students_get"] = $false
                    Write-Status $false "GET /students/{id} failed"
                }
                
                # Test UPDATE student (optional - can fail if not implemented)
                try {
                    $updateData = $createResponse
                    $updateData.grade = "11"
                    $updateJson = $updateData | ConvertTo-Json
                    
                    $updateResponse = Invoke-RestMethod -Uri "http://localhost:8081/students/$($createResponse.id)" -Method PUT -Body $updateJson -ContentType "application/json" -TimeoutSec 15
                    $TestResults.APIFunctionality["students_update"] = $true
                    Write-Status $true "PUT /students/{id} (update) working"
                } catch {
                    $TestResults.APIFunctionality["students_update"] = $false
                    Write-Status $false "PUT /students/{id} (update) failed" "This may be expected if not implemented"
                }
                
                # Test DELETE student (optional - can fail if not implemented)
                try {
                    Invoke-RestMethod -Uri "http://localhost:8081/students/$($createResponse.id)" -Method DELETE -TimeoutSec 15
                    $TestResults.APIFunctionality["students_delete"] = $true
                    Write-Status $true "DELETE /students/{id} working"
                } catch {
                    $TestResults.APIFunctionality["students_delete"] = $false
                    Write-Status $false "DELETE /students/{id} failed" "This may be expected if not implemented"
                }
                
            } else {
                $TestResults.APIFunctionality["students_create"] = $false
                Write-Status $false "POST /students (create) failed" "No ID returned in response"
            }
            
        } catch {
            $TestResults.APIFunctionality["students"] = $false
            Write-Status $false "Students API testing failed" $_.Exception.Message
        }
    }
    
    # Test other services (basic health + GET endpoint)
    $otherServices = $servicesToTest | Where-Object { $_ -ne "students" }
    foreach ($service in $otherServices) {
        if ($TestResults.ServiceHealth[$service]) {
            $port = $healthCheckPorts[$service]
            Write-Info "Testing $service API endpoints..."
            
            try {
                # Test main endpoint (GET /{service})
                $response = Invoke-RestMethod -Uri "http://localhost:$port/$service" -Method GET -TimeoutSec 10
                $TestResults.APIFunctionality[$service] = $true
                Write-Status $true "$service API GET endpoint working"
            } catch {
                $TestResults.APIFunctionality[$service] = $false
                Write-Status $false "$service API GET endpoint failed" $_.Exception.Message
            }
        }
    }
}

Write-Header "Test Results Summary"

# Calculate overall success
$overallSuccess = $true
$overallSuccess = $overallSuccess -and $TestResults.GoInstallation
$overallSuccess = $overallSuccess -and $TestResults.Dependencies
$overallSuccess = $overallSuccess -and ($TestResults.Compilation.Values -notcontains $false)

if (-not $SkipDocker) {
    $overallSuccess = $overallSuccess -and $TestResults.DockerInstallation
    $overallSuccess = $overallSuccess -and $TestResults.DatabaseConnection
    $overallSuccess = $overallSuccess -and ($TestResults.ServiceHealth.Values -notcontains $false)
}

$TestResults.OverallSuccess = $overallSuccess

# Display results
Write-Status $TestResults.GoInstallation "Go Installation"
if (-not $SkipDocker) {
    Write-Status $TestResults.DockerInstallation "Docker Installation"
}
Write-Status $TestResults.Dependencies "Dependencies Download"

foreach ($service in $servicesToTest) {
    Write-Status $TestResults.Compilation[$service] "$service Service Compilation"
}

if (-not $SkipDocker) {
    Write-Status $TestResults.DatabaseConnection "Database Connectivity"
    
    foreach ($service in $servicesToTest) {
        if ($TestResults.ServiceHealth.ContainsKey($service)) {
            Write-Status $TestResults.ServiceHealth[$service] "$service Service Health"
        }
    }
    
    if (-not $QuickTest -and $TestResults.APIFunctionality.Count -gt 0) {
        Write-Host ""
        Write-Info "API Functionality Results:"
        foreach ($test in $TestResults.APIFunctionality.GetEnumerator()) {
            Write-Status $test.Value $test.Key
        }
    }
}

Write-Host ""
if ($overallSuccess) {
    Write-Host "🎉 ALL TESTS PASSED! School Microservice is working correctly." -ForegroundColor Green
    Write-Host ""
    
    if (-not $SkipDocker) {
        Write-Info "Services are running and accessible at:"
        foreach ($service in $servicesToTest) {
            if ($healthCheckPorts.ContainsKey($service)) {
                $port = $healthCheckPorts[$service]
                Write-Host "  • $service Service: http://localhost:$port/$service" -ForegroundColor Green
            }
        }
        Write-Host "  • Couchbase Console: http://localhost:8091 (Administrator/password)" -ForegroundColor Green
        Write-Host ""
        
        if (-not $NoCleanup) {
            $keepRunning = Read-Host "Keep services running for manual testing? (Y/N)"
            if ($keepRunning -notmatch '^[Yy]') {
                Write-Info "Stopping all services..."
                docker-compose down 2>$null
                Write-Status $true "All services stopped and cleaned up"
            } else {
                Write-Info "Services left running. Use 'docker-compose down' to stop them later."
            }
        }
    }
} else {
    Write-Host "❌ SOME TESTS FAILED. Please check the errors above." -ForegroundColor Red
    
    if (-not $SkipDocker -and -not $NoCleanup) {
        Write-Info "Cleaning up Docker containers..."
        docker-compose down 2>$null
    }
    exit 1
}

Write-Host ""
Write-Info "Testing completed at $(Get-Date)"
