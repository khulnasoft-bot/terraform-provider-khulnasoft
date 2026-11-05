package khulnasoft

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestKhulnasoftRegistryDatasource(t *testing.T) {
	t.Parallel()
	name := acctest.RandomWithPrefix("terraform-test")
	url := "https://docker.io"
	rtype := "HUB"
	username := ""
	password := ""
	autopull := false
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckKhulnasoftRegistryDataSource(name, url, rtype, username, password, autopull),
				Check:  testAccCheckKhulnasoftRegistryDataSourceExists("data.khulnasoft_integration_registries.testregistries"),
			},
		},
	})
}

func testAccCheckKhulnasoftRegistryDataSource(name, url, rtype, username, password string, autopull bool) string {
	return fmt.Sprintf(`
	resource "khulnasoft_integration_registry" "new" {
		name = "%s"
		url = "%s"
		type = "%s"
		username = "%s"
		password = "%s"
		auto_pull = "%v"
	}

	data "khulnasoft_integration_registries" "testregistries" {
		name = khulnasoft_integration_registry.new.name
		depends_on = [
			khulnasoft_integration_registry.new
        ]
	}
	`, name, url, rtype, username, password, autopull)

}

func testAccCheckKhulnasoftRegistryDataSourceExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return NewNotFoundErrorf("%s in state", n)
		}

		if rs.Primary.ID == "" {
			return NewNotFoundErrorf("ID for %s in state", n)
		}

		return nil
	}
}
