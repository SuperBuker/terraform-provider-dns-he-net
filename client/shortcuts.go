package client

import (
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
