package models

import "fmt"

type HINFO struct {
	ID     *uint
	ZoneID uint
	Domain string
	TTL    uint // seconds
	Data   string
}

func ToHINFO(r Record) HINFO {
	return HINFO{
		ID:     r.ID,
		ZoneID: r.ZoneID,
		Domain: r.Domain,
		TTL:    r.TTL,
		Data:   r.Data,
	}
}

func (r HINFO) Serialise() map[string]string {
	return map[string]string{
		"Type":                "HINFO",
		"hosted_dns_zoneid":   fmt.Sprint(r.ZoneID),
		"hosted_dns_recordid": toString(r.ID),
		//"Priority": "",
		"Name":    r.Domain,
		"Content": r.Data,
		"TTL":     fmt.Sprint(r.TTL),
	}
}

func (r HINFO) Refs() map[string]string {
	return map[string]string{
		"hosted_dns_zoneid":   fmt.Sprint(r.ZoneID),
		"hosted_dns_recordid": toString(r.ID),
	}
}

func (r HINFO) GetID() (uint, bool) {
	if r.ID == nil {
		return 0, false
	}

	return *r.ID, true
}

func (r HINFO) GetZoneID() uint {
	return r.ZoneID
}

func (r HINFO) Type() string {
	return "HINFO"
}
