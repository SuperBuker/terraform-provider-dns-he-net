package tfmodels

import (
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"
	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/utils"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// LOC maps the data source schema data.
type LOC struct {
	ID     types.Int64  `tfsdk:"id"`
	ZoneID types.Int64  `tfsdk:"zone_id"`
	Domain types.String `tfsdk:"domain"`
	TTL    types.Int64  `tfsdk:"ttl"`
	Data   types.String `tfsdk:"data"`
}

func (loc *LOC) SetRecord(recordLOC models.LOC) error {
	loc.ID = utils.TypeInt(recordLOC.ID)
	loc.ZoneID = types.Int64Value(int64(recordLOC.ZoneID))
	loc.Domain = types.StringValue(recordLOC.Domain)
	loc.TTL = types.Int64Value(int64(recordLOC.TTL))
	loc.Data = types.StringValue(recordLOC.Data)

	return nil
}

func (loc LOC) GetRecord() (models.LOC, error) {
	return models.LOC{
		ID:     utils.NativeUInt(loc.ID),
		ZoneID: uint(loc.ZoneID.ValueInt64()),
		Domain: loc.Domain.ValueString(),
		TTL:    uint(loc.TTL.ValueInt64()),
		Data:   loc.Data.ValueString(),
	}, nil
}
