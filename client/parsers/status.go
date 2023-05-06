package parsers

import (
	"strings"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

// ParseStatus returns the status message from the HTML body.
func ParseStatus(doc *html.Node) string {
	q := `//div[@id="dns_status"]`
	node := htmlquery.FindOne(doc, q)

	if node != nil {
		return strings.TrimSpace(htmlquery.InnerText(node))
	}

	return ""
}
