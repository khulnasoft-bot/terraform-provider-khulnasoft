package khulnasoft

import (
	"context"
	"os"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/mitchellh/go-homedir"
)

var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider
var testVersion = "1.0"

func init() {
	testAccProvider = Provider(testVersion)
	testAccProviders = map[string]*schema.Provider{
		"khulnasoft": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	t.Parallel()
	if err := Provider(testVersion).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	t.Parallel()
	var _ *schema.Provider = Provider(testVersion)
}

func testAccPreCheck(t *testing.T) {
	configPath, _ := homedir.Expand("~/.khulnasoft/tf.config")
	if _, err := os.Stat(configPath); !os.IsNotExist(err) {
		return
	}

	// Check if using API key authentication
	if os.Getenv("KHULNASOFT_USE_API_KEY") != "" {
		useAPIKey, err := strconv.ParseBool(os.Getenv("KHULNASOFT_USE_API_KEY"))
		if err != nil {
			t.Fatalf("Invalid KHULNASOFT_USE_API_KEY value: %v", err)
		}

		if useAPIKey {
			if os.Getenv("KHULNASOFT_API_KEY_ID") == "" {
				t.Fatal("KHULNASOFT_API_KEY_ID must be set for acceptance tests when using API key authentication")
			}
			if os.Getenv("KHULNASOFT_API_SECRET") == "" {
				t.Fatal("KHULNASOFT_API_SECRET must be set for acceptance tests when using API key authentication")
			}
			if os.Getenv("KHULNASOFT_URL") == "" {
				t.Fatal("KHULNASOFT_URL must be set for acceptance tests")
			}
			return
		}
	}

	// Default to username/password authentication
	if os.Getenv("KHULNASOFT_USER") == "" {
		t.Fatal("KHULNASOFT_USER must be set for acceptance tests")
	}

	if os.Getenv("KHULNASOFT_PASSWORD") == "" {
		t.Fatal("KHULNASOFT_PASSWORD must be set for acceptance tests")
	}

}

func TestProvider_APIKeyAuth(t *testing.T) {
	t.Parallel()

	// Test with API key authentication
	t.Setenv("KHULNASOFT_API_KEY_ID", "test-api-key")
	t.Setenv("KHULNASOFT_API_SECRET", "test-api-secret")
	t.Setenv("KHULNASOFT_URL", "https://test.khulnasoft.com")

	provider := Provider(testVersion)
	err := provider.InternalValidate()
	if err != nil {
		t.Fatalf("Provider validation failed: %s", err)
	}

	// Test that the schema contains the API key fields
	if _, ok := provider.Schema["khulnasoft_api_key_id"]; !ok {
		t.Error("khulnasoft_api_key_id field not found in schema")
	}
	if _, ok := provider.Schema["khulnasoft_api_secret"]; !ok {
		t.Error("khulnasoft_api_secret field not found in schema")
	}
}

func TestProvider_MixedAuthError(t *testing.T) {
	t.Parallel()

	// Test that using both authentication methods produces an error
	provider := Provider(testVersion)

	// Create a ResourceData-like structure for testing
	d := schema.TestResourceDataRaw(t, provider.Schema, map[string]interface{}{
		"username":               "test-user",
		"password":               "test-password",
		"khulnasoft_api_key_id":  "test-api-key",
		"khulnasoft_api_secret": "test-api-secret",
		"khulnasoft_url":        "https://test.khulnasoft.com",
	})

	_, diags := providerConfigure(context.Background(), d)

	if !diags.HasError() {
		t.Error("Expected error when using both username/password and API key authentication")
	}
}

func TestProvider_APIKeyValidation(t *testing.T) {
	t.Parallel()

	// Test that API key fields must be provided together
	provider := Provider(testVersion)

	// Test with only API key ID (should fail)
	d1 := schema.TestResourceDataRaw(t, provider.Schema, map[string]interface{}{
		"khulnasoft_api_key_id": "test-api-key",
		"khulnasoft_url":       "https://test.khulnasoft.com",
	})

	_, diags1 := providerConfigure(context.Background(), d1)

	if !diags1.HasError() {
		t.Error("Expected error when providing only khulnasoft_api_key_id")
	}

	// Test with only API secret (should fail)
	d2 := schema.TestResourceDataRaw(t, provider.Schema, map[string]interface{}{
		"khulnasoft_api_secret": "test-api-secret",
		"khulnasoft_url":       "https://test.khulnasoft.com",
	})

	_, diags2 := providerConfigure(context.Background(), d2)

	if !diags2.HasError() {
		t.Error("Expected error when providing only khulnasoft_api_secret")
	}
}
