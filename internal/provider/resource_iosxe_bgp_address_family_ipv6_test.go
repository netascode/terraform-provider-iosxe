// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIosxeBGPAddressFamilyIPv6(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxeBGPAddressFamilyIPv6PrerequisitesConfig + testAccIosxeBGPAddressFamilyIPv6Config_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("iosxe_bgp_address_family_ipv6.test", "af_name", "unicast"),
				),
			},
			{
				ResourceName:  "iosxe_bgp_address_family_ipv6.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XE-native:native/router/Cisco-IOS-XE-bgp:bgp=65000/address-family/no-vrf/ipv6=unicast",
			},
		},
	})
}

const testAccIosxeBGPAddressFamilyIPv6PrerequisitesConfig = `
resource "iosxe_restconf" "PreReq0" {
	path = "Cisco-IOS-XE-native:native/ipv6"
	attributes = {
		"unicast-routing" = ""
	}
}

resource "iosxe_restconf" "PreReq1" {
	path = "Cisco-IOS-XE-native:native/router/Cisco-IOS-XE-bgp:bgp=65000"
	attributes = {
		"id" = "65000"
	}
}

`

func testAccIosxeBGPAddressFamilyIPv6Config_minimum() string {
	return `
	resource "iosxe_bgp_address_family_ipv6" "test" {
		asn = "65000"
		af_name = "unicast"
		depends_on = [iosxe_restconf.PreReq0, iosxe_restconf.PreReq1, ]
	}
	`
}

func testAccIosxeBGPAddressFamilyIPv6Config_all() string {
	return `
	resource "iosxe_bgp_address_family_ipv6" "test" {
		asn = "65000"
		af_name = "unicast"
		depends_on = [iosxe_restconf.PreReq0, iosxe_restconf.PreReq1, ]
	}
	`
}
