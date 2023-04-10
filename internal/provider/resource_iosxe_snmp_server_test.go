// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIosxeSNMPServer(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxeSNMPServerConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("iosxe_snmp_server.test", "chassis_id", "R1"),
					resource.TestCheckResourceAttr("iosxe_snmp_server.test", "contact", "Contact1"),
					resource.TestCheckResourceAttr("iosxe_snmp_server.test", "ifindex_persist", "true"),
					resource.TestCheckResourceAttr("iosxe_snmp_server.test", "location", "Location1"),
					resource.TestCheckResourceAttr("iosxe_snmp_server.test", "packetsize", "2000"),
					resource.TestCheckResourceAttr("iosxe_snmp_server.test", "queue_length", "100"),
					resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_logging_getop", "true"),
					resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_logging_setop", "true"),
					resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps", "true"),
					resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_snmp_authentication", "true"),
					resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_snmp_coldstart", "true"),
					resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_snmp_linkdown", "true"),
					resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_snmp_linkup", "true"),
					resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_snmp_warmstart", "true"),
					resource.TestCheckResourceAttr("iosxe_snmp_server.test", "source_interface_informs_gigabit_ethernet", "1"),
					resource.TestCheckResourceAttr("iosxe_snmp_server.test", "source_interface_traps_gigabit_ethernet", "1"),
					resource.TestCheckResourceAttr("iosxe_snmp_server.test", "trap_source_gigabit_ethernet", "1"),
					resource.TestCheckResourceAttr("iosxe_snmp_server.test", "snmp_communities.0.name", "COM1"),
					resource.TestCheckResourceAttr("iosxe_snmp_server.test", "snmp_communities.0.view", "VIEW1"),
					resource.TestCheckResourceAttr("iosxe_snmp_server.test", "snmp_communities.0.permission", "ro"),
					resource.TestCheckResourceAttr("iosxe_snmp_server.test", "snmp_communities.0.ipv6", "ACL1"),
					resource.TestCheckResourceAttr("iosxe_snmp_server.test", "snmp_communities.0.access_list_name", "1"),
					resource.TestCheckResourceAttr("iosxe_snmp_server.test", "contexts.0.name", "CON1"),
					resource.TestCheckResourceAttr("iosxe_snmp_server.test", "views.0.name", "VIEW1"),
					resource.TestCheckResourceAttr("iosxe_snmp_server.test", "views.0.mib", "interfaces"),
					resource.TestCheckResourceAttr("iosxe_snmp_server.test", "views.0.inc_exl", "included"),
				),
			},
			{
				ResourceName:  "iosxe_snmp_server.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XE-native:native/snmp-server",
			},
		},
	})
}

func testAccIosxeSNMPServerConfig_minimum() string {
	return `
	resource "iosxe_snmp_server" "test" {
	}
	`
}

func testAccIosxeSNMPServerConfig_all() string {
	return `
	resource "iosxe_snmp_server" "test" {
		chassis_id = "R1"
		contact = "Contact1"
		ifindex_persist = true
		location = "Location1"
		packetsize = 2000
		queue_length = 100
		enable_logging_getop = true
		enable_logging_setop = true
		enable_traps = true
		enable_traps_snmp_authentication = true
		enable_traps_snmp_coldstart = true
		enable_traps_snmp_linkdown = true
		enable_traps_snmp_linkup = true
		enable_traps_snmp_warmstart = true
		source_interface_informs_gigabit_ethernet = "1"
		source_interface_traps_gigabit_ethernet = "1"
		trap_source_gigabit_ethernet = "1"
		snmp_communities = [{
			name = "COM1"
			view = "VIEW1"
			permission = "ro"
			ipv6 = "ACL1"
			access_list_name = "1"
		}]
		contexts = [{
			name = "CON1"
		}]
		views = [{
			name = "VIEW1"
			mib = "interfaces"
			inc_exl = "included"
		}]
	}
	`
}
