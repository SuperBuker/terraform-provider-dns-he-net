package models

import "fmt"

type ALIAS struct {
	ID     *uint
	ZoneID uint
	Domain string
	TTL    uint // seconds
	Data   string
}

func ToALIAS(r Record) ALIAS {
	return ALIAS{
		ID:     r.ID,
		ZoneID: r.ZoneID,
		Domain: r.Domain,
		TTL:    r.TTL,
		Data:   r.Data,
	}
}

func (r ALIAS) Equals(rx RecordX) bool {
	if rx == nil {
		return false
	} else if rx.Type() != "ALIAS" {
		return false
	} else if rec, ok := rx.(Record); ok {
		// Convert from Record
		rx = ToALIAS(rec)
	}

	ralias := rx.(ALIAS)

	return r.ZoneID == ralias.ZoneID &&
		r.Domain == ralias.Domain &&
		r.TTL == ralias.TTL &&
		r.Data == ralias.Data
}

func (r ALIAS) Serialise() map[string]string {
	return map[string]string{
		"Type":                "ALIAS",
		"hosted_dns_zoneid":   fmt.Sprint(r.ZoneID),
		"hosted_dns_recordid": toString(r.ID),
		//"Priority":            "",
		"Name":    r.Domain,
		"Content": r.Data,
		"TTL":     fmt.Sprint(r.TTL),
	}
}

func (r ALIAS) Refs() map[string]string {
	return map[string]string{
		"hosted_dns_zoneid":   fmt.Sprint(r.ZoneID),
		"hosted_dns_recordid": toString(r.ID),
	}
}

func (r ALIAS) GetID() (uint, bool) {
	if r.ID == nil {
		return 0, false
	}

	return *r.ID, true
}

func (r ALIAS) GetZoneID() uint {
	return r.ZoneID
}

func (r ALIAS) Type() string {
	return "ALIAS"
}
