package auth

import (
	"fmt"
	"time"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

const otpUrl = "otpauth://totp/dns.he.net:%s?secret=%s&issuer=dns.he.net"

type Auth struct {
	User     string
	Password string
	OTPKey   *otp.Key
}

func NewAuth(user, pass, otpSecret string) (Auth, error) {
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
