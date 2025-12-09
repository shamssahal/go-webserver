.PHONY: build run test clean help docker-build docker-run docker-stop docker-clean

# Binary name
BINARY_NAME=go-server
MAIN_PATH=./cmd/api
DOCKER_IMAGE=go-server
DOCKER_TAG=latest

# Build the application
build:
	@echo "Building..."
	@go build -o bin/$(BINARY_NAME) $(MAIN_PATH)
	@echo "Build complete: bin/$(BINARY_NAME)"

# Run the application (depends on build)
run: build
	@echo "Running..."
	@./bin/$(BINARY_NAME)

# Run tests
test:
	@echo "Running tests..."
	@go test -v ./...

# Clean build artifacts
clean:
	@echo "Cleaning..."
	@rm -rf bin/
	@go clean
	@echo "Clean complete"

# Docker build
docker-build:
	@echo "Building Docker image..."
	@docker build -t $(DOCKER_IMAGE):$(DOCKER_TAG) .
	@echo "Docker image built: $(DOCKER_IMAGE):$(DOCKER_TAG)"

# Docker run with docker-compose (always fresh build, no cache)
docker-run:
	@echo "Building fresh image (no cache)..."
	@docker compose build --no-cache
	@echo "Starting Docker container..."
	@docker compose up -d
	@echo "Container started. Access at http://localhost:3000"

# Docker stop
docker-stop:
	@echo "Stopping Docker container..."
	@docker compose down
	@echo "Container stopped"

# Docker clean (remove image)
docker-clean:
	@echo "Removing Docker image..."
	@docker rmi $(DOCKER_IMAGE):$(DOCKER_TAG) || true
	@echo "Docker image removed"

# Display help
help:
	@echo "Available targets:"
	@echo "  build         - Build the application binary"
	@echo "  run           - Run the application"
	@echo "  test          - Run all tests"
	@echo "  clean         - Remove build artifacts"
	@echo "  docker-build  - Build Docker image"
	@echo "  docker-run    - Run Docker container with docker-compose"
	@echo "  docker-stop   - Stop Docker container"
	@echo "  docker-clean  - Remove Docker image"
	@echo "  help          - Display this help message"
