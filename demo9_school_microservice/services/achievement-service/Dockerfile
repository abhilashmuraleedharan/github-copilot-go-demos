# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum first for better Docker cache usage
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project
COPY . .

# Set the working directory to the project root
WORKDIR /app

# Build the application using the working main.go file
RUN go build -o /app/main ./services/achievement-service

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates curl

WORKDIR /root/

# Copy the binary from builder stage
COPY --from=builder /app/main .

# Expose port
EXPOSE 8084

# Command to run
CMD ["./main"]
