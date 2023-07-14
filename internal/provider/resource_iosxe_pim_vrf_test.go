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

func TestAccIosxePIMVRF(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_pim_vrf.test", "vrf", "VRF1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_pim_vrf.test", "autorp", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_pim_vrf.test", "autorp_listener", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_pim_vrf.test", "bsr_candidate_loopback", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_pim_vrf.test", "bsr_candidate_mask", "30"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_pim_vrf.test", "bsr_candidate_priority", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_pim_vrf.test", "ssm_range", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_pim_vrf.test", "ssm_default", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_pim_vrf.test", "rp_address", "19.19.19.19"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_pim_vrf.test", "rp_address_override", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_pim_vrf.test", "rp_address_bidir", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_pim_vrf.test", "cache_rpf_oif", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_pim_vrf.test", "rp_addresses.0.access_list", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_pim_vrf.test", "rp_addresses.0.rp_address", "10.10.10.10"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_pim_vrf.test", "rp_addresses.0.override", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_pim_vrf.test", "rp_addresses.0.bidir", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_pim_vrf.test", "rp_candidates.0.interface", "Loopback100"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_pim_vrf.test", "rp_candidates.0.interval", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_pim_vrf.test", "rp_candidates.0.priority", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_pim_vrf.test", "rp_candidates.0.bidir", "false"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxePIMVRFPrerequisitesConfig + testAccIosxePIMVRFConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				ResourceName:  "iosxe_pim_vrf.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XE-native:native/ip/pim/Cisco-IOS-XE-multicast:vrf=VRF1",
			},
		},
	})
}

const testAccIosxePIMVRFPrerequisitesConfig = `
resource "iosxe_restconf" "PreReq0" {
	path = "Cisco-IOS-XE-native:native/vrf/definition=VRF1"
	delete = false
	attributes = {
		"name" = "VRF1"
		"address-family/ipv4" = ""
	}
}

resource "iosxe_restconf" "PreReq1" {
	path = "Cisco-IOS-XE-native:native/interface/Loopback=100"
	attributes = {
		"name" = "100"
		"vrf/forwarding" = "VRF1"
		"ip/address/primary/address" = "200.200.200.200"
		"ip/address/primary/mask" = "255.255.255.255"
	}
	depends_on = [iosxe_restconf.PreReq0, ]
}

`

func testAccIosxePIMVRFConfig_minimum() string {
	config := `resource "iosxe_pim_vrf" "test" {` + "\n"
	config += `	vrf = "VRF1"` + "\n"
	config += `	depends_on = [iosxe_restconf.PreReq0, iosxe_restconf.PreReq1, ]` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxePIMVRFConfig_all() string {
	config := `resource "iosxe_pim_vrf" "test" {` + "\n"
	config += `	vrf = "VRF1"` + "\n"
	config += `	autorp = false` + "\n"
	config += `	autorp_listener = false` + "\n"
	config += `	bsr_candidate_loopback = 100` + "\n"
	config += `	bsr_candidate_mask = 30` + "\n"
	config += `	bsr_candidate_priority = 10` + "\n"
	config += `	ssm_range = "10"` + "\n"
	config += `	ssm_default = false` + "\n"
	config += `	rp_address = "19.19.19.19"` + "\n"
	config += `	rp_address_override = false` + "\n"
	config += `	rp_address_bidir = false` + "\n"
	config += `	cache_rpf_oif = true` + "\n"
	config += `	rp_addresses = [{` + "\n"
	config += `		access_list = "10"` + "\n"
	config += `		rp_address = "10.10.10.10"` + "\n"
	config += `		override = false` + "\n"
	config += `		bidir = false` + "\n"
	config += `	}]` + "\n"
	config += `	rp_candidates = [{` + "\n"
	config += `		interface = "Loopback100"` + "\n"
	config += `		interval = 100` + "\n"
	config += `		priority = 10` + "\n"
	config += `		bidir = false` + "\n"
	config += `	}]` + "\n"
	config += `	depends_on = [iosxe_restconf.PreReq0, iosxe_restconf.PreReq1, ]` + "\n"
	config += `}` + "\n"
	return config
}
