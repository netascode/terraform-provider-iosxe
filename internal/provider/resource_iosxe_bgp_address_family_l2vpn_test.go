//go:build testAll

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIosxeBGPAddressFamilyL2VPN(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxeBGPAddressFamilyL2VPNPrerequisitesConfig + testAccIosxeBGPAddressFamilyL2VPNConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("iosxe_bgp_address_family_l2vpn.test", "af_name", "evpn"),
				),
			},
			{
				ResourceName:  "iosxe_bgp_address_family_l2vpn.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XE-native:native/router/Cisco-IOS-XE-bgp:bgp=65000/address-family/no-vrf/l2vpn=evpn",
			},
		},
	})
}

const testAccIosxeBGPAddressFamilyL2VPNPrerequisitesConfig = `
resource "iosxe_restconf" "PreReq0" {
  path = "Cisco-IOS-XE-native:native/router/Cisco-IOS-XE-bgp:bgp=65000"
  attributes = {
      "id" = "65000"
  }
}

`

func testAccIosxeBGPAddressFamilyL2VPNConfig_minimum() string {
	return `
	resource "iosxe_bgp_address_family_l2vpn" "test" {
		asn = "65000"
		af_name = "evpn"
  		depends_on = [iosxe_restconf.PreReq0, ]
	}
	`
}

func testAccIosxeBGPAddressFamilyL2VPNConfig_all() string {
	return `
	resource "iosxe_bgp_address_family_l2vpn" "test" {
		asn = "65000"
		af_name = "evpn"
  		depends_on = [iosxe_restconf.PreReq0, ]
	}
	`
}
