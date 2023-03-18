package parsers

import (
	"bytes"
	"log"
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

func GetDomains(data []byte) []models.Domain {
	doc, err := htmlquery.Parse(bytes.NewReader(data))

	if err != nil {
		log.Fatal(err)
		return nil // err
	}

	q := `//table[@id="domains_table"]/tbody/tr`
	nodes := htmlquery.Find(doc, q)
	// TODO: Check if node is nil

	records := make([]models.Domain, len(nodes))

	for i, node := range nodes {
		record := parseDomainNode(node)
		records[i] = record
	}

	return records // nil
}
