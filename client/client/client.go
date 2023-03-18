package client

import (
	"context"
	"errors"
	"time"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/auth"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/parsers"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/status"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/utils"

	"github.com/go-resty/resty/v2"
)

type Client struct {
	auth    auth.Auth
	client  *resty.Client
	account string
	status  auth.Status
}

func NewClient(ctx context.Context, authAuth auth.Auth) (*Client, error) {
	client := newClient(ctx, authAuth)

	// Manually trigger authentication
	if cookies, err := client.autheticate(ctx); err == nil {
		client.client.SetCookies(cookies)
		return client, nil
	} else {
		return nil, err
	}
}

func newClient(ctx context.Context, authAuth auth.Auth) *Client {
	client := &Client{
		auth:   authAuth,
		client: resty.New(),
	}

	// Handle authentication
	client.client.OnBeforeRequest(func(c *resty.Client, req *resty.Request) error {
		var hasCookies bool

		for _, cookie := range c.Cookies {
			if cookie.Expires.Before(time.Now()) {
				cookie.MaxAge = 0
			} else if !hasCookies {
				hasCookies = true
			}
		}

		if hasCookies && client.status == auth.Ok {
			// pass
		} else if cookies, err := client.autheticate(req.Context()); err == nil {
			c.SetCookies(cookies)
		} else {
			return err
		}

		return nil
	})

	// Parse body errors
	client.client.OnAfterResponse(func(c *resty.Client, resp *resty.Response) (err error) {
		if resp.StatusCode() == 200 {
			err = status.Check(resp.Body())

			// Update client status
			if err == nil {
				// pass
			} else if errors.Is(err, &status.ErrNoAuth{}) {
				client.status = auth.NoAuth
			} else if errors.Is(err, &status.ErrOTPAuth{}) {
				client.status = auth.OTP
			}
		}
		return
	})

	// Parse responses
	client.client.OnAfterResponse(func(c *resty.Client, resp *resty.Response) (err error) {
		if resp.StatusCode() != 200 {
			//pass
		} else if !utils.IsNil(resp.Request.Result) {
			switch resp.Request.Result.(type) {
			case *[]models.Domain:
				resp.Request.Result = parsers.GetDomains(resp.Body())
			case *[]models.Record:
				resp.Request.Result, err = parsers.GetRecords(resp.Body())
			}
		}
		return
	})

	return client
}
