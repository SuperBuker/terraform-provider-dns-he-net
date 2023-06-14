# NS record.
resource "dns-he-net_ns" "example" {
  zone_id = 123456
  domain  = "example.com"
  ttl     = 172800
  data    = "ns2.he.net"
}

