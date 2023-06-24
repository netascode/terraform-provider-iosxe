//go:build testAll

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIosxeInterfaceVLAN(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxeInterfaceVLANPrerequisitesConfig + testAccIosxeInterfaceVLANConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("iosxe_interface_vlan.test", "name", "10"),
					resource.TestCheckResourceAttr("iosxe_interface_vlan.test", "autostate", "false"),
					resource.TestCheckResourceAttr("iosxe_interface_vlan.test", "description", "My Interface Description"),
					resource.TestCheckResourceAttr("iosxe_interface_vlan.test", "shutdown", "false"),
					resource.TestCheckResourceAttr("iosxe_interface_vlan.test", "ip_proxy_arp", "false"),
					resource.TestCheckResourceAttr("iosxe_interface_vlan.test", "ip_redirects", "false"),
					resource.TestCheckResourceAttr("iosxe_interface_vlan.test", "unreachables", "false"),
					resource.TestCheckResourceAttr("iosxe_interface_vlan.test", "vrf_forwarding", "VRF1"),
					resource.TestCheckResourceAttr("iosxe_interface_vlan.test", "ipv4_address", "10.1.1.1"),
					resource.TestCheckResourceAttr("iosxe_interface_vlan.test", "ipv4_address_mask", "255.255.255.0"),
					resource.TestCheckResourceAttr("iosxe_interface_vlan.test", "ip_dhcp_relay_source_interface", "Loopback100"),
					resource.TestCheckResourceAttr("iosxe_interface_vlan.test", "ip_access_group_in", "1"),
					resource.TestCheckResourceAttr("iosxe_interface_vlan.test", "ip_access_group_in_enable", "true"),
					resource.TestCheckResourceAttr("iosxe_interface_vlan.test", "ip_access_group_out", "1"),
					resource.TestCheckResourceAttr("iosxe_interface_vlan.test", "ip_access_group_out_enable", "true"),
					resource.TestCheckResourceAttr("iosxe_interface_vlan.test", "helper_addresses.0.address", "10.10.10.10"),
					resource.TestCheckResourceAttr("iosxe_interface_vlan.test", "helper_addresses.0.global", "false"),
					resource.TestCheckResourceAttr("iosxe_interface_vlan.test", "helper_addresses.0.vrf", "VRF1"),
				),
			},
			{
				ResourceName:  "iosxe_interface_vlan.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XE-native:native/interface/Vlan=10",
			},
		},
	})
}

const testAccIosxeInterfaceVLANPrerequisitesConfig = `
resource "iosxe_restconf" "PreReq0" {
  path = "Cisco-IOS-XE-native:native/vrf/definition=VRF1"
  delete = false
  attributes = {
      "name" = "VRF1"
      "address-family/ipv4" = ""
  }
}

`

func testAccIosxeInterfaceVLANConfig_minimum() string {
	return `
	resource "iosxe_interface_vlan" "test" {
		name = 10
  		depends_on = [iosxe_restconf.PreReq0, ]
	}
	`
}

func testAccIosxeInterfaceVLANConfig_all() string {
	return `
	resource "iosxe_interface_vlan" "test" {
		name = 10
		autostate = false
		description = "My Interface Description"
		shutdown = false
		ip_proxy_arp = false
		ip_redirects = false
		unreachables = false
		vrf_forwarding = "VRF1"
		ipv4_address = "10.1.1.1"
		ipv4_address_mask = "255.255.255.0"
		ip_dhcp_relay_source_interface = "Loopback100"
		ip_access_group_in = "1"
		ip_access_group_in_enable = true
		ip_access_group_out = "1"
		ip_access_group_out_enable = true
		helper_addresses = [{
			address = "10.10.10.10"
			global = false
			vrf = "VRF1"
		}]
  		depends_on = [iosxe_restconf.PreReq0, ]
	}
	`
}
