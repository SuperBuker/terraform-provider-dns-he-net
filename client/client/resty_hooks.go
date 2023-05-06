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

// Client hooks //

func (c *Client) authValidation(ctx context.Context) func(*resty.Client, *resty.Request) error {
	return func(rc *resty.Client, req *resty.Request) error {
		var hasCookies bool

		for _, cookie := range rc.Cookies {
			if cookie.Expires.Before(time.Now()) {
				cookie.MaxAge = 0
			} else if !hasCookies {
				hasCookies = true
			}
		}

		if hasCookies && c.status == auth.Ok {
			// pass
		} else if cookies, err := c.autheticate(req.Context()); err == nil {
			if len(rc.Cookies) != 0 {
				rc.Cookies = nil
				c.log.Info(req.Context(), "clearing cookies")
			}

			rc.SetCookies(cookies)

			if err := c.auth.Save(c.account, cookies); err != nil {
				fields := logging.Fields{"error": err}
				c.log.Error(ctx, "error happened when saving cookies", fields)
			}
		} else {
			return err
		}

		return nil
	}
}

func (c *Client) statusCheckLog(resp *resty.Response) error {
	msg, msgErrs, err := status.Check(result.Body(resp))

	if len(msg) != 0 {
		fields := logging.Fields{"status": msg}
		c.log.Info(resp.Request.Context(), "api message", fields)
	}

	for _, msgErr := range msgErrs {
		fields := logging.Fields{"error": msgErr}
		c.log.Error(resp.Request.Context(), "api error message", fields)
	}

	return err
}

func (c *Client) statusCheck(_ *resty.Client, resp *resty.Response) (err error) {
	if resp.StatusCode() == 200 {
		err = c.statusCheckLog(resp)

		// Update client status
		if err == nil {
			// pass
		} else if errors.Is(err, &status.ErrNoAuth{}) {
			c.status = auth.NoAuth
		} else if errors.Is(err, &status.ErrOTPAuth{}) {
			c.status = auth.OTP
		}
	}
	return
}

// retryIfAuth returns true if the request should be retried, includes pause.
func (c *Client) retryIfAuth(resp *resty.Response, err error) (retry bool) {
	retry = errors.Is(err, &status.ErrNoAuth{})

	if c.auth.OTPKey == nil || !retry {
		// pass
	} else if !(resp.Request.Attempt == 1 || resp.Request.Attempt > retries) {
		// quick exit on first and last attempt
		fields := logging.Fields{"attempt": resp.Request.Attempt}
		c.log.Info(resp.Request.Context(), "pausing before retrying request", fields)
		time.Sleep(30 * time.Second)
	}

	return
}

// Result hooks //

// initResult initializes the Result field to enable ResultX
func initResult(_ *resty.Client, resp *resty.Response) (err error) {
	if resp.StatusCode() == 200 {
		err = result.Init(resp)
	}
	return
}

// unwrapResult unwraps the ResultX, parses the body if known type
// and sets the resp.Result
func unwrapResult(_ *resty.Client, resp *resty.Response) (err error) {
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
		default:
			resp.Request.Result = res
		}
	}
	return
}
