package utils

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func TypeInt(i *uint) basetypes.Int64Value {
	if i == nil {
		return types.Int64Null()
	}

	return types.Int64Value(int64(*i))
}

func TypeInt16(i *uint16) basetypes.Int64Value {
	if i == nil {
		return types.Int64Null()
	}

	return types.Int64Value(int64(*i))
}
