package models

import "fmt"

type NAPTR struct {
	ID     *uint
	ZoneID uint
	Domain string
	TTL    uint // seconds
	Data   string
}

func ToNAPTR(r Record) NAPTR {
	return NAPTR{
		ID:     r.ID,
		ZoneID: r.ZoneID,
		Domain: r.Domain,
		TTL:    r.TTL,
		Data:   r.Data,
	}
}

func (r NAPTR) Equals(rx RecordX) bool {
	if rx == nil {
		return false
	} else if rx.Type() != "NAPTR" {
		return false
	} else if rec, ok := rx.(Record); ok {
		// Convert from Record
		rx = ToNAPTR(rec)
	}

	rnaptr := rx.(NAPTR)

	return r.ZoneID == rnaptr.ZoneID &&
		r.Domain == rnaptr.Domain &&
		r.TTL == rnaptr.TTL &&
		r.Data == rnaptr.Data
}

func (r NAPTR) Serialise() map[string]string {
	return map[string]string{
		"Type":                "NAPTR",
		"hosted_dns_zoneid":   fmt.Sprint(r.ZoneID),
		"hosted_dns_recordid": toString(r.ID),
		//"Priority": "",
		"Name":    r.Domain,
		"Content": r.Data,
		"TTL":     fmt.Sprint(r.TTL),
	}
}

func (r NAPTR) Refs() map[string]string {
	return map[string]string{
		"hosted_dns_zoneid":   fmt.Sprint(r.ZoneID),
		"hosted_dns_recordid": toString(r.ID),
	}
}

func (r NAPTR) GetID() (uint, bool) {
	if r.ID == nil {
		return 0, false
	}

	return *r.ID, true
}

func (r NAPTR) GetZoneID() uint {
	return r.ZoneID
}

func (r NAPTR) Type() string {
	return "NAPTR"
}
