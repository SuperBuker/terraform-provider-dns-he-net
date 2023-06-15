# HINFO record.
resource "dns-he-net_hinfo" "example" {
  zone_id = 123456
  domain  = "hinfo.example.com"
  ttl     = 86400
  data    = "\"armv7 Linux\""
}

