@echo off
REM Couchbase Demo Launcher for Windows
REM This batch file makes it easy to run the PowerShell demo script

echo.
echo ========================================
echo  Couchbase School Management Demo
echo ========================================
echo.

REM Check if PowerShell is available
powershell -Command "Get-Host" >nul 2>&1
if %errorlevel% neq 0 (
    echo ERROR: PowerShell is not available on this system.
    echo Please install PowerShell or run the script directly.
    pause
    exit /b 1
)

REM Check if Docker is running
docker ps >nul 2>&1
if %errorlevel% neq 0 (
    echo WARNING: Docker does not seem to be running.
    echo Please start Docker Desktop and run: docker-compose up -d
    echo.
    choice /M "Continue anyway"
    if %errorlevel% neq 1 goto :end
)

REM Set execution policy if needed
echo Setting PowerShell execution policy...
powershell -Command "Set-ExecutionPolicy -ExecutionPolicy RemoteSigned -Scope CurrentUser -Force" >nul 2>&1

REM Display menu
:menu
echo.
echo What would you like to do?
echo.
echo 1. Run Full Demo (Setup + All Demos)
echo 2. Setup Only (Initialize Couchbase)
echo 3. Demos Only (CRUD Operations)  
echo 4. Test Connection Only
echo 5. Interactive Mode
echo 6. Exit
echo.

choice /C 123456 /M "Enter your choice"

if %errorlevel%==1 (
    echo.
    echo Running full demo...
    powershell -ExecutionPolicy Bypass -File "couchbase-demo.ps1" -Action "setup"
    if %errorlevel%==0 (
        powershell -ExecutionPolicy Bypass -File "couchbase-demo.ps1" -Action "demo"
    )
    goto :continue
)

if %errorlevel%==2 (
    echo.
    echo Running setup only...
    powershell -ExecutionPolicy Bypass -File "couchbase-demo.ps1" -Action "setup"
    goto :continue
)

if %errorlevel%==3 (
    echo.
    echo Running demos only...
    powershell -ExecutionPolicy Bypass -File "couchbase-demo.ps1" -Action "demo"
    goto :continue
)

if %errorlevel%==4 (
    echo.
    echo Testing connection...
    powershell -ExecutionPolicy Bypass -File "couchbase-demo.ps1" -Action "test"
    goto :continue
)

if %errorlevel%==5 (
    echo.
    echo Starting interactive mode...
    powershell -ExecutionPolicy Bypass -File "couchbase-demo.ps1"
    goto :end
)

if %errorlevel%==6 goto :end

:continue
echo.
echo Demo completed!
echo.
choice /M "Return to menu"
if %errorlevel%==1 goto :menu

:end
echo.
echo Goodbye!
pause
