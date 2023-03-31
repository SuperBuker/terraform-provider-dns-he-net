package filters

import (
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"
)

func RecordById(records []models.Record, id uint) (models.Record, bool) {
	for _, record := range records {
		if record.Id == nil {
			//pass
		} else if *record.Id == id {
			return record, true
		}
	}

	return models.Record{}, false
}

func LatestRecord(records []models.Record) (r models.Record, ok bool) {
	var id uint

	for _, record := range records {
		if record.Id == nil {
			//pass
		} else if *record.Id >= id { // record my have id == zero
			r = record
			id = *record.Id

			if !ok {
				ok = true
			}
		}
	}
	return
}

func Record(records []models.Record, domain *string, typ *string) []models.Record {
	var fn func(record models.Record) bool
	var out []models.Record

	if _d, _t := domain != nil, typ != nil; _d && _t {
		fn = func(record models.Record) bool {
			return record.Domain == *domain && record.RecordType == *typ
		}
	} else if _d && !_t {
		fn = func(record models.Record) bool {
			return record.Domain == *domain
		}
	} else if !_d && _t {
		fn = func(record models.Record) bool {
			return record.RecordType == *typ
		}
	} else {
		out = make([]models.Record, len(records))
		copy(out, records)
		return out
	}

	out = make([]models.Record, 0)

	for _, record := range records {
		if fn(record) {
			out = append(out, record)
		}
	}
	return out
}
