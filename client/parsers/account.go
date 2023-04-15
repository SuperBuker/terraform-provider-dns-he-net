package parsers

import (
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

// GetAccount returns the account name from the HTML body.
func GetAccount(doc *html.Node) (string, error) {
	q := `//form[@name="remove_domain"]/input[@name="account"]`
	node := htmlquery.FindOne(doc, q)

	if table := htmlquery.FindOne(doc, q); table == nil {
		return "", &ErrNotFound{q}
	}

	return htmlquery.SelectAttr(node, "value"), nil
}
