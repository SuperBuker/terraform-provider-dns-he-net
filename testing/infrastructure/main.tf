provider "dns-he-net" {
  username   = var.account.username
  password   = var.account.password
  otp_secret = var.account.otp_secret
  store_type = var.account.store_type
}


module "dns-he-net-testing" {
  source = "./modules/dns-he-net-testing"

  account = var.account
  datasources_zone = var.datasources_zone
  resources_zone = var.resources_zone
  config_file = var.config_file

  providers = {
    dns-he-net = dns-he-net
  }
}
