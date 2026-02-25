package parsers

import (
	"strconv"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

// parseArpaZoneNode parses a delegated prefix node.
func parseArpaZoneNode(node *html.Node) (models.Zone, error) {
	c := htmlquery.FindOne(node, arpaZoneIDQ)
	if c == nil {
		return models.Zone{}, &ErrNotFound{arpaZoneIDQ}
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

// GetArpaZones returns the ARPA zones from the HTML body.
func GetArpaZones(doc *html.Node) ([]models.Zone, error) {
	if table := htmlquery.FindOne(doc, arpaZoneTableQ); table == nil {
		return nil, &ErrNotFound{arpaZoneTableQ}
	}

	nodes := htmlquery.Find(doc, arpaZoneQ)
	if len(nodes) <= 1 {
		return []models.Zone{}, nil // empty table
	}

	nodes = nodes[1:] // skip header
	arpaZones := make([]models.Zone, len(nodes))

	for i, node := range nodes {
		arpaZone, err := parseArpaZoneNode(node)
		if err != nil {
			return nil, errParsingNode(arpaZoneIDQ, "arpaZoneID", err)
		}
		arpaZones[i] = arpaZone
	}

	return arpaZones, nil
}
