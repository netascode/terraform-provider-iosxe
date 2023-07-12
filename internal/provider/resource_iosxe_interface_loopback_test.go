// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIosxeInterfaceLoopback(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_loopback.test", "name", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_loopback.test", "description", "My Interface Description"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_loopback.test", "shutdown", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_loopback.test", "ip_proxy_arp", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_loopback.test", "ip_redirects", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_loopback.test", "unreachables", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_loopback.test", "vrf_forwarding", "VRF1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_loopback.test", "ipv4_address", "200.1.1.1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_loopback.test", "ipv4_address_mask", "255.255.255.255"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_loopback.test", "ip_access_group_in", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_loopback.test", "ip_access_group_in_enable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_loopback.test", "ip_access_group_out", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_loopback.test", "ip_access_group_out_enable", "true"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxeInterfaceLoopbackPrerequisitesConfig + testAccIosxeInterfaceLoopbackConfig_minimum(),
			},
			{
				Config: testAccIosxeInterfaceLoopbackPrerequisitesConfig + testAccIosxeInterfaceLoopbackConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				ResourceName:  "iosxe_interface_loopback.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XE-native:native/interface/Loopback=100",
			},
		},
	})
}

const testAccIosxeInterfaceLoopbackPrerequisitesConfig = `
resource "iosxe_restconf" "PreReq0" {
	path = "Cisco-IOS-XE-native:native/vrf/definition=VRF1"
	delete = false
	attributes = {
		"name" = "VRF1"
		"address-family/ipv4" = ""
	}
}

`

func testAccIosxeInterfaceLoopbackConfig_minimum() string {
	config := `resource "iosxe_interface_loopback" "test" {` + "\n"
	config += `	name = 100` + "\n"
	config += `	depends_on = [iosxe_restconf.PreReq0, ]` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxeInterfaceLoopbackConfig_all() string {
	config := `resource "iosxe_interface_loopback" "test" {` + "\n"
	config += `	name = 100` + "\n"
	config += `	description = "My Interface Description"` + "\n"
	config += `	shutdown = false` + "\n"
	config += `	ip_proxy_arp = false` + "\n"
	config += `	ip_redirects = false` + "\n"
	config += `	unreachables = false` + "\n"
	config += `	vrf_forwarding = "VRF1"` + "\n"
	config += `	ipv4_address = "200.1.1.1"` + "\n"
	config += `	ipv4_address_mask = "255.255.255.255"` + "\n"
	config += `	ip_access_group_in = "1"` + "\n"
	config += `	ip_access_group_in_enable = true` + "\n"
	config += `	ip_access_group_out = "1"` + "\n"
	config += `	ip_access_group_out_enable = true` + "\n"
	config += `	depends_on = [iosxe_restconf.PreReq0, ]` + "\n"
	config += `}` + "\n"
	return config
}
