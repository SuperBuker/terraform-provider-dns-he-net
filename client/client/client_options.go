package client

import (
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/logging"
	"github.com/go-resty/resty/v2"
)

type Option struct {
	ClientFn     func(*resty.Client) *resty.Client
	AuthClientFn func(*resty.Client) *resty.Client
	LogFn        func(logging.Logger) logging.Logger
}

type Options []Option

// Proc Options
func (opts Options) ApplyClient(client *resty.Client) *resty.Client {
	for _, option := range opts {
		if option.ClientFn != nil {
			client = option.ClientFn(client)
		}
	}
	return client
}

func (opts Options) ApplyAuthClient(client *resty.Client) *resty.Client {
	for _, option := range opts {
		if option.AuthClientFn != nil {
			client = option.AuthClientFn(client)
		}
	}
	return client
}

func (opts Options) ApplyLogger(logger logging.Logger) logging.Logger {
	for _, option := range opts {
		if option.LogFn != nil {
			logger = option.LogFn(logger)
		}
	}
	return logger
}

// WithProxy sets the proxy for the client
func WithProxy(proxy string) Option {
	fn := func(c *resty.Client) *resty.Client {
		return c.SetProxy(proxy)
	}

	return Option{
		ClientFn:     fn,
		AuthClientFn: fn,
	}
}

// WithUserAgent sets the user agent for the client
func WithUserAgent(ua string) Option {
	fn := func(c *resty.Client) *resty.Client {
		return c.SetHeader("User-Agent", ua)
	}

	return Option{
		ClientFn:     fn,
		AuthClientFn: fn,
	}
}

// WithDebug sets the debug mode for the client
func WithDebug() Option {
	fn := func(c *resty.Client) *resty.Client {
		return c.SetDebug(true)
	}

	return Option{
		ClientFn:     fn,
		AuthClientFn: fn,
	}
}
