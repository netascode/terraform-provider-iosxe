// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

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
var _ resource.Resource = &SNMPServerResource{}
var _ resource.ResourceWithImportState = &SNMPServerResource{}

func NewSNMPServerResource() resource.Resource {
	return &SNMPServerResource{}
}

type SNMPServerResource struct {
	clients map[string]*restconf.Client
}

func (r *SNMPServerResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_snmp_server"
}

func (r *SNMPServerResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This resource can manage the SNMP Server configuration.",

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
			"chassis_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("String to uniquely identify this chassis").String,
				Optional:            true,
				Computed:            true,
			},
			"contact": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Text for mib object sysContact").String,
				Optional:            true,
				Computed:            true,
			},
			"ifindex_persist": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Persist interface indices").String,
				Optional:            true,
				Computed:            true,
			},
			"location": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Text for mib object sysLocation").String,
				Optional:            true,
				Computed:            true,
			},
			"packetsize": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Largest SNMP packet size").AddIntegerRangeDescription(484, 17892).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(484, 17892),
				},
			},
			"queue_length": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Message queue length for each TRAP host").AddIntegerRangeDescription(1, 5000).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 5000),
				},
			},
			"enable_logging_getop": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable SNMP GET Operation logging").String,
				Optional:            true,
				Computed:            true,
			},
			"enable_logging_setop": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable SNMP SET Operation logging").String,
				Optional:            true,
				Computed:            true,
			},
			"enable_informs": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable SNMP Informs").String,
				Optional:            true,
				Computed:            true,
			},
			"enable_traps": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable SNMP Traps").String,
				Optional:            true,
				Computed:            true,
			},
			"enable_traps_snmp_authentication": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable authentication trap").String,
				Optional:            true,
				Computed:            true,
			},
			"enable_traps_snmp_coldstart": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable coldStart trap").String,
				Optional:            true,
				Computed:            true,
			},
			"enable_traps_snmp_linkdown": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable linkDown trap").String,
				Optional:            true,
				Computed:            true,
			},
			"enable_traps_snmp_linkup": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable linkUp trap").String,
				Optional:            true,
				Computed:            true,
			},
			"enable_traps_snmp_warmstart": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable warmStart trap").String,
				Optional:            true,
				Computed:            true,
			},
			"source_interface_informs_gigabit_ethernet": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("GigabitEthernet IEEE 802.3z").String,
				Optional:            true,
				Computed:            true,
			},
			"source_interface_informs_ten_gigabit_ethernet": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Ten Gigabit Ethernet").String,
				Optional:            true,
				Computed:            true,
			},
			"source_interface_informs_forty_gigabit_ethernet": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Forty GigabitEthernet ").String,
				Optional:            true,
				Computed:            true,
			},
			"source_interface_informs_hundred_gig_e": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Hundred GigabitEthernet").String,
				Optional:            true,
				Computed:            true,
			},
			"source_interface_informs_loopback": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Loopback interface").AddIntegerRangeDescription(0, 2147483647).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 2147483647),
				},
			},
			"source_interface_informs_port_channel": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Ethernet Channel of interfaces").AddIntegerRangeDescription(0, 4294967295).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 4294967295),
				},
			},
			"source_interface_informs_port_channel_subinterface": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Optional:            true,
				Computed:            true,
			},
			"source_interface_informs_vlan": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Iosxr Vlans").AddIntegerRangeDescription(0, 65535).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 65535),
				},
			},
			"source_interface_traps_gigabit_ethernet": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("GigabitEthernet IEEE 802.3z").String,
				Optional:            true,
				Computed:            true,
			},
			"source_interface_traps_ten_gigabit_ethernet": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Ten Gigabit Ethernet").String,
				Optional:            true,
				Computed:            true,
			},
			"source_interface_traps_forty_gigabit_ethernet": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Forty GigabitEthernet ").String,
				Optional:            true,
				Computed:            true,
			},
			"source_interface_traps_hundred_gig_e": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Hundred GigabitEthernet").String,
				Optional:            true,
				Computed:            true,
			},
			"source_interface_traps_loopback": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Loopback interface").AddIntegerRangeDescription(0, 2147483647).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 2147483647),
				},
			},
			"source_interface_traps_port_channel": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Ethernet Channel of interfaces").AddIntegerRangeDescription(0, 4294967295).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 4294967295),
				},
			},
			"source_interface_traps_port_channel_subinterface": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Optional:            true,
				Computed:            true,
			},
			"source_interface_traps_vlan": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Iosxr Vlans").AddIntegerRangeDescription(0, 65535).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 65535),
				},
			},
			"trap_source_gigabit_ethernet": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("GigabitEthernet IEEE 802.3z").String,
				Optional:            true,
				Computed:            true,
			},
			"trap_source_ten_gigabit_ethernet": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Ten Gigabit Ethernet").String,
				Optional:            true,
				Computed:            true,
			},
			"trap_source_forty_gigabit_ethernet": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Forty GigabitEthernet ").String,
				Optional:            true,
				Computed:            true,
			},
			"trap_source_hundred_gig_e": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Hundred GigabitEthernet").String,
				Optional:            true,
				Computed:            true,
			},
			"trap_source_loopback": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Loopback interface").AddIntegerRangeDescription(0, 2147483647).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 2147483647),
				},
			},
			"trap_source_port_channel": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Ethernet Channel of interfaces").AddIntegerRangeDescription(0, 4294967295).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 4294967295),
				},
			},
			"trap_source_port_channel_subinterface": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Optional:            true,
				Computed:            true,
			},
			"trap_source_vlan": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Iosxr Vlans").AddIntegerRangeDescription(0, 65535).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 65535),
				},
			},
			"snmp_communities": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable SNMP; set community string and access privs").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("").String,
							Optional:            true,
							Computed:            true,
						},
						"view": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Restrict this community to a named MIB view").String,
							Optional:            true,
							Computed:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 214),
							},
						},
						"permission": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("").AddStringEnumDescription("ro", "rw").String,
							Optional:            true,
							Computed:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("ro", "rw"),
							},
						},
						"ipv6": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Specify IPv6 Named Access-List").String,
							Optional:            true,
							Computed:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 219),
							},
						},
						"access_list_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Access-list name").String,
							Optional:            true,
							Computed:            true,
						},
					},
				},
			},
			"contexts": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Create/Delete a context apart from default").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("").String,
							Optional:            true,
							Computed:            true,
						},
					},
				},
			},
			"views": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Define an SNMPv2 MIB view").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("").String,
							Optional:            true,
							Computed:            true,
						},
						"mib": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("").String,
							Optional:            true,
							Computed:            true,
						},
						"inc_exl": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("").AddStringEnumDescription("excluded", "included").String,
							Optional:            true,
							Computed:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("excluded", "included"),
							},
						},
					},
				},
			},
		},
	}
}

func (r *SNMPServerResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.clients = req.ProviderData.(map[string]*restconf.Client)
}

func (r *SNMPServerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan SNMPServer

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

func (r *SNMPServerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state SNMPServer

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.ValueString()))

	res, err := r.clients[state.Device.ValueString()].GetData(state.Id.ValueString())
	if res.StatusCode == 404 {
		state = SNMPServer{Device: state.Device, Id: state.Id}
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

func (r *SNMPServerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state SNMPServer

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

func (r *SNMPServerResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state SNMPServer

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

func (r *SNMPServerResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
