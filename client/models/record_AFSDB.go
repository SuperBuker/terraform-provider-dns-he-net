package models

import "fmt"

type AFSDB struct {
	ID      *uint
	ZoneID  uint
	Domain  string
	TTL     uint // seconds
	Data    string
	Dynamic bool
}

func ToAFSDB(r Record) AFSDB {
	return AFSDB{
		ID:      r.ID,
		ZoneID:  r.ZoneID,
		Domain:  r.Domain,
		TTL:     r.TTL,
		Data:    r.Data,
		Dynamic: r.Dynamic,
	}
}

func (r AFSDB) Serialise() map[string]string {
	return map[string]string{
		"Type":                "AFSDB",
		"hosted_dns_zoneid":   fmt.Sprint(r.ZoneID),
		"hosted_dns_recordid": toString(r.ID),
		//"Priority": "",
		"Name":    r.Domain,
		"Content": r.Data,
		"TTL":     fmt.Sprint(r.TTL),
		"dynamic": b2s[r.Dynamic],
	}
}

func (r AFSDB) Refs() map[string]string {
	return map[string]string{
		"hosted_dns_zoneid":   fmt.Sprint(r.ZoneID),
		"hosted_dns_recordid": toString(r.ID),
	}
}

func (r AFSDB) GetID() (uint, bool) {
	if r.ID == nil {
		return 0, false
	}

	return *r.ID, true
}

func (r AFSDB) GetZoneID() uint {
	return r.ZoneID
}

func (r AFSDB) Type() string {
	return "AFSDB"
}
