package parsers

import (
	"strconv"
	"strings"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

// parseNetworkPrefixNode parses a delegated prefix node.
func parseNetworkPrefixNode(node *html.Node) (models.NetworkPrefix, error) {
	var prefixID int
	var isValid bool

	classID := htmlquery.SelectAttr(node, "class")
	if classID != "broken_zone" {
		c := htmlquery.FindOne(node, prefixIDQ)
		if c == nil {
			return models.NetworkPrefix{}, &ErrNotFound{prefixIDQ}
		}

		var err error
		prefixID, err = strconv.Atoi(htmlquery.SelectAttr(c, "value"))
		if err != nil {
			return models.NetworkPrefix{}, err
		}
		isValid = true
	}

	c := htmlquery.FindOne(node, prefixNameQ)
	if c == nil {
		return models.NetworkPrefix{}, &ErrNotFound{prefixNameQ}
	}

	// Retrieve and clean up ipv6 prefix
	prefixValue := htmlquery.InnerText(c)
	prefixValue = strings.Split(prefixValue, " ")[0] // Removes trailing "(Incomplete) if present"
	prefixValue = strings.Trim(prefixValue, "\n ")

	return models.NetworkPrefix{
		ID:      uint(prefixID),
		Value:   prefixValue,
		Enabled: isValid,
	}, nil
}

// GetNetworkPrefixes returns the network prefixes from the HTML body.
func GetNetworkPrefixes(doc *html.Node) ([]models.NetworkPrefix, error) {
	if table := htmlquery.FindOne(doc, prefixesTableQ); table == nil {
		return nil, &ErrNotFound{prefixesTableQ}
	}

	nodes := htmlquery.Find(doc, prefixQ)

	if len(nodes) <= 1 {
		return []models.NetworkPrefix{}, nil // empty table
	}

	nodes = nodes[1:] // skip header
	prefixes := make([]models.NetworkPrefix, len(nodes))

	for i, node := range nodes {
		prefix, err := parseNetworkPrefixNode(node)

		if err != nil {
			return nil, errParsingNode(prefixQ, "prefixID", err)
		}
		prefixes[i] = prefix
	}

	return prefixes, nil
}
