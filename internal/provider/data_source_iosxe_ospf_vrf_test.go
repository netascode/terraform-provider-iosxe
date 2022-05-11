// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceIosxeOSPFVRF(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxeOSPFVRFPrerequisitesConfig + testAccDataSourceIosxeOSPFVRFConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.iosxe_ospf_vrf.test", "bfd_all_interfaces", "true"),
					resource.TestCheckResourceAttr("data.iosxe_ospf_vrf.test", "default_information_originate", "true"),
					resource.TestCheckResourceAttr("data.iosxe_ospf_vrf.test", "default_information_originate_always", "true"),
					resource.TestCheckResourceAttr("data.iosxe_ospf_vrf.test", "default_metric", "21"),
					resource.TestCheckResourceAttr("data.iosxe_ospf_vrf.test", "distance", "120"),
					resource.TestCheckResourceAttr("data.iosxe_ospf_vrf.test", "domain_tag", "10"),
					resource.TestCheckResourceAttr("data.iosxe_ospf_vrf.test", "mpls_ldp_autoconfig", "true"),
					resource.TestCheckResourceAttr("data.iosxe_ospf_vrf.test", "mpls_ldp_sync", "true"),
					resource.TestCheckResourceAttr("data.iosxe_ospf_vrf.test", "neighbor.0.ip", "2.2.2.2"),
					resource.TestCheckResourceAttr("data.iosxe_ospf_vrf.test", "neighbor.0.priority", "10"),
					resource.TestCheckResourceAttr("data.iosxe_ospf_vrf.test", "neighbor.0.cost", "100"),
					resource.TestCheckResourceAttr("data.iosxe_ospf_vrf.test", "network.0.ip", "3.3.3.0"),
					resource.TestCheckResourceAttr("data.iosxe_ospf_vrf.test", "network.0.wildcard", "0.0.0.255"),
					resource.TestCheckResourceAttr("data.iosxe_ospf_vrf.test", "network.0.area", "0"),
					resource.TestCheckResourceAttr("data.iosxe_ospf_vrf.test", "priority", "100"),
					resource.TestCheckResourceAttr("data.iosxe_ospf_vrf.test", "router_id", "1.2.3.4"),
					resource.TestCheckResourceAttr("data.iosxe_ospf_vrf.test", "shutdown", "false"),
					resource.TestCheckResourceAttr("data.iosxe_ospf_vrf.test", "summary_address.0.ip", "3.3.3.0"),
					resource.TestCheckResourceAttr("data.iosxe_ospf_vrf.test", "summary_address.0.mask", "255.255.255.0"),
				),
			},
		},
	})
}

const testAccDataSourceIosxeOSPFVRFPrerequisitesConfig = `
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

`

const testAccDataSourceIosxeOSPFVRFConfig = `

resource "iosxe_ospf_vrf" "test" {
  process_id = 1
  vrf = "VRF1"
  bfd_all_interfaces = true
  default_information_originate = true
  default_information_originate_always = true
  default_metric = 21
  distance = 120
  domain_tag = 10
  mpls_ldp_autoconfig = true
  mpls_ldp_sync = true
  neighbor = [{
    ip = "2.2.2.2"
    priority = 10
    cost = 100
  }]
  network = [{
    ip = "3.3.3.0"
    wildcard = "0.0.0.255"
    area = "0"
  }]
  priority = 100
  router_id = "1.2.3.4"
  shutdown = false
  summary_address = [{
    ip = "3.3.3.0"
    mask = "255.255.255.0"
  }]
  depends_on = [iosxe_restconf.PreReq0, iosxe_restconf.PreReq1, ]
}

data "iosxe_ospf_vrf" "test" {
  process_id = 1
  vrf = "VRF1"
  depends_on = [iosxe_ospf_vrf.test]
}
`
