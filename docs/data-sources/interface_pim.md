---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "iosxe_interface_pim Data Source - terraform-provider-iosxe"
subcategory: "Multicast"
description: |-
  This data source can read the Interface PIM configuration.
---

# iosxe_interface_pim (Data Source)

This data source can read the Interface PIM configuration.

## Example Usage

```terraform
data "iosxe_interface_pim" "example" {
  type = "Loopback"
  name = "100"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String)
- `type` (String) Interface type

### Optional

- `device` (String) A device name from the provider configuration.

### Read-Only

- `bfd` (Boolean) Configure BFD
- `border` (Boolean) Border of PIM domain
- `bsr_border` (Boolean) Border of PIM domain
- `dense_mode` (Boolean) Enable PIM dense-mode operation
- `dr_priority` (Number) PIM router DR priority
- `id` (String) The path of the retrieved object.
- `passive` (Boolean) Enable PIM passive interface operation
- `sparse_dense_mode` (Boolean) Enable PIM sparse-dense-mode operation
- `sparse_mode` (Boolean) Enable PIM sparse-mode operation

