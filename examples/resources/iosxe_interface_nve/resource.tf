resource "iosxe_interface_nve" "example" {
  name                           = 1
  description                    = "My Interface Description"
  shutdown                       = false
  host_reachability_protocol_bgp = true
  source_interface_loopback      = 100
  vni_vrfs = [
    {
      vni_range = "10000"
      vrf       = "VRF1"
    }
  ]
  vnis = [
    {
      vni_range            = "10000"
      ipv4_multicast_group = "255.1.1.1"
      ingress_replication  = true
    }
  ]
}
