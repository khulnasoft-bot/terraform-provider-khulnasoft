package khulnasoft

import (
	"fmt"
	"testing"

	"github.com/khulnasoft/terraform-provider-khulnasoft/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const (
	testRoleMappingSaasResourceName = "khulnasoft_role_mapping_saas.test"
)

func TestAccKhulnasoftRoleMappingSaas_basic(t *testing.T) {
	if !isSaasEnv() {
		t.Skip("Skipping role mapping SaaS test - not a SaaS environment")
	}

	roleName := acctest.RandomWithPrefix("tf-role")
	permSetName := acctest.RandomWithPrefix("tf-pset")
	samlGroups := []string{"DevTeam", "SecTeam"}
	permissionActions, _ := testAccPermissionSetActionFixtures(t)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckRoleMappingSaasDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccRoleMappingSaasConfig(permSetName, roleName, samlGroups, permissionActions),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoleMappingSaasExists(testRoleMappingSaasResourceName),
					resource.TestCheckResourceAttr(testRoleMappingSaasResourceName, "csp_role", roleName),
					resource.TestCheckResourceAttr(testRoleMappingSaasResourceName, "saml_groups.#", "2"),
				),
			},
			{
				ResourceName:      testRoleMappingSaasResourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccRoleMappingSaasConfig(permSetName, roleName string, groups, permissionActions []string) string {
	groupsStr := ""
	for _, g := range groups {
		groupsStr += fmt.Sprintf("\"%s\", ", g)
	}
	groupsStr = groupsStr[:len(groupsStr)-2]

	return fmt.Sprintf(`
resource "khulnasoft_permission_set_saas" "ps" {
  name        = "%s"
  description = "TF-generated permission set"
  actions     = [%s]
}

resource "khulnasoft_role" "r" {
  role_name   = "%s"
  description = "TF-generated role"
  permission  = khulnasoft_permission_set_saas.ps.name
  scopes      = ["Global"]
}

resource "khulnasoft_role_mapping_saas" "test" {
  saml_groups = [%s]
  csp_role    = khulnasoft_role.r.role_name
}
`, permSetName, permissionSetActionsConfig(permissionActions), roleName, groupsStr)
}

func testAccCheckRoleMappingSaasExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Resource not found in state: %s", n)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("ID not set for resource: %s", n)
		}
		cli := testAccProvider.Meta().(*client.Client)
		_, err := cli.GetRoleMappingSaas(rs.Primary.ID)
		return err
	}
}

func testAccCheckRoleMappingSaasDestroy(s *terraform.State) error {
	cli := testAccProvider.Meta().(*client.Client)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "khulnasoft_role_mapping_saas" {
			continue
		}
		_, err := cli.GetRoleMappingSaas(rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("role mapping %q still exists", rs.Primary.ID)
		}
	}
	return nil
}
