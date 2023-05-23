package utils_test

import (
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/utils"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"
)

func TestNativeUInt(t *testing.T) {
	x := 0

	assert.Equal(t, uint(x), *utils.NativeUInt(types.Int64Value(int64(x))))
	assert.Nil(t, utils.NativeUInt(types.Int64Null()))
	assert.Nil(t, utils.NativeUInt(types.Int64Unknown()))
}

func TestNativeUInt16(t *testing.T) {
	x := uint16(0)

	assert.Equal(t, x, *utils.NativeUInt16(types.Int64Value(int64(x))))
	assert.Nil(t, utils.NativeUInt16(types.Int64Null()))
	assert.Nil(t, utils.NativeUInt16(types.Int64Unknown()))
}
