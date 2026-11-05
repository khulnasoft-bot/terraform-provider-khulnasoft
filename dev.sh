#!/bin/bash
# KhulnaSoft Terraform Provider Development Helper Script
# This script helps developers get started and run common development tasks

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Utility functions
print_header() {
    echo -e "${BLUE}========================================${NC}"
    echo -e "${BLUE}  KhulnaSoft Terraform Provider${NC}"
    echo -e "${BLUE}  Development Helper Script${NC}"
    echo -e "${BLUE}========================================${NC}"
    echo ""
}

print_step() {
    echo -e "${GREEN}[STEP]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

print_info() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

# Check if command exists
command_exists() {
    command -v "$1" >/dev/null 2>&1
}

# Setup development environment
setup() {
    print_header
    print_step "Setting up development environment..."

    # Check Go installation
    if ! command_exists go; then
        print_error "Go is not installed. Please install Go 1.18 or later."
        exit 1
    fi

    # Check Go version
    GO_VERSION=$(go version | awk '{print $3}' | sed 's/go//')
    print_info "Go version: $GO_VERSION"

    # Setup Go modules
    print_step "Downloading dependencies..."
    go mod download
    go mod verify

    # Install development tools
    print_step "Installing development tools..."
    make tools

    # Setup environment file
    if [ ! -f .env.dev ]; then
        print_step "Creating development environment file..."
        make env-setup
        print_warning "Please edit .env.dev with your actual KhulnaSoft credentials"
    fi

    print_step "Setup complete!"
    echo ""
    print_info "Common development commands:"
    echo "  make help          - Show all available Makefile targets"
    echo "  make dev           - Setup development environment"
    echo "  make build         - Build the provider"
    echo "  make test          - Run unit tests"
    echo "  make fmt           - Format code"
    echo "  make lint          - Run linter"
    echo "  make docs          - Generate documentation"
    echo "  make install       - Install provider to Terraform"
}

# Run tests
run_tests() {
    print_header
    print_step "Running tests..."

    if [ "$1" = "acc" ]; then
        print_info "Running acceptance tests (requires TF_ACC=1 and credentials)"
        if [ -z "$KHULNASOFT_USER" ] || [ -z "$KHULNASOFT_PASSWORD" ] || [ -z "$KHULNASOFT_URL" ]; then
            print_warning "Acceptance tests require KHULNASOFT_USER, KHULNASOFT_PASSWORD, and KHULNASOFT_URL environment variables"
            print_info "Or set TF_ACC=1 and run tests"
        fi
        make testacc
    else
        print_info "Running unit tests..."
        make test
    fi
}

# Build provider
build() {
    print_header
    print_step "Building provider..."
    make build
    print_info "Build complete!"
}

# Check code quality
check_quality() {
    print_header
    print_step "Checking code quality..."

    print_info "Formatting code..."
    make fmt

    print_info "Running linter..."
    make lint

    print_info "Running go vet..."
    make vet

    print_info "Code quality check complete!"
}

# Generate documentation
gen_docs() {
    print_header
    print_step "Generating documentation..."
    make docs
    print_info "Documentation generated!"
}

# Install provider
install_provider() {
    print_header
    print_step "Installing provider to Terraform..."
    make install
    print_info "Provider installed!"
}

# Show status
show_status() {
    print_header
    print_step "Project Status"

    echo "Go version: $(go version)"
    echo "Current directory: $(pwd)"
    echo "Provider version: $(grep 'VERSION :=' GNUmakefile | awk '{print $3}')"
    echo "OS/Architecture: $(go env GOOS)_$(go env GOARCH)"

    if [ -f ".env.dev" ]; then
        print_info "Development environment file exists"
    else
        print_warning "Development environment file not found. Run './dev.sh setup'"
    fi

    echo ""
    print_info "Environment variables:"
    echo "  KHULNASOFT_USER: ${KHULNASOFT_USER:-<not set>}"
    echo "  KHULNASOFT_PASSWORD: ${KHULNASOFT_PASSWORD:-<not set>}"
    echo "  KHULNASOFT_URL: ${KHULNASOFT_URL:-<not set>}"
    echo "  TF_ACC: ${TF_ACC:-<not set>}"
}

# Main script logic
main() {
    case "${1:-help}" in
        "setup")
            setup
            ;;
        "test")
            run_tests
            ;;
        "testacc")
            run_tests acc
            ;;
        "build")
            build
            ;;
        "quality")
            check_quality
            ;;
        "docs")
            gen_docs
            ;;
        "install")
            install_provider
            ;;
        "status")
            show_status
            ;;
        "clean")
            print_header
            print_step "Cleaning project..."
            make clean
            print_info "Clean complete!"
            ;;
        "help"|"-h"|"--help")
            print_header
            echo "Usage: $0 [COMMAND]"
            echo ""
            echo "Commands:"
            echo "  setup     - Setup development environment"
            echo "  build     - Build the provider"
            echo "  test      - Run unit tests"
            echo "  testacc   - Run acceptance tests"
            echo "  quality   - Check code quality (fmt, lint, vet)"
            echo "  docs      - Generate documentation"
            echo "  install   - Install provider to Terraform"
            echo "  clean     - Clean build artifacts"
            echo "  status    - Show project status"
            echo "  help      - Show this help message"
            echo ""
            echo "Environment variables:"
            echo "  KHULNASOFT_USER     - KhulnaSoft username"
            echo "  KHULNASOFT_PASSWORD - KhulnaSoft password"
            echo "  KHULNASOFT_URL      - KhulnaSoft instance URL"
            echo "  TF_ACC             - Set to 1 for acceptance tests"
            ;;
        *)
            print_error "Unknown command: $1"
            echo ""
            echo "Run '$0 help' for available commands."
            exit 1
            ;;
    esac
}

# Run main function with all arguments
main "$@"
