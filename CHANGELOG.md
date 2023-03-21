## 0.1.13

- BREAKING CHANGE: Remove default value of `ip_access_group_in_enable` and `ip_access_group_out_enable` attributes of `iosxe_interface_ethernet` resource
- BREAKING CHANGE: Remove default value of `ip_access_group_in_enable` and `ip_access_group_out_enable` attributes of `iosxe_interface_loopback` resource
- BREAKING CHANGE: Remove default value of `ip_access_group_in_enable` and `ip_access_group_out_enable` attributes of `iosxe_interface_port_channel_subinterface` resource
- BREAKING CHANGE: Remove default value of `ip_access_group_in_enable` and `ip_access_group_out_enable` attributes of `iosxe_interface_port_channel` resource
- BREAKING CHANGE: Remove default value of `ip_access_group_in_enable` and `ip_access_group_out_enable` attributes of `iosxe_interface_vlan` resource

## 0.1.12

- Add iosxe_snmp_server_user resource and data source
- BREAKING CHANGE: Rename access_config_ipv6_acl attribute of iosxe_snmp_server_group resource to access_ipv6_acl

## 0.1.11

- Add iosxe_access_list_standard resource and data source
- Add iosxe_access_list_extended resource and data source
- Add ip_access_group_* attributes to iosxe_interface_ethernet resource and data source
- Add ip_access_group_* attributes to iosxe_interface_loopback resource and data source
- Add ip_access_group_* attributes to iosxe_interface_vlan resource and data source
- Add ip_access_group_* attributes to iosxe_interface_port_channel resource and data source
- Add ip_access_group_* attributes to iosxe_interface_port_channel_subinterface resource and data source

## 0.1.10

- Add iosxe_service resource and data source
- Add iosxe_logging resource and data source
- Add iosxe_logging_ipv4_host_transport resource and data source
- Add iosxe_logging_ipv4_host_vrf_transport resource and data source
- Add iosxe_logging_ipv6_host_transport resource and data source
- Add iosxe_logging_ipv6_host_vrf_transport resource and data source
- Add attributes to iosxe_system resource and data source
- Add iosxe_snmp_server resource and data source
- Add iosxe_snmp_server_group resource and data source

## 0.1.9

- Add iosxe_template resource and data source
- Add source_template option to iosxe_interface_ethernet resource

## 0.1.8

- Delete and recreate a VRF if route distinguisher changes
- Allow provider config without url attribute in case devices attribute is being used

## 0.1.7

- Add multicast_routing_switch attribute to iosxe_system resource

## 0.1.6

- Add DHCP helper options to iosxe_interface_ethernet resource
- Add DHCP helper options to iosxe_interface_vlan resource
- Add DHCP helper options to iosxe_interface_port_channel resource
- Add DHCP helper options to iosxe_interface_port_channel_subinterface resource
- Add iosxe_dhcp resource and data source
- Add iosxe_msdp resource and data source
- Add iosxe_msdp_vrf resource and data source
- Add ip_routing attribute to iosxe_system resource
- Add rp_address attributes to iosxe_pim and iosxe_pim_vrf resources
- Add mtu attribute to iosxe_system resource
- Add additional route target options to iosxe_vrf resource
- Add multicast routing options to iosxe_system resource

## 0.1.5

- Add iosxe_interface_ethernet resource and data source
- Add iosxe_interface_pim resource and data source
- BREAKING CHANGE: Remove pim_sparse_mode attribute from iosxe_interface_loopback resource
- BREAKING CHANGE: Remove pim_sparse_mode attribute from iosxe_interface_vlan resource
- Add iosxe_interface_switchport resource and data source
- Add iosxe_interface_ospf resource and data source
- Add iosxe_interface_ospf_process resource and data source
- Add iosxe_pim resource and data source
- Add iosxe_pim_vrf resource and data source
- Add iosxe_interface_port_channel resource and data source
- Add iosxe_interface_port_channel_subinterface resource and data source

## 0.1.4

- Fix handling of null values in resource list items

## 0.1.3

- Add unnumbered option to iosxe_interface_vlan resource
- Add option to configure lists with iosxe_restconf
- Add iosxe_bgp_ipv4_unicast_vrf_neighbor resource and data source
- BREAKING CHANGE: Rename iosxe_bgp_l2vpn_neighbor resource and data source to iosxe_bgp_l2vpn_evpn_neighbor and remove af_name attribute
- Improve handling of lists within resources

## 0.1.2

- Fix iosxe_vlan_configuration resource to be compatible with older IOS-XE releases
- Optimize RESTCONF retries

## 0.1.1

- Add iosxe_vlan resource and data source
- Add iosxe_vlan_configuration resource and data source
- Add stitching option to vrf route targets
- Add iosxe_interface_vlan resource and data source
- Add iosxe_interface_loopback resource and data source
- Add iosxe_interface_nve resource and data source
- Add iosxe_evpn resource and data source
- Add iosxe_evpn_instance resource and data source
- Add iosxe_bgp resource and data source
- Add iosxe_bgp_address_family_ipv4_vrf resource and data source
- Add iosxe_bgp_address_family_ipv6_vrf resource and data source
- Add iosxe_bgp_neighbor resource and data source
- Add iosxe_bgp_address_family_l2vpn resource and data source
- Add iosxe_bgp_l2vpn_neighbor resource and data source
- Add ipv6_unicast_routing option to iosxe_system resource

## 0.1.0

- Initial release
