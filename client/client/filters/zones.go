package filters

import (
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"
)

// ZoneById returns a zone by its ID.
// If the zone is not found, it returns an empty string and false.
func ZoneById(zones []models.Zone, id uint) (models.Zone, bool) {
	for _, zone := range zones {
		if zone.ID == id {
			return zone, true
		}
	}

	return models.Zone{}, false
}

// ZoneByName returns a zone by its second-level zone.
// If the zone is not found, it returns an empty string and false.
func ZoneByName(zones []models.Zone, name string) (models.Zone, bool) {
	for _, zone := range zones {
		if zone.Name == name {
			return zone, true
		}
	}

	return models.Zone{}, false
}

// LatestZone returns the latest zone (highest ID) in a slice of zones.
// If the slice is empty, it returns an empty string and false.
func LatestZone(zones []models.Zone) (d models.Zone, ok bool) {
	for _, zone := range zones {
		if zone.ID > d.ID {
			d = zone

			if !ok {
				ok = true
			}
		}
	}

	return
}
