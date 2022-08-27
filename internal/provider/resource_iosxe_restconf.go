package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/terraform-provider-iosxe/internal/provider/helpers"
)

type resourceRestconfType struct{}

func (t resourceRestconfType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Manages IOS-XE objects via RESTCONF calls. This resource can only manage a single object. It is able to read the state and therefore reconcile configuration drift.",

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
			"path": {
				MarkdownDescription: "A RESTCONF path, e.g. `openconfig-interfaces:interfaces`.",
				Type:                types.StringType,
				Required:            true,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					resource.RequiresReplace(),
				},
			},
			"delete": {
				MarkdownDescription: "Delete object during destroy operation. Default value is `true`.",
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					helpers.BooleanDefaultModifier(true),
				},
			},
			"attributes": {
				Type:                types.MapType{ElemType: types.StringType},
				MarkdownDescription: "Map of key-value pairs which represents the attributes and its values.",
				Optional:            true,
				Computed:            true,
			},
			"lists": {
				MarkdownDescription: "YANG lists.",
				Optional:            true,
				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
					"name": {
						MarkdownDescription: "YANG list name.",
						Type:                types.StringType,
						Required:            true,
					},
					"key": {
						MarkdownDescription: "YANG list key attribute. In case of multiple keys, those should be separated by a comma (`,`).",
						Type:                types.StringType,
						Optional:            true,
					},
					"items": {
						MarkdownDescription: "Items of YANG lists.",
						Optional:            true,
						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
							"attributes": {
								Type:                types.MapType{ElemType: types.StringType},
								MarkdownDescription: "Map of key-value pairs which represents the attributes and its values.",
								Optional:            true,
								Computed:            true,
							},
						}),
					},
					"values": {
						MarkdownDescription: "YANG leaf-list values.",
						Type:                types.ListType{ElemType: types.StringType},
						Optional:            true,
					},
				}),
			},
		},
	}, nil
}

func (t resourceRestconfType) NewResource(ctx context.Context, in provider.Provider) (resource.Resource, diag.Diagnostics) {
	provider, diags := convertProviderType(in)

	return resourceRestconf{
		provider: provider,
	}, diags
}

type resourceRestconf struct {
	provider iosxeProvider
}

func (r resourceRestconf) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan Restconf

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.getPath()))

	body := plan.toBody(ctx)
	res, err := r.provider.clients[plan.Device.Value].PatchData(plan.getPathShort(), body)
	if len(res.Errors.Error) > 0 && res.Errors.Error[0].ErrorMessage == "patch to a nonexistent resource" {
		_, err = r.provider.clients[plan.Device.Value].PutData(plan.getPath(), body)
	}
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PATCH), got error: %s", err))
		return
	}

	plan.Id = plan.Path
	plan.Attributes.Unknown = false

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.getPath()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r resourceRestconf) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state Restconf

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.getPath()))

	res, err := r.provider.clients[state.Device.Value].GetData(state.getPath())
	if res.StatusCode == 404 {
		state.Attributes.Elems = map[string]attr.Value{}
		state.Lists = make([]RestconfList, 0)
	} else {
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to read object, got error: %s", err))
			return
		}

		state.fromBody(ctx, res.Res)
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.getPath()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r resourceRestconf) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state Restconf

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

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.getPath()))

	body := plan.toBody(ctx)
	res, err := r.provider.clients[plan.Device.Value].PatchData(plan.getPathShort(), body)
	if len(res.Errors.Error) > 0 && res.Errors.Error[0].ErrorMessage == "patch to a nonexistent resource" {
		_, err = r.provider.clients[plan.Device.Value].PutData(plan.getPath(), body)
	}
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PATCH), got error: %s", err))
		return
	}

	deletedListItems := plan.getDeletedListItems(ctx, state)
	tflog.Debug(ctx, fmt.Sprintf("List items to delete: %+v", deletedListItems))

	for _, i := range deletedListItems {
		res, err := r.provider.clients[state.Device.Value].DeleteData(i)
		if err != nil && res.StatusCode != 404 {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object, got error: %s", err))
			return
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.getPath()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r resourceRestconf) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state Restconf

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.getPath()))

	if state.Delete.Value {
		res, err := r.provider.clients[state.Device.Value].DeleteData(state.getPath())
		if err != nil && res.StatusCode != 404 {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object, got error: %s", err))
			return
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.getPath()))

	resp.State.RemoveResource(ctx)
}

func (r resourceRestconf) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Import", req.ID))

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("path"), req.ID)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)

	tflog.Debug(ctx, fmt.Sprintf("%s: Import finished successfully", req.ID))
}
