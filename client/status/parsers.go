package status

import (
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/auth"
)

// TODO: extend ErrOTPAuthFailed?
var knownIssues = map[string]error{
	"Incorrect":                         &ErrAuthFailed{"Incorrect"},                                                      // login error
	"The token supplied is invalid.":    &ErrOTPAuthFailed{"The token supplied is invalid."},                              // invalid totp error
	"This token has already been used.": &ErrOTPAuthFailed{"This token has already been used. You may not reuse tokens."}, // reused totp error
	"You may not reuse tokens.":         nil,                                                                              // reused totp error, part 2

}

// fromAuthStatus returns an error asssociated to the auth status.
func fromAuthStatus(status auth.Status) (err error) {
	switch status {
	case auth.NoAuth:
		err = &ErrNoAuth{}
	case auth.Ok:
		// pass
	case auth.OTP:
		err = &ErrMissingOTPAuth{}
	case auth.Unknown:
		err = &ErrUnknownAuth{}
	case auth.Other:
		err = &ErrAuthFailed{}
	}

	return
}

func filterIssues(issues []string) ([]string, []error) {
	idx := 0
	errs := make([]error, 0)
	for _, issue := range issues {
		if err, ok := knownIssues[issue]; ok {
			if err != nil {
				errs = append(errs, err)
			}
		} else {
			issues[idx] = issue
			idx++
		}
	}

	issues = issues[:idx]
	return issues, errs
}

// fromIssue parses the errors in the response and returns them as &ErrHeNet{}.
func fromIssue(issues []string) (errs []error) {
	issues, errs = filterIssues(issues) // In-place filter + catalog

	for _, issue := range issues {
		errs = append(errs, &ErrHeNet{issue})
	}

	return
}
