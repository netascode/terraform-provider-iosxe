---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "iosxe_interface_port_channel Data Source - terraform-provider-iosxe"
subcategory: "Interface"
description: |-
  This data source can read the Interface Port Channel configuration.
---

# iosxe_interface_port_channel (Data Source)

This data source can read the Interface Port Channel configuration.

## Example Usage

```terraform
data "iosxe_interface_port_channel" "example" {
  name = 10
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (Number)

### Optional

- `device` (String) A device name from the provider configuration.

### Read-Only

- `description` (String) Interface specific description
- `helper_addresses` (Attributes List) Specify a destination address for UDP broadcasts (see [below for nested schema](#nestedatt--helper_addresses))
- `id` (String) The path of the retrieved object.
- `ip_dhcp_relay_source_interface` (String) Set source interface for relayed messages
- `ipv4_address` (String)
- `ipv4_address_mask` (String)
- `shutdown` (Boolean) Shutdown the selected interface
- `switchport` (Boolean)
- `vrf_forwarding` (String) Configure forwarding table

<a id="nestedatt--helper_addresses"></a>
### Nested Schema for `helper_addresses`

Read-Only:

- `address` (String)
- `global` (Boolean) Helper-address is global
- `vrf` (String) VRF name for helper-address (if different from interface VRF)


