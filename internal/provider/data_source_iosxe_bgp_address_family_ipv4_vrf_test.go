// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceIosxeBGPAddressFamilyIPv4VRF(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxeBGPAddressFamilyIPv4VRFPrerequisitesConfig + testAccDataSourceIosxeBGPAddressFamilyIPv4VRFConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.iosxe_bgp_address_family_ipv4_vrf.test", "vrfs.0.name", "VRF1"),
					resource.TestCheckResourceAttr("data.iosxe_bgp_address_family_ipv4_vrf.test", "vrfs.0.advertise_l2vpn_evpn", "true"),
					resource.TestCheckResourceAttr("data.iosxe_bgp_address_family_ipv4_vrf.test", "vrfs.0.redistribute_connected", "true"),
					resource.TestCheckResourceAttr("data.iosxe_bgp_address_family_ipv4_vrf.test", "vrfs.0.redistribute_static", "true"),
				),
			},
		},
	})
}

const testAccDataSourceIosxeBGPAddressFamilyIPv4VRFPrerequisitesConfig = `
resource "iosxe_restconf" "PreReq0" {
  path = "Cisco-IOS-XE-native:native/vrf/definition=VRF1"
  delete = false
  attributes = {
      name = "VRF1"
      rd = "1:1"
  }
}

resource "iosxe_restconf" "PreReq1" {
  path = "Cisco-IOS-XE-native:native/vrf/definition=VRF1/address-family"
  delete = false
  attributes = {
      ipv4 = ""
  }
  depends_on = [iosxe_restconf.PreReq0, ]
}

resource "iosxe_restconf" "PreReq2" {
  path = "Cisco-IOS-XE-native:native/router/Cisco-IOS-XE-bgp:bgp=65000"
  attributes = {
      id = "65000"
  }
  depends_on = [iosxe_restconf.PreReq1, ]
}

`

const testAccDataSourceIosxeBGPAddressFamilyIPv4VRFConfig = `

resource "iosxe_bgp_address_family_ipv4_vrf" "test" {
  asn = "65000"
  af_name = "unicast"
  vrfs = [{
    name = "VRF1"
    advertise_l2vpn_evpn = true
    redistribute_connected = true
    redistribute_static = true
  }]
  depends_on = [iosxe_restconf.PreReq0, iosxe_restconf.PreReq1, iosxe_restconf.PreReq2, ]
}

data "iosxe_bgp_address_family_ipv4_vrf" "test" {
  asn = "65000"
  af_name = "unicast"
  depends_on = [iosxe_bgp_address_family_ipv4_vrf.test]
}
`
