// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIosxeSNMPServerUser(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server_user.test", "username", "USER1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server_user.test", "grpname", "GROUP1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server_user.test", "v3_auth_algorithm", "sha"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server_user.test", "v3_auth_password", "Cisco123"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server_user.test", "v3_auth_priv_aes_algorithm", "128"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server_user.test", "v3_auth_priv_aes_password", "Cisco123"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server_user.test", "v3_auth_priv_aes_access_ipv6_acl", "V6ACL1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server_user.test", "v3_auth_priv_aes_access_acl_name", "ACL123"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxeSNMPServerUserConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
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
	config := `resource "iosxe_snmp_server_user" "test" {` + "\n"
	config += `	username = "USER1"` + "\n"
	config += `	grpname = "GROUP1"` + "\n"
	config += `	v3_auth_algorithm = "sha"` + "\n"
	config += `	v3_auth_password = "Cisco123"` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxeSNMPServerUserConfig_all() string {
	config := `resource "iosxe_snmp_server_user" "test" {` + "\n"
	config += `	username = "USER1"` + "\n"
	config += `	grpname = "GROUP1"` + "\n"
	config += `	v3_auth_algorithm = "sha"` + "\n"
	config += `	v3_auth_password = "Cisco123"` + "\n"
	config += `	v3_auth_priv_aes_algorithm = "128"` + "\n"
	config += `	v3_auth_priv_aes_password = "Cisco123"` + "\n"
	config += `	v3_auth_priv_aes_access_ipv6_acl = "V6ACL1"` + "\n"
	config += `	v3_auth_priv_aes_access_acl_name = "ACL123"` + "\n"
	config += `}` + "\n"
	return config
}
