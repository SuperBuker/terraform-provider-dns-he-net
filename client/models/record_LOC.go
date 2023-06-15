package models

import "fmt"

type LOC struct {
	ID     *uint
	ZoneID uint
	Domain string
	TTL    uint // seconds
	Data   string
}

func ToLOC(r Record) LOC {
	return LOC{
		ID:     r.ID,
		ZoneID: r.ZoneID,
		Domain: r.Domain,
		TTL:    r.TTL,
		Data:   r.Data,
	}
}

func (r LOC) Equals(rx RecordX) bool {
	if rx == nil {
		return false
	} else if rx.Type() != "LOC" {
		return false
	} else if rec, ok := rx.(Record); ok {
		// Convert from Record
		rx = ToLOC(rec)
	}

	rloc := rx.(LOC)

	return r.ZoneID == rloc.ZoneID &&
		r.Domain == rloc.Domain &&
		r.TTL == rloc.TTL &&
		r.Data == rloc.Data
}

func (r LOC) Serialise() map[string]string {
	return map[string]string{
		"Type":                "LOC",
		"hosted_dns_zoneid":   fmt.Sprint(r.ZoneID),
		"hosted_dns_recordid": toString(r.ID),
		//"Priority": "",
		"Name":    r.Domain,
		"Content": r.Data,
		"TTL":     fmt.Sprint(r.TTL),
	}
}

func (r LOC) Refs() map[string]string {
	return map[string]string{
		"hosted_dns_zoneid":   fmt.Sprint(r.ZoneID),
		"hosted_dns_recordid": toString(r.ID),
	}
}

func (r LOC) GetID() (uint, bool) {
	if r.ID == nil {
		return 0, false
	}

	return *r.ID, true
}

func (r LOC) GetZoneID() uint {
	return r.ZoneID
}

func (r LOC) Type() string {
	return "LOC"
}
