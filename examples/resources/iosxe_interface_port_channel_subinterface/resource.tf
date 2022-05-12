resource "iosxe_interface_port_channel_subinterface" "example" {
  name                        = "10.666"
  description                 = "My Interface Description"
  shutdown                    = false
  vrf_forwarding              = "VRF1"
  ipv4_address                = "192.0.2.2"
  ipv4_address_mask           = "255.255.255.0"
  encapsulation_dot1q_vlan_id = 666
}
