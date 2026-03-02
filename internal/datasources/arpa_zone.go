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
	_ datasource.DataSource              = &arpaZone{}
	_ datasource.DataSourceWithConfigure = &arpaZone{}
)

// NewArpaZone initialises the arpa DataSource.
func NewArpaZone() datasource.DataSource {
	return &arpaZone{}
}

// arpaZone is the data source implementation.
type arpaZone struct {
	client *client.Client
}

// arpaZoneModel maps the data source schema data.
type arpaZoneModel struct {
	ZoneID types.Int64  `tfsdk:"zone_id"`
	Name   types.String `tfsdk:"name"`
}

func (a *arpaZoneModel) setZone(zone models.Zone) error {
	a.ZoneID = types.Int64Value(int64(zone.ID))
	a.Name = types.StringValue(zone.Name)

	return nil
}

/*func (d *arpaZoneModel) getZone() (models.Zone, error) {
	return models.Zone{
		ID:   uint(d.ID.ValueInt64()),
		Name: d.Name.ValueString(),
	}, nil
}*/

// Metadata returns the data source type name.
func (arpaZone) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_arpa_zone"
}

// Schema defines the schema for the data source.
func (arpaZone) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description:         "DNS arpa zone",
		MarkdownDescription: "DNS ARPA zone",
		Attributes: map[string]schema.Attribute{
			"zone_id": schema.Int64Attribute{
				Description:         "dns.he.net ARPA zone id",
				MarkdownDescription: "dns.he.net ARPA zone id",
				Required:            true,
			},
			"name": schema.StringAttribute{
				Description:         "ARPA zone name",
				MarkdownDescription: "ARPA zone name",
				Computed:            true,
			},
		},
	}
}

// Configure adds the provider configured client to the data source.
func (a *arpaZone) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if cli, ok := configure(ctx, req, resp); ok {
		a.client = cli
	}
}

// Read refreshes the Terraform state with the latest data.
func (a arpaZone) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state arpaZoneModel

	// Retrieve values from state
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Terraform log
	ctxLog := tflog.SetField(ctx, "account_id", a.client.GetAccount())
	tflog.Debug(ctxLog, "Retrieving ARPA zones")

	arpas, err := a.client.GetArpaZones(ctx)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to fetch ARPA zones",
			err.Error(),
		)
		return
	}

	// Terraform log
	ctxLog = tflog.SetField(ctx, "arpa_zones_count", len(arpas))
	tflog.Debug(ctxLog, "Retrieved ARPA zones")

	arpa, ok := filters.ZoneById(arpas, uint(state.ZoneID.ValueInt64()))
	if !ok {
		resp.Diagnostics.AddError(
			"Unable to find ARPA zone",
			fmt.Sprintf("ARPA zone ID:%q doesn't exist", state.ZoneID.String()),
		)
		return
	}

	if err := state.setZone(arpa); err != nil {
		resp.Diagnostics.AddError(
			"Unable to set ARPA zone",
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
