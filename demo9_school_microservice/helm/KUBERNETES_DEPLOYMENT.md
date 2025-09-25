# [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-09-24
# Kubernetes Deployment Guide

This guide covers deploying the School Management Microservice to a Kubernetes cluster using Helm charts.

## Prerequisites

- **Kubernetes Cluster**: Version 1.18+ with RBAC enabled
- **Helm**: Version 3.0+
- **kubectl**: Configured to connect to your cluster
- **Docker Registry**: Access to push/pull container images

## Quick Deployment

### 1. Build and Push Docker Image

```bash
# Build the optimized Kubernetes image
docker build -t your-registry/school-microservice:1.0.0 .

# Push to your registry
docker push your-registry/school-microservice:1.0.0
```

### 2. Deploy with Helm

```bash
# Install the Helm chart
helm install school-microservice ./helm/school-microservice \
  --namespace school-demo \
  --create-namespace \
  --set image.repository=your-registry/school-microservice \
  --set image.tag=1.0.0 \
  --set secrets.couchbase.password=your-secure-password

# Check deployment status
kubectl get pods -n school-demo
kubectl get services -n school-demo
```

### 3. Access the Service

```bash
# Port forward for local access
kubectl port-forward -n school-demo svc/school-microservice 8080:80

# Test the service
curl http://localhost:8080/health
```

## Configuration Options

### Environment-specific Values

Create environment-specific values files:

**values-dev.yaml:**
```yaml
replicaCount: 1
resources:
  limits:
    cpu: 200m
    memory: 256Mi
  requests:
    cpu: 100m
    memory: 128Mi
autoscaling:
  enabled: false
ingress:
  hosts:
    - host: school-api-dev.example.com
```

**values-prod.yaml:**
```yaml
replicaCount: 3
resources:
  limits:
    cpu: 1
    memory: 1Gi
  requests:
    cpu: 500m
    memory: 512Mi
autoscaling:
  enabled: true
  minReplicas: 3
  maxReplicas: 20
ingress:
  hosts:
    - host: school-api.example.com
```

### Deploy to Different Environments

```bash
# Development
helm install school-dev ./helm/school-microservice \
  -f helm/school-microservice/values-dev.yaml \
  --namespace school-dev

# Production
helm install school-prod ./helm/school-microservice \
  -f helm/school-microservice/values-prod.yaml \
  --namespace school-prod
```

## Security Features

The Helm chart includes enterprise-grade security features:

- **Distroless base image** for minimal attack surface
- **Non-root user execution** (UID 65532)
- **Read-only root filesystem**
- **Security contexts** with dropped capabilities
- **Network policies** for traffic isolation
- **Pod security standards** compliance
- **Secrets management** for sensitive data

## High Availability

The deployment supports high availability through:

- **Multiple replicas** with anti-affinity rules
- **Pod disruption budgets** to ensure availability during updates
- **Horizontal Pod Autoscaling** based on CPU/memory metrics
- **Health checks** with proper liveness and readiness probes
- **Rolling updates** with zero downtime

## Monitoring Integration

Built-in support for monitoring:

```yaml
monitoring:
  enabled: true
  serviceMonitor:
    enabled: true
  prometheusRule:
    enabled: true
```

## Ingress Configuration

Configure ingress for external access:

```yaml
ingress:
  enabled: true
  className: nginx
  hosts:
    - host: school-api.example.com
      paths:
        - path: /
          pathType: Prefix
  tls:
    - secretName: school-microservice-tls
      hosts:
        - school-api.example.com
```

## Database Options

### Option 1: Bundled Couchbase (Development)
```yaml
couchbase:
  enabled: true
  persistence:
    enabled: true
    size: 10Gi
```

### Option 2: External Couchbase (Production)
```yaml
couchbase:
  enabled: false
config:
  couchbase:
    host: "external-couchbase.example.com"
    port: "8091"
```

## Upgrading

```bash
# Upgrade to new version
helm upgrade school-microservice ./helm/school-microservice \
  --set image.tag=1.1.0 \
  --namespace school-demo

# Check rollout status
kubectl rollout status deployment/school-microservice -n school-demo
```

## Troubleshooting

### Common Issues

1. **Image Pull Errors**
   ```bash
   kubectl describe pod <pod-name> -n school-demo
   ```

2. **Database Connection Issues**
   ```bash
   kubectl logs <pod-name> -n school-demo
   kubectl exec -it <pod-name> -n school-demo -- /bin/sh
   ```

3. **Service Discovery Issues**
   ```bash
   kubectl get endpoints -n school-demo
   kubectl get services -n school-demo
   ```

### Scaling Operations

```bash
# Manual scaling
kubectl scale deployment school-microservice --replicas=5 -n school-demo

# Check HPA status
kubectl get hpa -n school-demo
kubectl describe hpa school-microservice -n school-demo
```

## Cleanup

```bash
# Uninstall the release
helm uninstall school-microservice --namespace school-demo

# Delete namespace (if desired)
kubectl delete namespace school-demo
```