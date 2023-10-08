# Refactoring variables and datasources/resources

locals {
  # Account config
  account = {
    user     = var.account.username
    password = var.account.password
    otp      = var.account.otp_secret
    id       = data.dns-he-net_zones.zones.id
  }

  # Get Datasources zone
  datasources_zone = data.dns-he-net_zones.zones.zones[index(data.dns-he-net_zones.zones.zones.*.name, var.datasources_zone)]

  # Get Resources zone
  resources_zone = try(
    data.dns-he-net_zones.zones.zones[index(data.dns-he-net_zones.zones.zones.*.name, var.resources_zone)],
    {
      id   = ""
      name = ""
    },
  )

  # Static records
  SOA = data.dns-he-net_records.records.records[index(data.dns-he-net_records.records.records.*.record_type, "SOA")]
  NS  = data.dns-he-net_records.records.records[index(data.dns-he-net_records.records.records.*.data, "ns1.he.net")]

  # Get Records config
  records = {
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
      id     = local.NS.id
      domain = local.NS.domain
      ttl    = local.NS.ttl
      data   = local.NS.data
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
      id     = local.SOA.id
      domain = local.SOA.domain
      ttl    = local.SOA.ttl
      data   = local.SOA.data
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
  }

  test_config = {
    datasources = {
      account     = var.account.mask_creds ? null : local.account
      zone        = local.datasources_zone
      zones_count = length(data.dns-he-net_zones.zones.zones)
      records     = local.records
    }
    resources = {
      account = var.account.mask_creds ? null : local.account
      zone    = local.resources_zone
    }
  }
}
