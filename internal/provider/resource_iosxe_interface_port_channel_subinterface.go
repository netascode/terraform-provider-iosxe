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

type resourceInterfacePortChannelSubinterfaceType struct{}

func (t resourceInterfacePortChannelSubinterfaceType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This resource can manage the Interface Port Channel Subinterface configuration.",

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
			"name": {
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Type:                types.StringType,
				Required:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.StringPatternValidator(0, 0, `[1-9][0-9]*\.[1-9][0-9]*`),
				},
				PlanModifiers: tfsdk.AttributePlanModifiers{
					tfsdk.RequiresReplace(),
				},
			},
			"description": {
				MarkdownDescription: helpers.NewAttributeDescription("Interface specific description").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.StringPatternValidator(0, 200, `.*`),
				},
			},
			"shutdown": {
				MarkdownDescription: helpers.NewAttributeDescription("Shutdown the selected interface").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"vrf_forwarding": {
				MarkdownDescription: helpers.NewAttributeDescription("Configure forwarding table").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
			},
			"ipv4_address": {
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.StringPatternValidator(0, 0, `(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?`),
				},
			},
			"ipv4_address_mask": {
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.StringPatternValidator(0, 0, `(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?`),
				},
			},
			"encapsulation_dot1q_vlan_id": {
				MarkdownDescription: helpers.NewAttributeDescription("").AddIntegerRangeDescription(1, 4094).String,
				Type:                types.Int64Type,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.IntegerRangeValidator(1, 4094),
				},
			},
			"helper_addresses": {
				MarkdownDescription: helpers.NewAttributeDescription("Specify a destination address for UDP broadcasts").String,
				Optional:            true,
				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
					"address": {
						MarkdownDescription: helpers.NewAttributeDescription("").String,
						Type:                types.StringType,
						Optional:            true,
						Computed:            true,
						Validators: []tfsdk.AttributeValidator{
							helpers.StringPatternValidator(0, 0, `(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?`),
						},
					},
					"global": {
						MarkdownDescription: helpers.NewAttributeDescription("Helper-address is global").String,
						Type:                types.BoolType,
						Optional:            true,
						Computed:            true,
					},
					"vrf": {
						MarkdownDescription: helpers.NewAttributeDescription("VRF name for helper-address (if different from interface VRF)").String,
						Type:                types.StringType,
						Optional:            true,
						Computed:            true,
					},
				}, tfsdk.ListNestedAttributesOptions{}),
			},
		},
	}, nil
}

func (t resourceInterfacePortChannelSubinterfaceType) NewResource(ctx context.Context, in tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	provider, diags := convertProviderType(in)

	return resourceInterfacePortChannelSubinterface{
		provider: provider,
	}, diags
}

type resourceInterfacePortChannelSubinterface struct {
	provider provider
}

func (r resourceInterfacePortChannelSubinterface) Create(ctx context.Context, req tfsdk.CreateResourceRequest, resp *tfsdk.CreateResourceResponse) {
	var plan InterfacePortChannelSubinterface

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

	plan.setUnknownValues()

	plan.Id = types.String{Value: plan.getPath()}

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.getPath()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r resourceInterfacePortChannelSubinterface) Read(ctx context.Context, req tfsdk.ReadResourceRequest, resp *tfsdk.ReadResourceResponse) {
	var state InterfacePortChannelSubinterface

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.Value))

	res, err := r.provider.clients[state.Device.Value].GetData(state.Id.Value)
	if res.StatusCode == 404 {
		state = InterfacePortChannelSubinterface{Device: state.Device, Id: state.Id}
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

func (r resourceInterfacePortChannelSubinterface) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {
	var plan, state InterfacePortChannelSubinterface

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

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.Value))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r resourceInterfacePortChannelSubinterface) Delete(ctx context.Context, req tfsdk.DeleteResourceRequest, resp *tfsdk.DeleteResourceResponse) {
	var state InterfacePortChannelSubinterface

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.Value))

	res, err := r.provider.clients[state.Device.Value].DeleteData(state.Id.Value)
	if err != nil && res.StatusCode != 404 {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object, got error: %s", err))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.Value))

	resp.State.RemoveResource(ctx)
}

func (r resourceInterfacePortChannelSubinterface) ImportState(ctx context.Context, req tfsdk.ImportResourceStateRequest, resp *tfsdk.ImportResourceStateResponse) {
	tfsdk.ResourceImportStatePassthroughID(ctx, tftypes.NewAttributePath().WithAttributeName("id"), req, resp)
}
