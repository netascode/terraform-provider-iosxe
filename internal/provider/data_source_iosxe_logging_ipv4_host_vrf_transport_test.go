// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceIosxeLoggingIPv4HostVRFTransport(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_logging_ipv4_host_vrf_transport.test", "transport_udp_ports.0.port_number", "10000"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_logging_ipv4_host_vrf_transport.test", "transport_tcp_ports.0.port_number", "10001"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_logging_ipv4_host_vrf_transport.test", "transport_tls_ports.0.port_number", "10002"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxeLoggingIPv4HostVRFTransportPrerequisitesConfig + testAccDataSourceIosxeLoggingIPv4HostVRFTransportConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

const testAccDataSourceIosxeLoggingIPv4HostVRFTransportPrerequisitesConfig = `
resource "iosxe_restconf" "PreReq0" {
	path = "Cisco-IOS-XE-native:native/vrf/definition=VRF1"
	delete = false
	attributes = {
		"name" = "VRF1"
		"address-family/ipv4" = ""
		"address-family/ipv6" = ""
	}
}

`

func testAccDataSourceIosxeLoggingIPv4HostVRFTransportConfig() string {
	config := `resource "iosxe_logging_ipv4_host_vrf_transport" "test" {` + "\n"
	config += `	delete_mode = "attributes"` + "\n"
	config += `	ipv4_host = "2.2.2.2"` + "\n"
	config += `	vrf = "VRF1"` + "\n"
	config += `	transport_udp_ports = [{` + "\n"
	config += `		port_number = 10000` + "\n"
	config += `	}]` + "\n"
	config += `	transport_tcp_ports = [{` + "\n"
	config += `		port_number = 10001` + "\n"
	config += `	}]` + "\n"
	config += `	transport_tls_ports = [{` + "\n"
	config += `		port_number = 10002` + "\n"
	config += `	}]` + "\n"
	config += `	depends_on = [iosxe_restconf.PreReq0, ]` + "\n"
	config += `}` + "\n"

	config += `
		data "iosxe_logging_ipv4_host_vrf_transport" "test" {
			ipv4_host = "2.2.2.2"
			vrf = "VRF1"
			depends_on = [iosxe_logging_ipv4_host_vrf_transport.test]
		}
	`
	return config
}
