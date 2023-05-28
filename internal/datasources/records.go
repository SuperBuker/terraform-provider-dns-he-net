package datasources

import (
	"context"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/client"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"
	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/utils"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &records{}
	_ datasource.DataSourceWithConfigure = &records{}
)

// NewRecordIndex initialises the RecordIndex DataSource.
func NewRecordIndex() datasource.DataSource {
	return &records{}
}

// records is the data source implementation.
type records struct {
	client *client.Client
}

// recordsModel maps the data source schema data.
type recordsModel struct {
	ID      types.Int64   `tfsdk:"id"`
	Records []recordModel `tfsdk:"records"`
}

type recordModel struct {
	ID         types.Int64  `tfsdk:"id"`
	ZoneID     types.Int64  `tfsdk:"zone_id"`
	Domain     types.String `tfsdk:"domain"`
	RecordType types.String `tfsdk:"record_type"`
	TTL        types.Int64  `tfsdk:"ttl"`
	Priority   types.Int64  `tfsdk:"priority"`
	Data       types.String `tfsdk:"data"`
	Dynamic    types.Bool   `tfsdk:"dynamic"`
	Locked     types.Bool   `tfsdk:"locked"`
}

func (a *recordModel) setRecord(record models.Record) error {
	a.ID = utils.TypeInt(record.ID)
	a.ZoneID = types.Int64Value(int64(record.ZoneID))
	a.Domain = types.StringValue(record.Domain)
	a.RecordType = types.StringValue(record.RecordType)
	a.TTL = types.Int64Value(int64(record.TTL))
	a.Priority = utils.TypeInt16(record.Priority)
	a.Data = types.StringValue(record.Data)
	a.Dynamic = types.BoolValue(record.Dynamic)
	a.Locked = types.BoolValue(record.Locked)

	return nil
}

func (a *recordModel) getRecord() (models.Record, error) {
	return models.Record{
		ID:         utils.NativeUInt(a.ID),
		ZoneID:     uint(a.ZoneID.ValueInt64()),
		Domain:     a.Domain.ValueString(),
		RecordType: a.RecordType.ValueString(),
		TTL:        uint(a.TTL.ValueInt64()),
		Priority:   utils.NativeUInt16(a.Priority),
		Data:       a.Data.ValueString(),
		Dynamic:    a.Dynamic.ValueBool(),
		Locked:     a.Locked.ValueBool(),
	}, nil
}

// Metadata returns the data source type name.
func (records) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_records" // TODO: maybe rename
}

// Schema defines the schema for the data source.
func (records) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Description:         "dns.he.net zone id",
				MarkdownDescription: "dns.he.net zone id",
				Required:            true,
			},
			"records": schema.ListNestedAttribute{
				Description: "List of records.",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							Description:         "dns.he.net record id",
							MarkdownDescription: "dns.he.net record id",
							Computed:            true,
						},
						"zone_id": schema.Int64Attribute{
							Description:         "dns.he.net zone id",
							MarkdownDescription: "dns.he.net zone id",
							Computed:            true,
						},
						"domain": schema.StringAttribute{
							Description:         "Name of the DNS record",
							MarkdownDescription: "Name of the DNS record",
							Computed:            true,
						},
						"record_type": schema.StringAttribute{
							Description:         "DNS record type",
							MarkdownDescription: "DNS record type",
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
						"data": schema.StringAttribute{
							Description:         "Value of the DNS record, e.g. IP address",
							MarkdownDescription: "Value of the DNS record, e.g. IP address",
							Computed:            true,
						},
						"dynamic": schema.BoolAttribute{
							Description:         "Enable DDNS for this record",
							MarkdownDescription: "Enable DDNS for this record",
							Computed:            true,
						},
						"locked": schema.BoolAttribute{
							Description:         "Record immutable flag",
							MarkdownDescription: "Record immutable flag",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

// Configure adds the provider configured client to the data source.
func (d *records) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if cli, ok := configure(ctx, req, resp); ok {
		d.client = cli
	}
}

// Read refreshes the Terraform state with the latest data.
func (d records) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state recordsModel

	// Retrieve values from state
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Terraform log
	ctxLog := tflog.SetField(ctx, "zone_id", state.ID.String())
	tflog.Debug(ctxLog, "Retrieving DNS records")

	records, err := d.client.GetRecords(ctx, uint(state.ID.ValueInt64()))
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

	// Map response body to model
	for _, record := range records {
		recordstate := recordModel{}

		if err := recordstate.setRecord(record); err != nil {
			resp.Diagnostics.AddError(
				"Unable to set DNS record",
				err.Error(),
			)
			return
		}

		state.Records = append(state.Records, recordstate)
	}

	// Set state
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
