//go:build testAll

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIosxeInterfaceNVE(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxeInterfaceNVEConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("iosxe_interface_nve.test", "name", "1"),
					resource.TestCheckResourceAttr("iosxe_interface_nve.test", "description", "My Interface Description"),
					resource.TestCheckResourceAttr("iosxe_interface_nve.test", "shutdown", "false"),
					resource.TestCheckResourceAttr("iosxe_interface_nve.test", "host_reachability_protocol_bgp", "true"),
					resource.TestCheckResourceAttr("iosxe_interface_nve.test", "source_interface_loopback", "100"),
					resource.TestCheckResourceAttr("iosxe_interface_nve.test", "vnis.0.vni_range", "10000"),
					resource.TestCheckResourceAttr("iosxe_interface_nve.test", "vnis.0.ipv4_multicast_group", "225.1.1.1"),
				),
			},
			{
				ResourceName:  "iosxe_interface_nve.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XE-native:native/interface/nve=1",
			},
		},
	})
}

func testAccIosxeInterfaceNVEConfig_minimum() string {
	return `
	resource "iosxe_interface_nve" "test" {
		name = 1
	}
	`
}

func testAccIosxeInterfaceNVEConfig_all() string {
	return `
	resource "iosxe_interface_nve" "test" {
		name = 1
		description = "My Interface Description"
		shutdown = false
		host_reachability_protocol_bgp = true
		source_interface_loopback = 100
		vnis = [{
		vni_range = "10000"
		ipv4_multicast_group = "225.1.1.1"
		}]
	}
	`
}