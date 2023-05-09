package client

import (
	"context"
	"errors"
	"net/http"
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

// Client middlewares //

func (c *Client) authValidation(rc *resty.Client, req *resty.Request) error {
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
			rc.Cookies = make([]*http.Cookie, 0)
			req.Header.Del("Cookie")
			c.log.Info(req.Context(), "clearing cookies")
		}

		rc.SetCookies(cookies)

		// Error not returned, because it's a minor issue
		if err := c.auth.Save(c.account, cookies); err != nil {
			fields := logging.Fields{"error": err}
			c.log.Error(req.Context(), "error happened when saving cookies", fields)
		}
	} else {
		return err
	}

	return nil
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
	return errors.Is(err, &status.ErrNoAuth{})
}

// Result middlewares //

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

// Debug middlewares //

// debugReqStatus logs the request status
func (c *Client) debugReqStatus(rc *resty.Client, req *resty.Request) error {
	fields := logging.Fields{"status": c.status}
	c.log.Debug(context.Background(), "request client status", fields)

	fields = logging.Fields{"cookies": rc.Cookies}
	c.log.Debug(context.Background(), "request client cookies", fields)

	fields = logging.Fields{"cookies": rc.Header}
	c.log.Debug(context.Background(), "request client headers", fields)

	fields = logging.Fields{"cookies": req.Cookies}
	c.log.Debug(context.Background(), "request cookies", fields)

	fields = logging.Fields{"headers": req.Header}
	c.log.Debug(context.Background(), "request headers", fields)

	if req.RawRequest != nil {
		fields = logging.Fields{"cookies": req.RawRequest.Cookies()}
		c.log.Debug(context.Background(), "raw request cookies", fields)

		fields = logging.Fields{"headers": req.RawRequest.Header}
		c.log.Debug(context.Background(), "raw request headers", fields)
	}

	return nil
}

// debugRespStatus logs the response status
func (c *Client) debugRespStatus(rc *resty.Client, resp *resty.Response) error {
	fields := logging.Fields{"status": c.status}
	c.log.Debug(context.Background(), "response client status", fields)

	fields = logging.Fields{"cookies": rc.Cookies}
	c.log.Debug(context.Background(), "response client cookies", fields)

	fields = logging.Fields{"cookies": rc.Header}
	c.log.Debug(context.Background(), "response client headers", fields)

	fields = logging.Fields{"cookies": resp.Request.Cookies}
	c.log.Debug(context.Background(), "response cookies", fields)

	fields = logging.Fields{"headers": resp.Request.Header}
	c.log.Debug(context.Background(), "response headers", fields)

	if resp.Request.RawRequest != nil {
		fields = logging.Fields{"cookies": resp.Request.RawRequest.Cookies()}
		c.log.Debug(context.Background(), "raw response cookies", fields)

		fields = logging.Fields{"headers": resp.Request.RawRequest.Header}
		c.log.Debug(context.Background(), "raw response headers", fields)
	}

	return nil
}
