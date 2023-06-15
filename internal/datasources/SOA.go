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
	_ datasource.DataSource              = &soa{}
	_ datasource.DataSourceWithConfigure = &soa{}
)

// NewSOA initialises the SOA DataSource.
func NewSOA() datasource.DataSource {
	return &soa{}
}

// soa is the data source implementation.
type soa struct {
	client *client.Client
}

// Metadata returns the data source type name.
func (soa) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_soa" // TODO: maybe rename
}

// Schema defines the schema for the data source.
func (soa) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description:         "DNS SOA record",
		MarkdownDescription: "DNS SOA record",
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
			"mname": schema.StringAttribute{
				Description:         "Primary server for this zone",
				MarkdownDescription: "Primary server for this zone",
				Computed:            true,
			},
			"rname": schema.StringAttribute{
				Description:         "Party responsible for the domain",
				MarkdownDescription: "Party responsible for the domain",
				Computed:            true,
			},
			"serial": schema.Int64Attribute{
				Description:         "Numeric representation of changes",
				MarkdownDescription: "Numeric representation of changes",
				Computed:            true,
			},
			"refresh": schema.Int64Attribute{
				Description:         "Time to wait before refreshing SOA record",
				MarkdownDescription: "Time to wait before refreshing SOA record",
				Computed:            true,
			},
			"retry": schema.Int64Attribute{
				Description:         "Time to wait before retrying a failed refresh",
				MarkdownDescription: "Time to wait before retrying a failed refresh",
				Computed:            true,
			},
			"expire": schema.Int64Attribute{
				Description:         "Upper time limit before a zone isn't considered authoritative",
				MarkdownDescription: "Upper time limit before a zone isn't considered authoritative",
				Computed:            true,
			},
		},
	}
}

// Configure adds the provider configured client to the data source.
func (d *soa) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if cli, ok := configure(ctx, req, resp); ok {
		d.client = cli
	}
}

// Read refreshes the Terraform state with the latest data.
func (d soa) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state tfmodels.SOA

	// Retrieve values from state
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Retrieves record from dns.he.net, handles logging and errors
	recordX, ok := readRecord(ctx, d.client, state.ID, state.ZoneID, "SOA", resp)
	if !ok {
		return
	}

	recordSOA, ok := recordX.(models.SOA)
	if !ok {
		resp.Diagnostics.AddError(
			"Unable to cast SOA record",
			fmt.Sprintf("unexpected record type %T", recordX),
		)
		return
	}

	if err := state.SetRecord(recordSOA); err != nil {
		resp.Diagnostics.AddError(
			"Unable to set SOA record",
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
