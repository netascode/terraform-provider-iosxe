// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIosxeInterfaceNVE(t *testing.T) {
	if os.Getenv("C9000V") == "" {
		t.Skip("skipping test, set environment variable C9000V")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_nve.test", "name", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_nve.test", "description", "My Interface Description"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_nve.test", "shutdown", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_nve.test", "host_reachability_protocol_bgp", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_nve.test", "source_interface_loopback", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_nve.test", "vnis.0.vni_range", "10000"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_nve.test", "vnis.0.ipv4_multicast_group", "225.1.1.1"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxeInterfaceNVEConfig_minimum(),
			},
			{
				Config: testAccIosxeInterfaceNVEConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				ResourceName:  "iosxe_interface_nve.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XE-native:native/interface/nve=1",
			},
		},
	})
}

func testAccIosxeInterfaceNVEConfig_minimum() string {
	config := `resource "iosxe_interface_nve" "test" {` + "\n"
	config += `	name = 1` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxeInterfaceNVEConfig_all() string {
	config := `resource "iosxe_interface_nve" "test" {` + "\n"
	config += `	name = 1` + "\n"
	config += `	description = "My Interface Description"` + "\n"
	config += `	shutdown = false` + "\n"
	config += `	host_reachability_protocol_bgp = true` + "\n"
	config += `	source_interface_loopback = 100` + "\n"
	config += `	vnis = [{` + "\n"
	config += `		vni_range = "10000"` + "\n"
	config += `		ipv4_multicast_group = "225.1.1.1"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}
