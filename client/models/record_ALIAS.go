package models

import "fmt"

type ALIAS struct {
	Id       *uint
	ParentId uint
	Domain   string
	TTL      uint // seconds
	Data     string
}

func ToALIAS(r Record) ALIAS {
	return ALIAS{
		Id:       r.Id,
		ParentId: r.ParentId,
		Domain:   r.Domain,
		TTL:      r.TTL,
		Data:     r.Data,
	}
}

func (r ALIAS) Serialise() map[string]string {
	return map[string]string{
		"Type":                "ALIAS",
		"hosted_dns_zoneid":   fmt.Sprint(r.ParentId),
		"hosted_dns_recordid": toString(r.Id),
		//"Priority":            "",
		"Name":    r.Domain,
		"Content": r.Data,
		"TTL":     fmt.Sprint(r.TTL),
	}
}

func (r ALIAS) Refs() map[string]string {
	return map[string]string{
		"hosted_dns_zoneid":   fmt.Sprint(r.ParentId),
		"hosted_dns_recordid": toString(r.Id),
	}
}

func (r ALIAS) GetId() (uint, bool) {
	if r.Id == nil {
		return 0, false
	}

	return *r.Id, true
}

func (r ALIAS) GetParentId() uint {
	return r.ParentId
}

func (r ALIAS) Type() string {
	return "ALIAS"
}
