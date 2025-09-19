#!/bin/bash

# TinyPay Server Deployment Script
# This script builds and deploys the TinyPay server with Docker Compose

set -e  # Exit on any error

echo "🚀 Starting TinyPay Server deployment..."

# Check if Docker is running
if ! docker info > /dev/null 2>&1; then
    echo "❌ Docker is not running. Please start Docker and try again."
    exit 1
fi

# Check if docker-compose is available
if ! command -v docker-compose &> /dev/null; then
    echo "❌ docker-compose is not installed. Please install it and try again."
    exit 1
fi

# Stop existing services
echo "🛑 Stopping existing services..."
docker-compose down

# Build and start services
echo "🔨 Building and starting services..."
docker-compose up --build -d

# Wait for services to be ready
echo "⏳ Waiting for services to be ready..."
sleep 10

# Check service status
echo "📊 Checking service status..."
docker-compose ps

# Test health endpoint
echo "🏥 Testing health endpoint..."
if curl -f http://localhost/api/health > /dev/null 2>&1; then
    echo "✅ Health check passed!"
else
    echo "❌ Health check failed. Checking logs..."
    docker-compose logs --tail=20
    exit 1
fi

# Test OpenAPI endpoint
echo "📚 Testing OpenAPI endpoint..."
if curl -f http://localhost/openapi.yaml > /dev/null 2>&1; then
    echo "✅ OpenAPI endpoint is working!"
else
    echo "⚠️  OpenAPI endpoint test failed, but continuing..."
fi

echo "🎉 Deployment completed successfully!"
echo "📖 Access the API documentation at: http://localhost/docs"
echo "🔍 Check service logs with: docker-compose logs -f"
echo "🛑 Stop services with: docker-compose down"