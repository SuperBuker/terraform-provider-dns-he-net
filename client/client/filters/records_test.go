package filters_test

import (
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/client/filters"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"

	"github.com/stretchr/testify/assert"
)

var nums = []uint{
	0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
}

var records = []models.Record{
	{ID: &nums[0], ZoneID: 1, Domain: "a.example.com", RecordType: "A", TTL: 300, Priority: nil, Data: "0.0.0.0", Dynamic: false, Locked: false},
	{ID: &nums[1], ZoneID: 1, Domain: "b.example.com", RecordType: "A", TTL: 300, Priority: nil, Data: "0.0.0.0", Dynamic: false, Locked: false},
	{ID: &nums[2], ZoneID: 1, Domain: "c.example.com", RecordType: "A", TTL: 300, Priority: nil, Data: "0.0.0.0", Dynamic: false, Locked: false},
	{ID: &nums[3], ZoneID: 1, Domain: "d.example.com", RecordType: "A", TTL: 300, Priority: nil, Data: "0.0.0.0", Dynamic: false, Locked: false},
	{ID: &nums[4], ZoneID: 1, Domain: "e.example.com", RecordType: "A", TTL: 300, Priority: nil, Data: "0.0.0.0", Dynamic: false, Locked: false},
	{ID: &nums[5], ZoneID: 1, Domain: "aaaa.example.com", RecordType: "AAAA", TTL: 300, Priority: nil, Data: "0.0.0.0", Dynamic: false, Locked: false},
	{ID: &nums[6], ZoneID: 1, Domain: "bbbb.example.com", RecordType: "AAAA", TTL: 300, Priority: nil, Data: "0.0.0.0", Dynamic: false, Locked: false},
	{ID: &nums[7], ZoneID: 1, Domain: "cccc.example.com", RecordType: "AAAA", TTL: 300, Priority: nil, Data: "0.0.0.0", Dynamic: false, Locked: false},
	{ID: &nums[8], ZoneID: 1, Domain: "dddd.example.com", RecordType: "AAAA", TTL: 300, Priority: nil, Data: "0.0.0.0", Dynamic: false, Locked: false},
	{ID: &nums[9], ZoneID: 1, Domain: "eeee.example.com", RecordType: "AAAA", TTL: 300, Priority: nil, Data: "0.0.0.0", Dynamic: false, Locked: false},
}

var recordsNil = []models.Record{
	{ID: nil, ZoneID: 1, Domain: "a.example.com", RecordType: "A", TTL: 300, Priority: nil, Data: "0.0.0.0", Dynamic: false, Locked: false},
	{ID: nil, ZoneID: 1, Domain: "aaaa.example.com", RecordType: "AAAA", TTL: 300, Priority: nil, Data: "0.0.0.0", Dynamic: false, Locked: false},
}

func TestRecordById(t *testing.T) {
	for i := uint(0); i < 10; i++ {
		record, ok := filters.RecordById(records, i)
		assert.Equal(t, records[i], record)
		assert.True(t, ok)
	}

	record, ok := filters.RecordById(records, 10)
	assert.Equal(t, models.Record{}, record)
	assert.False(t, ok)

	for i := uint(0); i < 10; i++ {
		record, ok := filters.RecordById(recordsNil, i)
		assert.Equal(t, models.Record{}, record)
		assert.False(t, ok)
	}
}

func TestMatchRecord(t *testing.T) {
	for i, record := range records {
		rec, ok := filters.MatchRecord(records, record)
		assert.True(t, ok)
		assert.Equal(t, records[i], rec)
	}
}

func TestLatestRecord(t *testing.T) {
	record, ok := filters.LatestRecord(records)
	assert.Equal(t, records[9], record)
	assert.True(t, ok)

	record, ok = filters.LatestRecord(recordsNil)
	assert.Equal(t, models.Record{}, record)
	assert.False(t, ok)
}

func TestRecord(t *testing.T) {
	records_ := filters.Record(records, nil, nil)
	assert.Equal(t, records, records_)
	assert.False(t, &records == &records_)

	domain := "d.example.com"
	records_ = filters.Record(records, &domain, nil)
	assert.Equal(t, []models.Record{records[3]}, records_)
	assert.False(t, &records == &records_)

	typ := "AAAA"
	records_ = filters.Record(records, &domain, &typ)
	assert.Equal(t, []models.Record{}, records_)
	assert.False(t, &records == &records_)

	records_ = filters.Record(records, nil, &typ)
	assert.Equal(t, []models.Record{records[5], records[6], records[7], records[8], records[9]}, records_)
	assert.False(t, &records == &records_)
}
