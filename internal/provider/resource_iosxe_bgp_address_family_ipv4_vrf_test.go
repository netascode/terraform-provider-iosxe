// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIosxeBGPAddressFamilyIPv4VRF(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxeBGPAddressFamilyIPv4VRFPrerequisitesConfig + testAccIosxeBGPAddressFamilyIPv4VRFConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("iosxe_bgp_address_family_ipv4_vrf.test", "af_name", "unicast"),
					resource.TestCheckResourceAttr("iosxe_bgp_address_family_ipv4_vrf.test", "vrfs.0.name", "VRF1"),
					resource.TestCheckResourceAttr("iosxe_bgp_address_family_ipv4_vrf.test", "vrfs.0.advertise_l2vpn_evpn", "true"),
					resource.TestCheckResourceAttr("iosxe_bgp_address_family_ipv4_vrf.test", "vrfs.0.redistribute_connected", "true"),
					resource.TestCheckResourceAttr("iosxe_bgp_address_family_ipv4_vrf.test", "vrfs.0.redistribute_static", "true"),
				),
			},
			{
				ResourceName:  "iosxe_bgp_address_family_ipv4_vrf.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XE-native:native/router/Cisco-IOS-XE-bgp:bgp=65000/address-family/with-vrf/ipv4=unicast",
			},
		},
	})
}

const testAccIosxeBGPAddressFamilyIPv4VRFPrerequisitesConfig = `
resource "iosxe_restconf" "PreReq0" {
  path = "Cisco-IOS-XE-native:native/vrf/definition=VRF1"
  delete = false
  attributes = {
      "name" = "VRF1"
      "rd" = "1:1"
      "address-family/ipv4" = ""
  }
}

resource "iosxe_restconf" "PreReq1" {
  path = "Cisco-IOS-XE-native:native/router/Cisco-IOS-XE-bgp:bgp=65000"
  attributes = {
      "id" = "65000"
  }
  depends_on = [iosxe_restconf.PreReq0, ]
}

`

func testAccIosxeBGPAddressFamilyIPv4VRFConfig_minimum() string {
	return `
	resource "iosxe_bgp_address_family_ipv4_vrf" "test" {
		asn = "65000"
		af_name = "unicast"
  		depends_on = [iosxe_restconf.PreReq0, iosxe_restconf.PreReq1, ]
	}
	`
}

func testAccIosxeBGPAddressFamilyIPv4VRFConfig_all() string {
	return `
	resource "iosxe_bgp_address_family_ipv4_vrf" "test" {
		asn = "65000"
		af_name = "unicast"
		vrfs = [{
			name = "VRF1"
			advertise_l2vpn_evpn = true
			redistribute_connected = true
			redistribute_static = true
		}]
  		depends_on = [iosxe_restconf.PreReq0, iosxe_restconf.PreReq1, ]
	}
	`
}
