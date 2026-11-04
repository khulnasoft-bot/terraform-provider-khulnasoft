package khulnasoft

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestDataKhulnasoftVmwareAssurancePolicy(t *testing.T) {
	t.Parallel()
	name := "Default"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckKhulnasoftVmwareAssurancePolicyDataSource(name),
				Check:  testAccCheckKhulnasoftVmwareAssurancePolicyDataSourceExists("data.khulnasoft_vmware_assurance_policy.defaultvmw"),
			},
		},
	})
}

func testAccCheckKhulnasoftVmwareAssurancePolicyDataSource(name string) string {
	return fmt.Sprintf(`
	data "khulnasoft_vmware_assurance_policy" "defaultvmw" {
		name = "%s"
	}
	`, name)

}

func testAccCheckKhulnasoftVmwareAssurancePolicyDataSourceExists(n string) resource.TestCheckFunc {
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
