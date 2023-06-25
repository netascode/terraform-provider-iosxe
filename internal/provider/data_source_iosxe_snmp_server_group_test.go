// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceIosxeSNMPServerGroup(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxeSNMPServerGroupConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.iosxe_snmp_server_group.test", "v3_security.0.security_level", "priv"),
					resource.TestCheckResourceAttr("data.iosxe_snmp_server_group.test", "v3_security.0.context_node", "CON1"),
					resource.TestCheckResourceAttr("data.iosxe_snmp_server_group.test", "v3_security.0.match_node", "exact"),
					resource.TestCheckResourceAttr("data.iosxe_snmp_server_group.test", "v3_security.0.read_node", "VIEW1"),
					resource.TestCheckResourceAttr("data.iosxe_snmp_server_group.test", "v3_security.0.write_node", "VIEW2"),
					resource.TestCheckResourceAttr("data.iosxe_snmp_server_group.test", "v3_security.0.notify_node", "VIEW3"),
					resource.TestCheckResourceAttr("data.iosxe_snmp_server_group.test", "v3_security.0.access_ipv6_acl", "V6ACL1"),
					resource.TestCheckResourceAttr("data.iosxe_snmp_server_group.test", "v3_security.0.access_acl_name", "ACL1"),
				),
			},
		},
	})
}

const testAccDataSourceIosxeSNMPServerGroupConfig = `

resource "iosxe_snmp_server_group" "test" {
	delete_mode = "attributes"
	name = "GROUP1"
	v3_security = [{
		security_level = "priv"
		context_node = "CON1"
		match_node = "exact"
		read_node = "VIEW1"
		write_node = "VIEW2"
		notify_node = "VIEW3"
		access_ipv6_acl = "V6ACL1"
		access_acl_name = "ACL1"
	}]
}

data "iosxe_snmp_server_group" "test" {
	name = "GROUP1"
	depends_on = [iosxe_snmp_server_group.test]
}
`
