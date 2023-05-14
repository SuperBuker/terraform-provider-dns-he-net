package tfmodels

import (
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"
	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/utils"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SPF maps the data source schema data.
type SPF struct {
	ID       types.Int64  `tfsdk:"id"`
	ParentID types.Int64  `tfsdk:"parent_id"`
	Domain   types.String `tfsdk:"domain"`
	TTL      types.Int64  `tfsdk:"ttl"`
	Data     types.String `tfsdk:"data"`
}

func (spf *SPF) SetRecord(recordSPF models.SPF) error {
	spf.ID = utils.TypeInt(recordSPF.Id)
	spf.ParentID = types.Int64Value(int64(recordSPF.ParentId))
	spf.Domain = types.StringValue(recordSPF.Domain)
	spf.TTL = types.Int64Value(int64(recordSPF.TTL))
	spf.Data = types.StringValue(recordSPF.Data)

	return nil
}

func (spf SPF) GetRecord() (models.SPF, error) {
	return models.SPF{
		Id:       utils.NativeUInt(spf.ID),
		ParentId: uint(spf.ParentID.ValueInt64()),
		Domain:   spf.Domain.ValueString(),
		TTL:      uint(spf.TTL.ValueInt64()),
		Data:     spf.Data.ValueString(),
	}, nil
}
