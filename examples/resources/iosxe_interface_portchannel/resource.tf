resource "iosxe_interface_portchannel" "example" {
  name              = 100
  description       = "My Interface Description"
  shutdown          = false
  vrf_forwarding    = "VRF1"
  ipv4_address      = "192.0.2.1"
  ipv4_address_mask = "255.255.255.0"
}
