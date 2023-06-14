# MX record.
resource "dns-he-net_mx" "example" {
  zone_id  = 123456
  domain   = "example.com"
  ttl      = 86400
  priority = 1
  data     = "mx.example.com"
}

