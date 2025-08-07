# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go mod and source files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Set the working directory to the project root
WORKDIR /app

# Build the application using the working main.go file
RUN go build -o /app/main ./services/academic-service

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates curl

WORKDIR /root/

# Copy the binary from builder stage
COPY --from=builder /app/main .

# Expose port
EXPOSE 8083

# Command to run
CMD ["./main"]
