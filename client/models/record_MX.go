package models

import (
	"fmt"
)

type MX struct {
	Id       *uint
	ParentId uint
	Domain   string
	TTL      uint // seconds
	Priority uint16
	Data     string
}

func ToMX(r Record) (MX, error) {
	if r.Priority == nil {
		return MX{}, &ErrFormat{"Priority", "must be a positive integer"}
	}

	return MX{
		Id:       r.Id,
		ParentId: r.ParentId,
		Domain:   r.Domain,
		TTL:      r.TTL,
		Priority: *r.Priority,
		Data:     r.Data,
	}, nil
}

func (r MX) Serialise() map[string]string {
	return map[string]string{
		"Type":                "MX",
		"hosted_dns_zoneid":   fmt.Sprint(r.ParentId),
		"hosted_dns_recordid": toString(r.Id),
		"Priority":            fmt.Sprint(r.Priority),
		"Name":                r.Domain,
		"Content":             r.Data,
		"TTL":                 fmt.Sprint(r.TTL),
	}
}

func (r MX) Refs() map[string]string {
	return map[string]string{
		"hosted_dns_zoneid":   fmt.Sprint(r.ParentId),
		"hosted_dns_recordid": toString(r.Id),
	}
}

func (r MX) GetId() (uint, bool) {
	if r.Id == nil {
		return 0, false
	}

	return *r.Id, true
}

func (r MX) GetParentId() uint {
	return r.ParentId
}

func (r MX) Type() string {
	return "MX"
}
