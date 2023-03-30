package models

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type SOA struct {
	Id       *uint
	ParentId uint
	Domain   string
	TTL      uint // seconds
	MName    string
	RName    string
	Serial   uint
	Refresh  uint
	Retry    uint
	Expire   uint
}

func parseSOAData(data string) (SOA, error) {
	s := strings.Fields(data)

	if len(s) != 7 {
		return SOA{}, errors.New("unparseable SOA payload")
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
		Id:       r.Id,
		ParentId: r.ParentId,
		Domain:   r.Domain,
		TTL:      r.TTL,
		MName:    soa.MName,
		RName:    soa.RName,
		Serial:   soa.Serial,
		Refresh:  soa.Refresh,
		Retry:    soa.Retry,
		Expire:   soa.Expire,
	}, nil
}

func (r SOA) Serialise() map[string]string {
	return map[string]string{
		"Type":                "SOA",
		"hosted_dns_zoneid":   fmt.Sprint(r.ParentId),
		"hosted_dns_recordid": toString(r.Id),
		//"Priority": "",
		"Name": r.Domain,
		//"Content": r.Data, No need to serialise
		"TTL": fmt.Sprint(r.TTL),
	}
}

func (r SOA) Refs() map[string]string {
	return map[string]string{
		"hosted_dns_zoneid":   fmt.Sprint(r.ParentId),
		"hosted_dns_recordid": toString(r.Id),
	}
}

func (r SOA) GetId() (uint, bool) {
	if r.Id == nil {
		return 0, false
	}

	return *r.Id, true
}

func (r SOA) GetParentId() uint {
	return r.ParentId
}

func (r SOA) Type() string {
	return "SOA"
}
