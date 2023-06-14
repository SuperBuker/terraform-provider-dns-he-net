# RP record.
resource "dns-he-net_rp" "example" {
  zone_id = 123456
  domain  = "example.com"
  ttl     = 86400
  data    = "bofher.example.com bofher.example.com"
}

