package status

import (
	"errors"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/parsers"
	"golang.org/x/net/html"
)

// Check checks all possible errors in the response.
//   - If the user is not fully logged in.
//   - If there are other contained errors.
func Check(doc *html.Node) error {
	status := parsers.LoginStatus(doc)
	issue := parsers.ParseError(doc)

	errS := fromAuthStatus(status)

	errI := fromIssue(issue)

	if errS != nil && errI != nil {
		return errors.Join(errS, errI)
	} else if errS != nil {
		return errS
	}

	return errI
}
