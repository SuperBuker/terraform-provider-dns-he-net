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
	{Id: &nums[0], ParentId: 1, Domain: "a.example.com", RecordType: "A", TTL: 300, Priority: nil, Data: "0.0.0.0", Dynamic: false, Locked: false},
	{Id: &nums[1], ParentId: 1, Domain: "b.example.com", RecordType: "A", TTL: 300, Priority: nil, Data: "0.0.0.0", Dynamic: false, Locked: false},
	{Id: &nums[2], ParentId: 1, Domain: "c.example.com", RecordType: "A", TTL: 300, Priority: nil, Data: "0.0.0.0", Dynamic: false, Locked: false},
	{Id: &nums[3], ParentId: 1, Domain: "d.example.com", RecordType: "A", TTL: 300, Priority: nil, Data: "0.0.0.0", Dynamic: false, Locked: false},
	{Id: &nums[4], ParentId: 1, Domain: "e.example.com", RecordType: "A", TTL: 300, Priority: nil, Data: "0.0.0.0", Dynamic: false, Locked: false},
	{Id: &nums[5], ParentId: 1, Domain: "aaaa.example.com", RecordType: "AAAA", TTL: 300, Priority: nil, Data: "0.0.0.0", Dynamic: false, Locked: false},
	{Id: &nums[6], ParentId: 1, Domain: "bbbb.example.com", RecordType: "AAAA", TTL: 300, Priority: nil, Data: "0.0.0.0", Dynamic: false, Locked: false},
	{Id: &nums[7], ParentId: 1, Domain: "cccc.example.com", RecordType: "AAAA", TTL: 300, Priority: nil, Data: "0.0.0.0", Dynamic: false, Locked: false},
	{Id: &nums[8], ParentId: 1, Domain: "dddd.example.com", RecordType: "AAAA", TTL: 300, Priority: nil, Data: "0.0.0.0", Dynamic: false, Locked: false},
	{Id: &nums[9], ParentId: 1, Domain: "eeee.example.com", RecordType: "AAAA", TTL: 300, Priority: nil, Data: "0.0.0.0", Dynamic: false, Locked: false},
}

var recordsNil = []models.Record{
	{Id: nil, ParentId: 1, Domain: "a.example.com", RecordType: "A", TTL: 300, Priority: nil, Data: "0.0.0.0", Dynamic: false, Locked: false},
	{Id: nil, ParentId: 1, Domain: "aaaa.example.com", RecordType: "AAAA", TTL: 300, Priority: nil, Data: "0.0.0.0", Dynamic: false, Locked: false},
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

func TestLatestRecord(t *testing.T) {
	record, ok := filters.LatestRecord(records)
	assert.Equal(t, records[9], record)
	assert.True(t, ok)

	record, ok = filters.LatestRecord(recordsNil)
	assert.Equal(t, models.Record{}, record)
	assert.False(t, ok)
}

func TestRecord(t *testing.T) {
	_records := filters.Record(records, nil, nil)
	assert.Equal(t, records, _records)
	assert.False(t, &records == &_records)

	domain := "d.example.com"
	_records = filters.Record(records, &domain, nil)
	assert.Equal(t, []models.Record{records[3]}, _records)
	assert.False(t, &records == &_records)

	typ := "AAAA"
	_records = filters.Record(records, &domain, &typ)
	assert.Equal(t, []models.Record{}, _records)
	assert.False(t, &records == &_records)

	_records = filters.Record(records, nil, &typ)
	assert.Equal(t, []models.Record{records[5], records[6], records[7], records[8], records[9]}, _records)
	assert.False(t, &records == &_records)
}
