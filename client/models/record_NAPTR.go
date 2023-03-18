package models

import "fmt"

type NAPTR struct {
	Id       *uint
	ParentId uint
	Domain   string
	TTL      uint // seconds
	Data     string
}

func ToNAPTR(r Record) NAPTR {
	return NAPTR{
		Id:       r.Id,
		ParentId: r.ParentId,
		Domain:   r.Domain,
		TTL:      r.TTL,
		Data:     r.Data,
	}
}

func (r NAPTR) Serialise() map[string]string {
	return map[string]string{
		"Type":                "NAPTR",
		"hosted_dns_zoneid":   fmt.Sprint(r.ParentId),
		"hosted_dns_recordid": toString(r.Id),
		//"Priority": "",
		"Name":    r.Domain,
		"Content": r.Data,
		"TTL":     fmt.Sprint(r.TTL),
	}
}

func (r NAPTR) Refs() map[string]string {
	return map[string]string{
		"hosted_dns_zoneid":   fmt.Sprint(r.ParentId),
		"hosted_dns_recordid": toString(r.Id),
	}
}

func (r NAPTR) GetId() (uint, bool) {
	if r.Id == nil {
		return 0, false
	}

	return *r.Id, true
}

func (r NAPTR) GetParentId() uint {
	return r.ParentId
}

func (r NAPTR) Type() string {
	return "NAPTR"
}
