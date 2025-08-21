# [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-08-21
# Multi-stage Dockerfile for Achievements Service - Kubernetes Ready
FROM golang:1.21-alpine AS builder

WORKDIR /app
RUN apk add --no-cache git
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags='-w -s -extldflags "-static"' \
    -a -installsuffix cgo \
    -o achievements-service ./services/achievements

FROM alpine:3.18
RUN apk --no-cache add ca-certificates tzdata wget && \
    addgroup -g 1001 appgroup && \
    adduser -u 1001 -G appgroup -s /bin/sh -D appuser

WORKDIR /app
COPY --from=builder /app/achievements-service .
RUN chown -R appuser:appgroup /app
USER appuser

HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:8085/health || exit 1

EXPOSE 8085
CMD ["./achievements-service"]
