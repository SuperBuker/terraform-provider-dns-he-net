package tfmodels

import (
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"
	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/utils"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// AFSDB maps the record schema data.
type AFSDB struct {
	ID       types.Int64  `tfsdk:"id"`
	ParentID types.Int64  `tfsdk:"parent_id"`
	Domain   types.String `tfsdk:"domain"`
	TTL      types.Int64  `tfsdk:"ttl"`
	Data     types.String `tfsdk:"data"`
	Dynamic  types.Bool   `tfsdk:"dynamic"`
}

func (afsdb *AFSDB) SetRecord(recordAFSDB models.AFSDB) error {
	afsdb.ID = utils.TypeInt(recordAFSDB.Id)
	afsdb.ParentID = types.Int64Value(int64(recordAFSDB.ParentId))
	afsdb.Domain = types.StringValue(recordAFSDB.Domain)
	afsdb.TTL = types.Int64Value(int64(recordAFSDB.TTL))
	afsdb.Data = types.StringValue(recordAFSDB.Data)
	afsdb.Dynamic = types.BoolValue(recordAFSDB.Dynamic)

	return nil
}

func (afsdb AFSDB) GetRecord() (models.AFSDB, error) {
	return models.AFSDB{
		Id:       utils.NativeUInt(afsdb.ID),
		ParentId: uint(afsdb.ParentID.ValueInt64()),
		Domain:   afsdb.Domain.ValueString(),
		TTL:      uint(afsdb.TTL.ValueInt64()),
		Data:     afsdb.Data.ValueString(),
		Dynamic:  afsdb.Dynamic.ValueBool(),
	}, nil
}