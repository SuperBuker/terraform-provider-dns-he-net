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

type Client struct {
	client *resty.Client
}

func New(cli *resty.Client) Client {
	return Client{
		client: cli,
	}
}

func (c Client) update(ctx context.Context, form map[string]string) (string, error) {
	resp, err := c.client.R().
		SetFormData(form).
		SetContext(ctx).
		Post(endpoint)

	if err != nil {
		return "", &ErrAPI{err}
	}

	return strings.TrimSpace(resp.String()), nil
}

func (c Client) UpdateIP(ctx context.Context, hostname, password, myip string) (bool, error) {
	form := map[string]string{
		"hostname": hostname,
		"password": password,
		"myip":     myip,
	}

	msg, err := c.update(ctx, form)

	if err != nil {
		return false, err
	}

	return processResponse(msg)
}

func (c Client) UpdateTXT(ctx context.Context, hostname, password, txt string) (bool, error) {
	form := map[string]string{
		"hostname": hostname,
		"password": password,
		"txt":      txt,
	}

	msg, err := c.update(ctx, form)

	if err != nil {
		return false, err
	}

	return processResponse(msg)
}

func (c Client) CheckAuth(ctx context.Context, hostname, password string) (bool, error) {
	form := map[string]string{
		"hostname": hostname,
		"password": password,
		"myip":     "a.b.c.d",
	}

	msg, err := c.update(ctx, form)

	if err == nil {
		_, err = processResponse(msg)

		if err == nil {
			return true, nil // Ideally this should return some error...
		} else if errT := new(ErrField); errors.As(err, &errT) {
			return true, nil
		} else if errT := new(ErrAuthFailed); errors.As(err, &errT) {
			return false, nil
		}

	}

	return false, err
}
