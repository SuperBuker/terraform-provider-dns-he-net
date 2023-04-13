package auth

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
	"path/filepath"
	"strings"
)

func configFilePath(a *Auth, cs CookieStore) string {
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

func initIV(iv []byte) error {
	_, err := io.ReadFull(rand.Reader, iv)
	return err
}

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

func decrypt(a *Auth, cipherData []byte) ([]byte, error) {
	secret := buildSecret(a)
	block, err := aes.NewCipher(secret)
	if err != nil {
		return nil, &ErrFileEncrypt{err}
	}

	if len(cipherData) < aes.BlockSize {
		return nil, &ErrFileEncrypt{errors.New("ciphertext too short")}
	}

	iv := cipherData[:aes.BlockSize]
	data := make([]byte, len(cipherData)-aes.BlockSize)

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(data, cipherData[aes.BlockSize:])

	return data, nil
}

func encrypt(a *Auth, data []byte) ([]byte, error) {
	secret := buildSecret(a)

	block, err := aes.NewCipher(secret)
	if err != nil {
		return nil, &ErrFileEncrypt{err}
	}

	cipherData := make([]byte, aes.BlockSize+len(data))

	iv := cipherData[:aes.BlockSize]
	if err := initIV(iv); err != nil {
		return nil, &ErrFileEncrypt{err}
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherData[aes.BlockSize:], data)

	return cipherData, nil
}

func addChecksum(data []byte) []byte {
	out := make([]byte, sha256.Size+len(data))
	sum := sha256.Sum256(data)
	copy(out, sum[:])
	copy(out[sha256.Size:], data)

	return out
}

func extractChecksum(data []byte) ([]byte, error) {
	out := make([]byte, len(data)-sha256.Size)
	sum := sha256.Sum256(data[sha256.Size:])

	if !bytes.Equal(sum[:], data[:sha256.Size]) {
		return nil, &ErrFileChecksum{}
	}

	copy(out, data[sha256.Size:])

	return out, nil
}
