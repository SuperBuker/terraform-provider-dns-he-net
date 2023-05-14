package parsers

import (
	"strconv"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

// parseZoneNode parses a zone node.
func parseZoneNode(node *html.Node) (models.Zone, error) {
	q := `//td[@style]/img[@name][@value]`

	c := htmlquery.FindOne(node, q)
	recordId, err := strconv.Atoi(htmlquery.SelectAttr(c, "value"))

	if err != nil {
		return models.Zone{}, err
	}

	return models.Zone{
		ID:   uint(recordId),
		Name: htmlquery.SelectAttr(c, "name"),
	}, nil
}

// GetZones returns the zones from the HTML body.
func GetZones(doc *html.Node) ([]models.Zone, error) {
	q := `//table[@id="domains_table"]`

	if table := htmlquery.FindOne(doc, q); table == nil {
		return nil, &ErrNotFound{q}
	}

	q = `//table[@id="domains_table"]/tbody/tr`
	nodes := htmlquery.Find(doc, q)

	if nodes == nil {
		return []models.Zone{}, nil // empty table
	}

	records := make([]models.Zone, len(nodes))

	for i, node := range nodes {
		record, err := parseZoneNode(node)

		if err != nil {
			return nil, &ErrParsing{
				`//table[@id="domains_table"]/tbody/tr // recordId`,
				err,
			}
		}

		records[i] = record
	}

	return records, nil
}
