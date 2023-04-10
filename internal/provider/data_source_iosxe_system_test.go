// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceIosxeSystem(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxeSystemPrerequisitesConfig + testAccDataSourceIosxeSystemConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.iosxe_system.test", "hostname", "ROUTER-1"),
					resource.TestCheckResourceAttr("data.iosxe_system.test", "ipv6_unicast_routing", "true"),
					resource.TestCheckResourceAttr("data.iosxe_system.test", "ip_source_route", "false"),
					resource.TestCheckResourceAttr("data.iosxe_system.test", "ip_domain_lookup", "false"),
					resource.TestCheckResourceAttr("data.iosxe_system.test", "ip_domain_name", "test.com"),
					resource.TestCheckResourceAttr("data.iosxe_system.test", "login_delay", "10"),
					resource.TestCheckResourceAttr("data.iosxe_system.test", "login_on_failure", "true"),
					resource.TestCheckResourceAttr("data.iosxe_system.test", "login_on_failure_log", "true"),
					resource.TestCheckResourceAttr("data.iosxe_system.test", "login_on_success", "true"),
					resource.TestCheckResourceAttr("data.iosxe_system.test", "login_on_success_log", "true"),
					resource.TestCheckResourceAttr("data.iosxe_system.test", "multicast_routing", "true"),
					resource.TestCheckResourceAttr("data.iosxe_system.test", "multicast_routing_distributed", "true"),
					resource.TestCheckResourceAttr("data.iosxe_system.test", "multicast_routing_vrfs.0.vrf", "VRF1"),
					resource.TestCheckResourceAttr("data.iosxe_system.test", "multicast_routing_vrfs.0.distributed", "true"),
				),
			},
		},
	})
}

const testAccDataSourceIosxeSystemPrerequisitesConfig = `
resource "iosxe_restconf" "PreReq0" {
  path = "Cisco-IOS-XE-native:native/vrf/definition=VRF1"
  delete = false
  attributes = {
      "name" = "VRF1"
      "address-family/ipv4" = ""
  }
}

`

const testAccDataSourceIosxeSystemConfig = `

resource "iosxe_system" "test" {
  hostname = "ROUTER-1"
  ipv6_unicast_routing = true
  ip_source_route = false
  ip_domain_lookup = false
  ip_domain_name = "test.com"
  login_delay = 10
  login_on_failure = true
  login_on_failure_log = true
  login_on_success = true
  login_on_success_log = true
  multicast_routing = true
  multicast_routing_distributed = true
  multicast_routing_vrfs = [{
    vrf = "VRF1"
    distributed = true
  }]
  depends_on = [iosxe_restconf.PreReq0, ]
}

data "iosxe_system" "test" {
  depends_on = [iosxe_system.test]
}
`
