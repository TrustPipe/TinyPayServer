# Makefile for TinyPay Server

.PHONY: help generate build run clean test docs

# Default target
help:
	@echo "Available targets:"
	@echo "  generate    - Generate Go code from OpenAPI specification"
	@echo "  build       - Build the server binary"
	@echo "  run         - Run the server"
	@echo "  clean       - Clean generated files and binaries"
	@echo "  test        - Run tests"
	@echo "  docs        - Open API documentation in browser"
	@echo "  install     - Install oapi-codegen tool"

# Install oapi-codegen tool
install:
	@echo "Installing oapi-codegen..."
	go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
	@echo "oapi-codegen installed successfully"

# Generate Go code from OpenAPI specification
generate:
	@echo "Generating Go code from OpenAPI specification..."
	@mkdir -p api
	oapi-codegen -package=api -generate gin api/openapi.yaml > api/server.gen.go
	oapi-codegen -package=api -generate types api/openapi.yaml > api/types.gen.go
	oapi-codegen -package=api -generate client api/openapi.yaml > api/client.gen.go
	oapi-codegen -package=api -generate spec api/openapi.yaml > api/spec.gen.go
	@echo "Code generation completed"

# Build the server
build: generate
	@echo "Building TinyPay server..."
	go build -o tinypay-server .
	@echo "Build completed: tinypay-server"

# Run the server
run: build
	@echo "Starting TinyPay server..."
	./tinypay-server

# Run the server in development mode
dev:
	@echo "Starting TinyPay server in development mode..."
	go run .

# Clean generated files and binaries
clean:
	@echo "Cleaning generated files and binaries..."
	rm -f api/server.gen.go api/types.gen.go api/client.gen.go api/spec.gen.go
	rm -f tinypay-server
	@echo "Clean completed"

# Run tests
test:
	@echo "Running tests..."
	go test ./...

# Open API documentation in browser (macOS)
docs:
	@echo "Opening API documentation..."
	@echo "Please start the server first with 'make run' or 'make dev'"
	@echo "Then visit: http://localhost:9090/docs"
	open http://localhost:9090/docs 2>/dev/null || echo "Please open http://localhost:9090/docs in your browser"

# Validate OpenAPI specification
validate:
	@echo "Validating OpenAPI specification..."
	@command -v swagger >/dev/null 2>&1 || { echo "swagger-codegen-cli not found. Install it first."; exit 1; }
	swagger validate api/openapi.yaml

# Format Go code
fmt:
	@echo "Formatting Go code..."
	go fmt ./...

# Run linter
lint:
	@echo "Running linter..."
	@command -v golangci-lint >/dev/null 2>&1 || { echo "golangci-lint not found. Install it first."; exit 1; }
	golangci-lint run

# Show project structure
structure:
	@echo "Project structure:"
	tree -I 'node_modules|.git|vendor' .

# Show OpenAPI info
info:
	@echo "TinyPay API Information:"
	@echo "  OpenAPI Spec: api/openapi.yaml"
	@echo "  Documentation: http://localhost:9090/docs"
	@echo "  Health Check: http://localhost:9090/api/health"
	@echo "  API Endpoints:"
	@echo "    POST /api/payments - Create payment"
	@echo "    GET  /api/payments/{hash} - Get transaction status"