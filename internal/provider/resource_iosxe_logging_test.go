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

func TestAccIosxeLogging(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_logging.test", "monitor_severity", "informational"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_logging.test", "buffered_size", "16000"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_logging.test", "buffered_severity", "informational"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_logging.test", "console_severity", "informational"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_logging.test", "facility", "local0"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_logging.test", "history_size", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_logging.test", "history_severity", "informational"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_logging.test", "trap", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_logging.test", "trap_severity", "informational"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_logging.test", "origin_id_type", "hostname"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_logging.test", "source_interface", "Loopback0"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_logging.test", "source_interfaces_vrf.0.vrf", "VRF1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_logging.test", "source_interfaces_vrf.0.interface_name", "Loopback100"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_logging.test", "ipv4_hosts.0.ipv4_host", "1.1.1.1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_logging.test", "ipv4_vrf_hosts.0.ipv4_host", "1.1.1.1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_logging.test", "ipv4_vrf_hosts.0.vrf", "VRF1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_logging.test", "ipv6_hosts.0.ipv6_host", "2001::1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_logging.test", "ipv6_vrf_hosts.0.ipv6_host", "2001::1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_logging.test", "ipv6_vrf_hosts.0.vrf", "VRF1"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxeLoggingPrerequisitesConfig + testAccIosxeLoggingConfig_minimum(),
			},
			{
				Config: testAccIosxeLoggingPrerequisitesConfig + testAccIosxeLoggingConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				ResourceName:  "iosxe_logging.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XE-native:native/logging",
			},
		},
	})
}

const testAccIosxeLoggingPrerequisitesConfig = `
resource "iosxe_restconf" "PreReq0" {
	path = "Cisco-IOS-XE-native:native/vrf/definition=VRF1"
	delete = false
	attributes = {
		"name" = "VRF1"
		"address-family/ipv4" = ""
		"address-family/ipv6" = ""
	}
}

resource "iosxe_restconf" "PreReq1" {
	path = "Cisco-IOS-XE-native:native/interface/Loopback=100"
	attributes = {
		"name" = "100"
		"vrf/forwarding" = "VRF1"
	}
	depends_on = [iosxe_restconf.PreReq0, ]
}

`

func testAccIosxeLoggingConfig_minimum() string {
	config := `resource "iosxe_logging" "test" {` + "\n"
	config += `	depends_on = [iosxe_restconf.PreReq0, iosxe_restconf.PreReq1, ]` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxeLoggingConfig_all() string {
	config := `resource "iosxe_logging" "test" {` + "\n"
	config += `	monitor_severity = "informational"` + "\n"
	config += `	buffered_size = 16000` + "\n"
	config += `	buffered_severity = "informational"` + "\n"
	config += `	console_severity = "informational"` + "\n"
	config += `	facility = "local0"` + "\n"
	config += `	history_size = 100` + "\n"
	config += `	history_severity = "informational"` + "\n"
	config += `	trap = true` + "\n"
	config += `	trap_severity = "informational"` + "\n"
	config += `	origin_id_type = "hostname"` + "\n"
	config += `	source_interface = "Loopback0"` + "\n"
	config += `	source_interfaces_vrf = [{` + "\n"
	config += `		vrf = "VRF1"` + "\n"
	config += `		interface_name = "Loopback100"` + "\n"
	config += `	}]` + "\n"
	config += `	ipv4_hosts = [{` + "\n"
	config += `		ipv4_host = "1.1.1.1"` + "\n"
	config += `	}]` + "\n"
	config += `	ipv4_vrf_hosts = [{` + "\n"
	config += `		ipv4_host = "1.1.1.1"` + "\n"
	config += `		vrf = "VRF1"` + "\n"
	config += `	}]` + "\n"
	config += `	ipv6_hosts = [{` + "\n"
	config += `		ipv6_host = "2001::1"` + "\n"
	config += `	}]` + "\n"
	config += `	ipv6_vrf_hosts = [{` + "\n"
	config += `		ipv6_host = "2001::1"` + "\n"
	config += `		vrf = "VRF1"` + "\n"
	config += `	}]` + "\n"
	config += `	depends_on = [iosxe_restconf.PreReq0, iosxe_restconf.PreReq1, ]` + "\n"
	config += `}` + "\n"
	return config
}
