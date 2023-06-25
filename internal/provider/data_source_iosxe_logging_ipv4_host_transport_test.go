// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceIosxeLoggingIPv4HostTransport(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxeLoggingIPv4HostTransportConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.iosxe_logging_ipv4_host_transport.test", "transport_udp_ports.0.port_number", "10000"),
					resource.TestCheckResourceAttr("data.iosxe_logging_ipv4_host_transport.test", "transport_tcp_ports.0.port_number", "10001"),
					resource.TestCheckResourceAttr("data.iosxe_logging_ipv4_host_transport.test", "transport_tls_ports.0.port_number", "10002"),
				),
			},
		},
	})
}

const testAccDataSourceIosxeLoggingIPv4HostTransportConfig = `

resource "iosxe_logging_ipv4_host_transport" "test" {
	delete_mode = "attributes"
	ipv4_host = "2.2.2.2"
	transport_udp_ports = [{
		port_number = 10000
	}]
	transport_tcp_ports = [{
		port_number = 10001
	}]
	transport_tls_ports = [{
		port_number = 10002
	}]
}

data "iosxe_logging_ipv4_host_transport" "test" {
	ipv4_host = "2.2.2.2"
	depends_on = [iosxe_logging_ipv4_host_transport.test]
}
`
