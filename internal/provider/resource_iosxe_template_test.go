//go:build testAll

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIosxeTemplate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxeTemplateConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("iosxe_template.test", "template_name", "TEMP1"),
					resource.TestCheckResourceAttr("iosxe_template.test", "dot1x_pae", "both"),
					resource.TestCheckResourceAttr("iosxe_template.test", "dot1x_max_reauth_req", "3"),
					resource.TestCheckResourceAttr("iosxe_template.test", "dot1x_max_req", "3"),
					resource.TestCheckResourceAttr("iosxe_template.test", "service_policy_input", "SP1"),
					resource.TestCheckResourceAttr("iosxe_template.test", "service_policy_output", "SP2"),
					resource.TestCheckResourceAttr("iosxe_template.test", "switchport_mode_trunk", "true"),
					resource.TestCheckResourceAttr("iosxe_template.test", "switchport_mode_access", "false"),
					resource.TestCheckResourceAttr("iosxe_template.test", "switchport_nonegotiate", "false"),
					resource.TestCheckResourceAttr("iosxe_template.test", "switchport_block_unicast", "false"),
					resource.TestCheckResourceAttr("iosxe_template.test", "switchport_port_security", "true"),
					resource.TestCheckResourceAttr("iosxe_template.test", "switchport_port_security_aging_static", "false"),
					resource.TestCheckResourceAttr("iosxe_template.test", "switchport_port_security_aging_time", "100"),
					resource.TestCheckResourceAttr("iosxe_template.test", "switchport_port_security_aging_type", "true"),
					resource.TestCheckResourceAttr("iosxe_template.test", "switchport_port_security_aging_type_inactivity", "true"),
					resource.TestCheckResourceAttr("iosxe_template.test", "switchport_port_security_maximum_range.0.range", "100"),
					resource.TestCheckResourceAttr("iosxe_template.test", "switchport_port_security_maximum_range.0.vlan", "true"),
					resource.TestCheckResourceAttr("iosxe_template.test", "switchport_port_security_maximum_range.0.vlan_access", "true"),
					resource.TestCheckResourceAttr("iosxe_template.test", "switchport_port_security_violation_protect", "false"),
					resource.TestCheckResourceAttr("iosxe_template.test", "switchport_port_security_violation_restrict", "false"),
					resource.TestCheckResourceAttr("iosxe_template.test", "switchport_port_security_violation_shutdown", "false"),
					resource.TestCheckResourceAttr("iosxe_template.test", "switchport_access_vlan", "200"),
					resource.TestCheckResourceAttr("iosxe_template.test", "switchport_voice_vlan", "201"),
					resource.TestCheckResourceAttr("iosxe_template.test", "switchport_private_vlan_host_association_primary_range", "301"),
					resource.TestCheckResourceAttr("iosxe_template.test", "switchport_private_vlan_host_association_secondary_range", "302"),
					resource.TestCheckResourceAttr("iosxe_template.test", "switchport_trunk_allowed_vlans", "500-599"),
					resource.TestCheckResourceAttr("iosxe_template.test", "switchport_trunk_native_vlan_vlan_id", "10"),
					resource.TestCheckResourceAttr("iosxe_template.test", "mab", "true"),
					resource.TestCheckResourceAttr("iosxe_template.test", "mab_eap", "true"),
					resource.TestCheckResourceAttr("iosxe_template.test", "access_session_closed", "true"),
					resource.TestCheckResourceAttr("iosxe_template.test", "access_session_monitor", "true"),
					resource.TestCheckResourceAttr("iosxe_template.test", "access_session_port_control", "auto"),
					resource.TestCheckResourceAttr("iosxe_template.test", "access_session_control_direction", "both"),
					resource.TestCheckResourceAttr("iosxe_template.test", "access_session_host_mode", "single-host"),
					resource.TestCheckResourceAttr("iosxe_template.test", "access_session_interface_template_sticky", "true"),
					resource.TestCheckResourceAttr("iosxe_template.test", "access_session_interface_template_sticky_timer", "100"),
					resource.TestCheckResourceAttr("iosxe_template.test", "authentication_periodic", "true"),
					resource.TestCheckResourceAttr("iosxe_template.test", "authentication_timer_reauthenticate_server", "true"),
					resource.TestCheckResourceAttr("iosxe_template.test", "spanning_tree_bpduguard_enable", "true"),
					resource.TestCheckResourceAttr("iosxe_template.test", "spanning_tree_portfast", "true"),
					resource.TestCheckResourceAttr("iosxe_template.test", "spanning_tree_portfast_disable", "false"),
					resource.TestCheckResourceAttr("iosxe_template.test", "spanning_tree_portfast_edge", "false"),
					resource.TestCheckResourceAttr("iosxe_template.test", "spanning_tree_portfast_network", "false"),
					resource.TestCheckResourceAttr("iosxe_template.test", "storm_control_broadcast_level_pps_threshold", "10"),
					resource.TestCheckResourceAttr("iosxe_template.test", "storm_control_broadcast_level_bps_threshold", "10"),
					resource.TestCheckResourceAttr("iosxe_template.test", "storm_control_broadcast_level_threshold", "10"),
					resource.TestCheckResourceAttr("iosxe_template.test", "storm_control_multicast_level_pps_threshold", "10"),
					resource.TestCheckResourceAttr("iosxe_template.test", "storm_control_multicast_level_bps_threshold", "10000"),
					resource.TestCheckResourceAttr("iosxe_template.test", "storm_control_multicast_level_threshold", "10"),
					resource.TestCheckResourceAttr("iosxe_template.test", "storm_control_action_shutdown", "true"),
					resource.TestCheckResourceAttr("iosxe_template.test", "storm_control_action_trap", "true"),
					resource.TestCheckResourceAttr("iosxe_template.test", "load_interval", "30"),
					resource.TestCheckResourceAttr("iosxe_template.test", "ip_dhcp_snooping_limit_rate", "10"),
					resource.TestCheckResourceAttr("iosxe_template.test", "ip_dhcp_snooping_trust", "true"),
					resource.TestCheckResourceAttr("iosxe_template.test", "ip_access_group.0.direction", "in"),
					resource.TestCheckResourceAttr("iosxe_template.test", "ip_access_group.0.access_list", "ACL1"),
					resource.TestCheckResourceAttr("iosxe_template.test", "subscriber_aging_probe", "true"),
					resource.TestCheckResourceAttr("iosxe_template.test", "device_tracking", "true"),
					resource.TestCheckResourceAttr("iosxe_template.test", "device_tracking_vlan_range", "100-199"),
					resource.TestCheckResourceAttr("iosxe_template.test", "cts_manual", "true"),
					resource.TestCheckResourceAttr("iosxe_template.test", "cts_manual_policy_static_sgt", "100"),
					resource.TestCheckResourceAttr("iosxe_template.test", "cts_manual_policy_static_trusted", "false"),
					resource.TestCheckResourceAttr("iosxe_template.test", "cts_manual_propagate_sgt", "false"),
					resource.TestCheckResourceAttr("iosxe_template.test", "cts_role_based_enforcement", "false"),
				),
			},
			{
				ResourceName:  "iosxe_template.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XE-native:native/template/Cisco-IOS-XE-template:template_details=TEMP1",
			},
		},
	})
}

func testAccIosxeTemplateConfig_minimum() string {
	return `
	resource "iosxe_template" "test" {
		template_name = "TEMP1"
	}
	`
}

func testAccIosxeTemplateConfig_all() string {
	return `
	resource "iosxe_template" "test" {
		template_name = "TEMP1"
		dot1x_pae = "both"
		dot1x_max_reauth_req = 3
		dot1x_max_req = 3
		service_policy_input = "SP1"
		service_policy_output = "SP2"
		switchport_mode_trunk = true
		switchport_mode_access = false
		switchport_nonegotiate = false
		switchport_block_unicast = false
		switchport_port_security = true
		switchport_port_security_aging_static = false
		switchport_port_security_aging_time = 100
		switchport_port_security_aging_type = true
		switchport_port_security_aging_type_inactivity = true
		switchport_port_security_maximum_range = [{
		range = 100
		vlan = true
		vlan_access = true
		}]
		switchport_port_security_violation_protect = false
		switchport_port_security_violation_restrict = false
		switchport_port_security_violation_shutdown = false
		switchport_access_vlan = 200
		switchport_voice_vlan = 201
		switchport_private_vlan_host_association_primary_range = 301
		switchport_private_vlan_host_association_secondary_range = 302
		switchport_trunk_allowed_vlans = "500-599"
		switchport_trunk_native_vlan_vlan_id = 10
		mab = true
		mab_eap = true
		access_session_closed = true
		access_session_monitor = true
		access_session_port_control = "auto"
		access_session_control_direction = "both"
		access_session_host_mode = "single-host"
		access_session_interface_template_sticky = true
		access_session_interface_template_sticky_timer = 100
		authentication_periodic = true
		authentication_timer_reauthenticate_server = true
		spanning_tree_bpduguard_enable = true
		spanning_tree_portfast = true
		spanning_tree_portfast_disable = false
		spanning_tree_portfast_edge = false
		spanning_tree_portfast_network = false
		storm_control_broadcast_level_pps_threshold = "10"
		storm_control_broadcast_level_bps_threshold = 10
		storm_control_broadcast_level_threshold = 10
		storm_control_multicast_level_pps_threshold = "10"
		storm_control_multicast_level_bps_threshold = 10000
		storm_control_multicast_level_threshold = 10
		storm_control_action_shutdown = true
		storm_control_action_trap = true
		load_interval = 30
		ip_dhcp_snooping_limit_rate = 10
		ip_dhcp_snooping_trust = true
		ip_access_group = [{
		direction = "in"
		access_list = "ACL1"
		}]
		subscriber_aging_probe = true
		device_tracking = true
		device_tracking_vlan_range = "100-199"
		cts_manual = true
		cts_manual_policy_static_sgt = 100
		cts_manual_policy_static_trusted = false
		cts_manual_propagate_sgt = false
		cts_role_based_enforcement = false
	}
	`
}