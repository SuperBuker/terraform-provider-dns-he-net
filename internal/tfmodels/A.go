package tfmodels

import (
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"
	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/utils"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// A maps the record schema data.
type A struct {
	ID      types.Int64  `tfsdk:"id"`
	ZoneID  types.Int64  `tfsdk:"zone_id"`
	Domain  types.String `tfsdk:"domain"`
	TTL     types.Int64  `tfsdk:"ttl"`
	Data    types.String `tfsdk:"data"`
	Dynamic types.Bool   `tfsdk:"dynamic"`
}

func (a *A) SetRecord(recordA models.A) error {
	a.ID = utils.TypeInt(recordA.ID)
	a.ZoneID = types.Int64Value(int64(recordA.ZoneID))
	a.Domain = types.StringValue(recordA.Domain)
	a.TTL = types.Int64Value(int64(recordA.TTL))
	a.Data = types.StringValue(recordA.Data)
	a.Dynamic = types.BoolValue(recordA.Dynamic)

	return nil
}

func (a A) GetRecord() (models.A, error) {
	return models.A{
		ID:      utils.NativeUInt(a.ID),
		ZoneID:  uint(a.ZoneID.ValueInt64()),
		Domain:  a.Domain.ValueString(),
		TTL:     uint(a.TTL.ValueInt64()),
		Data:    a.Data.ValueString(),
		Dynamic: a.Dynamic.ValueBool(),
	}, nil
}
