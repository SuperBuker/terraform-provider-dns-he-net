package status

import (
	"errors"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/parsers"
	"golang.org/x/net/html"
)

// Check checks all possible errors in the response.
//   - If the user is not fully logged in.
//   - If there are other contained errors.
//   - If there are error messges in the response.
func Check(doc *html.Node) (string, []string, error) {
	// Parse statusMsg message
	statusMsg := parsers.ParseStatus(doc)

	// Parse login status & error message
	authStatus := parsers.LoginStatus(doc)
	issues := parsers.ParseError(doc)

	// Parse to error
	errS := fromAuthStatus(authStatus)
	errI := fromIssue(issues)

	if errS != nil && errI != nil {
		return statusMsg, issues, errors.Join(errS, errI)
	} else if errS != nil {
		return statusMsg, issues, errS
	}

	return statusMsg, issues, errI
}
