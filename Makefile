.PHONY: build run test clean help

# Binary name
BINARY_NAME=go-server
MAIN_PATH=./cmd/api

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

# Display help
help:
	@echo "Available targets:"
	@echo "  build  - Build the application binary"
	@echo "  run    - Run the application"
	@echo "  test   - Run all tests"
	@echo "  clean  - Remove build artifacts"
	@echo "  help   - Display this help message"
