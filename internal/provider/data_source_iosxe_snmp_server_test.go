// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceIosxeSNMPServer(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_snmp_server.test", "chassis_id", "R1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_snmp_server.test", "contact", "Contact1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_snmp_server.test", "ifindex_persist", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_snmp_server.test", "location", "Location1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_snmp_server.test", "packetsize", "2000"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_snmp_server.test", "queue_length", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_snmp_server.test", "enable_logging_getop", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_snmp_server.test", "enable_logging_setop", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_snmp_server.test", "enable_traps", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_snmp_server.test", "enable_traps_snmp_authentication", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_snmp_server.test", "enable_traps_snmp_coldstart", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_snmp_server.test", "enable_traps_snmp_linkdown", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_snmp_server.test", "enable_traps_snmp_linkup", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_snmp_server.test", "enable_traps_snmp_warmstart", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_snmp_server.test", "source_interface_informs_loopback", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_snmp_server.test", "source_interface_traps_loopback", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_snmp_server.test", "trap_source_loopback", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_snmp_server.test", "snmp_communities.0.name", "COM1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_snmp_server.test", "snmp_communities.0.view", "VIEW1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_snmp_server.test", "snmp_communities.0.permission", "ro"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_snmp_server.test", "snmp_communities.0.ipv6", "ACL1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_snmp_server.test", "snmp_communities.0.access_list_name", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_snmp_server.test", "contexts.0.name", "CON1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_snmp_server.test", "views.0.name", "VIEW1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_snmp_server.test", "views.0.mib", "interfaces"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_snmp_server.test", "views.0.inc_exl", "included"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxeSNMPServerPrerequisitesConfig + testAccDataSourceIosxeSNMPServerConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

const testAccDataSourceIosxeSNMPServerPrerequisitesConfig = `
resource "iosxe_restconf" "PreReq0" {
	path = "Cisco-IOS-XE-native:native/interface/Loopback=1"
	attributes = {
		"name" = "1"
	}
}

`

func testAccDataSourceIosxeSNMPServerConfig() string {
	config := `resource "iosxe_snmp_server" "test" {` + "\n"
	config += `	delete_mode = "attributes"` + "\n"
	config += `	chassis_id = "R1"` + "\n"
	config += `	contact = "Contact1"` + "\n"
	config += `	ifindex_persist = true` + "\n"
	config += `	location = "Location1"` + "\n"
	config += `	packetsize = 2000` + "\n"
	config += `	queue_length = 100` + "\n"
	config += `	enable_logging_getop = true` + "\n"
	config += `	enable_logging_setop = true` + "\n"
	config += `	enable_traps = true` + "\n"
	config += `	enable_traps_snmp_authentication = true` + "\n"
	config += `	enable_traps_snmp_coldstart = true` + "\n"
	config += `	enable_traps_snmp_linkdown = true` + "\n"
	config += `	enable_traps_snmp_linkup = true` + "\n"
	config += `	enable_traps_snmp_warmstart = true` + "\n"
	config += `	source_interface_informs_loopback = 1` + "\n"
	config += `	source_interface_traps_loopback = 1` + "\n"
	config += `	trap_source_loopback = 1` + "\n"
	config += `	snmp_communities = [{` + "\n"
	config += `		name = "COM1"` + "\n"
	config += `		view = "VIEW1"` + "\n"
	config += `		permission = "ro"` + "\n"
	config += `		ipv6 = "ACL1"` + "\n"
	config += `		access_list_name = "1"` + "\n"
	config += `	}]` + "\n"
	config += `	contexts = [{` + "\n"
	config += `		name = "CON1"` + "\n"
	config += `	}]` + "\n"
	config += `	views = [{` + "\n"
	config += `		name = "VIEW1"` + "\n"
	config += `		mib = "interfaces"` + "\n"
	config += `		inc_exl = "included"` + "\n"
	config += `	}]` + "\n"
	config += `	depends_on = [iosxe_restconf.PreReq0, ]` + "\n"
	config += `}` + "\n"

	config += `
		data "iosxe_snmp_server" "test" {
			depends_on = [iosxe_snmp_server.test]
		}
	`
	return config
}
