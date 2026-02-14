.PHONY: build run test clean help

BINARY_NAME=go-passman
BINARY_UNIX=$(BINARY_NAME)
BINARY_WIN=$(BINARY_NAME).exe

help:
	@echo "Available targets:"
	@echo "  build       - Build the application"
	@echo "  run         - Run the application"
	@echo "  test        - Run tests"
	@echo "  clean       - Clean build artifacts"
	@echo "  fmt         - Format code"
	@echo "  lint        - Run linter"

LDFLAGS=-s -w

build:
	@echo "Building $(BINARY_NAME) (slim)..."
	@go build -ldflags="$(LDFLAGS)" -trimpath -o $(BINARY_WIN)
	@echo "Build complete: $(BINARY_WIN)"

build-unix:
	@echo "Building $(BINARY_UNIX) (slim)..."
	@go build -ldflags="$(LDFLAGS)" -trimpath -o $(BINARY_UNIX)
	@echo "Build complete: $(BINARY_UNIX)"

run: build
	@./$(BINARY_WIN) --help

test:
	@go test -v ./...

clean:
	@echo "Cleaning..."
	@go clean
	@rm -f $(BINARY_NAME) $(BINARY_WIN)
	@echo "Clean complete"

fmt:
	@go fmt ./...

lint:
	@golangci-lint run ./...
