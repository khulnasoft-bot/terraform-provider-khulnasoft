package khulnasoft

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/khulnasoft/terraform-provider-khulnasoft/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAcckhulnasoftAPIKeysDataSource_byID(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckkhulnasoftAPIKeyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAcckhulnasoftAPIKeyAndDataSourceByID(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckkhulnasoftAPIKeyCreated("khulnasoft_khulnasoft_api_key.test"),
					testAccCheckkhulnasoftAPIKeyDataByID("data.khulnasoft_khulnasoft_api_keys.apikeys"),
				),
			},
		},
	})
}

func testAcckhulnasoftAPIKeyAndDataSourceByID() string {
	return `
	resource "khulnasoft_khulnasoft_api_key" "test" {
		description  = "Terraform test key"
		enabled      = true
		roles        = ["role1"]
		ip_addresses = ["127.0.0.1"]
		expiration   = 30
	}
	data "khulnasoft_khulnasoft_api_keys" "apikeys" {
		id = khulnasoft_khulnasoft_api_key.test.id
	}

	output "first_desc" {
		value = data.khulnasoft_khulnasoft_api_keys.apikeys.apikeys[0].description
	}
	output "test_api_key_id" {
		value = data.khulnasoft_khulnasoft_api_keys.apikeys.apikeys[0].id
	}
	`
}

func testAccCheckkhulnasoftAPIKeyCreated(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("resource %q not found in state", n)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("API Key wasn’t created (empty ID)")
		}
		return nil
	}
}

func testAccCheckkhulnasoftAPIKeyDataByID(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("data source %q not found in state", n)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("data source %q has no ID", n)
		}

		descKey := "apikeys.0.description"
		if v, ok := rs.Primary.Attributes[descKey]; !ok || v == "" {
			return fmt.Errorf("attribute %q is empty or not found", descKey)
		}
		return nil
	}
}

func testAccCheckkhulnasoftAPIKeyDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "khulnasoft_khulnasoft_api_key" {
			continue
		}

		id := rs.Primary.ID
		if id == "" {
			continue
		}

		// Try fetching the deleted key, expect failure
		keyID, err := strconv.Atoi(id)
		if err != nil {
			return fmt.Errorf("invalid ID %q: %v", id, err)
		}

		_, err = client.GetApiKey(keyID)
		if err == nil {
			return fmt.Errorf("API Key still exists (ID %d)", keyID)
		}
	}

	return nil
}
