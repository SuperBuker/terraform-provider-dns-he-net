# Static A record.
resource "dns-he-net_aaaa" "example" {
  zone_id = 123456
  domain  = "example.com"
  ttl     = 86400
  data    = "::1"
}

# Dynamic A record.
resource "dns-he-net_aaaa" "example" {
  zone_id = 123456
  domain  = "example.com"
  ttl     = 86400
  dynamic = true
}

# Dynamic A record with preset value.
resource "dns-he-net_aaaa" "example" {
  zone_id = 123456
  domain  = "example.com"
  ttl     = 86400
  data    = "::1"
  dynamic = true
}
