// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceIosxeInterfacePIM(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxeInterfacePIMPrerequisitesConfig + testAccDataSourceIosxeInterfacePIMConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.iosxe_interface_pim.test", "passive", "false"),
					resource.TestCheckResourceAttr("data.iosxe_interface_pim.test", "dense_mode", "false"),
					resource.TestCheckResourceAttr("data.iosxe_interface_pim.test", "sparse_mode", "true"),
					resource.TestCheckResourceAttr("data.iosxe_interface_pim.test", "sparse_dense_mode", "false"),
					resource.TestCheckResourceAttr("data.iosxe_interface_pim.test", "bfd", "false"),
					resource.TestCheckResourceAttr("data.iosxe_interface_pim.test", "border", "false"),
					resource.TestCheckResourceAttr("data.iosxe_interface_pim.test", "bsr_border", "false"),
					resource.TestCheckResourceAttr("data.iosxe_interface_pim.test", "dr_priority", "10"),
				),
			},
		},
	})
}

const testAccDataSourceIosxeInterfacePIMPrerequisitesConfig = `
resource "iosxe_restconf" "PreReq0" {
  path = "Cisco-IOS-XE-native:native/interface/Loopback=100"
  attributes = {
      "name" = "100"
  }
}

`

const testAccDataSourceIosxeInterfacePIMConfig = `

resource "iosxe_interface_pim" "test" {
  type = "Loopback"
  name = "100"
  passive = false
  dense_mode = false
  sparse_mode = true
  sparse_dense_mode = false
  bfd = false
  border = false
  bsr_border = false
  dr_priority = 10
  depends_on = [iosxe_restconf.PreReq0, ]
}

data "iosxe_interface_pim" "test" {
  type = "Loopback"
  name = "100"
  depends_on = [iosxe_interface_pim.test]
}
`
