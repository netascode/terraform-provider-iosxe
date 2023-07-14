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

func TestAccIosxeBGPIPv4UnicastVRFNeighbor(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_bgp_ipv4_unicast_vrf_neighbor.test", "ip", "3.3.3.3"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_bgp_ipv4_unicast_vrf_neighbor.test", "remote_as", "65000"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_bgp_ipv4_unicast_vrf_neighbor.test", "description", "BGP Neighbor Description"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_bgp_ipv4_unicast_vrf_neighbor.test", "shutdown", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_bgp_ipv4_unicast_vrf_neighbor.test", "update_source_loopback", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_bgp_ipv4_unicast_vrf_neighbor.test", "activate", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_bgp_ipv4_unicast_vrf_neighbor.test", "send_community", "both"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_bgp_ipv4_unicast_vrf_neighbor.test", "route_reflector_client", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_bgp_ipv4_unicast_vrf_neighbor.test", "route_maps.0.in_out", "in"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_bgp_ipv4_unicast_vrf_neighbor.test", "route_maps.0.route_map_name", "RM1"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxeBGPIPv4UnicastVRFNeighborPrerequisitesConfig + testAccIosxeBGPIPv4UnicastVRFNeighborConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				ResourceName:  "iosxe_bgp_ipv4_unicast_vrf_neighbor.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XE-native:native/router/Cisco-IOS-XE-bgp:bgp=65000/address-family/with-vrf/ipv4=unicast/vrf=VRF1/ipv4-unicast/neighbor=3.3.3.3",
			},
		},
	})
}

const testAccIosxeBGPIPv4UnicastVRFNeighborPrerequisitesConfig = `
resource "iosxe_restconf" "PreReq0" {
	path = "Cisco-IOS-XE-native:native/vrf/definition=VRF1"
	delete = false
	attributes = {
		"name" = "VRF1"
		"rd" = "1:1"
		"address-family/ipv4" = ""
	}
}

resource "iosxe_restconf" "PreReq1" {
	path = "Cisco-IOS-XE-native:native/router/Cisco-IOS-XE-bgp:bgp=65000"
	attributes = {
		"id" = "65000"
	}
}

resource "iosxe_restconf" "PreReq2" {
	path = "Cisco-IOS-XE-native:native/router/Cisco-IOS-XE-bgp:bgp=65000/address-family/with-vrf/ipv4=unicast"
	attributes = {
		"af-name" = "unicast"
	}
	lists = [
		{
			name = "vrf"
			key = "name"
			items = [
				{
					"name" = "VRF1"
				},
			] 
		},
	]
	depends_on = [iosxe_restconf.PreReq0, iosxe_restconf.PreReq1, ]
}

resource "iosxe_restconf" "PreReq3" {
	path = "Cisco-IOS-XE-native:native/interface/Loopback=100"
	attributes = {
		"name" = "100"
	}
}

`

func testAccIosxeBGPIPv4UnicastVRFNeighborConfig_minimum() string {
	config := `resource "iosxe_bgp_ipv4_unicast_vrf_neighbor" "test" {` + "\n"
	config += `	asn = "65000"` + "\n"
	config += `	vrf = "VRF1"` + "\n"
	config += `	ip = "3.3.3.3"` + "\n"
	config += `	depends_on = [iosxe_restconf.PreReq0, iosxe_restconf.PreReq1, iosxe_restconf.PreReq2, iosxe_restconf.PreReq3, ]` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxeBGPIPv4UnicastVRFNeighborConfig_all() string {
	config := `resource "iosxe_bgp_ipv4_unicast_vrf_neighbor" "test" {` + "\n"
	config += `	asn = "65000"` + "\n"
	config += `	vrf = "VRF1"` + "\n"
	config += `	ip = "3.3.3.3"` + "\n"
	config += `	remote_as = "65000"` + "\n"
	config += `	description = "BGP Neighbor Description"` + "\n"
	config += `	shutdown = false` + "\n"
	config += `	update_source_loopback = "100"` + "\n"
	config += `	activate = true` + "\n"
	config += `	send_community = "both"` + "\n"
	config += `	route_reflector_client = false` + "\n"
	config += `	route_maps = [{` + "\n"
	config += `		in_out = "in"` + "\n"
	config += `		route_map_name = "RM1"` + "\n"
	config += `	}]` + "\n"
	config += `	depends_on = [iosxe_restconf.PreReq0, iosxe_restconf.PreReq1, iosxe_restconf.PreReq2, iosxe_restconf.PreReq3, ]` + "\n"
	config += `}` + "\n"
	return config
}
