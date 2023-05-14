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
	_ datasource.DataSource              = &domain{}
	_ datasource.DataSourceWithConfigure = &domain{}
)

// NewDomain initialises the Domain DataSource.
func NewDomain() datasource.DataSource {
	return &domain{}
}

// cdomain is the data source implementation.
type domain struct {
	client *client.Client
}

// domainModel maps the data source schema data.
type domainModel struct {
	ID     types.Int64  `tfsdk:"id"`
	Domain types.String `tfsdk:"domain"`
}

func (d *domainModel) setDomain(domain models.Zone) error {
	d.ID = types.Int64Value(int64(domain.Id))
	d.Domain = types.StringValue(domain.Domain)

	return nil
}

func (d *domainModel) getDomain() (models.Zone, error) {
	return models.Zone{
		Id:     uint(d.ID.ValueInt64()),
		Domain: d.Domain.ValueString(),
	}, nil
}

// Metadata returns the data source type name.
func (domain) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_domain" // TODO: maybe rename
}

// Schema defines the schema for the data source.
func (domain) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
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
	}
}

// Configure adds the provider configured client to the data source.
func (d *domain) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
func (d domain) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state domainModel

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

	domain, ok := filters.DomainById(domains, uint(state.ID.ValueInt64()))
	if !ok {
		resp.Diagnostics.AddError(
			"Unable to find root domain",
			fmt.Sprintf(`root domain "%s" doesn't exist`, state.ID.String()),
		)
		return
	}

	if err := state.setDomain(domain); err != nil {
		resp.Diagnostics.AddError(
			"Unable to set root domain",
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
