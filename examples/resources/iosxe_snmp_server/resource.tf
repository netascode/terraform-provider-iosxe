resource "iosxe_snmp_server" "example" {
  chassis_id                                = "R1"
  contact                                   = "Contact1"
  ifindex_persist                           = true
  location                                  = "Location1"
  packetsize                                = 2000
  queue_length                              = 100
  source_interface_informs_gigabit_ethernet = "1"
  source_interface_traps_gigabit_ethernet   = "1"
  trap_source_gigabit_ethernet              = "1"
  snmp_communities = [
    {
      name             = "COM1"
      view             = "VIEW1"
      permission       = "ro"
      ipv6             = "ACL1"
      access_list_name = "1"
    }
  ]
  contexts = [
    {
      name = "CON1"
    }
  ]
  views = [
    {
      name    = "VIEW1"
      mib     = "interfaces"
      inc_exl = "included"
    }
  ]
}