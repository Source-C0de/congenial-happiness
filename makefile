# ==============================================================================
# Go Project ContactHub Backend
# Author: Fahmim Mohammod Shahriar
# Date: 2026-04-11
# Description: SOTA Makefile for Go projects with linting, testing, and builds.
# ==============================================================================

# --------------------------
# Configuration
# --------------------------
# Binary name
BINARY_NAME ?= contacthub
# Main package path (usually ./cmd/app or just .)
MAIN_PATH ?= ./cmd/server/main.go
# Version handling (defaults to git tag or short commit)
VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
BUILD_TIME ?= $(shell date -u '+%Y-%m-%dT%H:%M:%SZ')
GIT_COMMIT ?= $(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")

# Build flags
LDFLAGS := -ldflags "-s -w -X main.Version=$(VERSION) -X main.BuildTime=$(BUILD_TIME) -X main.GitCommit=$(GIT_COMMIT)"

# Directories
BUILD_DIR := build
DIST_DIR := dist
COVER_OUT := coverage.out

# Go commands
GO := go
GOFMT := gofmt
GOLINT := golangci-lint # Recommended linter
GOVULNCHECK := govulncheck

# OS/Arch for cross-compilation (default to current)
OS ?= $(shell uname -s | tr '[:upper:]' '[:lower:]')
ARCH ?= $(shell uname -m)
# Fix common arch names for Go
ifeq ($(ARCH),x86_64)
	ARCH := amd64
endif
ifeq ($(ARCH),aarch64)
	ARCH := arm64
endif

# --------------------------
# Phony Targets
# --------------------------
.PHONY: all build clean test lint format vet vuln-check run install-deps help docker-build

# --------------------------
# Default Target
# --------------------------
all: lint test build

# --------------------------
# Help System
# --------------------------
help: ## 📚 Show this help message
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

# --------------------------
# Development Workflow
# --------------------------
run: ## 🚀 Run the application locally
	@echo "🚀 Running $(BINARY_NAME)..."
	$(GO) run $(MAIN_PATH)

debug: ## 🐛 Run with debug flags (no optimization)
	@echo "🐛 Running in debug mode..."
	$(GO) run -gcflags="all=-N -l" $(MAIN_PATH)

install-deps: ## 📦 Install development dependencies (linters, etc.)
	@echo "📦 Installing tools..."
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install golang.org/x/vuln/cmd/govulncheck@latest
	@echo "✅ Tools installed."

# --------------------------
# Code Quality
# --------------------------
format: ## 🎨 Format code using gofmt
	@echo "🎨 Formatting code..."
	$(GOFMT) -w .

lint: ## 🔍 Lint code using golangci-lint
	@echo "🔍 Linting code..."
	@if command -v $(GOLINT) &> /dev/null; then \
		$(GOLINT) run ./...; \
	else \
		echo "⚠️  golangci-lint not found. Run 'make install-deps'"; \
		exit 1; \
	fi

vet: ## 👁️  Run go vet (static analysis)
	@echo "👁️  Running go vet..."
	$(GO) vet ./...

vuln-check: ## 🛡️  Check for vulnerabilities
	@echo "🛡️  Checking for vulnerabilities..."
	@if command -v $(GOVULNCHECK) &> /dev/null; then \
		$(GOVULNCHECK) ./...; \
	else \
		echo "⚠️  govulncheck not found. Run 'make install-deps'"; \
		exit 1; \
	fi

# --------------------------
# Testing
# --------------------------
test: ## 🧪 Run unit tests
	@echo "🧪 Running tests..."
	$(GO) test -v -race -count=1 ./...

test-cover: ## 📊 Run tests with coverage report
	@echo "📊 Running tests with coverage..."
	$(GO) test -v -race -coverprofile=$(COVER_OUT) -covermode=atomic ./...
	@echo "✅ Coverage report generated: $(COVER_OUT)"
	@$(GO) tool cover -html=$(COVER_OUT) -o coverage.html
	@echo "🌐 Open coverage.html to view details."

# --------------------------
# Building
# --------------------------
build: ## 🏗️  Build the binary for current OS/Arch
	@echo "🏗️  Building $(BINARY_NAME) v$(VERSION) for $(OS)/$(ARCH)..."
	@mkdir -p $(BUILD_DIR)
	CGO_ENABLED=0 GOOS=$(OS) GOARCH=$(ARCH) $(GO) build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PATH)
	@echo "✅ Binary created: $(BUILD_DIR)/$(BINARY_NAME)"

release: ## 📦 Build binaries for multiple platforms (Linux, macOS, Windows)
	@echo "📦 Building release binaries..."
	@mkdir -p $(DIST_DIR)
	# Linux AMD64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GO) build $(LDFLAGS) -o $(DIST_DIR)/$(BINARY_NAME)-linux-amd64 $(MAIN_PATH)
	# macOS AMD64
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GO) build $(LDFLAGS) -o $(DIST_DIR)/$(BINARY_NAME)-darwin-amd64 $(MAIN_PATH)
	# macOS ARM64 (M1/M2)
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 $(GO) build $(LDFLAGS) -o $(DIST_DIR)/$(BINARY_NAME)-darwin-arm64 $(MAIN_PATH)
	# Windows AMD64
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GO) build $(LDFLAGS) -o $(DIST_DIR)/$(BINARY_NAME)-windows-amd64.exe $(MAIN_PATH)
	@echo "✅ Release binaries created in $(DIST_DIR)/"

# --------------------------
# Docker
# --------------------------
docker-build: ## 🐳 Build Docker image
	@echo "🐳 Building Docker image..."
	docker build -t $(BINARY_NAME):$(VERSION) .

docker-run: ## 🐳 Run Docker container
	@echo "🐳 Running Docker container..."
	docker run --rm -p 8080:8080 $(BINARY_NAME):$(VERSION)

# --------------------------
# Cleanup
# --------------------------
clean: ## 🧹 Clean build artifacts
	@echo "🧹 Cleaning..."
	rm -rf $(BUILD_DIR) $(DIST_DIR) $(COVER_OUT) coverage.html
	@echo "✅ Clean complete."

# --------------------------
# CI/CD Helper (Used by GitHub Actions/GitLab CI)
# --------------------------
ci: lint vet vuln-check test build ## 🤖 Run all CI checks
	@echo "✅ CI pipeline passed."


# Claude 
tlaude: claude --dangerously-skip-permissions
