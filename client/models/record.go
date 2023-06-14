package models

import "strings"

type Record struct {
	ID         *uint
	ZoneID     uint
	Domain     string
	RecordType string // to improve
	TTL        uint   // seconds
	Priority   *uint16
	Data       string
	Dynamic    bool
	Locked     bool
}

func (r Record) ToX() (RecordX, error) {
	switch r.Type() {
	case "SOA":
		return ToSOA(r)
	case "A":
		return ToA(r), nil
	case "AAAA":
		return ToAAAA(r), nil
	case "CNAME":
		return ToCNAME(r), nil
	case "ALIAS":
		return ToALIAS(r), nil
	case "MX":
		return ToMX(r)
	case "NS":
		return ToNS(r), nil
	case "TXT":
		return ToTXT(r), nil
	case "CAA":
		return ToCAA(r), nil
	case "AFSDB":
		return ToAFSDB(r), nil
	case "HINFO":
		return ToHINFO(r), nil
	case "RP":
		return ToRP(r), nil
	case "LOC":
		return ToLOC(r), nil
	case "NAPTR":
		return ToNAPTR(r), nil
	case "PTR":
		return ToPTR(r), nil
	case "SSHFP":
		return ToSSHFP(r), nil
	case "SPF":
		return ToSPF(r), nil
	case "SRV":
		return ToSRV(r)
	}
	return nil, nil // This needs an error, whatever
}

func (r Record) Equals(rx RecordX) bool {
	if rx == nil {
		// pass
	} else if rx.Type() != r.Type() {
		// pass
	} else if rec, err := r.ToX(); err == nil {
		return rec.Equals(rx)
	}

	return false
}

func (r Record) Serialise() map[string]string {
	if rx, err := r.ToX(); err == nil {
		return rx.Serialise()
	}

	return nil
}

func (r Record) Refs() map[string]string {
	if rx, err := r.ToX(); err == nil {
		return rx.Refs()
	}

	return nil
}

func (r Record) GetID() (uint, bool) {
	if r.ID == nil {
		return 0, false
	}

	return *r.ID, true
}

func (r Record) GetZoneID() uint {
	return r.ZoneID
}

func (r Record) Type() string {
	return strings.ToUpper(r.RecordType)
}
