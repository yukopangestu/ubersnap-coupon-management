# Build stage
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Download dependencies first (better caching)
COPY go.mod ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
# -o main: output file name
# ./cmd/api: entry point
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/api

# Final stage
FROM alpine:latest

WORKDIR /root/

# Copy binary from builder
COPY --from=builder /app/main .

# Expose port
EXPOSE 8080

# Command to run
CMD ["./main"]
