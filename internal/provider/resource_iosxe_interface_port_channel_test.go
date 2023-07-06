// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIosxeInterfacePortChannel(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_port_channel.test", "name", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_port_channel.test", "description", "My Interface Description"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_port_channel.test", "shutdown", "false"))
	if os.Getenv("C8000V") != "" {
		checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_port_channel.test", "ip_proxy_arp", "false"))
	}
	if os.Getenv("C8000V") != "" {
		checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_port_channel.test", "ip_redirects", "false"))
	}
	if os.Getenv("C8000V") != "" {
		checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_port_channel.test", "unreachables", "false"))
	}
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_port_channel.test", "vrf_forwarding", "VRF1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_port_channel.test", "ipv4_address", "192.0.2.1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_port_channel.test", "ipv4_address_mask", "255.255.255.0"))
	if os.Getenv("C9000V") != "" {
		checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_port_channel.test", "switchport", "false"))
	}
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_port_channel.test", "ip_access_group_in", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_port_channel.test", "ip_access_group_in_enable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_port_channel.test", "ip_access_group_out", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_port_channel.test", "ip_access_group_out_enable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_port_channel.test", "ip_dhcp_relay_source_interface", "Loopback100"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_port_channel.test", "helper_addresses.0.address", "10.10.10.10"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_port_channel.test", "helper_addresses.0.global", "false"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxeInterfacePortChannelPrerequisitesConfig + testAccIosxeInterfacePortChannelConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				ResourceName:  "iosxe_interface_port_channel.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XE-native:native/interface/Port-channel=10",
			},
		},
	})
}

const testAccIosxeInterfacePortChannelPrerequisitesConfig = `
resource "iosxe_restconf" "PreReq0" {
	path = "Cisco-IOS-XE-native:native/vrf/definition=VRF1"
	delete = false
	attributes = {
		"name" = "VRF1"
		"address-family/ipv4" = ""
	}
}

`

func testAccIosxeInterfacePortChannelConfig_minimum() string {
	config := `resource "iosxe_interface_port_channel" "test" {` + "\n"
	config += `	name = 10` + "\n"
	config += `	depends_on = [iosxe_restconf.PreReq0, ]` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxeInterfacePortChannelConfig_all() string {
	config := `resource "iosxe_interface_port_channel" "test" {` + "\n"
	config += `	name = 10` + "\n"
	config += `	description = "My Interface Description"` + "\n"
	config += `	shutdown = false` + "\n"
	if os.Getenv("C8000V") != "" {
		config += `	ip_proxy_arp = false` + "\n"
	}
	if os.Getenv("C8000V") != "" {
		config += `	ip_redirects = false` + "\n"
	}
	if os.Getenv("C8000V") != "" {
		config += `	unreachables = false` + "\n"
	}
	config += `	vrf_forwarding = "VRF1"` + "\n"
	config += `	ipv4_address = "192.0.2.1"` + "\n"
	config += `	ipv4_address_mask = "255.255.255.0"` + "\n"
	if os.Getenv("C9000V") != "" {
		config += `	switchport = false` + "\n"
	}
	config += `	ip_access_group_in = "1"` + "\n"
	config += `	ip_access_group_in_enable = true` + "\n"
	config += `	ip_access_group_out = "1"` + "\n"
	config += `	ip_access_group_out_enable = true` + "\n"
	config += `	ip_dhcp_relay_source_interface = "Loopback100"` + "\n"
	config += `	helper_addresses = [{` + "\n"
	config += `		address = "10.10.10.10"` + "\n"
	config += `		global = false` + "\n"
	config += `	}]` + "\n"
	config += `	depends_on = [iosxe_restconf.PreReq0, ]` + "\n"
	config += `}` + "\n"
	return config
}
