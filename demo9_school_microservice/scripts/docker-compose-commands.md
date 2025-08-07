# Docker Compose Commands

This document provides comprehensive Docker Compose commands for managing the School Management System.

## Quick Commands

### Start All Services
```bash
# Start all services in the background
docker-compose up -d

# Start all services with logs visible
docker-compose up

# Start specific services only
docker-compose up -d couchbase api-gateway student-service
```

### Stop Services
```bash
# Stop all services
docker-compose down

# Stop services but keep volumes
docker-compose stop

# Stop and remove all containers, networks, and volumes
docker-compose down -v

# Stop and remove everything including images
docker-compose down -v --rmi all
```

### Build and Deploy
```bash
# Build all images and start services
docker-compose up -d --build

# Build specific service
docker-compose build student-service

# Force rebuild without cache
docker-compose build --no-cache

# Pull latest images and start
docker-compose pull && docker-compose up -d
```

### Service Management
```bash
# View running services
docker-compose ps

# View logs for all services
docker-compose logs

# Follow logs for specific service
docker-compose logs -f api-gateway

# View logs for multiple services
docker-compose logs student-service teacher-service

# Restart specific service
docker-compose restart student-service

# Scale specific service (if supported)
docker-compose up -d --scale student-service=3
```

### Health Checks and Monitoring
```bash
# Check service health status
docker-compose ps --format "table {{.Name}}\t{{.Status}}\t{{.Ports}}"

# Execute command in running container
docker-compose exec api-gateway sh

# Execute one-off command
docker-compose run --rm student-service sh

# View resource usage
docker stats $(docker-compose ps -q)
```

### Database Operations
```bash
# Start only Couchbase
docker-compose up -d couchbase

# Wait for Couchbase to be ready
docker-compose exec couchbase couchbase-cli server-info -c localhost -u Administrator -p password

# Access Couchbase container
docker-compose exec couchbase bash
```

## Development Workflow

### Complete Development Setup
```bash
# 1. Clean start
docker-compose down -v
docker system prune -f

# 2. Build and start all services
docker-compose up -d --build

# 3. Wait for services to be ready (check health)
./scripts/wait-for-services.sh

# 4. Load sample data
./scripts/load-sample-data.sh

# 5. Run tests
./scripts/run-tests.sh
```

### Debugging Workflow
```bash
# View logs with timestamps
docker-compose logs -t

# Follow logs for troubleshooting
docker-compose logs -f --tail=100

# Inspect specific container
docker-compose exec student-service env
docker-compose exec student-service ps aux
docker-compose exec student-service netstat -tlnp
```

### Production Deployment
```bash
# Production deployment with resource limits
docker-compose -f docker-compose.yml -f docker-compose.prod.yml up -d

# Health check before deployment
docker-compose config --quiet && echo "Configuration is valid"

# Rolling update (zero downtime)
docker-compose up -d --no-deps --build student-service
```

## Environment-Specific Commands

### Windows PowerShell
```powershell
# Start services
docker-compose up -d

# View logs
docker-compose logs

# Stop services
docker-compose down

# Clean everything
docker-compose down -v; docker system prune -f
```

### Linux/macOS
```bash
# Start with resource monitoring
docker-compose up -d && watch docker stats $(docker-compose ps -q)

# Complete cleanup
docker-compose down -v && docker system prune -af

# Start with custom environment file
docker-compose --env-file .env.production up -d
```

## Troubleshooting Commands

### Service Issues
```bash
# Check if ports are available
netstat -an | grep :8080
ss -tlnp | grep :8080  # Linux

# Restart problematic service
docker-compose restart student-service

# Recreate specific service
docker-compose up -d --force-recreate student-service

# Check service dependencies
docker-compose config --services
```

### Network Issues
```bash
# Inspect network
docker network ls
docker network inspect schoolmgmt-network

# Test connectivity between services
docker-compose exec api-gateway ping student-service
docker-compose exec api-gateway nslookup student-service
```

### Volume Issues
```bash
# List volumes
docker volume ls

# Inspect volume
docker volume inspect schoolmgmt_couchbase_data

# Backup volume
docker run --rm -v schoolmgmt_couchbase_data:/data -v $(pwd):/backup alpine tar czf /backup/couchbase-backup.tar.gz /data

# Restore volume
docker run --rm -v schoolmgmt_couchbase_data:/data -v $(pwd):/backup alpine tar xzf /backup/couchbase-backup.tar.gz -C /
```

### Performance Monitoring
```bash
# Monitor resource usage
docker-compose exec api-gateway top
docker-compose exec api-gateway free -h
docker-compose exec api-gateway df -h

# Monitor logs for errors
docker-compose logs | grep -i error
docker-compose logs | grep -i "connection refused"
```

## Advanced Commands

### Multi-Environment Management
```bash
# Development environment
docker-compose -f docker-compose.yml -f docker-compose.dev.yml up -d

# Testing environment
docker-compose -f docker-compose.yml -f docker-compose.test.yml up -d

# Production environment
docker-compose -f docker-compose.yml -f docker-compose.prod.yml up -d
```

### CI/CD Commands
```bash
# Build and test
docker-compose -f docker-compose.yml -f docker-compose.test.yml up --build --abort-on-container-exit

# Health check validation
timeout 300 bash -c 'until docker-compose exec api-gateway curl -f http://localhost:8080/health; do sleep 5; done'

# Automated deployment
docker-compose pull && docker-compose up -d --remove-orphans
```

### Backup and Recovery
```bash
# Full system backup
docker-compose exec couchbase cbbackup http://localhost:8091 /backup -u Administrator -p password

# Export configuration
docker-compose config > docker-compose-backup.yml

# Export environment
docker-compose exec api-gateway env > api-gateway.env
```

## Quick Reference

| Command | Description |
|---------|-------------|
| `docker-compose up -d` | Start all services in background |
| `docker-compose down` | Stop all services |
| `docker-compose logs -f SERVICE` | Follow logs for specific service |
| `docker-compose restart SERVICE` | Restart specific service |
| `docker-compose exec SERVICE sh` | Access service container |
| `docker-compose ps` | List running services |
| `docker-compose build --no-cache` | Rebuild all images |
| `docker-compose down -v` | Stop and remove volumes |

## Environment Variables

Create a `.env` file in the project root:

```env
# Service Ports
API_GATEWAY_PORT=8080
STUDENT_SERVICE_PORT=8081
TEACHER_SERVICE_PORT=8082
ACADEMIC_SERVICE_PORT=8083
ACHIEVEMENT_SERVICE_PORT=8084

# Couchbase Configuration
COUCHBASE_ADMIN_USER=Administrator
COUCHBASE_ADMIN_PASSWORD=password
COUCHBASE_BUCKET=schoolmgmt

# Environment
ENVIRONMENT=development
LOG_LEVEL=debug
```

Use with: `docker-compose --env-file .env up -d`
