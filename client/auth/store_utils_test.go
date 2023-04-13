package auth

import (
	"crypto/sha256"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIV(t *testing.T) {
	iv := make([]byte, 16)
	err := initIV(iv)

	require.NoError(t, err)
	assert.NotEqual(t, make([]byte, 16), iv)
}

func TestBuildSecret(t *testing.T) {
	auth, err := NewAuth("1", "2", "3", -1)
	require.NoError(t, err)

	sum := sha256.Sum256([]byte("1:2:3"))
	assert.Equal(t, sum[:24], buildSecret(&auth))
}

func TestEncryption(t *testing.T) {
	// OK
	data := make([]byte, 50)
	err := initIV(data)
	require.NoError(t, err)

	auth, err := NewAuth("1", "2", "3", -1)
	require.NoError(t, err)

	enc, err := encrypt(&auth, data)
	require.NoError(t, err)

	data2, err := decrypt(&auth, enc)
	require.NoError(t, err)

	assert.Equal(t, data, data2)

	// Error
	data2, err = decrypt(&auth, enc[:15])
	require.Error(t, err)
	assert.Equal(t, "file encryption/decryption failed", err.Error())
	assert.Nil(t, data2)
}

func TestChecksum(t *testing.T) {
	// OK
	data := make([]byte, 50)
	err := initIV(data)
	require.NoError(t, err)

	dataSum := addChecksum(data)

	data2, err := extractChecksum(dataSum)
	require.NoError(t, err)

	assert.Equal(t, data, data2)

	// Error
	err = initIV(dataSum)
	require.NoError(t, err)

	data2, err = extractChecksum(dataSum)
	require.Error(t, err)
	assert.ErrorIs(t, err, &ErrFileChecksum{})
	assert.Nil(t, data2)
}
