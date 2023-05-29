package ddns

import "strings"

func processResponse(msg string) (bool, error) {
	msgT := strings.Split(msg, " ")[0] // Extract first word.

	switch msgT {
	case "good":
		return true, nil
	case "nochg":
		return false, nil
	case "badauth":
		return false, &ErrAuthFailed{}
	case "abuse":
		return false, &ErrAbuse{}
	case "noipv4", "noipv6", "notxt", "badip":
		return false, &ErrField{msgT}
	default: //nolint:goerr113	// This is a generic error.
		return false, &ErrUnknown{msg}
	}
}
