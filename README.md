<a href="https://terraform.io">
    <img src="Terraform_PrimaryLogo_Color_RGB.png" alt="Terraform logo" title="Terraform" height="100" />
</a>
<a href="https://www.khulnasoft.com/">
    <img src="https://avatars3.githubusercontent.com/u/43526139?s=200&v=4" alt="Khulnasoft logo" title="Khulnasoft" height="100" />
</a>

Khulnasoft Provider for Terraform
===========================

This is the Khulnasoft provider for [Terraform](https://www.terraform.io/).

Useful links:
- [Khulnasoft Documentation](https://docs.khulnasoft.com)
- [Khulnasoft Provider Documentation](https://registry.terraform.io/providers/khulnasoft/khulnasoft/latest/docs)
- [Terraform Documentation](https://www.terraform.io/docs/language/index.html)
- [Terraform Provider Development](DEVELOPMENT.md)

The provider lets you declaratively define the configuration for your Khulnasoft Enterprise platform.


## Contents

- [Khulnasoft Provider for Terraform](#khulnasoft-provider-for-terraform)
  - [Contents](#contents)
  - [Requirements](#requirements)
  - [Using the Khulnasoft provider](#using-the-khulnasoft-provider)
  - [Using the Khulnasoft provider SaaS solution](#using-the-khulnasoft-provider-saas-solution)
  - [Contributing](#contributing)


## Requirements

-	[Terraform](https://www.terraform.io/downloads.html) v0.12.x or higher
-	[Go](https://golang.org/doc/install) v1.16.x (to build the provider plugin)
- [Khulnasoft Enterprise Platform](https://www.khulnasoft.com/khulnasoft-cloud-native-security-platform/)

## Using the Khulnasoft provider

To quickly get started using the Khulnasoft provider for Terraform, configure the provider as shown below. Full provider documentation with details on all options available is located on the [Terraform Registry site](https://registry.terraform.io/providers/khulnasoft/khulnasoft/latest/docs).

```hcl
terraform {
  required_providers {
    khulnasoft = {
      version = "0.8.26"
      source  = "khulnasoft/khulnasoft"
    }
  }
}

provider "khulnasoft" {
  username = "IaC"
  khulnasoft_url = "https://khulnasoft.com"
  password = "@password"
}
```

## Authentication Methods

The Khulnasoft provider supports two authentication methods:

### Username/Password Authentication (Default)

```hcl
provider "khulnasoft" {
  username = "IaC"
  password = "@password"
  khulnasoft_url = "https://khulnasoft.com"
}
```

### API Key Authentication

```hcl
provider "khulnasoft" {
  khulnasoft_api_key_id = "your-api-key-id"
  khulnasoft_api_secret = "your-api-secret"
  khulnasoft_url = "https://khulnasoft.com"
}
```

**Note:** You cannot use both authentication methods simultaneously. Choose either username/password OR API key authentication.

## Using the Khulnasoft provider SaaS solution

To quickly get started using the Khulnasoft SaaS provider for Terraform, configure the provider as shown above. The khulnasoft_url should point to cloud.khulnasoft.com for the Khulnasoft Customers and the Dev/QA Teams need to provide their Urls respectively.

**_NOTE:_**  SaaS authentication is supported from version 0.8.4+

## Contributing

The KhulnaSoft Provider for Terraform is the work of many contributors. We appreciate your help!

To contribute, please read the [contribution guidelines](CONTRIBUTING.md). You may also [report an issue](https://github.com/khulnasoft/terraform-provider-khulnasoft/issues/new/choose). Once you've filed an issue.

## Development

This project includes comprehensive development tooling to help contributors get started quickly and maintain code quality.

### Quick Start

1. **Setup development environment:**
   ```bash
   ./dev.sh setup
   ```

2. **Build the provider:**
   ```bash
   make build
   ```

3. **Run tests:**
   ```bash
   make test          # Unit tests
   make testacc       # Acceptance tests (requires credentials)
   ```

4. **Install to Terraform:**
   ```bash
   make install
   ```

### Development Script

The `dev.sh` script provides a user-friendly interface for common development tasks:

```bash
./dev.sh setup     # Setup development environment
./dev.sh build     # Build the provider
./dev.sh test      # Run unit tests
./dev.sh testacc   # Run acceptance tests
./dev.sh quality   # Check code quality (format, lint, vet)
./dev.sh docs      # Generate documentation
./dev.sh install   # Install to Terraform
./dev.sh status    # Show project status
./dev.sh help      # Show all available commands
```

### Makefile Targets

The Makefile provides comprehensive targets for development:

#### Build & Install
- `make build` - Build the provider binary
- `make install` - Install provider to Terraform plugins directory
- `make clean` - Clean build artifacts

#### Testing
- `make test` - Run unit tests
- `make testacc` - Run acceptance tests
- `make test-coverage` - Run tests with coverage report

#### Code Quality
- `make fmt` - Format Go code
- `make lint` - Run linter
- `make vet` - Run go vet
- `make quality` - Run all quality checks

#### Documentation
- `make docs` - Generate documentation
- `make docs-serve` - Serve documentation locally

#### Development
- `make dev` - Setup development environment
- `make tools` - Install development tools
- `make deps` - Update dependencies

#### Release & Version Management
- `make release` - Create a new release
- `make tag` - Create and push git tag
- `make version-bump-patch` - Bump patch version

### Environment Variables

For development and testing, set these environment variables:

```bash
# Authentication (choose one method)
export KHULNASOFT_USER="your-username"
export KHULNASOFT_PASSWORD="your-password"
# OR for API key authentication:
export KHULNASOFT_API_KEY_ID="your-api-key-id"
export KHULNASOFT_API_SECRET="your-api-secret"

# Instance configuration
export KHULNASOFT_URL="https://your-khulnasoft-instance.com"

# TLS settings
export KHULNASOFT_TLS_VERIFY="true"
export KHULNASOFT_CA_CERT_PATH="/path/to/ca-cert.pem"

# Testing
export TF_ACC=1  # Enable acceptance tests
```

### Getting Help

- `make help` - Show all available Makefile targets
- `./dev.sh help` - Show development script commands
- `make info` - Show project information

### Development Workflow

1. **Make changes** to the codebase
2. **Format code:** `make fmt`
3. **Run tests:** `make test`
4. **Check quality:** `make quality`
5. **Build:** `make build`
6. **Test manually** with Terraform
7. **Commit changes** with descriptive commit messages

For more detailed development information, see [DEVELOPMENT.md](DEVELOPMENT.md).
