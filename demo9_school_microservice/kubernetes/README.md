# [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-08-21
# School Microservice Kubernetes Deployment Guide

## Prerequisites

1. **Kubernetes Cluster** (v1.20+)
2. **Helm 3.x** installed
3. **kubectl** configured for your cluster
4. **NGINX Ingress Controller** (optional, for external access)
5. **Prometheus Operator** (optional, for monitoring)

## Quick Deployment

### 1. Deploy with Helm

```bash
# Add Couchbase Helm repository (if using Couchbase operator)
helm repo add couchbase https://couchbase-partners.github.io/helm-charts/
helm repo update

# Create namespace
kubectl create namespace school-demo

# Install the chart
helm install school-microservice ./helm/school-microservice \
  --namespace school-demo \
  --create-namespace

# Or install with custom values
helm install school-microservice ./helm/school-microservice \
  --namespace school-demo \
  --values ./helm/school-microservice/values-production.yaml
```

### 2. Verify Deployment

```bash
# Check all resources
kubectl get all -n school-demo

# Check pod status
kubectl get pods -n school-demo

# Check services
kubectl get svc -n school-demo

# Check ingress
kubectl get ingress -n school-demo
```

### 3. Access Services

```bash
# Port forward for testing (if no ingress)
kubectl port-forward svc/school-microservice-students 8081:8081 -n school-demo
kubectl port-forward svc/school-microservice-teachers 8082:8082 -n school-demo
kubectl port-forward svc/school-microservice-classes 8083:8083 -n school-demo
kubectl port-forward svc/school-microservice-academics 8084:8084 -n school-demo
kubectl port-forward svc/school-microservice-achievements 8085:8085 -n school-demo

# Test health endpoints
curl http://localhost:8081/health
curl http://localhost:8082/health
curl http://localhost:8083/health
curl http://localhost:8084/health
curl http://localhost:8085/health
```

## Configuration Options

### Environment-specific Values

Create different values files for different environments:

```bash
# values-development.yaml
helm install school-microservice ./helm/school-microservice \
  --namespace school-demo \
  --values ./helm/school-microservice/values-development.yaml

# values-staging.yaml
helm install school-microservice ./helm/school-microservice \
  --namespace school-demo \
  --values ./helm/school-microservice/values-staging.yaml

# values-production.yaml
helm install school-microservice ./helm/school-microservice \
  --namespace school-demo \
  --values ./helm/school-microservice/values-production.yaml
```

### Custom Configuration

```yaml
# Custom values example
image:
  registry: your-registry.com
  tag: "v1.0.1"

resources:
  limits:
    cpu: 1000m
    memory: 1Gi
  requests:
    cpu: 500m
    memory: 512Mi

autoscaling:
  enabled: true
  minReplicas: 3
  maxReplicas: 20

couchbase:
  cluster:
    size: 5
    resources:
      limits:
        cpu: 4
        memory: 8Gi
```

## Monitoring and Observability

### Enable Prometheus Monitoring

```yaml
# In values.yaml
serviceMonitor:
  enabled: true
  namespace: monitoring
  labels:
    prometheus: kube-prometheus
```

### View Logs

```bash
# View logs for specific service
kubectl logs -f deployment/school-microservice-students -n school-demo

# View logs for all services
kubectl logs -f -l app.kubernetes.io/name=school-microservice -n school-demo
```

## Scaling

### Manual Scaling

```bash
# Scale specific service
kubectl scale deployment school-microservice-students --replicas=5 -n school-demo

# Scale all services
kubectl scale deployment -l app.kubernetes.io/name=school-microservice --replicas=3 -n school-demo
```

### Auto Scaling

HPA is enabled by default and will scale based on CPU and memory usage.

## Updates and Rollbacks

### Rolling Updates

```bash
# Update with new image
helm upgrade school-microservice ./helm/school-microservice \
  --namespace school-demo \
  --set image.tag=v1.0.2

# Check rollout status
kubectl rollout status deployment/school-microservice-students -n school-demo
```

### Rollback

```bash
# View rollout history
kubectl rollout history deployment/school-microservice-students -n school-demo

# Rollback to previous version
kubectl rollout undo deployment/school-microservice-students -n school-demo

# Rollback to specific revision
kubectl rollout undo deployment/school-microservice-students --to-revision=2 -n school-demo
```

## Troubleshooting

### Common Issues

1. **Pods not starting**: Check resource limits and node capacity
2. **Database connection errors**: Verify Couchbase cluster is running
3. **Service discovery issues**: Check DNS and service names
4. **Ingress not working**: Verify ingress controller and annotations

### Debug Commands

```bash
# Describe problematic pods
kubectl describe pod <pod-name> -n school-demo

# Check events
kubectl get events -n school-demo --sort-by='.lastTimestamp'

# Check resource usage
kubectl top pods -n school-demo
kubectl top nodes

# Access pod shell for debugging
kubectl exec -it <pod-name> -n school-demo -- /bin/sh
```

## Security Considerations

1. **Network Policies**: Enabled by default to restrict traffic
2. **Pod Security**: Runs as non-root user with read-only filesystem
3. **Secrets Management**: Database credentials stored in Kubernetes secrets
4. **RBAC**: Minimal service account permissions
5. **Image Security**: Multi-stage builds with minimal base images

## Performance Tuning

### Resource Optimization

```yaml
resources:
  limits:
    cpu: 500m      # Adjust based on load testing
    memory: 512Mi  # Monitor memory usage patterns
  requests:
    cpu: 100m      # Start conservative
    memory: 128Mi  # Monitor startup requirements
```

### Database Optimization

```yaml
couchbase:
  buckets:
    - name: school
      memoryQuota: 2048  # Increase for better performance
      replicas: 2        # Increase for high availability
```

## Backup and Disaster Recovery

### Database Backup

```bash
# Use Couchbase backup tools
cbbackupmgr config --archive /backup --repo school-backup
cbbackupmgr backup --archive /backup --repo school-backup --cluster couchbase://school-couchbase-cluster:8091
```

### Application Backup

```bash
# Backup Helm releases
helm get values school-microservice -n school-demo > school-microservice-backup.yaml

# Backup Kubernetes resources
kubectl get all -n school-demo -o yaml > school-microservice-k8s-backup.yaml
```
