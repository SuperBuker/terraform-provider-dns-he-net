package datasources

import (
	"context"
	"fmt"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/client"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/client/filters"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &zone{}
	_ datasource.DataSourceWithConfigure = &zone{}
)

// NewZone initialises the Zone DataSource.
func NewZone() datasource.DataSource {
	return &zone{}
}

// zone is the data source implementation.
type zone struct {
	client *client.Client
}

// zoneModel maps the data source schema data.
type zoneModel struct {
	ID   types.Int64  `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

func (d *zoneModel) setZone(zone models.Zone) error {
	d.ID = types.Int64Value(int64(zone.ID))
	d.Name = types.StringValue(zone.Name)

	return nil
}

/*func (d *zoneModel) getZone() (models.Zone, error) {
	return models.Zone{
		ID:   uint(d.ID.ValueInt64()),
		Name: d.Name.ValueString(),
	}, nil
}*/

// Metadata returns the data source type name.
func (zone) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_zone" // TODO: maybe rename
}

// Schema defines the schema for the data source.
func (zone) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Description:         "dns.he.net zone id",
				MarkdownDescription: "dns.he.net zone id",
				Required:            true,
			},
			"name": schema.StringAttribute{
				Description:         "zone name",
				MarkdownDescription: "zone name",
				Computed:            true,
			},
		},
	}
}

// Configure adds the provider configured client to the data source.
func (d *zone) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if cli, ok := configure(ctx, req, resp); ok {
		d.client = cli
	}
}

// Read refreshes the Terraform state with the latest data.
func (d zone) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state zoneModel

	// Retrieve values from state
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Terraform log
	ctxLog := tflog.SetField(ctx, "account_id", d.client.GetAccount())
	tflog.Debug(ctxLog, "Retrieving zones")

	zones, err := d.client.GetZones(ctx)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to fetch zones",
			err.Error(),
		)
		return
	}

	// Terraform log
	ctxLog = tflog.SetField(ctx, "zones_count", len(zones))
	tflog.Debug(ctxLog, "Retrieved zones")

	zone, ok := filters.ZoneById(zones, uint(state.ID.ValueInt64()))
	if !ok {
		resp.Diagnostics.AddError(
			"Unable to find zone",
			fmt.Sprintf("zone ID:%q doesn't exist", state.ID.String()),
		)
		return
	}

	if err := state.setZone(zone); err != nil {
		resp.Diagnostics.AddError(
			"Unable to set zone",
			err.Error(),
		)
		return
	}

	// Set state
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
