// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/terraform-provider-iosxe/internal/provider/helpers"
)

type resourcePIMType struct{}

func (t resourcePIMType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This resource can manage the PIM configuration.",

		Attributes: map[string]tfsdk.Attribute{
			"device": {
				MarkdownDescription: "A device name from the provider configuration.",
				Type:                types.StringType,
				Optional:            true,
			},
			"id": {
				MarkdownDescription: "The path of the object.",
				Type:                types.StringType,
				Computed:            true,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					tfsdk.UseStateForUnknown(),
				},
			},
			"autorp": {
				MarkdownDescription: helpers.NewAttributeDescription("Configure AutoRP global operations").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"autorp_listener": {
				MarkdownDescription: helpers.NewAttributeDescription("Allow AutoRP packets across sparse mode interface").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"bsr_candidate_loopback": {
				MarkdownDescription: helpers.NewAttributeDescription("Loopback interface").AddIntegerRangeDescription(0, 2147483647).String,
				Type:                types.Int64Type,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.IntegerRangeValidator(0, 2147483647),
				},
			},
			"bsr_candidate_mask": {
				MarkdownDescription: helpers.NewAttributeDescription("Hash Mask length for RP selection").AddIntegerRangeDescription(0, 32).String,
				Type:                types.Int64Type,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.IntegerRangeValidator(0, 32),
				},
			},
			"bsr_candidate_priority": {
				MarkdownDescription: helpers.NewAttributeDescription("Priority value for candidate bootstrap router").AddIntegerRangeDescription(0, 255).String,
				Type:                types.Int64Type,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.IntegerRangeValidator(0, 255),
				},
			},
			"bsr_candidate_accept_rp_candidate": {
				MarkdownDescription: helpers.NewAttributeDescription("BSR RP candidate filter").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
			},
			"ssm_range": {
				MarkdownDescription: helpers.NewAttributeDescription("ACL for group range to be used for SSM").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
			},
			"ssm_default": {
				MarkdownDescription: helpers.NewAttributeDescription("Use 232/8 group range for SSM").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"rp_address": {
				MarkdownDescription: helpers.NewAttributeDescription("IP address of Rendezvous-point for group").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.StringPatternValidator(0, 0, `(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?`),
				},
			},
			"rp_address_override": {
				MarkdownDescription: helpers.NewAttributeDescription("Overrides dynamically learnt RP mappings").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"rp_address_bidir": {
				MarkdownDescription: helpers.NewAttributeDescription("Group range treated in bidirectional shared-tree mode").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"rp_addresses": {
				MarkdownDescription: helpers.NewAttributeDescription("PIM RP-address (Rendezvous Point)").String,
				Optional:            true,
				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
					"access_list": {
						MarkdownDescription: helpers.NewAttributeDescription("IP Access-list").String,
						Type:                types.StringType,
						Optional:            true,
						Computed:            true,
					},
					"rp_address": {
						MarkdownDescription: helpers.NewAttributeDescription("IP address of Rendezvous-point for group").String,
						Type:                types.StringType,
						Optional:            true,
						Computed:            true,
						Validators: []tfsdk.AttributeValidator{
							helpers.StringPatternValidator(0, 0, `(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?`),
						},
					},
					"override": {
						MarkdownDescription: helpers.NewAttributeDescription("Overrides dynamically learnt RP mappings").String,
						Type:                types.BoolType,
						Optional:            true,
						Computed:            true,
					},
					"bidir": {
						MarkdownDescription: helpers.NewAttributeDescription("Group range treated in bidirectional shared-tree mode").String,
						Type:                types.BoolType,
						Optional:            true,
						Computed:            true,
					},
				}, tfsdk.ListNestedAttributesOptions{}),
			},
			"rp_candidates": {
				MarkdownDescription: helpers.NewAttributeDescription("To be a PIM version 2 RP candidate").String,
				Optional:            true,
				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
					"interface": {
						MarkdownDescription: helpers.NewAttributeDescription("Autonomic-Networking virtual interface").String,
						Type:                types.StringType,
						Optional:            true,
						Computed:            true,
					},
					"group_list": {
						MarkdownDescription: helpers.NewAttributeDescription("IP Access list").String,
						Type:                types.StringType,
						Optional:            true,
						Computed:            true,
					},
					"interval": {
						MarkdownDescription: helpers.NewAttributeDescription("RP candidate advertisement interval").AddIntegerRangeDescription(1, 16383).String,
						Type:                types.Int64Type,
						Optional:            true,
						Computed:            true,
						Validators: []tfsdk.AttributeValidator{
							helpers.IntegerRangeValidator(1, 16383),
						},
					},
					"priority": {
						MarkdownDescription: helpers.NewAttributeDescription("RP candidate priority").AddIntegerRangeDescription(0, 255).String,
						Type:                types.Int64Type,
						Optional:            true,
						Computed:            true,
						Validators: []tfsdk.AttributeValidator{
							helpers.IntegerRangeValidator(0, 255),
						},
					},
					"bidir": {
						MarkdownDescription: helpers.NewAttributeDescription("Group range treated in bidirectional shared-tree mode").String,
						Type:                types.BoolType,
						Optional:            true,
						Computed:            true,
					},
				}, tfsdk.ListNestedAttributesOptions{}),
			},
		},
	}, nil
}

func (t resourcePIMType) NewResource(ctx context.Context, in tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	provider, diags := convertProviderType(in)

	return resourcePIM{
		provider: provider,
	}, diags
}

type resourcePIM struct {
	provider provider
}

func (r resourcePIM) Create(ctx context.Context, req tfsdk.CreateResourceRequest, resp *tfsdk.CreateResourceResponse) {
	var plan PIM

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.getPath()))

	// Create object
	body := plan.toBody()

	res, err := r.provider.clients[plan.Device.Value].PatchData(plan.getPathShort(), body)
	if len(res.Errors.Error) > 0 && res.Errors.Error[0].ErrorMessage == "patch to a nonexistent resource" {
		_, err = r.provider.clients[plan.Device.Value].PutData(plan.getPath(), body)
	}
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PATCH), got error: %s", err))
		return
	}

	emptyLeafsDelete := plan.getEmptyLeafsDelete()
	tflog.Debug(ctx, fmt.Sprintf("List of empty leafs to delete: %+v", emptyLeafsDelete))

	for _, i := range emptyLeafsDelete {
		res, err := r.provider.clients[plan.Device.Value].DeleteData(i)
		if err != nil && res.StatusCode != 404 {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object, got error: %s", err))
			return
		}
	}

	plan.setUnknownValues()

	plan.Id = types.String{Value: plan.getPath()}

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.getPath()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r resourcePIM) Read(ctx context.Context, req tfsdk.ReadResourceRequest, resp *tfsdk.ReadResourceResponse) {
	var state PIM

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.Value))

	res, err := r.provider.clients[state.Device.Value].GetData(state.Id.Value)
	if res.StatusCode == 404 {
		state = PIM{Device: state.Device, Id: state.Id}
	} else {
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
			return
		}

		state.updateFromBody(res.Res)
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.Value))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r resourcePIM) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {
	var plan, state PIM

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Read state
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.Id.Value))

	body := plan.toBody()
	res, err := r.provider.clients[plan.Device.Value].PatchData(plan.getPathShort(), body)
	if len(res.Errors.Error) > 0 && res.Errors.Error[0].ErrorMessage == "patch to a nonexistent resource" {
		_, err = r.provider.clients[plan.Device.Value].PutData(plan.getPath(), body)
	}
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PATCH), got error: %s", err))
		return
	}

	plan.setUnknownValues()

	deletedListItems := plan.getDeletedListItems(state)
	tflog.Debug(ctx, fmt.Sprintf("List items to delete: %+v", deletedListItems))

	for _, i := range deletedListItems {
		res, err := r.provider.clients[state.Device.Value].DeleteData(i)
		if err != nil && res.StatusCode != 404 {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object, got error: %s", err))
			return
		}
	}

	emptyLeafsDelete := plan.getEmptyLeafsDelete()
	tflog.Debug(ctx, fmt.Sprintf("List of empty leafs to delete: %+v", emptyLeafsDelete))

	for _, i := range emptyLeafsDelete {
		res, err := r.provider.clients[plan.Device.Value].DeleteData(i)
		if err != nil && res.StatusCode != 404 {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object, got error: %s", err))
			return
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.Value))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r resourcePIM) Delete(ctx context.Context, req tfsdk.DeleteResourceRequest, resp *tfsdk.DeleteResourceResponse) {
	var state PIM

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.Value))

	res, err := r.provider.clients[state.Device.Value].DeleteData(state.Id.Value)
	if err != nil && res.StatusCode != 404 && res.StatusCode != 400 {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object, got error: %s", err))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.Value))

	resp.State.RemoveResource(ctx)
}

func (r resourcePIM) ImportState(ctx context.Context, req tfsdk.ImportResourceStateRequest, resp *tfsdk.ImportResourceStateResponse) {
	tfsdk.ResourceImportStatePassthroughID(ctx, tftypes.NewAttributePath().WithAttributeName("id"), req, resp)
}
