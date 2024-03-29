---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "iosxe_mdt_subscription Data Source - terraform-provider-iosxe"
subcategory: "MDT"
description: |-
  This data source can read the MDT Subscription configuration.
---

# iosxe_mdt_subscription (Data Source)

This data source can read the MDT Subscription configuration.

## Example Usage

```terraform
data "iosxe_mdt_subscription" "example" {
  subscription_id = 101
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `subscription_id` (Number) Unique subscription identifier.

### Optional

- `device` (String) A device name from the provider configuration.

### Read-Only

- `encoding` (String) Update notification encoding
- `filter_xpath` (String) XPath expression describing the set of objects wanted as part of the subscription
- `id` (String) The path of the retrieved object.
- `receivers` (Attributes List) Configuration of receivers of configured subscriptions. Use of this is deprecated. Use mdt-receiver-names instead. (see [below for nested schema](#nestedatt--receivers))
- `source_address` (String) The source address for the notifications
- `source_vrf` (String) Network instance name for the VRF
- `stream` (String) The name of the event stream being subscribed to
- `update_policy_on_change` (Boolean) If true, there is no initial update notification with the current value of all the data. NOT CURRENTLY SUPPORTED. If specified, must be false
- `update_policy_periodic` (Number) Period of update notifications in hundredths of a second

<a id="nestedatt--receivers"></a>
### Nested Schema for `receivers`

Read-Only:

- `address` (String) IP address of the receiver
- `port` (Number) Network port of the receiver
- `protocol` (String) Receiver transport protocol.
