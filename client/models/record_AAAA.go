package models

import "fmt"

type AAAA struct {
	ID      *uint
	ZoneID  uint
	Domain  string
	TTL     uint // seconds
	Data    string
	Dynamic bool
}

func ToAAAA(r Record) AAAA {
	return AAAA{
		ID:      r.ID,
		ZoneID:  r.ZoneID,
		Domain:  r.Domain,
		TTL:     r.TTL,
		Data:    r.Data,
		Dynamic: r.Dynamic,
	}
}

func (r AAAA) Equals(rx RecordX) bool {
	if rx == nil {
		return false
	} else if rx.Type() != "AAAA" {
		return false
	} else if rec, ok := rx.(Record); ok {
		// Convert from Record
		rx = ToAAAA(rec)
	}

	raaaa := rx.(AAAA)

	return r.ZoneID == raaaa.ZoneID &&
		r.Domain == raaaa.Domain &&
		r.TTL == raaaa.TTL &&
		r.Data == raaaa.Data &&
		r.Dynamic == raaaa.Dynamic
}

func (r AAAA) Serialise() map[string]string {
	return map[string]string{
		"Type":                "aaaa",
		"hosted_dns_zoneid":   fmt.Sprint(r.ZoneID),
		"hosted_dns_recordid": toString(r.ID),
		//"Priority":            "",
		"Name":    r.Domain,
		"Content": r.Data,
		"TTL":     fmt.Sprint(r.TTL),
		"dynamic": b2s[r.Dynamic],
	}
}

func (r AAAA) Refs() map[string]string {
	return map[string]string{
		"hosted_dns_zoneid":   fmt.Sprint(r.ZoneID),
		"hosted_dns_recordid": toString(r.ID),
	}
}

func (r AAAA) GetID() (uint, bool) {
	if r.ID == nil {
		return 0, false
	}

	return *r.ID, true
}

func (r AAAA) GetZoneID() uint {
	return r.ZoneID
}

func (r AAAA) Type() string {
	return "AAAA"
}
