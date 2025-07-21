.PHONY: build install clean test help

# Build the application
build:
	go build -o gitty

# Build with version info
build-version:
	./scripts/build.sh

# Install globally
install: build
	sudo cp gitty /usr/local/bin/

# Create release
release:
	@if [ -z "$(VERSION)" ]; then \
		echo "Usage: make release VERSION=v1.0.0"; \
		exit 1; \
	fi
	./scripts/release.sh $(VERSION)

# Run tests
test:
	go test ./...

# Test installation scripts
test-install:
	./scripts/test-install.sh

# Format code
fmt:
	go fmt ./...

# Lint code
lint:
	golangci-lint run

# Clean and rebuild
rebuild: clean build

# Clean build artifacts
clean:
	rm -f gitty

# Run tests
test:
	go test ./...

# Show help
help:
	@echo "Available commands:"
	@echo "  make build   - Build the gitty application"
	@echo "  make install - Build and install globally"
	@echo "  make clean   - Remove build artifacts"
	@echo "  make test    - Run tests"
	@echo "  make help    - Show this help"

# Default target
all: build 