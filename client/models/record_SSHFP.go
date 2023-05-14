package models

import "fmt"

type SSHFP struct {
	ID     *uint
	ZoneID uint
	Domain string
	TTL    uint // seconds
	Data   string
}

func ToSSHFP(r Record) SSHFP {
	return SSHFP{
		ID:     r.ID,
		ZoneID: r.ZoneID,
		Domain: r.Domain,
		TTL:    r.TTL,
		Data:   r.Data,
	}
}

func (r SSHFP) Serialise() map[string]string {
	return map[string]string{
		"Type":                "SSHFP",
		"hosted_dns_zoneid":   fmt.Sprint(r.ZoneID),
		"hosted_dns_recordid": toString(r.ID),
		//"Priority": "",
		"Name":    r.Domain,
		"Content": r.Data,
		"TTL":     fmt.Sprint(r.TTL),
	}
}

func (r SSHFP) Refs() map[string]string {
	return map[string]string{
		"hosted_dns_zoneid":   fmt.Sprint(r.ZoneID),
		"hosted_dns_recordid": toString(r.ID),
	}
}

func (r SSHFP) GetID() (uint, bool) {
	if r.ID == nil {
		return 0, false
	}

	return *r.ID, true
}

func (r SSHFP) GetZoneID() uint {
	return r.ZoneID
}

func (r SSHFP) Type() string {
	return "SSHFP"
}
