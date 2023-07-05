// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIosxeDHCP(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_dhcp.test", "relay_information_trust_all", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_dhcp.test", "relay_information_option_default", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_dhcp.test", "relay_information_option_vpn", "true"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxeDHCPConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				ResourceName:  "iosxe_dhcp.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XE-native:native/ip/dhcp",
			},
		},
	})
}

func testAccIosxeDHCPConfig_minimum() string {
	config := `resource "iosxe_dhcp" "test" {` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxeDHCPConfig_all() string {
	config := `resource "iosxe_dhcp" "test" {` + "\n"
	config += `	relay_information_trust_all = false` + "\n"
	config += `	relay_information_option_default = false` + "\n"
	config += `	relay_information_option_vpn = true` + "\n"
	config += `}` + "\n"
	return config
}
