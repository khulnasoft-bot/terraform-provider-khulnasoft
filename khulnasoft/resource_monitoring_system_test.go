package khulnasoft

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccKhulnasoftMonitoringSystem(t *testing.T) {
	t.Skip()
	t.Parallel()

	name := "Prometheus"
	msType := "prometheus"
	token := "tf-acc-test-token"
	enabled := true
	interval := 30

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: CheckDestroy("khulnasoft_monitoring_system.prom_mon"),
		Steps: []resource.TestStep{
			{
				Config: testAccMonitoringSystemResourceConfig(name, msType, token, enabled, interval),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckKhulnasoftMonitoringSystemExists("khulnasoft_monitoring_system.prom_mon"),
					resource.TestCheckResourceAttr("khulnasoft_monitoring_system.prom_mon", "name", name),
					resource.TestCheckResourceAttr("khulnasoft_monitoring_system.prom_mon", "type", msType),
					resource.TestCheckResourceAttr("khulnasoft_monitoring_system.prom_mon", "enabled", fmt.Sprintf("%t", enabled)),
					resource.TestCheckResourceAttr("khulnasoft_monitoring_system.prom_mon", "interval", fmt.Sprintf("%d", interval)),
				),
			},
			{
				ResourceName:      "khulnasoft_monitoring_system.prom_mon",
				ImportState:       true,
				ImportStateVerify: true,
				// If token isn't returned by Read or is write-only, ignore it during import verification.
				ImportStateVerifyIgnore: []string{"token", "last_updated"},
			},
		},
	})
}

func testAccMonitoringSystemResourceConfig(name, msType, token string, enabled bool, interval int) string {
	return fmt.Sprintf(`
resource "khulnasoft_monitoring_system" "prom_mon" {
  name        = "%s"
  type        = "%s"
  token       = "%s"
  enabled     = %t
  interval    = %d
}
`, name, msType, token, enabled, interval)
}

func testAccCheckKhulnasoftMonitoringSystemExists(n string) resource.TestCheckFunc {
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
