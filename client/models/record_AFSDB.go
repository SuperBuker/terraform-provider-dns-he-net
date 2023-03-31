package models

import "fmt"

type AFSDB struct {
	Id       *uint
	ParentId uint
	Domain   string
	TTL      uint // seconds
	Data     string
	Dynamic  bool
}

func ToAFSDB(r Record) AFSDB {
	return AFSDB{
		Id:       r.Id,
		ParentId: r.ParentId,
		Domain:   r.Domain,
		TTL:      r.TTL,
		Data:     r.Data,
		Dynamic:  r.Dynamic,
	}
}

func (r AFSDB) Serialise() map[string]string {
	return map[string]string{
		"Type":                "AFSDB",
		"hosted_dns_zoneid":   fmt.Sprint(r.ParentId),
		"hosted_dns_recordid": toString(r.Id),
		//"Priority": "",
		"Name":    r.Domain,
		"Content": r.Data,
		"TTL":     fmt.Sprint(r.TTL),
		"Dynamic": b2s[r.Dynamic],
	}
}

func (r AFSDB) Refs() map[string]string {
	return map[string]string{
		"hosted_dns_zoneid":   fmt.Sprint(r.ParentId),
		"hosted_dns_recordid": toString(r.Id),
	}
}

func (r AFSDB) GetId() (uint, bool) {
	if r.Id == nil {
		return 0, false
	}

	return *r.Id, true
}

func (r AFSDB) GetParentId() uint {
	return r.ParentId
}

func (r AFSDB) Type() string {
	return "AFSDB"
}
