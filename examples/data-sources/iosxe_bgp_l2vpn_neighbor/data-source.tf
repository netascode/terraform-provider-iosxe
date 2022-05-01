data "iosxe_bgp_l2vpn_neighbor" "example" {
  asn     = "65000"
  af_name = "evpn"
  ip      = "3.3.3.3"
}
