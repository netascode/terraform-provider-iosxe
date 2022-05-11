resource "iosxe_interface_ethernet" "example" {
  type              = "GigabitEthernet"
  name              = "3"
  description       = "My Interface Description"
  shutdown          = true
  vrf_forwarding    = "VRF1"
  ipv4_address      = "15.1.1.1"
  ipv4_address_mask = "255.255.255.252"
}
