terraform {
  required_providers {
    khulnasoft = {
      version = "0.15.0"
      source  = "khulnasoft/khulnasoft"
    }
  }
}

provider "khulnasoft" {
  username = "IaC"                 // Alternatively sourced from $KHULNASOFT_USER
  khulnasoft_url = "https://khulnasofturl.com" // Alternatively sourced from $KHULNASOFT_URL
  password = "@password"           // Alternatively sourced from $KHULNASOFT_PASSWORD

  // If you are using unverifiable certificates (e.g. self-signed) you may need to disable certificate verification
  verify_tls = false // Alternatively sourced from $KHULNASOFT_TLS_VERIFY

  // Alternatively, you can provide these configurations from a config file, and configure the provider as below
  // config_path = '/path/to/tf.config' // defaults to '~/.khulnasoft/tf.config' -- Alternatively sourced from $KHULNASOFT_CONFIG
  // validate = false // Skip provider credential validation

  //Alternatively, you can use API key authentication as below instead of username/password authentication.
  khulnasoft_api_key      = var.khulnasoft_api_key    // Alternatively sourced from $KHULNASOFT_API
  khulnasoft_api_secret   = var.khulnasoft_api_secret // Alternatively sourced from $KHULNASOFT_SECRET
  validity          = 240                    // Alternatively sourced from $KHULNASOFT_TOKEN_VALIDITY
  allowed_endpoints = ["ANY"]                // Alternatively sourced from $KHULNASOFT_ALLOWED_ENDPOINTS
  csp_roles         = ["Admin"]              // Alternatively sourced from $KHULNASOFT_CSP_ROLES
}
