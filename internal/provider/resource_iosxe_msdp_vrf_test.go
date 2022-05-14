// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIosxeMSDPVRF(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxeMSDPVRFPrerequisitesConfig + testAccIosxeMSDPVRFConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("iosxe_msdp_vrf.test", "vrf", "VRF1"),
					resource.TestCheckResourceAttr("iosxe_msdp_vrf.test", "originator_id", "Loopback100"),
					resource.TestCheckResourceAttr("iosxe_msdp_vrf.test", "peers.0.addr", "10.1.1.1"),
					resource.TestCheckResourceAttr("iosxe_msdp_vrf.test", "peers.0.remote_as", "65000"),
					resource.TestCheckResourceAttr("iosxe_msdp_vrf.test", "peers.0.connect_source_loopback", "100"),
					resource.TestCheckResourceAttr("iosxe_msdp_vrf.test", "passwords.0.addr", "10.1.1.1"),
					resource.TestCheckResourceAttr("iosxe_msdp_vrf.test", "passwords.0.encryption", "0"),
					resource.TestCheckResourceAttr("iosxe_msdp_vrf.test", "passwords.0.password", "Cisco123"),
				),
			},
			{
				ResourceName:  "iosxe_msdp_vrf.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XE-native:native/ip/Cisco-IOS-XE-multicast:msdp/vrf=VRF1",
			},
		},
	})
}

const testAccIosxeMSDPVRFPrerequisitesConfig = `
resource "iosxe_restconf" "PreReq0" {
  path = "Cisco-IOS-XE-native:native/vrf/definition=VRF1"
  delete = false
  attributes = {
      name = "VRF1"
  }
}

resource "iosxe_restconf" "PreReq1" {
  path = "Cisco-IOS-XE-native:native/vrf/definition=VRF1/address-family"
  delete = false
  attributes = {
      ipv4 = ""
  }
  depends_on = [iosxe_restconf.PreReq0, ]
}

resource "iosxe_restconf" "PreReq2" {
  path = "Cisco-IOS-XE-native:native/interface/Loopback=100"
  attributes = {
      name = "100"
  }
}

resource "iosxe_restconf" "PreReq3" {
  path = "Cisco-IOS-XE-native:native/interface/Loopback=100/vrf"
  attributes = {
      forwarding = "VRF1"
  }
  depends_on = [iosxe_restconf.PreReq1, iosxe_restconf.PreReq2, ]
}

`

func testAccIosxeMSDPVRFConfig_minimum() string {
	return `
	resource "iosxe_msdp_vrf" "test" {
		vrf = "VRF1"
  		depends_on = [iosxe_restconf.PreReq0, iosxe_restconf.PreReq1, iosxe_restconf.PreReq2, iosxe_restconf.PreReq3, ]
	}
	`
}

func testAccIosxeMSDPVRFConfig_all() string {
	return `
	resource "iosxe_msdp_vrf" "test" {
		vrf = "VRF1"
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
  		depends_on = [iosxe_restconf.PreReq0, iosxe_restconf.PreReq1, iosxe_restconf.PreReq2, iosxe_restconf.PreReq3, ]
	}
	`
}
