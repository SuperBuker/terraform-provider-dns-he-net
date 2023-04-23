package auth

import (
	"net/http"
	"os"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/utils"
	"github.com/kirsle/configdir"
)

type AuthStore int8

const (
	Dummy AuthStore = iota
	Simple
	Encrypted
)

func storeSelector(cs AuthStore) authStore {
	switch cs {
	case Simple:
		return simpleStore()
	case Encrypted:
		return encryptedStore()
	default:
		return dummyStore()
	}
}

type authStore struct {
	Load func(a *Auth) (string, []*http.Cookie, error)
	Save func(a *Auth, account string, cookies []*http.Cookie) error
}

// Load and Save functions for a dummy cookie store.
// Load() returns a not implemented error, Save() skips the operation.
func dummyStore() authStore {
	return authStore{
		Load: func(a *Auth) (string, []*http.Cookie, error) {
			return "", nil, &utils.ErrNotImplemented{}
		},
		Save: func(a *Auth, account string, cookies []*http.Cookie) error {
			return nil
		},
	}
}

// Load and Save functions for a simple file-based cookie store.
func simpleStore() authStore {
	return authStore{
		Load: func(a *Auth) (string, []*http.Cookie, error) {
			bytes, err := os.ReadFile(configFilePath(a, Simple))
			if err != nil {
				return "", nil, &ErrFileIO{err}
			}

			return deserialise(bytes)
		},
		Save: func(a *Auth, account string, cookies []*http.Cookie) error {
			// Ensure it exists.
			if err := configdir.MakePath(configPath); err != nil {
				return &ErrFileIO{err}
			}

			data, err := serialise(account, cookies)
			if err != nil {
				return err // Returns custom error
			}

			if err = os.WriteFile(configFilePath(a, Simple), data, 0644); err != nil {
				return &ErrFileIO{err}
			}

			return nil
		},
	}
}

// Load and Save functions for a encrypted and checksum validated cookie store.
func encryptedStore() authStore {
	return authStore{
		Load: func(a *Auth) (string, []*http.Cookie, error) {
			cipherData, err := os.ReadFile(configFilePath(a, Encrypted))
			if err != nil {
				return "", nil, &ErrFileIO{err}
			}

			sumData, err := decrypt(a, cipherData)
			if err != nil {
				return "", nil, err // Returns custom error
			}

			bytes, err := extractChecksum(sumData)
			if err != nil {
				return "", nil, err // Returns custom error
			}

			return deserialise(bytes)
		},
		Save: func(a *Auth, account string, cookies []*http.Cookie) error {
			// Ensure it exists.
			if err := configdir.MakePath(configPath); err != nil {
				return &ErrFileIO{err}
			}

			data, err := serialise(account, cookies)
			if err != nil {
				return err // Returns custom error
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
