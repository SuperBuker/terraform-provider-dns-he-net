package utils_test

import (
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/utils"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"
)

func TestTypeInt(t *testing.T) {
	x := uint(0)

	assert.Equal(t, types.Int64Value(0), utils.TypeInt(&x))
	assert.Equal(t, types.Int64Null(), utils.TypeInt(nil))
}

func TestTypeInt16(t *testing.T) {
	x := uint16(0)

	assert.Equal(t, types.Int64Value(0), utils.TypeInt16(&x))
	assert.Equal(t, types.Int64Null(), utils.TypeInt16(nil))
}
