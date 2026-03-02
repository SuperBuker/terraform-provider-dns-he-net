# Refactoring variables and datasources/resources

locals {
  # Account config
  account = {
    user     = var.account.username
    password = var.account.password
    otp      = var.account.otp_secret
    id       = data.dns-he-net_domain_zones.domains.id
  }

  # Get Datasources zones
  datasources_domain_zone = data.dns-he-net_domain_zones.domains.zones[index(data.dns-he-net_domain_zones.domains.zones.*.name, var.datasources_domain_zone)]
  datasources_domain_zone_pending_delegation = data.dns-he-net_domain_zones.domains.zones[index(data.dns-he-net_domain_zones.domains.zones.*.name, var.datasources_domain_zone_pending_delegation)]
  datasources_arpa_zone   = data.dns-he-net_arpa_zones.arpas.zones[index(data.dns-he-net_arpa_zones.arpas.zones.*.name, var.datasources_arpa_zone)]

  # Get Resources zone
  resources_domain_zone = try(
    data.dns-he-net_domain_zones.domains.zones[index(data.dns-he-net_domain_zones.domains.zones.*.name, var.resources_domain_zone)],
    {
      zone_id = ""
      name    = ""
    },
  )

  resources_arpa_zone = try(
    data.dns-he-net_arpa_zones.arpas.zones[index(data.dns-he-net_arpa_zones.arpas.zones.*.name, var.resources_arpa_zone)],
    {
      zone_id = ""
      name    = ""
    },
  )

  datasources_network_prefix = try(
    data.dns-he-net_network_prefixes.network_prefixes.network_prefixes[index(data.dns-he-net_network_prefixes.network_prefixes.network_prefixes.*.value, var.datasources_network_prefix)],
    {
      id    = ""
      value = ""
    },
  )

  # Static records
  domain_SOA_lite = data.dns-he-net_records.domain_records.records[index(data.dns-he-net_records.domain_records.records.*.record_type, "SOA")]
  arpa_SOA_lite   = data.dns-he-net_records.arpa_records.records[index(data.dns-he-net_records.arpa_records.records.*.record_type, "SOA")]
  domain_NS       = data.dns-he-net_records.domain_records.records[index(data.dns-he-net_records.domain_records.records.*.data, "ns1.he.net")]
  arpa_NS         = data.dns-he-net_records.arpa_records.records[index(data.dns-he-net_records.arpa_records.records.*.data, "ns1.he.net")]

  # Get Records config
  records = {
    domain = {
      A = {
        id     = resource.dns-he-net_a.record-a.id
        domain = resource.dns-he-net_a.record-a.domain
        ttl    = resource.dns-he-net_a.record-a.ttl
        data   = resource.dns-he-net_a.record-a.data
      },
      AAAA = {
        id     = resource.dns-he-net_aaaa.record-aaaa.id
        domain = resource.dns-he-net_aaaa.record-aaaa.domain
        ttl    = resource.dns-he-net_aaaa.record-aaaa.ttl
        data   = resource.dns-he-net_aaaa.record-aaaa.data
      },
      AFSDB = {
        id     = resource.dns-he-net_afsdb.record-afsdb.id
        domain = resource.dns-he-net_afsdb.record-afsdb.domain
        ttl    = resource.dns-he-net_afsdb.record-afsdb.ttl
        data   = resource.dns-he-net_afsdb.record-afsdb.data
      },
      ALIAS = {
        id     = resource.dns-he-net_alias.record-alias.id
        domain = resource.dns-he-net_alias.record-alias.domain
        ttl    = resource.dns-he-net_alias.record-alias.ttl
        data   = resource.dns-he-net_alias.record-alias.data
      },
      CAA = {
        id     = resource.dns-he-net_caa.record-caa.id
        domain = resource.dns-he-net_caa.record-caa.domain
        ttl    = resource.dns-he-net_caa.record-caa.ttl
        data   = resource.dns-he-net_caa.record-caa.data
      },
      CNAME = {
        id     = resource.dns-he-net_cname.record-cname.id
        domain = resource.dns-he-net_cname.record-cname.domain
        ttl    = resource.dns-he-net_cname.record-cname.ttl
        data   = resource.dns-he-net_cname.record-cname.data
      },
      HINFO = {
        id     = resource.dns-he-net_hinfo.record-hinfo.id
        domain = resource.dns-he-net_hinfo.record-hinfo.domain
        ttl    = resource.dns-he-net_hinfo.record-hinfo.ttl
        data   = resource.dns-he-net_hinfo.record-hinfo.data
      },
      LOC = {
        id     = resource.dns-he-net_loc.record-loc.id
        domain = resource.dns-he-net_loc.record-loc.domain
        ttl    = resource.dns-he-net_loc.record-loc.ttl
        data   = resource.dns-he-net_loc.record-loc.data
      },
      MX = {
        id     = resource.dns-he-net_mx.record-mx.id
        domain = resource.dns-he-net_mx.record-mx.domain
        ttl    = resource.dns-he-net_mx.record-mx.ttl
        data   = resource.dns-he-net_mx.record-mx.data
      },
      NS = {
        id     = local.domain_NS.id
        domain = local.domain_NS.domain
        ttl    = local.domain_NS.ttl
        data   = local.domain_NS.data
      },
      PTR = {
        id     = resource.dns-he-net_ptr.record-ptr.id
        domain = resource.dns-he-net_ptr.record-ptr.domain
        ttl    = resource.dns-he-net_ptr.record-ptr.ttl
        data   = resource.dns-he-net_ptr.record-ptr.data
      },
      RP = {
        id     = resource.dns-he-net_rp.record-rp.id
        domain = resource.dns-he-net_rp.record-rp.domain
        ttl    = resource.dns-he-net_rp.record-rp.ttl
        data   = resource.dns-he-net_rp.record-rp.data
      },
      SOA = {
        id     = data.dns-he-net_soa.domain_soa.id
        domain = data.dns-he-net_soa.domain_soa.domain
        ttl    = data.dns-he-net_soa.domain_soa.ttl
        data   = local.domain_SOA_lite.data
        extra_args = {
          mname   = data.dns-he-net_soa.domain_soa.mname
          rname   = data.dns-he-net_soa.domain_soa.rname
          refresh = tostring(data.dns-he-net_soa.domain_soa.refresh)
          retry   = tostring(data.dns-he-net_soa.domain_soa.retry)
          expire  = tostring(data.dns-he-net_soa.domain_soa.expire)
        }
      },
      SRV = {
        id     = resource.dns-he-net_srv.record-srv.id
        domain = resource.dns-he-net_srv.record-srv.domain
        ttl    = resource.dns-he-net_srv.record-srv.ttl
        data   = resource.dns-he-net_srv.record-srv.target
      },
      SSHFP = {
        id     = resource.dns-he-net_sshfp.record-sshfp.id
        domain = resource.dns-he-net_sshfp.record-sshfp.domain
        ttl    = resource.dns-he-net_sshfp.record-sshfp.ttl
        data   = resource.dns-he-net_sshfp.record-sshfp.data
      },
      TXT = {
        id     = resource.dns-he-net_txt.record-txt.id
        domain = resource.dns-he-net_txt.record-txt.domain
        ttl    = resource.dns-he-net_txt.record-txt.ttl
        data   = resource.dns-he-net_txt.record-txt.data
      }
    },
    arpa = {
      CNAME = {
        id     = resource.dns-he-net_cname.record-arpa-cname.id
        domain = resource.dns-he-net_cname.record-arpa-cname.domain
        ttl    = resource.dns-he-net_cname.record-arpa-cname.ttl
        data   = resource.dns-he-net_cname.record-arpa-cname.data
      },
      PTR = {
        id     = resource.dns-he-net_ptr.record-arpa-ptr.id
        domain = resource.dns-he-net_ptr.record-arpa-ptr.domain
        ttl    = resource.dns-he-net_ptr.record-arpa-ptr.ttl
        data   = resource.dns-he-net_ptr.record-arpa-ptr.data
      },
      SOA = {
        id     = data.dns-he-net_soa.arpa_soa.id
        domain = data.dns-he-net_soa.arpa_soa.domain
        ttl    = data.dns-he-net_soa.arpa_soa.ttl
        data   = local.arpa_SOA_lite.data
        extra_args = {
          mname   = data.dns-he-net_soa.arpa_soa.mname
          rname   = data.dns-he-net_soa.arpa_soa.rname
          refresh = tostring(data.dns-he-net_soa.arpa_soa.refresh)
          retry   = tostring(data.dns-he-net_soa.arpa_soa.retry)
          expire  = tostring(data.dns-he-net_soa.arpa_soa.expire)
        }
      },
      NS = {
        id     = local.arpa_NS.id
        domain = local.arpa_NS.domain
        ttl    = local.arpa_NS.ttl
        data   = local.arpa_NS.data

      },
      TXT = {
        id     = resource.dns-he-net_txt.record-arpa-txt.id
        domain = resource.dns-he-net_txt.record-arpa-txt.domain
        ttl    = resource.dns-he-net_txt.record-arpa-txt.ttl
        data   = resource.dns-he-net_txt.record-arpa-txt.data
      }
    }
  }

  test_config = {
    datasources = {
      account = var.account.mask_creds ? null : local.account
      domain_zones = {
        ok = {
          id           = local.datasources_domain_zone.zone_id
          name         = local.datasources_domain_zone.name
          records      = local.records.domain
          record_count = length(data.dns-he-net_records.domain_records.records)

        },
        pending_delegation = {
          id           = local.datasources_domain_zone_pending_delegation.zone_id
          name         = local.datasources_domain_zone_pending_delegation.name
          record_count = 0

        }
      },
      domain_zones_count = length(data.dns-he-net_domain_zones.domains.zones)
      network_prefixes = {
        ok = {
          id      = local.datasources_network_prefix.id
          value   = local.datasources_network_prefix.value
          enabled = true
        },
      },
      network_prefixes_count = length(data.dns-he-net_network_prefixes.network_prefixes)
      arpa_zones = {
        ok = {
          id           = local.datasources_arpa_zone.zone_id
          name         = local.datasources_arpa_zone.name
          records      = local.records.arpa
          record_count = length(data.dns-he-net_records.arpa_records.records)

        },
      },
      arpa_zones_count = length(data.dns-he-net_arpa_zones.arpas.zones)
    }
    resources = {
      account     = var.account.mask_creds ? null : local.account
      domain_zone = local.resources_domain_zone
      arpa_zone   = local.resources_arpa_zone
    }
  }
}
