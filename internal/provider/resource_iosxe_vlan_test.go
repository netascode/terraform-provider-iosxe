//go:build testAll

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIosxeVLAN(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxeVLANConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("iosxe_vlan.test", "vlan_id", "123"),
					resource.TestCheckResourceAttr("iosxe_vlan.test", "name", "Vlan123"),
					resource.TestCheckResourceAttr("iosxe_vlan.test", "shutdown", "false"),
				),
			},
			{
				ResourceName:  "iosxe_vlan.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XE-native:native/vlan/ios-vlan:vlan-list=123",
			},
		},
	})
}

func testAccIosxeVLANConfig_minimum() string {
	return `
	resource "iosxe_vlan" "test" {
		vlan_id = 123
	}
	`
}

func testAccIosxeVLANConfig_all() string {
	return `
	resource "iosxe_vlan" "test" {
		vlan_id = 123
		name = "Vlan123"
		shutdown = false
	}
	`
}
