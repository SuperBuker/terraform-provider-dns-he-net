package resources

import (
	"context"
	"fmt"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/client"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/client/filters"
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
	_ resource.Resource                = &a{}
	_ resource.ResourceWithConfigure   = &a{}
	_ resource.ResourceWithImportState = &a{}
	_ resource.ResourceWithConfigure   = &a{}
)

// NewA initialises the A Resource.
func NewA() resource.Resource {
	return &a{}
}

// a is the data source implementation.
type a struct {
	client *client.Client
}

// Metadata returns the resource type name.
func (a) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_a" // TODO: maybe rename
}

// Schema defines the schema for the resource.
func (a) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"zone_id": schema.Int64Attribute{
				Required:            true,
				MarkdownDescription: "",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
			},
			"domain": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "",
				Validators: []validator.String{
					domainValidator,
				},
				// Misses top domain validation
			},
			"ttl": schema.Int64Attribute{
				Required:            true,
				MarkdownDescription: "",
				Validators: []validator.Int64{
					int64validator.Between(300, 86400),
				},
			},
			"data": schema.StringAttribute{
				Computed:            true,
				Optional:            true,
				MarkdownDescription: "",
				Validators: []validator.String{
					ipv4Validator,
				},
				PlanModifiers: []planmodifier.String{
					planmodifiers.UseStateOrDftForUnknown("0.0.0.0"),
				},
			},
			"dynamic": schema.BoolAttribute{
				Computed:            true, // Isn't really computed...
				Optional:            true,
				MarkdownDescription: "",
				Default:             booldefault.StaticBool(false),
			},
		},
	}
}

// Configure adds the provider configured client to the resource.
func (r *a) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

	r.client = cli
}

// Create creates the resource and sets the initial Terraform state.
func (r a) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var state tfmodels.A

	// Retrieve values from state
	diags := req.Plan.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	recordA, err := state.GetRecord()
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to build A record",
			err.Error(),
		)
		return
	}

	recordX, err := r.client.SetRecord(ctx, recordA)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to create A record",
			err.Error(),
		)
		return
	}

	recordA, ok := recordX.(models.A)
	if !ok {
		resp.Diagnostics.AddError(
			"Unable to cast A record",
			fmt.Sprintf("unexpacted record type %T", recordX),
		)
		return
	}

	if err = state.SetRecord(recordA); err != nil {
		resp.Diagnostics.AddError(
			"Unable to set A record",
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
func (r a) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state tfmodels.A

	// Retrieve values from state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	records, err := r.client.GetRecords(ctx, uint(state.ZoneID.ValueInt64())) //GetOne(state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to fetch DNS records",
			err.Error(),
		)
		return
	}

	record, ok := filters.RecordById(records, uint(state.ID.ValueInt64()))
	if !ok {
		resp.Diagnostics.AddError(
			"Unable to find A record",
			fmt.Sprintf("record %q in zone %q doesn't exist", state.ID.String(), state.ZoneID.String()),
		)
		return
	}

	recordX, err := record.ToX()
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to cast A record",
			err.Error(),
		)
		return
	}

	recordA, ok := recordX.(models.A)
	if !ok {
		resp.Diagnostics.AddError(
			"Unable to cast A record",
			fmt.Sprintf("unexpacted record type %T", recordX),
		)
		return
	}

	if err = state.SetRecord(recordA); err != nil {
		resp.Diagnostics.AddError(
			"Unable to set A record",
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
func (r a) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state tfmodels.A

	// Retrieve values from state
	diags := req.Plan.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	recordA, err := state.GetRecord()
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to build A record",
			err.Error(),
		)
		return
	}

	recordX, err := r.client.SetRecord(ctx, recordA)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to update A record",
			err.Error(),
		)
		return
	}

	recordA, ok := recordX.(models.A)
	if !ok {
		resp.Diagnostics.AddError(
			"Unable to cast A record",
			fmt.Sprintf("unexpacted record type %T", recordX),
		)
		return
	}

	if err = state.SetRecord(recordA); err != nil {
		resp.Diagnostics.AddError(
			"Unable to set A record",
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
func (r a) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state tfmodels.A

	// Retrieve values from state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	recordA, err := state.GetRecord()
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to build A record",
			err.Error(),
		)
		return
	}

	err = r.client.DeleteRecord(ctx, recordA)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to delete A record",
			err.Error(),
		)
	}
}

func (a) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var config tfmodels.A

	// Retrieve values from config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	recordA, err := config.GetRecord()
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to build A record",
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
			"Invalid A record configuration",
			"Static A records must have Data configured.",
		)
		return
	}
}

func (a) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	importRecordState(ctx, req, resp)
}
