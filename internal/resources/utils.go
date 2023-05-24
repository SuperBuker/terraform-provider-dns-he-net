package resources

import (
	"context"
	"regexp"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Validators //

var ipv4Validator = stringvalidator.RegexMatches(regexp.MustCompile(`^((25[0-5]|(2[0-4]|1\d|[1-9]|)\d)\.?\b){4}$`), "value must be a valid IPv4 address")
var domainValidator = stringvalidator.RegexMatches(regexp.MustCompile(`^([a-zA-Z0-9][a-zA-Z0-9\-]{0,61}[a-zA-Z0-9]\.)+[a-zA-Z]{2,}$`), "value must be a valid domain name")

// Common functions //

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
