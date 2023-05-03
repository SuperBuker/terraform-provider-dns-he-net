package tfmodels

import (
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"
	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/utils"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// NAPTR maps the data source schema data.
type NAPTR struct {
	ID       types.Int64  `tfsdk:"id"`
	ParentID types.Int64  `tfsdk:"parent_id"`
	Domain   types.String `tfsdk:"domain"`
	TTL      types.Int64  `tfsdk:"ttl"`
	Data     types.String `tfsdk:"data"`
}

func (naptr *NAPTR) SetRecord(recordNAPTR models.NAPTR) error {
	naptr.ID = utils.TypeInt(recordNAPTR.Id)
	naptr.ParentID = types.Int64Value(int64(recordNAPTR.ParentId))
	naptr.Domain = types.StringValue(recordNAPTR.Domain)
	naptr.TTL = types.Int64Value(int64(recordNAPTR.TTL))
	naptr.Data = types.StringValue(recordNAPTR.Data)

	return nil
}

func (naptr NAPTR) GetRecord() (models.NAPTR, error) {
	return models.NAPTR{
		Id:       utils.NativeUInt(naptr.ID),
		ParentId: uint(naptr.ParentID.ValueInt64()),
		Domain:   naptr.Domain.ValueString(),
		TTL:      uint(naptr.TTL.ValueInt64()),
		Data:     naptr.Data.ValueString(),
	}, nil
}
