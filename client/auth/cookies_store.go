package auth

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"

	"github.com/kirsle/configdir"
)

type CookieStore int8

const (
	Dummy CookieStore = iota
	Simple
	Encrypted
)

func storeSelector(cs CookieStore) cookieStore {
	switch cs {
	case Simple:
		return simpleStore()
	case Encrypted:
		return encryptedStore()
	default:
		return dummyStore()
	}
}

type cookieStore struct {
	Load func(a *Auth) ([]*http.Cookie, error)
	Save func(a *Auth, cookies []*http.Cookie) error
}

func dummyStore() cookieStore {
	return cookieStore{
		Load: func(a *Auth) ([]*http.Cookie, error) {
			return nil, errors.New("not implemented")
		},
		Save: func(a *Auth, cookies []*http.Cookie) error {
			return nil
		},
	}
}

func simpleStore() cookieStore {
	return cookieStore{
		Load: func(a *Auth) ([]*http.Cookie, error) {
			data, err := os.ReadFile(configFilePath(a, Simple))
			if err != nil {
				return nil, err
			}

			var cookies []*http.Cookie

			if err = json.Unmarshal(data, &cookies); err != nil {
				return nil, err
			} else {
				// TODO: remove expired cookies
				return cookies, err
			}
		},
		Save: func(a *Auth, cookies []*http.Cookie) error {
			// Ensure it exists.
			if err := configdir.MakePath(configPath); err != nil {
				return err
			}

			data, err := json.Marshal(cookies)
			if err != nil {
				return err
			}

			return os.WriteFile(configFilePath(a, Simple), data, 0644)
		},
	}
}

func encryptedStore() cookieStore {
	return cookieStore{
		Load: func(a *Auth) ([]*http.Cookie, error) {
			cipherData, err := os.ReadFile(configFilePath(a, Encrypted))
			if err != nil {
				return nil, err
			}

			sumData, err := decrypt(a, cipherData)
			if err != nil {
				return nil, err
			}

			data, err := extractChecksum(sumData)
			if err != nil {
				return nil, err
			}

			var cookies []*http.Cookie

			if err = json.Unmarshal(data, &cookies); err != nil {
				return nil, err
			} else {
				// TODO: remove expired cookies
				return cookies, err
			}
		},
		Save: func(a *Auth, cookies []*http.Cookie) error {
			// Ensure it exists.
			if err := configdir.MakePath(configPath); err != nil {
				return err
			}

			data, err := json.Marshal(cookies)
			if err != nil {
				return err
			}

			sumData := addChecksum(data)

			cipherData, err := encrypt(a, sumData)
			if err != nil {
				return err
			}

			return os.WriteFile(configFilePath(a, Encrypted), cipherData, 0644)
		},
	}
}
