package models

import "fmt"

type SSHFP struct {
	Id       *uint
	ParentId uint
	Domain   string
	TTL      uint // seconds
	Data     string
}

func ToSSHFP(r Record) SSHFP {
	return SSHFP{
		Id:       r.Id,
		ParentId: r.ParentId,
		Domain:   r.Domain,
		TTL:      r.TTL,
		Data:     r.Data,
	}
}

func (r SSHFP) Serialise() map[string]string {
	return map[string]string{
		"Type":                "SSHFP",
		"hosted_dns_zoneid":   fmt.Sprint(r.ParentId),
		"hosted_dns_recordid": toString(r.Id),
		//"Priority": "",
		"Name":    r.Domain,
		"Content": r.Data,
		"TTL":     fmt.Sprint(r.TTL),
	}
}

func (r SSHFP) Refs() map[string]string {
	return map[string]string{
		"hosted_dns_zoneid":   fmt.Sprint(r.ParentId),
		"hosted_dns_recordid": toString(r.Id),
	}
}

func (r SSHFP) GetId() (uint, bool) {
	if r.Id == nil {
		return 0, false
	}

	return *r.Id, true
}

func (r SSHFP) GetParentId() uint {
	return r.ParentId
}

func (r SSHFP) Type() string {
	return "SSHFP"
}