# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Make sure scripts are executable
RUN chmod +x scripts/test-api.sh

# Build the application
RUN GOOS=linux go build -o main cmd/server/main.go

# Runtime stage
FROM alpine:latest

# Install ca-certificates and curl for HTTPS calls and testing
RUN apk --no-cache add ca-certificates curl

WORKDIR /root/

# Copy the binary from builder stage
COPY --from=builder /app/main .

# Copy scripts for testing
COPY --from=builder /app/scripts ./scripts

# Expose port
EXPOSE 8080

# Run the application
CMD ["./main"]