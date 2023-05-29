package models

import "fmt"

type DDNSKey struct {
	Domain string
	ZoneID uint
	Key    string
}

func (dk DDNSKey) Serialise() map[string]string {
	return map[string]string{
		"hosted_dns_zoneid": fmt.Sprint(dk.ZoneID),
		"Name":              dk.Domain,
		"Key":               dk.Key,
		"Key2":              dk.Key,
	}
}

func (dk DDNSKey) Refs() map[string]string {
	return map[string]string{
		"hosted_dns_zoneid": fmt.Sprint(dk.ZoneID),
		"Name":              dk.Domain,
	}
}

func (dk DDNSKey) GetDomain() string {
	return dk.Domain
}

func (dk DDNSKey) GetZoneID() uint {
	return dk.ZoneID
}
