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
	_ datasource.DataSource              = &domainZone{}
	_ datasource.DataSourceWithConfigure = &domainZone{}
)

// NewDomainZone initialises the Domain DataSource.
func NewDomainZone() datasource.DataSource {
	return &domainZone{}
}

// domainZone is the data source implementation.
type domainZone struct {
	client *client.Client
}

// domainZoneModel maps the data source schema data.
type domainZoneModel struct {
	ZoneID types.Int64  `tfsdk:"zone_id"`
	Name   types.String `tfsdk:"name"`
}

func (d *domainZoneModel) setZone(zone models.Zone) error {
	d.ZoneID = types.Int64Value(int64(zone.ID))
	d.Name = types.StringValue(zone.Name)

	return nil
}

/*func (d *domainZoneModel) getZone() (models.Zone, error) {
	return models.Zone{
		ID:   uint(d.ID.ValueInt64()),
		Name: d.Name.ValueString(),
	}, nil
}*/

// Metadata returns the data source type name.
func (domainZone) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_domain_zone"
}

// Schema defines the schema for the data source.
func (domainZone) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description:         "DNS domain zone",
		MarkdownDescription: "DNS domain zone",
		Attributes: map[string]schema.Attribute{
			"zone_id": schema.Int64Attribute{
				Description:         "dns.he.net domain zone id",
				MarkdownDescription: "dns.he.net domain zone id",
				Required:            true,
			},
			"name": schema.StringAttribute{
				Description:         "domain zone name",
				MarkdownDescription: "domain zone name",
				Computed:            true,
			},
		},
	}
}

// Configure adds the provider configured client to the data source.
func (d *domainZone) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if cli, ok := configure(ctx, req, resp); ok {
		d.client = cli
	}
}

// Read refreshes the Terraform state with the latest data.
func (d domainZone) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state domainZoneModel

	// Retrieve values from state
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Terraform log
	ctxLog := tflog.SetField(ctx, "account_id", d.client.GetAccount())
	tflog.Debug(ctxLog, "Retrieving domain zones")

	domains, err := d.client.GetDomainZones(ctx)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to fetch domain zones",
			err.Error(),
		)
		return
	}

	// Terraform log
	ctxLog = tflog.SetField(ctx, "domain_zones_count", len(domains))
	tflog.Debug(ctxLog, "Retrieved domain zones")

	domain, ok := filters.ZoneById(domains, uint(state.ZoneID.ValueInt64()))
	if !ok {
		resp.Diagnostics.AddError(
			"Unable to find domain zone",
			fmt.Sprintf("domain zone ID:%q doesn't exist", state.ZoneID.String()),
		)
		return
	}

	if err := state.setZone(domain); err != nil {
		resp.Diagnostics.AddError(
			"Unable to set domain zone",
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
