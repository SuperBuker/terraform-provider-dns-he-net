provider "dns-he-net" {
  username   = var.account.username
  password   = var.account.password
  otp_secret = var.account.otp_secret
  store_type = var.account.store_type
}


module "dns-he-net-testing" {
  source = "./modules/dns-he-net-testing"

  account = var.account
  datasources_domain_zone = var.datasources_domain_zone
  datasources_domain_zone_pending_delegation = var.datasources_domain_zone_pending_delegation
  datasources_arpa_zone = var.datasources_arpa_zone
  datasources_arpa_domain_example = var.datasources_arpa_domain_example
  datasources_network_prefix = var.datasources_network_prefix
  resources_domain_zone = var.resources_domain_zone
  resources_arpa_zone = var.resources_arpa_zone
  config_file = var.config_file

  providers = {
    dns-he-net = dns-he-net
  }
}
