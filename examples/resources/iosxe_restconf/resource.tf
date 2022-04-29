resource "iosxe_restconf" "example" {
  path = "Cisco-IOS-XE-native:native/banner/login"
  attributes = {
    banner = "My Banner"
  }
}
