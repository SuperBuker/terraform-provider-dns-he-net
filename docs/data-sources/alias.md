---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dns-he-net_alias Data Source - dns-he-net"
subcategory: ""
description: |-
  DNS ALIAS record
---

# dns-he-net_alias (Data Source)

DNS ALIAS record

## Example Usage

```terraform
# Retrieve ALIAS record.
resource "dns-he-net_alias" "example" {
  id      = 123456789
  zone_id = 123456
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (Number) dns.he.net record id
- `zone_id` (Number) dns.he.net zone id

### Read-Only

- `data` (String) Value of the DNS record: a hostname
- `domain` (String) Name of the DNS record
- `ttl` (Number) Time-To-Live of the DNS record
