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

func TestAccDataSourceIosxeInterfacePIM(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_pim.test", "passive", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_pim.test", "dense_mode", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_pim.test", "sparse_mode", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_pim.test", "sparse_dense_mode", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_pim.test", "bfd", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_pim.test", "border", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_pim.test", "bsr_border", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_pim.test", "dr_priority", "10"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxeInterfacePIMPrerequisitesConfig + testAccDataSourceIosxeInterfacePIMConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
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

func testAccDataSourceIosxeInterfacePIMConfig() string {
	config := `resource "iosxe_interface_pim" "test" {` + "\n"
	config += `	type = "Loopback"` + "\n"
	config += `	name = "100"` + "\n"
	config += `	passive = false` + "\n"
	config += `	dense_mode = false` + "\n"
	config += `	sparse_mode = true` + "\n"
	config += `	sparse_dense_mode = false` + "\n"
	config += `	bfd = false` + "\n"
	config += `	border = false` + "\n"
	config += `	bsr_border = false` + "\n"
	config += `	dr_priority = 10` + "\n"
	config += `	depends_on = [iosxe_restconf.PreReq0, ]` + "\n"
	config += `}` + "\n"

	config += `
		data "iosxe_interface_pim" "test" {
			type = "Loopback"
			name = "100"
			depends_on = [iosxe_interface_pim.test]
		}
	`
	return config
}
