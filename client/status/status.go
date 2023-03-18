package status

import (
	"errors"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/parsers"
)

func Check(data []byte) error {
	status, err := parsers.LoginStatus(data)

	if err != nil {
		return err
	}

	var issue string
	issue, err = parsers.ParseError(data)

	if err != nil {
		return err
	}

	errS := fromAuthStatus(status)

	errI := fromIssue(issue)

	if errS != nil && errI != nil {
		return errors.Join(errS, errI)
	} else if errS != nil {
		return errS
	}

	return errI
}
