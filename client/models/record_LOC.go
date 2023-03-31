package models

import "fmt"

type LOC struct {
	Id       *uint
	ParentId uint
	Domain   string
	TTL      uint // seconds
	Data     string
}

func ToLOC(r Record) LOC {
	return LOC{
		Id:       r.Id,
		ParentId: r.ParentId,
		Domain:   r.Domain,
		TTL:      r.TTL,
		Data:     r.Data,
	}
}

func (r LOC) Serialise() map[string]string {
	return map[string]string{
		"Type":                "LOC",
		"hosted_dns_zoneid":   fmt.Sprint(r.ParentId),
		"hosted_dns_recordid": toString(r.Id),
		//"Priority": "",
		"Name":    r.Domain,
		"Content": r.Data,
		"TTL":     fmt.Sprint(r.TTL),
	}
}

func (r LOC) Refs() map[string]string {
	return map[string]string{
		"hosted_dns_zoneid":   fmt.Sprint(r.ParentId),
		"hosted_dns_recordid": toString(r.Id),
	}
}

func (r LOC) GetId() (uint, bool) {
	if r.Id == nil {
		return 0, false
	}

	return *r.Id, true
}

func (r LOC) GetParentId() uint {
	return r.ParentId
}

func (r LOC) Type() string {
	return "LOC"
}
