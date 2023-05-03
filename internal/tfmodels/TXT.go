package tfmodels

import (
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"
	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/utils"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// TXT maps the data source schema data.
type TXT struct {
	ID       types.Int64  `tfsdk:"id"`
	ParentID types.Int64  `tfsdk:"parent_id"`
	Domain   types.String `tfsdk:"domain"`
	TTL      types.Int64  `tfsdk:"ttl"`
	Data     types.String `tfsdk:"data"`
	Dynamic  types.Bool   `tfsdk:"dynamic"`
}

func (txt *TXT) SetRecord(recordTXT models.TXT) error {
	txt.ID = utils.TypeInt(recordTXT.Id)
	txt.ParentID = types.Int64Value(int64(recordTXT.ParentId))
	txt.Domain = types.StringValue(recordTXT.Domain)
	txt.TTL = types.Int64Value(int64(recordTXT.TTL))
	txt.Data = types.StringValue(recordTXT.Data)
	txt.Dynamic = types.BoolValue(recordTXT.Dynamic)

	return nil
}

func (txt TXT) GetRecord() (models.TXT, error) {
	return models.TXT{
		Id:       utils.NativeUInt(txt.ID),
		ParentId: uint(txt.ParentID.ValueInt64()),
		Domain:   txt.Domain.ValueString(),
		TTL:      uint(txt.TTL.ValueInt64()),
		Data:     txt.Data.ValueString(),
		Dynamic:  txt.Dynamic.ValueBool(),
	}, nil
}
