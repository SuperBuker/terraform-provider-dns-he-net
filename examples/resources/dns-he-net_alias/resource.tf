# ALIAS record.
resource "dns-he-net_alias" "example" {
  zone_id = 123456
  domain  = "alias.example.com"
  ttl     = 86400
  data    = "example.com"
}

