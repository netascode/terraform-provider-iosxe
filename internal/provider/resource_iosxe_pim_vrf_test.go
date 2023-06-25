// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIosxePIMVRF(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxePIMVRFPrerequisitesConfig + testAccIosxePIMVRFConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("iosxe_pim_vrf.test", "vrf", "VRF1"),
					resource.TestCheckResourceAttr("iosxe_pim_vrf.test", "autorp", "false"),
					resource.TestCheckResourceAttr("iosxe_pim_vrf.test", "autorp_listener", "false"),
					resource.TestCheckResourceAttr("iosxe_pim_vrf.test", "bsr_candidate_loopback", "100"),
					resource.TestCheckResourceAttr("iosxe_pim_vrf.test", "bsr_candidate_mask", "30"),
					resource.TestCheckResourceAttr("iosxe_pim_vrf.test", "bsr_candidate_priority", "10"),
					resource.TestCheckResourceAttr("iosxe_pim_vrf.test", "ssm_range", "10"),
					resource.TestCheckResourceAttr("iosxe_pim_vrf.test", "ssm_default", "false"),
					resource.TestCheckResourceAttr("iosxe_pim_vrf.test", "rp_address", "19.19.19.19"),
					resource.TestCheckResourceAttr("iosxe_pim_vrf.test", "rp_address_override", "false"),
					resource.TestCheckResourceAttr("iosxe_pim_vrf.test", "rp_address_bidir", "false"),
					resource.TestCheckResourceAttr("iosxe_pim_vrf.test", "rp_addresses.0.access_list", "10"),
					resource.TestCheckResourceAttr("iosxe_pim_vrf.test", "rp_addresses.0.rp_address", "10.10.10.10"),
					resource.TestCheckResourceAttr("iosxe_pim_vrf.test", "rp_addresses.0.override", "false"),
					resource.TestCheckResourceAttr("iosxe_pim_vrf.test", "rp_addresses.0.bidir", "false"),
					resource.TestCheckResourceAttr("iosxe_pim_vrf.test", "rp_candidates.0.interface", "Loopback100"),
					resource.TestCheckResourceAttr("iosxe_pim_vrf.test", "rp_candidates.0.interval", "100"),
					resource.TestCheckResourceAttr("iosxe_pim_vrf.test", "rp_candidates.0.priority", "10"),
					resource.TestCheckResourceAttr("iosxe_pim_vrf.test", "rp_candidates.0.bidir", "false"),
				),
			},
			{
				ResourceName:  "iosxe_pim_vrf.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XE-native:native/ip/pim/Cisco-IOS-XE-multicast:vrf=VRF1",
			},
		},
	})
}

const testAccIosxePIMVRFPrerequisitesConfig = `
resource "iosxe_restconf" "PreReq0" {
	path = "Cisco-IOS-XE-native:native/vrf/definition=VRF1"
	delete = false
	attributes = {
		"name" = "VRF1"
		"address-family/ipv4" = ""
	}
}

resource "iosxe_restconf" "PreReq1" {
	path = "Cisco-IOS-XE-native:native/interface/Loopback=100"
	attributes = {
		"name" = "100"
		"vrf/forwarding" = "VRF1"
		"ip/address/primary/address" = "200.200.200.200"
		"ip/address/primary/mask" = "255.255.255.255"
	}
	depends_on = [iosxe_restconf.PreReq0, ]
}

`

func testAccIosxePIMVRFConfig_minimum() string {
	return `
	resource "iosxe_pim_vrf" "test" {
		vrf = "VRF1"
		depends_on = [iosxe_restconf.PreReq0, iosxe_restconf.PreReq1, ]
	}
	`
}

func testAccIosxePIMVRFConfig_all() string {
	return `
	resource "iosxe_pim_vrf" "test" {
		vrf = "VRF1"
		autorp = false
		autorp_listener = false
		bsr_candidate_loopback = 100
		bsr_candidate_mask = 30
		bsr_candidate_priority = 10
		ssm_range = "10"
		ssm_default = false
		rp_address = "19.19.19.19"
		rp_address_override = false
		rp_address_bidir = false
		rp_addresses = [{
			access_list = "10"
			rp_address = "10.10.10.10"
			override = false
			bidir = false
		}]
		rp_candidates = [{
			interface = "Loopback100"
			interval = 100
			priority = 10
			bidir = false
		}]
		depends_on = [iosxe_restconf.PreReq0, iosxe_restconf.PreReq1, ]
	}
	`
}
