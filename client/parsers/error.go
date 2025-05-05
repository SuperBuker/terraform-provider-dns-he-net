package parsers

import (
	"strings"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

// ParseError returns the error message from the HTML body.
func ParseError(doc *html.Node) (issues []string) {
	node := htmlquery.FindOne(doc, errorQ)

	if node != nil {
		issues = strings.Split(htmlquery.InnerText(node), "  ")

		for i, issue := range issues {
			// Issues clean up
			issue = strings.TrimSpace(issue)
			issue = strings.ReplaceAll(issue, "\n", " ")
			issues[i] = issue
		}
	}

	return
}
