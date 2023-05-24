package ddns

import (
	"context"
	"errors"
	"strings"

	"github.com/go-resty/resty/v2"
)

const (
	endpoint = "https://dyn.dns.he.net/nic/update"
)

func update(ctx context.Context, form map[string]string) (string, error) {
	// TODO: Set user-agent.
	resp, err := resty.New().R().
		SetFormData(form).
		SetContext(ctx).
		Post(endpoint)

	if err != nil {
		return "", &ErrAPI{err}
	}

	return strings.TrimSpace(resp.String()), nil
}

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

func UpdateIP(ctx context.Context, hostname, password, myip string) (bool, error) {
	form := map[string]string{
		"hostname": hostname,
		"password": password,
		"myip":     myip,
	}

	msg, err := update(ctx, form)

	if err != nil {
		return false, err
	}

	return processResponse(msg)
}

func UpdateTXT(ctx context.Context, hostname, password, txt string) (bool, error) {
	form := map[string]string{
		"hostname": hostname,
		"password": password,
		"txt":      txt,
	}

	msg, err := update(ctx, form)

	if err != nil {
		return false, err
	}

	return processResponse(msg)
}

func CheckAuth(ctx context.Context, hostname, password string) (bool, error) {
	form := map[string]string{
		"hostname": hostname,
		"password": password,
		"myip":     "a.b.c.d",
	}

	msg, err := update(ctx, form)

	if err == nil {
		_, err = processResponse(msg)

		if err == nil {
			return true, nil
		} else if errT := new(ErrField); errors.As(err, &errT) {
			return true, nil
		} else if errT := new(ErrAuthFailed); errors.As(err, &errT) {
			return false, nil
		}

	}

	return false, err
}
