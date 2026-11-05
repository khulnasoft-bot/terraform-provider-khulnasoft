terraform {
  required_providers {
    khulnasoft = {
      version = "0.8.30"
      source  = "khulnasoft/khulnasoft"
    }
  }
}

# Example using username and password authentication
provider "khulnasoft" {
  username = "IaC"                 // Alternatively sourced from $KHULNASOFT_USER
  khulnasoft_url = "https://khulnasoft.com" // Alternatively sourced from $KHULNASOFT_URL
  password = "@password"           // Alternatively sourced from $KHULNASOFT_PASSWORD

  // If you are using unverifiable certificates (e.g. self-signed) you may need to disable certificate verification
  verify_tls = false // Alternatively sourced from $KHULNASOFT_TLS_VERIFY

  // Alternatively, you can provide these configurations from a config file, and configure the provider as below
  // config_path = '/path/to/tf.config' // defaults to '~/.khulnasoft/tf.config' -- Alternatively sourced from $KHULNASOFT_CONFIG
}

# Example using API key authentication (alternative to username/password)
# provider "khulnasoft" {
#   khulnasoft_api_key_id = "your-api-key-id"    // Alternatively sourced from $KHULNASOFT_API_KEY_ID
#   khulnasoft_api_secret = "your-api-secret"   // Alternatively sourced from $KHULNASOFT_API_SECRET
#   khulnasoft_url = "https://khulnasoft.com"   // Alternatively sourced from $KHULNASOFT_URL
#
#   // If you are using unverifiable certificates (e.g. self-signed) you may need to disable certificate verification
#   verify_tls = false // Alternatively sourced from $KHULNASOFT_TLS_VERIFY
# }
