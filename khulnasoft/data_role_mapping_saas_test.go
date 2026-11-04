package khulnasoft

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestKhulnasoftRolesMappingSaasDatasource(t *testing.T) {
	if !isSaasEnv() {
		t.Skip("Skipping SaaS roles mapping data source test - not a SaaS environment")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckKhulnasoftRolesMappingSaasBasicConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckKhulnasoftRolesMappingSaasExists("data.khulnasoft_roles_mapping_saas.test"),
					resource.TestCheckResourceAttrSet("data.khulnasoft_roles_mapping_saas.test", "roles_mapping.#"),
				),
			},
		},
	})
}

func testAccCheckKhulnasoftRolesMappingSaasBasicConfig() string {
	return `
	data "khulnasoft_roles_mapping_saas" "test" {}
	`
}

func testAccCheckKhulnasoftRolesMappingSaasExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set for %s", n)
		}
		return nil
	}
}
