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
	_ datasource.DataSource              = &domainZones{}
	_ datasource.DataSourceWithConfigure = &domainZones{}
)

// NewDomainZoneIndex initialises the DomainIndex DataSource.
func NewDomainZoneIndex() datasource.DataSource {
	return &domainZones{}
}

// domainZones is the data source implementation.
type domainZones struct {
	client *client.Client
}

// domainsModel maps the data source schema data.
type domainsModel struct {
	ID          types.String      `tfsdk:"id"`
	Zones []domainZoneModel `tfsdk:"zones"`
}

// Metadata returns the data source type name.
func (domainZones) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_domain_zones"
}

// Schema defines the schema for the data source.
func (domainZones) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description:         "DNS domain zones in account",
		MarkdownDescription: "DNS domain zones in account",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "dns.he.net account id",
				MarkdownDescription: "dns.he.net account id",
				Computed:            true,
			},
			"zones": schema.ListNestedAttribute{
				Description:         "Domain zones list",
				MarkdownDescription: "Domain zones list",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
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
				},
			},
		},
	}
}

// Configure adds the provider configured client to the data source.
func (d *domainZones) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if cli, ok := configure(ctx, req, resp); ok {
		d.client = cli
	}
}

// Read refreshes the Terraform state with the latest data.
func (d domainZones) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state domainsModel

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

	// Map response body to model
	for _, domain := range domains {
		domainState := domainZoneModel{}

		if err := domainState.setZone(domain); err != nil {
			resp.Diagnostics.AddError(
				"Unable to set domain zone",
				err.Error(),
			)
			return
		}

		state.Zones = append(state.Zones, domainState)
	}

	state.ID = types.StringValue(d.client.GetAccount())

	// Set state
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
