package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApplyToSlice(t *testing.T) {

	assert.Equal(t,
		[]int{1, 2, 3},
		ApplyToSlice(
			func(i int) int {
				return i + 1
			},
			[]int{0, 1, 2},
		),
	)

	assert.Equal(t,
		[]uint{1, 2, 3},
		ApplyToSlice(
			func(i uint) uint {
				return i + 1
			},
			[]uint{0, 1, 2},
		),
	)

	assert.Equal(t,
		[]string{`"1"`, `"2"`, `"3"`},
		ApplyToSlice(
			func(s string) string {
				return `"` + s + `"`
			},
			[]string{"1", "2", "3"},
		),
	)

}
