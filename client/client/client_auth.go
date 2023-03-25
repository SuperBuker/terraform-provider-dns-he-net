package client

import (
	"context"
	"errors"
	"net/http"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/auth"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/client/authx"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/client/result"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/parsers"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/status"

	"github.com/go-resty/resty/v2"
)

func (c *Client) autheticate(ctx context.Context) ([]*http.Cookie, error) {
	// New client to not to mess with the regular one
	client := resty.New()

	// To simplify handling
	client.OnAfterResponse(func(c *resty.Client, resp *resty.Response) (err error) {
		if resp.StatusCode() == 200 {
			err = result.Init(resp)
		}
		return
	})

	client.OnAfterResponse(func(c *resty.Client, resp *resty.Response) (err error) {
		// Set cookies
		if cookies := resp.Cookies(); len(cookies) > 0 {
			c.SetCookies(cookies)
		}

		// Parse errors
		if resp.StatusCode() == 200 {
			err = status.Check(result.Body(resp))
		}
		return
	})

	resp, err := client.R().
		SetFormData(authx.Creds(c.auth)).
		Post(endpoint)

	if err == nil {
		c.status = auth.Ok

		if accountTmp, err := parsers.GetAccount(result.Body(resp)); err == nil {
			c.account = accountTmp
		}
		return resp.Cookies(), nil
	} else if !errors.Is(err, &status.ErrOTPAuth{}) {
		return nil, err
	}

	c.status = auth.OTP

	form, err := authx.Totp(c.auth)

	if err != nil {
		return nil, err
	}

	resp2, err := client.R().
		SetFormData(form).
		Post(endpoint)

	if err != nil {
		return nil, err
	}

	if accountTmp, err := parsers.GetAccount(result.Body(resp2)); err == nil {
		c.account = accountTmp
	}

	c.status = auth.Ok
	return resp.Cookies(), nil
}
