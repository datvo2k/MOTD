# Variables
BINARY_NAME=motd
GO=go

# Default target
all: build

# Build the binary
build:
	$(GO) build -o $(BINARY_NAME)

# Clean build artifacts
clean:
	rm -f $(BINARY_NAME)
	go clean

# Run the application
run: build
	./$(BINARY_NAME)

# Build for multiple platforms
cross-build:
	GOOS=linux GOARCH=amd64 $(GO) build -o $(BINARY_NAME)_linux_amd64
	GOOS=windows GOARCH=amd64 $(GO) build -o $(BINARY_NAME)_windows_amd64.exe
	GOOS=darwin GOARCH=amd64 $(GO) build -o $(BINARY_NAME)_darwin_amd64

.PHONY: all build clean run cross-build