provider "dns-he-net" {
  username   = var.dhn_username
  password   = var.dhn_password
  otp_secret = var.dhn_2fa
  store_type = var.dhn_store_type
}

data "dns-he-net_zones" "zones" {}

resource "dns-he-net_a" "record-a" {
  zone_id = data.dns-he-net_zones.zones.zones[0].id
  domain  = "server0.example.com"
  ttl     = 300
  dynamic = true
}

resource "dns-he-net_cname" "example" {
  zone_id = data.dns-he-net_zones.zones.zones[0].id
  domain  = "example.com"
  ttl     = 86400
  data    = resource.dns-he-net_a.record-a.domain
}