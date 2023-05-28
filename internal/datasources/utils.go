package datasources

import (
	"context"
	"fmt"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/client"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/client/filters"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Common functions //

func configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) (*client.Client, bool) {
	if req.ProviderData == nil {
		return nil, false
	}

	cli, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"unable to configure client",
			"client casting failed",
		)
		return nil, false
	}

	return cli, true
}

func readRecord(ctx context.Context, cli *client.Client, ID types.Int64, zoneID types.Int64, typ string, resp *datasource.ReadResponse) (models.RecordX, bool) {
	// Terraform log
	ctxLog := tflog.SetField(ctx, "zone_id", zoneID.String())
	tflog.Debug(ctxLog, "Retrieving DNS records")

	records, err := cli.GetRecords(ctx, uint(zoneID.ValueInt64())) //GetOne(state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to fetch DNS records",
			err.Error(),
		)
		return nil, false
	}

	// Terraform log
	ctxLog = tflog.SetField(ctxLog, "record_count", len(records))
	tflog.Debug(ctxLog, "Retrieved DNS records")

	record, ok := filters.RecordById(records, uint(ID.ValueInt64()))
	if !ok {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to find %s record", typ),
			fmt.Sprintf("record %q in zone %q doesn't exist", ID.String(), zoneID.String()),
		)
		return nil, false
	}

	recordX, err := record.ToX()
	if err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to cast %s record", typ),
			err.Error(),
		)
		return nil, false
	}

	return recordX, true
}
