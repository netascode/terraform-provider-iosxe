// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIosxeRouteMap(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxeRouteMapConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("iosxe_route_map.test", "name", "RM1"),
					resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.seq", "10"),
					resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.operation", "permit"),
					resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.description", "Entry 10"),
					resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.continue", "false"),
					resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.match_interfaces.0", "GigabitEthernet1"),
					resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.match_ip_address_access_lists.0", "ACL1"),
					resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.match_ip_next_hop_access_lists.0", "ACL1"),
					resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.match_ipv6_address_access_lists", "ACL1"),
					resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.match_ipv6_next_hop_access_lists", "ACL1"),
					resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.match_route_type_external", "true"),
					resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.match_route_type_external_type_1", "true"),
					resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.match_route_type_external_type_2", "true"),
					resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.match_route_type_internal", "true"),
					resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.match_route_type_level_1", "true"),
					resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.match_route_type_level_2", "true"),
					resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.match_route_type_local", "true"),
					resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.match_source_protocol_bgp.0", "65000"),
					resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.match_source_protocol_connected", "true"),
					resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.match_source_protocol_eigrp.0", "10"),
					resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.match_source_protocol_isis", "true"),
					resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.match_source_protocol_lisp", "true"),
					resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.match_source_protocol_ospf.0", "10"),
					resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.match_source_protocol_ospfv3.0", "10"),
					resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.match_source_protocol_rip", "true"),
					resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.match_source_protocol_static", "true"),
					resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.match_tags.0", "100"),
					resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.match_track", "1"),
					resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.match_as_paths.0", "10"),
					resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.match_community_lists.0", "COMM1"),
					resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.match_community_list_exact_match", "true"),
					resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.match_extcommunity_lists.0", "EXTCOMM1"),
					resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.match_local_preferences.0", "100"),
					resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.set_default_interfaces.0", "GigabitEthernet1"),
					resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.set_global", "false"),
					resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.set_interfaces.0", "GigabitEthernet1"),
					resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.set_ip_address", "PFL1"),
					resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.set_ip_default_global_next_hop_address.0", "1.2.3.4"),
					resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.set_ip_default_next_hop_address.0", "1.2.3.4"),
					resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.set_ip_global_next_hop_address.0", "1.2.3.4"),
					resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.set_ip_next_hop_address.0", "1.2.3.4"),
					resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.set_ip_qos_group", "1"),
					resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.set_ipv6_address", "PFL2"),
					resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.set_ipv6_default_global_next_hop", "2001::1"),
					resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.set_ipv6_default_next_hop.0", "2001::1"),
					resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.set_ipv6_next_hop.0", "2001::1"),
					resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.set_level_1", "true"),
					resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.set_metric_value", "110"),
					resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.set_metric_delay", "10"),
					resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.set_metric_reliability", "90"),
					resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.set_metric_loading", "10"),
					resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.set_metric_mtu", "1500"),
					resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.set_metric_type", "external"),
					resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.set_tag", "100"),
					resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.set_as_path_prepend_as", "65001 65001"),
					resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.set_as_path_prepend_last_as", "5"),
					resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.set_as_path_tag", "true"),
					resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.set_communities.0", "1:2"),
					resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.set_communities_additive", "true"),
					resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.set_community_list_delete", "true"),
					resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.set_community_list_name", "COMML1"),
					resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.set_extcomunity_rt.0", "10:10"),
					resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.set_extcomunity_soo", "10:10"),
					resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.set_extcomunity_vpn_distinguisher", "10:10"),
					resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.set_extcomunity_vpn_distinguisher_additive", "true"),
					resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.set_local_preference", "110"),
					resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.set_weight", "10000"),
				),
			},
			{
				ResourceName:  "iosxe_route_map.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XE-native:native/route-map=RM1",
			},
		},
	})
}

func testAccIosxeRouteMapConfig_minimum() string {
	return `
	resource "iosxe_route_map" "test" {
		name = "RM1"
	}
	`
}

func testAccIosxeRouteMapConfig_all() string {
	return `
	resource "iosxe_route_map" "test" {
		name = "RM1"
		entries = [{
		seq = 10
		operation = "permit"
		description = "Entry 10"
		continue = false
		match_interfaces = ["GigabitEthernet1"]
		match_ip_address_access_lists = ["ACL1"]
		match_ip_next_hop_access_lists = ["ACL1"]
		match_ipv6_address_access_lists = "ACL1"
		match_ipv6_next_hop_access_lists = "ACL1"
		match_route_type_external = true
		match_route_type_external_type_1 = true
		match_route_type_external_type_2 = true
		match_route_type_internal = true
		match_route_type_level_1 = true
		match_route_type_level_2 = true
		match_route_type_local = true
		match_source_protocol_bgp = ["65000"]
		match_source_protocol_connected = true
		match_source_protocol_eigrp = ["10"]
		match_source_protocol_isis = true
		match_source_protocol_lisp = true
		match_source_protocol_ospf = ["10"]
		match_source_protocol_ospfv3 = ["10"]
		match_source_protocol_rip = true
		match_source_protocol_static = true
		match_tags = [100]
		match_track = 1
		match_as_paths = [10]
		match_community_lists = ["COMM1"]
		match_community_list_exact_match = true
		match_extcommunity_lists = ["EXTCOMM1"]
		match_local_preferences = [100]
		set_default_interfaces = ["GigabitEthernet1"]
		set_global = false
		set_interfaces = ["GigabitEthernet1"]
		set_ip_address = "PFL1"
		set_ip_default_global_next_hop_address = ["1.2.3.4"]
		set_ip_default_next_hop_address = ["1.2.3.4"]
		set_ip_global_next_hop_address = ["1.2.3.4"]
		set_ip_next_hop_address = ["1.2.3.4"]
		set_ip_qos_group = 1
		set_ipv6_address = "PFL2"
		set_ipv6_default_global_next_hop = "2001::1"
		set_ipv6_default_next_hop = ["2001::1"]
		set_ipv6_next_hop = ["2001::1"]
		set_level_1 = true
		set_metric_value = 110
		set_metric_delay = "10"
		set_metric_reliability = 90
		set_metric_loading = 10
		set_metric_mtu = 1500
		set_metric_type = "external"
		set_tag = 100
		set_as_path_prepend_as = "65001 65001"
		set_as_path_prepend_last_as = 5
		set_as_path_tag = true
		set_communities = ["1:2"]
		set_communities_additive = true
		set_community_list_delete = true
		set_community_list_name = "COMML1"
		set_extcomunity_rt = ["10:10"]
		set_extcomunity_soo = "10:10"
		set_extcomunity_vpn_distinguisher = "10:10"
		set_extcomunity_vpn_distinguisher_additive = true
		set_local_preference = 110
		set_weight = 10000
		}]
	}
	`
}
