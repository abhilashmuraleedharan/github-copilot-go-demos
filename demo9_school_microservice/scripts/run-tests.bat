@echo off
:: [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-08-21
:: Test Configuration and Coverage Script for Windows

setlocal enabledelayedexpansion

:: Configuration
set TEST_TIMEOUT=30s
set TEST_VERBOSE=true
set TEST_RACE_DETECTION=true
set TEST_COVERAGE=true
set COVERAGE_OUTPUT_FILE=coverage.out
set COVERAGE_HTML_FILE=coverage.html
set COVERAGE_THRESHOLD_TOTAL=80
set COVERAGE_THRESHOLD_PACKAGE=75
set STUDENT_SERVICE_DIR=.\services\students

echo 🧪 School Microservice - Student Service Test Suite
echo ==================================================
echo.

:check_dependencies
echo 🔍 Checking test dependencies...

:: Check if Go is installed
go version >nul 2>&1
if errorlevel 1 (
    echo ❌ Go is not installed or not in PATH
    exit /b 1
)

:: Get Go version
for /f "tokens=3" %%i in ('go version') do set go_version=%%i
echo 📋 Go version: %go_version%

echo ✅ Dependencies check complete
echo.

:parse_args
if "%1"=="" goto run_all_tests
if "%1"=="handlers" goto run_handlers_tests
if "%1"=="repository" goto run_repository_tests
if "%1"=="models" goto run_models_tests
if "%1"=="benchmarks" goto run_benchmarks
if "%1"=="stats" goto show_test_stats
if "%1"=="clean" goto cleanup
if "%1"=="all" goto run_all_tests
goto run_all_tests

:run_handlers_tests
echo 📦 Testing package: handlers
echo Directory: %STUDENT_SERVICE_DIR%\handlers
echo.

cd %STUDENT_SERVICE_DIR%\handlers
go test -v -race -timeout=%TEST_TIMEOUT% -coverprofile=handlers_coverage.out -covermode=atomic .
set test_result=!errorlevel!

if !test_result! equ 0 (
    echo ✅ handlers tests PASSED
    
    if exist handlers_coverage.out (
        echo 📊 Generating coverage report...
        go tool cover -func=handlers_coverage.out
        go tool cover -html=handlers_coverage.out -o=handlers_coverage.html
        echo 📄 HTML coverage report: handlers_coverage.html
    )
) else (
    echo ❌ handlers tests FAILED
)

cd ..\..\..\..
echo.
goto end

:run_repository_tests
echo 📦 Testing package: repository
echo Directory: %STUDENT_SERVICE_DIR%\repository
echo.

cd %STUDENT_SERVICE_DIR%\repository
go test -v -race -timeout=%TEST_TIMEOUT% -coverprofile=repository_coverage.out -covermode=atomic .
set test_result=!errorlevel!

if !test_result! equ 0 (
    echo ✅ repository tests PASSED
    
    if exist repository_coverage.out (
        echo 📊 Generating coverage report...
        go tool cover -func=repository_coverage.out
        go tool cover -html=repository_coverage.out -o=repository_coverage.html
        echo 📄 HTML coverage report: repository_coverage.html
    )
) else (
    echo ❌ repository tests FAILED
)

cd ..\..\..\..
echo.
goto end

:run_models_tests
echo 📦 Testing package: models
echo Directory: %STUDENT_SERVICE_DIR%\models
echo.

cd %STUDENT_SERVICE_DIR%\models
go test -v -race -timeout=%TEST_TIMEOUT% -coverprofile=models_coverage.out -covermode=atomic .
set test_result=!errorlevel!

if !test_result! equ 0 (
    echo ✅ models tests PASSED
    
    if exist models_coverage.out (
        echo 📊 Generating coverage report...
        go tool cover -func=models_coverage.out
        go tool cover -html=models_coverage.out -o=models_coverage.html
        echo 📄 HTML coverage report: models_coverage.html
    )
) else (
    echo ❌ models tests FAILED
)

cd ..\..\..\..
echo.
goto end

:run_all_tests
echo 🔄 Running all Student Service tests...
echo.

cd %STUDENT_SERVICE_DIR%

:: Run all tests with combined coverage
go test -v -race -timeout=%TEST_TIMEOUT% -coverprofile=%COVERAGE_OUTPUT_FILE% -covermode=atomic ./...
set test_result=!errorlevel!

if !test_result! equ 0 (
    echo ✅ All Student Service tests PASSED
    
    if exist %COVERAGE_OUTPUT_FILE% (
        echo.
        echo 📊 Combined Coverage Report:
        go tool cover -func=%COVERAGE_OUTPUT_FILE%
        
        echo.
        echo 📄 Generating HTML coverage report...
        go tool cover -html=%COVERAGE_OUTPUT_FILE% -o=%COVERAGE_HTML_FILE%
        echo 📄 HTML coverage report: %COVERAGE_HTML_FILE%
        echo.
        echo 🌐 Open coverage report in browser:
        echo    file:///%cd:\=/%/%COVERAGE_HTML_FILE%
    )
) else (
    echo ❌ Some Student Service tests FAILED
    cd ..\..
    exit /b 1
)

cd ..\..

if "%1"=="all" goto run_benchmarks
goto show_test_stats

:run_benchmarks
echo ⚡ Running benchmarks...
echo.

cd %STUDENT_SERVICE_DIR%

:: Run benchmarks for all packages
go test -bench=. -benchmem ./...

cd ..\..
echo.

if "%1"=="benchmarks" goto end
goto show_test_stats

:show_test_stats
echo.
echo 📈 Test Statistics Summary
echo ==========================

cd %STUDENT_SERVICE_DIR%

:: Count test files
set test_files=0
for /r %%f in (*_test.go) do set /a test_files+=1
echo 📁 Test files: !test_files!

:: Count test functions (approximate)
findstr /r /c:"^func Test" *_test.go >nul 2>&1 && (
    for /f %%i in ('findstr /r /c:"^func Test" *_test.go ^| find /c /v ""') do set test_functions=%%i
) || set test_functions=0
echo 🧪 Test functions: !test_functions!

:: Count benchmark functions (approximate)
findstr /r /c:"^func Benchmark" *_test.go >nul 2>&1 && (
    for /f %%i in ('findstr /r /c:"^func Benchmark" *_test.go ^| find /c /v ""') do set benchmark_functions=%%i
) || set benchmark_functions=0
echo ⚡ Benchmark functions: !benchmark_functions!

echo.

cd ..\..

if "%1"=="stats" goto end

echo 🎉 Test suite execution complete!
echo.
echo 📋 Next steps:
echo    • Review coverage report: %COVERAGE_HTML_FILE%
echo    • Check for any failing tests
echo    • Optimize slow benchmarks if needed
echo    • Add more tests for uncovered code paths

goto end

:cleanup
echo 🧹 Cleaning up test artifacts...

cd %STUDENT_SERVICE_DIR%

:: Remove coverage files
del /q *.out 2>nul
del /q *coverage.html 2>nul

:: Remove from subdirectories
del /q handlers\*.out 2>nul
del /q handlers\*coverage.html 2>nul
del /q repository\*.out 2>nul
del /q repository\*coverage.html 2>nul
del /q models\*.out 2>nul
del /q models\*coverage.html 2>nul

cd ..\..

echo ✅ Cleanup complete

:end
echo.
pause
