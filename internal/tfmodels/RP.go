package tfmodels

import (
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"
	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/utils"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// RP maps the data source schema data.
type RP struct {
	ID       types.Int64  `tfsdk:"id"`
	ParentID types.Int64  `tfsdk:"parent_id"`
	Domain   types.String `tfsdk:"domain"`
	TTL      types.Int64  `tfsdk:"ttl"`
	Data     types.String `tfsdk:"data"`
}

func (rp *RP) SetRecord(recordRP models.RP) error {
	rp.ID = utils.TypeInt(recordRP.Id)
	rp.ParentID = types.Int64Value(int64(recordRP.ParentId))
	rp.Domain = types.StringValue(recordRP.Domain)
	rp.TTL = types.Int64Value(int64(recordRP.TTL))
	rp.Data = types.StringValue(recordRP.Data)

	return nil
}

func (rp RP) GetRecord() (models.RP, error) {
	return models.RP{
		Id:       utils.NativeUInt(rp.ID),
		ParentId: uint(rp.ParentID.ValueInt64()),
		Domain:   rp.Domain.ValueString(),
		TTL:      uint(rp.TTL.ValueInt64()),
		Data:     rp.Data.ValueString(),
	}, nil
}
