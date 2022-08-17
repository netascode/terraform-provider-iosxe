// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/terraform-provider-iosxe/internal/provider/helpers"
)

type resourceEVPNInstanceType struct{}

func (t resourceEVPNInstanceType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This resource can manage the EVPN Instance configuration.",

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
			"evpn_instance_num": {
				MarkdownDescription: helpers.NewAttributeDescription("evpn instance number").AddIntegerRangeDescription(1, 65535).String,
				Type:                types.Int64Type,
				Required:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.IntegerRangeValidator(1, 65535),
				},
				PlanModifiers: tfsdk.AttributePlanModifiers{
					tfsdk.RequiresReplace(),
				},
			},
			"vlan_based_replication_type_ingress": {
				MarkdownDescription: helpers.NewAttributeDescription("Ingress replication").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"vlan_based_replication_type_static": {
				MarkdownDescription: helpers.NewAttributeDescription("Static replication").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"vlan_based_replication_type_p2mp": {
				MarkdownDescription: helpers.NewAttributeDescription("p2mp replication").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"vlan_based_replication_type_mp2mp": {
				MarkdownDescription: helpers.NewAttributeDescription("mp2mp replication").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"vlan_based_encapsulation": {
				MarkdownDescription: helpers.NewAttributeDescription("Data encapsulation method").AddStringEnumDescription("mpls", "vxlan").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.StringEnumValidator("mpls", "vxlan"),
				},
			},
			"vlan_based_auto_route_target": {
				MarkdownDescription: helpers.NewAttributeDescription("Automatically set a route-target").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"vlan_based_rd": {
				MarkdownDescription: helpers.NewAttributeDescription("ASN:nn or IP-address:nn").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
			},
			"vlan_based_route_target": {
				MarkdownDescription: helpers.NewAttributeDescription("ASN:nn or IP-address:nn").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
			},
			"vlan_based_route_target_both": {
				MarkdownDescription: helpers.NewAttributeDescription("ASN:nn or IP-address:nn").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
			},
			"vlan_based_route_target_import": {
				MarkdownDescription: helpers.NewAttributeDescription("ASN:nn or IP-address:nn").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
			},
			"vlan_based_route_target_export": {
				MarkdownDescription: helpers.NewAttributeDescription("ASN:nn or IP-address:nn").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
			},
			"vlan_based_ip_local_learning_disable": {
				MarkdownDescription: helpers.NewAttributeDescription("Disable IP local learning from dataplane").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"vlan_based_ip_local_learning_enable": {
				MarkdownDescription: helpers.NewAttributeDescription("Enable IP local learning from dataplane").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"vlan_based_default_gateway_advertise": {
				MarkdownDescription: helpers.NewAttributeDescription("Advertise Default Gateway MAC/IP routes").AddStringEnumDescription("disable", "enable").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.StringEnumValidator("disable", "enable"),
				},
			},
			"vlan_based_re_originate_route_type5": {
				MarkdownDescription: helpers.NewAttributeDescription("Re-originate route-type 5").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
		},
	}, nil
}

func (t resourceEVPNInstanceType) NewResource(ctx context.Context, in tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	provider, diags := convertProviderType(in)

	return resourceEVPNInstance{
		provider: provider,
	}, diags
}

type resourceEVPNInstance struct {
	provider provider
}

func (r resourceEVPNInstance) Create(ctx context.Context, req tfsdk.CreateResourceRequest, resp *tfsdk.CreateResourceResponse) {
	var plan EVPNInstance

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.getPath()))

	// Create object
	body := plan.toBody(ctx)

	res, err := r.provider.clients[plan.Device.Value].PatchData(plan.getPathShort(), body)
	if len(res.Errors.Error) > 0 && res.Errors.Error[0].ErrorMessage == "patch to a nonexistent resource" {
		_, err = r.provider.clients[plan.Device.Value].PutData(plan.getPath(), body)
	}
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PATCH), got error: %s", err))
		return
	}

	emptyLeafsDelete := plan.getEmptyLeafsDelete(ctx)
	tflog.Debug(ctx, fmt.Sprintf("List of empty leafs to delete: %+v", emptyLeafsDelete))

	for _, i := range emptyLeafsDelete {
		res, err := r.provider.clients[plan.Device.Value].DeleteData(i)
		if err != nil && res.StatusCode != 404 {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object, got error: %s", err))
			return
		}
	}

	plan.setUnknownValues(ctx)

	plan.Id = types.String{Value: plan.getPath()}

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.getPath()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r resourceEVPNInstance) Read(ctx context.Context, req tfsdk.ReadResourceRequest, resp *tfsdk.ReadResourceResponse) {
	var state EVPNInstance

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.Value))

	res, err := r.provider.clients[state.Device.Value].GetData(state.Id.Value)
	if res.StatusCode == 404 {
		state = EVPNInstance{Device: state.Device, Id: state.Id}
	} else {
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
			return
		}

		state.updateFromBody(ctx, res.Res)
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.Value))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r resourceEVPNInstance) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {
	var plan, state EVPNInstance

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

	body := plan.toBody(ctx)
	res, err := r.provider.clients[plan.Device.Value].PatchData(plan.getPathShort(), body)
	if len(res.Errors.Error) > 0 && res.Errors.Error[0].ErrorMessage == "patch to a nonexistent resource" {
		_, err = r.provider.clients[plan.Device.Value].PutData(plan.getPath(), body)
	}
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PATCH), got error: %s", err))
		return
	}

	plan.setUnknownValues(ctx)

	deletedListItems := plan.getDeletedListItems(ctx, state)
	tflog.Debug(ctx, fmt.Sprintf("List items to delete: %+v", deletedListItems))

	for _, i := range deletedListItems {
		res, err := r.provider.clients[state.Device.Value].DeleteData(i)
		if err != nil && res.StatusCode != 404 {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object, got error: %s", err))
			return
		}
	}

	emptyLeafsDelete := plan.getEmptyLeafsDelete(ctx)
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

func (r resourceEVPNInstance) Delete(ctx context.Context, req tfsdk.DeleteResourceRequest, resp *tfsdk.DeleteResourceResponse) {
	var state EVPNInstance

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

func (r resourceEVPNInstance) ImportState(ctx context.Context, req tfsdk.ImportResourceStateRequest, resp *tfsdk.ImportResourceStateResponse) {
	tfsdk.ResourceImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
