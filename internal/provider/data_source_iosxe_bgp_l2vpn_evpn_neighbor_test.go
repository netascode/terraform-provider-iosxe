// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceIosxeBGPL2VPNEVPNNeighbor(t *testing.T) {
	if os.Getenv("C9000V") == "" {
		t.Skip("skipping test, set environment variable C9000V")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_bgp_l2vpn_evpn_neighbor.test", "activate", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_bgp_l2vpn_evpn_neighbor.test", "send_community", "both"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_bgp_l2vpn_evpn_neighbor.test", "route_reflector_client", "false"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxeBGPL2VPNEVPNNeighborPrerequisitesConfig + testAccDataSourceIosxeBGPL2VPNEVPNNeighborConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

const testAccDataSourceIosxeBGPL2VPNEVPNNeighborPrerequisitesConfig = `
resource "iosxe_restconf" "PreReq0" {
	path = "Cisco-IOS-XE-native:native/router/Cisco-IOS-XE-bgp:bgp=65000"
	attributes = {
		"id" = "65000"
	}
}

resource "iosxe_restconf" "PreReq1" {
	path = "Cisco-IOS-XE-native:native/router/Cisco-IOS-XE-bgp:bgp=65000/address-family/no-vrf/l2vpn=evpn"
	attributes = {
		"af-name" = "evpn"
	}
	depends_on = [iosxe_restconf.PreReq0, ]
}

resource "iosxe_restconf" "PreReq2" {
	path = "Cisco-IOS-XE-native:native/router/Cisco-IOS-XE-bgp:bgp=65000/neighbor=3.3.3.3"
	attributes = {
		"id" = "3.3.3.3"
		"remote-as" = "65000"
	}
	depends_on = [iosxe_restconf.PreReq0, ]
}

`

func testAccDataSourceIosxeBGPL2VPNEVPNNeighborConfig() string {
	config := `resource "iosxe_bgp_l2vpn_evpn_neighbor" "test" {` + "\n"
	config += `	delete_mode = "attributes"\n`
	config += `	asn = "65000"` + "\n"
	config += `	ip = "3.3.3.3"` + "\n"
	config += `	activate = true` + "\n"
	config += `	send_community = "both"` + "\n"
	config += `	route_reflector_client = false` + "\n"
	config += `	depends_on = [iosxe_restconf.PreReq0, iosxe_restconf.PreReq1, iosxe_restconf.PreReq2, ]` + "\n"
	config += `}` + "\n"

	config += `
		data "iosxe_bgp_l2vpn_evpn_neighbor" "test" {
			asn = "65000"
			ip = "3.3.3.3"
			depends_on = [iosxe_bgp_l2vpn_evpn_neighbor.test]
		}
	`
	return config
}
