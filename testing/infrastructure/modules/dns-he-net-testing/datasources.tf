# This file defines all the datasources existing in the test environment.
# Some DNS records should not be manipulated.

data "dns-he-net_arpa_zones" "arpas" {
}

data "dns-he-net_domain_zones" "domains" {
}

data "dns-he-net_network_prefixes" "network_prefixes" {
}

data "dns-he-net_records" "arpa_records" {
  id = local.datasources_arpa_zone.zone_id
}

data "dns-he-net_records" "domain_records" {
  id = local.datasources_domain_zone.zone_id
}

data "dns-he-net_soa" "arpa_soa" {
  id = local.arpa_SOA_lite.id
  zone_id = local.arpa_SOA_lite.zone_id
}

data "dns-he-net_soa" "domain_soa" {
  id = local.domain_SOA_lite.id
  zone_id = local.domain_SOA_lite.zone_id
}
