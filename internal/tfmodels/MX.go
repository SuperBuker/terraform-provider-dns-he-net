package tfmodels

import (
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"
	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/utils"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// MX maps the data source schema data.
type MX struct {
	ID       types.Int64  `tfsdk:"id"`
	ParentID types.Int64  `tfsdk:"parent_id"`
	Domain   types.String `tfsdk:"domain"`
	TTL      types.Int64  `tfsdk:"ttl"`
	Priority types.Int64  `tfsdk:"priority"`
	Data     types.String `tfsdk:"data"`
}

func (mx *MX) SetRecord(recordMX models.MX) error {
	mx.ID = utils.TypeInt(recordMX.Id)
	mx.ParentID = types.Int64Value(int64(recordMX.ZoneID))
	mx.Domain = types.StringValue(recordMX.Domain)
	mx.TTL = types.Int64Value(int64(recordMX.TTL))
	mx.Priority = types.Int64Value(int64(recordMX.Priority))
	mx.Data = types.StringValue(recordMX.Data)

	return nil
}

func (mx MX) GetRecord() (models.MX, error) {
	return models.MX{
		Id:       utils.NativeUInt(mx.ID),
		ZoneID:   uint(mx.ParentID.ValueInt64()),
		Domain:   mx.Domain.ValueString(),
		TTL:      uint(mx.TTL.ValueInt64()),
		Priority: uint16(mx.Priority.ValueInt64()),
		Data:     mx.Data.ValueString(),
	}, nil
}
