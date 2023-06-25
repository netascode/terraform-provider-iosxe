// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-restconf"
	"github.com/netascode/terraform-provider-iosxe/internal/provider/helpers"
)

func NewAccessListExtendedResource() resource.Resource {
	return &AccessListExtendedResource{}
}

type AccessListExtendedResource struct {
	clients map[string]*restconf.Client
}

func (r *AccessListExtendedResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_access_list_extended"
}

func (r *AccessListExtendedResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This resource can manage the Access List Extended configuration.",

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The path of the object.",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"entries": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"sequence": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("").AddIntegerRangeDescription(1, 2147483647).String,
							Required:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 2147483647),
							},
						},
						"remark": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Access list entry comment").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 100),
							},
						},
						"ace_rule_action": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("").AddStringEnumDescription("deny", "permit").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("deny", "permit"),
							},
						},
						"ace_rule_protocol": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("").String,
							Optional:            true,
						},
						"service_object_group": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Service object group name").String,
							Optional:            true,
						},
						"source_prefix": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.RegexMatches(regexp.MustCompile(`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?`), ""),
							},
						},
						"source_prefix_mask": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.RegexMatches(regexp.MustCompile(`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?`), ""),
							},
						},
						"source_any": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Any source host").String,
							Optional:            true,
						},
						"source_host": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("A single source host").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.RegexMatches(regexp.MustCompile(`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?`), ""),
							},
						},
						"source_object_group": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Source network object group").String,
							Optional:            true,
						},
						"source_port_equal": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Match only packets on a given port number up to 10 ports").String,
							Optional:            true,
						},
						"source_port_greater_than": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Match only packets with a greater port number").String,
							Optional:            true,
						},
						"source_port_lesser_than": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Match only packets with a lower port number").String,
							Optional:            true,
						},
						"source_port_range_from": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Match only packets in the range of port numbers").String,
							Optional:            true,
						},
						"source_port_range_to": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Match only packets in the range of port numbers").String,
							Optional:            true,
						},
						"destination_prefix": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.RegexMatches(regexp.MustCompile(`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?`), ""),
							},
						},
						"destination_prefix_mask": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.RegexMatches(regexp.MustCompile(`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?`), ""),
							},
						},
						"destination_any": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Any destination host").String,
							Optional:            true,
						},
						"destination_host": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("A single destination host").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.RegexMatches(regexp.MustCompile(`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?`), ""),
							},
						},
						"destination_object_group": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Destination network object group").String,
							Optional:            true,
						},
						"destination_port_equal": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Match only packets on a given port number up to 10 ports").String,
							Optional:            true,
						},
						"destination_port_greater_than": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Match only packets with a greater port number").String,
							Optional:            true,
						},
						"destination_port_lesser_than": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Match only packets with a lower port number").String,
							Optional:            true,
						},
						"destination_port_range_from": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Match only packets in the range of port numbers").String,
							Optional:            true,
						},
						"destination_port_range_to": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Match only packets in the range of port numbers").String,
							Optional:            true,
						},
						"ack": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Match on the ACK bit").String,
							Optional:            true,
						},
						"fin": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Match on the FIN bit").String,
							Optional:            true,
						},
						"psh": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Match on the PSH bit").String,
							Optional:            true,
						},
						"rst": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Match on the RST bit").String,
							Optional:            true,
						},
						"syn": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Match on the SYN bit").String,
							Optional:            true,
						},
						"urg": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Match on the URG bit").String,
							Optional:            true,
						},
						"established": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Match established connections").String,
							Optional:            true,
						},
						"dscp": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Match packets with given dscp value").String,
							Optional:            true,
						},
						"fragments": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Check non-initial fragments").String,
							Optional:            true,
						},
						"precedence": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Match packets with given precedence value").String,
							Optional:            true,
						},
						"tos": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Match packets with given TOS value").String,
							Optional:            true,
						},
					},
				},
			},
		},
	}
}

func (r *AccessListExtendedResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.clients = req.ProviderData.(map[string]*restconf.Client)
}

func (r *AccessListExtendedResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan AccessListExtended

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if _, ok := r.clients[plan.Device.ValueString()]; !ok {
		resp.Diagnostics.AddAttributeError(path.Root("device"), "Invalid device", fmt.Sprintf("Device '%s' does not exist in provider configuration.", plan.Device.ValueString()))
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

	plan.Id = types.StringValue(plan.getPath())

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.getPath()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *AccessListExtendedResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state AccessListExtended

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if _, ok := r.clients[state.Device.ValueString()]; !ok {
		resp.Diagnostics.AddAttributeError(path.Root("device"), "Invalid device", fmt.Sprintf("Device '%s' does not exist in provider configuration.", state.Device.ValueString()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.ValueString()))

	res, err := r.clients[state.Device.ValueString()].GetData(state.Id.ValueString())
	if res.StatusCode == 404 {
		state = AccessListExtended{Device: state.Device, Id: state.Id}
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

func (r *AccessListExtendedResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state AccessListExtended

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

	if _, ok := r.clients[plan.Device.ValueString()]; !ok {
		resp.Diagnostics.AddAttributeError(path.Root("device"), "Invalid device", fmt.Sprintf("Device '%s' does not exist in provider configuration.", plan.Device.ValueString()))
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

func (r *AccessListExtendedResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state AccessListExtended

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if _, ok := r.clients[state.Device.ValueString()]; !ok {
		resp.Diagnostics.AddAttributeError(path.Root("device"), "Invalid device", fmt.Sprintf("Device '%s' does not exist in provider configuration.", state.Device.ValueString()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))
	deleteMode := "all"

	if deleteMode == "all" {
		res, err := r.clients[state.Device.ValueString()].DeleteData(state.Id.ValueString())
		if err != nil && res.StatusCode != 404 && res.StatusCode != 400 {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object, got error: %s", err))
			return
		}
	} else {
		deletePaths := state.getDeletePaths(ctx)
		tflog.Debug(ctx, fmt.Sprintf("Paths to delete: %+v", deletePaths))

		for _, i := range deletePaths {
			res, err := r.clients[state.Device.ValueString()].DeleteData(i)
			if err != nil && res.StatusCode != 404 {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object, got error: %s", err))
				return
			}
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

func (r *AccessListExtendedResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
