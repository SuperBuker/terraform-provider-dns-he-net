# Static TXT record.
resource "dns-he-net_txt" "example" {
  zone_id = 123456
  domain  = "example.com"
  ttl     = 86400
  data    = "\"Just for the record\""
}

# Dynamic TXT record.
resource "dns-he-net_txt" "example" {
  zone_id = 123456
  domain  = "example.com"
  ttl     = 86400
  dynamic = true
}

# Dynamic TXT record with preset value.
resource "dns-he-net_txt" "example" {
  zone_id = 123456
  domain  = "example.com"
  ttl     = 86400
  data    = "\"Just for the record\""
  dynamic = true
}
