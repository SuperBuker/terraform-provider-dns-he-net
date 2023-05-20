package parsers

import (
	"strings"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

// ParseStatus returns the dns status from the HTML body.
func ParseStatus(doc *html.Node) string {
	node := htmlquery.FindOne(doc, statusQ)

	if node != nil {
		return strings.TrimSpace(htmlquery.InnerText(node))
	}

	return ""
}
