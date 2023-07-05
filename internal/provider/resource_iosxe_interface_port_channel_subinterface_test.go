// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIosxeInterfacePortChannelSubinterface(t *testing.T) {
	if os.Getenv("C8000V") == "" {
		t.Skip("skipping test, set environment variable C8000V")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_port_channel_subinterface.test", "name", "10.666"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_port_channel_subinterface.test", "description", "My Interface Description"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_port_channel_subinterface.test", "shutdown", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_port_channel_subinterface.test", "ip_proxy_arp", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_port_channel_subinterface.test", "ip_redirects", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_port_channel_subinterface.test", "unreachables", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_port_channel_subinterface.test", "vrf_forwarding", "VRF1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_port_channel_subinterface.test", "ipv4_address", "192.0.2.2"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_port_channel_subinterface.test", "ipv4_address_mask", "255.255.255.0"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_port_channel_subinterface.test", "encapsulation_dot1q_vlan_id", "666"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_port_channel_subinterface.test", "ip_access_group_in", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_port_channel_subinterface.test", "ip_access_group_in_enable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_port_channel_subinterface.test", "ip_access_group_out", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_port_channel_subinterface.test", "ip_access_group_out_enable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_port_channel_subinterface.test", "helper_addresses.0.address", "10.10.10.10"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_port_channel_subinterface.test", "helper_addresses.0.global", "false"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxeInterfacePortChannelSubinterfacePrerequisitesConfig + testAccIosxeInterfacePortChannelSubinterfaceConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				ResourceName:  "iosxe_interface_port_channel_subinterface.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XE-native:native/interface/Port-channel-subinterface/Port-channel=10.666",
			},
		},
	})
}

const testAccIosxeInterfacePortChannelSubinterfacePrerequisitesConfig = `
resource "iosxe_restconf" "PreReq0" {
	path = "Cisco-IOS-XE-native:native/vrf/definition=VRF1"
	delete = false
	attributes = {
		"name" = "VRF1"
		"address-family/ipv4" = ""
	}
}

resource "iosxe_restconf" "PreReq1" {
	path = "Cisco-IOS-XE-native:native/interface/Port-channel=10"
	attributes = {
		"name" = "10"
	}
}

resource "iosxe_restconf" "PreReq2" {
	path = "Cisco-IOS-XE-native:native/interface/Port-channel-subinterface/Port-channel=10.666"
	attributes = {
		"name" = "10.666"
	}
	depends_on = [iosxe_restconf.PreReq1, ]
}

`

func testAccIosxeInterfacePortChannelSubinterfaceConfig_minimum() string {
	config := `resource "iosxe_interface_port_channel_subinterface" "test" {` + "\n"
	config += `	name = "10.666"` + "\n"
	config += `	depends_on = [iosxe_restconf.PreReq0, iosxe_restconf.PreReq1, iosxe_restconf.PreReq2, ]` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxeInterfacePortChannelSubinterfaceConfig_all() string {
	config := `resource "iosxe_interface_port_channel_subinterface" "test" {` + "\n"
	config += `	name = "10.666"` + "\n"
	config += `	description = "My Interface Description"` + "\n"
	config += `	shutdown = false` + "\n"
	config += `	ip_proxy_arp = false` + "\n"
	config += `	ip_redirects = false` + "\n"
	config += `	unreachables = false` + "\n"
	config += `	vrf_forwarding = "VRF1"` + "\n"
	config += `	ipv4_address = "192.0.2.2"` + "\n"
	config += `	ipv4_address_mask = "255.255.255.0"` + "\n"
	config += `	encapsulation_dot1q_vlan_id = 666` + "\n"
	config += `	ip_access_group_in = "1"` + "\n"
	config += `	ip_access_group_in_enable = true` + "\n"
	config += `	ip_access_group_out = "1"` + "\n"
	config += `	ip_access_group_out_enable = true` + "\n"
	config += `	helper_addresses = [{` + "\n"
	config += `		address = "10.10.10.10"` + "\n"
	config += `		global = false` + "\n"
	config += `	}]` + "\n"
	config += `	depends_on = [iosxe_restconf.PreReq0, iosxe_restconf.PreReq1, iosxe_restconf.PreReq2, ]` + "\n"
	config += `}` + "\n"
	return config
}
