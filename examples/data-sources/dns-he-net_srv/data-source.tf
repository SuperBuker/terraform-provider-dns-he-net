# Retrieve SRV record.
resource "dns-he-net_srv" "example" {
  id      = 123456789
  zone_id = 123456
}
