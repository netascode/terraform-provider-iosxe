// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIosxeLogging(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxeLoggingPrerequisitesConfig + testAccIosxeLoggingConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("iosxe_logging.test", "monitor_severity", "informational"),
					resource.TestCheckResourceAttr("iosxe_logging.test", "buffered_size", "16000"),
					resource.TestCheckResourceAttr("iosxe_logging.test", "buffered_severity", "informational"),
					resource.TestCheckResourceAttr("iosxe_logging.test", "console_severity", "informational"),
					resource.TestCheckResourceAttr("iosxe_logging.test", "facility", "local0"),
					resource.TestCheckResourceAttr("iosxe_logging.test", "history_size", "100"),
					resource.TestCheckResourceAttr("iosxe_logging.test", "history_severity", "informational"),
					resource.TestCheckResourceAttr("iosxe_logging.test", "trap", "true"),
					resource.TestCheckResourceAttr("iosxe_logging.test", "trap_severity", "informational"),
					resource.TestCheckResourceAttr("iosxe_logging.test", "origin_id_type", "hostname"),
					resource.TestCheckResourceAttr("iosxe_logging.test", "source_interface", "Loopback0"),
					resource.TestCheckResourceAttr("iosxe_logging.test", "source_interfaces_vrf.0.vrf", "VRF1"),
					resource.TestCheckResourceAttr("iosxe_logging.test", "source_interfaces_vrf.0.interface_name", "Loopback100"),
					resource.TestCheckResourceAttr("iosxe_logging.test", "ipv4_hosts.0.ipv4_host", "1.1.1.1"),
					resource.TestCheckResourceAttr("iosxe_logging.test", "ipv4_vrf_hosts.0.ipv4_host", "1.1.1.1"),
					resource.TestCheckResourceAttr("iosxe_logging.test", "ipv4_vrf_hosts.0.vrf", "VRF1"),
					resource.TestCheckResourceAttr("iosxe_logging.test", "ipv6_hosts.0.ipv6_host", "2001::1"),
					resource.TestCheckResourceAttr("iosxe_logging.test", "ipv6_vrf_hosts.0.ipv6_host", "2001::1"),
					resource.TestCheckResourceAttr("iosxe_logging.test", "ipv6_vrf_hosts.0.vrf", "VRF1"),
				),
			},
			{
				ResourceName:  "iosxe_logging.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XE-native:native/logging",
			},
		},
	})
}

const testAccIosxeLoggingPrerequisitesConfig = `
resource "iosxe_restconf" "PreReq0" {
  path = "Cisco-IOS-XE-native:native/vrf/definition=VRF1"
  delete = false
  attributes = {
      "name" = "VRF1"
      "address-family/ipv4" = ""
      "address-family/ipv6" = ""
  }
}

resource "iosxe_restconf" "PreReq1" {
  path = "Cisco-IOS-XE-native:native/interface/Loopback=100"
  attributes = {
      "name" = "100"
      "vrf/forwarding" = "VRF1"
  }
  depends_on = [iosxe_restconf.PreReq0, ]
}

`

func testAccIosxeLoggingConfig_minimum() string {
	return `
	resource "iosxe_logging" "test" {
  		depends_on = [iosxe_restconf.PreReq0, iosxe_restconf.PreReq1, ]
	}
	`
}

func testAccIosxeLoggingConfig_all() string {
	return `
	resource "iosxe_logging" "test" {
		monitor_severity = "informational"
		buffered_size = 16000
		buffered_severity = "informational"
		console_severity = "informational"
		facility = "local0"
		history_size = 100
		history_severity = "informational"
		trap = true
		trap_severity = "informational"
		origin_id_type = "hostname"
		source_interface = "Loopback0"
		source_interfaces_vrf = [{
			vrf = "VRF1"
			interface_name = "Loopback100"
		}]
		ipv4_hosts = [{
			ipv4_host = "1.1.1.1"
		}]
		ipv4_vrf_hosts = [{
			ipv4_host = "1.1.1.1"
			vrf = "VRF1"
		}]
		ipv6_hosts = [{
			ipv6_host = "2001::1"
		}]
		ipv6_vrf_hosts = [{
			ipv6_host = "2001::1"
			vrf = "VRF1"
		}]
  		depends_on = [iosxe_restconf.PreReq0, iosxe_restconf.PreReq1, ]
	}
	`
}
