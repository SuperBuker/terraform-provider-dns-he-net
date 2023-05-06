package client

import (
	"context"
	"errors"
	"time"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/auth"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/client/result"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/logging"
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
	log     logging.Logger
}

// NewClient returns a new client, requires a context and an auth.Auth.
// Autehticates the client against the API.
func NewClient(ctx context.Context, authAuth auth.Auth, log logging.Logger) (*Client, error) {
	client := newClient(ctx, authAuth, log)

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
			fields := logging.Fields{"error": err}
			log.Error(ctx, "error happened when saving cookies", fields)
		}

		return client, nil
	} else {
		return nil, err
	}
}

// newClient returns a new client, handles the go-resty client configuration.
func newClient(ctx context.Context, authAuth auth.Auth, log logging.Logger) *Client {
	client := &Client{
		auth:   authAuth,
		client: resty.New().SetRetryCount(3),
		log:    log,
	}

	// Handle authentication
	client.client.OnBeforeRequest(func(rc *resty.Client, req *resty.Request) error {
		var hasCookies bool

		for _, cookie := range rc.Cookies {
			if cookie.Expires.Before(time.Now()) {
				cookie.MaxAge = 0
			} else if !hasCookies {
				hasCookies = true
			}
		}

		if hasCookies && client.status == auth.Ok {
			// pass
		} else if cookies, err := client.autheticate(req.Context()); err == nil {
			if len(rc.Cookies) != 0 {
				rc.Cookies = nil
				client.log.Info(req.Context(), "clearing cookies")
			}

			rc.SetCookies(cookies)

			if err := client.auth.Save(client.account, cookies); err != nil {
				fields := logging.Fields{"error": err}
				log.Error(ctx, "error happened when saving cookies", fields)
			}
		} else {
			return err
		}

		return nil
	})

	// Parse html
	client.client.OnAfterResponse(func(_ *resty.Client, resp *resty.Response) (err error) {
		if resp.StatusCode() == 200 {
			err = result.Init(resp)
		}
		return
	})

	// Parse body errors
	client.client.OnAfterResponse(func(_ *resty.Client, resp *resty.Response) (err error) {
		if resp.StatusCode() == 200 {
			var msg string
			msg, err = status.Check(result.Body(resp))

			if len(msg) != 0 {
				fields := logging.Fields{"status": msg}
				client.log.Info(resp.Request.Context(), "api message", fields)
			}

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
	client.client.OnAfterResponse(func(_ *resty.Client, resp *resty.Response) (err error) {
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
		func(_ *resty.Response, err error) bool {
			return errors.Is(err, &status.ErrNoAuth{})
		},
	)

	return client
}

// GetAccount returns the account ID.
func (c *Client) GetAccount() string {
	return c.account
}
