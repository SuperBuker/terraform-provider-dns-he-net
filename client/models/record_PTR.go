package models

import "fmt"

type PTR struct {
	Id       *uint
	ParentId uint
	Domain   string
	TTL      uint // seconds
	Data     string
}

func ToPTR(r Record) PTR {
	return PTR{
		Id:       r.Id,
		ParentId: r.ParentId,
		Domain:   r.Domain,
		TTL:      r.TTL,
		Data:     r.Data,
	}
}

func (r PTR) Serialise() map[string]string {
	return map[string]string{
		"Type":                "PTR",
		"hosted_dns_zoneid":   fmt.Sprint(r.ParentId),
		"hosted_dns_recordid": toString(r.Id),
		//"Priority": "",
		"Name":    r.Domain,
		"Content": r.Data,
		"TTL":     fmt.Sprint(r.TTL),
	}
}

func (r PTR) Refs() map[string]string {
	return map[string]string{
		"hosted_dns_zoneid":   fmt.Sprint(r.ParentId),
		"hosted_dns_recordid": toString(r.Id),
	}
}

func (r PTR) GetId() (uint, bool) {
	if r.Id == nil {
		return 0, false
	}

	return *r.Id, true
}

func (r PTR) GetParentId() uint {
	return r.ParentId
}

func (r PTR) Type() string {
	return "PTR"
}
