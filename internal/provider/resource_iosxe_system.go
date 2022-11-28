// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-restconf"
	"github.com/netascode/terraform-provider-iosxe/internal/provider/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &SystemResource{}
var _ resource.ResourceWithImportState = &SystemResource{}

func NewSystemResource() resource.Resource {
	return &SystemResource{}
}

type SystemResource struct {
	clients map[string]*restconf.Client
}

func (r *SystemResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_system"
}

func (r *SystemResource) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This resource can manage the System configuration.",

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
					resource.UseStateForUnknown(),
				},
			},
			"hostname": {
				MarkdownDescription: helpers.NewAttributeDescription("Set system's network name").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.StringPatternValidator(1, 246, `([0-9.]*[A-Za-z\-_]+[A-Za-z0-9.\-_]*)`),
				},
			},
			"ip_routing": {
				MarkdownDescription: helpers.NewAttributeDescription("Enable or disable IP routing").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"ipv6_unicast_routing": {
				MarkdownDescription: helpers.NewAttributeDescription("Enable unicast routing").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"mtu": {
				MarkdownDescription: helpers.NewAttributeDescription("").AddIntegerRangeDescription(1500, 9198).String,
				Type:                types.Int64Type,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.IntegerRangeValidator(1500, 9198),
				},
			},
			"ip_source_route": {
				MarkdownDescription: helpers.NewAttributeDescription("Process packets with source routing header options").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"ip_domain_lookup": {
				MarkdownDescription: helpers.NewAttributeDescription("Enable IP Domain Name System hostname translation").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"ip_domain_name": {
				MarkdownDescription: helpers.NewAttributeDescription("Define the default domain name").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
			},
			"login_delay": {
				MarkdownDescription: helpers.NewAttributeDescription("Set delay between successive fail login").AddIntegerRangeDescription(1, 10).String,
				Type:                types.Int64Type,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.IntegerRangeValidator(1, 10),
				},
			},
			"login_on_failure": {
				MarkdownDescription: helpers.NewAttributeDescription("Set options for failed login attempt").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"login_on_failure_log": {
				MarkdownDescription: helpers.NewAttributeDescription("Generate syslogs on failure logins").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"login_on_success": {
				MarkdownDescription: helpers.NewAttributeDescription("Set options for successful login attempt").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"login_on_success_log": {
				MarkdownDescription: helpers.NewAttributeDescription("Generate syslogs on successful logins").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"multicast_routing": {
				MarkdownDescription: helpers.NewAttributeDescription("Enable IP multicast forwarding").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"multicast_routing_switch": {
				MarkdownDescription: helpers.NewAttributeDescription("Enable IP multicast forwarding, some XE devices use this option instead of `multicast_routing`.").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"multicast_routing_distributed": {
				MarkdownDescription: helpers.NewAttributeDescription("Distributed multicast switching").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"multicast_routing_vrfs": {
				MarkdownDescription: helpers.NewAttributeDescription("Select VPN Routing/Forwarding instance").String,
				Optional:            true,
				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
					"vrf": {
						MarkdownDescription: helpers.NewAttributeDescription("").String,
						Type:                types.StringType,
						Optional:            true,
						Computed:            true,
					},
					"distributed": {
						MarkdownDescription: helpers.NewAttributeDescription("Distributed multicast switching").String,
						Type:                types.BoolType,
						Optional:            true,
						Computed:            true,
					},
				}),
			},
		},
	}, nil
}

func (r *SystemResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.clients = req.ProviderData.(map[string]*restconf.Client)
}

func (r *SystemResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan System

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.getPath()))

	// Create object
	body := plan.toBody(ctx)

	res, err := r.clients[plan.Device.ValueString()].PatchData(plan.getPathShort(), body)
	if len(res.Errors.Error) > 0 && res.Errors.Error[0].ErrorMessage == "patch to a nonexistent resource" {
		_, err = r.clients[plan.Device.ValueString()].PutData(plan.getPath(), body)
	}
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PATCH), got error: %s", err))
		return
	}

	emptyLeafsDelete := plan.getEmptyLeafsDelete(ctx)
	tflog.Debug(ctx, fmt.Sprintf("List of empty leafs to delete: %+v", emptyLeafsDelete))

	for _, i := range emptyLeafsDelete {
		res, err := r.clients[plan.Device.ValueString()].DeleteData(i)
		if err != nil && res.StatusCode != 404 {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object, got error: %s", err))
			return
		}
	}

	plan.setUnknownValues(ctx)

	plan.Id = types.StringValue(plan.getPath())

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.getPath()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *SystemResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state System

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.ValueString()))

	res, err := r.clients[state.Device.ValueString()].GetData(state.Id.ValueString())
	if res.StatusCode == 404 {
		state = System{Device: state.Device, Id: state.Id}
	} else {
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
			return
		}

		state.updateFromBody(ctx, res.Res)
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r *SystemResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state System

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

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.Id.ValueString()))

	body := plan.toBody(ctx)
	res, err := r.clients[plan.Device.ValueString()].PatchData(plan.getPathShort(), body)
	if len(res.Errors.Error) > 0 && res.Errors.Error[0].ErrorMessage == "patch to a nonexistent resource" {
		_, err = r.clients[plan.Device.ValueString()].PutData(plan.getPath(), body)
	}
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PATCH), got error: %s", err))
		return
	}

	plan.setUnknownValues(ctx)

	deletedListItems := plan.getDeletedListItems(ctx, state)
	tflog.Debug(ctx, fmt.Sprintf("List items to delete: %+v", deletedListItems))

	for _, i := range deletedListItems {
		res, err := r.clients[state.Device.ValueString()].DeleteData(i)
		if err != nil && res.StatusCode != 404 {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object, got error: %s", err))
			return
		}
	}

	emptyLeafsDelete := plan.getEmptyLeafsDelete(ctx)
	tflog.Debug(ctx, fmt.Sprintf("List of empty leafs to delete: %+v", emptyLeafsDelete))

	for _, i := range emptyLeafsDelete {
		res, err := r.clients[plan.Device.ValueString()].DeleteData(i)
		if err != nil && res.StatusCode != 404 {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object, got error: %s", err))
			return
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *SystemResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state System

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

func (r *SystemResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
