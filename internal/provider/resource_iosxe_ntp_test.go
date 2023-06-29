// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIosxeNTP(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxeNTPPrerequisitesConfig + testAccIosxeNTPConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("iosxe_ntp.test", "authenticate", "true"),
					resource.TestCheckResourceAttr("iosxe_ntp.test", "logging", "false"),
					resource.TestCheckResourceAttr("iosxe_ntp.test", "access_group_peer_acl", "SACL1"),
					resource.TestCheckResourceAttr("iosxe_ntp.test", "access_group_query_only_acl", "SACL1"),
					resource.TestCheckResourceAttr("iosxe_ntp.test", "access_group_serve_acl", "SACL1"),
					resource.TestCheckResourceAttr("iosxe_ntp.test", "access_group_serve_only_acl", "SACL1"),
					resource.TestCheckResourceAttr("iosxe_ntp.test", "authentication_keys.0.number", "1"),
					resource.TestCheckResourceAttr("iosxe_ntp.test", "authentication_keys.0.md5", "060506324F41584B564347"),
					resource.TestCheckResourceAttr("iosxe_ntp.test", "authentication_keys.0.encryption_type", "7"),
					resource.TestCheckResourceAttr("iosxe_ntp.test", "master", "true"),
					resource.TestCheckResourceAttr("iosxe_ntp.test", "master_stratum", "5"),
					resource.TestCheckResourceAttr("iosxe_ntp.test", "passive", "false"),
					resource.TestCheckResourceAttr("iosxe_ntp.test", "update_calendar", "false"),
					resource.TestCheckResourceAttr("iosxe_ntp.test", "trap_source_gigabit_ethernet", "1"),
					resource.TestCheckResourceAttr("iosxe_ntp.test", "servers.0.ip_address", "1.2.3.4"),
					resource.TestCheckResourceAttr("iosxe_ntp.test", "servers.0.source", "GigabitEthernet3"),
					resource.TestCheckResourceAttr("iosxe_ntp.test", "servers.0.key", "1"),
					resource.TestCheckResourceAttr("iosxe_ntp.test", "servers.0.prefer", "true"),
					resource.TestCheckResourceAttr("iosxe_ntp.test", "servers.0.version", "2"),
					resource.TestCheckResourceAttr("iosxe_ntp.test", "server_vrfs.0.name", "VRF1"),
					resource.TestCheckResourceAttr("iosxe_ntp.test", "server_vrfs.0.servers.0.ip_address", "3.4.5.6"),
					resource.TestCheckResourceAttr("iosxe_ntp.test", "server_vrfs.0.servers.0.key", "1"),
					resource.TestCheckResourceAttr("iosxe_ntp.test", "server_vrfs.0.servers.0.prefer", "true"),
					resource.TestCheckResourceAttr("iosxe_ntp.test", "server_vrfs.0.servers.0.version", "2"),
					resource.TestCheckResourceAttr("iosxe_ntp.test", "peers.0.ip_address", "5.2.3.4"),
					resource.TestCheckResourceAttr("iosxe_ntp.test", "peers.0.source", "GigabitEthernet3"),
					resource.TestCheckResourceAttr("iosxe_ntp.test", "peers.0.key", "1"),
					resource.TestCheckResourceAttr("iosxe_ntp.test", "peers.0.prefer", "true"),
					resource.TestCheckResourceAttr("iosxe_ntp.test", "peers.0.version", "2"),
					resource.TestCheckResourceAttr("iosxe_ntp.test", "peer_vrfs.0.name", "VRF1"),
					resource.TestCheckResourceAttr("iosxe_ntp.test", "peer_vrfs.0.peers.0.ip_address", "5.4.5.6"),
					resource.TestCheckResourceAttr("iosxe_ntp.test", "peer_vrfs.0.peers.0.key", "1"),
					resource.TestCheckResourceAttr("iosxe_ntp.test", "peer_vrfs.0.peers.0.prefer", "true"),
					resource.TestCheckResourceAttr("iosxe_ntp.test", "peer_vrfs.0.peers.0.version", "2"),
				),
			},
			{
				ResourceName:  "iosxe_ntp.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XE-native:native/ntp",
			},
		},
	})
}

const testAccIosxeNTPPrerequisitesConfig = `
resource "iosxe_restconf" "PreReq0" {
	path = "Cisco-IOS-XE-native:native/vrf/definition=VRF1"
	delete = false
	attributes = {
		"name" = "VRF1"
		"address-family/ipv4" = ""
	}
}

resource "iosxe_restconf" "PreReq1" {
	path = "Cisco-IOS-XE-native:native/ip/access-list/Cisco-IOS-XE-acl:standard=SACL1"
	attributes = {
		"name" = "SACL1"
	}
}

`

func testAccIosxeNTPConfig_minimum() string {
	return `
	resource "iosxe_ntp" "test" {
		depends_on = [iosxe_restconf.PreReq0, iosxe_restconf.PreReq1, ]
	}
	`
}

func testAccIosxeNTPConfig_all() string {
	return `
	resource "iosxe_ntp" "test" {
		authenticate = true
		logging = false
		access_group_peer_acl = "SACL1"
		access_group_query_only_acl = "SACL1"
		access_group_serve_acl = "SACL1"
		access_group_serve_only_acl = "SACL1"
		authentication_keys = [{
			number = 1
			md5 = "060506324F41584B564347"
			encryption_type = 7
		}]
		master = true
		master_stratum = 5
		passive = false
		update_calendar = false
		trap_source_gigabit_ethernet = "1"
		servers = [{
			ip_address = "1.2.3.4"
			source = "GigabitEthernet3"
			key = 1
			prefer = true
			version = 2
		}]
		server_vrfs = [{
			name = "VRF1"
			servers = [{
				ip_address = "3.4.5.6"
				key = 1
				prefer = true
				version = 2
			}]
		}]
		peers = [{
			ip_address = "5.2.3.4"
			source = "GigabitEthernet3"
			key = 1
			prefer = true
			version = 2
		}]
		peer_vrfs = [{
			name = "VRF1"
			peers = [{
				ip_address = "5.4.5.6"
				key = 1
				prefer = true
				version = 2
			}]
		}]
		depends_on = [iosxe_restconf.PreReq0, iosxe_restconf.PreReq1, ]
	}
	`
}
