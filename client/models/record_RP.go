package models

import "fmt"

type RP struct {
	ID     *uint
	ZoneID uint
	Domain string
	TTL    uint // seconds
	Data   string
}

func ToRP(r Record) RP {
	return RP{
		ID:     r.ID,
		ZoneID: r.ZoneID,
		Domain: r.Domain,
		TTL:    r.TTL,
		Data:   r.Data,
	}
}

func (r RP) Equals(rx RecordX) bool {
	if rx == nil {
		return false
	} else if rx.Type() != "RP" {
		return false
	} else if rec, ok := rx.(Record); ok {
		// Convert from Record
		rx = ToRP(rec)
	}

	rrp := rx.(RP)

	return r.ZoneID == rrp.ZoneID &&
		r.Domain == rrp.Domain &&
		r.TTL == rrp.TTL &&
		r.Data == rrp.Data
}

func (r RP) Serialise() map[string]string {
	return map[string]string{
		"Type":                "RP",
		"hosted_dns_zoneid":   fmt.Sprint(r.ZoneID),
		"hosted_dns_recordid": toString(r.ID),
		//"Priority": "",
		"Name":    r.Domain,
		"Content": r.Data,
		"TTL":     fmt.Sprint(r.TTL),
	}
}

func (r RP) Refs() map[string]string {
	return map[string]string{
		"hosted_dns_zoneid":   fmt.Sprint(r.ZoneID),
		"hosted_dns_recordid": toString(r.ID),
	}
}

func (r RP) GetID() (uint, bool) {
	if r.ID == nil {
		return 0, false
	}

	return *r.ID, true
}

func (r RP) GetZoneID() uint {
	return r.ZoneID
}

func (r RP) Type() string {
	return "RP"
}
