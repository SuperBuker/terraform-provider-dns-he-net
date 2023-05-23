package utils

import (
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func NativeUInt(i basetypes.Int64Value) *uint {
	if i.IsUnknown() || i.IsNull() {
		return nil
	}
	n := uint(i.ValueInt64())
	return &n
}

func NativeUInt16(i basetypes.Int64Value) *uint16 {
	if i.IsUnknown() || i.IsNull() {
		return nil
	}
	n := uint16(i.ValueInt64())
	return &n
}
