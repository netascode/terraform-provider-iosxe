// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceIosxeInterfaceOSPFProcess(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_ospf_process.test", "area.0.area_id", "1"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxeInterfaceOSPFProcessPrerequisitesConfig + testAccDataSourceIosxeInterfaceOSPFProcessConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

const testAccDataSourceIosxeInterfaceOSPFProcessPrerequisitesConfig = `
resource "iosxe_restconf" "PreReq0" {
	path = "Cisco-IOS-XE-native:native/router/Cisco-IOS-XE-ospf:router-ospf/ospf/process-id=1"
	attributes = {
		"id" = "1"
	}
}

`

func testAccDataSourceIosxeInterfaceOSPFProcessConfig() string {
	config := `resource "iosxe_interface_ospf_process" "test" {` + "\n"
	config += `	type = "GigabitEthernet"` + "\n"
	config += `	name = "2"` + "\n"
	config += `	process_id = 1` + "\n"
	config += `	area = [{` + "\n"
	config += `		area_id = "1"` + "\n"
	config += `	}]` + "\n"
	config += `	depends_on = [iosxe_restconf.PreReq0, ]` + "\n"
	config += `}` + "\n"

	config += `
		data "iosxe_interface_ospf_process" "test" {
			type = "GigabitEthernet"
			name = "2"
			process_id = 1
			depends_on = [iosxe_interface_ospf_process.test]
		}
	`
	return config
}
