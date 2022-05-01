resource "iosxe_interface_vlan" "example" {
  name              = 10
  autostate         = false
  description       = "My Interface Description"
  shutdown          = false
  vrf_forwarding    = "VRF1"
  ipv4_address      = "10.1.1.1"
  ipv4_address_mask = "255.255.255.0"
  pim_sparse_mode   = true
}
