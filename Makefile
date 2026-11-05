# KhulnaSoft Terraform Provider Makefile
# Comprehensive build, test, and development targets

# Project configuration
BINARY_NAME := terraform-provider-khulnasoft
VERSION ?= 0.9.0
NAMESPACE := khulnasoft
HOSTNAME := reegregistry.terraform.io
OS_ARCH := $(shell go env GOOS)_$(shell go env GOARCH)
PLUGIN_DIR := ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${BINARY_NAME}/${VERSION}/${OS_ARCH}

# Build configuration
LD_FLAGS := "-X main.version=v${VERSION} -X github.com/khulnasoft/terraform-provider-khulnasoft/client.version=v${VERSION}"
MAIN_PACKAGE := .
PACKAGES := ./...

# Tools
GOLANGCI_LINT ?= golangci-lint
TFPLUGINDOCS ?= tfplugindocs
GORELEASER ?= goreleaser

# Default target
.DEFAULT_GOAL := help

# Help target
help: ## Show this help message
	@echo "KhulnaSoft Terraform Provider - Available targets:"
	@echo ""
	@echo "  all                     Default target (build)"
	@echo "  help                    Show this help message"
	@echo "  build                   Build the provider binary"
	@echo "  install                 Build and install the provider to Terraform plugins directory"
	@echo "  clean                   Clean build artifacts"
	@echo "  test                    Run unit tests"
	@echo "  testacc                 Run acceptance tests"
	@echo "  test-coverage           Run tests with coverage report"
	@echo "  fmt                     Format Go code"
	@echo "  lint                    Run linter"
	@echo "  vet                     Run go vet"
	@echo "  quality                 Run all code quality checks"
	@echo "  docs                    Generate documentation"
	@echo "  docs-serve              Serve documentation locally"
	@echo "  dev                     Setup development environment"
	@echo "  tools                   Install development tools"
	@echo "  deps                    Update dependencies"
	@echo "  info                    Show project information"
	@echo "  release                 Create a new release"
	@echo "  tag                     Create and push git tag"
	@echo "  version-bump-patch      Bump patch version"
	@echo "  version-bump-minor      Bump minor version"
	@echo "  version-bump-major      Bump major version"
	@echo "  check                   Check if required tools are installed"
	@echo "  validate                Validate the provider"
	@echo "  security                Run security checks"
	@echo "  quick                   Quick development workflow"
	@echo ""
	@echo "For more information, see README.md"

.PHONY: help build install clean test testacc test-coverage fmt lint vet docs docs-serve release tag dev watch info tools

# Build targets
build: ## Build the provider binary
	@echo "Building $(BINARY_NAME) v$(VERSION)..."
	go build -ldflags $(LD_FLAGS) -o $(BINARY_NAME) $(MAIN_PACKAGE)

build-cross: ## Build for multiple platforms (requires goreleaser)
	@echo "Building for multiple platforms..."
	$(GORELEASER) build --snapshot --clean

install: build ## Build and install the provider to Terraform plugins directory
	@echo "Installing $(BINARY_NAME) to $(PLUGIN_DIR)..."
	mkdir -p $(PLUGIN_DIR)
	mv $(BINARY_NAME) $(PLUGIN_DIR)/

# Clean targets
clean: ## Clean build artifacts and temporary files
	@echo "Cleaning build artifacts..."
	go clean
	rm -f $(BINARY_NAME)
	rm -rf dist/
	rm -rf ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${BINARY_NAME}/

clean-all: clean ## Clean everything including vendor and modules
	@echo "Cleaning everything..."
	rm -rf vendor/
	go mod tidy

# Test targets
test: ## Run unit tests
	@echo "Running unit tests..."
	go test $(PACKAGES) -v -cover

testacc: ## Run acceptance tests (requires TF_ACC=1 and credentials)
	@echo "Running acceptance tests..."
	TF_ACC=1 go test $(PACKAGES) -v $(TESTARGS) -timeout 120m

test-coverage: ## Run tests with coverage
	@echo "Running tests with coverage..."
	go test $(PACKAGES) -v -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

test-all: test testacc ## Run all tests

# Code quality targets
fmt: ## Format Go code
	@echo "Formatting Go code..."
	go fmt $(PACKAGES)

lint: ## Run linter
	@echo "Running linter..."
	$(GOLANGCI_LINT) run

vet: ## Run go vet
	@echo "Running go vet..."
	go vet $(PACKAGES)

fix: ## Auto-fix linting issues
	@echo "Auto-fixing linting issues..."
	$(GOLANGCI_LINT) run --fix

quality: fmt lint vet ## Run all code quality checks

# Documentation targets
docs: ## Generate documentation
	@echo "Generating documentation..."
	$(TFPLUGINDOCS) generate

docs-serve: ## Serve documentation locally
	@echo "Serving documentation locally..."
	$(TFPLUGINDOCS) serve

docs-check: ## Check if documentation is up to date
	@echo "Checking documentation..."
	$(TFPLUGINDOCS) validate

# Release targets
release-dry: ## Dry run of release process
	@echo "Dry run of release..."
	$(GORELEASER) release --snapshot --clean

release: ## Create a new release
	@echo "Creating release v$(VERSION)..."
	$(GORELEASER) release --clean

tag: ## Create and push a git tag
	@echo "Creating git tag v$(VERSION)..."
	git tag -a v$(VERSION) -m "Release v$(VERSION)"
	git push origin v$(VERSION)

# Development targets
dev: ## Setup development environment
	@echo "Setting up development environment..."
	go mod download
	go mod verify
	go install github.com/golangci-lint/cmd/golangci-lint@latest
	go install github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs@latest

deps: ## Install/update dependencies
	@echo "Installing dependencies..."
	go get -u ./...
	go mod tidy
	go mod vendor

tools: ## Install development tools
	@echo "Installing development tools..."
	go install honnef.co/go/tools/cmd/staticcheck@latest
	go install golang.org/x/tools/cmd/goimports@latest
	go install github.com/client9/misspell/cmd/misspell@latest

# Utility targets
info: ## Show project information
	@echo "Project Information:"
	@echo "  Binary:       $(BINARY_NAME)"
	@echo "  Version:      $(VERSION)"
	@echo "  Go Version:   $(shell go version)"
	@echo "  OS/Arch:      $(OS_ARCH)"
	@echo "  Plugin Dir:   $(PLUGIN_DIR)"
	@echo "  Packages:     $(PACKAGES)"

check: ## Check if required tools are installed
	@echo "Checking required tools..."
	@command -v go >/dev/null 2>&1 || { echo "Go is required but not installed. Aborting."; exit 1; }
	@command -v $(GOLANGCI_LINT) >/dev/null 2>&1 || echo "Warning: golangci-lint not found. Run 'make tools' to install."
	@command -v $(TFPLUGINDOCS) >/dev/null 2>&1 || echo "Warning: tfplugindocs not found. Run 'make tools' to install."
	@echo "All required tools are available!"

# Docker targets (if needed)
docker-build: ## Build Docker image for development
	@echo "Building Docker image..."
	docker build -t $(BINARY_NAME)-dev .

docker-run: ## Run Docker container with development environment
	@echo "Running Docker container..."
	docker run -it --rm -v $(PWD):/app -w /app $(BINARY_NAME)-dev bash

# CI/CD targets
ci: quality test ## Run CI pipeline (quality + unit tests)
ci-all: quality test-all ## Run full CI pipeline (quality + all tests)

# Quick development workflow
quick: fmt build test ## Quick development workflow (format, build, test)

# Watch mode (requires entr or similar tool)
watch: ## Watch for changes and run tests (requires entr)
	@echo "Watching for changes and running tests..."
	@command -v entr >/dev/null 2>&1 || { echo "entr is required for watch mode. Install with: brew install entr"; exit 1; }
	find . -name "*.go" | entr -r make test

# Security scanning
security: ## Run security checks
	@echo "Running security checks..."
	go install github.com/securecodewarrior/go-vuln-check@latest
	govulncheck ./...

# Performance profiling
profile: ## Generate performance profile
	@echo "Generating performance profile..."
	go test -cpuprofile=cpu.prof -memprofile=mem.prof -bench=. $(PACKAGES)
	go tool pprof cpu.prof
	go tool pprof mem.prof

# Update dependencies
update-deps: ## Update all dependencies
	@echo "Updating dependencies..."
	go get -u all
	go mod tidy
	go mod vendor

# Backup current state
backup: ## Create a backup of current state
	@echo "Creating backup..."
	tar -czf backup-$(shell date +%Y%m%d-%H%M%S).tar.gz --exclude='.git' --exclude='backup-*.tar.gz' .

# Restore from backup (use: make restore BACKUP=backup-20231201-1200.tar.gz)
restore: ## Restore from backup
	@echo "Restoring from $(BACKUP)..."
	tar -xzf $(BACKUP)

# Environment setup
env-setup: ## Setup environment variables for development
	@echo "Setting up environment variables..."
	@echo "export TF_LOG=DEBUG" > .env.dev
	@echo "export KHULNASOFT_URL=https://your-khulnasoft-instance.com" >> .env.dev
	@echo "export KHULNASOFT_USER=your-username" >> .env.dev
	@echo "export KHULNASOFT_PASSWORD=your-password" >> .env.dev
	@echo "Environment setup complete. Edit .env.dev with your actual values."

# Validation targets
validate: ## Validate the provider
	@echo "Validating provider..."
	go vet $(PACKAGES)
	$(GOLANGCI_LINT) run
	$(TFPLUGINDOCS) validate

validate-all: validate test ## Full validation including tests

# Legacy compatibility (call GNUmakefile targets)
gmake-%: ## Call GNUmakefile targets (e.g., make gmake-build)
	$(MAKE) -f GNUmakefile $*

# Version management
version: ## Show current version
	@echo "Current version: $(VERSION)"

version-bump-patch: ## Bump patch version (e.g., 1.0.0 -> 1.0.1)
	$(eval NEW_VERSION := $(shell echo $(VERSION) | awk -F. '{print $$1"."$$2"."$$3+1}'))
	@echo "Bumping version from $(VERSION) to $(NEW_VERSION)"
	@echo "VERSION := $(NEW_VERSION)" > .version.tmp
	@sed '1d' GNUmakefile > GNUmakefile.tmp && cat .version.tmp GNUmakefile.tmp > GNUmakefile && rm -f .version.tmp GNUmakefile.tmp
	@echo "Version updated in GNUmakefile"

version-bump-minor: ## Bump minor version (e.g., 1.0.0 -> 1.1.0)
	$(eval NEW_VERSION := $(shell echo $(VERSION) | awk -F. '{print $$1"."$$2+1".0"}'))
	@echo "Bumping version from $(VERSION) to $(NEW_VERSION)"
	@echo "VERSION := $(NEW_VERSION)" > .version.tmp
	@sed '1d' GNUmakefile > GNUmakefile.tmp && cat .version.tmp GNUmakefile.tmp > GNUmakefile && rm -f .version.tmp GNUmakefile.tmp
	@echo "Version updated in GNUmakefile"

version-bump-major: ## Bump major version (e.g., 1.0.0 -> 2.0.0)
	$(eval NEW_VERSION := $(shell echo $(VERSION) | awk -F. '{print $$1+1".0.0"}'))
	@echo "Bumping version from $(VERSION) to $(NEW_VERSION)"
	@echo "VERSION := $(NEW_VERSION)" > .version.tmp
	@sed '1d' GNUmakefile > GNUmakefile.tmp && cat .version.tmp GNUmakefile.tmp > GNUmakefile && rm -f .version.tmp GNUmakefile.tmp
	@echo "Version updated in GNUmakefile"
