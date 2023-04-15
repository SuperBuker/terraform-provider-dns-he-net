package authx

import "github.com/SuperBuker/terraform-provider-dns-he-net/client/auth"

// Creds returns a map of the values required to log in.
func Creds(auth auth.Auth) (m map[string]string) {
	m = auth.GetAuthForm()
	m["submit"] = "Login!"
	return m
}

// Totp returns a map of the values required to complete the log in.
func Totp(auth auth.Auth) (map[string]string, error) {
	key, err := auth.GetCode()

	if err != nil {
		return nil, err
	}

	return map[string]string{
		"tfacode": key,
		"submit":  "Submit",
	}, nil
}
