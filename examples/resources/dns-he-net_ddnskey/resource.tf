# DDNS update key.
resource "dns-he-net_ddnskey" "example" {
  zone_id = 123456
  domain  = "example.com"
  key     = "secret key"
}

