package datasources

import (
	"context"
	"fmt"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/client"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/client/filters"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"
	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/tfmodels"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &rp{}
	_ datasource.DataSourceWithConfigure = &rp{}
)

// NewRP initialises the RP DataSource.
func NewRP() datasource.DataSource {
	return &rp{}
}

// rp is the data source implementation.
type rp struct {
	client *client.Client
}

// Metadata returns the data source type name.
func (rp) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_rp" // TODO: maybe rename
}

// Schema defines the schema for the data source.
func (rp) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
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
			"data": schema.StringAttribute{
				Description:         "Value of the DNS record: *TODO*",
				MarkdownDescription: "Value of the DNS record: *TODO*",
				Computed:            true,
			},
		},
	}
}

// Configure adds the provider configured client to the data source.
func (d *rp) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
func (d rp) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state tfmodels.RP

	// Retrieve values from state
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Terraform log
	ctxLog := tflog.SetField(ctx, "zone_id", state.ZoneID.String())
	tflog.Debug(ctxLog, "Retrieving DNS records")

	records, err := d.client.GetRecords(ctx, uint(state.ZoneID.ValueInt64())) //GetOne(state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to fetch DNS records",
			err.Error(),
		)
		return
	}

	// Terraform log
	ctxLog = tflog.SetField(ctxLog, "record_count", len(records))
	tflog.Debug(ctxLog, "Retrieved DNS records")

	record, ok := filters.RecordById(records, uint(state.ID.ValueInt64()))
	if !ok {
		resp.Diagnostics.AddError(
			"Unable to find RP record",
			fmt.Sprintf("record %q in zone %q doesn't exist", state.ID.String(), state.ZoneID.String()),
		)
		return
	}

	recordX, err := record.ToX()
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to cast RP record",
			err.Error(),
		)
		return
	}

	recordRP, ok := recordX.(models.RP)
	if !ok {
		resp.Diagnostics.AddError(
			"Unable to cast RP record",
			fmt.Sprintf("unexpacted record type %T", recordX),
		)
		return
	}

	if err = state.SetRecord(recordRP); err != nil {
		resp.Diagnostics.AddError(
			"Unable to set RP record",
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
