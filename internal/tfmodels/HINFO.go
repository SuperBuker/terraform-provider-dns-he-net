package tfmodels

import (
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"
	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/utils"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// HINFO maps the data source schema data.
type HINFO struct {
	ID       types.Int64  `tfsdk:"id"`
	ParentID types.Int64  `tfsdk:"parent_id"`
	Domain   types.String `tfsdk:"domain"`
	TTL      types.Int64  `tfsdk:"ttl"`
	Data     types.String `tfsdk:"data"`
}

func (hinfo *HINFO) SetRecord(recordHINFO models.HINFO) error {
	hinfo.ID = utils.TypeInt(recordHINFO.Id)
	hinfo.ParentID = types.Int64Value(int64(recordHINFO.ParentId))
	hinfo.Domain = types.StringValue(recordHINFO.Domain)
	hinfo.TTL = types.Int64Value(int64(recordHINFO.TTL))
	hinfo.Data = types.StringValue(recordHINFO.Data)

	return nil
}

func (hinfo HINFO) GetRecord() (models.HINFO, error) {
	return models.HINFO{
		Id:       utils.NativeUInt(hinfo.ID),
		ParentId: uint(hinfo.ParentID.ValueInt64()),
		Domain:   hinfo.Domain.ValueString(),
		TTL:      uint(hinfo.TTL.ValueInt64()),
		Data:     hinfo.Data.ValueString(),
	}, nil
}
