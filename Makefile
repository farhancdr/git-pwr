# Project variables
BINARY_NAME = git-pwr
MAIN_PACKAGE = ./main.go
BUILD_DIR = build

# Default target
all: build

# Build the binary
build:
	go build -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PACKAGE)

# Run the CLI
run:
	go run $(MAIN_PACKAGE)

# Install the binary to your GOPATH/bin or GOBIN
install:
	mv $(BUILD_DIR)/$(BINARY_NAME) $(GOPATH)/bin/$(BINARY_NAME) || mv $(BUILD_DIR)/$(BINARY_NAME) $(GOBIN)/$(BINARY_NAME)

# Run tests (if any)
test:
	go test ./...

# Format code
fmt:
	go fmt ./...

# Clean build artifacts
clean:
	rm -rf $(BUILD_DIR)

# Rebuild from scratch
rebuild: clean build

.PHONY: all build run install test fmt clean rebuild
