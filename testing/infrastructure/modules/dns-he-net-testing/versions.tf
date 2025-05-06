terraform {
  required_version = ">= 1.3.0"

  required_providers {
    dns-he-net = {
      source  = "SuperBuker/dns-he-net"
      version = "0.0.9"
    }
  }
}
