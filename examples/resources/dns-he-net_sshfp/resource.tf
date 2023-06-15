# SSHFP record.
resource "dns-he-net_sshfp" "example" {
  zone_id = 123456
  domain  = "sshfp.example.com"
  ttl     = 86400
  data    = "4 2 123456789abcdef67890123456789abcdef67890123456789abcdef123456789"
}

