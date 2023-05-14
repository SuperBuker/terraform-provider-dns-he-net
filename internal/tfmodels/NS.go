package tfmodels

import (
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"
	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/utils"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// NS maps the data source schema data.
type NS struct {
	ID       types.Int64  `tfsdk:"id"`
	ParentID types.Int64  `tfsdk:"parent_id"`
	Domain   types.String `tfsdk:"domain"`
	TTL      types.Int64  `tfsdk:"ttl"`
	Data     types.String `tfsdk:"data"`
}

func (ns *NS) SetRecord(recordNS models.NS) error {
	ns.ID = utils.TypeInt(recordNS.Id)
	ns.ParentID = types.Int64Value(int64(recordNS.ZoneID))
	ns.Domain = types.StringValue(recordNS.Domain)
	ns.TTL = types.Int64Value(int64(recordNS.TTL))
	ns.Data = types.StringValue(recordNS.Data)

	return nil
}

func (ns NS) GetRecord() (models.NS, error) {
	return models.NS{
		Id:     utils.NativeUInt(ns.ID),
		ZoneID: uint(ns.ParentID.ValueInt64()),
		Domain: ns.Domain.ValueString(),
		TTL:    uint(ns.TTL.ValueInt64()),
		Data:   ns.Data.ValueString(),
	}, nil
}
