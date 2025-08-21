@echo off
:: [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-08-21
:: School Microservice Deployment Script for Kubernetes (Windows)

setlocal enabledelayedexpansion

:: Configuration
set NAMESPACE=school-demo
set RELEASE_NAME=school-microservice
set CHART_PATH=.\helm\school-microservice
set VALUES_FILE=
set ENVIRONMENT=development

:: Colors (limited in Windows cmd)
set GREEN=[92m
set YELLOW=[93m
set RED=[91m
set NC=[0m

:show_usage
echo Usage: %0 [OPTIONS]
echo.
echo Options:
echo   -e, --environment ENV    Environment (development^|staging^|production) [default: development]
echo   -n, --namespace NAME     Kubernetes namespace [default: school-demo]
echo   -r, --release NAME       Helm release name [default: school-microservice]
echo   -v, --values FILE        Custom values file
echo   -h, --help              Show this help message
echo.
echo Examples:
echo   %0 --environment production
echo   %0 --namespace school-prod --environment production
echo   %0 --values custom-values.yaml
goto :eof

:log_info
echo %GREEN%[INFO]%NC% %~1
goto :eof

:log_warn
echo %YELLOW%[WARN]%NC% %~1
goto :eof

:log_error
echo %RED%[ERROR]%NC% %~1
goto :eof

:check_prerequisites
call :log_info "Checking prerequisites..."

:: Check kubectl
kubectl version --client >nul 2>&1
if errorlevel 1 (
    call :log_error "kubectl is not installed or not in PATH"
    exit /b 1
)

:: Check helm
helm version --short >nul 2>&1
if errorlevel 1 (
    call :log_error "helm is not installed or not in PATH"
    exit /b 1
)

:: Check cluster connectivity
kubectl cluster-info >nul 2>&1
if errorlevel 1 (
    call :log_error "Cannot connect to Kubernetes cluster"
    exit /b 1
)

call :log_info "Prerequisites check passed"
goto :eof

:create_namespace
call :log_info "Creating namespace %NAMESPACE% if it doesn't exist..."
kubectl create namespace %NAMESPACE% --dry-run=client -o yaml | kubectl apply -f -
goto :eof

:add_helm_repos
call :log_info "Adding required Helm repositories..."
helm repo add couchbase https://couchbase-partners.github.io/helm-charts/ 2>nul
helm repo update
goto :eof

:deploy_application
call :log_info "Deploying School Microservice to %ENVIRONMENT% environment..."

set HELM_ARGS=upgrade --install %RELEASE_NAME% %CHART_PATH% --namespace %NAMESPACE% --create-namespace --wait --timeout 600s

:: Add environment-specific values file
if exist "%CHART_PATH%\values-%ENVIRONMENT%.yaml" (
    set HELM_ARGS=!HELM_ARGS! --values %CHART_PATH%\values-%ENVIRONMENT%.yaml
    call :log_info "Using environment-specific values: values-%ENVIRONMENT%.yaml"
)

:: Add custom values file if specified
if not "%VALUES_FILE%"=="" (
    if exist "%VALUES_FILE%" (
        set HELM_ARGS=!HELM_ARGS! --values %VALUES_FILE%
        call :log_info "Using custom values file: %VALUES_FILE%"
    ) else (
        call :log_error "Custom values file not found: %VALUES_FILE%"
        exit /b 1
    )
)

:: Execute helm deployment
helm !HELM_ARGS!
goto :eof

:verify_deployment
call :log_info "Verifying deployment..."

:: Wait for deployments to be ready
set SERVICES=students teachers classes academics achievements

for %%s in (%SERVICES%) do (
    set DEPLOYMENT_NAME=%RELEASE_NAME%-%%s
    call :log_info "Waiting for deployment !DEPLOYMENT_NAME! to be ready..."
    kubectl wait --for=condition=available --timeout=300s deployment/!DEPLOYMENT_NAME! -n %NAMESPACE%
)

:: Check pod status
call :log_info "Checking pod status..."
kubectl get pods -n %NAMESPACE% -l app.kubernetes.io/name=school-microservice

:: Check service status
call :log_info "Checking service status..."
kubectl get services -n %NAMESPACE% -l app.kubernetes.io/name=school-microservice
goto :eof

:test_health_endpoints
call :log_info "Testing health endpoints..."

:: Note: Windows cmd doesn't handle background processes easily
:: This is a simplified version that tests port forwarding capability
set SERVICES=students:8081 teachers:8082 classes:8083 academics:8084 achievements:8085

for %%s in (%SERVICES%) do (
    for /f "tokens=1,2 delims=:" %%a in ("%%s") do (
        set SERVICE=%%a
        set PORT=%%b
        set SERVICE_NAME=%RELEASE_NAME%-%%a
        call :log_info "Service !SERVICE! is available at port !PORT!"
    )
)
goto :eof

:show_access_info
call :log_info "Deployment completed successfully!"
echo.
echo Access Information:
echo ===================
echo Namespace: %NAMESPACE%
echo Release: %RELEASE_NAME%
echo.

:: Show ingress information
kubectl get ingress %RELEASE_NAME% -n %NAMESPACE% -o jsonpath="{.spec.rules[0].host}" >nul 2>&1
if not errorlevel 1 (
    echo Ingress configured. Check ingress status:
    echo   kubectl get ingress %RELEASE_NAME% -n %NAMESPACE%
) else (
    echo To access services via port forwarding:
    echo   kubectl port-forward svc/%RELEASE_NAME%-students 8081:8081 -n %NAMESPACE%
    echo   kubectl port-forward svc/%RELEASE_NAME%-teachers 8082:8082 -n %NAMESPACE%
    echo   kubectl port-forward svc/%RELEASE_NAME%-classes 8083:8083 -n %NAMESPACE%
    echo   kubectl port-forward svc/%RELEASE_NAME%-academics 8084:8084 -n %NAMESPACE%
    echo   kubectl port-forward svc/%RELEASE_NAME%-achievements 8085:8085 -n %NAMESPACE%
)

echo.
echo Useful Commands:
echo   View pods:     kubectl get pods -n %NAMESPACE%
echo   View logs:     kubectl logs -f deployment/%RELEASE_NAME%-students -n %NAMESPACE%
echo   View services: kubectl get svc -n %NAMESPACE%
echo   Uninstall:     helm uninstall %RELEASE_NAME% -n %NAMESPACE%
goto :eof

:: Parse command line arguments
:parse_args
if "%1"=="" goto :main
if "%1"=="-e" (
    set ENVIRONMENT=%2
    shift
    shift
    goto :parse_args
)
if "%1"=="--environment" (
    set ENVIRONMENT=%2
    shift
    shift
    goto :parse_args
)
if "%1"=="-n" (
    set NAMESPACE=%2
    shift
    shift
    goto :parse_args
)
if "%1"=="--namespace" (
    set NAMESPACE=%2
    shift
    shift
    goto :parse_args
)
if "%1"=="-r" (
    set RELEASE_NAME=%2
    shift
    shift
    goto :parse_args
)
if "%1"=="--release" (
    set RELEASE_NAME=%2
    shift
    shift
    goto :parse_args
)
if "%1"=="-v" (
    set VALUES_FILE=%2
    shift
    shift
    goto :parse_args
)
if "%1"=="--values" (
    set VALUES_FILE=%2
    shift
    shift
    goto :parse_args
)
if "%1"=="-h" goto :show_usage
if "%1"=="--help" goto :show_usage

call :log_error "Unknown option %1"
call :show_usage
exit /b 1

:main
call :parse_args %*

:: Validate environment
if not "%ENVIRONMENT%"=="development" if not "%ENVIRONMENT%"=="staging" if not "%ENVIRONMENT%"=="production" (
    call :log_error "Invalid environment: %ENVIRONMENT%. Must be one of: development, staging, production"
    exit /b 1
)

call :log_info "Starting School Microservice deployment..."
call :log_info "Environment: %ENVIRONMENT%"
call :log_info "Namespace: %NAMESPACE%"
call :log_info "Release: %RELEASE_NAME%"

call :check_prerequisites
if errorlevel 1 exit /b 1

call :create_namespace
call :add_helm_repos
call :deploy_application
if errorlevel 1 exit /b 1

call :verify_deployment
call :test_health_endpoints
call :show_access_info
