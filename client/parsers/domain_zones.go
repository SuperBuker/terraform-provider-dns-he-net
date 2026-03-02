package parsers

import (
	"strconv"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

// parseDomainZoneNode parses a zone node.
func parseDomainZoneNode(node *html.Node) (models.Zone, error) {
	c := htmlquery.FindOne(node, domainZoneIDQ)
	if c == nil {
		return models.Zone{}, &ErrNotFound{domainZoneIDQ}
	}

	zoneID, err := strconv.Atoi(htmlquery.SelectAttr(c, "value"))

	if err != nil {
		return models.Zone{}, err
	}

	return models.Zone{
		ID:   uint(zoneID),
		Name: htmlquery.SelectAttr(c, "name"),
	}, nil
}

// GetDomainZones returns the zones from the HTML body.
func GetDomainZones(doc *html.Node) ([]models.Zone, error) {
	if table := htmlquery.FindOne(doc, domainZonesTableQ); table == nil {
		return nil, &ErrNotFound{domainZonesTableQ}
	}

	nodes := htmlquery.Find(doc, domainZoneQ)

	if nodes == nil {
		return []models.Zone{}, nil // empty table
	}

	zones := make([]models.Zone, len(nodes))

	for i, node := range nodes {
		zone, err := parseDomainZoneNode(node)
		if err != nil {
			return nil, errParsingNode(domainZoneIDQ, "domainZoneID", err)
		}

		zones[i] = zone
	}

	return zones, nil
}
