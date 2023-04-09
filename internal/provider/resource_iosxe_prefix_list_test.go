// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIosxePrefixList(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxePrefixListConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("iosxe_prefix_list.test", "prefixes.0.name", "PREFIX_LIST_1"),
					resource.TestCheckResourceAttr("iosxe_prefix_list.test", "prefixes.0.seq", "10"),
					resource.TestCheckResourceAttr("iosxe_prefix_list.test", "prefixes.0.action", "permit"),
					resource.TestCheckResourceAttr("iosxe_prefix_list.test", "prefixes.0.ip", "10.0.0.0/8"),
					resource.TestCheckResourceAttr("iosxe_prefix_list.test", "prefixes.0.ge", "24"),
					resource.TestCheckResourceAttr("iosxe_prefix_list.test", "prefixes.0.le", "32"),
				),
			},
			{
				ResourceName:  "iosxe_prefix_list.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XE-native:native/ip/prefix-lists",
			},
		},
	})
}

func testAccIosxePrefixListConfig_minimum() string {
	return `
	resource "iosxe_prefix_list" "test" {
	}
	`
}

func testAccIosxePrefixListConfig_all() string {
	return `
	resource "iosxe_prefix_list" "test" {
		prefixes = [{
		name = "PREFIX_LIST_1"
		seq = 10
		action = "permit"
		ip = "10.0.0.0/8"
		ge = 24
		le = 32
		}]
	}
	`
}
