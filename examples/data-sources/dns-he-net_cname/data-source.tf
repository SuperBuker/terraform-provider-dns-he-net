# Retrieve CNAME record.
resource "dns-he-net_cname" "example" {
  id      = 123456789
  zone_id = 123456
}
