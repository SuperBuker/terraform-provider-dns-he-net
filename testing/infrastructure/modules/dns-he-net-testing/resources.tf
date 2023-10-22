# This file defines all the resources that are going to be created
# in the test environment.

resource "dns-he-net_a" "record-a" {
  zone_id = local.datasources_zone.id
  domain  = "example-a.${local.datasources_zone.name}"
  ttl     = 300
  dynamic = true
}

resource "dns-he-net_aaaa" "record-aaaa" {
  zone_id = local.datasources_zone.id
  domain  = "example-aaaa.${local.datasources_zone.name}"
  ttl     = 300
  dynamic = true
}

resource "dns-he-net_afsdb" "record-afsdb" {
  zone_id = local.datasources_zone.id
  domain  = local.datasources_zone.name
  ttl     = 300
  data    = "2 green.${local.datasources_zone.name}"
}

resource "dns-he-net_alias" "record-alias" {
  zone_id = local.datasources_zone.id
  domain  = "example-alias.${local.datasources_zone.name}"
  ttl     = 300
  data    = local.datasources_zone.name
}

resource "dns-he-net_caa" "record-caa" {
  zone_id = local.datasources_zone.id
  domain  = local.datasources_zone.name
  ttl     = 86400
  data    = "0 issuewild \";\""
}

resource "dns-he-net_cname" "record-cname" {
  zone_id = local.datasources_zone.id
  domain  = "example-cname.${local.datasources_zone.name}"
  ttl     = 300
  data    = "example.com"
}

resource "dns-he-net_hinfo" "record-hinfo" {
  zone_id = local.datasources_zone.id
  domain  = "example-hinfo.${local.datasources_zone.name}"
  ttl     = 86400
  data    = "\"armv7 Linux\""
}

resource "dns-he-net_loc" "record-loc" {
  zone_id = local.datasources_zone.id
  domain  = "example-loc.${local.datasources_zone.name}"
  ttl     = 86400
  data    = "40 27 53.86104 N 3 39 2.59092 W 712.8m 0.00m 0.00m 0.00m"
}

resource "dns-he-net_mx" "record-mx" {
  zone_id  = local.datasources_zone.id
  domain   = local.datasources_zone.name
  ttl      = 3600
  priority = 1
  data     = "mx.example.com"
}

resource "dns-he-net_naptr" "record-naptr" {
  zone_id = local.datasources_zone.id
  domain  = local.datasources_zone.name
  ttl     = 86400
  data    = "100 10 \"S\" \"SIP+D2U\" \"!^.*$!sip:bofher@${local.datasources_zone.name}!\" _sip._udp.${local.datasources_zone.name}."
}

#resource "dns-he-net_ns" "record-ns" {} # It's a datasource

resource "dns-he-net_ptr" "record-ptr" {
  zone_id = local.datasources_zone.id
  domain  = "example-ptr.${local.datasources_zone.name}"
  ttl     = 300
  data    = local.datasources_zone.name
}

resource "dns-he-net_rp" "record-rp" {
  zone_id = local.datasources_zone.id
  domain  = local.datasources_zone.name
  ttl     = 86400
  data    = "bofher.${local.datasources_zone.name} bofher.${local.datasources_zone.name}"
}

#resource "dns-he-net_soa" "record-soa" {} # It's a datasource

resource "dns-he-net_spf" "record-spf" {
  zone_id = local.datasources_zone.id
  domain  = local.datasources_zone.name
  ttl     = 86400
  data    = "\"v=spf1 include:_spf.example.com ~all\""
}

resource "dns-he-net_srv" "record-srv" {
  zone_id  = local.datasources_zone.id
  domain   = "_bofher._tcp.${local.datasources_zone.name}"
  ttl      = 28800
  priority = 0
  weight   = 0
  port     = 22
  target   = local.datasources_zone.name
}

resource "dns-he-net_sshfp" "record-sshfp" {
  zone_id = local.datasources_zone.id
  domain  = "example-sshfp.${local.datasources_zone.name}"
  ttl     = 86400
  data    = "4 2 123456789abcdef67890123456789abcdef67890123456789abcdef123456789"
}

resource "dns-he-net_txt" "record-txt" {
  zone_id = local.datasources_zone.id
  domain  = "bofher.${local.datasources_zone.name}"
  ttl     = 300
  data    = "\"Just for the record\""
  dynamic = true
}
