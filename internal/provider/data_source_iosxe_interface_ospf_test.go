// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceIosxeInterfaceOSPF(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_ospf.test", "cost", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_ospf.test", "dead_interval", "30"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_ospf.test", "hello_interval", "5"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_ospf.test", "mtu_ignore", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_ospf.test", "network_type_broadcast", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_ospf.test", "network_type_non_broadcast", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_ospf.test", "network_type_point_to_multipoint", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_ospf.test", "network_type_point_to_point", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_ospf.test", "priority", "10"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxeInterfaceOSPFPrerequisitesConfig + testAccDataSourceIosxeInterfaceOSPFConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

const testAccDataSourceIosxeInterfaceOSPFPrerequisitesConfig = `
resource "iosxe_restconf" "PreReq0" {
	path = "Cisco-IOS-XE-native:native/interface/Loopback=1"
	attributes = {
		"name" = "1"
	}
}

`

func testAccDataSourceIosxeInterfaceOSPFConfig() string {
	config := `resource "iosxe_interface_ospf" "test" {` + "\n"
	config += `	delete_mode = "attributes"` + "\n"
	config += `	type = "Loopback"` + "\n"
	config += `	name = "1"` + "\n"
	config += `	cost = 10` + "\n"
	config += `	dead_interval = 30` + "\n"
	config += `	hello_interval = 5` + "\n"
	config += `	mtu_ignore = false` + "\n"
	config += `	network_type_broadcast = false` + "\n"
	config += `	network_type_non_broadcast = false` + "\n"
	config += `	network_type_point_to_multipoint = false` + "\n"
	config += `	network_type_point_to_point = true` + "\n"
	config += `	priority = 10` + "\n"
	config += `	depends_on = [iosxe_restconf.PreReq0, ]` + "\n"
	config += `}` + "\n"

	config += `
		data "iosxe_interface_ospf" "test" {
			type = "Loopback"
			name = "1"
			depends_on = [iosxe_interface_ospf.test]
		}
	`
	return config
}
