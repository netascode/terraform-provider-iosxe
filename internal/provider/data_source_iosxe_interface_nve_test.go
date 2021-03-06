//go:build testAll

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceIosxeInterfaceNVE(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxeInterfaceNVEConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.iosxe_interface_nve.test", "description", "My Interface Description"),
					resource.TestCheckResourceAttr("data.iosxe_interface_nve.test", "shutdown", "false"),
					resource.TestCheckResourceAttr("data.iosxe_interface_nve.test", "host_reachability_protocol_bgp", "true"),
					resource.TestCheckResourceAttr("data.iosxe_interface_nve.test", "source_interface_loopback", "100"),
					resource.TestCheckResourceAttr("data.iosxe_interface_nve.test", "vnis.0.vni_range", "10000"),
					resource.TestCheckResourceAttr("data.iosxe_interface_nve.test", "vnis.0.ipv4_multicast_group", "225.1.1.1"),
				),
			},
		},
	})
}

const testAccDataSourceIosxeInterfaceNVEConfig = `

resource "iosxe_interface_nve" "test" {
  name = 1
  description = "My Interface Description"
  shutdown = false
  host_reachability_protocol_bgp = true
  source_interface_loopback = 100
  vnis = [{
    vni_range = "10000"
    ipv4_multicast_group = "225.1.1.1"
  }]
}

data "iosxe_interface_nve" "test" {
  name = 1
  depends_on = [iosxe_interface_nve.test]
}
`
