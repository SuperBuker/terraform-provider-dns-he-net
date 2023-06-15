# NAPTR record.
resource "dns-he-net_naptr" "example" {
  zone_id = 123456
  domain  = "example.com"
  ttl     = 86400
  data    = "100 10 \"S\" \"SIP+D2U\" \"!^.*$!sip:bofher@example.com!\" _sip._udp.example.com."
}

