# Static A record.
resource "dns-he-net_a" "example" {
  zone_id = 123456
  domain  = "example.com"
  ttl     = 86400
  data    = "1.2.3.4"
}

# Dynamic A record.
resource "dns-he-net_a" "example" {
  zone_id = 123456
  domain  = "example.com"
  ttl     = 86400
  dynamic = true
}

# Dynamic A record with preset value.
resource "dns-he-net_a" "example" {
  zone_id = 123456
  domain  = "example.com"
  ttl     = 86400
  data    = "1.2.3.4"
  dynamic = true
}
