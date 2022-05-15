resource "iosxe_pim_vrf" "example" {
  vrf                               = "VRF1"
  autorp                            = false
  autorp_listener                   = false
  bsr_candidate_loopback            = 100
  bsr_candidate_mask                = 30
  bsr_candidate_priority            = 10
  bsr_candidate_accept_rp_candidate = "10"
  ssm_range                         = "10"
  ssm_default                       = false
  rp_address                        = "19.19.19.19"
  rp_address_override               = false
  rp_address_bidir                  = true
  rp_addresses = [
    {
      access_list = "10"
      rp_address  = "10.10.10.10"
      override    = false
      bidir       = false
    }
  ]
  rp_candidates = [
    {
      interface  = "Loopback100"
      group_list = "10"
      interval   = 100
      priority   = 10
      bidir      = false
    }
  ]
}
