package tfmodels

import (
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"
	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/utils"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SSHFP maps the data source schema data.
type SSHFP struct {
	ID       types.Int64  `tfsdk:"id"`
	ParentID types.Int64  `tfsdk:"parent_id"`
	Domain   types.String `tfsdk:"domain"`
	TTL      types.Int64  `tfsdk:"ttl"`
	Data     types.String `tfsdk:"data"`
}

func (sshfp *SSHFP) SetRecord(recordSSHFP models.SSHFP) error {
	sshfp.ID = utils.TypeInt(recordSSHFP.Id)
	sshfp.ParentID = types.Int64Value(int64(recordSSHFP.ZoneID))
	sshfp.Domain = types.StringValue(recordSSHFP.Domain)
	sshfp.TTL = types.Int64Value(int64(recordSSHFP.TTL))
	sshfp.Data = types.StringValue(recordSSHFP.Data)

	return nil
}

func (sshfp SSHFP) GetRecord() (models.SSHFP, error) {
	return models.SSHFP{
		Id:     utils.NativeUInt(sshfp.ID),
		ZoneID: uint(sshfp.ParentID.ValueInt64()),
		Domain: sshfp.Domain.ValueString(),
		TTL:    uint(sshfp.TTL.ValueInt64()),
		Data:   sshfp.Data.ValueString(),
	}, nil
}
