package models

import "fmt"

type SPF struct {
	ID     *uint
	ZoneID uint
	Domain string
	TTL    uint // seconds
	Data   string
}

func ToSPF(r Record) SPF {
	return SPF{
		ID:     r.ID,
		ZoneID: r.ZoneID,
		Domain: r.Domain,
		TTL:    r.TTL,
		Data:   r.Data,
	}
}

func (r SPF) Equals(rx RecordX) bool {
	if rx == nil {
		return false
	} else if rx.Type() != "SPF" {
		return false
	} else if rec, ok := rx.(Record); ok {
		// Convert from Record
		rx = ToSPF(rec)
	}

	rspf := rx.(SPF)

	return r.ZoneID == rspf.ZoneID &&
		r.Domain == rspf.Domain &&
		r.TTL == rspf.TTL &&
		r.Data == rspf.Data
}

func (r SPF) Serialise() map[string]string {
	return map[string]string{
		"Type":                "SPF",
		"hosted_dns_zoneid":   fmt.Sprint(r.ZoneID),
		"hosted_dns_recordid": toString(r.ID),
		//"Priority": "",
		"Name":    r.Domain,
		"Content": r.Data,
		"TTL":     fmt.Sprint(r.TTL),
	}
}

func (r SPF) Refs() map[string]string {
	return map[string]string{
		"hosted_dns_zoneid":   fmt.Sprint(r.ZoneID),
		"hosted_dns_recordid": toString(r.ID),
	}
}

func (r SPF) GetID() (uint, bool) {
	if r.ID == nil {
		return 0, false
	}

	return *r.ID, true
}

func (r SPF) GetZoneID() uint {
	return r.ZoneID
}

func (r SPF) Type() string {
	return "SPF"
}
