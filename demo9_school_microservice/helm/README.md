# School Microservice Helm Chart

## Overview

This Helm chart deploys the School Management Microservice along with Couchbase database to a Kubernetes cluster in the `school-demo` namespace.

## Prerequisites

- Kubernetes 1.19+ cluster
- Helm 3.0+
- kubectl configured to access your cluster
- Docker image of the school microservice built and available

## Building the Docker Image

Before deploying to Kubernetes, build the Docker image:

```bash
# Build the image
docker build -t school-microservice:1.0.0 .

# Tag for your registry (if pushing to a registry)
docker tag school-microservice:1.0.0 <your-registry>/school-microservice:1.0.0

# Push to registry (if using remote registry)
docker push <your-registry>/school-microservice:1.0.0
```

For local Kubernetes clusters (like Minikube or Docker Desktop), you can use the locally built image directly.

## Installation

### 1. Install with Default Values

Deploy to the `school-demo` namespace:

```bash
cd helm
helm install school-release school-microservice --create-namespace
```

### 2. Install with Custom Values

Create a custom values file `my-values.yaml`:

```yaml
namespace: school-demo

schoolService:
  replicaCount: 3
  image:
    repository: your-registry/school-microservice
    tag: "1.0.0"

couchbase:
  auth:
    username: "YourAdmin"
    password: "YourSecurePassword"
  persistence:
    size: 10Gi
```

Install with custom values:

```bash
helm install school-release school-microservice -f my-values.yaml --create-namespace
```

### 3. Upgrade an Existing Release

```bash
helm upgrade school-release school-microservice -f my-values.yaml
```

## Configuration

### Key Configuration Parameters

| Parameter | Description | Default |
|-----------|-------------|---------|
| `namespace` | Kubernetes namespace for deployment | `school-demo` |
| `schoolService.replicaCount` | Number of service replicas | `2` |
| `schoolService.image.repository` | Docker image repository | `school-microservice` |
| `schoolService.image.tag` | Docker image tag | `1.0.0` |
| `schoolService.service.type` | Kubernetes service type | `ClusterIP` |
| `schoolService.service.port` | Service port | `8080` |
| `couchbase.enabled` | Enable Couchbase deployment | `true` |
| `couchbase.auth.username` | Couchbase admin username | `Administrator` |
| `couchbase.auth.password` | Couchbase admin password | `password` |
| `couchbase.bucket.name` | Couchbase bucket name | `school` |
| `couchbase.persistence.enabled` | Enable persistent storage | `true` |
| `couchbase.persistence.size` | PVC size for Couchbase | `5Gi` |
| `ingress.enabled` | Enable ingress | `false` |

### Full Configuration Reference

See [values.yaml](school-microservice/values.yaml) for all available configuration options.

## Accessing the Application

### Using Port Forwarding (ClusterIP)

```bash
# Forward service port to localhost
kubectl port-forward -n school-demo svc/school-service 8080:8080

# Access the API
curl http://localhost:8080/health
```

### Using NodePort

Update `values.yaml`:

```yaml
schoolService:
  service:
    type: NodePort
```

Get the NodePort:

```bash
export NODE_PORT=$(kubectl get --namespace school-demo -o jsonpath="{.spec.ports[0].nodePort}" services school-service)
export NODE_IP=$(kubectl get nodes --namespace school-demo -o jsonpath="{.items[0].status.addresses[0].address}")
echo "http://$NODE_IP:$NODE_PORT"
```

### Using Ingress

Enable ingress in `values.yaml`:

```yaml
ingress:
  enabled: true
  className: nginx
  hosts:
    - host: school-api.example.com
      paths:
        - path: /
          pathType: Prefix
```

## Verification

### Check Deployment Status

```bash
# Check all resources
kubectl get all -n school-demo

# Check pods
kubectl get pods -n school-demo

# Check services
kubectl get svc -n school-demo

# Check Couchbase initialization job
kubectl get jobs -n school-demo
```

### View Logs

```bash
# Service logs
kubectl logs -n school-demo -l app=school-service -f

# Couchbase logs
kubectl logs -n school-demo -l app=couchbase -f

# Couchbase initialization logs
kubectl logs -n school-demo job/couchbase-init
```

### Test the API

```bash
# Health check
curl http://localhost:8080/health

# Create a student
curl -X POST http://localhost:8080/api/students \
  -H "Content-Type: application/json" \
  -d '{
    "id": "student001",
    "firstName": "Alice",
    "lastName": "Johnson",
    "grade": 10,
    "dateOfBirth": "2008-05-15",
    "email": "alice.j@school.com"
  }'

# Get students
curl http://localhost:8080/api/students
```

## Accessing Couchbase Web Console

```bash
# Port forward Couchbase console
kubectl port-forward -n school-demo svc/couchbase-service 8091:8091

# Open in browser: http://localhost:8091
# Username: Administrator (or your configured username)
# Password: password (or your configured password)
```

## Uninstallation

```bash
# Uninstall the release
helm uninstall school-release -n school-demo

# Delete the namespace (optional)
kubectl delete namespace school-demo
```

## Troubleshooting

### Pods Not Starting

```bash
# Describe pod for details
kubectl describe pod -n school-demo <pod-name>

# Check events
kubectl get events -n school-demo --sort-by='.lastTimestamp'
```

### Image Pull Errors

If using a private registry, create an image pull secret:

```bash
kubectl create secret docker-registry regcred \
  --docker-server=<your-registry-server> \
  --docker-username=<your-username> \
  --docker-password=<your-password> \
  -n school-demo
```

Update `values.yaml`:

```yaml
schoolService:
  image:
    pullSecrets:
      - regcred
```

### Couchbase Initialization Fails

Check the initialization job logs:

```bash
kubectl logs -n school-demo job/couchbase-init

# If job fails, delete it and it will retry
kubectl delete job -n school-demo couchbase-init
```

### Service Connection Issues

Verify service endpoints:

```bash
kubectl get endpoints -n school-demo

# Test internal connectivity
kubectl run -it --rm debug --image=curlimages/curl --restart=Never -n school-demo -- sh
# Inside the pod:
curl http://school-service:8080/health
curl http://couchbase-service:8091/ui/index.html
```

## Helm Commands Reference

```bash
# List releases
helm list -n school-demo

# Get release status
helm status school-release -n school-demo

# Get release values
helm get values school-release -n school-demo

# Get all release information
helm get all school-release -n school-demo

# Rollback to previous version
helm rollback school-release -n school-demo

# Dry run (test without installing)
helm install school-release school-microservice --dry-run --debug

# Template rendering (see generated manifests)
helm template school-release school-microservice
```

## Production Considerations

### Security

1. **Use Secrets for Sensitive Data**: Don't store passwords in values.yaml
   ```bash
   kubectl create secret generic couchbase-creds \
     --from-literal=username=admin \
     --from-literal=password=secure-password \
     -n school-demo
   ```

2. **Enable Pod Security**: Configure security contexts and network policies

3. **Use RBAC**: Create service accounts with minimal required permissions

### High Availability

1. **Increase Replicas**:
   ```yaml
   schoolService:
     replicaCount: 3
   ```

2. **Configure Pod Disruption Budgets**:
   ```yaml
   apiVersion: policy/v1
   kind: PodDisruptionBudget
   metadata:
     name: school-service-pdb
   spec:
     minAvailable: 2
     selector:
       matchLabels:
         app: school-service
   ```

3. **Use Couchbase Operator**: For production, consider using the [Couchbase Autonomous Operator](https://docs.couchbase.com/operator/current/overview.html) instead of a simple deployment

### Monitoring

1. **Add Prometheus Annotations**:
   ```yaml
   schoolService:
     podAnnotations:
       prometheus.io/scrape: "true"
       prometheus.io/port: "8080"
       prometheus.io/path: "/metrics"
   ```

2. **Configure Resource Limits**: Ensure proper resource allocation

3. **Set Up Alerts**: Create alerting rules for critical metrics

## Support

For issues and questions:
- Check the [troubleshooting guide](#troubleshooting)
- Review Kubernetes events: `kubectl get events -n school-demo`
- Check application logs: `kubectl logs -n school-demo -l app=school-service`

## License

See LICENSE file in the root directory.
