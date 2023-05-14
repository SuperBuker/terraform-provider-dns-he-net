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
	_ datasource.DataSource              = &aaaa{}
	_ datasource.DataSourceWithConfigure = &aaaa{}
)

// NewAAAA initialises the AAAA DataSource.
func NewAAAA() datasource.DataSource {
	return &aaaa{}
}

// aaaa is the data source implementation.
type aaaa struct {
	client *client.Client
}

// Metadata returns the data source type name.
func (aaaa) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_aaaa" // TODO: maybe rename
}

// Schema defines the schema for the data source.
func (aaaa) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Description:         "dns.he.net record id",
				MarkdownDescription: "dns.he.net record id",
				Required:            true,
			},
			"parent_id": schema.Int64Attribute{
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
				Description:         "Value of the DNS record: IPv6 address",
				MarkdownDescription: "Value of the DNS record: IPv6 address",
				Computed:            true,
			},
			"dynamic": schema.BoolAttribute{
				Description:         "Enable DDNS for this record",
				MarkdownDescription: "Enable DDNS for this record",
				Computed:            true,
			},
		},
	}
}

// Configure adds the provider configured client to the data source.
func (d *aaaa) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
func (d aaaa) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state tfmodels.AAAA

	// Retrieve values from state
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Terraform log
	ctxLog := tflog.SetField(ctx, "zone_id", state.ParentID.String())
	tflog.Debug(ctxLog, "Retrieving DNS records")

	records, err := d.client.GetRecords(ctx, uint(state.ParentID.ValueInt64())) //GetOne(state.ID.ValueString())
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
			"Unable to find AAAA record",
			fmt.Sprintf(`record "%s" in domain "%s" doesn't exist`, state.ID.String(), state.ParentID.String()),
		)
		return
	}

	recordX, err := record.ToX()
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to cast AAAA record",
			err.Error(),
		)
		return
	}

	recordAAAA, ok := recordX.(models.AAAA)
	if !ok {
		resp.Diagnostics.AddError(
			"Unable to cast AAAA record",
			fmt.Sprintf("unexpacted record type %T", recordX),
		)
		return
	}

	if err = state.SetRecord(recordAAAA); err != nil {
		resp.Diagnostics.AddError(
			"Unable to set AAAA record",
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
