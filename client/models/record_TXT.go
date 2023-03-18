package models

import "fmt"

type TXT struct {
	Id       *uint
	ParentId uint
	Domain   string
	TTL      uint // seconds
	Data     string
	Dynamic  bool
}

func ToTXT(r Record) TXT {
	return TXT{
		Id:       r.Id,
		ParentId: r.ParentId,
		Domain:   r.Domain,
		TTL:      r.TTL,
		Data:     r.Data,
		Dynamic:  r.Dynamic,
	}
}

func (r TXT) Serialise() map[string]string {
	return map[string]string{
		"Type":                "TXT",
		"hosted_dns_zoneid":   fmt.Sprint(r.ParentId),
		"hosted_dns_recordid": toString(r.Id),
		//"Priority": "",
		"Name":    r.Domain,
		"Content": r.Data,
		"TTL":     fmt.Sprint(r.TTL),
		"Dynamic": b2s[r.Dynamic],
	}
}

func (r TXT) Refs() map[string]string {
	return map[string]string{
		"hosted_dns_zoneid":   fmt.Sprint(r.ParentId),
		"hosted_dns_recordid": toString(r.Id),
	}
}

func (r TXT) GetId() (uint, bool) {
	if r.Id == nil {
		return 0, false
	}

	return *r.Id, true
}

func (r TXT) GetParentId() uint {
	return r.ParentId
}

func (r TXT) Type() string {
	return "TXT"
}
