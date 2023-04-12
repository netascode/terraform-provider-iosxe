---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "iosxe_bgp_address_family_ipv4 Data Source - terraform-provider-iosxe"
subcategory: "BGP"
description: |-
  This data source can read the BGP Address Family IPv4 configuration.
---

# iosxe_bgp_address_family_ipv4 (Data Source)

This data source can read the BGP Address Family IPv4 configuration.

## Example Usage

```terraform
data "iosxe_bgp_address_family_ipv4" "example" {
  asn     = "65000"
  af_name = "unicast"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `af_name` (String)
- `asn` (String)

### Optional

- `device` (String) A device name from the provider configuration.

### Read-Only

- `id` (String) The path of the retrieved object.

