package parsers

import (
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/auth"
	"golang.org/x/net/html"

	"github.com/antchfx/htmlquery"
)

// LoginStatus returns the login status from the HTML body.
func LoginStatus(doc *html.Node) auth.Status {
	if doc == nil {
		return auth.Unknown
	}

	q := `//a[@id="_tlogout"]`
	node := htmlquery.FindOne(doc, q)

	if node != nil {
		return auth.Ok
	}

	q = `//input[@id="tfacode"]`
	node = htmlquery.FindOne(doc, q)

	if node != nil {
		return auth.OTP
	}

	q = `//form[@name="login"]`
	node = htmlquery.FindOne(doc, q)

	if node != nil {
		return auth.NoAuth
	}

	return auth.Unknown
}
