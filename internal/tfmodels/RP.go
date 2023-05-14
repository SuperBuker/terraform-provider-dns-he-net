package tfmodels

import (
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"
	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/utils"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// RP maps the data source schema data.
type RP struct {
	ID     types.Int64  `tfsdk:"id"`
	ZoneID types.Int64  `tfsdk:"zone_id"`
	Domain types.String `tfsdk:"domain"`
	TTL    types.Int64  `tfsdk:"ttl"`
	Data   types.String `tfsdk:"data"`
}

func (rp *RP) SetRecord(recordRP models.RP) error {
	rp.ID = utils.TypeInt(recordRP.ID)
	rp.ZoneID = types.Int64Value(int64(recordRP.ZoneID))
	rp.Domain = types.StringValue(recordRP.Domain)
	rp.TTL = types.Int64Value(int64(recordRP.TTL))
	rp.Data = types.StringValue(recordRP.Data)

	return nil
}

func (rp RP) GetRecord() (models.RP, error) {
	return models.RP{
		ID:     utils.NativeUInt(rp.ID),
		ZoneID: uint(rp.ZoneID.ValueInt64()),
		Domain: rp.Domain.ValueString(),
		TTL:    uint(rp.TTL.ValueInt64()),
		Data:   rp.Data.ValueString(),
	}, nil
}
