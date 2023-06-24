// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceIosxeInterfaceLoopback(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxeInterfaceLoopbackPrerequisitesConfig + testAccDataSourceIosxeInterfaceLoopbackConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.iosxe_interface_loopback.test", "description", "My Interface Description"),
					resource.TestCheckResourceAttr("data.iosxe_interface_loopback.test", "shutdown", "false"),
					resource.TestCheckResourceAttr("data.iosxe_interface_loopback.test", "ip_proxy_arp", "false"),
					resource.TestCheckResourceAttr("data.iosxe_interface_loopback.test", "ip_redirects", "false"),
					resource.TestCheckResourceAttr("data.iosxe_interface_loopback.test", "unreachables", "false"),
					resource.TestCheckResourceAttr("data.iosxe_interface_loopback.test", "vrf_forwarding", "VRF1"),
					resource.TestCheckResourceAttr("data.iosxe_interface_loopback.test", "ipv4_address", "200.1.1.1"),
					resource.TestCheckResourceAttr("data.iosxe_interface_loopback.test", "ipv4_address_mask", "255.255.255.255"),
					resource.TestCheckResourceAttr("data.iosxe_interface_loopback.test", "ip_access_group_in", "1"),
					resource.TestCheckResourceAttr("data.iosxe_interface_loopback.test", "ip_access_group_in_enable", "true"),
					resource.TestCheckResourceAttr("data.iosxe_interface_loopback.test", "ip_access_group_out", "1"),
					resource.TestCheckResourceAttr("data.iosxe_interface_loopback.test", "ip_access_group_out_enable", "true"),
				),
			},
		},
	})
}

const testAccDataSourceIosxeInterfaceLoopbackPrerequisitesConfig = `
resource "iosxe_restconf" "PreReq0" {
  path = "Cisco-IOS-XE-native:native/vrf/definition=VRF1"
  delete = false
  attributes = {
      "name" = "VRF1"
      "address-family/ipv4" = ""
  }
}

`

const testAccDataSourceIosxeInterfaceLoopbackConfig = `

resource "iosxe_interface_loopback" "test" {
  name = 100
  description = "My Interface Description"
  shutdown = false
  ip_proxy_arp = false
  ip_redirects = false
  unreachables = false
  vrf_forwarding = "VRF1"
  ipv4_address = "200.1.1.1"
  ipv4_address_mask = "255.255.255.255"
  ip_access_group_in = "1"
  ip_access_group_in_enable = true
  ip_access_group_out = "1"
  ip_access_group_out_enable = true
  depends_on = [iosxe_restconf.PreReq0, ]
}

data "iosxe_interface_loopback" "test" {
  name = 100
  depends_on = [iosxe_interface_loopback.test]
}
`
