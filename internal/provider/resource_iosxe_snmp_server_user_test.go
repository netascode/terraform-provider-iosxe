// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

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
