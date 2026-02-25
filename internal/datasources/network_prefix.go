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
	_ datasource.DataSource              = &networkPrefix{}
	_ datasource.DataSourceWithConfigure = &networkPrefix{}
)

// NewNetworkPrefix initialises the NetworkPrefix DataSource.
func NewNetworkPrefix() datasource.DataSource {
	return &networkPrefix{}
}

// networkPrefix is the data source implementation.
type networkPrefix struct {
	client *client.Client
}

// networkPrefixModel maps the data source schema data.
type networkPrefixModel struct {
	ID      types.Int64  `tfsdk:"id"`
	Value   types.String `tfsdk:"value"`
	Enabled types.Bool   `tfsdk:"enabled"`
}

func (np *networkPrefixModel) setNetworkPrefix(prefix models.NetworkPrefix) error {
	np.ID = types.Int64Value(int64(prefix.ID))
	np.Value = types.StringValue(prefix.Value)
	np.Enabled = types.BoolValue(prefix.Enabled)

	return nil
}

/*func (np *networkPrefixModel) getNetworkPrefix() (models.NetworkPrefix, error) {
	return models.NetworkPrefix{
		ID:      uint(np.ID.ValueInt64()),
		Value:   np.Value.ValueString(),
		Enabled: np.Enabled.ValueBool(),
	}, nil
}*/

// Metadata returns the data source type name.
func (networkPrefix) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_network_prefix"
}

// Schema defines the schema for the data source.
func (networkPrefix) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description:         "DNS network prefix",
		MarkdownDescription: "DNS network prefix",
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
	}
}

// Configure adds the provider configured client to the data source.
func (d *networkPrefix) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if cli, ok := configure(ctx, req, resp); ok {
		d.client = cli
	}
}

// Read refreshes the Terraform state with the latest data.
func (d networkPrefix) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state networkPrefixModel

	// Retrieve values from state
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Terraform log
	ctxLog := tflog.SetField(ctx, "account_id", d.client.GetAccount())
	tflog.Debug(ctxLog, "Retrieving network prefixes")

	networkPrefixes, err := d.client.GetNetworkPrefixes(ctx)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to fetch network prefixes",
			err.Error(),
		)
		return
	}

	// Terraform log
	ctxLog = tflog.SetField(ctx, "network_prefixes_count", len(networkPrefixes))
	tflog.Debug(ctxLog, "Retrieved network prefixes")

	networkPrefix, ok := filters.NetworkPrefixById(networkPrefixes, uint(state.ID.ValueInt64()))
	if !ok {
		resp.Diagnostics.AddError(
			"Unable to find network prefix",
			fmt.Sprintf("network prefix ID:%q doesn't exist", state.ID.String()),
		)
		return
	}

	if err := state.setNetworkPrefix(networkPrefix); err != nil {
		resp.Diagnostics.AddError(
			"Unable to set network prefix",
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
