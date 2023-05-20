package parsers

import (
	"strings"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"
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

// GetStatusMessage returns the status message from the HTML body.
func GetStatusMessage(doc *html.Node) (models.StatusMessage, error) {
	if msg := ParseStatus(doc); len(msg) != 0 {
		return models.StatusMessage{
			Data: msg,
		}, nil
	} else {
		return models.StatusMessage{}, &ErrNotFound{statusQ}
	}
}
