package tfmodels

import (
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"
	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/utils"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ALIAS maps the data source schema data.
type ALIAS struct {
	ID       types.Int64  `tfsdk:"id"`
	ParentID types.Int64  `tfsdk:"parent_id"`
	Domain   types.String `tfsdk:"domain"`
	TTL      types.Int64  `tfsdk:"ttl"`
	Data     types.String `tfsdk:"data"`
}

func (alias *ALIAS) SetRecord(recordALIAS models.ALIAS) error {
	alias.ID = utils.TypeInt(recordALIAS.Id)
	alias.ParentID = types.Int64Value(int64(recordALIAS.ZoneID))
	alias.Domain = types.StringValue(recordALIAS.Domain)
	alias.TTL = types.Int64Value(int64(recordALIAS.TTL))
	alias.Data = types.StringValue(recordALIAS.Data)

	return nil
}

func (alias ALIAS) GetRecord() (models.ALIAS, error) {
	return models.ALIAS{
		Id:     utils.NativeUInt(alias.ID),
		ZoneID: uint(alias.ParentID.ValueInt64()),
		Domain: alias.Domain.ValueString(),
		TTL:    uint(alias.TTL.ValueInt64()),
		Data:   alias.Data.ValueString(),
	}, nil
}
