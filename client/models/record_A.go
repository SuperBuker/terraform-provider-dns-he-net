package models

import "fmt"

type A struct {
	Id       *uint
	ParentId uint
	Domain   string
	TTL      uint // seconds
	Data     string
	Dynamic  bool
}

func ToA(r Record) A {
	return A{
		Id:       r.Id,
		ParentId: r.ParentId,
		Domain:   r.Domain,
		TTL:      r.TTL,
		Data:     r.Data,
		Dynamic:  r.Dynamic,
	}
}

func (r A) Serialise() map[string]string {
	return map[string]string{
		"Type":                "a",
		"hosted_dns_zoneid":   fmt.Sprint(r.ParentId),
		"hosted_dns_recordid": toString(r.Id),
		//"Priority":            "",
		"Name":    r.Domain,
		"Content": r.Data,
		"TTL":     fmt.Sprint(r.TTL),
		"Dynamic": b2s[r.Dynamic],
	}
}

func (r A) Refs() map[string]string {
	return map[string]string{
		"hosted_dns_zoneid":   fmt.Sprint(r.ParentId),
		"hosted_dns_recordid": toString(r.Id),
	}
}

func (r A) GetId() (uint, bool) {
	if r.Id == nil {
		return 0, false
	}

	return *r.Id, true
}

func (r A) GetParentId() uint {
	return r.ParentId
}

func (r A) Type() string {
	return "A"
}
