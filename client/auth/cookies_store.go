package auth

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/utils"
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

// Load and Save functions for a dummy cookie store.
// Load() returns a not implemented error, Save() skips the operation.
func dummyStore() cookieStore {
	return cookieStore{
		Load: func(a *Auth) ([]*http.Cookie, error) {
			return nil, &utils.ErrNotImplemented{}
		},
		Save: func(a *Auth, cookies []*http.Cookie) error {
			return nil
		},
	}
}

// Load and Save functions for a simple file-based cookie store.
func simpleStore() cookieStore {
	return cookieStore{
		Load: func(a *Auth) ([]*http.Cookie, error) {
			data, err := os.ReadFile(configFilePath(a, Simple))
			if err != nil {
				return nil, &ErrFileIO{err}
			}

			var cookies []*http.Cookie

			if err = json.Unmarshal(data, &cookies); err != nil {
				return nil, &ErrFileEncoding{err}
			} else {
				return cookies, nil
			}
		},
		Save: func(a *Auth, cookies []*http.Cookie) error {
			// Ensure it exists.
			if err := configdir.MakePath(configPath); err != nil {
				return &ErrFileIO{err}
			}

			data, err := json.Marshal(cookies)
			if err != nil {
				return &ErrFileEncoding{err}
			}

			if err = os.WriteFile(configFilePath(a, Simple), data, 0644); err != nil {
				return &ErrFileIO{err}
			}

			return nil
		},
	}
}

// Load and Save functions for a encrypted and checksum validated cookie store.
func encryptedStore() cookieStore {
	return cookieStore{
		Load: func(a *Auth) ([]*http.Cookie, error) {
			cipherData, err := os.ReadFile(configFilePath(a, Encrypted))
			if err != nil {
				return nil, &ErrFileIO{err}
			}

			sumData, err := decrypt(a, cipherData)
			if err != nil {
				return nil, err // Returns custom error
			}

			data, err := extractChecksum(sumData)
			if err != nil {
				return nil, err // Returns custom error
			}

			var cookies []*http.Cookie

			if err = json.Unmarshal(data, &cookies); err != nil {
				return nil, &ErrFileEncoding{err}
			} else {
				return cookies, nil
			}
		},
		Save: func(a *Auth, cookies []*http.Cookie) error {
			// Ensure it exists.
			if err := configdir.MakePath(configPath); err != nil {
				return &ErrFileIO{err}
			}

			data, err := json.Marshal(cookies)
			if err != nil {
				return &ErrFileEncoding{err}
			}

			sumData := addChecksum(data)

			cipherData, err := encrypt(a, sumData)
			if err != nil {
				return err // Returns custom error
			}

			if err = os.WriteFile(configFilePath(a, Encrypted), cipherData, 0644); err != nil {
				return &ErrFileIO{err}
			}

			return nil
		},
	}
}
