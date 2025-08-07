# Kubernetes Deployment

This directory contains the Kubernetes deployment configurations for the School Management System microservices.

## Directory Structure

```
k8s/
├── README.md                    # This file
├── deploy.sh                    # Linux/macOS deployment script
├── deploy.bat                   # Windows deployment script
├── dockerfiles/                 # Docker build files
│   ├── api-gateway.Dockerfile
│   ├── student-service.Dockerfile
│   ├── teacher-service.Dockerfile
│   ├── academic-service.Dockerfile
│   └── achievement-service.Dockerfile
└── helm/                        # Helm charts
    └── school-management/       # Main Helm chart
        ├── Chart.yaml
        ├── values.yaml
        └── templates/
            ├── _helpers.tpl
            ├── namespace.yaml
            ├── deployments.yaml
            ├── services.yaml
            ├── ingress.yaml
            ├── secrets.yaml
            ├── couchbase.yaml
            ├── networkpolicy.yaml
            ├── servicemonitor.yaml
            └── hpa.yaml
```

## Prerequisites

- **Docker**: For building container images
- **Kubernetes cluster**: minikube, kind, Docker Desktop, or cloud provider (GKE, EKS, AKS)
- **Helm 3.x**: For deployment management
- **kubectl**: Configured to connect to your cluster

## Quick Start

### Option 1: Using Deployment Scripts (Recommended)

#### Linux/macOS:
```bash
# Make the script executable
chmod +x k8s/deploy.sh

# Deploy everything (builds images and deploys to cluster)
./k8s/deploy.sh

# Or specific commands:
./k8s/deploy.sh build    # Build images only
./k8s/deploy.sh status   # Check deployment status
./k8s/deploy.sh cleanup  # Remove deployment
```

#### Windows:
```cmd
# Deploy everything (builds images and deploys to cluster)
k8s\deploy.bat

# Or specific commands:
k8s\deploy.bat build    # Build images only
k8s\deploy.bat status   # Check deployment status
k8s\deploy.bat cleanup  # Remove deployment
```

### Option 2: Manual Deployment

1. **Build Docker Images**:
   ```bash
   # From the project root directory
   docker build -f k8s/dockerfiles/api-gateway.Dockerfile -t school-mgmt/api-gateway:latest .
   docker build -f k8s/dockerfiles/student-service.Dockerfile -t school-mgmt/student-service:latest .
   docker build -f k8s/dockerfiles/teacher-service.Dockerfile -t school-mgmt/teacher-service:latest .
   docker build -f k8s/dockerfiles/academic-service.Dockerfile -t school-mgmt/academic-service:latest .
   docker build -f k8s/dockerfiles/achievement-service.Dockerfile -t school-mgmt/achievement-service:latest .
   ```

2. **Load Images to Kind (if using kind)**:
   ```bash
   kind load docker-image school-mgmt/api-gateway:latest
   kind load docker-image school-mgmt/student-service:latest
   kind load docker-image school-mgmt/teacher-service:latest
   kind load docker-image school-mgmt/academic-service:latest
   kind load docker-image school-mgmt/achievement-service:latest
   ```

3. **Deploy with Helm**:
   ```bash
   # Create namespace
   kubectl create namespace school-demo
   
   # Deploy the application
   helm install school-mgmt ./k8s/helm/school-management --namespace school-demo --wait
   ```

4. **Access the Application**:
   ```bash
   # Port forward to access the API Gateway
   kubectl port-forward -n school-demo svc/school-mgmt-api-gateway 8080:8080
   
   # Access the application at http://localhost:8080
kubectl get pods -n school-demo
```

## 📁 Structure

```
k8s/
├── helm/
│   └── school-management/
│       ├── Chart.yaml
│       ├── values.yaml
│       └── templates/
│           ├── api-gateway/
│           ├── student-service/
│           ├── teacher-service/
│           ├── academic-service/
│           ├── achievement-service/
│           ├── couchbase/
│           └── common/
├── dockerfiles/
│   ├── api-gateway.Dockerfile
│   ├── student-service.Dockerfile
│   ├── teacher-service.Dockerfile
│   ├── academic-service.Dockerfile
│   └── achievement-service.Dockerfile
└── README.md
```

## 🔧 Configuration

### Environment Variables
All services support configuration via environment variables:
- `PORT` - Service port
- `SERVICE_NAME` - Service identifier
- `COUCHBASE_URL` - Database connection string
- `LOG_LEVEL` - Logging level

### Resource Limits
Default resource allocations:
- CPU: 100m-500m
- Memory: 128Mi-512Mi
- Storage: 1Gi (for Couchbase)

## 📊 Monitoring

Health checks are configured for all services:
- **Liveness Probe**: `/health` endpoint
- **Readiness Probe**: `/health` endpoint  
- **Startup Probe**: 30s timeout

## 🔄 Scaling

Services can be scaled using Helm values:
```yaml
replicaCount:
  apiGateway: 2
  studentService: 3
  teacherService: 2
  academicService: 2
  achievementService: 2
```

## 🛠️ Development

### Building Images
```bash
# Build all images
./scripts/build-images.sh

# Build specific service
docker build -f k8s/dockerfiles/student-service.Dockerfile -t school-mgmt/student-service:latest .
```

### Local Testing
```bash
# Deploy to local cluster (minikube/kind)
helm install school-management ./helm/school-management -n school-demo --set global.environment=development
```

## 🚀 Production Deployment

### Image Registry
Update `values.yaml` with your container registry:
```yaml
global:
  imageRegistry: "your-registry.com/school-mgmt"
  imageTag: "v1.0.0"
```

### Ingress Configuration
Enable ingress for external access:
```yaml
ingress:
  enabled: true
  className: "nginx"
  hosts:
    - host: school-api.your-domain.com
      paths:
        - path: /
          pathType: Prefix
```

### Persistence
Configure persistent storage for Couchbase:
```yaml
couchbase:
  persistence:
    enabled: true
    storageClass: "fast-ssd"
    size: 20Gi
```

## 🔐 Security

### Service Accounts
Each service runs with its own service account with minimal required permissions.

### Network Policies
Network policies restrict communication between services to required connections only.

### Secrets Management
Database credentials and API keys are stored in Kubernetes secrets.

## 📈 Observability

### Logging
All services output structured JSON logs that can be collected by:
- Fluentd/Fluent Bit
- Elasticsearch
- CloudWatch/Azure Monitor

### Metrics
Services expose Prometheus metrics at `/metrics` endpoint.

### Tracing
OpenTelemetry compatible tracing is available for request tracking.

## 🛡️ High Availability

### Service Mesh
Optional Istio service mesh integration for:
- Traffic management
- Security policies
- Observability

### Database Clustering
Couchbase can be deployed in cluster mode for high availability:
```yaml
couchbase:
  cluster:
    enabled: true
    nodes: 3
```

## 📚 Additional Resources

- [Helm Chart Documentation](./helm/school-management/README.md)
- [Kubernetes Best Practices](https://kubernetes.io/docs/concepts/configuration/overview/)
- [Couchbase Operator](https://docs.couchbase.com/operator/current/overview.html)

---

For detailed deployment instructions, see the individual service documentation in the `helm/school-management/templates/` directory.
