package datasources

import (
	"context"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/client"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &zones{}
	_ datasource.DataSourceWithConfigure = &zones{}
)

// NewZoneIndex initialises the ZoneIndex DataSource.
func NewZoneIndex() datasource.DataSource {
	return &zones{}
}

// zones is the data source implementation.
type zones struct {
	client *client.Client
}

// zonesModel maps the data source schema data.
type zonesModel struct {
	ID    types.String `tfsdk:"id"`
	Zones []zoneModel  `tfsdk:"zones"`
}

// Metadata returns the data source type name.
func (zones) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_zones" // TODO: maybe rename
}

// Schema defines the schema for the data source.
func (zones) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description:         "DNS zones in account",
		MarkdownDescription: "DNS zones in account",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "dns.he.net account id",
				MarkdownDescription: "dns.he.net account id",
				Computed:            true,
			},
			"zones": schema.ListNestedAttribute{
				Description:         "Zones list",
				MarkdownDescription: "Zones list",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
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
				},
			},
		},
	}
}

// Configure adds the provider configured client to the data source.
func (d *zones) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if cli, ok := configure(ctx, req, resp); ok {
		d.client = cli
	}
}

// Read refreshes the Terraform state with the latest data.
func (d zones) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state zonesModel

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

	// Map response body to model
	for _, zone := range zones {
		zoneState := zoneModel{}

		if err := zoneState.setZone(zone); err != nil {
			resp.Diagnostics.AddError(
				"Unable to set zone",
				err.Error(),
			)
			return
		}

		state.Zones = append(state.Zones, zoneState)
	}

	state.ID = types.StringValue(d.client.GetAccount())

	// Set state
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
