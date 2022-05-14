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

type resourceBGPIPv4UnicastVRFNeighborType struct{}

func (t resourceBGPIPv4UnicastVRFNeighborType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This resource can manage the BGP IPv4 Unicast VRF Neighbor configuration.",

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
			"asn": {
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Type:                types.StringType,
				Required:            true,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					tfsdk.RequiresReplace(),
				},
			},
			"vrf": {
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Type:                types.StringType,
				Required:            true,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					tfsdk.RequiresReplace(),
				},
			},
			"ip": {
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Type:                types.StringType,
				Required:            true,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					tfsdk.RequiresReplace(),
				},
			},
			"remote_as": {
				MarkdownDescription: helpers.NewAttributeDescription("Specify a BGP peer-group remote-as").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
			},
			"description": {
				MarkdownDescription: helpers.NewAttributeDescription("Neighbor specific description").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
			},
			"shutdown": {
				MarkdownDescription: helpers.NewAttributeDescription("Administratively shut down this neighbor").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"update_source_loopback": {
				MarkdownDescription: helpers.NewAttributeDescription("Loopback interface").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
			},
			"activate": {
				MarkdownDescription: helpers.NewAttributeDescription("Enable the address family for this neighbor").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"send_community": {
				MarkdownDescription: helpers.NewAttributeDescription("").AddStringEnumDescription("both", "extended", "standard").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.StringEnumValidator("both", "extended", "standard"),
				},
			},
			"route_reflector_client": {
				MarkdownDescription: helpers.NewAttributeDescription("Configure a neighbor as Route Reflector client").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
		},
	}, nil
}

func (t resourceBGPIPv4UnicastVRFNeighborType) NewResource(ctx context.Context, in tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	provider, diags := convertProviderType(in)

	return resourceBGPIPv4UnicastVRFNeighbor{
		provider: provider,
	}, diags
}

type resourceBGPIPv4UnicastVRFNeighbor struct {
	provider provider
}

func (r resourceBGPIPv4UnicastVRFNeighbor) Create(ctx context.Context, req tfsdk.CreateResourceRequest, resp *tfsdk.CreateResourceResponse) {
	var plan BGPIPv4UnicastVRFNeighbor

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

func (r resourceBGPIPv4UnicastVRFNeighbor) Read(ctx context.Context, req tfsdk.ReadResourceRequest, resp *tfsdk.ReadResourceResponse) {
	var state BGPIPv4UnicastVRFNeighbor

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.Value))

	res, err := r.provider.clients[state.Device.Value].GetData(state.Id.Value)
	if res.StatusCode == 404 {
		state = BGPIPv4UnicastVRFNeighbor{Device: state.Device, Id: state.Id}
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

func (r resourceBGPIPv4UnicastVRFNeighbor) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {
	var plan, state BGPIPv4UnicastVRFNeighbor

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

func (r resourceBGPIPv4UnicastVRFNeighbor) Delete(ctx context.Context, req tfsdk.DeleteResourceRequest, resp *tfsdk.DeleteResourceResponse) {
	var state BGPIPv4UnicastVRFNeighbor

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

func (r resourceBGPIPv4UnicastVRFNeighbor) ImportState(ctx context.Context, req tfsdk.ImportResourceStateRequest, resp *tfsdk.ImportResourceStateResponse) {
	tfsdk.ResourceImportStatePassthroughID(ctx, tftypes.NewAttributePath().WithAttributeName("id"), req, resp)
}
