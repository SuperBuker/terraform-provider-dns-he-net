package resources

import (
	"context"
	"fmt"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/client"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"
	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/planmodifiers"
	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/tfmodels"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &txt{}
	_ resource.ResourceWithConfigure   = &txt{}
	_ resource.ResourceWithImportState = &txt{}
	_ resource.ResourceWithConfigure   = &txt{}
)

// NewTXT initialises the A Resource.
func NewTXT() resource.Resource {
	return &txt{}
}

// txt is the data source implementation.
type txt struct {
	client *client.Client
}

// Metadata returns the resource type name.
func (txt) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_txt" // TODO: maybe rename
}

// Schema defines the schema for the resource.
func (txt) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description:         "DNS TXT record",
		MarkdownDescription: "DNS TXT record",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Computed:            true,
				Description:         "dns.he.net record id",
				MarkdownDescription: "dns.he.net record id",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"zone_id": schema.Int64Attribute{
				Required:            true,
				Description:         "dns.he.net zone id",
				MarkdownDescription: "dns.he.net zone id",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
			},
			"domain": schema.StringAttribute{
				Required:            true,
				Description:         "Name of the DNS record",
				MarkdownDescription: "Name of the DNS record",
				Validators: []validator.String{
					domainValidator,
				},
				// Misses top domain validation
			},
			"ttl": schema.Int64Attribute{
				Required:            true,
				Description:         "Time-To-Live of the DNS record",
				MarkdownDescription: "Time-To-Live of the DNS record",
				Validators: []validator.Int64{
					int64validator.Between(300, 86400),
				},
			},
			"data": schema.StringAttribute{
				Computed:            true,
				Optional:            true,
				Description:         "Value of the DNS record: arbitrary text",
				MarkdownDescription: "Value of the DNS record: arbitrary text",
				Validators: []validator.String{
					txtValidator,
				},
				PlanModifiers: []planmodifier.String{
					planmodifiers.UseStateOrDftForUnknown(`""`),
				},
			},
			"dynamic": schema.BoolAttribute{
				Computed:            true, // Isn't really computed...
				Optional:            true,
				Description:         "Enable DDNS for this record",
				MarkdownDescription: "Enable DDNS for this record",
				Default:             booldefault.StaticBool(false),
			},
		},
	}
}

// Configure adds the provider configured client to the resource.
func (r *txt) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if cli, ok := configure(ctx, req, resp); ok {
		r.client = cli
	}
}

// Create creates the resource and sets the initial Terraform state.
func (r txt) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var state tfmodels.TXT

	// Retrieve values from state
	diags := req.Plan.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	recordA, err := state.GetRecord()
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to build TXT record",
			err.Error(),
		)
		return
	}

	recordX, err := r.client.SetRecord(ctx, recordA)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to create TXT record",
			err.Error(),
		)
		return
	}

	recordA, ok := recordX.(models.TXT)
	if !ok {
		resp.Diagnostics.AddError(
			"Unable to cast TXT record",
			fmt.Sprintf("unexpected record type %T", recordX),
		)
		return
	}

	if err = state.SetRecord(recordA); err != nil {
		resp.Diagnostics.AddError(
			"Unable to set TXT record",
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

// Read refreshes the Terraform state with the latest data.
func (r txt) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state tfmodels.TXT

	// Retrieve values from state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Retrieves record from dns.he.net, handles logging and errors
	recordX, ok := readRecord(ctx, r.client, state.ID, state.ZoneID, "TXT", resp)
	if !ok {
		return
	}

	recordA, ok := recordX.(models.TXT)
	if !ok {
		resp.Diagnostics.AddError(
			"Unable to cast TXT record",
			fmt.Sprintf("unexpected record type %T", recordX),
		)
		return
	}

	if err := state.SetRecord(recordA); err != nil {
		resp.Diagnostics.AddError(
			"Unable to set TXT record",
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

// Update updates the resource and sets the updated Terraform state on success.
func (r txt) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state tfmodels.TXT

	// Retrieve values from state
	diags := req.Plan.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	recordA, err := state.GetRecord()
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to build TXT record",
			err.Error(),
		)
		return
	}

	recordX, err := r.client.SetRecord(ctx, recordA)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to update TXT record",
			err.Error(),
		)
		return
	}

	recordA, ok := recordX.(models.TXT)
	if !ok {
		resp.Diagnostics.AddError(
			"Unable to cast TXT record",
			fmt.Sprintf("unexpected record type %T", recordX),
		)
		return
	}

	if err = state.SetRecord(recordA); err != nil {
		resp.Diagnostics.AddError(
			"Unable to set TXT record",
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

// Delete deletes the resource and removes the Terraform state on success.
func (r txt) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state tfmodels.TXT

	// Retrieve values from state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	recordA, err := state.GetRecord()
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to build TXT record",
			err.Error(),
		)
		return
	}

	err = r.client.DeleteRecord(ctx, recordA)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to delete TXT record",
			err.Error(),
		)
	}
}

func (txt) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var config tfmodels.TXT

	// Retrieve values from config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	recordA, err := config.GetRecord()
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to build TXT record",
			err.Error(),
		)
		return
	}

	// Validate configuration
	if !config.Data.IsUnknown() && !config.Data.IsNull() {
		// pass
	} else if recordA.Dynamic {
		resp.Diagnostics.AddAttributeWarning(
			path.Root("data"),
			"Missing Attribute Configuration",
			"Applying default configuration",
		)
	} else {
		resp.Diagnostics.AddError(
			"Invalid TXT record configuration",
			"Static TXT records must have Data configured.",
		)
		return
	}
}

func (txt) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	importRecordState(ctx, req, resp)
}
