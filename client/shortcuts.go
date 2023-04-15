// DNS.HE.NET HTTP client.

package client

import (
	"context"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/auth"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/client"
)

var NewClient = client.NewClient

var NewAuth = auth.NewAuth

var CookieStore = struct {
	Dummy     auth.CookieStore
	Simple    auth.CookieStore
	Encrypted auth.CookieStore
}{
	auth.Dummy,
	auth.Simple,
	auth.Encrypted,
}

func NewClientWithCreds(username, password, secret string, cs auth.CookieStore) (func(context.Context) (*client.Client, error), error) {
	auth, err := auth.NewAuth(username, password, secret, cs)
	if err != nil {
		return nil, err
	}

	return func(ctx context.Context) (*client.Client, error) {
		return client.NewClient(ctx, auth)
	}, nil
}
