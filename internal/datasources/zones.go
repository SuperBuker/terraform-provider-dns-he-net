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
	_ datasource.DataSource              = &domains{}
	_ datasource.DataSourceWithConfigure = &domains{}
)

// NewDomainIndex initialises the DomainIndex DataSource.
func NewDomainIndex() datasource.DataSource {
	return &domains{}
}

// domains is the data source implementation.
type domains struct {
	client *client.Client
}

// domainModel maps the data source schema data.
type domainsModel struct {
	ID      types.String  `tfsdk:"id"`
	Domains []domainModel `tfsdk:"domains"`
}

// Metadata returns the data source type name.
func (domains) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_domains" // TODO: maybe rename
}

// Schema defines the schema for the data source.
func (domains) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "dns.he.net account id",
				MarkdownDescription: "dns.he.net account id",
				Computed:            true,
			},
			"domains": schema.ListNestedAttribute{
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
						"domain": schema.StringAttribute{
							Description:         "zone root domain name",
							MarkdownDescription: "zone root domain name",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

// Configure adds the provider configured client to the data source.
func (d *domains) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	cli, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"unable to configure client",
			"client casting failed",
		)
		return
	}

	d.client = cli
}

// Read refreshes the Terraform state with the latest data.
func (d domains) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state domainsModel

	// Retrieve values from state
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Terraform log
	ctxLog := tflog.SetField(ctx, "account_id", d.client.GetAccount())
	tflog.Debug(ctxLog, "Retrieving root domains")

	domains, err := d.client.GetDomains(ctx)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to fetch root domains",
			err.Error(),
		)
		return
	}

	// Terraform log
	ctxLog = tflog.SetField(ctx, "zones_count", len(domains))
	tflog.Debug(ctxLog, "Retrieved root domains")

	// Map response body to model
	for _, domain := range domains {
		domainState := domainModel{}

		if err := domainState.setDomain(domain); err != nil {
			resp.Diagnostics.AddError(
				"Unable to set root domain",
				err.Error(),
			)
			return
		}

		state.Domains = append(state.Domains, domainState)
	}

	state.ID = types.StringValue(d.client.GetAccount())

	// Set state
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
