package models

import "fmt"

type PTR struct {
	ID     *uint
	ZoneID uint
	Domain string
	TTL    uint // seconds
	Data   string
}

func ToPTR(r Record) PTR {
	return PTR{
		ID:     r.ID,
		ZoneID: r.ZoneID,
		Domain: r.Domain,
		TTL:    r.TTL,
		Data:   r.Data,
	}
}

func (r PTR) Equals(rx RecordX) bool {
	if rx == nil {
		return false
	} else if rx.Type() != "PTR" {
		return false
	} else if rec, ok := rx.(Record); ok {
		// Convert from Record
		rx = ToPTR(rec)
	}

	rptr := rx.(PTR)

	return r.ZoneID == rptr.ZoneID &&
		r.Domain == rptr.Domain &&
		r.TTL == rptr.TTL &&
		r.Data == rptr.Data
}

func (r PTR) Serialise() map[string]string {
	return map[string]string{
		"Type":                "PTR",
		"hosted_dns_zoneid":   fmt.Sprint(r.ZoneID),
		"hosted_dns_recordid": toString(r.ID),
		//"Priority": "",
		"Name":    r.Domain,
		"Content": r.Data,
		"TTL":     fmt.Sprint(r.TTL),
	}
}

func (r PTR) Refs() map[string]string {
	return map[string]string{
		"hosted_dns_zoneid":   fmt.Sprint(r.ZoneID),
		"hosted_dns_recordid": toString(r.ID),
	}
}

func (r PTR) GetID() (uint, bool) {
	if r.ID == nil {
		return 0, false
	}

	return *r.ID, true
}

func (r PTR) GetZoneID() uint {
	return r.ZoneID
}

func (r PTR) Type() string {
	return "PTR"
}
