package utils_test

import (
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/utils"
	"github.com/stretchr/testify/assert"
)

func TestNil(t *testing.T) {
	var i *int

	matrix := []struct {
		input  interface{}
		result bool
	}{
		{
			input:  nil,
			result: true,
		},
		{
			input:  i,
			result: true,
		},
		{
			input:  "",
			result: false,
		},
		{
			input:  0,
			result: false,
		},
		{
			input:  0.0,
			result: false,
		},
		{
			input:  false,
			result: false,
		},
		{
			input:  []interface{}{},
			result: false,
		},
		{
			input:  map[string]interface{}{},
			result: false,
		},
		{
			input:  []interface{}{},
			result: false,
		},
	}

	for _, m := range matrix {
		assert.Equal(t, m.result, utils.IsNil(m.input), "input: %v", m.input)
	}
}
