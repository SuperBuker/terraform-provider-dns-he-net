package parsers

import (
	"strings"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

// ParseError returns the error message from the HTML body.
func ParseError(doc *html.Node) (issues []string) {
	q := `//div[@id="dns_err"]`
	node := htmlquery.FindOne(doc, q)

	if node != nil {
		issues = strings.Split(htmlquery.InnerText(node), "  ")

		for i, issue := range issues {
			// Issues clean up
			issue = strings.TrimSpace(issue)
			issue = strings.Replace(issue, "\n", " ", -1)
			issues[i] = issue
		}
	}

	return
}
