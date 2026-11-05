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
-	[Go](https://golang.org/doc/install) v1.18.x (to build the provider plugin)
- [Khulnasoft Enterprise Platform](https://www.khulnasoft.com/khulnasoft-cloud-native-security-platform/)

## Using the Khulnasoft provider

To quickly get started using the Khulnasoft provider for Terraform, configure the provider as shown below. Full provider documentation with details on all options available is located on the [Terraform Registry site](https://registry.terraform.io/providers/khulnasoft/khulnasoft/latest/docs).

```hcl
terraform {
  required_providers {
    khulnasoft = {
      version = "0.15.1"
      source  = "khulnasoft/khulnasoft"
    }
  }
}

provider "khulnasoft" {
  username = "IaC"
  khulnasoft_url = "https://khulnasofturl.com"
  password = "@password"
}

//Alternative API Based authentication
provider "khulnasoft" {
  khulnasoft_api_key      = var.khulnasoft_api_key
  khulnasoft_api_secret   = var.khulnasoft_api_secret
  validity          = 240
  allowed_endpoints = ["ANY"]
  csp_roles         = ["Admin"]
}
```
## Using the Khulnasoft provider SaaS solution

To quickly get started using the Khulnasoft SaaS provider for Terraform, configure the provider as shown above. The khulnasoft_url should point to cloud.khulnasoft.com for the Khulnasoft Customers and the Dev/QA Teams need to provide their Urls respectively.

**_NOTE:_**  SaaS authentication is supported from version 0.8.4+

## Contributing

The Khulnasoft Provider for Terraform is the work of many contributors. We appreciate your help!

To contribute, please read the [contribution guidelines](CONTRIBUTING.md). You may also [report an issue](https://github.com/khulnasoft/terraform-provider-khulnasoft/issues/new/choose). Once you've filed an issue.
