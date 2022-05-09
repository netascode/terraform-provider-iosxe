---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "iosxe_interface_vlan Resource - terraform-provider-iosxe"
subcategory: "Interface"
description: |-
  This resource can manage the Interface VLAN configuration.
---

# iosxe_interface_vlan (Resource)

This resource can manage the Interface VLAN configuration.

## Example Usage

```terraform
resource "iosxe_interface_vlan" "example" {
  name              = 10
  autostate         = false
  description       = "My Interface Description"
  shutdown          = false
  vrf_forwarding    = "VRF1"
  ipv4_address      = "10.1.1.1"
  ipv4_address_mask = "255.255.255.0"
  pim_sparse_mode   = true
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (Number) - Range: `1`-`4094`

### Optional

- `autostate` (Boolean) Enable auto-state determination for VLAN
- `description` (String) Interface specific description
- `device` (String) A device name from the provider configuration.
- `ipv4_address` (String)
- `ipv4_address_mask` (String)
- `pim_sparse_mode` (Boolean) Enable PIM sparse-mode operation
- `shutdown` (Boolean) Shutdown the selected interface
- `unnumbered` (String) Enable IP processing without an explicit address
- `vrf_forwarding` (String) Configure forwarding table

### Read-Only

- `id` (String) The path of the object.

## Import

Import is supported using the following syntax:

```shell
terraform import iosxe_interface_vlan.example "Cisco-IOS-XE-native:native/interface/Vlan=10"
```