//go:build testAll

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceIosxeBGPAddressFamilyL2VPN(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxeBGPAddressFamilyL2VPNPrerequisitesConfig + testAccDataSourceIosxeBGPAddressFamilyL2VPNConfig,
				Check:  resource.ComposeTestCheckFunc(),
			},
		},
	})
}

const testAccDataSourceIosxeBGPAddressFamilyL2VPNPrerequisitesConfig = `
resource "iosxe_restconf" "PreReq0" {
  path = "Cisco-IOS-XE-native:native/router/Cisco-IOS-XE-bgp:bgp=65000"
  attributes = {
      "id" = "65000"
  }
}

`

const testAccDataSourceIosxeBGPAddressFamilyL2VPNConfig = `

resource "iosxe_bgp_address_family_l2vpn" "test" {
  asn = "65000"
  af_name = "evpn"
  depends_on = [iosxe_restconf.PreReq0, ]
}

data "iosxe_bgp_address_family_l2vpn" "test" {
  asn = "65000"
  af_name = "evpn"
  depends_on = [iosxe_bgp_address_family_l2vpn.test]
}
`
