package khulnasoft

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestKhulnasoftUsersSaasManagement(t *testing.T) {

	if !isSaasEnv() {
		t.Skip("Skipping saas user test because its on prem env")
	}

	t.Parallel()
	userID := acctest.RandomWithPrefix("terrafrom.user")
	email := fmt.Sprintf("%s@khulnasoft.com", userID)
	groups := acctest.RandomWithPrefix("firstGroup")
	newGroups := acctest.RandomWithPrefix("secondGroup")

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccUsersSaasDestroy,
		Steps: []resource.TestStep{
			{
				// Config returns the test resource
				Config: testAccCheckKhulnasoftUsersSaas(email, groups, true, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckKhulnasoftUsersSaassExists("khulnasoft_user_saas.new"),
				),
			},
			{
				// Config returns the test resource
				Config: testAccCheckKhulnasoftUsersSaas1(email, newGroups, true, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckKhulnasoftUsersSaassExists("khulnasoft_user_saas.new"),
				),
			},
			{
				ResourceName:            "khulnasoft_user_saas.new",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"groups"}, //TODO: add groups to read
			},
		},
	})
}

func testAccCheckKhulnasoftUsersSaas(email, groups string, groupAdmin, accountAdmin bool) string {
	return fmt.Sprintf(`

	resource "khulnasoft_group" "new" {
		name    = "%s"
    }

	resource "khulnasoft_user_saas" "new" {
		email    = "%s"
		csp_roles = []
		groups {
			name = "%s"
			group_admin = %v
		}
		account_admin = %v
		depends_on = ["khulnasoft_group.new"]
	  }`, groups, email, groups, groupAdmin, accountAdmin)
}

func testAccCheckKhulnasoftUsersSaas1(email, groups string, groupAdmin, accountAdmin bool) string {
	return fmt.Sprintf(`

	resource "khulnasoft_group" "new" {
		name    = "%s"
    }
	
	resource "khulnasoft_user_saas" "new" {
		email    = "%s"
		csp_roles = []
		groups {
			name = "%s"
			group_admin = %v
		}
		account_admin = %v
		depends_on = ["khulnasoft_group.new"]
	  }`, groups, email, groups, groupAdmin, accountAdmin)
}

func testAccCheckKhulnasoftUsersSaassExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return NewNotFoundErrorf("%s in state", n)
		}

		if rs.Primary.ID == "" {
			return NewNotFoundErrorf("Id for %s in state", n)
		}
		return nil
	}
}

func testAccUsersSaasDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "khulnasoft_user.new" && rs.Type != "khulnasoft_group.new" {
			continue
		}

		if rs.Primary.ID != "" {
			return fmt.Errorf("Object %q still exists", rs.Primary.ID)
		}
		return nil
	}
	return nil
}
