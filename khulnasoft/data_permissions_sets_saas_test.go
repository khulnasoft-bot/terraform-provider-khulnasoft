package khulnasoft

import (
    "fmt"
    "testing"

    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
    "github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestKhulnasoftPermissionsSetSaasDatasource(t *testing.T) {
    if !isSaasEnv() {
        t.Skip("Skipping permission set test because its not a SaaS environment")
    }

    resource.Test(t, resource.TestCase{
        PreCheck:  func() { testAccPreCheck(t) },
        Providers: testAccProviders,
        Steps: []resource.TestStep{
            {
                Config: testAccCheckKhulnasoftPermissionsSetSaasBasicConfig(),
                Check: resource.ComposeTestCheckFunc(
                    // Check that the data source was created
                    testAccCheckKhulnasoftPermissionsSetSaasExists("data.khulnasoft_permissions_sets_saas.test"),
                    // Check list attribute is populated
                    resource.TestCheckResourceAttrSet("data.khulnasoft_permissions_sets_saas.test", "permissions_sets.#"),
                    // Custom check for data validity
                    testAccCheckPermissionsSetSaasAttributes("data.khulnasoft_permissions_sets_saas.test"),
                ),
            },
        },
    })
}

func TestKhulnasoftPermissionsSetSaasDatasource_Structure(t *testing.T) {
    if !isSaasEnv() {
        t.Skip("Skipping permission set test because its not a SaaS environment")
    }

    resource.Test(t, resource.TestCase{
        PreCheck:  func() { testAccPreCheck(t) },
        Providers: testAccProviders,
        Steps: []resource.TestStep{
            {
                Config: testAccCheckKhulnasoftPermissionsSetSaasBasicConfig(),
                Check: resource.ComposeTestCheckFunc(
                    // Verify ID exists
                    resource.TestCheckResourceAttrSet(
                        "data.khulnasoft_permissions_sets_saas.test",
                        "id",
                    ),
                    // Verify permissions_sets list exists
                    resource.TestCheckResourceAttrSet(
                        "data.khulnasoft_permissions_sets_saas.test",
                        "permissions_sets.#",
                    ),
                ),
            },
        },
    })
}

// Basic Config
func testAccCheckKhulnasoftPermissionsSetSaasBasicConfig() string {
    return `
    data "khulnasoft_permissions_sets_saas" "test" {
    }
    `
}

// Check data source exists
func testAccCheckKhulnasoftPermissionsSetSaasExists(n string) resource.TestCheckFunc {
    return func(s *terraform.State) error {
        rs, ok := s.RootModule().Resources[n]
        if !ok {
            return fmt.Errorf("Not found: %s", n)
        }

        if rs.Primary.ID == "" {
            return fmt.Errorf("No ID is set")
        }

        return nil
    }
}

// Check attributes match schema
func testAccCheckPermissionsSetSaasAttributes(resourceName string) resource.TestCheckFunc {
    return func(s *terraform.State) error {
        rs, ok := s.RootModule().Resources[resourceName]
        if !ok {
            return fmt.Errorf("Not found: %s", resourceName)
        }

        numSets, ok := rs.Primary.Attributes["permissions_sets.#"]
        if !ok {
            return fmt.Errorf("No permissions_sets found")
        }

        // If we have permission sets, verify their structure
        if numSets != "0" {
            // Get the first permission set and verify required attributes
            if name, ok := rs.Primary.Attributes["permissions_sets.0.name"]; !ok || name == "" {
                return fmt.Errorf("permissions_sets.0.name is empty or missing")
            }

            if desc, exists := rs.Primary.Attributes["permissions_sets.0.description"]; exists {
                if desc == "" {
                    return fmt.Errorf("permissions_sets.0.description is empty")
                }
            }

            if _, ok := rs.Primary.Attributes["permissions_sets.0.actions.#"]; !ok {
                return fmt.Errorf("permissions_sets.0.actions is missing")
            }
        }

        return nil
    }
}