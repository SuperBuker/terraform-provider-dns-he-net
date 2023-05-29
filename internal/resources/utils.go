package resources

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/client"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/client/filters"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Validators //

var ipv4Validator = stringvalidator.RegexMatches(regexp.MustCompile(`^((25[0-5]|(2[0-4]|1\d|[1-9]|)\d)\.?\b){4}$`), "value must be a valid IPv4 address")
var ipv6Validator = stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)$`), "value must be a valid IPv6 address")
var domainValidator = stringvalidator.RegexMatches(regexp.MustCompile(`^([a-zA-Z0-9][a-zA-Z0-9\-]{0,61}[a-zA-Z0-9]\.)+[a-zA-Z]{2,}$`), "value must be a valid domain name")

// Common functions //

func configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) (*client.Client, bool) {
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

func readRecord(ctx context.Context, cli *client.Client, ID types.Int64, zoneID types.Int64, typ string, resp *resource.ReadResponse) (models.RecordX, bool) {
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

func importRecordState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ZoneID & ID and save to attributes

	sIDs := strings.Split(req.ID, "-")
	if len(sIDs) != 2 {
		resp.Diagnostics.AddError(
			"Error importing item",
			`Could not import item, unexpected error: Identifier should have the format "ZoneID-ID"`,
		)
		return
	}

	ids := make([]int64, 2)
	for i, sID := range sIDs {
		id, err := strconv.ParseInt(sID, 10, 64)

		if err != nil {
			resp.Diagnostics.AddError(
				"Error importing item",
				`Could not import item, unexpected error (Identifier should have the format "ZoneID-ID"): `+err.Error(),
			)
			return
		}

		ids[i] = id
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("zone_id"), ids[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), ids[1])...)
}

func validateName(name string) error {
	if len(name) > 255 {
		return fmt.Errorf("name is too long (%d > 255)", len(name))
	}

	return nil
}
