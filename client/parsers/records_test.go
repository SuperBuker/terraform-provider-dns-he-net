package parsers_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/parsers"
	"github.com/antchfx/htmlquery"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var records = []models.Record{
	{ID: nil, ZoneID: 1234567, Domain: "example.com", RecordType: "SOA", TTL: 172800, Priority: nil, Data: "ns1.he.net. hostmaster.he.net. 2023031805 10800 1800 604800 86400", Dynamic: false, Locked: true},
	{ID: nil, ZoneID: 1234567, Domain: "example.com", RecordType: "NS", TTL: 86400, Priority: nil, Data: "ns2.he.net", Dynamic: false, Locked: false},
	{ID: nil, ZoneID: 1234567, Domain: "example.com", RecordType: "NS", TTL: 86400, Priority: nil, Data: "ns3.he.net", Dynamic: false, Locked: false},
	{ID: nil, ZoneID: 1234567, Domain: "example.com", RecordType: "NS", TTL: 86400, Priority: nil, Data: "ns5.he.net", Dynamic: false, Locked: false},
	{ID: nil, ZoneID: 1234567, Domain: "example.com", RecordType: "NS", TTL: 86400, Priority: nil, Data: "ns4.he.net", Dynamic: false, Locked: false},
	{ID: nil, ZoneID: 1234567, Domain: "a.example.com", RecordType: "A", TTL: 300, Priority: nil, Data: "1.2.3.4", Dynamic: false, Locked: false},
	{ID: nil, ZoneID: 1234567, Domain: "b.example.com", RecordType: "A", TTL: 300, Priority: nil, Data: "1.2.3.4", Dynamic: true, Locked: false},
	{ID: nil, ZoneID: 1234567, Domain: "c.example.com", RecordType: "A", TTL: 300, Priority: nil, Data: "1.2.3.4", Dynamic: false, Locked: false},
	{ID: nil, ZoneID: 1234567, Domain: "d.example.com", RecordType: "A", TTL: 300, Priority: nil, Data: "1.2.3.4", Dynamic: true, Locked: false},
	{ID: nil, ZoneID: 1234567, Domain: "e.example.com", RecordType: "A", TTL: 300, Priority: nil, Data: "1.2.3.4", Dynamic: false, Locked: false},
	{ID: nil, ZoneID: 1234567, Domain: "f.example.com", RecordType: "A", TTL: 300, Priority: nil, Data: "1.2.3.4", Dynamic: true, Locked: false},
	{ID: nil, ZoneID: 1234567, Domain: "g.example.com", RecordType: "A", TTL: 300, Priority: nil, Data: "1.2.3.4", Dynamic: true, Locked: false},
	{ID: nil, ZoneID: 1234567, Domain: "h.example.com", RecordType: "A", TTL: 300, Priority: nil, Data: "1.2.3.4", Dynamic: false, Locked: false},
	{ID: nil, ZoneID: 1234567, Domain: "a.example.com", RecordType: "AAAA", TTL: 86400, Priority: nil, Data: "2001:1234:5678::1", Dynamic: false, Locked: false},
	{ID: nil, ZoneID: 1234567, Domain: "b.example.com", RecordType: "AAAA", TTL: 86400, Priority: nil, Data: "2001:1234:5678::1", Dynamic: false, Locked: false},
	{ID: nil, ZoneID: 1234567, Domain: "c.example.com", RecordType: "AAAA", TTL: 86400, Priority: nil, Data: "2001:1234:5678::1", Dynamic: false, Locked: false},
	{ID: nil, ZoneID: 1234567, Domain: "d.example.com", RecordType: "AAAA", TTL: 86400, Priority: nil, Data: "2001:1234:5678::1", Dynamic: false, Locked: false},
	{ID: nil, ZoneID: 1234567, Domain: "e.example.com", RecordType: "AAAA", TTL: 86400, Priority: nil, Data: "2001:1234:5678::1", Dynamic: false, Locked: false},
	{ID: nil, ZoneID: 1234567, Domain: "f.example.com", RecordType: "AAAA", TTL: 86400, Priority: nil, Data: "2001:1234:5678::1", Dynamic: false, Locked: false},
	{ID: nil, ZoneID: 1234567, Domain: "example.com", RecordType: "MX", TTL: 3600, Priority: getUint16(1), Data: "mx.email.com", Dynamic: false, Locked: false},
	{ID: nil, ZoneID: 1234567, Domain: "example.com", RecordType: "MX", TTL: 3600, Priority: getUint16(5), Data: "alt1.mx.email.com", Dynamic: false, Locked: false},
	{ID: nil, ZoneID: 1234567, Domain: "example.com", RecordType: "MX", TTL: 3600, Priority: getUint16(5), Data: "alt2.mx.email.com", Dynamic: false, Locked: false},
	{ID: nil, ZoneID: 1234567, Domain: "example.com", RecordType: "MX", TTL: 3600, Priority: getUint16(10), Data: "alt3.mx.email.com", Dynamic: false, Locked: false},
	{ID: nil, ZoneID: 1234567, Domain: "example.com", RecordType: "MX", TTL: 3600, Priority: getUint16(10), Data: "alt4.mx.email.com", Dynamic: false, Locked: false},
	{ID: nil, ZoneID: 1234567, Domain: "aa.example.com", RecordType: "CNAME", TTL: 86400, Priority: nil, Data: "example.com", Dynamic: false, Locked: false},
	{ID: nil, ZoneID: 1234567, Domain: "bb.example.com", RecordType: "CNAME", TTL: 86400, Priority: nil, Data: "example.com", Dynamic: false, Locked: false},
	{ID: nil, ZoneID: 1234567, Domain: "cc.example.com", RecordType: "CNAME", TTL: 86400, Priority: nil, Data: "example.com", Dynamic: false, Locked: false},
	{ID: nil, ZoneID: 1234567, Domain: "dd.example.com", RecordType: "CNAME", TTL: 86400, Priority: nil, Data: "example.com", Dynamic: false, Locked: false},
	{ID: nil, ZoneID: 1234567, Domain: "ee.example.com", RecordType: "CNAME", TTL: 86400, Priority: nil, Data: "e.service.com", Dynamic: false, Locked: false},
	{ID: nil, ZoneID: 1234567, Domain: "ff.example.com", RecordType: "CNAME", TTL: 86400, Priority: nil, Data: "example.com", Dynamic: false, Locked: false},
	{ID: nil, ZoneID: 1234567, Domain: "example.com", RecordType: "CAA", TTL: 86400, Priority: nil, Data: "0 iodef \"webmaster@example.com\"", Dynamic: false, Locked: false},
	{ID: nil, ZoneID: 1234567, Domain: "example.com", RecordType: "CAA", TTL: 86400, Priority: nil, Data: "0 issue \"letsencrypt.org\"", Dynamic: false, Locked: false},
	{ID: nil, ZoneID: 1234567, Domain: "example.com", RecordType: "CAA", TTL: 86400, Priority: nil, Data: "0 issuewild \";\"", Dynamic: false, Locked: false},
	{ID: nil, ZoneID: 1234567, Domain: "_https._tcp.example.com", RecordType: "SRV", TTL: 300, Priority: getUint16(0), Data: "0 443 ff.example.com", Dynamic: false, Locked: false},
	{ID: nil, ZoneID: 1234567, Domain: "txt.example.com", RecordType: "TXT", TTL: 86400, Priority: nil, Data: "\"Some data 0\"", Dynamic: false, Locked: false},
	{ID: nil, ZoneID: 1234567, Domain: "txt.example.com", RecordType: "TXT", TTL: 86400, Priority: nil, Data: "\"Some data 1\"", Dynamic: false, Locked: false},
	{ID: nil, ZoneID: 1234567, Domain: "txt.example.com", RecordType: "TXT", TTL: 86400, Priority: nil, Data: "\"Some data 2\"", Dynamic: false, Locked: false},
	{ID: nil, ZoneID: 1234567, Domain: "txt.example.com", RecordType: "TXT", TTL: 86400, Priority: nil, Data: "\"Some data 3\"", Dynamic: false, Locked: false},
	{ID: nil, ZoneID: 1234567, Domain: "txt.example.com", RecordType: "TXT", TTL: 86400, Priority: nil, Data: "\"Some data 4\"", Dynamic: false, Locked: false},
}

func TestRecords(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		data, err := os.ReadFile("../testing_data/html/records.html")
		require.NoError(t, err)

		doc, err := htmlquery.Parse(bytes.NewReader(data))
		require.NoError(t, err)

		_records, err := parsers.GetRecords(doc)
		require.NoError(t, err)

		for i, record := range _records {
			assert.Equal(t, records[i], record)
		}
	})

	t.Run("missing data", func(t *testing.T) {
		data := []byte("<html></html>")
		doc, err := htmlquery.Parse(bytes.NewReader(data))
		require.NoError(t, err)

		_records, err := parsers.GetRecords(doc)
		require.Error(t, err)
		targetErr := &parsers.ErrNotFound{}
		assert.ErrorAs(t, err, &targetErr)

		assert.Nil(t, _records)
		assert.Equal(t, errNotFoundString(recordsTableQ), err.Error())
	})

	t.Run("empty table", func(t *testing.T) {
		data, err := os.ReadFile("../testing_data/html/records_empty.html")
		require.NoError(t, err)

		doc, err := htmlquery.Parse(bytes.NewReader(data))
		require.NoError(t, err)

		_records, err := parsers.GetRecords(doc)
		require.NoError(t, err)

		assert.Equal(t, []models.Record{}, _records)
	})
}

func getUint16(i uint16) *uint16 {
	return &i
}

func init() {
	// Add if to the records.
	for i, x := range records {
		var id = uint(123456789 + i)
		x.ID = &id
		records[i] = x
	}
}
