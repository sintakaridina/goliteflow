# Makefile for GoliteFlow

# Variables
VERSION ?= 1.0.0
BUILD_TIME := $(shell date -u '+%Y-%m-%d_%H:%M:%S')
GIT_COMMIT := $(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")
GO_VERSION := $(shell go version | awk '{print $$3}')
LDFLAGS := -s -w -X main.Version=$(VERSION) -X main.BuildTime=$(BUILD_TIME) -X main.GitCommit=$(GIT_COMMIT) -X main.GoVersion=$(GO_VERSION)

# Build directory
BUILD_DIR := build

# Default target
.PHONY: all
all: clean test build

# Help target
.PHONY: help
help:
	@echo "GoliteFlow Makefile"
	@echo ""
	@echo "Available targets:"
	@echo "  build       - Build the application"
	@echo "  test        - Run tests"
	@echo "  lint        - Run linter"
	@echo "  clean       - Clean build artifacts"
	@echo "  release     - Build release binaries"
	@echo "  install     - Install to GOPATH/bin"
	@echo "  docker      - Build Docker image"
	@echo "  help        - Show this help"
	@echo ""
	@echo "Variables:"
	@echo "  VERSION     - Version to build (default: 1.0.0)"

# Build the application
.PHONY: build
build:
	@echo "Building GoliteFlow v$(VERSION)..."
	@echo "Build Time: $(BUILD_TIME)"
	@echo "Git Commit: $(GIT_COMMIT)"
	@echo "Go Version: $(GO_VERSION)"
	@go build -ldflags "$(LDFLAGS)" -o goliteflow ./cmd/goliteflow
	@echo "✓ Built goliteflow"

# Run tests
.PHONY: test
test:
	@echo "Running tests..."
	@go test -v -race -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "✓ Tests completed"

# Run linter
.PHONY: lint
lint:
	@echo "Running linter..."
	@golangci-lint run
	@echo "✓ Linting completed"

# Clean build artifacts
.PHONY: clean
clean:
	@echo "Cleaning build artifacts..."
	@rm -rf $(BUILD_DIR)
	@rm -f goliteflow
	@rm -f coverage.out coverage.html
	@echo "✓ Cleaned"

# Build release binaries
.PHONY: release
release: clean
	@echo "Building release binaries for v$(VERSION)..."
	@mkdir -p $(BUILD_DIR)
	@$(MAKE) build-release PLATFORM=linux/amd64
	@$(MAKE) build-release PLATFORM=linux/arm64
	@$(MAKE) build-release PLATFORM=windows/amd64
	@$(MAKE) build-release PLATFORM=windows/arm64
	@$(MAKE) build-release PLATFORM=darwin/amd64
	@$(MAKE) build-release PLATFORM=darwin/arm64
	@$(MAKE) create-checksums
	@$(MAKE) create-release-notes
	@echo "✓ Release build completed"

# Build for specific platform
.PHONY: build-release
build-release:
	@echo "Building for $(PLATFORM)..."
	@$(eval OS := $(word 1,$(subst /, ,$(PLATFORM))))
	@$(eval ARCH := $(word 2,$(subst /, ,$(PLATFORM))))
	@mkdir -p $(BUILD_DIR)/goliteflow_$(VERSION)_$(OS)_$(ARCH)
	@GOOS=$(OS) GOARCH=$(ARCH) CGO_ENABLED=0 go build -ldflags "$(LDFLAGS)" -o $(BUILD_DIR)/goliteflow_$(VERSION)_$(OS)_$(ARCH)/goliteflow$(if $(filter windows,$(OS)),.exe,) ./cmd/goliteflow
	@cd $(BUILD_DIR) && tar -czf goliteflow_$(VERSION)_$(OS)_$(ARCH).tar.gz goliteflow_$(VERSION)_$(OS)_$(ARCH)/
	@echo "✓ Built for $(PLATFORM)"

# Create checksums
.PHONY: create-checksums
create-checksums:
	@echo "Creating checksums..."
	@cd $(BUILD_DIR) && sha256sum *.tar.gz *.zip > goliteflow_$(VERSION)_checksums.txt
	@echo "✓ Checksums created"

# Create release notes
.PHONY: create-release-notes
create-release-notes:
	@echo "Creating release notes..."
	@echo "# GoliteFlow v$(VERSION)" > $(BUILD_DIR)/RELEASE_NOTES.md
	@echo "" >> $(BUILD_DIR)/RELEASE_NOTES.md
	@echo "## Downloads" >> $(BUILD_DIR)/RELEASE_NOTES.md
	@echo "" >> $(BUILD_DIR)/RELEASE_NOTES.md
	@echo "### Linux" >> $(BUILD_DIR)/RELEASE_NOTES.md
	@echo "- **AMD64**: [goliteflow_$(VERSION)_linux_amd64.tar.gz](goliteflow_$(VERSION)_linux_amd64.tar.gz)" >> $(BUILD_DIR)/RELEASE_NOTES.md
	@echo "- **ARM64**: [goliteflow_$(VERSION)_linux_arm64.tar.gz](goliteflow_$(VERSION)_linux_arm64.tar.gz)" >> $(BUILD_DIR)/RELEASE_NOTES.md
	@echo "" >> $(BUILD_DIR)/RELEASE_NOTES.md
	@echo "### Windows" >> $(BUILD_DIR)/RELEASE_NOTES.md
	@echo "- **AMD64**: [goliteflow_$(VERSION)_windows_amd64.zip](goliteflow_$(VERSION)_windows_amd64.zip)" >> $(BUILD_DIR)/RELEASE_NOTES.md
	@echo "- **ARM64**: [goliteflow_$(VERSION)_windows_arm64.zip](goliteflow_$(VERSION)_windows_arm64.zip)" >> $(BUILD_DIR)/RELEASE_NOTES.md
	@echo "" >> $(BUILD_DIR)/RELEASE_NOTES.md
	@echo "### macOS" >> $(BUILD_DIR)/RELEASE_NOTES.md
	@echo "- **AMD64**: [goliteflow_$(VERSION)_darwin_amd64.tar.gz](goliteflow_$(VERSION)_darwin_amd64.tar.gz)" >> $(BUILD_DIR)/RELEASE_NOTES.md
	@echo "- **ARM64**: [goliteflow_$(VERSION)_darwin_arm64.tar.gz](goliteflow_$(VERSION)_darwin_arm64.tar.gz)" >> $(BUILD_DIR)/RELEASE_NOTES.md
	@echo "✓ Release notes created"

# Build release binaries for README download links
.PHONY: release-binaries
release-binaries: clean
	@echo "Building release binaries for README download links..."
	@mkdir -p releases
	@echo "Building goliteflow-linux-amd64..."
	@GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "-s -w" -o releases/goliteflow-linux-amd64 ./cmd/goliteflow
	@echo "Building goliteflow-windows-amd64.exe..."
	@GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "-s -w" -o releases/goliteflow-windows-amd64.exe ./cmd/goliteflow
	@echo "Building goliteflow-darwin-arm64..."
	@GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 go build -ldflags "-s -w" -o releases/goliteflow-darwin-arm64 ./cmd/goliteflow
	@chmod +x releases/goliteflow-linux-amd64 releases/goliteflow-darwin-arm64
	@echo "✓ Release binaries created in releases/ directory"
	@ls -la releases/

# Install to GOPATH/bin
.PHONY: install
install: build
	@echo "Installing goliteflow..."
	@go install -ldflags "$(LDFLAGS)" ./cmd/goliteflow
	@echo "✓ Installed to $(GOPATH)/bin/goliteflow"

# Build Docker image
.PHONY: docker
docker:
	@echo "Building Docker image..."
	@docker build -t sintakaridina/goliteflow:$(VERSION) .
	@docker tag sintakaridina/goliteflow:$(VERSION) sintakaridina/goliteflow:latest
	@echo "✓ Docker image built"

# Format code
.PHONY: fmt
fmt:
	@echo "Formatting code..."
	@go fmt ./...
	@echo "✓ Code formatted"

# Vet code
.PHONY: vet
vet:
	@echo "Vetting code..."
	@go vet ./...
	@echo "✓ Code vetted"

# Run all checks
.PHONY: check
check: fmt vet lint test
	@echo "✓ All checks passed"

# Create tag
.PHONY: tag
tag:
	@echo "Creating tag v$(VERSION)..."
	@git tag -a v$(VERSION) -m "Release v$(VERSION)"
	@echo "✓ Tag v$(VERSION) created"

# Push tag
.PHONY: push-tag
push-tag:
	@echo "Pushing tag v$(VERSION)..."
	@git push origin v$(VERSION)
	@echo "✓ Tag v$(VERSION) pushed"

# Full release process
.PHONY: full-release
full-release: check release tag push-tag
	@echo "✓ Full release process completed for v$(VERSION)"
