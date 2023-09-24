---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dns-he-net_sshfp Data Source - dns-he-net"
subcategory: ""
description: |-
  DNS SSHFP record
---

# dns-he-net_sshfp (Data Source)

DNS SSHFP record

## Example Usage

```terraform
# Retrieve SSHFP record.
resource "dns-he-net_sshfp" "example" {
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

- `data` (String) Value of the DNS record: algorithm, (hash) type and fingerprint
- `domain` (String) Name of the DNS record
- `ttl` (Number) Time-To-Live of the DNS record