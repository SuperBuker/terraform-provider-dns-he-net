package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

/*
This code has been copied from https://gist.github.com/dopey/c69559607800d2f2f90b1b1ed4e550fb
*/

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ01256789"

// GenerateRandomString generates a random string of the specified length using the defined character set.
func GenerateRandomString(n int) string {
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			panic(fmt.Errorf("failed to generate random number: %w", err))
		}
		ret[i] = letters[num.Int64()]
	}

	return string(ret)
}
