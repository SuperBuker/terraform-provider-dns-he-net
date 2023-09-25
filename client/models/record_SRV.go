package models

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type SRV struct {
	ID       *uint
	ZoneID   uint
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
		return SRV{}, &ErrFormat{"", "unparseable SRV payload"}
	}

	srv := SRV{
		Target: s[2],
	}

	for i, x := range []*uint16{&srv.Weight, &srv.Port} {
		if a, err := strconv.Atoi(s[i]); err != nil {
			return SRV{}, err
		} else if a < 0 || a > math.MaxUint16 {
			field := []string{"Weight", "Port"}[i]
			return SRV{}, &ErrFormat{field, "value out of range"}
		} else {
			*x = uint16(a)
		}
	}

	return srv, nil
}

func ToSRV(r Record) (SRV, error) {
	if r.Priority == nil {
		return SRV{}, &ErrFormat{"Priority", "must be a positive integer"}
	}

	srv, err := parseSRVData(r.Data)
	if err != nil {
		return SRV{}, err
	}

	return SRV{
		ID:       r.ID,
		ZoneID:   r.ZoneID,
		Domain:   r.Domain,
		TTL:      r.TTL,
		Priority: *r.Priority,
		Weight:   srv.Weight,
		Port:     srv.Port,
		Target:   srv.Target,
	}, nil
}

func (r SRV) Equals(rx RecordX) bool {
	if rx == nil {
		return false
	} else if rx.Type() != "SRV" {
		return false
	} else if rec, ok := rx.(Record); ok {
		// Convert from Record
		var err error
		rx, err = ToSRV(rec)

		if err != nil {
			return false
		}
	}

	rsrv := rx.(SRV)

	return r.ZoneID == rsrv.ZoneID &&
		r.Domain == rsrv.Domain &&
		r.TTL == rsrv.TTL &&
		r.Priority == rsrv.Priority &&
		r.Weight == rsrv.Weight &&
		r.Port == rsrv.Port &&
		r.Target == rsrv.Target
}

func (r SRV) Serialise() map[string]string {
	return map[string]string{
		"Type":                "SRV",
		"hosted_dns_zoneid":   fmt.Sprint(r.ZoneID),
		"hosted_dns_recordid": toString(r.ID),
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
		"hosted_dns_zoneid":   fmt.Sprint(r.ZoneID),
		"hosted_dns_recordid": toString(r.ID),
	}
}

func (r SRV) GetID() (uint, bool) {
	if r.ID == nil {
		return 0, false
	}

	return *r.ID, true
}

func (r SRV) GetZoneID() uint {
	return r.ZoneID
}

func (r SRV) Type() string {
	return "SRV"
}
