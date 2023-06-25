// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceIosxeInterfaceEthernet(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxeInterfaceEthernetPrerequisitesConfig + testAccDataSourceIosxeInterfaceEthernetConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.iosxe_interface_ethernet.test", "description", "My Interface Description"),
					resource.TestCheckResourceAttr("data.iosxe_interface_ethernet.test", "shutdown", "false"),
					resource.TestCheckResourceAttr("data.iosxe_interface_ethernet.test", "ip_proxy_arp", "false"),
					resource.TestCheckResourceAttr("data.iosxe_interface_ethernet.test", "ip_redirects", "false"),
					resource.TestCheckResourceAttr("data.iosxe_interface_ethernet.test", "unreachables", "false"),
					resource.TestCheckResourceAttr("data.iosxe_interface_ethernet.test", "ipv4_address", "15.1.1.1"),
					resource.TestCheckResourceAttr("data.iosxe_interface_ethernet.test", "ipv4_address_mask", "255.255.255.252"),
					resource.TestCheckResourceAttr("data.iosxe_interface_ethernet.test", "ip_dhcp_relay_source_interface", "Loopback100"),
					resource.TestCheckResourceAttr("data.iosxe_interface_ethernet.test", "ip_access_group_in", "1"),
					resource.TestCheckResourceAttr("data.iosxe_interface_ethernet.test", "ip_access_group_in_enable", "true"),
					resource.TestCheckResourceAttr("data.iosxe_interface_ethernet.test", "ip_access_group_out", "1"),
					resource.TestCheckResourceAttr("data.iosxe_interface_ethernet.test", "ip_access_group_out_enable", "true"),
					resource.TestCheckResourceAttr("data.iosxe_interface_ethernet.test", "helper_addresses.0.address", "10.10.10.10"),
					resource.TestCheckResourceAttr("data.iosxe_interface_ethernet.test", "helper_addresses.0.global", "false"),
					resource.TestCheckResourceAttr("data.iosxe_interface_ethernet.test", "helper_addresses.0.vrf", "VRF1"),
					resource.TestCheckResourceAttr("data.iosxe_interface_ethernet.test", "source_template.0.template_name", "TEMP1"),
					resource.TestCheckResourceAttr("data.iosxe_interface_ethernet.test", "source_template.0.merge", "false"),
				),
			},
		},
	})
}

const testAccDataSourceIosxeInterfaceEthernetPrerequisitesConfig = `
resource "iosxe_restconf" "PreReq0" {
	path = "Cisco-IOS-XE-native:native/vrf/definition=VRF1"
	delete = false
	attributes = {
		"name" = "VRF1"
		"address-family/ipv4" = ""
	}
}

`

const testAccDataSourceIosxeInterfaceEthernetConfig = `

resource "iosxe_interface_ethernet" "test" {
	delete_mode = "attributes"
	type = "GigabitEthernet"
	name = "3"
	description = "My Interface Description"
	shutdown = false
	ip_proxy_arp = false
	ip_redirects = false
	unreachables = false
	ipv4_address = "15.1.1.1"
	ipv4_address_mask = "255.255.255.252"
	ip_dhcp_relay_source_interface = "Loopback100"
	ip_access_group_in = "1"
	ip_access_group_in_enable = true
	ip_access_group_out = "1"
	ip_access_group_out_enable = true
	helper_addresses = [{
		address = "10.10.10.10"
		global = false
		vrf = "VRF1"
	}]
	source_template = [{
		template_name = "TEMP1"
		merge = false
	}]
	depends_on = [iosxe_restconf.PreReq0, ]
}

data "iosxe_interface_ethernet" "test" {
	type = "GigabitEthernet"
	name = "3"
	depends_on = [iosxe_interface_ethernet.test]
}
`
