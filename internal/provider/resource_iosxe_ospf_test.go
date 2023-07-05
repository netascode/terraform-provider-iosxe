// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIosxeOSPF(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_ospf.test", "process_id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_ospf.test", "bfd_all_interfaces", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_ospf.test", "default_information_originate", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_ospf.test", "default_information_originate_always", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_ospf.test", "default_metric", "21"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_ospf.test", "distance", "120"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_ospf.test", "domain_tag", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_ospf.test", "neighbor.0.ip", "2.2.2.2"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_ospf.test", "neighbor.0.priority", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_ospf.test", "neighbor.0.cost", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_ospf.test", "network.0.ip", "3.3.3.0"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_ospf.test", "network.0.wildcard", "0.0.0.255"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_ospf.test", "network.0.area", "0"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_ospf.test", "priority", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_ospf.test", "router_id", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_ospf.test", "shutdown", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_ospf.test", "summary_address.0.ip", "3.3.3.0"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_ospf.test", "summary_address.0.mask", "255.255.255.0"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxeOSPFConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				ResourceName:  "iosxe_ospf.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XE-native:native/router/Cisco-IOS-XE-ospf:router-ospf/ospf/process-id=1",
			},
		},
	})
}

func testAccIosxeOSPFConfig_minimum() string {
	config := `resource "iosxe_ospf" "test" {` + "\n"
	config += `	process_id = 1` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxeOSPFConfig_all() string {
	config := `resource "iosxe_ospf" "test" {` + "\n"
	config += `	process_id = 1` + "\n"
	config += `	bfd_all_interfaces = true` + "\n"
	config += `	default_information_originate = true` + "\n"
	config += `	default_information_originate_always = true` + "\n"
	config += `	default_metric = 21` + "\n"
	config += `	distance = 120` + "\n"
	config += `	domain_tag = 10` + "\n"
	config += `	neighbor = [{` + "\n"
	config += `		ip = "2.2.2.2"` + "\n"
	config += `		priority = 10` + "\n"
	config += `		cost = 100` + "\n"
	config += `	}]` + "\n"
	config += `	network = [{` + "\n"
	config += `		ip = "3.3.3.0"` + "\n"
	config += `		wildcard = "0.0.0.255"` + "\n"
	config += `		area = "0"` + "\n"
	config += `	}]` + "\n"
	config += `	priority = 100` + "\n"
	config += `	router_id = "1.2.3.4"` + "\n"
	config += `	shutdown = false` + "\n"
	config += `	summary_address = [{` + "\n"
	config += `		ip = "3.3.3.0"` + "\n"
	config += `		mask = "255.255.255.0"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}
