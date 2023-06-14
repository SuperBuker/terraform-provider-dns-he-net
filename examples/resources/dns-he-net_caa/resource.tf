# CAA record.
resource "dns-he-net_caa" "example" {
  zone_id = 123456
  domain  = "example.com"
  ttl     = 86400
  data    = "0 iodef \"bofher@example.com\""
}

