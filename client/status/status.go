package status

import (
	"errors"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/parsers"
	"golang.org/x/net/html"
)

// Check checks all possible errors in the response.
//   - If the user is not fully logged in.
//   - If there are other contained errors.
func Check(doc *html.Node) (string, error) {
	// Parse status message
	status := parsers.ParseStatus(doc)

	// Parse login status & error message
	authStatus := parsers.LoginStatus(doc)
	issue := parsers.ParseError(doc)

	// Parse to error
	errS := fromAuthStatus(authStatus)
	errI := fromIssue(issue)

	if errS != nil && errI != nil {
		return status, errors.Join(errS, errI)
	} else if errS != nil {
		return status, errS
	}

	return status, errI
}
