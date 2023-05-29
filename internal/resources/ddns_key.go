package resources

import (
	"context"
	"fmt"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/client"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/client/filters"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/ddns"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource               = &ddnsKey{}
	_ resource.ResourceWithConfigure  = &ddnsKey{}
	_ resource.ResourceWithModifyPlan = &ddnsKey{}
)

// NewAccount initialises the DDNS Key Resource.
func NewDDNSKey() resource.Resource {
	return &ddnsKey{}
}

// account is the data source implementation.
type ddnsKey struct {
	client *client.Client
}

// accountModel maps the data source schema data.
type ddnsKeyModel struct {
	ID     types.String `tfsdk:"id"`
	Domain types.String `tfsdk:"domain"`
	ZoneID types.Int64  `tfsdk:"zone_id"`
	Key    types.String `tfsdk:"key"`
}

func (dk *ddnsKeyModel) set(ndk models.DDNSKey) {

	dk.Domain = types.StringValue(ndk.Domain)
	dk.ZoneID = types.Int64Value(int64(ndk.ZoneID))
	dk.Key = types.StringValue(ndk.Key)

	return
}

func (dk ddnsKeyModel) get() models.DDNSKey {
	return models.DDNSKey{
		Domain: dk.Domain.ValueString(),
		ZoneID: uint(dk.ZoneID.ValueInt64()),
		Key:    dk.Key.ValueString(),
	}
}

// Metadata returns the resource type name.
func (ddnsKey) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ddnskey" // TODO: maybe rename
}

// Schema defines the schema for the resource.
func (ddnsKey) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"domain": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"zone_id": schema.Int64Attribute{
				Required:            true,
				MarkdownDescription: "",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
			},
			"key": schema.StringAttribute{
				Required:            true,
				Sensitive:           true,
				MarkdownDescription: "",
				// TODO: Missing validation
			},
		},
	}
}

// Configure adds the provider configured client to the resource.
func (dk *ddnsKey) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

	dk.client = cli
}

// Create creates the resource and sets the initial Terraform state.
func (dk ddnsKey) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var state ddnsKeyModel

	// Retrieve values from state
	diags := req.Plan.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	ddnsKey := state.get()

	_, err := dk.client.SetDDNSKey(ctx, ddnsKey)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to set DDNS key",
			err.Error(),
		)
		return
	}

	// Set ID
	state.ID = state.Domain

	// Set state, it's mandatory
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read refreshes the Terraform state with the latest data.
func (dk ddnsKey) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ddnsKeyModel

	// Retrieve values from state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	zones, err := dk.client.GetZones(ctx)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to fetch DNS records",
			err.Error(),
		)
		return
	}

	_, ok := filters.ZoneById(zones, uint(state.ZoneID.ValueInt64()))
	if !ok {
		resp.Diagnostics.AddError(
			"Unable to find zone",
			fmt.Sprintf("zone %q doesn't exist", state.ZoneID.String()),
		)
		return
	}

	// Set ID
	if state.ID.IsUnknown() {
		state.ID = state.Domain
	}

	// Check if the state key is still valid, remove otherwhise
	ok, err = dk.client.DDNS().CheckAuth(ctx, state.Domain.ValueString(), state.Key.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to validate DDNS key",
			err.Error(),
		)
		return
	} else if !ok {
		state.Key = types.StringUnknown()
	}

	// Set state
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update updates the resource and sets the updated Terraform state on success.
func (dk ddnsKey) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state ddnsKeyModel

	// Retrieve values from state
	diags := req.Plan.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	ddnsKey := state.get()

	_, err := dk.client.SetDDNSKey(ctx, ddnsKey)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to set DDNS key",
			err.Error(),
		)
		return
	}

	// Set state, it's mandatory
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Delete deletes the resource and removes the Terraform state on success.
func (ddnsKey) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Maybe set a random key value
}

func (ddnsKey) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	var planData, stateData ddnsKeyModel

	if req.State.Raw.IsNull() {
		// Create record
		return
	}

	if req.Plan.Raw.IsNull() {
		// Destroying record
		resp.Diagnostics.AddWarning(
			"Resource Destruction Considerations",
			"Applying this resource destruction will only remove the resource from the Terraform state "+
				"and will not call the deletion API due to API limitations. Currently there's no way to "+
				"fully destroy this resource, just alter the update key.",
		)
		return
	}

	// Retrieve values from plan
	diags := req.Plan.Get(ctx, &planData)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Retrieve values from state
	diags = req.State.Get(ctx, &stateData)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set ID in plan if it's not set
	if !stateData.ID.IsUnknown() {
		planData.ID = stateData.Domain
	}

	// Set plan
	diags = resp.Plan.Set(ctx, &planData)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
