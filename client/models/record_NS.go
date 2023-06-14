package models

import "fmt"

type NS struct {
	ID     *uint
	ZoneID uint
	Domain string
	TTL    uint // seconds
	Data   string
}

func ToNS(r Record) NS {
	return NS{
		ID:     r.ID,
		ZoneID: r.ZoneID,
		Domain: r.Domain,
		TTL:    r.TTL,
		Data:   r.Data,
	}
}

func (r NS) Equals(rx RecordX) bool {
	if rx == nil {
		return false
	} else if rx.Type() != "NS" {
		return false
	} else if rec, ok := rx.(Record); ok {
		// Convert from Record
		rx = ToNS(rec)
	}

	rns := rx.(NS)

	return r.ZoneID == rns.ZoneID &&
		r.Domain == rns.Domain &&
		r.TTL == rns.TTL &&
		r.Data == rns.Data
}

func (r NS) Serialise() map[string]string {
	return map[string]string{
		"Type":                "NS",
		"hosted_dns_zoneid":   fmt.Sprint(r.ZoneID),
		"hosted_dns_recordid": toString(r.ID),
		//"Priority": "",
		"Name":    r.Domain,
		"Content": r.Data,
		"TTL":     fmt.Sprint(r.TTL),
	}
}

func (r NS) Refs() map[string]string {
	return map[string]string{
		"hosted_dns_zoneid":   fmt.Sprint(r.ZoneID),
		"hosted_dns_recordid": toString(r.ID),
	}
}

func (r NS) GetID() (uint, bool) {
	if r.ID == nil {
		return 0, false
	}

	return *r.ID, true
}

func (r NS) GetZoneID() uint {
	return r.ZoneID
}
func (r NS) Type() string {
	return "NS"
}
