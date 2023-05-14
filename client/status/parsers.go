package status

import (
	"errors"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/auth"
)

// fromAuthStatus returns an error asssociated to the auth status.
func fromAuthStatus(status auth.Status) (err error) {
	switch status {
	case auth.NoAuth:
		err = &ErrNoAuth{}
	case auth.Ok:
		// pass
	case auth.OTP:
		err = &ErrOTPAuth{}
	case auth.Unknown:
		err = &ErrUnknownAuth{}
	case auth.Other:
		err = &ErrAuthFailed{}
	}

	return
}

// fromIssue parses the errors in the response and returns them as &ErrHeNet{}.
func fromIssue(issues []string) (err error) {
	if issues == nil {
		//pass
	} else if len(issues) == 1 {
		err = &ErrHeNet{issues[0]}
	} else {
		errs := make([]error, len(issues))
		for i, issue := range issues {
			errs[i] = &ErrHeNet{issue}
		}

		err = errors.Join(errs...)
	}

	return
}
