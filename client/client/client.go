package client

import (
	"context"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/auth"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/logging"

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
	client.client.OnBeforeRequest(client.authValidation(ctx))

	// Parse html
	client.client.OnAfterResponse(initResult)

	// Parse body errors
	client.client.OnAfterResponse(client.statusCheck)

	// Parse responses
	client.client.OnAfterResponse(unwrapResult)

	// Set retry condition
	client.client.AddRetryCondition(client.retryIfAuth)

	return client
}

// GetAccount returns the account ID.
func (c *Client) GetAccount() string {
	return c.account
}
