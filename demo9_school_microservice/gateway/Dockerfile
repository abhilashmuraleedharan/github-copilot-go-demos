# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy go mod and source files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o api-gateway main.go

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates curl

WORKDIR /root/

# Copy the binary from builder stage
COPY --from=builder /app/api-gateway .

# Expose port
EXPOSE 8080

# Command to run
CMD ["./api-gateway"]
