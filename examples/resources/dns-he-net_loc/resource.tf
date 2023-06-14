# LOC record.
resource "dns-he-net_loc" "example" {
  zone_id = 123456
  domain  = "loc.example.com"
  ttl     = 86400
  data    = "51 56 0.123 N 5 54 0.000 E 4.00m 1.00m 10000.00m 10.00m"
}

