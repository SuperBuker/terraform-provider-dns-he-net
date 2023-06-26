package datasources

import (
	"context"
	"fmt"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/client"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"
	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/tfmodels"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &loc{}
	_ datasource.DataSourceWithConfigure = &loc{}
)

// NewLOC initialises the LOC DataSource.
func NewLOC() datasource.DataSource {
	return &loc{}
}

// loc is the data source implementation.
type loc struct {
	client *client.Client
}

// Metadata returns the data source type name.
func (loc) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_loc" // TODO: maybe rename
}

// Schema defines the schema for the data source.
func (loc) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description:         "DNS LOC record",
		MarkdownDescription: "DNS LOC record",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Description:         "dns.he.net record id",
				MarkdownDescription: "dns.he.net record id",
				Required:            true,
			},
			"zone_id": schema.Int64Attribute{
				Description:         "dns.he.net zone id",
				MarkdownDescription: "dns.he.net zone id",
				Required:            true,
			},
			"domain": schema.StringAttribute{
				Description:         "Name of the DNS record",
				MarkdownDescription: "Name of the DNS record",
				Computed:            true,
			},
			"ttl": schema.Int64Attribute{
				Description:         "Time-To-Live of the DNS record",
				MarkdownDescription: "Time-To-Live of the DNS record",
				Computed:            true,
			},
			"data": schema.StringAttribute{
				Description:         "Value of the DNS record: WGS 84 system coordinates",
				MarkdownDescription: "Value of the DNS record: WGS 84 system coordinates",
				Computed:            true,
			},
		},
	}
}

// Configure adds the provider configured client to the data source.
func (d *loc) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if cli, ok := configure(ctx, req, resp); ok {
		d.client = cli
	}
}

// Read refreshes the Terraform state with the latest data.
func (d loc) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state tfmodels.LOC

	// Retrieve values from state
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Retrieves record from dns.he.net, handles logging and errors
	recordX, ok := readRecord(ctx, d.client, state.ID, state.ZoneID, "LOC", resp)
	if !ok {
		return
	}

	recordLOC, ok := recordX.(models.LOC)
	if !ok {
		resp.Diagnostics.AddError(
			"Unable to cast LOC record",
			fmt.Sprintf("unexpected record type %T", recordX),
		)
		return
	}

	if err := state.SetRecord(recordLOC); err != nil {
		resp.Diagnostics.AddError(
			"Unable to set LOC record",
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
