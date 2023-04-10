// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIosxeInterfaceOSPFProcess(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxeInterfaceOSPFProcessPrerequisitesConfig + testAccIosxeInterfaceOSPFProcessConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("iosxe_interface_ospf_process.test", "process_id", "1"),
					resource.TestCheckResourceAttr("iosxe_interface_ospf_process.test", "area.0.area_id", "1"),
				),
			},
			{
				ResourceName:  "iosxe_interface_ospf_process.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XE-native:native/interface/GigabitEthernet=2/ip/Cisco-IOS-XE-ospf:router-ospf/ospf/process-id=1",
			},
		},
	})
}

const testAccIosxeInterfaceOSPFProcessPrerequisitesConfig = `
resource "iosxe_restconf" "PreReq0" {
  path = "Cisco-IOS-XE-native:native/router/Cisco-IOS-XE-ospf:router-ospf/ospf/process-id=1"
  attributes = {
      "id" = "1"
  }
}

`

func testAccIosxeInterfaceOSPFProcessConfig_minimum() string {
	return `
	resource "iosxe_interface_ospf_process" "test" {
		type = "GigabitEthernet"
		name = "2"
		process_id = 1
  		depends_on = [iosxe_restconf.PreReq0, ]
	}
	`
}

func testAccIosxeInterfaceOSPFProcessConfig_all() string {
	return `
	resource "iosxe_interface_ospf_process" "test" {
		type = "GigabitEthernet"
		name = "2"
		process_id = 1
		area = [{
		area_id = "1"
		}]
  		depends_on = [iosxe_restconf.PreReq0, ]
	}
	`
}
