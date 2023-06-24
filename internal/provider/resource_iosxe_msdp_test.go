// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIosxeMSDP(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxeMSDPPrerequisitesConfig + testAccIosxeMSDPConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("iosxe_msdp.test", "originator_id", "Loopback100"),
					resource.TestCheckResourceAttr("iosxe_msdp.test", "peers.0.addr", "10.1.1.1"),
					resource.TestCheckResourceAttr("iosxe_msdp.test", "peers.0.remote_as", "65000"),
					resource.TestCheckResourceAttr("iosxe_msdp.test", "peers.0.connect_source_loopback", "100"),
					resource.TestCheckResourceAttr("iosxe_msdp.test", "passwords.0.addr", "10.1.1.1"),
					resource.TestCheckResourceAttr("iosxe_msdp.test", "passwords.0.encryption", "0"),
					resource.TestCheckResourceAttr("iosxe_msdp.test", "passwords.0.password", "Cisco123"),
				),
			},
			{
				ResourceName:  "iosxe_msdp.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XE-native:native/ip/Cisco-IOS-XE-multicast:msdp",
			},
		},
	})
}

const testAccIosxeMSDPPrerequisitesConfig = `
resource "iosxe_restconf" "PreReq0" {
  path = "Cisco-IOS-XE-native:native/interface/Loopback=100"
  attributes = {
      "name" = "100"
  }
}

`

func testAccIosxeMSDPConfig_minimum() string {
	return `
	resource "iosxe_msdp" "test" {
  		depends_on = [iosxe_restconf.PreReq0, ]
	}
	`
}

func testAccIosxeMSDPConfig_all() string {
	return `
	resource "iosxe_msdp" "test" {
		originator_id = "Loopback100"
		peers = [{
			addr = "10.1.1.1"
			remote_as = 65000
			connect_source_loopback = 100
		}]
		passwords = [{
			addr = "10.1.1.1"
			encryption = 0
			password = "Cisco123"
		}]
  		depends_on = [iosxe_restconf.PreReq0, ]
	}
	`
}
