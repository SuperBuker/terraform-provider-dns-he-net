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
	_ datasource.DataSource              = &networkPrefixes{}
	_ datasource.DataSourceWithConfigure = &networkPrefixes{}
)

// NewNetworkPrefixIndex initialises the NetworkPrefixIndex DataSource.
func NewNetworkPrefixIndex() datasource.DataSource {
	return &networkPrefixes{}
}

// networkPrefixes is the data source implementation.
type networkPrefixes struct {
	client *client.Client
}

// networkPrefixesModel maps the data source schema data.
type networkPrefixesModel struct {
	ID              types.String         `tfsdk:"id"`
	NetworkPrefixes []networkPrefixModel `tfsdk:"network_prefixes"`
}

// Metadata returns the data source type name.
func (networkPrefixes) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_network_prefixes"
}

// Schema defines the schema for the data source.
func (networkPrefixes) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description:         "DNS prefixes in account",
		MarkdownDescription: "DNS prefixes in account",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "dns.he.net account id",
				MarkdownDescription: "dns.he.net account id",
				Computed:            true,
			},
			"network_prefixes": schema.ListNestedAttribute{
				Description:         "Network prefixes list",
				MarkdownDescription: "Network prefixes list",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							Description:         "dns.he.net network prefix id",
							MarkdownDescription: "dns.he.net network prefix id",
							Required:            true,
						},
						"value": schema.StringAttribute{
							Description:         "network prefix value",
							MarkdownDescription: "network prefix value",
							Computed:            true,
						},
						"enabled": schema.BoolAttribute{
							Description:         "network prefix is enabled",
							MarkdownDescription: "network prefix is enabled",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

// Configure adds the provider configured client to the data source.
func (d *networkPrefixes) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if cli, ok := configure(ctx, req, resp); ok {
		d.client = cli
	}
}

// Read refreshes the Terraform state with the latest data.
func (d networkPrefixes) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state networkPrefixesModel

	// Retrieve values from state
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Terraform log
	ctxLog := tflog.SetField(ctx, "account_id", d.client.GetAccount())
	tflog.Debug(ctxLog, "Retrieving network prefixes")

	network_prefixes, err := d.client.GetNetworkPrefixes(ctx)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to fetch network prefixes",
			err.Error(),
		)
		return
	}

	// Terraform log
	ctxLog = tflog.SetField(ctx, "network_prefixes_count", len(network_prefixes))
	tflog.Debug(ctxLog, "Retrieved network prefixes")

	// Map response body to model
	for _, prefix := range network_prefixes {
		prefixState := networkPrefixModel{}

		if err := prefixState.setNetworkPrefix(prefix); err != nil {
			resp.Diagnostics.AddError(
				"Unable to set network prefix",
				err.Error(),
			)
			return
		}

		state.NetworkPrefixes = append(state.NetworkPrefixes, prefixState)
	}

	state.ID = types.StringValue(d.client.GetAccount())

	// Set state
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
