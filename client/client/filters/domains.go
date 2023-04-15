package filters

import (
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"
)

// DomainById returns a domain by its ID.
// If the domain is not found, it returns an empty string and false.
func DomainById(domains []models.Domain, id uint) (models.Domain, bool) {
	for _, domain := range domains {
		if domain.Id == id {
			return domain, true
		}
	}

	return models.Domain{}, false
}

// DomainByTLD returns a domain by its top-level domain.
// If the domain is not found, it returns an empty string and false.
func DomainByTLD(domains []models.Domain, tld string) (models.Domain, bool) {
	for _, domain := range domains {
		if domain.Domain == tld {
			return domain, true
		}
	}

	return models.Domain{}, false
}

// LatestDomain returns the latest domain (highest ID) in a slice of domains.
// If the slice is empty, it returns an empty string and false.
func LatestDomain(domains []models.Domain) (d models.Domain, ok bool) {
	for _, domain := range domains {
		if domain.Id > d.Id {
			d = domain

			if !ok {
				ok = true
			}
		}
	}

	return
}
