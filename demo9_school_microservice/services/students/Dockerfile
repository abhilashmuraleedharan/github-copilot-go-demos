# [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-08-21
# Multi-stage Dockerfile for Students Service - Kubernetes Ready
FROM golang:1.21-alpine AS builder

# Set working directory
WORKDIR /app

# Install git for private repositories (if needed)
RUN apk add --no-cache git

# Copy go mod files first for better caching
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the students service with optimizations
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags='-w -s -extldflags "-static"' \
    -a -installsuffix cgo \
    -o students-service ./services/students

# Final stage - minimal image
FROM alpine:3.18

# Install ca-certificates for HTTPS calls, timezone data, and wget for health checks
RUN apk --no-cache add ca-certificates tzdata wget && \
    addgroup -g 1001 appgroup && \
    adduser -u 1001 -G appgroup -s /bin/sh -D appuser

# Set working directory
WORKDIR /app

# Copy the binary from builder stage
COPY --from=builder /app/students-service .

# Change ownership to non-root user
RUN chown -R appuser:appgroup /app

# Switch to non-root user for security
USER appuser

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:8081/health || exit 1

# Expose port
EXPOSE 8081

# Run the application
CMD ["./students-service"]
