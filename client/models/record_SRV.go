package models

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type SRV struct {
	Id       *uint
	ParentId uint
	Domain   string
	TTL      uint // seconds
	Priority uint16
	Weight   uint16
	Port     uint16
	Target   string
}

func parseSRVData(data string) (SRV, error) {
	s := strings.Fields(data)

	if len(s) != 3 {
		return SRV{}, errors.New("unparseable SRV payload")
	}

	srv := SRV{
		Target: s[2],
	}

	for i, x := range []*uint16{&srv.Weight, &srv.Port} {
		if a, err := strconv.Atoi(s[i]); err != nil {
			return SRV{}, err
		} else {
			*x = uint16(a)
		}
	}

	return srv, nil
}

func ToSRV(r Record) (SRV, error) {
	if r.Priority == nil {
		return SRV{}, errors.New("invalid priority, must be a positive integer")
	}

	srv, err := parseSRVData(r.Data)
	if err != nil {
		return SRV{}, err
	}

	return SRV{
		Id:       r.Id,
		ParentId: r.ParentId,
		Domain:   r.Domain,
		TTL:      r.TTL,
		Priority: *r.Priority,
		Weight:   srv.Weight,
		Port:     srv.Port,
		Target:   srv.Target,
	}, nil
}

func (r SRV) Serialise() map[string]string {
	return map[string]string{
		"Type":                "SRV",
		"hosted_dns_zoneid":   fmt.Sprint(r.ParentId),
		"hosted_dns_recordid": toString(r.Id),
		"Priority":            fmt.Sprint(r.Priority),
		"Name":                r.Domain,
		"Weight":              fmt.Sprint(r.Weight),
		"Port":                fmt.Sprint(r.Port),
		"Target":              r.Target,
		"TTL":                 fmt.Sprint(r.TTL),
	}
}

func (r SRV) Refs() map[string]string {
	return map[string]string{
		"hosted_dns_zoneid":   fmt.Sprint(r.ParentId),
		"hosted_dns_recordid": toString(r.Id),
	}
}

func (r SRV) GetId() (uint, bool) {
	if r.Id == nil {
		return 0, false
	}

	return *r.Id, true
}

func (r SRV) GetParentId() uint {
	return r.ParentId
}

func (r SRV) Type() string {
	return "SRV"
}
