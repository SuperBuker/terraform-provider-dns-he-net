package parsers

import (
	"strconv"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

// parseDomainNode parses a domain node.
func parseDomainNode(node *html.Node) (models.Domain, error) {
	q := `//td[@style]/img[@name][@value]`

	c := htmlquery.FindOne(node, q)
	recordId, err := strconv.Atoi(htmlquery.SelectAttr(c, "value"))

	if err != nil {
		return models.Domain{}, err
	}

	return models.Domain{
		Id:     uint(recordId),
		Domain: htmlquery.SelectAttr(c, "name"),
	}, nil
}

// GetDomains returns the domains from the HTML body.
func GetDomains(doc *html.Node) ([]models.Domain, error) {
	q := `//table[@id="domains_table"]`

	if table := htmlquery.FindOne(doc, q); table == nil {
		return nil, &ErrNotFound{q}
	}

	q = `//table[@id="domains_table"]/tbody/tr`
	nodes := htmlquery.Find(doc, q)

	if nodes == nil {
		return []models.Domain{}, nil // empty table
	}

	records := make([]models.Domain, len(nodes))

	for i, node := range nodes {
		record, err := parseDomainNode(node)

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
