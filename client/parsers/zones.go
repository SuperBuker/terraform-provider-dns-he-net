package parsers

import (
	"strconv"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

// parseZoneNode parses a zone node.
func parseZoneNode(node *html.Node) (models.Zone, error) {
	c := htmlquery.FindOne(node, zoneIDQ)
	zoneID, err := strconv.Atoi(htmlquery.SelectAttr(c, "value"))

	if err != nil {
		return models.Zone{}, err
	}

	return models.Zone{
		ID:   uint(zoneID),
		Name: htmlquery.SelectAttr(c, "name"),
	}, nil
}

// GetZones returns the zones from the HTML body.
func GetZones(doc *html.Node) ([]models.Zone, error) {
	if table := htmlquery.FindOne(doc, zonesTableQ); table == nil {
		return nil, &ErrNotFound{zonesTableQ}
	}

	nodes := htmlquery.Find(doc, zoneQ)

	if nodes == nil {
		return []models.Zone{}, nil // empty table
	}

	zones := make([]models.Zone, len(nodes))

	for i, node := range nodes {
		zone, err := parseZoneNode(node)

		if err != nil {
			return nil, errParsingNode(zoneQ, "zoneID", err)
		}

		zones[i] = zone
	}

	return zones, nil
}
