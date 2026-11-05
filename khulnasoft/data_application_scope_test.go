package khulnasoft

import (
	"fmt"
	"testing"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestDataApplicationScope(t *testing.T) {
	t.Parallel()
	name := "Global"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckApplicationScopeDataSource(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApplicationScopeExists("data.khulnasoft_application_scope.defaultiap"),
					// Basic attribute checks
					resource.TestCheckResourceAttr("data.khulnasoft_application_scope.defaultiap", "name", name),
					resource.TestCheckResourceAttrSet("data.khulnasoft_application_scope.defaultiap", "description"),
					resource.TestCheckResourceAttrSet("data.khulnasoft_application_scope.defaultiap", "author"),
					// Categories existence check
					resource.TestCheckResourceAttrSet("data.khulnasoft_application_scope.defaultiap", "categories.#"),
				),
			},
		},
	})
}

// Test specific for codebuild data source functionality
func TestDataApplicationScopeWithCodeBuild(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				// First create a scope with codebuild
				Config: testAccCheckApplicationScopeWithCodeBuild(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApplicationScopeExists("khulnasoft_application_scope.test_codebuild"),
				),
			},
			{
				// Then test reading it as a data source
				Config: testAccCheckApplicationScopeWithCodeBuildDataSource(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApplicationScopeExists("data.khulnasoft_application_scope.test_codebuild_ds"),
					// Verify codebuild specific attributes
					resource.TestCheckResourceAttr(
						"data.khulnasoft_application_scope.test_codebuild_ds",
						"categories.0.artifacts.0.codebuild.0.expression",
						"v1",
					),
					resource.TestCheckResourceAttr(
						"data.khulnasoft_application_scope.test_codebuild_ds",
						"categories.0.artifacts.0.codebuild.0.variables.0.attribute",
						"khulnasoft.topic",
					),
					resource.TestCheckResourceAttr(
						"data.khulnasoft_application_scope.test_codebuild_ds",
						"categories.0.artifacts.0.codebuild.0.variables.0.value",
						"topic1",
					),
				),
			},
		},
	})
}

func testAccCheckApplicationScopeDataSource(name string) string {
	return fmt.Sprintf(`
	data "khulnasoft_application_scope" "defaultiap" {
		name = "%s"
	}
	`, name)
}

func testAccCheckApplicationScopeWithCodeBuild() string {
	return `
	resource "khulnasoft_application_scope" "test_codebuild" {
		description = "test codebuild application scope"
		name        = "test_codebuild"
		
		categories {
			artifacts {
				codebuild {
					expression = "v1"
					variables {
						attribute = "khulnasoft.topic"
						value     = "topic1"
					}
				}
			}
		}
	}`
}

func testAccCheckApplicationScopeWithCodeBuildDataSource() string {
	return `
	resource "khulnasoft_application_scope" "test_codebuild" {
		description = "test codebuild application scope"
		name        = "test_codebuild"
		
		categories {
			artifacts {
				codebuild {
					expression = "v1"
					variables {
						attribute = "khulnasoft.topic"
						value     = "topic1"
					}
				}
			}
		}
	}

	data "khulnasoft_application_scope" "test_codebuild_ds" {
		name = khulnasoft_application_scope.test_codebuild.name
	}`
}
