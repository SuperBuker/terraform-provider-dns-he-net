# AFSDB record.
resource "dns-he-net_afsdb" "example" {
  zone_id = 123456
  domain  = "example.com"
  ttl     = 86400
  data    = "2 green.example.com"
}

