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
var _ resource.Resource = &ServiceResource{}
var _ resource.ResourceWithImportState = &ServiceResource{}

func NewServiceResource() resource.Resource {
	return &ServiceResource{}
}

type ServiceResource struct {
	clients map[string]*restconf.Client
}

func (r *ServiceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_service"
}

func (r *ServiceResource) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This resource can manage the Service configuration.",

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
			"pad": {
				MarkdownDescription: helpers.NewAttributeDescription("Enable PAD commands").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"password_encryption": {
				MarkdownDescription: helpers.NewAttributeDescription("Encrypt system passwords").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"password_recovery": {
				MarkdownDescription: helpers.NewAttributeDescription("Enable password recovery").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"timestamps": {
				MarkdownDescription: helpers.NewAttributeDescription("Timestamp debug/log messages").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"timestamps_debug": {
				MarkdownDescription: helpers.NewAttributeDescription("Timestamp debug messages").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"timestamps_debug_datetime": {
				MarkdownDescription: helpers.NewAttributeDescription("Timestamp with date and time").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"timestamps_debug_datetime_msec": {
				MarkdownDescription: helpers.NewAttributeDescription("Include milliseconds in timestamp").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"timestamps_debug_datetime_localtime": {
				MarkdownDescription: helpers.NewAttributeDescription("Use local time zone for timestamps").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"timestamps_debug_datetime_show_timezone": {
				MarkdownDescription: helpers.NewAttributeDescription("Add time zone information to timestamp").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"timestamps_debug_datetime_year": {
				MarkdownDescription: helpers.NewAttributeDescription("Include year in timestamp").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"timestamps_debug_uptime": {
				MarkdownDescription: helpers.NewAttributeDescription("Timestamp with system uptime").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"timestamps_log": {
				MarkdownDescription: helpers.NewAttributeDescription("Timestamp log messages").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"timestamps_log_datetime": {
				MarkdownDescription: helpers.NewAttributeDescription("Timestamp with date and time").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"timestamps_log_datetime_msec": {
				MarkdownDescription: helpers.NewAttributeDescription("Include milliseconds in timestamp").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"timestamps_log_datetime_localtime": {
				MarkdownDescription: helpers.NewAttributeDescription("Use local time zone for timestamps").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"timestamps_log_datetime_show_timezone": {
				MarkdownDescription: helpers.NewAttributeDescription("Add time zone information to timestamp").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"timestamps_log_datetime_year": {
				MarkdownDescription: helpers.NewAttributeDescription("Include year in timestamp").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"timestamps_log_uptime": {
				MarkdownDescription: helpers.NewAttributeDescription("Timestamp with system uptime").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"dhcp": {
				MarkdownDescription: helpers.NewAttributeDescription("Enable DHCP server and relay agent").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"tcp_keepalives_in": {
				MarkdownDescription: helpers.NewAttributeDescription("Generate keepalives on idle incoming network connections").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"tcp_keepalives_out": {
				MarkdownDescription: helpers.NewAttributeDescription("Generate keepalives on idle outgoing network connections").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
		},
	}, nil
}

func (r *ServiceResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.clients = req.ProviderData.(map[string]*restconf.Client)
}

func (r *ServiceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan Service

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

func (r *ServiceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state Service

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.ValueString()))

	res, err := r.clients[state.Device.ValueString()].GetData(state.Id.ValueString())
	if res.StatusCode == 404 {
		state = Service{Device: state.Device, Id: state.Id}
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

func (r *ServiceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state Service

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

func (r *ServiceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state Service

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))

	res, err := r.clients[state.Device.ValueString()].DeleteData(state.Id.ValueString())
	if err != nil && res.StatusCode != 404 && res.StatusCode != 400 {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object, got error: %s", err))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

func (r *ServiceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
