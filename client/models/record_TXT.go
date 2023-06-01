package models

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/utils"
)

var txtInData = regexp.MustCompile(`(?:"([ -~]{255})" )|"([ -~]{1,255})"$`)
var txtOutData = regexp.MustCompile(`^"([ -~]*)"$`)

type TXT struct {
	ID      *uint
	ZoneID  uint
	Domain  string
	TTL     uint // seconds
	Data    string
	Dynamic bool
}

// concatTXTData concatenates the received field to reconstruct the original value.
// By default, the provider splits the data in chunks of 255 characters.
func concatTXTData(data string) string {
	var b strings.Builder
	for _, x := range txtInData.FindAllStringSubmatch(data, -1) {
		b.WriteString(x[1] + x[2]) // Either one of them is empty
	}

	return `"` + b.String() + `"`
}

// splitTXTData splits the data in chunks of 255 characters which are sent quoted and
// separated by a whitespace. This processing is automatically done by the provider
// on the server side, we just handle it on advance to mimic the regular website
// requests.
func splitTXTData(data string) string {
	if txtOutData.MatchString(data) {
		data = data[1 : len(data)-1] // Remove the first and last "
	}

	fn := func(s string) string {
		return `"` + s + `"`
	}

	return strings.Join(utils.ApplyToSlice(fn, utils.SplitByLen(data, 255)), " ")
}

func ToTXT(r Record) TXT {
	return TXT{
		ID:      r.ID,
		ZoneID:  r.ZoneID,
		Domain:  r.Domain,
		TTL:     r.TTL,
		Data:    concatTXTData(r.Data),
		Dynamic: r.Dynamic,
	}
}

func (r TXT) Serialise() map[string]string {
	return map[string]string{
		"Type":                "TXT",
		"hosted_dns_zoneid":   fmt.Sprint(r.ZoneID),
		"hosted_dns_recordid": toString(r.ID),
		//"Priority": "",
		"Name":    r.Domain,
		"Content": splitTXTData(r.Data),
		"TTL":     fmt.Sprint(r.TTL),
		"dynamic": b2s[r.Dynamic],
	}
}

func (r TXT) Refs() map[string]string {
	return map[string]string{
		"hosted_dns_zoneid":   fmt.Sprint(r.ZoneID),
		"hosted_dns_recordid": toString(r.ID),
	}
}

func (r TXT) GetID() (uint, bool) {
	if r.ID == nil {
		return 0, false
	}

	return *r.ID, true
}

func (r TXT) GetZoneID() uint {
	return r.ZoneID
}

func (r TXT) Type() string {
	return "TXT"
}
