resource "iosxe_interface_loopback" "example" {
  name              = 100
  description       = "My Interface Description"
  shutdown          = false
  vrf_forwarding    = "VRF1"
  ipv4_address      = "200.1.1.1"
  ipv4_address_mask = "255.255.255.255"
  pim_sparse_mode   = true
}
