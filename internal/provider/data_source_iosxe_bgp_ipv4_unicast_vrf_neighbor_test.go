// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceIosxeBGPIPv4UnicastVRFNeighbor(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxeBGPIPv4UnicastVRFNeighborPrerequisitesConfig + testAccDataSourceIosxeBGPIPv4UnicastVRFNeighborConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.iosxe_bgp_ipv4_unicast_vrf_neighbor.test", "remote_as", "65000"),
					resource.TestCheckResourceAttr("data.iosxe_bgp_ipv4_unicast_vrf_neighbor.test", "description", "BGP Neighbor Description"),
					resource.TestCheckResourceAttr("data.iosxe_bgp_ipv4_unicast_vrf_neighbor.test", "shutdown", "false"),
					resource.TestCheckResourceAttr("data.iosxe_bgp_ipv4_unicast_vrf_neighbor.test", "update_source_loopback", "100"),
					resource.TestCheckResourceAttr("data.iosxe_bgp_ipv4_unicast_vrf_neighbor.test", "activate", "true"),
					resource.TestCheckResourceAttr("data.iosxe_bgp_ipv4_unicast_vrf_neighbor.test", "send_community", "both"),
					resource.TestCheckResourceAttr("data.iosxe_bgp_ipv4_unicast_vrf_neighbor.test", "route_reflector_client", "false"),
					resource.TestCheckResourceAttr("data.iosxe_bgp_ipv4_unicast_vrf_neighbor.test", "route_maps.0.in_out", "in"),
					resource.TestCheckResourceAttr("data.iosxe_bgp_ipv4_unicast_vrf_neighbor.test", "route_maps.0.route_map_name", "RM1"),
				),
			},
		},
	})
}

const testAccDataSourceIosxeBGPIPv4UnicastVRFNeighborPrerequisitesConfig = `
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
}

resource "iosxe_restconf" "PreReq2" {
  path = "Cisco-IOS-XE-native:native/router/Cisco-IOS-XE-bgp:bgp=65000/address-family/with-vrf/ipv4=unicast"
  attributes = {
      "af-name" = "unicast"
  }
  lists = [
    {
      name = "vrf"
      key = "name"
      items = [
          {
            attributes = {
              "name" = "VRF1"
            }
          },
      ] 
    },
  ]
  depends_on = [iosxe_restconf.PreReq0, iosxe_restconf.PreReq1, ]
}

resource "iosxe_restconf" "PreReq3" {
  path = "Cisco-IOS-XE-native:native/interface/Loopback=100"
  attributes = {
      "name" = "100"
  }
}

`

const testAccDataSourceIosxeBGPIPv4UnicastVRFNeighborConfig = `

resource "iosxe_bgp_ipv4_unicast_vrf_neighbor" "test" {
  asn = "65000"
  vrf = "VRF1"
  ip = "3.3.3.3"
  remote_as = "65000"
  description = "BGP Neighbor Description"
  shutdown = false
  update_source_loopback = "100"
  activate = true
  send_community = "both"
  route_reflector_client = false
  route_maps = [{
    in_out = "in"
    route_map_name = "RM1"
  }]
  depends_on = [iosxe_restconf.PreReq0, iosxe_restconf.PreReq1, iosxe_restconf.PreReq2, iosxe_restconf.PreReq3, ]
}

data "iosxe_bgp_ipv4_unicast_vrf_neighbor" "test" {
  asn = "65000"
  vrf = "VRF1"
  ip = "3.3.3.3"
  depends_on = [iosxe_bgp_ipv4_unicast_vrf_neighbor.test]
}
`
