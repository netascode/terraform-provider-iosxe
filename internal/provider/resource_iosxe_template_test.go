// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIosxeTemplate(t *testing.T) {
	if os.Getenv("C9000V") == "" {
		t.Skip("skipping test, set environment variable C9000V")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_template.test", "template_name", "TEMP1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_template.test", "dot1x_pae", "both"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_template.test", "dot1x_max_reauth_req", "3"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_template.test", "dot1x_max_req", "3"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_template.test", "service_policy_input", "SP1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_template.test", "service_policy_output", "SP2"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_template.test", "switchport_mode_trunk", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_template.test", "switchport_mode_access", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_template.test", "switchport_nonegotiate", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_template.test", "switchport_block_unicast", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_template.test", "switchport_port_security", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_template.test", "switchport_port_security_aging_static", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_template.test", "switchport_port_security_aging_time", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_template.test", "switchport_port_security_aging_type", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_template.test", "switchport_port_security_aging_type_inactivity", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_template.test", "switchport_port_security_maximum_range.0.range", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_template.test", "switchport_port_security_maximum_range.0.vlan", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_template.test", "switchport_port_security_maximum_range.0.vlan_access", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_template.test", "switchport_port_security_violation_protect", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_template.test", "switchport_port_security_violation_restrict", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_template.test", "switchport_port_security_violation_shutdown", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_template.test", "switchport_access_vlan", "200"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_template.test", "switchport_voice_vlan", "201"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_template.test", "switchport_private_vlan_host_association_primary_range", "301"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_template.test", "switchport_private_vlan_host_association_secondary_range", "302"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_template.test", "switchport_trunk_allowed_vlans", "500-599"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_template.test", "switchport_trunk_native_vlan_vlan_id", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_template.test", "mab", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_template.test", "mab_eap", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_template.test", "access_session_closed", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_template.test", "access_session_monitor", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_template.test", "access_session_port_control", "auto"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_template.test", "access_session_control_direction", "both"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_template.test", "access_session_host_mode", "single-host"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_template.test", "access_session_interface_template_sticky", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_template.test", "access_session_interface_template_sticky_timer", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_template.test", "authentication_periodic", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_template.test", "authentication_timer_reauthenticate_server", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_template.test", "spanning_tree_bpduguard_enable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_template.test", "spanning_tree_portfast", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_template.test", "spanning_tree_portfast_disable", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_template.test", "spanning_tree_portfast_edge", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_template.test", "spanning_tree_portfast_network", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_template.test", "storm_control_broadcast_level_pps_threshold", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_template.test", "storm_control_broadcast_level_bps_threshold", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_template.test", "storm_control_broadcast_level_threshold", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_template.test", "storm_control_multicast_level_pps_threshold", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_template.test", "storm_control_multicast_level_bps_threshold", "10000"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_template.test", "storm_control_multicast_level_threshold", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_template.test", "storm_control_action_shutdown", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_template.test", "storm_control_action_trap", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_template.test", "load_interval", "30"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_template.test", "ip_dhcp_snooping_limit_rate", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_template.test", "ip_dhcp_snooping_trust", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_template.test", "ip_access_group.0.direction", "in"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_template.test", "ip_access_group.0.access_list", "ACL1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_template.test", "subscriber_aging_probe", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_template.test", "device_tracking", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_template.test", "device_tracking_vlan_range", "100-199"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_template.test", "cts_manual", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_template.test", "cts_manual_policy_static_sgt", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_template.test", "cts_manual_policy_static_trusted", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_template.test", "cts_manual_propagate_sgt", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_template.test", "cts_role_based_enforcement", "false"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxeTemplateConfig_minimum(),
			},
			{
				Config: testAccIosxeTemplateConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
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
	config := `resource "iosxe_template" "test" {` + "\n"
	config += `	template_name = "TEMP1"` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxeTemplateConfig_all() string {
	config := `resource "iosxe_template" "test" {` + "\n"
	config += `	template_name = "TEMP1"` + "\n"
	config += `	dot1x_pae = "both"` + "\n"
	config += `	dot1x_max_reauth_req = 3` + "\n"
	config += `	dot1x_max_req = 3` + "\n"
	config += `	service_policy_input = "SP1"` + "\n"
	config += `	service_policy_output = "SP2"` + "\n"
	config += `	switchport_mode_trunk = true` + "\n"
	config += `	switchport_mode_access = false` + "\n"
	config += `	switchport_nonegotiate = false` + "\n"
	config += `	switchport_block_unicast = false` + "\n"
	config += `	switchport_port_security = true` + "\n"
	config += `	switchport_port_security_aging_static = false` + "\n"
	config += `	switchport_port_security_aging_time = 100` + "\n"
	config += `	switchport_port_security_aging_type = true` + "\n"
	config += `	switchport_port_security_aging_type_inactivity = true` + "\n"
	config += `	switchport_port_security_maximum_range = [{` + "\n"
	config += `		range = 100` + "\n"
	config += `		vlan = true` + "\n"
	config += `		vlan_access = true` + "\n"
	config += `	}]` + "\n"
	config += `	switchport_port_security_violation_protect = false` + "\n"
	config += `	switchport_port_security_violation_restrict = false` + "\n"
	config += `	switchport_port_security_violation_shutdown = false` + "\n"
	config += `	switchport_access_vlan = 200` + "\n"
	config += `	switchport_voice_vlan = 201` + "\n"
	config += `	switchport_private_vlan_host_association_primary_range = 301` + "\n"
	config += `	switchport_private_vlan_host_association_secondary_range = 302` + "\n"
	config += `	switchport_trunk_allowed_vlans = "500-599"` + "\n"
	config += `	switchport_trunk_native_vlan_vlan_id = 10` + "\n"
	config += `	mab = true` + "\n"
	config += `	mab_eap = true` + "\n"
	config += `	access_session_closed = true` + "\n"
	config += `	access_session_monitor = true` + "\n"
	config += `	access_session_port_control = "auto"` + "\n"
	config += `	access_session_control_direction = "both"` + "\n"
	config += `	access_session_host_mode = "single-host"` + "\n"
	config += `	access_session_interface_template_sticky = true` + "\n"
	config += `	access_session_interface_template_sticky_timer = 100` + "\n"
	config += `	authentication_periodic = true` + "\n"
	config += `	authentication_timer_reauthenticate_server = true` + "\n"
	config += `	spanning_tree_bpduguard_enable = true` + "\n"
	config += `	spanning_tree_portfast = true` + "\n"
	config += `	spanning_tree_portfast_disable = false` + "\n"
	config += `	spanning_tree_portfast_edge = false` + "\n"
	config += `	spanning_tree_portfast_network = false` + "\n"
	config += `	storm_control_broadcast_level_pps_threshold = "10"` + "\n"
	config += `	storm_control_broadcast_level_bps_threshold = 10` + "\n"
	config += `	storm_control_broadcast_level_threshold = 10` + "\n"
	config += `	storm_control_multicast_level_pps_threshold = "10"` + "\n"
	config += `	storm_control_multicast_level_bps_threshold = 10000` + "\n"
	config += `	storm_control_multicast_level_threshold = 10` + "\n"
	config += `	storm_control_action_shutdown = true` + "\n"
	config += `	storm_control_action_trap = true` + "\n"
	config += `	load_interval = 30` + "\n"
	config += `	ip_dhcp_snooping_limit_rate = 10` + "\n"
	config += `	ip_dhcp_snooping_trust = true` + "\n"
	config += `	ip_access_group = [{` + "\n"
	config += `		direction = "in"` + "\n"
	config += `		access_list = "ACL1"` + "\n"
	config += `	}]` + "\n"
	config += `	subscriber_aging_probe = true` + "\n"
	config += `	device_tracking = true` + "\n"
	config += `	device_tracking_vlan_range = "100-199"` + "\n"
	config += `	cts_manual = true` + "\n"
	config += `	cts_manual_policy_static_sgt = 100` + "\n"
	config += `	cts_manual_policy_static_trusted = false` + "\n"
	config += `	cts_manual_propagate_sgt = false` + "\n"
	config += `	cts_role_based_enforcement = false` + "\n"
	config += `}` + "\n"
	return config
}
