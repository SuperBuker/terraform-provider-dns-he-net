package models

import "fmt"

type A struct {
	ID      *uint
	ZoneID  uint
	Domain  string
	TTL     uint // seconds
	Data    string
	Dynamic bool
}

func ToA(r Record) A {
	return A{
		ID:      r.ID,
		ZoneID:  r.ZoneID,
		Domain:  r.Domain,
		TTL:     r.TTL,
		Data:    r.Data,
		Dynamic: r.Dynamic,
	}
}

func (r A) Serialise() map[string]string {
	return map[string]string{
		"Type":                "a",
		"hosted_dns_zoneid":   fmt.Sprint(r.ZoneID),
		"hosted_dns_recordid": toString(r.ID),
		//"Priority":            "",
		"Name":    r.Domain,
		"Content": r.Data,
		"TTL":     fmt.Sprint(r.TTL),
		"dynamic": b2s[r.Dynamic],
	}
}

func (r A) Refs() map[string]string {
	return map[string]string{
		"hosted_dns_zoneid":   fmt.Sprint(r.ZoneID),
		"hosted_dns_recordid": toString(r.ID),
	}
}

func (r A) GetID() (uint, bool) {
	if r.ID == nil {
		return 0, false
	}

	return *r.ID, true
}

func (r A) GetZoneID() uint {
	return r.ZoneID
}

func (r A) Type() string {
	return "A"
}
