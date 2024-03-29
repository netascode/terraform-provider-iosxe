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
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIosxeInterfaceVLAN(t *testing.T) {
	if os.Getenv("C9000V") == "" {
		t.Skip("skipping test, set environment variable C9000V")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_vlan.test", "name", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_vlan.test", "autostate", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_vlan.test", "description", "My Interface Description"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_vlan.test", "shutdown", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_vlan.test", "ip_proxy_arp", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_vlan.test", "ip_redirects", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_vlan.test", "unreachables", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_vlan.test", "vrf_forwarding", "VRF1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_vlan.test", "ipv4_address", "10.1.1.1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_vlan.test", "ipv4_address_mask", "255.255.255.0"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_vlan.test", "ip_dhcp_relay_source_interface", "Loopback100"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_vlan.test", "ip_access_group_in", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_vlan.test", "ip_access_group_in_enable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_vlan.test", "ip_access_group_out", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_vlan.test", "ip_access_group_out_enable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_vlan.test", "helper_addresses.0.address", "10.10.10.10"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_vlan.test", "helper_addresses.0.global", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_vlan.test", "helper_addresses.0.vrf", "VRF1"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxeInterfaceVLANPrerequisitesConfig + testAccIosxeInterfaceVLANConfig_minimum(),
			},
			{
				Config: testAccIosxeInterfaceVLANPrerequisitesConfig + testAccIosxeInterfaceVLANConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				ResourceName:  "iosxe_interface_vlan.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XE-native:native/interface/Vlan=10",
			},
		},
	})
}

const testAccIosxeInterfaceVLANPrerequisitesConfig = `
resource "iosxe_restconf" "PreReq0" {
	path = "Cisco-IOS-XE-native:native/vrf/definition=VRF1"
	delete = false
	attributes = {
		"name" = "VRF1"
		"address-family/ipv4" = ""
	}
}

`

func testAccIosxeInterfaceVLANConfig_minimum() string {
	config := `resource "iosxe_interface_vlan" "test" {` + "\n"
	config += `	name = 10` + "\n"
	config += `	depends_on = [iosxe_restconf.PreReq0, ]` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxeInterfaceVLANConfig_all() string {
	config := `resource "iosxe_interface_vlan" "test" {` + "\n"
	config += `	name = 10` + "\n"
	config += `	autostate = false` + "\n"
	config += `	description = "My Interface Description"` + "\n"
	config += `	shutdown = false` + "\n"
	config += `	ip_proxy_arp = false` + "\n"
	config += `	ip_redirects = false` + "\n"
	config += `	unreachables = false` + "\n"
	config += `	vrf_forwarding = "VRF1"` + "\n"
	config += `	ipv4_address = "10.1.1.1"` + "\n"
	config += `	ipv4_address_mask = "255.255.255.0"` + "\n"
	config += `	ip_dhcp_relay_source_interface = "Loopback100"` + "\n"
	config += `	ip_access_group_in = "1"` + "\n"
	config += `	ip_access_group_in_enable = true` + "\n"
	config += `	ip_access_group_out = "1"` + "\n"
	config += `	ip_access_group_out_enable = true` + "\n"
	config += `	helper_addresses = [{` + "\n"
	config += `		address = "10.10.10.10"` + "\n"
	config += `		global = false` + "\n"
	config += `		vrf = "VRF1"` + "\n"
	config += `	}]` + "\n"
	config += `	depends_on = [iosxe_restconf.PreReq0, ]` + "\n"
	config += `}` + "\n"
	return config
}
