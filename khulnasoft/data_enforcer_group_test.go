package khulnasoft

import (
	"fmt"
	"testing"

	"github.com/khulnasoft/terraform-provider-khulnasoft/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestKhulnasoftEnforcerGroupDatasource(t *testing.T) {
	t.Parallel()
	var basicEnforcerGroup = client.EnforcerGroup{
		ID:          acctest.RandomWithPrefix("terraform-test"),
		Description: "Created",
		LogicalName: "terraform-eg",
		Enforce:     false,
		Gateways: []string{
			"3ef9a43f2693_gateway",
		},
		Type:              "agent",
		EnforcerImageName: "registry.khulnasoft.com/enforcer:6.5.22034",
		Orchestrator:      client.EnforcerOrchestrator{},
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckKhulnasoftEnforcerGroupDataSource(basicEnforcerGroup),
				Check:  testAccCheckKhulnasoftEnforcerGroupDataSourceExists("data.khulnasoft_enforcer_groups.testegdata"),
			},
		},
	})
}

func testAccCheckKhulnasoftEnforcerGroupDataSource(enforcerGroup client.EnforcerGroup) string {
	return fmt.Sprintf(`
	
	resource "khulnasoft_enforcer_groups" "testegdata" {
		group_id = "%s"
		description = "%s"
		logical_name = "%s"
		enforce = "%v"
		gateways = ["%s"]
		type = "%s"
		orchestrator {
			type = "%s"
            service_account = "%s"
			namespace = "%s"
			master = "%v"
		}
	}
	data "khulnasoft_enforcer_groups" "testegdata" {
		group_id = khulnasoft_enforcer_groups.testegdata.group_id
		depends_on = [
          khulnasoft_enforcer_groups.testegdata
        ]
	}
	`,
		enforcerGroup.ID,
		enforcerGroup.Description,
		enforcerGroup.LogicalName,
		enforcerGroup.Enforce,
		enforcerGroup.Gateways[0],
		enforcerGroup.Type,
		enforcerGroup.Orchestrator.Type,
		enforcerGroup.Orchestrator.ServiceAccount,
		enforcerGroup.Orchestrator.Namespace,
		enforcerGroup.Orchestrator.Master)

}

func testAccCheckKhulnasoftEnforcerGroupDataSourceExists(n string) resource.TestCheckFunc {
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
