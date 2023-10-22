package status

import (
	"errors"
	"sort"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/parsers"
	"golang.org/x/net/html"
)

// Check checks all possible errors in the response.
//   - If the user is not fully logged in.
//   - If there are other contained errors.
//   - If there are error messges in the response.
//   - Sorts errors output by severity
func Check(doc *html.Node) (string, []string, error) {
	// Parse statusMsg message
	statusMsg := parsers.ParseStatus(doc)

	// Parse login status & error message
	authStatus := parsers.LoginStatus(doc)
	issues := parsers.ParseError(doc)

	// Parse to error
	err := fromAuthStatus(authStatus)
	errs := fromIssue(issues)

	// Append to error slice
	if err != nil {
		errs = append(errs, err)
	}

	// Sort by severity
	return statusMsg, issues, errorMerging(errs)
}

func errorMerging(errs []error) error {
	if len(errs) == 0 {
		return nil
	} else if len(errs) == 1 {
		return errs[0]
	}

	sort.SliceStable(errs, func(i, j int) bool {
		// Descending order
		return errorScore(errs[i]) > errorScore(errs[j])
	})

	return errors.Join(errs...)
}
