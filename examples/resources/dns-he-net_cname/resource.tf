# CNAME record.
resource "dns-he-net_cname" "example" {
  zone_id = 123456
  domain  = "cname.example.com"
  ttl     = 86400
  data    = "example.com"
}

