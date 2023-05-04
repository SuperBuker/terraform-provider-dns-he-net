package client

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/auth"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/client/result"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/parsers"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/status"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/utils"

	"github.com/go-resty/resty/v2"
)

// Client is a client for the dns.he.net API.
type Client struct {
	auth    auth.Auth
	client  *resty.Client
	account string
	status  auth.Status
}

// NewClient returns a new client, requires a context and an auth.Auth.
// Autehticates the client against the API.
func NewClient(ctx context.Context, authAuth auth.Auth) (*Client, error) {
	client := newClient(ctx, authAuth)

	if account, cookies, err := authAuth.Load(); err == nil {
		// Load cookies from filestore
		client.client.SetCookies(cookies)
		client.status = auth.Ok
		client.account = account
		return client, nil
	} else if cookies, err = client.autheticate(ctx); err == nil {
		// Manually trigger authentication
		client.client.SetCookies(cookies)

		if err := client.auth.Save(client.account, cookies); err != nil {
			log.Printf("error happened when saving cookies: %v", err)
		}

		return client, nil
	} else {
		return nil, err
	}
}

// newClient returns a new client, handles the go-resty client configuration.
func newClient(ctx context.Context, authAuth auth.Auth) *Client {
	client := &Client{
		auth:   authAuth,
		client: resty.New().SetRetryCount(3),
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
			if len(c.Cookies) != 0 {
				c.Cookies = nil
				log.Printf("clearing cookies")
			}

			c.SetCookies(cookies)

			if err := client.auth.Save(client.account, cookies); err != nil {
				log.Printf("error happened when saving cookies: %v", err)
			}
		} else {
			return err
		}

		return nil
	})

	// Parse html
	client.client.OnAfterResponse(func(c *resty.Client, resp *resty.Response) (err error) {
		if resp.StatusCode() == 200 {
			err = result.Init(resp)
		}
		return
	})

	// Parse body errors
	client.client.OnAfterResponse(func(c *resty.Client, resp *resty.Response) (err error) {
		if resp.StatusCode() == 200 {
			err = status.Check(result.Body(resp))

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
		} else if res := result.Result(resp); !utils.IsNil(res) {
			switch res.(type) {
			case *[]models.Domain:
				body := result.Body(resp)
				resp.Request.Result, err = parsers.GetDomains(body)
			case *[]models.Record:
				body := result.Body(resp)
				resp.Request.Result, err = parsers.GetRecords(body)
			}
		}
		return
	})

	// Set retry condition
	client.client.AddRetryCondition(
		// RetryConditionFunc type is for retry condition function
		// input: non-nil Response OR request execution error
		func(r *resty.Response, err error) bool {
			return errors.Is(err, &status.ErrNoAuth{})
		},
	)

	return client
}

// GetAccount returns the account ID.
func (c *Client) GetAccount() string {
	return c.account
}
