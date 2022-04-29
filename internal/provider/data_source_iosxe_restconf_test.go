package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceIosxeRestconf(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxeRestconfConfigInterface,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.iosxe_restconf.test", "id", "Cisco-IOS-XE-native:native/banner/login"),
					resource.TestCheckResourceAttr("data.iosxe_restconf.test", "attributes.banner", "My Banner"),
				),
			},
		},
	})
}

const testAccDataSourceIosxeRestconfConfigInterface = `
resource "iosxe_restconf" "test" {
	path = "Cisco-IOS-XE-native:native/banner/login"
	attributes = {
		banner = "My Banner"
	}
}

data "iosxe_restconf" "test" {
	path = "Cisco-IOS-XE-native:native/banner/login"
	depends_on = [iosxe_restconf.test]
}
`
