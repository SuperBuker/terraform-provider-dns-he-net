package tfmodels

import (
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"
	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/utils"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// PTR maps the data source schema data.
type PTR struct {
	ID     types.Int64  `tfsdk:"id"`
	ZoneID types.Int64  `tfsdk:"zone_id"`
	Domain types.String `tfsdk:"domain"`
	TTL    types.Int64  `tfsdk:"ttl"`
	Data   types.String `tfsdk:"data"`
}

func (ptr *PTR) SetRecord(recordPTR models.PTR) error {
	ptr.ID = utils.TypeInt(recordPTR.ID)
	ptr.ZoneID = types.Int64Value(int64(recordPTR.ZoneID))
	ptr.Domain = types.StringValue(recordPTR.Domain)
	ptr.TTL = types.Int64Value(int64(recordPTR.TTL))
	ptr.Data = types.StringValue(recordPTR.Data)

	return nil
}

func (ptr PTR) GetRecord() (models.PTR, error) {
	return models.PTR{
		ID:     utils.NativeUInt(ptr.ID),
		ZoneID: uint(ptr.ZoneID.ValueInt64()),
		Domain: ptr.Domain.ValueString(),
		TTL:    uint(ptr.TTL.ValueInt64()),
		Data:   ptr.Data.ValueString(),
	}, nil
}
