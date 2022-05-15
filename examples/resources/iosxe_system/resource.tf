resource "iosxe_system" "example" {
  hostname             = "ROUTER-1"
  ip_routing           = true
  ipv6_unicast_routing = true
}
