resource "iosxe_interface_ethernet" "example" {
  type                           = "GigabitEthernet"
  name                           = "3"
  description                    = "My Interface Description"
  shutdown                       = false
  ipv4_address                   = "15.1.1.1"
  ipv4_address_mask              = "255.255.255.252"
  ip_dhcp_relay_source_interface = "Loopback100"
  helper_addresses = [
    {
      address = "10.10.10.10"
      global  = false
      vrf     = "VRF1"
    }
  ]
  source_template = [
    {
      template_name = "TEMP1"
      merge         = false
    }
  ]
}
