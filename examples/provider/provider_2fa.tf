# Configure the DNS Provider
provider "dns-he-net" {
  username   = "john.smith@email.com"
  password   = "123456"
  otp_secret = "JBSWY3DPEHPK3PXP"
  store_type = "encrypted"
}


# Retrieve zones
data "dns-he-net_zones" "zones" {
}

# Create a DNS A record set
resource "dns-he-net_a" "record-a" {
  zone_id = data.dns-he-net_zones.zones.zones[0].id
  domain  = "example.com"
  ttl     = 86400
  data    = "1.2.3.4"
}