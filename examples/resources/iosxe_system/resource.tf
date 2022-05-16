resource "iosxe_system" "example" {
  hostname                      = "ROUTER-1"
  ip_routing                    = true
  ipv6_unicast_routing          = true
  multicast_routing             = true
  multicast_routing_distributed = true
  multicast_routing_vrfs = [
    {
      vrf         = "VRF1"
      distributed = true
    }
  ]
}
