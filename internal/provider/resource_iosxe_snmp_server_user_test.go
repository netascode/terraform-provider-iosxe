// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIosxeSNMPServerUser(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxeSNMPServerUserConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("iosxe_snmp_server_user.test", "username", "USER1"),
					resource.TestCheckResourceAttr("iosxe_snmp_server_user.test", "grpname", "GROUP1"),
					resource.TestCheckResourceAttr("iosxe_snmp_server_user.test", "v3_auth_algorithm", "sha"),
					resource.TestCheckResourceAttr("iosxe_snmp_server_user.test", "v3_auth_password", "Cisco123"),
					resource.TestCheckResourceAttr("iosxe_snmp_server_user.test", "v3_auth_priv_aes_algorithm", "128"),
					resource.TestCheckResourceAttr("iosxe_snmp_server_user.test", "v3_auth_priv_aes_password", "Cisco123"),
					resource.TestCheckResourceAttr("iosxe_snmp_server_user.test", "v3_auth_priv_aes_access_ipv6_acl", "V6ACL1"),
					resource.TestCheckResourceAttr("iosxe_snmp_server_user.test", "v3_auth_priv_aes_access_acl_name", "ACL123"),
				),
			},
			{
				ResourceName:  "iosxe_snmp_server_user.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XE-native:native/snmp-server/Cisco-IOS-XE-snmp:user/names=USER1,GROUP1",
			},
		},
	})
}

func testAccIosxeSNMPServerUserConfig_minimum() string {
	return `
	resource "iosxe_snmp_server_user" "test" {
		username = "USER1"
		grpname = "GROUP1"
		v3_auth_algorithm = "sha"
		v3_auth_password = "Cisco123"
	}
	`
}

func testAccIosxeSNMPServerUserConfig_all() string {
	return `
	resource "iosxe_snmp_server_user" "test" {
		username = "USER1"
		grpname = "GROUP1"
		v3_auth_algorithm = "sha"
		v3_auth_password = "Cisco123"
		v3_auth_priv_aes_algorithm = "128"
		v3_auth_priv_aes_password = "Cisco123"
		v3_auth_priv_aes_access_ipv6_acl = "V6ACL1"
		v3_auth_priv_aes_access_acl_name = "ACL123"
	}
	`
}