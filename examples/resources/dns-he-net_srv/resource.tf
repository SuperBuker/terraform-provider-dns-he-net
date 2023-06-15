# SRV record.
resource "dns-he-net_srv" "example" {
  zone_id  = 123456
  domain   = "_bofher._tcp.example.com"
  ttl      = 86400
  priority = 0
  weight   = 0
  port     = 22
  target   = "example.com"
}

