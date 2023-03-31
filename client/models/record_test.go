package models_test

import (
	"fmt"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var records_serial = map[string]int{
	"SOA":   5,
	"NS":    6,
	"A":     7,
	"AAAA":  7,
	"MX":    7,
	"CNAME": 6,
	"ALIAS": 6,
	"CAA":   6,
	"SRV":   9,
	"TXT":   7,
	"AFSDB": 7,
	"HINFO": 6,
	"RP":    6,
	"LOC":   6,
	"NAPTR": 6,
	"PTR":   6,
	"SSHFP": 6,
	"SPF":   6,
}

var records_in = []models.Record{
	{Id: nil, ParentId: 1234567, Domain: "example.com", RecordType: "SOA", TTL: 172800, Priority: nil, Data: "ns1.he.net. hostmaster.he.net. 2023031805 10800 1800 604800 86400", Dynamic: false, Locked: true},
	{Id: nil, ParentId: 1234567, Domain: "example.com", RecordType: "NS", TTL: 86400, Priority: nil, Data: "ns2.he.net", Dynamic: false, Locked: false},
	{Id: nil, ParentId: 1234567, Domain: "a.example.com", RecordType: "A", TTL: 300, Priority: nil, Data: "1.2.3.4", Dynamic: false, Locked: false},
	{Id: nil, ParentId: 1234567, Domain: "a.example.com", RecordType: "AAAA", TTL: 86400, Priority: nil, Data: "2001:1234:5678::1", Dynamic: false, Locked: false},
	{Id: nil, ParentId: 1234567, Domain: "example.com", RecordType: "MX", TTL: 3600, Priority: getUint16(1), Data: "mx.email.com", Dynamic: false, Locked: false},
	{Id: nil, ParentId: 1234567, Domain: "aa.example.com", RecordType: "CNAME", TTL: 86400, Priority: nil, Data: "example.com", Dynamic: false, Locked: false},
	{Id: nil, ParentId: 1234567, Domain: "aa.example.com", RecordType: "ALIAS", TTL: 86400, Priority: nil, Data: "example.com", Dynamic: false, Locked: false},
	{Id: nil, ParentId: 1234567, Domain: "example.com", RecordType: "CAA", TTL: 86400, Priority: nil, Data: "0 iodef \"webmaster@example.com\"", Dynamic: false, Locked: false},
	{Id: nil, ParentId: 1234567, Domain: "_https._tcp.example.com", RecordType: "SRV", TTL: 300, Priority: getUint16(0), Data: "0 443 ff.example.com", Dynamic: false, Locked: false},
	{Id: nil, ParentId: 1234567, Domain: "txt.example.com", RecordType: "TXT", TTL: 86400, Priority: nil, Data: "\"Some data 0\"", Dynamic: false, Locked: false},
	{Id: nil, ParentId: 1234567, Domain: "example.com", RecordType: "AFSDB", TTL: 300, Priority: nil, Data: "1 afsdb.example.com", Dynamic: false, Locked: false},
	{Id: nil, ParentId: 1234567, Domain: "example.com", RecordType: "HINFO", TTL: 300, Priority: nil, Data: "i686 Linux", Dynamic: false, Locked: false},
	{Id: nil, ParentId: 1234567, Domain: "example.com", RecordType: "RP", TTL: 300, Priority: nil, Data: "user.example.com user.example.com", Dynamic: false, Locked: false},
	{Id: nil, ParentId: 1234567, Domain: "example.com", RecordType: "LOC", TTL: 300, Priority: nil, Data: "51 56 0.123 N 5 54 0.000 E 4.00m 1.00m 10000.00m 10.00m", Dynamic: false, Locked: false},
	{Id: nil, ParentId: 1234567, Domain: "example.com", RecordType: "NAPTR", TTL: 300, Priority: nil, Data: `100 10 "U" "E2U+sip" "!^.*$!sip:customer-service@example.com!"`, Dynamic: false, Locked: false},
	{Id: nil, ParentId: 1234567, Domain: "example.com", RecordType: "PTR", TTL: 300, Priority: nil, Data: "example.com", Dynamic: false, Locked: false},
	{Id: nil, ParentId: 1234567, Domain: "example.com", RecordType: "SSHFP", TTL: 300, Priority: nil, Data: "4 2 123456789abcdef67890123456789abcdef67890123456789abcdef123456789", Dynamic: false, Locked: false},
	{Id: nil, ParentId: 1234567, Domain: "example.com", RecordType: "SPF", TTL: 300, Priority: nil, Data: `"v=spf1 include:spf.email.com -all"`, Dynamic: false, Locked: false},
}

var records_out = []models.RecordX{
	models.SOA{Id: nil, ParentId: 1234567, Domain: "example.com", TTL: 172800, MName: "ns1.he.net.", RName: "hostmaster.he.net.", Serial: 2023031805, Refresh: 10800, Retry: 1800, Expire: 604800},
	models.NS{Id: nil, ParentId: 1234567, Domain: "example.com", TTL: 86400, Data: "ns2.he.net"},
	models.A{Id: nil, ParentId: 1234567, Domain: "a.example.com", TTL: 300, Data: "1.2.3.4", Dynamic: false},
	models.AAAA{Id: nil, ParentId: 1234567, Domain: "a.example.com", TTL: 86400, Data: "2001:1234:5678::1", Dynamic: false},
	models.MX{Id: nil, ParentId: 1234567, Domain: "example.com", TTL: 3600, Priority: 1, Data: "mx.email.com"},
	models.CNAME{Id: nil, ParentId: 1234567, Domain: "aa.example.com", TTL: 86400, Data: "example.com"},
	models.ALIAS{Id: nil, ParentId: 1234567, Domain: "aa.example.com", TTL: 86400, Data: "example.com"},
	models.CAA{Id: nil, ParentId: 1234567, Domain: "example.com", TTL: 86400, Data: "0 iodef \"webmaster@example.com\""},
	models.SRV{Id: nil, ParentId: 1234567, Domain: "_https._tcp.example.com", TTL: 300, Priority: 0, Weight: 0, Port: 443, Target: "ff.example.com"},
	models.TXT{Id: nil, ParentId: 1234567, Domain: "txt.example.com", TTL: 86400, Data: "\"Some data 0\"", Dynamic: false},
	models.AFSDB{Id: nil, ParentId: 1234567, Domain: "example.com", TTL: 300, Data: "1 afsdb.example.com", Dynamic: false},
	models.HINFO{Id: nil, ParentId: 1234567, Domain: "example.com", TTL: 300, Data: "i686 Linux"},
	models.RP{Id: nil, ParentId: 1234567, Domain: "example.com", TTL: 300, Data: "user.example.com user.example.com"},
	models.LOC{Id: nil, ParentId: 1234567, Domain: "example.com", TTL: 300, Data: "51 56 0.123 N 5 54 0.000 E 4.00m 1.00m 10000.00m 10.00m"},
	models.NAPTR{Id: nil, ParentId: 1234567, Domain: "example.com", TTL: 300, Data: `100 10 "U" "E2U+sip" "!^.*$!sip:customer-service@example.com!"`},
	models.PTR{Id: nil, ParentId: 1234567, Domain: "example.com", TTL: 300, Data: "example.com"},
	models.SSHFP{Id: nil, ParentId: 1234567, Domain: "example.com", TTL: 300, Data: "4 2 123456789abcdef67890123456789abcdef67890123456789abcdef123456789"},
	models.SPF{Id: nil, ParentId: 1234567, Domain: "example.com", TTL: 300, Data: `"v=spf1 include:spf.email.com -all"`},
}

func getUint16(i uint16) *uint16 {
	return &i
}

func TestRecord(t *testing.T) {
	for i, record_in := range records_in {
		record_out, err := record_in.ToX()
		require.NoError(t, err, record_in.RecordType)
		assert.Equal(t, records_out[i], record_out, record_in.RecordType)

		id, ok := record_out.GetId()
		assert.False(t, ok, record_in.RecordType)
		assert.Equal(t, uint(0), id, record_in.RecordType)

		assert.Equal(t, record_in.ParentId, record_out.GetParentId(), record_in.RecordType)

		assert.Equal(t, record_in.RecordType, record_out.Type(), record_in.RecordType)

		assert.Equal(t, map[string]string{
			"hosted_dns_zoneid":   fmt.Sprint(record_in.ParentId),
			"hosted_dns_recordid": "",
		}, record_out.Refs(), record_in.RecordType)

		assert.Equal(t, record_out.Refs(), record_in.Refs(), record_in.RecordType)

		assert.Equal(t, records_serial[record_in.RecordType], len(record_out.Serialise()), record_in.RecordType)

		assert.Equal(t, record_out.Serialise(), record_in.Serialise(), record_in.RecordType)

	}
}
