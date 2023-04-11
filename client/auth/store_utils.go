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

func configFilePath(a *Auth) string {
	return filepath.Join(configPath, fmt.Sprintf("cookies-%s.json.enc", a.User))
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
		return nil, err
	}

	if len(cipherData) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
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
		return nil, err
	}

	cipherData := make([]byte, aes.BlockSize+len(data))

	iv := cipherData[:aes.BlockSize]
	if err := initIV(iv); err != nil {
		return nil, err
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
		return nil, errors.New("checksum doesn't match")
	}

	copy(out, data[sha256.Size:])

	return out, nil
}
