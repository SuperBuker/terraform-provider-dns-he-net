# PTR record.
resource "dns-he-net_ptr" "example" {
  zone_id = 123456
  domain  = "ptr.example.com"
  ttl     = 86400
  data    = "example.com"
}

