FROM golang:1.25-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/tinypay-server .

# Create a minimal image
FROM alpine:latest

# Install ca-certificates for HTTPS requests
RUN apk --no-cache add ca-certificates

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/tinypay-server /app/tinypay-server

# Copy .env.example as reference
COPY .env.example /app/.env.example

# Expose the port (default 9090)
EXPOSE 9090

# Run the application
CMD ["/app/tinypay-server"]
