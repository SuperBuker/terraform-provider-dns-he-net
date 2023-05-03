package tfmodels

import (
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"
	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/utils"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SRV maps the data source schema data.
type SRV struct {
	ID       types.Int64  `tfsdk:"id"`
	ParentID types.Int64  `tfsdk:"parent_id"`
	Domain   types.String `tfsdk:"domain"`
	TTL      types.Int64  `tfsdk:"ttl"`
	Priority types.Int64  `tfsdk:"priority"`
	Weight   types.Int64  `tfsdk:"weight"`
	Port     types.Int64  `tfsdk:"port"`
	Target   types.String `tfsdk:"target"`
}

func (srv *SRV) SetRecord(recordSRV models.SRV) error {
	srv.ID = utils.TypeInt(recordSRV.Id)
	srv.ParentID = types.Int64Value(int64(recordSRV.ParentId))
	srv.Domain = types.StringValue(recordSRV.Domain)
	srv.TTL = types.Int64Value(int64(recordSRV.TTL))
	srv.Priority = types.Int64Value(int64(recordSRV.Priority))
	srv.Weight = types.Int64Value(int64(recordSRV.Weight))
	srv.Port = types.Int64Value(int64(recordSRV.Port))
	srv.Target = types.StringValue(recordSRV.Target)

	return nil
}

func (srv SRV) GetRecord() (models.SRV, error) {
	return models.SRV{
		Id:       utils.NativeUInt(srv.ID),
		ParentId: uint(srv.ParentID.ValueInt64()),
		Domain:   srv.Domain.ValueString(),
		TTL:      uint(srv.TTL.ValueInt64()),
		Priority: uint16(srv.Priority.ValueInt64()),
		Weight:   uint16(srv.Weight.ValueInt64()),
		Port:     uint16(srv.Port.ValueInt64()),
		Target:   srv.Target.ValueString(),
	}, nil
}
