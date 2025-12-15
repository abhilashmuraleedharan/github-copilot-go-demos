# [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-12-15
# Multi-stage build for Go application
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum* ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o school-microservice .

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from builder
COPY --from=builder /app/school-microservice .

# Expose port
EXPOSE 8080

# Run the application
CMD ["./school-microservice"]
