# Retrieve PTR record.
resource "dns-he-net_ptr" "example" {
  id      = 123456789
  zone_id = 123456
}
