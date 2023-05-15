// DNS.HE.NET HTTP client.

package client

import (
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/auth"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/client"
)

var NewClient = client.NewClient

var NewAuth = auth.NewAuth

var CookieStore = struct {
	Dummy     auth.AuthStore
	Simple    auth.AuthStore
	Encrypted auth.AuthStore
}{
	auth.Dummy,
	auth.Simple,
	auth.Encrypted,
}

var With = struct {
	Debug     func() client.Option
	Proxy     func(string) client.Option
	UserAgent func(string) client.Option
}{
	client.WithDebug,
	client.WithProxy,
	client.WithUserAgent,
}
