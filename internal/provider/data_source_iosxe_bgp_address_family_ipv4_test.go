//go:build testAll

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceIosxeBGPAddressFamilyIPv4(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxeBGPAddressFamilyIPv4PrerequisitesConfig + testAccDataSourceIosxeBGPAddressFamilyIPv4Config,
				Check:  resource.ComposeTestCheckFunc(),
			},
		},
	})
}

const testAccDataSourceIosxeBGPAddressFamilyIPv4PrerequisitesConfig = `
resource "iosxe_restconf" "PreReq0" {
  path = "Cisco-IOS-XE-native:native/router/Cisco-IOS-XE-bgp:bgp=65000"
  attributes = {
      "id" = "65000"
  }
}

`

const testAccDataSourceIosxeBGPAddressFamilyIPv4Config = `

resource "iosxe_bgp_address_family_ipv4" "test" {
  asn = "65000"
  af_name = "unicast"
  depends_on = [iosxe_restconf.PreReq0, ]
}

data "iosxe_bgp_address_family_ipv4" "test" {
  asn = "65000"
  af_name = "unicast"
  depends_on = [iosxe_bgp_address_family_ipv4.test]
}
`