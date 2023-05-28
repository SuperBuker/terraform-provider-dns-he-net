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
	_ datasource.DataSource              = &srv{}
	_ datasource.DataSourceWithConfigure = &srv{}
)

// NewSRV initialises the SRV DataSource.
func NewSRV() datasource.DataSource {
	return &srv{}
}

// srv is the data source implementation.
type srv struct {
	client *client.Client
}

// Metadata returns the data source type name.
func (srv) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_srv" // TODO: maybe rename
}

// Schema defines the schema for the data source.
func (srv) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
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
			"priority": schema.Int64Attribute{
				Description:         "DNS record priority",
				MarkdownDescription: "DNS record priority",
				Computed:            true,
			},
			"weight": schema.Int64Attribute{
				Description:         "DNS record weight",
				MarkdownDescription: "DNS record weight",
				Computed:            true,
			},
			"port": schema.Int64Attribute{
				Description:         "DNS record port",
				MarkdownDescription: "DNS record port",
				Computed:            true,
			},
			"target": schema.StringAttribute{
				Description:         "Value of the DNS record: *TODO*",
				MarkdownDescription: "Value of the DNS record: *TODO*",
				Computed:            true,
			},
		},
	}
}

// Configure adds the provider configured client to the data source.
func (d *srv) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if cli, ok := configure(ctx, req, resp); ok {
		d.client = cli
	}
}

// Read refreshes the Terraform state with the latest data.
func (d srv) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state tfmodels.SRV

	// Retrieve values from state
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Retrieves record from dns.he.net, handles logging and errors
	recordX, ok := readRecord(ctx, d.client, state.ID, state.ZoneID, "SRV", resp)
	if !ok {
		return
	}

	recordSRV, ok := recordX.(models.SRV)
	if !ok {
		resp.Diagnostics.AddError(
			"Unable to cast SRV record",
			fmt.Sprintf("unexpected record type %T", recordX),
		)
		return
	}

	if err := state.SetRecord(recordSRV); err != nil {
		resp.Diagnostics.AddError(
			"Unable to set SRV record",
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
