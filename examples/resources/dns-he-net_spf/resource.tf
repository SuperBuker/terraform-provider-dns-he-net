# SPF record.
resource "dns-he-net_spf" "example" {
  zone_id = 123456
  domain  = "example.com"
  ttl     = 86400
  data    = "\"v=spf1 include:_spf.email.com ~all\""
}

