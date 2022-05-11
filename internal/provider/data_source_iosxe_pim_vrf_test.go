// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceIosxePIMVRF(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxePIMVRFPrerequisitesConfig + testAccDataSourceIosxePIMVRFConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.iosxe_pim_vrf.test", "autorp", "false"),
					resource.TestCheckResourceAttr("data.iosxe_pim_vrf.test", "autorp_listener", "false"),
					resource.TestCheckResourceAttr("data.iosxe_pim_vrf.test", "bsr_candidate_loopback", "100"),
					resource.TestCheckResourceAttr("data.iosxe_pim_vrf.test", "bsr_candidate_mask", "30"),
					resource.TestCheckResourceAttr("data.iosxe_pim_vrf.test", "bsr_candidate_priority", "10"),
					resource.TestCheckResourceAttr("data.iosxe_pim_vrf.test", "bsr_candidate_accept_rp_candidate", "10"),
					resource.TestCheckResourceAttr("data.iosxe_pim_vrf.test", "ssm_range", "10"),
					resource.TestCheckResourceAttr("data.iosxe_pim_vrf.test", "ssm_default", "true"),
					resource.TestCheckResourceAttr("data.iosxe_pim_vrf.test", "rp_addresses.0.access_list", "10"),
					resource.TestCheckResourceAttr("data.iosxe_pim_vrf.test", "rp_addresses.0.rp_address", "10.10.10.10"),
					resource.TestCheckResourceAttr("data.iosxe_pim_vrf.test", "rp_addresses.0.override", "false"),
					resource.TestCheckResourceAttr("data.iosxe_pim_vrf.test", "rp_addresses.0.bidir", "false"),
					resource.TestCheckResourceAttr("data.iosxe_pim_vrf.test", "rp_candidates.0.interface", "Loopback100"),
					resource.TestCheckResourceAttr("data.iosxe_pim_vrf.test", "rp_candidates.0.group_list", "10"),
					resource.TestCheckResourceAttr("data.iosxe_pim_vrf.test", "rp_candidates.0.interval", "100"),
					resource.TestCheckResourceAttr("data.iosxe_pim_vrf.test", "rp_candidates.0.priority", "10"),
					resource.TestCheckResourceAttr("data.iosxe_pim_vrf.test", "rp_candidates.0.bidir", "false"),
				),
			},
		},
	})
}

const testAccDataSourceIosxePIMVRFPrerequisitesConfig = `
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
  depends_on = [iosxe_restconf.PreReq2, ]
}

resource "iosxe_restconf" "PreReq4" {
  path = "Cisco-IOS-XE-native:native/interface/Loopback=100/ip/address/primary"
  attributes = {
      address = "200.200.200.200"
      mask = "255.255.255.255"
  }
  depends_on = [iosxe_restconf.PreReq3, ]
}

`

const testAccDataSourceIosxePIMVRFConfig = `

resource "iosxe_pim_vrf" "test" {
  vrf = "VRF1"
  autorp = false
  autorp_listener = false
  bsr_candidate_loopback = 100
  bsr_candidate_mask = 30
  bsr_candidate_priority = 10
  bsr_candidate_accept_rp_candidate = "10"
  ssm_range = "10"
  ssm_default = true
  rp_addresses = [{
    access_list = "10"
    rp_address = "10.10.10.10"
    override = false
    bidir = false
  }]
  rp_candidates = [{
    interface = "Loopback100"
    group_list = "10"
    interval = 100
    priority = 10
    bidir = false
  }]
  depends_on = [iosxe_restconf.PreReq0, iosxe_restconf.PreReq1, iosxe_restconf.PreReq2, iosxe_restconf.PreReq3, iosxe_restconf.PreReq4, ]
}

data "iosxe_pim_vrf" "test" {
  vrf = "VRF1"
  depends_on = [iosxe_pim_vrf.test]
}
`
