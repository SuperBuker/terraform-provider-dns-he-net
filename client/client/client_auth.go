package client

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/auth"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/client/authx"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/client/result"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/logging"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/parsers"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/status"

	"github.com/go-resty/resty/v2"
)

// autheticate authenticates the client against the API on a separated go-resty
// client, then returns the cookies.
func (c *Client) autheticate(ctx context.Context) ([]*http.Cookie, error) {
	// New client to not to mess with the regular one
	client := resty.New()

	client = c.options.ApplyAuthClient(client)

	// Debug
	//client.OnBeforeRequest(c.debugReqStatus)

	// Debug
	//client.OnAfterResponse(c.debugRespStatus)

	// To simplify handling
	client.OnAfterResponse(initResult)

	client.OnAfterResponse(func(rc *resty.Client, resp *resty.Response) (err error) {
		// Parse errors
		if resp.StatusCode() == 200 {
			err = c.statusCheckLog(resp)
		}
		return
	})

	c.log.Info(ctx, "autheticating, please wait...")

	cookies, err := c.authBasic(ctx, client)

	if err == nil {
		c.status = auth.Ok
		// account already set
		return cookies, nil
	} else if !errors.Is(err, &status.ErrOTPAuth{}) {
		return nil, err
	}

	c.status = auth.OTP

	if c.auth.OTPKey == nil {
		c.log.Fatal(ctx, "OTP key not set")
		return nil, err
	}

	// WIP Must be moved to auth OTP
	for i := 0; i < retries; i++ {
		err = c.authOTP(ctx, client)
		if err == nil {
			c.status = auth.Ok
			// account already set

			return cookies, nil
		} else if !errors.Is(err, &status.ErrOTPAuth{}) {
			// unexpected error
			break
		} else if i < retries-1 {
			// awaiting next OTP window, quick exit on lastest attempt
			fields := logging.Fields{"attempt": i}
			c.log.Info(ctx, "pausing before retrying login", fields)
			time.Sleep(retryDelay)
		}
	}
	return nil, err
}

func (c *Client) authBasic(ctx context.Context, client *resty.Client) ([]*http.Cookie, error) {
	c.log.Debug(ctx, "auth request - basic creds")

	resp, err := client.R().
		SetContext(ctx).
		SetFormData(authx.Creds(c.auth)).
		Post(endpoint)

	if err == nil {
		// Requires parsing response
		c.setAccount(resp)
	} else if !errors.Is(err, &status.ErrOTPAuth{}) {
		return nil, err
	}

	return resp.Cookies(), err
}

func (c *Client) authOTP(ctx context.Context, client *resty.Client) error {
	form, err := authx.Totp(c.auth)

	if err != nil {
		return err
	}

	resp, err := client.R().
		SetContext(ctx).
		SetFormData(form).
		Post(endpoint)

	if err == nil {
		// Requires parsing response
		c.setAccount(resp)
	}

	return err
}

func (c *Client) setAccount(resp *resty.Response) {
	if account, err := parsers.GetAccount(result.Body(resp)); err == nil {
		fields := logging.Fields{"account": account}
		c.log.Debug(resp.Request.Context(), "authentication successful", fields)
		c.account = account
	} else {
		fields := logging.Fields{"error": err}
		c.log.Debug(resp.Request.Context(), "authentication successful, but failed parsing account", fields)
	}
}
