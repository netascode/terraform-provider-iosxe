// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIosxeInterfaceOSPFProcess(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_ospf_process.test", "process_id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_ospf_process.test", "area.0.area_id", "1"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxeInterfaceOSPFProcessPrerequisitesConfig + testAccIosxeInterfaceOSPFProcessConfig_minimum(),
			},
			{
				Config: testAccIosxeInterfaceOSPFProcessPrerequisitesConfig + testAccIosxeInterfaceOSPFProcessConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				ResourceName:  "iosxe_interface_ospf_process.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XE-native:native/interface/Loopback=1/ip/Cisco-IOS-XE-ospf:router-ospf/ospf/process-id=1",
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

resource "iosxe_restconf" "PreReq1" {
	path = "Cisco-IOS-XE-native:native/interface/Loopback=1"
	attributes = {
		"name" = "1"
	}
}

`

func testAccIosxeInterfaceOSPFProcessConfig_minimum() string {
	config := `resource "iosxe_interface_ospf_process" "test" {` + "\n"
	config += `	type = "Loopback"` + "\n"
	config += `	name = "1"` + "\n"
	config += `	process_id = 1` + "\n"
	config += `	depends_on = [iosxe_restconf.PreReq0, iosxe_restconf.PreReq1, ]` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxeInterfaceOSPFProcessConfig_all() string {
	config := `resource "iosxe_interface_ospf_process" "test" {` + "\n"
	config += `	type = "Loopback"` + "\n"
	config += `	name = "1"` + "\n"
	config += `	process_id = 1` + "\n"
	config += `	area = [{` + "\n"
	config += `		area_id = "1"` + "\n"
	config += `	}]` + "\n"
	config += `	depends_on = [iosxe_restconf.PreReq0, iosxe_restconf.PreReq1, ]` + "\n"
	config += `}` + "\n"
	return config
}
