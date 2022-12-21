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

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &PIMVRFResource{}
var _ resource.ResourceWithImportState = &PIMVRFResource{}

func NewPIMVRFResource() resource.Resource {
	return &PIMVRFResource{}
}

type PIMVRFResource struct {
	clients map[string]*restconf.Client
}

func (r *PIMVRFResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_pim_vrf"
}

func (r *PIMVRFResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This resource can manage the PIM VRF configuration.",

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
			"vrf": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"autorp": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure AutoRP global operations").String,
				Optional:            true,
				Computed:            true,
			},
			"autorp_listener": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Allow AutoRP packets across sparse mode interface").String,
				Optional:            true,
				Computed:            true,
			},
			"bsr_candidate_loopback": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Loopback interface").AddIntegerRangeDescription(0, 2147483647).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 2147483647),
				},
			},
			"bsr_candidate_mask": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Hash Mask length for RP selection").AddIntegerRangeDescription(0, 32).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 32),
				},
			},
			"bsr_candidate_priority": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Priority value for candidate bootstrap router").AddIntegerRangeDescription(0, 255).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 255),
				},
			},
			"bsr_candidate_accept_rp_candidate": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("BSR RP candidate filter").String,
				Optional:            true,
				Computed:            true,
			},
			"ssm_range": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("ACL for group range to be used for SSM").String,
				Optional:            true,
				Computed:            true,
			},
			"ssm_default": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Use 232/8 group range for SSM").String,
				Optional:            true,
				Computed:            true,
			},
			"rp_address": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IP address of Rendezvous-point for group").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile(`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?`), ""),
				},
			},
			"rp_address_override": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Overrides dynamically learnt RP mappings").String,
				Optional:            true,
				Computed:            true,
			},
			"rp_address_bidir": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Group range treated in bidirectional shared-tree mode").String,
				Optional:            true,
				Computed:            true,
			},
			"rp_addresses": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("PIM RP-address (Rendezvous Point)").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"access_list": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IP Access-list").String,
							Optional:            true,
							Computed:            true,
						},
						"rp_address": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IP address of Rendezvous-point for group").String,
							Optional:            true,
							Computed:            true,
							Validators: []validator.String{
								stringvalidator.RegexMatches(regexp.MustCompile(`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?`), ""),
							},
						},
						"override": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Overrides dynamically learnt RP mappings").String,
							Optional:            true,
							Computed:            true,
						},
						"bidir": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Group range treated in bidirectional shared-tree mode").String,
							Optional:            true,
							Computed:            true,
						},
					},
				},
			},
			"rp_candidates": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("To be a PIM version 2 RP candidate").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"interface": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Autonomic-Networking virtual interface").String,
							Optional:            true,
							Computed:            true,
						},
						"group_list": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IP Access list").String,
							Optional:            true,
							Computed:            true,
						},
						"interval": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("RP candidate advertisement interval").AddIntegerRangeDescription(1, 16383).String,
							Optional:            true,
							Computed:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 16383),
							},
						},
						"priority": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("RP candidate priority").AddIntegerRangeDescription(0, 255).String,
							Optional:            true,
							Computed:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 255),
							},
						},
						"bidir": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Group range treated in bidirectional shared-tree mode").String,
							Optional:            true,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (r *PIMVRFResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.clients = req.ProviderData.(map[string]*restconf.Client)
}

func (r *PIMVRFResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan PIMVRF

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

func (r *PIMVRFResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state PIMVRF

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.ValueString()))

	res, err := r.clients[state.Device.ValueString()].GetData(state.Id.ValueString())
	if res.StatusCode == 404 {
		state = PIMVRF{Device: state.Device, Id: state.Id}
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

func (r *PIMVRFResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state PIMVRF

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

func (r *PIMVRFResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state PIMVRF

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

func (r *PIMVRFResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
