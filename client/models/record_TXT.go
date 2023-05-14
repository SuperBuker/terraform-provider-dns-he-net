package models

import "fmt"

type TXT struct {
	ID      *uint
	ZoneID  uint
	Domain  string
	TTL     uint // seconds
	Data    string
	Dynamic bool
}

func ToTXT(r Record) TXT {
	return TXT{
		ID:      r.ID,
		ZoneID:  r.ZoneID,
		Domain:  r.Domain,
		TTL:     r.TTL,
		Data:    r.Data,
		Dynamic: r.Dynamic,
	}
}

func (r TXT) Serialise() map[string]string {
	return map[string]string{
		"Type":                "TXT",
		"hosted_dns_zoneid":   fmt.Sprint(r.ZoneID),
		"hosted_dns_recordid": toString(r.ID),
		//"Priority": "",
		"Name":    r.Domain,
		"Content": r.Data[1 : len(r.Data)-1],
		"TTL":     fmt.Sprint(r.TTL),
		"dynamic": b2s[r.Dynamic],
	}
}

func (r TXT) Refs() map[string]string {
	return map[string]string{
		"hosted_dns_zoneid":   fmt.Sprint(r.ZoneID),
		"hosted_dns_recordid": toString(r.ID),
	}
}

func (r TXT) GetID() (uint, bool) {
	if r.ID == nil {
		return 0, false
	}

	return *r.ID, true
}

func (r TXT) GetZoneID() uint {
	return r.ZoneID
}

func (r TXT) Type() string {
	return "TXT"
}
