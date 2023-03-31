package authx

import "github.com/SuperBuker/terraform-provider-dns-he-net/client/auth"

func Creds(auth auth.Auth) (m map[string]string) {
	m = auth.GetAuthForm()
	m["submit"] = "Login!"
	return m
}

func Totp(auth auth.Auth) (map[string]string, error) {
	k, err := auth.GetCode()

	if err != nil {
		return nil, err
	}

	return map[string]string{
		"tfacode": k,
		"submit":  "Submit",
	}, nil
}
