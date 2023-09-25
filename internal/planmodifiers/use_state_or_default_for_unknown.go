package planmodifiers

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// UseStateOrDftForUnknown returns a plan modifier that copies a known prior state
// value into the planned value, or a default value if not known prior state exists.
// Use this when it is known that an unconfigured value will remain the same after
// a resource update.
//
// To prevent Terraform errors, the framework automatically sets unconfigured
// and Computed attributes to an unknown value "(known after apply)" on update.
// Using this plan modifier will instead display the prior state value in the
// plan, unless a prior plan modifier adjusts the value.
func UseStateOrDftForUnknown(dft string) planmodifier.String {
	return useStateOrDftForUnknownModifier{
		dft: dft,
	}
}

// useStateForUnknownModifier implements the plan modifier.
type useStateOrDftForUnknownModifier struct {
	dft string
}

// Description returns a human-readable description of the plan modifier.
func (useStateOrDftForUnknownModifier) Description(context.Context) string {
	return "Once set, the value of this attribute in state will not change."
}

// MarkdownDescription returns a markdown description of the plan modifier.
func (useStateOrDftForUnknownModifier) MarkdownDescription(context.Context) string {
	return "Once set, the value of this attribute in state will not change."
}

// PlanModifyString implements the plan modification logic.
func (m useStateOrDftForUnknownModifier) PlanModifyString(_ context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	// Do nothing if there is a known planned value.
	if !req.PlanValue.IsUnknown() {
		return
	}

	// Do nothing if there is an unknown configuration value, otherwise interpolation gets messed up.
	if req.ConfigValue.IsUnknown() {
		return
	}

	// Set default if there is no state value.
	if req.StateValue.IsNull() {
		resp.PlanValue = types.StringValue(m.dft)
	} else {
		resp.PlanValue = req.StateValue
	}
}
