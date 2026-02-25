package parsers

import (
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"

	"golang.org/x/net/html"
)

// GetAllZones returns the Domain and ARPA zones from the HTML body.
func GetAllZones(doc *html.Node) ([]models.Zone, error) {
	domainZones, err := GetDomainZones(doc)
	if err != nil {
		return nil, err
	}

	arpaZones, err := GetArpaZones(doc)
	if err != nil {
		return nil, err
	}

	return append(domainZones, arpaZones...), nil
}
