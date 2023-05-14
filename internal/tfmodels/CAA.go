package tfmodels

import (
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"
	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/utils"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// CAA maps the record schema data.
type CAA struct {
	ID       types.Int64  `tfsdk:"id"`
	ParentID types.Int64  `tfsdk:"parent_id"`
	Domain   types.String `tfsdk:"domain"`
	TTL      types.Int64  `tfsdk:"ttl"`
	Data     types.String `tfsdk:"data"`
}

func (caa *CAA) SetRecord(recordCAA models.CAA) error {
	caa.ID = utils.TypeInt(recordCAA.Id)
	caa.ParentID = types.Int64Value(int64(recordCAA.ZoneID))
	caa.Domain = types.StringValue(recordCAA.Domain)
	caa.TTL = types.Int64Value(int64(recordCAA.TTL))
	caa.Data = types.StringValue(recordCAA.Data)

	return nil
}

func (caa CAA) GetRecord() (models.CAA, error) {
	return models.CAA{
		Id:     utils.NativeUInt(caa.ID),
		ZoneID: uint(caa.ParentID.ValueInt64()),
		Domain: caa.Domain.ValueString(),
		TTL:    uint(caa.TTL.ValueInt64()),
		Data:   caa.Data.ValueString(),
	}, nil
}
