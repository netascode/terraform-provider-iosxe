// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIosxeBGPIPv4UnicastNeighbor(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_bgp_ipv4_unicast_neighbor.test", "ip", "3.3.3.3"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_bgp_ipv4_unicast_neighbor.test", "activate", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_bgp_ipv4_unicast_neighbor.test", "send_community", "both"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_bgp_ipv4_unicast_neighbor.test", "route_reflector_client", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_bgp_ipv4_unicast_neighbor.test", "route_maps.0.in_out", "in"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_bgp_ipv4_unicast_neighbor.test", "route_maps.0.route_map_name", "RM1"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxeBGPIPv4UnicastNeighborPrerequisitesConfig + testAccIosxeBGPIPv4UnicastNeighborConfig_minimum(),
			},
			{
				Config: testAccIosxeBGPIPv4UnicastNeighborPrerequisitesConfig + testAccIosxeBGPIPv4UnicastNeighborConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				ResourceName:  "iosxe_bgp_ipv4_unicast_neighbor.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XE-native:native/router/Cisco-IOS-XE-bgp:bgp=65000/address-family/no-vrf/ipv4=unicast/ipv4-unicast/neighbor=3.3.3.3",
			},
		},
	})
}

const testAccIosxeBGPIPv4UnicastNeighborPrerequisitesConfig = `
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

func testAccIosxeBGPIPv4UnicastNeighborConfig_minimum() string {
	config := `resource "iosxe_bgp_ipv4_unicast_neighbor" "test" {` + "\n"
	config += `	asn = "65000"` + "\n"
	config += `	ip = "3.3.3.3"` + "\n"
	config += `	depends_on = [iosxe_restconf.PreReq0, iosxe_restconf.PreReq1, iosxe_restconf.PreReq2, iosxe_restconf.PreReq3, ]` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxeBGPIPv4UnicastNeighborConfig_all() string {
	config := `resource "iosxe_bgp_ipv4_unicast_neighbor" "test" {` + "\n"
	config += `	asn = "65000"` + "\n"
	config += `	ip = "3.3.3.3"` + "\n"
	config += `	activate = true` + "\n"
	config += `	send_community = "both"` + "\n"
	config += `	route_reflector_client = false` + "\n"
	config += `	route_maps = [{` + "\n"
	config += `		in_out = "in"` + "\n"
	config += `		route_map_name = "RM1"` + "\n"
	config += `	}]` + "\n"
	config += `	depends_on = [iosxe_restconf.PreReq0, iosxe_restconf.PreReq1, iosxe_restconf.PreReq2, iosxe_restconf.PreReq3, ]` + "\n"
	config += `}` + "\n"
	return config
}
