# [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-09-24
# Multi-stage Dockerfile for School Management Microservice - Kubernetes Optimized

# Build stage
FROM golang:1.20-alpine AS builder

# Set working directory
WORKDIR /app

# Install build dependencies and security updates
RUN apk add --no-cache git ca-certificates tzdata \
    && apk upgrade --no-cache

# Create non-root user for build process
RUN addgroup -g 1001 -S buildgroup && \
    adduser -u 1001 -S builduser -G buildgroup

# Copy go mod files first for better layer caching
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download && go mod verify

# Copy source code
COPY --chown=builduser:buildgroup . .

# Switch to build user
USER builduser

# Build the application with optimizations
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -a -installsuffix cgo \
    -ldflags='-w -s -extldflags "-static"' \
    -o school-service ./cmd/main.go

# Final stage - Distroless for maximum security
FROM gcr.io/distroless/static-debian11:nonroot

# Copy CA certificates and timezone data
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

# Set working directory
WORKDIR /app

# Copy built binary from builder stage
COPY --from=builder /app/school-service .

# Copy configuration template
COPY --from=builder /app/.env.example .env.example

# Use nonroot user (65532:65532)
USER nonroot:nonroot

# Expose port
EXPOSE 8080

# Health check (note: distroless doesn't have wget/curl, health check will be done via Kubernetes probes)
# HEALTHCHECK is not needed in Kubernetes as we'll use readiness/liveness probes

# Command to run the application
ENTRYPOINT ["./school-service"]