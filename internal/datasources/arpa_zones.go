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
	_ datasource.DataSource              = &arpaZones{}
	_ datasource.DataSourceWithConfigure = &arpaZones{}
)

// NewArpaZoneIndex initialises the ArpaIndex DataSource.
func NewArpaZoneIndex() datasource.DataSource {
	return &arpaZones{}
}

// arpaZones is the data source implementation.
type arpaZones struct {
	client *client.Client
}

// arpasModel maps the data source schema data.
type arpasModel struct {
	ID        types.String    `tfsdk:"id"`
	ArpaZones []arpaZoneModel `tfsdk:"arpa_zones"`
}

// Metadata returns the data source type name.
func (arpaZones) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_arpa_zones"
}

// Schema defines the schema for the data source.
func (arpaZones) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description:         "DNS arpa zones in account",
		MarkdownDescription: "DNS arpa zones in account",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "dns.he.net account id",
				MarkdownDescription: "dns.he.net account id",
				Computed:            true,
			},
			"arpa_zones": schema.ListNestedAttribute{
				Description:         "Arpa zones list",
				MarkdownDescription: "Arpa zones list",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"zone_id": schema.Int64Attribute{
							Description:         "dns.he.net arpa zone id",
							MarkdownDescription: "dns.he.net arpa zone id",
							Required:            true,
						},
						"name": schema.StringAttribute{
							Description:         "arpa zone name",
							MarkdownDescription: "arpa zone name",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

// Configure adds the provider configured client to the data source.
func (d *arpaZones) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if cli, ok := configure(ctx, req, resp); ok {
		d.client = cli
	}
}

// Read refreshes the Terraform state with the latest data.
func (d arpaZones) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state arpasModel

	// Retrieve values from state
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Terraform log
	ctxLog := tflog.SetField(ctx, "account_id", d.client.GetAccount())
	tflog.Debug(ctxLog, "Retrieving arpa zones")

	arpas, err := d.client.GetArpaZones(ctx)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to fetch arpa zones",
			err.Error(),
		)
		return
	}

	// Terraform log
	ctxLog = tflog.SetField(ctx, "arpa_zones_count", len(arpas))
	tflog.Debug(ctxLog, "Retrieved arpa zones")

	// Map response body to model
	for _, arpa := range arpas {
		arpaState := arpaZoneModel{}

		if err := arpaState.setZone(arpa); err != nil {
			resp.Diagnostics.AddError(
				"Unable to set arpa zone",
				err.Error(),
			)
			return
		}

		state.ArpaZones = append(state.ArpaZones, arpaState)
	}

	state.ID = types.StringValue(d.client.GetAccount())

	// Set state
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
