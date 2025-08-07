@echo off
REM School Management System - Kubernetes Deployment Script (Windows)
REM This script helps deploy the School Management System to a Kubernetes cluster

setlocal enabledelayedexpansion

REM Configuration
set NAMESPACE=school-demo
set CHART_NAME=school-management
set RELEASE_NAME=school-mgmt
set CHART_PATH=.\k8s\helm\school-management

REM Function to check if command exists
where kubectl >nul 2>nul
if %errorlevel% neq 0 (
    echo [ERROR] kubectl is not installed. Please install it and try again.
    exit /b 1
)

where helm >nul 2>nul
if %errorlevel% neq 0 (
    echo [ERROR] helm is not installed. Please install it and try again.
    exit /b 1
)

where docker >nul 2>nul
if %errorlevel% neq 0 (
    echo [ERROR] docker is not installed. Please install it and try again.
    exit /b 1
)

echo School Management System - Kubernetes Deployment
echo ================================================

REM Parse command line arguments
set COMMAND=%1
if "%COMMAND%"=="" set COMMAND=deploy

if "%COMMAND%"=="build" goto :build
if "%COMMAND%"=="deploy" goto :deploy
if "%COMMAND%"=="status" goto :status
if "%COMMAND%"=="cleanup" goto :cleanup
if "%COMMAND%"=="help" goto :help
if "%COMMAND%"=="-h" goto :help
if "%COMMAND%"=="--help" goto :help

echo [ERROR] Unknown command: %COMMAND%
echo Use '%~nx0 help' for usage information.
exit /b 1

:build
echo [INFO] Building Docker images...

echo [INFO] Building API Gateway image...
docker build -f k8s\dockerfiles\api-gateway.Dockerfile -t school-mgmt/api-gateway:latest .
if %errorlevel% neq 0 exit /b 1

echo [INFO] Building Student Service image...
docker build -f k8s\dockerfiles\student-service.Dockerfile -t school-mgmt/student-service:latest .
if %errorlevel% neq 0 exit /b 1

echo [INFO] Building Teacher Service image...
docker build -f k8s\dockerfiles\teacher-service.Dockerfile -t school-mgmt/teacher-service:latest .
if %errorlevel% neq 0 exit /b 1

echo [INFO] Building Academic Service image...
docker build -f k8s\dockerfiles\academic-service.Dockerfile -t school-mgmt/academic-service:latest .
if %errorlevel% neq 0 exit /b 1

echo [INFO] Building Achievement Service image...
docker build -f k8s\dockerfiles\achievement-service.Dockerfile -t school-mgmt/achievement-service:latest .
if %errorlevel% neq 0 exit /b 1

echo [SUCCESS] All Docker images built successfully
if "%COMMAND%"=="build" exit /b 0
goto :continue_deploy

:deploy
echo [INFO] Checking Kubernetes cluster connectivity...
kubectl cluster-info >nul 2>nul
if %errorlevel% neq 0 (
    echo [ERROR] Cannot connect to Kubernetes cluster. Please check your kubeconfig.
    exit /b 1
)
echo [SUCCESS] Connected to Kubernetes cluster

call :build

:continue_deploy
echo [INFO] Checking for kind cluster...
kubectl config current-context | findstr "kind" >nul 2>nul
if %errorlevel% equ 0 (
    echo [INFO] Detected kind cluster. Loading images...
    kind load docker-image school-mgmt/api-gateway:latest
    kind load docker-image school-mgmt/student-service:latest
    kind load docker-image school-mgmt/teacher-service:latest
    kind load docker-image school-mgmt/academic-service:latest
    kind load docker-image school-mgmt/achievement-service:latest
    echo [SUCCESS] Images loaded to kind cluster
)

echo [INFO] Deploying School Management System with Helm...

REM Create namespace if it doesn't exist
kubectl create namespace %NAMESPACE% --dry-run=client -o yaml | kubectl apply -f -

REM Deploy or upgrade the Helm release
helm upgrade --install %RELEASE_NAME% %CHART_PATH% --namespace %NAMESPACE% --set-string global.imageTag=latest --wait --timeout=10m
if %errorlevel% neq 0 (
    echo [ERROR] Helm deployment failed
    exit /b 1
)

echo [SUCCESS] Helm deployment completed

:status
echo [INFO] Checking deployment status...

REM Wait for all deployments to be ready
kubectl wait --for=condition=available --timeout=300s deployment --all -n %NAMESPACE%

echo.
echo [INFO] Pod status:
kubectl get pods -n %NAMESPACE%

echo.
echo [INFO] Service status:
kubectl get services -n %NAMESPACE%

REM Get ingress status if enabled
kubectl get ingress -n %NAMESPACE% >nul 2>nul
if %errorlevel% equ 0 (
    echo.
    echo [INFO] Ingress status:
    kubectl get ingress -n %NAMESPACE%
)

echo.
echo [SUCCESS] Deployment completed successfully!
echo.
echo [INFO] Access Information:

REM Get API Gateway service info
for /f "tokens=*" %%i in ('kubectl get service -n %NAMESPACE% -l app.kubernetes.io/component=api-gateway -o jsonpath="{.items[0].metadata.name}" 2^>nul') do set GATEWAY_SERVICE=%%i

if defined GATEWAY_SERVICE (
    for /f "tokens=*" %%i in ('kubectl get service %GATEWAY_SERVICE% -n %NAMESPACE% -o jsonpath="{.spec.type}"') do set SERVICE_TYPE=%%i
    
    if "!SERVICE_TYPE!"=="LoadBalancer" (
        for /f "tokens=*" %%i in ('kubectl get service %GATEWAY_SERVICE% -n %NAMESPACE% -o jsonpath="{.status.loadBalancer.ingress[0].ip}" 2^>nul') do set EXTERNAL_IP=%%i
        if defined EXTERNAL_IP (
            echo   API Gateway: http://!EXTERNAL_IP!:8080
        ) else (
            echo   [WARNING] LoadBalancer external IP is pending. Use port-forward for now.
            echo   Run: kubectl port-forward -n %NAMESPACE% svc/%GATEWAY_SERVICE% 8080:8080
        )
    ) else if "!SERVICE_TYPE!"=="NodePort" (
        for /f "tokens=*" %%i in ('kubectl get service %GATEWAY_SERVICE% -n %NAMESPACE% -o jsonpath="{.spec.ports[0].nodePort}"') do set NODE_PORT=%%i
        for /f "tokens=*" %%i in ('kubectl get nodes -o jsonpath="{.items[0].status.addresses[?(@.type==\"ExternalIP\")].address}" 2^>nul') do set NODE_IP=%%i
        if not defined NODE_IP (
            for /f "tokens=*" %%i in ('kubectl get nodes -o jsonpath="{.items[0].status.addresses[?(@.type==\"InternalIP\")].address}"') do set NODE_IP=%%i
        )
        echo   API Gateway: http://!NODE_IP!:!NODE_PORT!
    ) else (
        echo   Run: kubectl port-forward -n %NAMESPACE% svc/%GATEWAY_SERVICE% 8080:8080
        echo   Then access: http://localhost:8080
    )
)

echo.
echo [INFO] Useful commands:
echo   View logs: kubectl logs -f deployment/%RELEASE_NAME%-api-gateway -n %NAMESPACE%
echo   Scale services: kubectl scale deployment/%RELEASE_NAME%-api-gateway --replicas=3 -n %NAMESPACE%
echo   Delete deployment: helm uninstall %RELEASE_NAME% -n %NAMESPACE%

exit /b 0

:cleanup
echo [INFO] Checking Kubernetes cluster connectivity...
kubectl cluster-info >nul 2>nul
if %errorlevel% neq 0 (
    echo [ERROR] Cannot connect to Kubernetes cluster. Please check your kubeconfig.
    exit /b 1
)

echo [INFO] Cleaning up deployment...
helm uninstall %RELEASE_NAME% -n %NAMESPACE% 2>nul
kubectl delete namespace %NAMESPACE% --ignore-not-found=true
echo [SUCCESS] Cleanup completed
exit /b 0

:help
echo Usage: %~nx0 [command]
echo.
echo Commands:
echo   build    - Build Docker images only
echo   deploy   - Build images and deploy to Kubernetes (default)
echo   status   - Check deployment status and display access info
echo   cleanup  - Remove the deployment and namespace
echo   help     - Show this help message
echo.
echo Examples:
echo   %~nx0                 # Deploy the application
echo   %~nx0 build           # Build Docker images only
echo   %~nx0 status          # Check deployment status
echo   %~nx0 cleanup         # Clean up deployment
exit /b 0
