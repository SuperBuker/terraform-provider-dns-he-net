package parsers

import (
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/auth"
	"golang.org/x/net/html"

	"github.com/antchfx/htmlquery"
)

// loginStatusQuery is the query to find the login status.
type loginStatusQuery struct {
	status auth.Status
	query  string
}

// getLoginStatusTuples returns the login status tuples.
func getLoginStatusTuples() []loginStatusQuery {
	return []loginStatusQuery{
		{auth.Ok, loginOkQ},
		{auth.OTP, loginOtpQ},
		{auth.NoAuth, loginNoAuthQ}, // NOTE: this must be the last one
	}
}

// LoginStatus returns the login status from the HTML body.
func LoginStatus(doc *html.Node) auth.Status {
	if doc == nil {
		return auth.Unknown
	}

	// Note: the order of the tuples is important
	for _, tuple := range getLoginStatusTuples() {
		node := htmlquery.FindOne(doc, tuple.query)
		if node != nil {
			return tuple.status
		}
	}

	return auth.Unknown
}
