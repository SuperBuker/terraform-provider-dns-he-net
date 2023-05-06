package parsers

import (
	"strings"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

// ParseError returns the error message from the HTML body.
func ParseError(doc *html.Node) []string {
	q := `//div[@id="dns_err"]`
	node := htmlquery.FindOne(doc, q)

	if node != nil {
		return strings.Split(htmlquery.InnerText(node), "  ")
	}

	return nil
}
