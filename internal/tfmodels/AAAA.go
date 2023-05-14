package tfmodels

import (
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"
	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/utils"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// AAAA maps the record schema data.
type AAAA struct {
	ID      types.Int64  `tfsdk:"id"`
	ZoneID  types.Int64  `tfsdk:"zone_id"`
	Domain  types.String `tfsdk:"domain"`
	TTL     types.Int64  `tfsdk:"ttl"`
	Data    types.String `tfsdk:"data"`
	Dynamic types.Bool   `tfsdk:"dynamic"`
}

func (aaaa *AAAA) SetRecord(recordAAAA models.AAAA) error {
	aaaa.ID = utils.TypeInt(recordAAAA.ID)
	aaaa.ZoneID = types.Int64Value(int64(recordAAAA.ZoneID))
	aaaa.Domain = types.StringValue(recordAAAA.Domain)
	aaaa.TTL = types.Int64Value(int64(recordAAAA.TTL))
	aaaa.Data = types.StringValue(recordAAAA.Data)
	aaaa.Dynamic = types.BoolValue(recordAAAA.Dynamic)

	return nil
}

func (aaaa AAAA) GetRecord() (models.AAAA, error) {
	return models.AAAA{
		ID:      utils.NativeUInt(aaaa.ID),
		ZoneID:  uint(aaaa.ZoneID.ValueInt64()),
		Domain:  aaaa.Domain.ValueString(),
		TTL:     uint(aaaa.TTL.ValueInt64()),
		Data:    aaaa.Data.ValueString(),
		Dynamic: aaaa.Dynamic.ValueBool(),
	}, nil
}
