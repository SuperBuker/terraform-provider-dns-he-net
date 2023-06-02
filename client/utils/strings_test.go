package utils_test

import (
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/utils"
	"github.com/stretchr/testify/assert"
)

func TestSplitByLen(t *testing.T) {
	matrix := []struct {
		inputString string
		inputLen    int
		result      []string
	}{
		{
			inputString: "1234567890",
			inputLen:    3,
			result:      []string{"123", "456", "789", "0"},
		},
		{
			inputString: "1234567890",
			inputLen:    1,
			result:      []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"},
		},
		{
			inputString: "1234567890",
			inputLen:    10,
			result:      []string{"1234567890"},
		},
		{
			inputString: "1234567890",
			inputLen:    11,
			result:      []string{"1234567890"},
		},
		{
			inputString: "1234567890",
			inputLen:    0,
			result:      nil,
		},
	}

	for _, m := range matrix {
		assert.Equal(t, m.result, utils.SplitByLen(m.inputString, m.inputLen), "inputs: %s, %d", m.inputString, m.inputLen)
	}
}
