package filters

import (
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"
)

// RecordById returns a record by its ID.
// If the record is not found, it returns an empty record and false.
func RecordById(records []models.Record, id uint) (models.Record, bool) {
	for _, record := range records {
		if record.ID == nil {
			//pass
		} else if *record.ID == id {
			return record, true
		}
	}

	return models.Record{}, false
}

// MatchRecord returns the record matching the input in a slice of records.
// If the slice doesn't contain any matching record, it returns an empty record
// and false.
func MatchRecord(records []models.Record, rx models.RecordX) (r models.Record, ok bool) {
	for _, record := range records {
		if rx.Equals(record) {
			return record, true
		}
	}

	return models.Record{}, false
}

// LatestRecord returns the latest record (highest ID) in a slice of records.
// If the slice doesn't contain any record with ID, it returns an empty record
// and false.
func LatestRecord(records []models.Record) (r models.Record, ok bool) {
	var id uint

	for _, record := range records {
		if record.ID == nil {
			//pass
		} else if *record.ID >= id { // record may have id == zero
			r = record
			id = *record.ID

			if !ok {
				ok = true
			}
		}
	}
	return
}

// Record returns a slice of records that match the domain name and/or type.
// Only the not nil fields are used for filtering
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
