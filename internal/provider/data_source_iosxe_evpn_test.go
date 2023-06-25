//go:build testAll

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceIosxeEVPN(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxeEVPNConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.iosxe_evpn.test", "replication_type_ingress", "false"),
					resource.TestCheckResourceAttr("data.iosxe_evpn.test", "replication_type_static", "true"),
					resource.TestCheckResourceAttr("data.iosxe_evpn.test", "replication_type_p2mp", "false"),
					resource.TestCheckResourceAttr("data.iosxe_evpn.test", "replication_type_mp2mp", "false"),
					resource.TestCheckResourceAttr("data.iosxe_evpn.test", "mac_duplication_limit", "10"),
					resource.TestCheckResourceAttr("data.iosxe_evpn.test", "mac_duplication_time", "100"),
					resource.TestCheckResourceAttr("data.iosxe_evpn.test", "ip_duplication_limit", "10"),
					resource.TestCheckResourceAttr("data.iosxe_evpn.test", "ip_duplication_time", "100"),
					resource.TestCheckResourceAttr("data.iosxe_evpn.test", "router_id_loopback", "100"),
					resource.TestCheckResourceAttr("data.iosxe_evpn.test", "default_gateway_advertise", "true"),
					resource.TestCheckResourceAttr("data.iosxe_evpn.test", "logging_peer_state", "true"),
					resource.TestCheckResourceAttr("data.iosxe_evpn.test", "route_target_auto_vni", "true"),
				),
			},
		},
	})
}

const testAccDataSourceIosxeEVPNConfig = `

resource "iosxe_evpn" "test" {
	delete_mode = "attributes"
	replication_type_ingress = false
	replication_type_static = true
	replication_type_p2mp = false
	replication_type_mp2mp = false
	mac_duplication_limit = 10
	mac_duplication_time = 100
	ip_duplication_limit = 10
	ip_duplication_time = 100
	router_id_loopback = 100
	default_gateway_advertise = true
	logging_peer_state = true
	route_target_auto_vni = true
}

data "iosxe_evpn" "test" {
	depends_on = [iosxe_evpn.test]
}
`
