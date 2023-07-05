// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceIosxeMSDP(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_msdp.test", "originator_id", "Loopback100"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_msdp.test", "passwords.0.addr", "10.1.1.1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_msdp.test", "passwords.0.encryption", "0"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_msdp.test", "passwords.0.password", "Cisco123"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_msdp.test", "peers.0.addr", "10.1.1.1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_msdp.test", "peers.0.remote_as", "65000"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_msdp.test", "peers.0.connect_source_loopback", "100"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxeMSDPPrerequisitesConfig + testAccDataSourceIosxeMSDPConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

const testAccDataSourceIosxeMSDPPrerequisitesConfig = `
resource "iosxe_restconf" "PreReq0" {
	path = "Cisco-IOS-XE-native:native/interface/Loopback=100"
	attributes = {
		"name" = "100"
	}
}

`

func testAccDataSourceIosxeMSDPConfig() string {
	config := `resource "iosxe_msdp" "test" {` + "\n"
	config += `	delete_mode = "attributes"\n`
	config += `	originator_id = "Loopback100"` + "\n"
	config += `	passwords = [{` + "\n"
	config += `		addr = "10.1.1.1"` + "\n"
	config += `		encryption = 0` + "\n"
	config += `		password = "Cisco123"` + "\n"
	config += `	}]` + "\n"
	config += `	peers = [{` + "\n"
	config += `		addr = "10.1.1.1"` + "\n"
	config += `		remote_as = 65000` + "\n"
	config += `		connect_source_loopback = 100` + "\n"
	config += `	}]` + "\n"
	config += `	depends_on = [iosxe_restconf.PreReq0, ]` + "\n"
	config += `}` + "\n"

	config += `
		data "iosxe_msdp" "test" {
			depends_on = [iosxe_msdp.test]
		}
	`
	return config
}
