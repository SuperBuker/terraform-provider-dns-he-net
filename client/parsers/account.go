package parsers

import (
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

// GetAccount returns the account name from the HTML body.
func GetAccount(doc *html.Node) (string, error) {
	node := htmlquery.FindOne(doc, accountQ)

	if table := htmlquery.FindOne(doc, accountQ); table == nil {
		return "", &ErrNotFound{accountQ}
	}

	return htmlquery.SelectAttr(node, "value"), nil
}
