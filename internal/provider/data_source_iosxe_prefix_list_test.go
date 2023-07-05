// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceIosxePrefixList(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_prefix_list.test", "prefixes.0.name", "PREFIX_LIST_1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_prefix_list.test", "prefixes.0.seq", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_prefix_list.test", "prefixes.0.action", "permit"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_prefix_list.test", "prefixes.0.ip", "10.0.0.0/8"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_prefix_list.test", "prefixes.0.ge", "24"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_prefix_list.test", "prefixes.0.le", "32"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxePrefixListConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

func testAccDataSourceIosxePrefixListConfig() string {
	config := `resource "iosxe_prefix_list" "test" {` + "\n"
	config += `	prefixes = [{` + "\n"
	config += `		name = "PREFIX_LIST_1"` + "\n"
	config += `		seq = 10` + "\n"
	config += `		action = "permit"` + "\n"
	config += `		ip = "10.0.0.0/8"` + "\n"
	config += `		ge = 24` + "\n"
	config += `		le = 32` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "iosxe_prefix_list" "test" {
			depends_on = [iosxe_prefix_list.test]
		}
	`
	return config
}
