// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceIosxeEVPNInstance(t *testing.T) {
	if os.Getenv("C9000V") == "" {
		t.Skip("skipping test, set environment variable C9000V")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_evpn_instance.test", "vlan_based_replication_type_ingress", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_evpn_instance.test", "vlan_based_replication_type_static", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_evpn_instance.test", "vlan_based_replication_type_p2mp", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_evpn_instance.test", "vlan_based_replication_type_mp2mp", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_evpn_instance.test", "vlan_based_encapsulation", "vxlan"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_evpn_instance.test", "vlan_based_auto_route_target", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_evpn_instance.test", "vlan_based_rd", "10:10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_evpn_instance.test", "vlan_based_route_target", "10:10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_evpn_instance.test", "vlan_based_route_target_both", "10:10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_evpn_instance.test", "vlan_based_route_target_import", "10:10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_evpn_instance.test", "vlan_based_route_target_export", "10:10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_evpn_instance.test", "vlan_based_ip_local_learning_disable", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_evpn_instance.test", "vlan_based_ip_local_learning_enable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_evpn_instance.test", "vlan_based_default_gateway_advertise", "enable"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_evpn_instance.test", "vlan_based_re_originate_route_type5", "true"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxeEVPNInstanceConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

func testAccDataSourceIosxeEVPNInstanceConfig() string {
	config := `resource "iosxe_evpn_instance" "test" {` + "\n"
	config += `	delete_mode = "attributes"` + "\n"
	config += `	evpn_instance_num = 10` + "\n"
	config += `	vlan_based_replication_type_ingress = false` + "\n"
	config += `	vlan_based_replication_type_static = true` + "\n"
	config += `	vlan_based_replication_type_p2mp = false` + "\n"
	config += `	vlan_based_replication_type_mp2mp = false` + "\n"
	config += `	vlan_based_encapsulation = "vxlan"` + "\n"
	config += `	vlan_based_auto_route_target = false` + "\n"
	config += `	vlan_based_rd = "10:10"` + "\n"
	config += `	vlan_based_route_target = "10:10"` + "\n"
	config += `	vlan_based_route_target_both = "10:10"` + "\n"
	config += `	vlan_based_route_target_import = "10:10"` + "\n"
	config += `	vlan_based_route_target_export = "10:10"` + "\n"
	config += `	vlan_based_ip_local_learning_disable = false` + "\n"
	config += `	vlan_based_ip_local_learning_enable = true` + "\n"
	config += `	vlan_based_default_gateway_advertise = "enable"` + "\n"
	config += `	vlan_based_re_originate_route_type5 = true` + "\n"
	config += `}` + "\n"

	config += `
		data "iosxe_evpn_instance" "test" {
			evpn_instance_num = 10
			depends_on = [iosxe_evpn_instance.test]
		}
	`
	return config
}
