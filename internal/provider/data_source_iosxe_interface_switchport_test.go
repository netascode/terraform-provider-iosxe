//go:build testAll

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceIosxeInterfaceSwitchport(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxeInterfaceSwitchportConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.iosxe_interface_switchport.test", "mode_access", "false"),
					resource.TestCheckResourceAttr("data.iosxe_interface_switchport.test", "mode_dot1q_tunnel", "false"),
					resource.TestCheckResourceAttr("data.iosxe_interface_switchport.test", "mode_private_vlan_trunk", "false"),
					resource.TestCheckResourceAttr("data.iosxe_interface_switchport.test", "mode_private_vlan_host", "false"),
					resource.TestCheckResourceAttr("data.iosxe_interface_switchport.test", "mode_private_vlan_promiscuous", "false"),
					resource.TestCheckResourceAttr("data.iosxe_interface_switchport.test", "mode_trunk", "true"),
					resource.TestCheckResourceAttr("data.iosxe_interface_switchport.test", "nonegotiate", "false"),
					resource.TestCheckResourceAttr("data.iosxe_interface_switchport.test", "access_vlan", "100"),
					resource.TestCheckResourceAttr("data.iosxe_interface_switchport.test", "trunk_allowed_vlans", "100,101"),
					resource.TestCheckResourceAttr("data.iosxe_interface_switchport.test", "trunk_native_vlan_tag", "false"),
					resource.TestCheckResourceAttr("data.iosxe_interface_switchport.test", "trunk_native_vlan", "100"),
					resource.TestCheckResourceAttr("data.iosxe_interface_switchport.test", "host", "false"),
				),
			},
		},
	})
}

const testAccDataSourceIosxeInterfaceSwitchportConfig = `

resource "iosxe_interface_switchport" "test" {
  type = "GigabitEthernet"
  name = "1/0/3"
  mode_access = false
  mode_dot1q_tunnel = false
  mode_private_vlan_trunk = false
  mode_private_vlan_host = false
  mode_private_vlan_promiscuous = false
  mode_trunk = true
  nonegotiate = false
  access_vlan = "100"
  trunk_allowed_vlans = "100,101"
  trunk_native_vlan_tag = false
  trunk_native_vlan = 100
  host = false
}

data "iosxe_interface_switchport" "test" {
  type = "GigabitEthernet"
  name = "1/0/3"
  depends_on = [iosxe_interface_switchport.test]
}
`
