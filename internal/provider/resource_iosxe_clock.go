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

func NewClockResource() resource.Resource {
	return &ClockResource{}
}

type ClockResource struct {
	clients map[string]*restconf.Client
}

func (r *ClockResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_clock"
}

func (r *ClockResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This resource can manage the Clock configuration.",

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
			"delete_mode": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure behavior when deleting/destroying the resource. Either delete the entire object (YANG container) being managed, or only delete the individual resource attributes configured explicitly and leave everything else as-is. Default value is `all`.").AddStringEnumDescription("all", "attributes").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("all", "attributes"),
				},
			},
			"calendar_valid": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Calendar time is authoritative").String,
				Optional:            true,
			},
			"summer_time_zone": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Name of time zone in summer").String,
				Optional:            true,
			},
			"summer_time_date": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure absolute summer time").String,
				Optional:            true,
			},
			"summer_time_date_start_day": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("").AddIntegerRangeDescription(1, 31).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 31),
				},
			},
			"summer_time_date_start_month": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").AddStringEnumDescription("Apr", "Aug", "Dec", "Feb", "Jan", "Jul", "Jun", "Mar", "May", "Nov", "Oct", "Sep").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("Apr", "Aug", "Dec", "Feb", "Jan", "Jul", "Jun", "Mar", "May", "Nov", "Oct", "Sep"),
				},
			},
			"summer_time_date_start_year": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("").AddIntegerRangeDescription(1993, 2035).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1993, 2035),
				},
			},
			"summer_time_date_start_time": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile(`([0-1]?[0-9]|2[0-4]):([0-5][0-9])(:[0-5][0-9])?`), ""),
				},
			},
			"summer_time_date_end_day": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("").AddIntegerRangeDescription(1, 31).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 31),
				},
			},
			"summer_time_date_end_month": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").AddStringEnumDescription("Apr", "Aug", "Dec", "Feb", "Jan", "Jul", "Jun", "Mar", "May", "Nov", "Oct", "Sep").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("Apr", "Aug", "Dec", "Feb", "Jan", "Jul", "Jun", "Mar", "May", "Nov", "Oct", "Sep"),
				},
			},
			"summer_time_date_end_year": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("").AddIntegerRangeDescription(1993, 2035).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1993, 2035),
				},
			},
			"summer_time_date_end_time": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile(`([0-1]?[0-9]|2[0-4]):([0-5][0-9])(:[0-5][0-9])?`), ""),
				},
			},
			"summer_time_date_offset": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("").AddIntegerRangeDescription(1, 1440).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 1440),
				},
			},
			"summer_time_recurring": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure recurring summer time").String,
				Optional:            true,
			},
			"summer_time_recurring_start_week": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Optional:            true,
			},
			"summer_time_recurring_start_weekday": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").AddStringEnumDescription("Fri", "Mon", "Sat", "Sun", "Thu", "Tue", "Wed").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("Fri", "Mon", "Sat", "Sun", "Thu", "Tue", "Wed"),
				},
			},
			"summer_time_recurring_start_month": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").AddStringEnumDescription("Apr", "Aug", "Dec", "Feb", "Jan", "Jul", "Jun", "Mar", "May", "Nov", "Oct", "Sep").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("Apr", "Aug", "Dec", "Feb", "Jan", "Jul", "Jun", "Mar", "May", "Nov", "Oct", "Sep"),
				},
			},
			"summer_time_recurring_start_time": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile(`([0-1]?[0-9]|2[0-4]):([0-5][0-9])(:[0-5][0-9])?`), ""),
				},
			},
			"summer_time_recurring_end_week": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Optional:            true,
			},
			"summer_time_recurring_end_weekday": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").AddStringEnumDescription("Fri", "Mon", "Sat", "Sun", "Thu", "Tue", "Wed").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("Fri", "Mon", "Sat", "Sun", "Thu", "Tue", "Wed"),
				},
			},
			"summer_time_recurring_end_month": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").AddStringEnumDescription("Apr", "Aug", "Dec", "Feb", "Jan", "Jul", "Jun", "Mar", "May", "Nov", "Oct", "Sep").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("Apr", "Aug", "Dec", "Feb", "Jan", "Jul", "Jun", "Mar", "May", "Nov", "Oct", "Sep"),
				},
			},
			"summer_time_recurring_end_time": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile(`([0-1]?[0-9]|2[0-4]):([0-5][0-9])(:[0-5][0-9])?`), ""),
				},
			},
			"summer_time_recurring_offset": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("").AddIntegerRangeDescription(1, 1440).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 1440),
				},
			},
		},
	}
}

func (r *ClockResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.clients = req.ProviderData.(map[string]*restconf.Client)
}

func (r *ClockResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan Clock

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

	emptyLeafsDelete := plan.getEmptyLeafsDelete(ctx)
	tflog.Debug(ctx, fmt.Sprintf("List of empty leafs to delete: %+v", emptyLeafsDelete))

	if YangPatch {
		edits := []restconf.YangPatchEdit{restconf.NewYangPatchEdit("merge", plan.getPath(), restconf.Body{Str: body})}
		for _, i := range emptyLeafsDelete {
			edits = append(edits, restconf.NewYangPatchEdit("remove", i, restconf.Body{}))
		}
		_, err := r.clients[plan.Device.ValueString()].YangPatchData("", "1", "", edits)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object, got error: %s", err))
			return
		}
	} else {
		res, err := r.clients[plan.Device.ValueString()].PatchData(plan.getPathShort(), body)
		if len(res.Errors.Error) > 0 && res.Errors.Error[0].ErrorMessage == "patch to a nonexistent resource" {
			_, err = r.clients[plan.Device.ValueString()].PutData(plan.getPath(), body)
		}
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PATCH), got error: %s", err))
			return
		}
		for _, i := range emptyLeafsDelete {
			res, err := r.clients[plan.Device.ValueString()].DeleteData(i)
			if err != nil && res.StatusCode != 404 {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object, got error: %s", err))
				return
			}
		}
	}

	plan.Id = types.StringValue(plan.getPath())

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.getPath()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *ClockResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state Clock

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
		state = Clock{Device: state.Device, Id: state.Id}
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

func (r *ClockResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state Clock

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

	deletedListItems := plan.getDeletedListItems(ctx, state)
	tflog.Debug(ctx, fmt.Sprintf("List items to delete: %+v", deletedListItems))

	emptyLeafsDelete := plan.getEmptyLeafsDelete(ctx)
	tflog.Debug(ctx, fmt.Sprintf("List of empty leafs to delete: %+v", emptyLeafsDelete))

	if YangPatch {
		edits := []restconf.YangPatchEdit{restconf.NewYangPatchEdit("merge", plan.getPath(), restconf.Body{Str: body})}
		for _, i := range deletedListItems {
			edits = append(edits, restconf.NewYangPatchEdit("remove", i, restconf.Body{}))
		}
		for _, i := range emptyLeafsDelete {
			edits = append(edits, restconf.NewYangPatchEdit("remove", i, restconf.Body{}))
		}
		_, err := r.clients[plan.Device.ValueString()].YangPatchData("", "1", "", edits)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object, got error: %s", err))
			return
		}
	} else {
		res, err := r.clients[plan.Device.ValueString()].PatchData(plan.getPathShort(), body)
		if len(res.Errors.Error) > 0 && res.Errors.Error[0].ErrorMessage == "patch to a nonexistent resource" {
			_, err = r.clients[plan.Device.ValueString()].PutData(plan.getPath(), body)
		}
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PATCH), got error: %s", err))
			return
		}
		for _, i := range deletedListItems {
			res, err := r.clients[state.Device.ValueString()].DeleteData(i)
			if err != nil && res.StatusCode != 404 {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object, got error: %s", err))
				return
			}
		}
		for _, i := range emptyLeafsDelete {
			res, err := r.clients[plan.Device.ValueString()].DeleteData(i)
			if err != nil && res.StatusCode != 404 {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object, got error: %s", err))
				return
			}
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *ClockResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state Clock

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
	if state.DeleteMode.ValueString() == "all" {
		deleteMode = "all"
	} else if state.DeleteMode.ValueString() == "attributes" {
		deleteMode = "attributes"
	}

	if deleteMode == "all" {
		res, err := r.clients[state.Device.ValueString()].DeleteData(state.Id.ValueString())
		if err != nil && res.StatusCode != 404 && res.StatusCode != 400 {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object, got error: %s", err))
			return
		}
	} else {
		deletePaths := state.getDeletePaths(ctx)
		tflog.Debug(ctx, fmt.Sprintf("Paths to delete: %+v", deletePaths))

		if YangPatch {
			edits := []restconf.YangPatchEdit{}
			for _, i := range deletePaths {
				edits = append(edits, restconf.NewYangPatchEdit("remove", i, restconf.Body{}))
			}
			_, err := r.clients[state.Device.ValueString()].YangPatchData("", "1", "", edits)
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object, got error: %s", err))
				return
			}
		} else {
			for _, i := range deletePaths {
				res, err := r.clients[state.Device.ValueString()].DeleteData(i)
				if err != nil && res.StatusCode != 404 {
					resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object, got error: %s", err))
					return
				}
			}
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

func (r *ClockResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
