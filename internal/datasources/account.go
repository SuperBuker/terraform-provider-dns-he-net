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
	_ datasource.DataSource              = &account{}
	_ datasource.DataSourceWithConfigure = &account{}
)

// NewAccount initialises the Account DataSource.
func NewAccount() datasource.DataSource {
	return &account{}
}

// account is the data source implementation.
type account struct {
	client *client.Client
}

// accountModel maps the data source schema data.
type accountModel struct {
	ID types.String `tfsdk:"id"`
}

// Metadata returns the data source type name.
func (account) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_account" // TODO: maybe rename
}

// Schema defines the schema for the data source.
func (account) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "dns.he.net account id",
				Computed:            true,
			},
		},
	}
}

// Configure adds the provider configured client to the data source.
func (d *account) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if cli, ok := configure(ctx, req, resp); ok {
		d.client = cli
	}
}

// Read refreshes the Terraform state with the latest data.
func (d account) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state accountModel

	// Retrieve values from state
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Terraform log
	tflog.Debug(ctx, "Retrieving account id from cache")

	state.ID = types.StringValue(d.client.GetAccount())

	// Set state
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
