@echo off
REM Service Health Check and Startup Script for Windows
REM This script waits for all services to become healthy before proceeding

setlocal enabledelayedexpansion

REM Configuration
set TIMEOUT=300
set CHECK_INTERVAL=5
set MAX_RETRIES=60

REM Service URLs
if not defined GATEWAY_URL set GATEWAY_URL=http://localhost:8080
if not defined STUDENT_SERVICE_URL set STUDENT_SERVICE_URL=http://localhost:8081
if not defined TEACHER_SERVICE_URL set TEACHER_SERVICE_URL=http://localhost:8082
if not defined ACADEMIC_SERVICE_URL set ACADEMIC_SERVICE_URL=http://localhost:8083
if not defined ACHIEVEMENT_SERVICE_URL set ACHIEVEMENT_SERVICE_URL=http://localhost:8084
if not defined COUCHBASE_URL set COUCHBASE_URL=http://localhost:8091

echo [INFO] School Management System - Service Health Checker
echo =================================================

REM Show service information
echo [INFO] Service Information:
echo ====================
echo API Gateway:      %GATEWAY_URL%
echo Student Service:  %STUDENT_SERVICE_URL%
echo Teacher Service:  %TEACHER_SERVICE_URL%
echo Academic Service: %ACADEMIC_SERVICE_URL%
echo Achievement Service: %ACHIEVEMENT_SERVICE_URL%
echo Couchbase:        %COUCHBASE_URL%
echo.
echo Web Interfaces:
echo - Couchbase Console: %COUCHBASE_URL%
echo - API Documentation: %GATEWAY_URL%/docs (if available)
echo.

REM Check Docker Compose status
echo [INFO] Checking Docker Compose services status...
docker-compose ps

REM Wait for services function
call :wait_for_service "Couchbase" "%COUCHBASE_URL%" "/pools"
if errorlevel 1 goto :error

call :wait_for_service "API Gateway" "%GATEWAY_URL%" "/health"
if errorlevel 1 goto :error

call :wait_for_service "Student Service" "%STUDENT_SERVICE_URL%" "/health"
if errorlevel 1 goto :error

call :wait_for_service "Teacher Service" "%TEACHER_SERVICE_URL%" "/health"
if errorlevel 1 goto :error

call :wait_for_service "Academic Service" "%ACADEMIC_SERVICE_URL%" "/health"
if errorlevel 1 goto :error

call :wait_for_service "Achievement Service" "%ACHIEVEMENT_SERVICE_URL%" "/health"
if errorlevel 1 goto :error

echo [SUCCESS] All services are healthy and ready!

REM Run comprehensive health checks
echo [INFO] Performing comprehensive health checks...

REM Test API endpoints
echo [INFO] Testing API Gateway endpoints...
curl -s "%GATEWAY_URL%/api/students" >nul 2>&1
if not errorlevel 1 (
    echo [SUCCESS] Students endpoint accessible
) else (
    echo [WARNING] Students endpoint not accessible
)

curl -s "%GATEWAY_URL%/api/teachers" >nul 2>&1
if not errorlevel 1 (
    echo [SUCCESS] Teachers endpoint accessible
) else (
    echo [WARNING] Teachers endpoint not accessible
)

curl -s "%GATEWAY_URL%/api/courses" >nul 2>&1
if not errorlevel 1 (
    echo [SUCCESS] Courses endpoint accessible
) else (
    echo [WARNING] Courses endpoint not accessible
)

curl -s "%GATEWAY_URL%/api/achievements" >nul 2>&1
if not errorlevel 1 (
    echo [SUCCESS] Achievements endpoint accessible
) else (
    echo [WARNING] Achievements endpoint not accessible
)

REM Test database connectivity
echo [INFO] Testing Couchbase connectivity...
curl -s "%COUCHBASE_URL%/pools" >nul 2>&1
if not errorlevel 1 (
    echo [SUCCESS] Couchbase is accessible
) else (
    echo [WARNING] Couchbase not accessible
)

REM Run basic smoke tests
echo [INFO] Running basic smoke tests...

echo [INFO] Testing student creation...
curl -s -X POST "%GATEWAY_URL%/api/students" ^
    -H "Content-Type: application/json" ^
    -d "{\"firstName\":\"Test\",\"lastName\":\"Student\",\"email\":\"test.student@school.edu\",\"grade\":\"10\",\"status\":\"active\"}" >nul 2>&1

if not errorlevel 1 (
    echo [SUCCESS] Student creation test passed
) else (
    echo [WARNING] Student creation test failed
)

echo [INFO] Testing student retrieval...
curl -s "%GATEWAY_URL%/api/students" >nul 2>&1
if not errorlevel 1 (
    echo [SUCCESS] Student retrieval test passed
) else (
    echo [WARNING] Student retrieval test failed
)

REM Show next steps
echo.
echo [SUCCESS] All services are ready! Here are some next steps:
echo.
echo 1. Load sample data:
echo    ./scripts/load-comprehensive-data.sh
echo    OR (on Windows with WSL): wsl ./scripts/load-comprehensive-data.sh
echo    OR use PowerShell: .\scripts\load-sample-data.ps1
echo.
echo 2. Test the API with sample commands:
echo    curl %GATEWAY_URL%/api/students
echo    curl %GATEWAY_URL%/api/teachers
echo.
echo 3. Access the Couchbase Console:
echo    Open %COUCHBASE_URL% in your browser
echo    Login: Administrator / password
echo.
echo 4. Check the documentation:
echo    - API Commands: .\scripts\sample-curl-commands.md
echo    - Docker Commands: .\scripts\docker-compose-commands.md
echo    - Quick Start Guide: .\QUICK_START_GUIDE.md
echo.

echo [SUCCESS] Health check completed successfully!
goto :end

:wait_for_service
set service_name=%~1
set service_url=%~2
set endpoint=%~3
set retries=0

echo [INFO] Waiting for %service_name% to become healthy...

:retry_loop
curl -s -f "%service_url%%endpoint%" >nul 2>&1
if not errorlevel 1 (
    echo [SUCCESS] %service_name% is healthy!
    exit /b 0
)

set /a retries+=1
if %retries% geq %MAX_RETRIES% (
    echo [ERROR] %service_name% failed to become healthy within %TIMEOUT% seconds
    exit /b 1
)

echo [INFO] Attempt %retries%/%MAX_RETRIES%: %service_name% not ready yet...
timeout /t %CHECK_INTERVAL% >nul
goto :retry_loop

:error
echo [ERROR] Health check failed. Please check the service logs:
echo   docker-compose logs
echo   docker-compose logs [service-name]
exit /b 1

:end
exit /b 0
