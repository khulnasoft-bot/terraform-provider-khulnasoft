## 0.9.0 (Unreleased)

FEATURES:

* **New Authentication Method**: Added comprehensive API key authentication support with `khulnasoft_api_key_id` and `khulnasoft_api_secret` configuration options
* **AWS CodeBuild Integration**: Added support for AWS CodeBuild in the application scope resource for enhanced cloud workload protection
* **Development Tools Enhancement**: Added comprehensive development tools including improved Makefile, linting, testing, and documentation generation

IMPROVEMENTS:

* **Runtime Policy**: Enhanced runtime policy handling with improved error handling and serverless application support
* **Dependencies**: Updated Go module dependencies for improved security and compatibility

BACKWARDS INCOMPATIBILITIES / NOTES:

* This release introduces new authentication options but maintains backward compatibility with existing username/password authentication
