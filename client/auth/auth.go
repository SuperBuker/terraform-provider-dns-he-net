package auth

import (
	"fmt"
	"net/http"
	"time"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"

	"github.com/kirsle/configdir"
)

const otpUrl = "otpauth://totp/dns.he.net:%s?secret=%s&issuer=dns.he.net"

var configPath = configdir.LocalConfig("terraform-provider-dns-he-net")

type Auth struct {
	User     string
	Password string
	OTPKey   *otp.Key
	store    cookieStore
}

func NewAuth(user, pass, otpSecret string, storeMode CookieStore) (Auth, error) {
	k := fmt.Sprintf(otpUrl, user, otpSecret)

	key, err := otp.NewKeyFromURL(k)

	if err != nil {
		return Auth{}, err
	}

	// Maybe testing here the code generation is a good idea.

	return Auth{
		User:     user,
		Password: pass,
		OTPKey:   key,
		store:    storeSelector(storeMode),
	}, nil
}

func (a *Auth) GetAuthForm() map[string]string {
	return map[string]string{
		"email": a.User,
		"pass":  a.Password,
	}
}

func (a *Auth) GetCode() (string, error) {
	return totp.GenerateCode(a.OTPKey.Secret(), time.Now())
}

func (a *Auth) LoadCookies() ([]*http.Cookie, error) {
	return a.store.Load(a)
}

func (a *Auth) SaveCookies(cookies []*http.Cookie) error {
	return a.store.Save(a, cookies)
}
