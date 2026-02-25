# DDNS update key.
resource "dns-he-net_ddnskey" "example" {
  zone_id = 123456
  domain  = "example.com"
  key     = "secret key"
}

// IMPORTANT: This resource is not deletable. It can only be created and updated.
// On deletion, a random key is set. Please ensure the DDNS flag is disabled in the Record.