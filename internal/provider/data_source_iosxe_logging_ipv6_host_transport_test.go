// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceIosxeLoggingIPv6HostTransport(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxeLoggingIPv6HostTransportConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.iosxe_logging_ipv6_host_transport.test", "transport_udp_ports.0.port_number", "10000"),
					resource.TestCheckResourceAttr("data.iosxe_logging_ipv6_host_transport.test", "transport_tcp_ports.0.port_number", "10001"),
					resource.TestCheckResourceAttr("data.iosxe_logging_ipv6_host_transport.test", "transport_tls_ports.0.port_number", "10002"),
				),
			},
		},
	})
}

const testAccDataSourceIosxeLoggingIPv6HostTransportConfig = `

resource "iosxe_logging_ipv6_host_transport" "test" {
  ipv6_host = "2001::1"
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

data "iosxe_logging_ipv6_host_transport" "test" {
  ipv6_host = "2001::1"
  depends_on = [iosxe_logging_ipv6_host_transport.test]
}
`
