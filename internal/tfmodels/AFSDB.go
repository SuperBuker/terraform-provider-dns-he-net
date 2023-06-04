package tfmodels

import (
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"
	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/utils"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// AFSDB maps the record schema data.
type AFSDB struct {
	ID     types.Int64  `tfsdk:"id"`
	ZoneID types.Int64  `tfsdk:"zone_id"`
	Domain types.String `tfsdk:"domain"`
	TTL    types.Int64  `tfsdk:"ttl"`
	Data   types.String `tfsdk:"data"`
}

func (afsdb *AFSDB) SetRecord(recordAFSDB models.AFSDB) error {
	afsdb.ID = utils.TypeInt(recordAFSDB.ID)
	afsdb.ZoneID = types.Int64Value(int64(recordAFSDB.ZoneID))
	afsdb.Domain = types.StringValue(recordAFSDB.Domain)
	afsdb.TTL = types.Int64Value(int64(recordAFSDB.TTL))
	afsdb.Data = types.StringValue(recordAFSDB.Data)

	return nil
}

func (afsdb AFSDB) GetRecord() (models.AFSDB, error) {
	return models.AFSDB{
		ID:     utils.NativeUInt(afsdb.ID),
		ZoneID: uint(afsdb.ZoneID.ValueInt64()),
		Domain: afsdb.Domain.ValueString(),
		TTL:    uint(afsdb.TTL.ValueInt64()),
		Data:   afsdb.Data.ValueString(),
	}, nil
}
