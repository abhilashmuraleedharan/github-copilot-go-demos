# Script Running Guide for School Management System

This guide provides detailed instructions for running all scripts across different platforms (Windows, Linux, macOS).

## Prerequisites

Before running any scripts, ensure you have:

1. **Docker and Docker Compose** installed
2. **curl** command available
3. **Git Bash** (Windows) or native terminal (Linux/macOS)
4. **PowerShell** (Windows) for `.ps1` scripts

## Script Overview

| Script | Platform | Purpose |
|--------|----------|---------|
| `wait-for-services.sh` | Linux/macOS/WSL | Wait for all services to be healthy |
| `wait-for-services.bat` | Windows CMD | Wait for all services to be healthy |
| `load-comprehensive-data.sh` | Linux/macOS/WSL | Load comprehensive sample data |
| `load-sample-data.sh` | Linux/macOS/WSL | Load basic sample data |
| `load-sample-data.ps1` | Windows PowerShell | Load sample data using PowerShell |
| `couchbase-demo.sh` | Linux/macOS/WSL | Couchbase setup and demo |
| `couchbase-demo.ps1` | Windows PowerShell | Couchbase setup and demo |
| `run-demo.bat` | Windows CMD | Complete demo runner |

## Quick Start Commands

### Start the System
```bash
# 1. Start all services
docker-compose up -d

# 2. Wait for services to be ready
# Linux/macOS/WSL:
./scripts/wait-for-services.sh

# Windows CMD:
scripts\wait-for-services.bat

# Windows PowerShell:
.\scripts\wait-for-services.ps1  # (if available)
```

### Load Sample Data
```bash
# Linux/macOS/WSL (comprehensive data):
./scripts/load-comprehensive-data.sh

# Linux/macOS/WSL (basic data):
./scripts/load-sample-data.sh

# Windows PowerShell:
.\scripts\load-sample-data.ps1

# Using Docker container (any platform):
docker-compose exec couchbase bash -c "curl -o load-data.sh <script-url> && chmod +x load-data.sh && ./load-data.sh"
```

## Platform-Specific Instructions

### Linux/macOS

#### Prerequisites
```bash
# Install required tools
sudo apt-get update && sudo apt-get install -y curl docker.io docker-compose  # Ubuntu/Debian
# OR
brew install curl docker docker-compose  # macOS with Homebrew
```

#### Running Scripts
```bash
# Make scripts executable
chmod +x scripts/*.sh

# Start system and wait for readiness
docker-compose up -d
./scripts/wait-for-services.sh

# Load comprehensive sample data
./scripts/load-comprehensive-data.sh

# Run Couchbase demo
./scripts/couchbase-demo.sh

# Test API endpoints
curl http://localhost:8080/api/students
curl http://localhost:8080/api/teachers
```

#### Troubleshooting on Linux/macOS
```bash
# Check script permissions
ls -la scripts/

# Make all scripts executable
chmod +x scripts/*.sh

# Check if Docker is running
docker --version
docker-compose --version

# View service logs
docker-compose logs
docker-compose logs api-gateway
```

### Windows

#### Option 1: Windows Subsystem for Linux (WSL) - Recommended
```powershell
# Install WSL if not already installed
wsl --install

# Use WSL to run Linux scripts
wsl
cd /mnt/d/demo/schoolmgmt  # Adjust path as needed
chmod +x scripts/*.sh
./scripts/wait-for-services.sh
./scripts/load-comprehensive-data.sh
```

#### Option 2: Git Bash
```bash
# Use Git Bash terminal
# Navigate to project directory
cd /d/demo/schoolmgmt

# Run scripts
bash scripts/wait-for-services.sh
bash scripts/load-comprehensive-data.sh
bash scripts/couchbase-demo.sh
```

#### Option 3: Windows CMD
```cmd
REM Use native Windows batch files
docker-compose up -d
scripts\wait-for-services.bat
scripts\run-demo.bat

REM For data loading, use PowerShell or WSL
```

#### Option 4: PowerShell
```powershell
# Start Docker services
docker-compose up -d

# Wait for services (if PowerShell version available)
.\scripts\wait-for-services.ps1

# Load sample data
.\scripts\load-sample-data.ps1

# Run Couchbase demo
.\scripts\couchbase-demo.ps1

# Test endpoints
Invoke-RestMethod -Uri "http://localhost:8080/api/students" -Method Get
```

### Windows Troubleshooting

#### PowerShell Execution Policy
```powershell
# Check current execution policy
Get-ExecutionPolicy

# Set execution policy to allow local scripts
Set-ExecutionPolicy -ExecutionPolicy RemoteSigned -Scope CurrentUser

# Or run with bypass
PowerShell -ExecutionPolicy Bypass -File .\scripts\load-sample-data.ps1
```

#### Line Ending Issues
```bash
# If you get line ending errors, convert files
# Using Git Bash or WSL:
dos2unix scripts/*.sh

# Or using sed:
sed -i 's/\r$//' scripts/*.sh
```

#### Path Issues
```cmd
REM Ensure you're in the correct directory
cd /d D:\demo\schoolmgmt

REM Use full paths if needed
D:\demo\schoolmgmt\scripts\wait-for-services.bat
```

## Script Details

### wait-for-services Scripts

**Purpose**: Wait for all microservices and Couchbase to become healthy

**Linux/macOS/WSL**:
```bash
./scripts/wait-for-services.sh

# Options:
./scripts/wait-for-services.sh check    # Only check health
./scripts/wait-for-services.sh smoke    # Run smoke tests
./scripts/wait-for-services.sh info     # Show service info
./scripts/wait-for-services.sh help     # Show help
```

**Windows CMD**:
```cmd
scripts\wait-for-services.bat
```

**Expected Output**:
```
[INFO] School Management System - Service Health Checker
=================================================
[INFO] Service Information:
API Gateway:      http://localhost:8080
Student Service:  http://localhost:8081
...
[SUCCESS] All services are healthy and ready!
```

### Data Loading Scripts

**Purpose**: Load sample data into Couchbase database

**Comprehensive Data (Linux/macOS/WSL)**:
```bash
./scripts/load-comprehensive-data.sh
```

**Basic Data (Linux/macOS/WSL)**:
```bash
./scripts/load-sample-data.sh
```

**PowerShell (Windows)**:
```powershell
.\scripts\load-sample-data.ps1

# With custom parameters:
.\scripts\load-sample-data.ps1 -CouchbaseHost "localhost" -BucketName "schoolmgmt"
```

**Expected Output**:
```
[INFO] Starting data loading process...
[SUCCESS] Couchbase is ready!
[SUCCESS] Bucket schoolmgmt created successfully
[SUCCESS] Sample data loaded successfully!
[SUCCESS] Indexes created successfully!
```

### Couchbase Demo Scripts

**Purpose**: Set up Couchbase and run demonstrations

**Linux/macOS/WSL**:
```bash
./scripts/couchbase-demo.sh
```

**PowerShell (Windows)**:
```powershell
.\scripts\couchbase-demo.ps1
```

## Docker Container Method (Universal)

If you have issues with local script execution, you can run scripts inside containers:

### Load Data from Container
```bash
# Copy script to container and run
docker-compose exec couchbase bash -c "
curl -o /tmp/load-data.sh https://raw.githubusercontent.com/your-repo/main/scripts/load-comprehensive-data.sh
chmod +x /tmp/load-data.sh
/tmp/load-data.sh
"

# Or mount local scripts
docker-compose exec -v $(pwd)/scripts:/scripts couchbase bash /scripts/load-comprehensive-data.sh
```

### Execute Commands in API Gateway Container
```bash
# Test from inside the API Gateway container
docker-compose exec api-gateway sh -c "
curl http://student-service:8081/health
curl http://teacher-service:8082/health
"
```

## Environment Variables

Set these environment variables to customize script behavior:

```bash
# Service URLs
export GATEWAY_URL="http://localhost:8080"
export STUDENT_SERVICE_URL="http://localhost:8081"
export TEACHER_SERVICE_URL="http://localhost:8082"
export ACADEMIC_SERVICE_URL="http://localhost:8083"
export ACHIEVEMENT_SERVICE_URL="http://localhost:8084"
export COUCHBASE_URL="http://localhost:8091"

# Couchbase Configuration
export COUCHBASE_HOST="localhost"
export COUCHBASE_PORT="8091"
export COUCHBASE_USER="Administrator"
export COUCHBASE_PASSWORD="password"
export BUCKET_NAME="schoolmgmt"

# Timeouts
export TIMEOUT=300
export CHECK_INTERVAL=5
```

**Windows CMD**:
```cmd
set GATEWAY_URL=http://localhost:8080
set COUCHBASE_HOST=localhost
set BUCKET_NAME=schoolmgmt
```

**PowerShell**:
```powershell
$env:GATEWAY_URL = "http://localhost:8080"
$env:COUCHBASE_HOST = "localhost"
$env:BUCKET_NAME = "schoolmgmt"
```

## Complete Workflow Examples

### Full Setup and Test (Linux/macOS/WSL)
```bash
#!/bin/bash
# complete-setup.sh

echo "Starting School Management System setup..."

# 1. Start services
docker-compose down -v  # Clean start
docker-compose up -d --build

# 2. Wait for services
./scripts/wait-for-services.sh

# 3. Load sample data
./scripts/load-comprehensive-data.sh

# 4. Test API endpoints
echo "Testing API endpoints..."
curl -s http://localhost:8080/api/students | jq .
curl -s http://localhost:8080/api/teachers | jq .
curl -s http://localhost:8080/api/courses | jq .

echo "Setup completed successfully!"
```

### Full Setup and Test (Windows PowerShell)
```powershell
# complete-setup.ps1

Write-Host "Starting School Management System setup..." -ForegroundColor Green

# 1. Start services
docker-compose down -v  # Clean start
docker-compose up -d --build

# 2. Wait for services
.\scripts\wait-for-services.bat

# 3. Load sample data
.\scripts\load-sample-data.ps1

# 4. Test API endpoints
Write-Host "Testing API endpoints..." -ForegroundColor Blue
$students = Invoke-RestMethod -Uri "http://localhost:8080/api/students"
$teachers = Invoke-RestMethod -Uri "http://localhost:8080/api/teachers"
$courses = Invoke-RestMethod -Uri "http://localhost:8080/api/courses"

Write-Host "Found $($students.Count) students, $($teachers.Count) teachers, $($courses.Count) courses" -ForegroundColor Green
Write-Host "Setup completed successfully!" -ForegroundColor Green
```

## Common Issues and Solutions

### Script Permission Denied
```bash
# Linux/macOS/WSL
chmod +x scripts/*.sh

# If still failing, check file ownership
sudo chown $USER:$USER scripts/*.sh
```

### Docker Permission Issues
```bash
# Linux - add user to docker group
sudo usermod -aG docker $USER
# Then logout and login again

# Or run with sudo
sudo docker-compose up -d
```

### Couchbase Connection Issues
```bash
# Check if Couchbase is running
curl http://localhost:8091/pools

# Check Docker container status
docker-compose ps couchbase
docker-compose logs couchbase

# Wait longer for Couchbase to start
sleep 60 && ./scripts/load-sample-data.sh
```

### Network Issues
```bash
# Check if ports are available
netstat -an | grep :8080  # Linux/macOS
netstat -an | findstr :8080  # Windows

# Check Docker network
docker network ls
docker network inspect schoolmgmt-network
```

### Windows PowerShell Issues
```powershell
# Enable PowerShell script execution
Set-ExecutionPolicy -ExecutionPolicy RemoteSigned -Scope CurrentUser

# Check PowerShell version (should be 5.1 or higher)
$PSVersionTable.PSVersion

# Install PowerShell 7 if needed
winget install Microsoft.PowerShell
```

## Script Customization

### Custom Data Loading
```bash
# Edit load-comprehensive-data.sh to add your own data
vim scripts/load-comprehensive-data.sh

# Add custom student data
cat >> /tmp/custom-students.json << 'EOF'
{"id":"student-custom","type":"student","firstName":"Custom","lastName":"Student",...}
EOF
```

### Custom Health Checks
```bash
# Add custom service checks to wait-for-services.sh
check_custom_service() {
    local service_name="Custom Service"
    local service_url="http://localhost:9000"
    wait_for_service "$service_name" "$service_url" "/health"
}
```

## Monitoring and Logging

### View All Logs
```bash
# Follow all service logs
docker-compose logs -f

# View specific service logs
docker-compose logs -f api-gateway
docker-compose logs -f couchbase

# Save logs to file
docker-compose logs > system-logs.txt
```

### Resource Monitoring
```bash
# Monitor resource usage
docker stats $(docker-compose ps -q)

# Check disk usage
docker system df

# Clean up unused resources
docker system prune -f
```

## Next Steps

After running the scripts successfully:

1. **Access Couchbase Console**: http://localhost:8091 (Administrator/password)
2. **Test API Endpoints**: Use the sample cURL commands from `scripts/sample-curl-commands.md`
3. **View Documentation**: Check `QUICK_START_GUIDE.md` and `API_GATEWAY_DOCUMENTATION.md`
4. **Deploy to Kubernetes**: Use the Helm charts in `k8s/helm/school-management/`

For additional help, refer to:
- `scripts/docker-compose-commands.md` - Docker Compose operations
- `scripts/sample-curl-commands.md` - API testing commands
- `TROUBLESHOOTING.md` - Common issues and solutions
- `k8s/README.md` - Kubernetes deployment guide
