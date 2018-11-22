package azurerm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccAzureRMIssueUSeSet(t *testing.T) {
	config := getConfig()
	updatedConfig := getUpdatedConfig()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check:  emptyCheck(),
			},
			{
				Config: updatedConfig,
				Check:  emptyCheck(),
			},
		},
	})
}

func emptyCheck() resource.TestCheckFunc {
	return func(s *terraform.State) error {
		return nil
	}
}

func getConfig() string {
	return fmt.Sprint(`
resource "azurerm_issue_use_set" "test" {
  locations_list      = ["eastus", "westus"]
  locations_set		  = ["eastus", "westus"]
}
`)
}

func getUpdatedConfig() string {
	return fmt.Sprint(`
resource "azurerm_issue_use_set" "test" {
  locations_list      = ["eastus", "westeurope"]
  locations_set		  = ["eastus", "westeurope"]
}
`)
}

func testCheckDestroy(s *terraform.State) error {
	return nil
}
