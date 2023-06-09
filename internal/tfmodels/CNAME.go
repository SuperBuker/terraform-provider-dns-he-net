package tfmodels

import (
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"
	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/utils"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// CNAME maps the record schema data.
type CNAME struct {
	ID     types.Int64  `tfsdk:"id"`
	ZoneID types.Int64  `tfsdk:"zone_id"`
	Domain types.String `tfsdk:"domain"`
	TTL    types.Int64  `tfsdk:"ttl"`
	Data   types.String `tfsdk:"data"`
}

func (cname *CNAME) SetRecord(recordCNAME models.CNAME) error {
	cname.ID = utils.TypeInt(recordCNAME.ID)
	cname.ZoneID = types.Int64Value(int64(recordCNAME.ZoneID))
	cname.Domain = types.StringValue(recordCNAME.Domain)
	cname.TTL = types.Int64Value(int64(recordCNAME.TTL))
	cname.Data = types.StringValue(recordCNAME.Data)

	return nil
}

func (cname CNAME) GetRecord() (models.CNAME, error) {
	return models.CNAME{
		ID:     utils.NativeUInt(cname.ID),
		ZoneID: uint(cname.ZoneID.ValueInt64()),
		Domain: cname.Domain.ValueString(),
		TTL:    uint(cname.TTL.ValueInt64()),
		Data:   cname.Data.ValueString(),
	}, nil
}
