resource "iosxe_bgp_l2vpn_neighbor" "example" {
  asn                    = "65000"
  af_name                = "evpn"
  ip                     = "3.3.3.3"
  activate               = true
  send_community         = "both"
  route_reflector_client = false
}
