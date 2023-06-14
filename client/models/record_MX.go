package models

import (
	"fmt"
)

type MX struct {
	ID       *uint
	ZoneID   uint
	Domain   string
	TTL      uint // seconds
	Priority uint16
	Data     string
}

func ToMX(r Record) (MX, error) {
	if r.Priority == nil {
		return MX{}, &ErrFormat{"Priority", "must be a positive integer"}
	}

	return MX{
		ID:       r.ID,
		ZoneID:   r.ZoneID,
		Domain:   r.Domain,
		TTL:      r.TTL,
		Priority: *r.Priority,
		Data:     r.Data,
	}, nil
}

func (r MX) Equals(rx RecordX) bool {
	if rx == nil {
		return false
	} else if rx.Type() != "MX" {
		return false
	} else if rec, ok := rx.(Record); ok {
		// Convert from Record
		var err error
		rx, err = ToMX(rec)

		if err != nil {
			return false
		}
	}

	rmx := rx.(MX)

	return r.ZoneID == rmx.ZoneID &&
		r.Domain == rmx.Domain &&
		r.TTL == rmx.TTL &&
		r.Priority == rmx.Priority &&
		r.Data == rmx.Data
}

func (r MX) Serialise() map[string]string {
	return map[string]string{
		"Type":                "MX",
		"hosted_dns_zoneid":   fmt.Sprint(r.ZoneID),
		"hosted_dns_recordid": toString(r.ID),
		"Priority":            fmt.Sprint(r.Priority),
		"Name":                r.Domain,
		"Content":             r.Data,
		"TTL":                 fmt.Sprint(r.TTL),
	}
}

func (r MX) Refs() map[string]string {
	return map[string]string{
		"hosted_dns_zoneid":   fmt.Sprint(r.ZoneID),
		"hosted_dns_recordid": toString(r.ID),
	}
}

func (r MX) GetID() (uint, bool) {
	if r.ID == nil {
		return 0, false
	}

	return *r.ID, true
}

func (r MX) GetZoneID() uint {
	return r.ZoneID
}

func (r MX) Type() string {
	return "MX"
}
