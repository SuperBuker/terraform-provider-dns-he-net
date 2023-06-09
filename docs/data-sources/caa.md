---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dns-he-net_caa Data Source - dns-he-net"
subcategory: ""
description: |-
  DNS CAA record
---

# dns-he-net_caa (Data Source)

DNS CAA record

## Example Usage

```terraform
# Retrieve CAA record.
resource "dns-he-net_caa" "example" {
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

- `data` (String) Value of the DNS record: flags, tag and value
- `domain` (String) Name of the DNS record
- `ttl` (Number) Time-To-Live of the DNS record
