resource "iosxe_interface_vlan" "example" {
  name                           = 10
  autostate                      = false
  description                    = "My Interface Description"
  shutdown                       = false
  vrf_forwarding                 = "VRF1"
  ipv4_address                   = "10.1.1.1"
  ipv4_address_mask              = "255.255.255.0"
  ip_dhcp_relay_source_interface = "Loopback100"
  helper_addresses = [
    {
      address = "10.10.10.10"
      global  = false
      vrf     = "VRF1"
    }
  ]
}
