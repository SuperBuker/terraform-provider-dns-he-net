package tfmodels

import (
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"
	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/utils"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SSHFP maps the data source schema data.
type SSHFP struct {
	ID     types.Int64  `tfsdk:"id"`
	ZoneID types.Int64  `tfsdk:"zone_id"`
	Domain types.String `tfsdk:"domain"`
	TTL    types.Int64  `tfsdk:"ttl"`
	Data   types.String `tfsdk:"data"`
}

func (sshfp *SSHFP) SetRecord(recordSSHFP models.SSHFP) error {
	sshfp.ID = utils.TypeInt(recordSSHFP.ID)
	sshfp.ZoneID = types.Int64Value(int64(recordSSHFP.ZoneID))
	sshfp.Domain = types.StringValue(recordSSHFP.Domain)
	sshfp.TTL = types.Int64Value(int64(recordSSHFP.TTL))
	sshfp.Data = types.StringValue(recordSSHFP.Data)

	return nil
}

func (sshfp SSHFP) GetRecord() (models.SSHFP, error) {
	return models.SSHFP{
		ID:     utils.NativeUInt(sshfp.ID),
		ZoneID: uint(sshfp.ZoneID.ValueInt64()),
		Domain: sshfp.Domain.ValueString(),
		TTL:    uint(sshfp.TTL.ValueInt64()),
		Data:   sshfp.Data.ValueString(),
	}, nil
}
