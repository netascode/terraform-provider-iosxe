resource "iosxe_interface_ospf_process" "example" {
  type       = "GigabitEthernet"
  name       = "2"
  process_id = 1
  area = [
    {
      area_id = "1"
    }
  ]
}
