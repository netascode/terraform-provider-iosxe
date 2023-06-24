// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIosxeInterfacePIM(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxeInterfacePIMPrerequisitesConfig + testAccIosxeInterfacePIMConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("iosxe_interface_pim.test", "passive", "false"),
					resource.TestCheckResourceAttr("iosxe_interface_pim.test", "dense_mode", "false"),
					resource.TestCheckResourceAttr("iosxe_interface_pim.test", "sparse_mode", "true"),
					resource.TestCheckResourceAttr("iosxe_interface_pim.test", "sparse_dense_mode", "false"),
					resource.TestCheckResourceAttr("iosxe_interface_pim.test", "bfd", "false"),
					resource.TestCheckResourceAttr("iosxe_interface_pim.test", "border", "false"),
					resource.TestCheckResourceAttr("iosxe_interface_pim.test", "bsr_border", "false"),
					resource.TestCheckResourceAttr("iosxe_interface_pim.test", "dr_priority", "10"),
				),
			},
			{
				ResourceName:  "iosxe_interface_pim.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XE-native:native/interface/Loopback=100/ip/pim",
			},
		},
	})
}

const testAccIosxeInterfacePIMPrerequisitesConfig = `
resource "iosxe_restconf" "PreReq0" {
  path = "Cisco-IOS-XE-native:native/interface/Loopback=100"
  attributes = {
      "name" = "100"
  }
}

`

func testAccIosxeInterfacePIMConfig_minimum() string {
	return `
	resource "iosxe_interface_pim" "test" {
		type = "Loopback"
		name = "100"
  		depends_on = [iosxe_restconf.PreReq0, ]
	}
	`
}

func testAccIosxeInterfacePIMConfig_all() string {
	return `
	resource "iosxe_interface_pim" "test" {
		type = "Loopback"
		name = "100"
		passive = false
		dense_mode = false
		sparse_mode = true
		sparse_dense_mode = false
		bfd = false
		border = false
		bsr_border = false
		dr_priority = 10
  		depends_on = [iosxe_restconf.PreReq0, ]
	}
	`
}
