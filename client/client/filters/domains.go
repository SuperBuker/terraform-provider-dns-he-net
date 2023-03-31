package filters

import (
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"
)

func DomainById(doamins []models.Domain, id uint) (models.Domain, bool) {
	for _, domain := range doamins {
		if domain.Id == id {
			return domain, true
		}
	}

	return models.Domain{}, false
}

func DomainByTLD(doamins []models.Domain, tld string) (models.Domain, bool) {
	for _, domain := range doamins {
		if domain.Domain == tld {
			return domain, true
		}
	}

	return models.Domain{}, false
}

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
