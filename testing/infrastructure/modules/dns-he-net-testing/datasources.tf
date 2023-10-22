# This file defines all the datasources existing in the test environment.
# Some DNS records should not be manipulated.

data "dns-he-net_zones" "zones" {
}


data "dns-he-net_records" "records" {
  id = local.datasources_zone.id
}
