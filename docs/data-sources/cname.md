---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dns-he-net_cname Data Source - dns-he-net"
subcategory: ""
description: |-
  DNS CNAME record
---

# dns-he-net_cname (Data Source)

DNS CNAME record

## Example Usage

```terraform
# Retrieve CNAME record.
resource "dns-he-net_cname" "example" {
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
