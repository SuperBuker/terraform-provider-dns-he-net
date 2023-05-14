package models

import (
	"fmt"
	"strconv"
	"strings"
)

type SOA struct {
	ID      *uint
	ZoneID  uint
	Domain  string
	TTL     uint // seconds
	MName   string
	RName   string
	Serial  uint
	Refresh uint
	Retry   uint
	Expire  uint
}

func parseSOAData(data string) (SOA, error) {
	s := strings.Fields(data)

	if len(s) != 7 {
		return SOA{}, &ErrFormat{"", "unparseable SOA payload"}
	}

	soa := SOA{
		MName: s[0],
		RName: s[1],
	}

	for i, x := range []*uint{&soa.Serial, &soa.Refresh, &soa.Retry, &soa.Expire} {
		if a, err := strconv.Atoi(s[i+2]); err != nil {
			return SOA{}, err
		} else {
			*x = uint(a)
		}
	}

	return soa, nil
}

func ToSOA(r Record) (SOA, error) {
	soa, err := parseSOAData(r.Data)
	if err != nil {
		return SOA{}, err
	}

	return SOA{
		ID:      r.ID,
		ZoneID:  r.ZoneID,
		Domain:  r.Domain,
		TTL:     r.TTL,
		MName:   soa.MName,
		RName:   soa.RName,
		Serial:  soa.Serial,
		Refresh: soa.Refresh,
		Retry:   soa.Retry,
		Expire:  soa.Expire,
	}, nil
}

func (r SOA) Serialise() map[string]string {
	return map[string]string{
		"Type":                "SOA",
		"hosted_dns_zoneid":   fmt.Sprint(r.ZoneID),
		"hosted_dns_recordid": toString(r.ID),
		//"Priority": "",
		"Name": r.Domain,
		//"Content": r.Data, No need to serialise
		"TTL": fmt.Sprint(r.TTL),
	}
}

func (r SOA) Refs() map[string]string {
	return map[string]string{
		"hosted_dns_zoneid":   fmt.Sprint(r.ZoneID),
		"hosted_dns_recordid": toString(r.ID),
	}
}

func (r SOA) GetID() (uint, bool) {
	if r.ID == nil {
		return 0, false
	}

	return *r.ID, true
}

func (r SOA) GetZoneID() uint {
	return r.ZoneID
}

func (r SOA) Type() string {
	return "SOA"
}
