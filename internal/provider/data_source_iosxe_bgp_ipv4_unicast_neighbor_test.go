// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceIosxeBGPIPv4UnicastNeighbor(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxeBGPIPv4UnicastNeighborPrerequisitesConfig + testAccDataSourceIosxeBGPIPv4UnicastNeighborConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.iosxe_bgp_ipv4_unicast_neighbor.test", "activate", "true"),
					resource.TestCheckResourceAttr("data.iosxe_bgp_ipv4_unicast_neighbor.test", "send_community", "both"),
					resource.TestCheckResourceAttr("data.iosxe_bgp_ipv4_unicast_neighbor.test", "route_reflector_client", "false"),
					resource.TestCheckResourceAttr("data.iosxe_bgp_ipv4_unicast_neighbor.test", "route_maps.0.in_out", "in"),
					resource.TestCheckResourceAttr("data.iosxe_bgp_ipv4_unicast_neighbor.test", "route_maps.0.route_map_name", "RM1"),
				),
			},
		},
	})
}

const testAccDataSourceIosxeBGPIPv4UnicastNeighborPrerequisitesConfig = `
resource "iosxe_restconf" "PreReq0" {
  path = "Cisco-IOS-XE-native:native/router/Cisco-IOS-XE-bgp:bgp=65000"
  attributes = {
      "id" = "65000"
  }
}

resource "iosxe_restconf" "PreReq1" {
  path = "Cisco-IOS-XE-native:native/router/Cisco-IOS-XE-bgp:bgp=65000/address-family/no-vrf/ipv4=unicast"
  attributes = {
      "af-name" = "unicast"
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

resource "iosxe_restconf" "PreReq3" {
  path = "Cisco-IOS-XE-native:native/interface/Loopback=100"
  attributes = {
      "name" = "100"
  }
}

`

const testAccDataSourceIosxeBGPIPv4UnicastNeighborConfig = `

resource "iosxe_bgp_ipv4_unicast_neighbor" "test" {
  asn = "65000"
  ip = "3.3.3.3"
  activate = true
  send_community = "both"
  route_reflector_client = false
  route_maps = [{
    in_out = "in"
    route_map_name = "RM1"
  }]
  depends_on = [iosxe_restconf.PreReq0, iosxe_restconf.PreReq1, iosxe_restconf.PreReq2, iosxe_restconf.PreReq3, ]
}

data "iosxe_bgp_ipv4_unicast_neighbor" "test" {
  asn = "65000"
  ip = "3.3.3.3"
  depends_on = [iosxe_bgp_ipv4_unicast_neighbor.test]
}
`
