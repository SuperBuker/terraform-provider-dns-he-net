# Required Provider
terraform {
  required_providers {
    dns-he-net = {
      source = "SuperBuker/dns-he-net"
    }
  }
}

# Provider Configuration
provider "dns-he-net" {
  username   = var.dhn_username
  password   = var.dhn_password
  otp_secret = var.dhn_2fa
  store_type = var.dhn_store_type
}
