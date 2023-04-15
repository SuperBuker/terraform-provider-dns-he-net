package status

import (
	"errors"
	"strings"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/auth"
)

// fromAuthStatus returns an error asssociated to the auth status.
func fromAuthStatus(status auth.Status) error {
	switch status {
	case auth.NoAuth:
		return &ErrNoAuth{}
	case auth.Ok:
		return nil
	case auth.OTP:
		return &ErrOTPAuth{}
	case auth.Unknown:
		return &ErrUnknownAuth{}
	case auth.Other:
		return &ErrAuth{}
	}

	return nil // Dead code
}

// fromIssue parses the errors in the response and returns them as &ErrHeNet{}.
func fromIssue(issue string) error {
	if len(issue) != 0 {
		issues := strings.Split(issue, "  ") // Two spaces
		if len(issues) == 1 {
			return &ErrHeNet{issue}
		}

		errs := make([]error, len(issues))
		for i, issue := range issues {
			errs[i] = &ErrHeNet{issue}
		}

		return errors.Join(errs...)
	}

	return nil
}
