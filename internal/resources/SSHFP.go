package resources

import (
	"context"
	"fmt"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/client"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"
	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/tfmodels"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &sshfp{}
	_ resource.ResourceWithConfigure   = &sshfp{}
	_ resource.ResourceWithImportState = &sshfp{}
)

// NewSSHFP initialises the A Resource.
func NewSSHFP() resource.Resource {
	return &sshfp{}
}

// sshfp is the data source implementation.
type sshfp struct {
	client *client.Client
}

// Metadata returns the resource type name.
func (sshfp) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sshfp" // TODO: maybe rename
}

// Schema defines the schema for the resource.
func (sshfp) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
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
				Required:            true,
				Description:         "Value of the DNS record: *TODO*",
				MarkdownDescription: "Value of the DNS record: *TODO*",
				Validators: []validator.String{
					sshfpValidator,
				},
			},
		},
	}
}

// Configure adds the provider configured client to the resource.
func (r *sshfp) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if cli, ok := configure(ctx, req, resp); ok {
		r.client = cli
	}
}

// Create creates the resource and sets the initial Terraform state.
func (r sshfp) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var state tfmodels.SSHFP

	// Retrieve values from state
	diags := req.Plan.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	recordSSHFP, err := state.GetRecord()
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to build SSHFP record",
			err.Error(),
		)
		return
	}

	recordX, err := r.client.SetRecord(ctx, recordSSHFP)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to create SSHFP record",
			err.Error(),
		)
		return
	}

	recordSSHFP, ok := recordX.(models.SSHFP)
	if !ok {
		resp.Diagnostics.AddError(
			"Unable to cast SSHFP record",
			fmt.Sprintf("unexpected record type %T", recordX),
		)
		return
	}

	if err = state.SetRecord(recordSSHFP); err != nil {
		resp.Diagnostics.AddError(
			"Unable to set SSHFP record",
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
func (r sshfp) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state tfmodels.SSHFP

	// Retrieve values from state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Retrieves record from dns.he.net, handles logging and errors
	recordX, ok := readRecord(ctx, r.client, state.ID, state.ZoneID, "SSHFP", resp)
	if !ok {
		return
	}

	recordSSHFP, ok := recordX.(models.SSHFP)
	if !ok {
		resp.Diagnostics.AddError(
			"Unable to cast SSHFP record",
			fmt.Sprintf("unexpected record type %T", recordX),
		)
		return
	}

	if err := state.SetRecord(recordSSHFP); err != nil {
		resp.Diagnostics.AddError(
			"Unable to set SSHFP record",
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
func (r sshfp) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state tfmodels.SSHFP

	// Retrieve values from state
	diags := req.Plan.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	recordSSHFP, err := state.GetRecord()
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to build SSHFP record",
			err.Error(),
		)
		return
	}

	recordX, err := r.client.SetRecord(ctx, recordSSHFP)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to update SSHFP record",
			err.Error(),
		)
		return
	}

	recordSSHFP, ok := recordX.(models.SSHFP)
	if !ok {
		resp.Diagnostics.AddError(
			"Unable to cast SSHFP record",
			fmt.Sprintf("unexpected record type %T", recordX),
		)
		return
	}

	if err = state.SetRecord(recordSSHFP); err != nil {
		resp.Diagnostics.AddError(
			"Unable to set SSHFP record",
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
func (r sshfp) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state tfmodels.SSHFP

	// Retrieve values from state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	recordSSHFP, err := state.GetRecord()
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to build SSHFP record",
			err.Error(),
		)
		return
	}

	err = r.client.DeleteRecord(ctx, recordSSHFP)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to delete SSHFP record",
			err.Error(),
		)
	}
}

func (sshfp) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	importRecordState(ctx, req, resp)
}
