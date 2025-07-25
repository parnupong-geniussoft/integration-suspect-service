# Use official Golang image as builder
FROM golang:alpine AS builder

WORKDIR /app

# Copy source code
COPY . .

# Download dependencies
RUN go mod download

# Build Go binary
RUN go build -o integration-suspect-service ./app/main.go

# Use minimal image for runtime
FROM alpine:latest

WORKDIR /app

# Copy binary and .env to /app
COPY --from=builder /app/integration-suspect-service .
COPY .env .

# Expose port your app uses (e.g. 8081)
EXPOSE 8081

# Run binary from /app
CMD ["./integration-suspect-service"]
