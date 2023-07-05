// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceIosxeNTP(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_ntp.test", "authenticate", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_ntp.test", "logging", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_ntp.test", "access_group_peer_acl", "SACL1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_ntp.test", "access_group_query_only_acl", "SACL1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_ntp.test", "access_group_serve_acl", "SACL1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_ntp.test", "access_group_serve_only_acl", "SACL1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_ntp.test", "authentication_keys.0.number", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_ntp.test", "authentication_keys.0.md5", "060506324F41584B564347"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_ntp.test", "authentication_keys.0.encryption_type", "7"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_ntp.test", "master", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_ntp.test", "master_stratum", "5"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_ntp.test", "passive", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_ntp.test", "update_calendar", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_ntp.test", "trap_source_gigabit_ethernet", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_ntp.test", "servers.0.ip_address", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_ntp.test", "servers.0.source", "GigabitEthernet3"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_ntp.test", "servers.0.key", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_ntp.test", "servers.0.prefer", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_ntp.test", "servers.0.version", "2"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_ntp.test", "server_vrfs.0.name", "VRF1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_ntp.test", "server_vrfs.0.servers.0.ip_address", "3.4.5.6"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_ntp.test", "server_vrfs.0.servers.0.key", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_ntp.test", "server_vrfs.0.servers.0.prefer", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_ntp.test", "server_vrfs.0.servers.0.version", "2"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_ntp.test", "peers.0.ip_address", "5.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_ntp.test", "peers.0.source", "GigabitEthernet3"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_ntp.test", "peers.0.key", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_ntp.test", "peers.0.prefer", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_ntp.test", "peers.0.version", "2"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_ntp.test", "peer_vrfs.0.name", "VRF1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_ntp.test", "peer_vrfs.0.peers.0.ip_address", "5.4.5.6"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_ntp.test", "peer_vrfs.0.peers.0.key", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_ntp.test", "peer_vrfs.0.peers.0.prefer", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_ntp.test", "peer_vrfs.0.peers.0.version", "2"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxeNTPPrerequisitesConfig + testAccDataSourceIosxeNTPConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

const testAccDataSourceIosxeNTPPrerequisitesConfig = `
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

func testAccDataSourceIosxeNTPConfig() string {
	config := `resource "iosxe_ntp" "test" {` + "\n"
	config += `	delete_mode = "attributes"` + "\n"
	config += `	authenticate = true` + "\n"
	config += `	logging = false` + "\n"
	config += `	access_group_peer_acl = "SACL1"` + "\n"
	config += `	access_group_query_only_acl = "SACL1"` + "\n"
	config += `	access_group_serve_acl = "SACL1"` + "\n"
	config += `	access_group_serve_only_acl = "SACL1"` + "\n"
	config += `	authentication_keys = [{` + "\n"
	config += `		number = 1` + "\n"
	config += `		md5 = "060506324F41584B564347"` + "\n"
	config += `		encryption_type = 7` + "\n"
	config += `	}]` + "\n"
	config += `	master = true` + "\n"
	config += `	master_stratum = 5` + "\n"
	config += `	passive = false` + "\n"
	config += `	update_calendar = false` + "\n"
	config += `	trap_source_gigabit_ethernet = "1"` + "\n"
	config += `	servers = [{` + "\n"
	config += `		ip_address = "1.2.3.4"` + "\n"
	config += `		source = "GigabitEthernet3"` + "\n"
	config += `		key = 1` + "\n"
	config += `		prefer = true` + "\n"
	config += `		version = 2` + "\n"
	config += `	}]` + "\n"
	config += `	server_vrfs = [{` + "\n"
	config += `		name = "VRF1"` + "\n"
	config += `		servers = [{` + "\n"
	config += `			ip_address = "3.4.5.6"` + "\n"
	config += `			key = 1` + "\n"
	config += `			prefer = true` + "\n"
	config += `			version = 2` + "\n"
	config += `		}]` + "\n"
	config += `	}]` + "\n"
	config += `	peers = [{` + "\n"
	config += `		ip_address = "5.2.3.4"` + "\n"
	config += `		source = "GigabitEthernet3"` + "\n"
	config += `		key = 1` + "\n"
	config += `		prefer = true` + "\n"
	config += `		version = 2` + "\n"
	config += `	}]` + "\n"
	config += `	peer_vrfs = [{` + "\n"
	config += `		name = "VRF1"` + "\n"
	config += `		peers = [{` + "\n"
	config += `			ip_address = "5.4.5.6"` + "\n"
	config += `			key = 1` + "\n"
	config += `			prefer = true` + "\n"
	config += `			version = 2` + "\n"
	config += `		}]` + "\n"
	config += `	}]` + "\n"
	config += `	depends_on = [iosxe_restconf.PreReq0, iosxe_restconf.PreReq1, ]` + "\n"
	config += `}` + "\n"

	config += `
		data "iosxe_ntp" "test" {
			depends_on = [iosxe_ntp.test]
		}
	`
	return config
}
