# Scripts Directory

This directory contains comprehensive utility scripts for the School Management System.

## Available Scripts

### Service Management

#### wait-for-services.sh / wait-for-services.bat
- **Purpose**: Wait for all microservices and Couchbase to become healthy
- **Platforms**: Linux/macOS/WSL (.sh) | Windows CMD (.bat)
- **Usage**: 
  ```bash
  ./scripts/wait-for-services.sh [check|smoke|info|help]
  scripts\wait-for-services.bat  # Windows
  ```
- **Features**:
  - Health checks for all services
  - Service connectivity testing
  - Comprehensive smoke tests
  - Colored output and progress indicators

### Data Loading

#### load-comprehensive-data.sh
- **Purpose**: Load comprehensive sample data into Couchbase
- **Platform**: Linux/macOS/WSL
- **Usage**: `./scripts/load-comprehensive-data.sh`
- **Features**:
  - Creates buckets and indexes
  - Loads students, teachers, courses, grades, achievements
  - Data verification and validation
  - Automatic error handling and retries

#### load-sample-data.sh
- **Purpose**: Load basic sample data (original version)
- **Platform**: Linux/macOS/WSL  
- **Usage**: `./scripts/load-sample-data.sh`
- **Features**:
  - Basic data loading functionality
  - Simple Couchbase setup

#### load-sample-data.ps1
- **Purpose**: PowerShell version for Windows data loading
- **Platform**: Windows PowerShell
- **Usage**: `.\scripts\load-sample-data.ps1`
- **Parameters**:
  ```powershell
  -CouchbaseHost "localhost"
  -CouchbasePort 8091
  -CouchbaseUser "Administrator"
  -CouchbasePassword "password"
  -BucketName "schoolmgmt"
  -Timeout 300
  ```

### Couchbase Operations

#### couchbase-demo.sh
- **Purpose**: Demonstrates Couchbase operations and data manipulation
- **Platform**: Linux/macOS/WSL
- **Usage**: `./scripts/couchbase-demo.sh`
- **Features**:
  - Sets up Couchbase cluster
  - Creates buckets and collections
  - Loads sample data
  - Demonstrates CRUD operations

#### couchbase-demo.ps1
- **Purpose**: PowerShell version of Couchbase demo for Windows
- **Platform**: Windows PowerShell
- **Usage**: `.\scripts\couchbase-demo.ps1`
- **Features**:
  - Same functionality as shell version
  - Windows-native PowerShell implementation
  - Colored output and error handling

### Windows Integration

#### run-demo.bat
- **Purpose**: Windows batch file to run the complete demo
- **Platform**: Windows CMD
- **Usage**: `scripts\run-demo.bat`
- **Features**:
  - Starts Docker Compose
  - Waits for services
  - Runs Couchbase setup
  - Executes sample operations

## Documentation

### docker-compose-commands.md
- **Purpose**: Comprehensive Docker Compose command reference
- **Content**:
  - Quick commands for service management
  - Development workflow examples
  - Production deployment commands
  - Troubleshooting and debugging commands
  - Environment-specific instructions

### sample-curl-commands.md
- **Purpose**: Complete cURL command reference for API testing
- **Content**:
  - CRUD operations for all services
  - Health check endpoints
  - Advanced queries and analytics
  - Batch operations
  - Error handling examples
  - PowerShell equivalents for Windows

### script-running-guide.md
- **Purpose**: Detailed guide for running scripts on all platforms
- **Content**:
  - Platform-specific instructions (Windows/Linux/macOS)
  - Troubleshooting common issues
  - Environment variable configuration
  - Complete workflow examples
  - Docker container execution methods

## Quick Start

### Linux/macOS/WSL
```bash
# 1. Start services
docker-compose up -d

# 2. Wait for services to be ready
./scripts/wait-for-services.sh

# 3. Load comprehensive sample data
./scripts/load-comprehensive-data.sh

# 4. Test the system
curl http://localhost:8080/api/students
```

### Windows PowerShell
```powershell
# 1. Start services
docker-compose up -d

# 2. Wait for services (use batch file)
.\scripts\wait-for-services.bat

# 3. Load sample data
.\scripts\load-sample-data.ps1

# 4. Test the system
Invoke-RestMethod -Uri "http://localhost:8080/api/students"
```

### Windows CMD
```cmd
REM 1. Start services
docker-compose up -d

REM 2. Wait for services and run demo
scripts\wait-for-services.bat
scripts\run-demo.bat

REM 3. Test endpoints manually or use PowerShell
```

## Prerequisites

Before running any scripts:

1. **Docker & Docker Compose**: Ensure both are installed and running
2. **curl**: Available on most systems, required for health checks
3. **Platform Tools**:
   - Linux/macOS: bash, standard Unix tools
   - Windows: PowerShell 5.1+ for .ps1 scripts, Git Bash for .sh scripts

## Script Permissions

### Linux/macOS/WSL
```bash
# Make scripts executable
chmod +x scripts/*.sh

# Verify permissions
ls -la scripts/
```

### Windows PowerShell
```powershell
# Allow local script execution
Set-ExecutionPolicy -ExecutionPolicy RemoteSigned -Scope CurrentUser

# Check execution policy
Get-ExecutionPolicy
```

## Environment Variables

### All Platforms
```bash
# Service URLs
GATEWAY_URL=http://localhost:8080
STUDENT_SERVICE_URL=http://localhost:8081
TEACHER_SERVICE_URL=http://localhost:8082
ACADEMIC_SERVICE_URL=http://localhost:8083
ACHIEVEMENT_SERVICE_URL=http://localhost:8084
COUCHBASE_URL=http://localhost:8091

# Couchbase Configuration
COUCHBASE_HOST=localhost
COUCHBASE_PORT=8091
COUCHBASE_USER=Administrator
COUCHBASE_PASSWORD=password
BUCKET_NAME=schoolmgmt

# Timeouts
TIMEOUT=300
CHECK_INTERVAL=5
```

### Windows CMD
```cmd
set GATEWAY_URL=http://localhost:8080
set COUCHBASE_HOST=localhost
set BUCKET_NAME=schoolmgmt
```

### PowerShell
```powershell
$env:GATEWAY_URL = "http://localhost:8080"
$env:COUCHBASE_HOST = "localhost"
$env:BUCKET_NAME = "schoolmgmt"
```

## Troubleshooting

### Common Issues

1. **Permission Denied**: 
   - Linux/macOS: `chmod +x scripts/*.sh`
   - Windows: Check PowerShell execution policy

2. **Connection Refused**: 
   - Verify services are running: `docker-compose ps`
   - Check service health: `curl http://localhost:8080/health`

3. **Timeout Errors**: 
   - Couchbase can take 1-2 minutes to start
   - Increase timeout values in scripts

4. **Port Conflicts**:
   - Check for conflicting services: `netstat -an | grep :8080`
   - Stop conflicting services or change ports

### Checking Service Health

```bash
# Individual service health checks
curl http://localhost:8080/health  # API Gateway
curl http://localhost:8081/health  # Student Service
curl http://localhost:8082/health  # Teacher Service
curl http://localhost:8083/health  # Academic Service
curl http://localhost:8084/health  # Achievement Service
curl http://localhost:8091/pools   # Couchbase
```

### Viewing Logs

```bash
# View all service logs
docker-compose logs

# Follow logs for specific service
docker-compose logs -f api-gateway
docker-compose logs -f couchbase

# Save logs to file
docker-compose logs > system-logs.txt
```

### Debug Mode

```bash
# Run scripts with debug output
bash -x scripts/wait-for-services.sh
bash -x scripts/load-comprehensive-data.sh

# PowerShell verbose mode
.\scripts\load-sample-data.ps1 -Verbose
```

## Script Execution Order

### Complete Setup Workflow
1. **Start Services**: `docker-compose up -d`
2. **Wait for Health**: `./scripts/wait-for-services.sh`
3. **Load Data**: `./scripts/load-comprehensive-data.sh`
4. **Verify Setup**: Test API endpoints
5. **Run Demos**: `./scripts/couchbase-demo.sh`

### Development Workflow
1. **Clean Start**: `docker-compose down -v`
2. **Build & Start**: `docker-compose up -d --build`
3. **Health Check**: `./scripts/wait-for-services.sh`
4. **Load Test Data**: `./scripts/load-comprehensive-data.sh`
5. **Run Tests**: Use sample cURL commands

### Production Deployment
1. **Deploy**: Use Kubernetes Helm charts (`k8s/helm/`)
2. **Health Check**: Verify all pods are ready
3. **Load Data**: Run data loading scripts against production endpoints
4. **Validate**: Run comprehensive tests

## Advanced Usage

### Custom Data Loading
```bash
# Edit data files before loading
vim scripts/load-comprehensive-data.sh

# Load custom data files
CUSTOM_DATA_DIR=/path/to/data ./scripts/load-comprehensive-data.sh
```

### Monitoring During Execution
```bash
# Monitor resource usage while scripts run
watch docker stats

# Monitor logs in real-time
docker-compose logs -f &
./scripts/load-comprehensive-data.sh
```

### Batch Operations
```bash
# Run multiple operations in sequence
./scripts/wait-for-services.sh && \
./scripts/load-comprehensive-data.sh && \
./scripts/couchbase-demo.sh
```

## Integration with CI/CD

### GitHub Actions Example
```yaml
- name: Setup School Management System
  run: |
    docker-compose up -d
    ./scripts/wait-for-services.sh
    ./scripts/load-comprehensive-data.sh
```

### Jenkins Pipeline Example
```groovy
stage('Setup SMS') {
    steps {
        sh 'docker-compose up -d'
        sh './scripts/wait-for-services.sh'
        sh './scripts/load-comprehensive-data.sh'
    }
}
```

## Couchbase Demo Features

The Couchbase demo scripts demonstrate:

### Interactive Menu Options
1. **Test Connection** - Verify Couchbase is accessible
2. **Initialize Cluster** - Set up Couchbase cluster and admin user
3. **Create Collections** - Create school management collections
4. **Create Indexes** - Create performance indexes
5. **Demo Student CRUD** - Student create, read, update operations
6. **Demo Teacher CRUD** - Teacher management operations
7. **Demo Academic CRUD** - Academic records operations
8. **Demo Achievement CRUD** - Achievement system operations
9. **Show Database Stats** - Display record counts
10. **Run All Setup** - Execute setup steps (1-4)
11. **Run All Demos** - Execute all demos (5-9)
12. **Full Demo** - Complete setup and demos

### What the Scripts Demonstrate

#### Student Management
```sql
-- Create student with full profile
INSERT INTO schoolmgmt.school.students (KEY, VALUE) VALUES ("student-001", {
    "id": "student-001",
    "first_name": "John",
    "last_name": "Doe",
    "email": "john.doe@school.edu",
    "grade": "10",
    "age": 16,
    "enrollment_date": "2024-08-01",
    "status": "active"
})
```

#### Teacher Management
```sql
-- Create teacher with department info
INSERT INTO schoolmgmt.school.teachers (KEY, VALUE) VALUES ("teacher-001", {
    "id": "teacher-001",
    "first_name": "Dr. Emma",
    "last_name": "Wilson",
    "department": "Mathematics",
    "subjects": ["Algebra", "Geometry"],
    "experience": 8
})
```

## Next Steps

After successful script execution:

1. **Access Web Interfaces**:
   - Couchbase Console: http://localhost:8091 (Administrator/password)
   - API Gateway: http://localhost:8080

2. **Test API Endpoints**:
   - Use commands from `sample-curl-commands.md`
   - Import Postman collection (if available)

3. **Review Documentation**:
   - API Documentation: `API_GATEWAY_DOCUMENTATION.md`
   - Quick Start Guide: `QUICK_START_GUIDE.md`
   - Design Document: `DESIGN_DOCUMENT.md`

4. **Deploy to Production**:
   - Kubernetes: Use Helm charts in `k8s/helm/school-management/`
   - Docker Swarm: Adapt docker-compose.yml for swarm mode

5. **Monitor and Maintain**:
   - Set up logging and monitoring
   - Configure backups
   - Implement health checks

For detailed instructions, refer to:
- `docker-compose-commands.md` - Docker Compose operations
- `sample-curl-commands.md` - API testing commands
- `script-running-guide.md` - Platform-specific running instructions
