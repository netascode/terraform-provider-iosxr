---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "iosxr_interface Resource - terraform-provider-iosxr"
subcategory: "Interface"
description: |-
  This resource can manage the Interface configuration.
---

# iosxr_interface (Resource)

This resource can manage the Interface configuration.

## Example Usage

```terraform
resource "iosxr_interface" "example" {
  interface_name          = "GigabitEthernet0/0/0/1"
  l2transport             = false
  point_to_point          = false
  multipoint              = false
  shutdown                = true
  mtu                     = 9000
  bandwidth               = 100000
  description             = "My Interface Description"
  vrf                     = "VRF1"
  ipv4_address            = "1.1.1.1"
  ipv4_netmask            = "255.255.255.0"
  ipv6_link_local_address = "fe80::1"
  ipv6_link_local_zone    = "0"
  ipv6_autoconfig         = false
  ipv6_enable             = true
  ipv6_addresses = [
    {
      address       = "2001::1"
      prefix_length = 64
      zone          = "0"
    }
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `interface_name` (String) Interface configuration subcommands

### Optional

- `bandwidth` (Number) Set the bandwidth of an interface
  - Range: `0`-`9223372036854775807`
- `description` (String) Set description for this interface
- `device` (String) A device name from the provider configuration.
- `ipv4_address` (String) IP address
- `ipv4_netmask` (String) IP subnet mask
- `ipv6_addresses` (Attributes List) IPv6 address (see [below for nested schema](#nestedatt--ipv6_addresses))
- `ipv6_autoconfig` (Boolean) Enable slaac on Mgmt interface
- `ipv6_enable` (Boolean) Enable IPv6 on interface
- `ipv6_link_local_address` (String) IPv6 address
- `ipv6_link_local_zone` (String) IPv6 address zone
  - Default value: `0`
- `l2transport` (Boolean) l2transport sub-interface
- `mtu` (Number) Set the MTU on an interface
  - Range: `64`-`65535`
- `multipoint` (Boolean) multipoint sub-interface
- `point_to_point` (Boolean) point-to-point sub-interface
- `shutdown` (Boolean) shutdown the given interface
- `unnumbered` (String) Enable IP processing without an explicit address
- `vrf` (String) Set VRF in which the interface operates

### Read-Only

- `id` (String) The path of the object.

<a id="nestedatt--ipv6_addresses"></a>
### Nested Schema for `ipv6_addresses`

Optional:

- `address` (String) IPv6 name or address
- `prefix_length` (Number) Prefix length in bits
  - Range: `0`-`128`
- `zone` (String) IPv6 address zone
  - Default value: `0`

## Import

Import is supported using the following syntax:

```shell
terraform import iosxr_interface.example "Cisco-IOS-XR-um-interface-cfg:/interfaces/interface[interface-name=GigabitEthernet0/0/0/1]"
```
