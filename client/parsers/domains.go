package parsers

import (
	"strconv"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

func parseDomainNode(node *html.Node) (record models.Domain) {
	q := `//td[@style]/img[@name][@value]`

	c := htmlquery.FindOne(node, q)
	recordId, _ := strconv.Atoi(htmlquery.SelectAttr(c, "value")) // WIP
	record.Id = uint(recordId)
	record.Domain = htmlquery.SelectAttr(c, "name")

	return
}

func GetDomains(doc *html.Node) ([]models.Domain, error) {
	q := `//table[@id="domains_table"]`

	if table := htmlquery.FindOne(doc, q); table == nil {
		return nil, &ErrNotFound{q}
	}

	q = `//table[@id="domains_table"]/tbody/tr`
	nodes := htmlquery.Find(doc, q)

	if nodes == nil {
		return []models.Domain{}, nil
	}

	records := make([]models.Domain, len(nodes))

	for i, node := range nodes {
		record := parseDomainNode(node)
		records[i] = record
	}

	return records, nil
}
