package auth

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"strings"
)

type serialisedStore struct {
	Account string         `json:"account"`
	Cookies []*http.Cookie `json:"cookies"`
}

// configFilePath returns the path to the cookie file for the given user.
// The cookie file is named depending on the aht username and cookie store type.
func configFilePath(a *Auth, cs AuthStore) string {
	var filename string

	switch cs {
	case Simple:
		filename = fmt.Sprintf("cookies-%s.json", a.User)
	case Encrypted:
		filename = fmt.Sprintf("cookies-%s.json.enc", a.User)
	default:
		return ""
	}

	return filepath.Join(configPath, filename)
}

// initIV creates a new nonce, returns error.
func initIV(iv []byte) error {
	_, err := io.ReadFull(rand.Reader, iv)
	return err
}

// buildSecret generates secret from Auth.
func buildSecret(a *Auth) []byte {
	input := []string{
		a.User,
		a.Password,
		a.OTPKey.Secret(),
	}

	output := make([]string, len(input))

	for i, str := range input {
		output[i] = strings.Replace(str, ":", "\\:", -1)
	}

	// Generetates a checksum from the input
	sum := sha256.Sum256([]byte(strings.Join(output, ":")))
	return sum[:24]
}

func serialise(account string, cookies []*http.Cookie) ([]byte, error) {
	data := serialisedStore{
		Account: account,
		Cookies: cookies,
	}

	bytes, err := json.Marshal(data)
	if err != nil {
		return nil, &ErrFileEncoding{err}
	}

	return bytes, nil
}

func deserialise(bytes []byte) (string, []*http.Cookie, error) {
	var data serialisedStore

	if err := json.Unmarshal(bytes, &data); err != nil {
		return "", nil, &ErrFileEncoding{err}
	} else {
		return data.Account, data.Cookies, nil
	}
}

// decrypt the given data using the given Auth, returns new slice and custom
// error.
func decrypt(a *Auth, cipherData []byte) ([]byte, error) {
	secret := buildSecret(a)
	block, err := aes.NewCipher(secret)
	if err != nil {
		return nil, &ErrFileEncryption{err}
	}

	if len(cipherData) < aes.BlockSize {
		return nil, &ErrFileEncryption{errors.New("ciphertext too short")}
	}

	iv := cipherData[:aes.BlockSize]
	data := make([]byte, len(cipherData)-aes.BlockSize)

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(data, cipherData[aes.BlockSize:])

	return data, nil
}

// encrypt the given data using the given Auth, returns new slice and custom
// error.
func encrypt(a *Auth, data []byte) ([]byte, error) {
	secret := buildSecret(a)

	block, err := aes.NewCipher(secret)
	if err != nil {
		return nil, &ErrFileEncryption{err}
	}

	cipherData := make([]byte, aes.BlockSize+len(data))

	iv := cipherData[:aes.BlockSize]
	if err := initIV(iv); err != nil {
		return nil, &ErrFileEncryption{err}
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherData[aes.BlockSize:], data)

	return cipherData, nil
}

// addChecksum prepends a checksum to the given data.
func addChecksum(data []byte) ([]byte, error) {
	if len(data) > 64*1024*1024 {
		return nil, &ErrFileEncryption{errors.New("data too large")}
	}

	out := make([]byte, sha256.Size+len(data))
	sum := sha256.Sum256(data)
	copy(out, sum[:])
	copy(out[sha256.Size:], data)

	return out, nil
}

// extractChecksum given data, returns new slice and error.
func extractChecksum(data []byte) ([]byte, error) {
	out := make([]byte, len(data)-sha256.Size)
	sum := sha256.Sum256(data[sha256.Size:])

	if !bytes.Equal(sum[:], data[:sha256.Size]) {
		return nil, &ErrFileChecksum{}
	}

	copy(out, data[sha256.Size:])

	return out, nil
}
