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
	_ resource.Resource                = &rp{}
	_ resource.ResourceWithConfigure   = &rp{}
	_ resource.ResourceWithImportState = &rp{}
)

// NewRP initialises the A Resource.
func NewRP() resource.Resource {
	return &rp{}
}

// rp is the data source implementation.
type rp struct {
	client *client.Client
}

// Metadata returns the resource type name.
func (rp) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_rp" // TODO: maybe rename
}

// Schema defines the schema for the resource.
func (rp) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description:         "DNS RP record",
		MarkdownDescription: "DNS RP record",
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
					rpValidator,
				},
			},
		},
	}
}

// Configure adds the provider configured client to the resource.
func (r *rp) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if cli, ok := configure(ctx, req, resp); ok {
		r.client = cli
	}
}

// Create creates the resource and sets the initial Terraform state.
func (r rp) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var state tfmodels.RP

	// Retrieve values from state
	diags := req.Plan.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	recordRP, err := state.GetRecord()
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to build RP record",
			err.Error(),
		)
		return
	}

	recordX, err := r.client.SetRecord(ctx, recordRP)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to create RP record",
			err.Error(),
		)
		return
	}

	recordRP, ok := recordX.(models.RP)
	if !ok {
		resp.Diagnostics.AddError(
			"Unable to cast RP record",
			fmt.Sprintf("unexpected record type %T", recordX),
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

// Read refreshes the Terraform state with the latest data.
func (r rp) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state tfmodels.RP

	// Retrieve values from state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Retrieves record from dns.he.net, handles logging and errors
	recordX, ok := readRecord(ctx, r.client, state.ID, state.ZoneID, "RP", resp)
	if !ok {
		return
	}

	recordRP, ok := recordX.(models.RP)
	if !ok {
		resp.Diagnostics.AddError(
			"Unable to cast RP record",
			fmt.Sprintf("unexpected record type %T", recordX),
		)
		return
	}

	if err := state.SetRecord(recordRP); err != nil {
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

// Update updates the resource and sets the updated Terraform state on success.
func (r rp) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state tfmodels.RP

	// Retrieve values from state
	diags := req.Plan.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	recordRP, err := state.GetRecord()
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to build RP record",
			err.Error(),
		)
		return
	}

	recordX, err := r.client.SetRecord(ctx, recordRP)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to update RP record",
			err.Error(),
		)
		return
	}

	recordRP, ok := recordX.(models.RP)
	if !ok {
		resp.Diagnostics.AddError(
			"Unable to cast RP record",
			fmt.Sprintf("unexpected record type %T", recordX),
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

// Delete deletes the resource and removes the Terraform state on success.
func (r rp) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state tfmodels.RP

	// Retrieve values from state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	recordRP, err := state.GetRecord()
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to build RP record",
			err.Error(),
		)
		return
	}

	err = r.client.DeleteRecord(ctx, recordRP)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to delete RP record",
			err.Error(),
		)
	}
}

func (rp) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	importRecordState(ctx, req, resp)
}
