// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceIosxeBGPAddressFamilyIPv6VRF(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_bgp_address_family_ipv6_vrf.test", "vrfs.0.name", "VRF1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_bgp_address_family_ipv6_vrf.test", "vrfs.0.advertise_l2vpn_evpn", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_bgp_address_family_ipv6_vrf.test", "vrfs.0.redistribute_connected", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_bgp_address_family_ipv6_vrf.test", "vrfs.0.redistribute_static", "true"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxeBGPAddressFamilyIPv6VRFPrerequisitesConfig + testAccDataSourceIosxeBGPAddressFamilyIPv6VRFConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

const testAccDataSourceIosxeBGPAddressFamilyIPv6VRFPrerequisitesConfig = `
resource "iosxe_restconf" "PreReq0" {
	path = "Cisco-IOS-XE-native:native/ipv6"
	attributes = {
		"unicast-routing" = ""
	}
}

resource "iosxe_restconf" "PreReq1" {
	path = "Cisco-IOS-XE-native:native/vrf/definition=VRF1"
	delete = false
	attributes = {
		"name" = "VRF1"
		"rd" = "1:1"
		"address-family/ipv6" = ""
	}
	depends_on = [iosxe_restconf.PreReq0, ]
}

resource "iosxe_restconf" "PreReq2" {
	path = "Cisco-IOS-XE-native:native/router/Cisco-IOS-XE-bgp:bgp=65000"
	attributes = {
		"id" = "65000"
	}
	depends_on = [iosxe_restconf.PreReq1, ]
}

`

func testAccDataSourceIosxeBGPAddressFamilyIPv6VRFConfig() string {
	config := `resource "iosxe_bgp_address_family_ipv6_vrf" "test" {` + "\n"
	config += `	delete_mode = "attributes"\n`
	config += `	asn = "65000"` + "\n"
	config += `	af_name = "unicast"` + "\n"
	config += `	vrfs = [{` + "\n"
	config += `		name = "VRF1"` + "\n"
	config += `		advertise_l2vpn_evpn = true` + "\n"
	config += `		redistribute_connected = true` + "\n"
	config += `		redistribute_static = true` + "\n"
	config += `	}]` + "\n"
	config += `	depends_on = [iosxe_restconf.PreReq0, iosxe_restconf.PreReq1, iosxe_restconf.PreReq2, ]` + "\n"
	config += `}` + "\n"

	config += `
		data "iosxe_bgp_address_family_ipv6_vrf" "test" {
			asn = "65000"
			af_name = "unicast"
			depends_on = [iosxe_bgp_address_family_ipv6_vrf.test]
		}
	`
	return config
}
