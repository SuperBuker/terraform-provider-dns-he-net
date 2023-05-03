package tfmodels

import (
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"
	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/utils"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SOA maps the data source schema data.
type SOA struct {
	ID       types.Int64  `tfsdk:"id"`
	ParentID types.Int64  `tfsdk:"parent_id"`
	Domain   types.String `tfsdk:"domain"`
	TTL      types.Int64  `tfsdk:"ttl"`
	MName    types.String `tfsdk:"mname"`
	RName    types.String `tfsdk:"rname"`
	Serial   types.Int64  `tfsdk:"serial"`
	Refresh  types.Int64  `tfsdk:"refresh"`
	Retry    types.Int64  `tfsdk:"retry"`
	Expire   types.Int64  `tfsdk:"expire"`
}

func (soa *SOA) SetRecord(recordSOA models.SOA) error {
	soa.ID = utils.TypeInt(recordSOA.Id)
	soa.ParentID = types.Int64Value(int64(recordSOA.ParentId))
	soa.Domain = types.StringValue(recordSOA.Domain)
	soa.TTL = types.Int64Value(int64(recordSOA.TTL))
	soa.MName = types.StringValue(recordSOA.MName)
	soa.RName = types.StringValue(recordSOA.RName)
	soa.Serial = types.Int64Value(int64(recordSOA.Serial))
	soa.Refresh = types.Int64Value(int64(recordSOA.Refresh))
	soa.Retry = types.Int64Value(int64(recordSOA.Retry))
	soa.Expire = types.Int64Value(int64(recordSOA.Expire))

	return nil
}

func (soa SOA) GetRecord() (models.SOA, error) {
	return models.SOA{
		Id:       utils.NativeUInt(soa.ID),
		ParentId: uint(soa.ParentID.ValueInt64()),
		Domain:   soa.Domain.ValueString(),
		TTL:      uint(soa.TTL.ValueInt64()),
		MName:    soa.MName.ValueString(),
		RName:    soa.RName.ValueString(),
		Serial:   uint(soa.Serial.ValueInt64()),
		Refresh:  uint(soa.Refresh.ValueInt64()),
		Retry:    uint(soa.Retry.ValueInt64()),
		Expire:   uint(soa.Expire.ValueInt64()),
	}, nil
}
