// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/terraform-provider-iosxe/internal/provider/helpers"
)

type resourceTemplateType struct{}

func (t resourceTemplateType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This resource can manage the Template configuration.",

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
			"template_name": {
				MarkdownDescription: helpers.NewAttributeDescription("Template name").String,
				Type:                types.StringType,
				Required:            true,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					resource.RequiresReplace(),
				},
			},
			"dot1x_pae": {
				MarkdownDescription: helpers.NewAttributeDescription("Set 802.1x interface pae type").AddStringEnumDescription("authenticator", "both", "supplicant").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.StringEnumValidator("authenticator", "both", "supplicant"),
				},
			},
			"dot1x_max_reauth_req": {
				MarkdownDescription: helpers.NewAttributeDescription("Max No. of Reauthentication Attempts").AddIntegerRangeDescription(1, 10).String,
				Type:                types.Int64Type,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.IntegerRangeValidator(1, 10),
				},
			},
			"dot1x_max_req": {
				MarkdownDescription: helpers.NewAttributeDescription("Max No. of Retries").AddIntegerRangeDescription(1, 10).String,
				Type:                types.Int64Type,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.IntegerRangeValidator(1, 10),
				},
			},
			"service_policy_input": {
				MarkdownDescription: helpers.NewAttributeDescription("policy-map name").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
			},
			"service_policy_output": {
				MarkdownDescription: helpers.NewAttributeDescription("policy-map name").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
			},
			"source_template": {
				MarkdownDescription: helpers.NewAttributeDescription("Get config from a template").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
			},
			"switchport_mode_trunk": {
				MarkdownDescription: helpers.NewAttributeDescription("Set trunking mode to TRUNK unconditionally").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"switchport_mode_access": {
				MarkdownDescription: helpers.NewAttributeDescription("Set trunking mode to ACCESS unconditionally").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"switchport_nonegotiate": {
				MarkdownDescription: helpers.NewAttributeDescription("Device will not engage in negotiation protocol on this interface").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"switchport_block_unicast": {
				MarkdownDescription: helpers.NewAttributeDescription("Block unknown unicast addresses").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"switchport_port_security": {
				MarkdownDescription: helpers.NewAttributeDescription("Security related command").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"switchport_port_security_aging_static": {
				MarkdownDescription: helpers.NewAttributeDescription("Enable aging for configured secure addresses").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"switchport_port_security_aging_time": {
				MarkdownDescription: helpers.NewAttributeDescription("Port-security aging time").AddIntegerRangeDescription(1, 1440).String,
				Type:                types.Int64Type,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.IntegerRangeValidator(1, 1440),
				},
			},
			"switchport_port_security_aging_type": {
				MarkdownDescription: helpers.NewAttributeDescription("Port-security aging type").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"switchport_port_security_aging_type_inactivity": {
				MarkdownDescription: helpers.NewAttributeDescription("Aging based on inactivity time period").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"switchport_port_security_maximum_range": {
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Optional:            true,
				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
					"range": {
						MarkdownDescription: helpers.NewAttributeDescription("Maximum addresses").AddIntegerRangeDescription(1, 3072).String,
						Type:                types.Int64Type,
						Optional:            true,
						Computed:            true,
						Validators: []tfsdk.AttributeValidator{
							helpers.IntegerRangeValidator(1, 3072),
						},
					},
					"vlan": {
						MarkdownDescription: helpers.NewAttributeDescription("Max secure addresses per vlan").String,
						Type:                types.BoolType,
						Optional:            true,
						Computed:            true,
					},
					"vlan_access": {
						MarkdownDescription: helpers.NewAttributeDescription("access vlan").String,
						Type:                types.BoolType,
						Optional:            true,
						Computed:            true,
					},
				}),
			},
			"switchport_port_security_violation_protect": {
				MarkdownDescription: helpers.NewAttributeDescription("Security violation protect mode").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"switchport_port_security_violation_restrict": {
				MarkdownDescription: helpers.NewAttributeDescription("Security violation restrict mode").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"switchport_port_security_violation_shutdown": {
				MarkdownDescription: helpers.NewAttributeDescription("Security violation shutdown mode").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"switchport_access_vlan": {
				MarkdownDescription: helpers.NewAttributeDescription("VLAN ID of the VLAN when this port is in access mode").AddIntegerRangeDescription(1, 4094).String,
				Type:                types.Int64Type,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.IntegerRangeValidator(1, 4094),
				},
			},
			"switchport_voice_vlan": {
				MarkdownDescription: helpers.NewAttributeDescription("Vlan for voice traffic").AddIntegerRangeDescription(1, 4094).String,
				Type:                types.Int64Type,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.IntegerRangeValidator(1, 4094),
				},
			},
			"switchport_private_vlan_host_association_primary_range": {
				MarkdownDescription: helpers.NewAttributeDescription("Primary normal range VLAN ID of the private VLAN port association").AddIntegerRangeDescription(2, 1001).String,
				Type:                types.Int64Type,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.IntegerRangeValidator(2, 1001),
				},
			},
			"switchport_private_vlan_host_association_secondary_range": {
				MarkdownDescription: helpers.NewAttributeDescription("Secondary normal range VLAN ID of the private VLAN host port association").AddIntegerRangeDescription(2, 1001).String,
				Type:                types.Int64Type,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.IntegerRangeValidator(2, 1001),
				},
			},
			"switchport_trunk_allowed_vlans": {
				MarkdownDescription: helpers.NewAttributeDescription("VLAN IDs of the allowed VLANs when this port is in trunking mode").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
			},
			"switchport_trunk_native_vlan_tag": {
				MarkdownDescription: helpers.NewAttributeDescription("Set native VLAN tagging state").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"switchport_trunk_native_vlan_vlan_id": {
				MarkdownDescription: helpers.NewAttributeDescription("VLAN ID of the native VLAN when this port is in trunking mode").AddIntegerRangeDescription(1, 4094).String,
				Type:                types.Int64Type,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.IntegerRangeValidator(1, 4094),
				},
			},
			"mab": {
				MarkdownDescription: helpers.NewAttributeDescription("MAC Authentication Bypass Interface Config Commands").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"mab_eap": {
				MarkdownDescription: helpers.NewAttributeDescription("Use EAP authentication for MAC Auth Bypass").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"access_session_closed": {
				MarkdownDescription: helpers.NewAttributeDescription("Enable closed access on port (disabled by default, i.e. open access)").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"access_session_monitor": {
				MarkdownDescription: helpers.NewAttributeDescription("Apply template to monitor access sessions on the port").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"access_session_port_control": {
				MarkdownDescription: helpers.NewAttributeDescription("Set the port-control value").AddStringEnumDescription("auto", "force-authorized", "force-unauthorized").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.StringEnumValidator("auto", "force-authorized", "force-unauthorized"),
				},
			},
			"access_session_control_direction": {
				MarkdownDescription: helpers.NewAttributeDescription("Set the control-direction on the interface").AddStringEnumDescription("both", "in").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.StringEnumValidator("both", "in"),
				},
			},
			"access_session_host_mode": {
				MarkdownDescription: helpers.NewAttributeDescription("Set the Host mode for authentication on this interface").AddStringEnumDescription("multi-auth", "multi-domain", "multi-host", "single-host").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.StringEnumValidator("multi-auth", "multi-domain", "multi-host", "single-host"),
				},
			},
			"access_session_interface_template_sticky": {
				MarkdownDescription: helpers.NewAttributeDescription("Interface templates set to sticky").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"access_session_interface_template_sticky_timer": {
				MarkdownDescription: helpers.NewAttributeDescription("Sticky timer value").AddIntegerRangeDescription(1, 65535).String,
				Type:                types.Int64Type,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.IntegerRangeValidator(1, 65535),
				},
			},
			"authentication_periodic": {
				MarkdownDescription: helpers.NewAttributeDescription("Enable or Disable Reauthentication for this port").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"authentication_timer_reauthenticate_server": {
				MarkdownDescription: helpers.NewAttributeDescription("Obtain re-authentication timeout value from the server").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"authentication_timer_reauthenticate_range": {
				MarkdownDescription: helpers.NewAttributeDescription("Enter a value between 1 and 65535").AddIntegerRangeDescription(1, 65535).String,
				Type:                types.Int64Type,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.IntegerRangeValidator(1, 65535),
				},
			},
			"spanning_tree_bpduguard_enable": {
				MarkdownDescription: helpers.NewAttributeDescription("Enable BPDU guard for this interface").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"spanning_tree_service_policy": {
				MarkdownDescription: helpers.NewAttributeDescription("help").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"spanning_tree_portfast": {
				MarkdownDescription: helpers.NewAttributeDescription("Portfast options for the interface").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"spanning_tree_portfast_disable": {
				MarkdownDescription: helpers.NewAttributeDescription("Disable portfast for this interface").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"spanning_tree_portfast_edge": {
				MarkdownDescription: helpers.NewAttributeDescription("Enable portfast edge on the interface").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"spanning_tree_portfast_network": {
				MarkdownDescription: helpers.NewAttributeDescription("Enable portfast network on the interface").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"storm_control_broadcast_level_pps_threshold": {
				MarkdownDescription: helpers.NewAttributeDescription("Enter threshold").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
			},
			"storm_control_broadcast_level_bps_threshold": {
				MarkdownDescription: helpers.NewAttributeDescription("Enter threshold").String,
				Type:                types.Float64Type,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.FloatRangeValidator(0, 1e+11),
				},
			},
			"storm_control_broadcast_level_threshold": {
				MarkdownDescription: helpers.NewAttributeDescription("Enter threshold").String,
				Type:                types.Float64Type,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.FloatRangeValidator(0, 10000),
				},
			},
			"storm_control_multicast_level_pps_threshold": {
				MarkdownDescription: helpers.NewAttributeDescription("Enter threshold").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
			},
			"storm_control_multicast_level_bps_threshold": {
				MarkdownDescription: helpers.NewAttributeDescription("Enter threshold").String,
				Type:                types.Float64Type,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.FloatRangeValidator(0, 1e+11),
				},
			},
			"storm_control_multicast_level_threshold": {
				MarkdownDescription: helpers.NewAttributeDescription("Enter threshold").String,
				Type:                types.Float64Type,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.FloatRangeValidator(0, 10000),
				},
			},
			"storm_control_action_shutdown": {
				MarkdownDescription: helpers.NewAttributeDescription("Shutdown this interface if a storm occurs").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"storm_control_action_trap": {
				MarkdownDescription: helpers.NewAttributeDescription("Send SNMP trap if a storm occurs").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"load_interval": {
				MarkdownDescription: helpers.NewAttributeDescription("Load interval delay in seconds").AddIntegerRangeDescription(30, 600).String,
				Type:                types.Int64Type,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.IntegerRangeValidator(30, 600),
				},
			},
			"ip_dhcp_snooping_limit_rate": {
				MarkdownDescription: helpers.NewAttributeDescription("DHCP snooping rate limit").AddIntegerRangeDescription(1, 2048).String,
				Type:                types.Int64Type,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.IntegerRangeValidator(1, 2048),
				},
			},
			"ip_dhcp_snooping_trust": {
				MarkdownDescription: helpers.NewAttributeDescription("DHCP Snooping trust config").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"ip_access_group": {
				MarkdownDescription: helpers.NewAttributeDescription("Access control list for IP packets").String,
				Optional:            true,
				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
					"direction": {
						MarkdownDescription: helpers.NewAttributeDescription("packet flow direction").AddStringEnumDescription("in", "out").String,
						Type:                types.StringType,
						Optional:            true,
						Computed:            true,
						Validators: []tfsdk.AttributeValidator{
							helpers.StringEnumValidator("in", "out"),
						},
					},
					"access_list": {
						MarkdownDescription: helpers.NewAttributeDescription("Access control list name").String,
						Type:                types.StringType,
						Optional:            true,
						Computed:            true,
					},
				}),
			},
			"subscriber_aging_inactivity_timer_value": {
				MarkdownDescription: helpers.NewAttributeDescription("Enter a value between 1 and 65535 in seconds").AddIntegerRangeDescription(1, 65535).String,
				Type:                types.Int64Type,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.IntegerRangeValidator(1, 65535),
				},
			},
			"subscriber_aging_inactivity_timer_probe": {
				MarkdownDescription: helpers.NewAttributeDescription("ARP probe").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"subscriber_aging_probe": {
				MarkdownDescription: helpers.NewAttributeDescription("ARP probe").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"device_tracking": {
				MarkdownDescription: helpers.NewAttributeDescription("Configure device-tracking on the interface").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"device_tracking_attach_policy": {
				MarkdownDescription: helpers.NewAttributeDescription("policy name for device tracking").String,
				Optional:            true,
				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
					"policy_name": {
						MarkdownDescription: helpers.NewAttributeDescription("").String,
						Type:                types.StringType,
						Optional:            true,
						Computed:            true,
					},
					"vlan_range": {
						MarkdownDescription: helpers.NewAttributeDescription("VLAN IDs of the VLANs for which this policy applies").String,
						Type:                types.StringType,
						Optional:            true,
						Computed:            true,
					},
				}),
			},
			"device_tracking_vlan_range": {
				MarkdownDescription: helpers.NewAttributeDescription("VLAN IDs of the VLANs for which this policy applies").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
			},
			"cts_manual": {
				MarkdownDescription: helpers.NewAttributeDescription("Supply local configuration for CTS parameters").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"cts_manual_policy_static_sgt": {
				MarkdownDescription: helpers.NewAttributeDescription("Source Security Group Tag to apply to untagged or non-trusted incoming traffic").AddIntegerRangeDescription(2, 65519).String,
				Type:                types.Int64Type,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.IntegerRangeValidator(2, 65519),
				},
			},
			"cts_manual_policy_static_trusted": {
				MarkdownDescription: helpers.NewAttributeDescription("Trust the Source Group Tags (SGT) that the peer uses for sending").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"cts_manual_propagate_sgt": {
				MarkdownDescription: helpers.NewAttributeDescription("CTS SGT Propagation configuration").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"cts_role_based_enforcement": {
				MarkdownDescription: helpers.NewAttributeDescription("Enable Role-based Access Control enforcement").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
		},
	}, nil
}

func (t resourceTemplateType) NewResource(ctx context.Context, in provider.Provider) (resource.Resource, diag.Diagnostics) {
	provider, diags := convertProviderType(in)

	return resourceTemplate{
		provider: provider,
	}, diags
}

type resourceTemplate struct {
	provider iosxeProvider
}

func (r resourceTemplate) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan Template

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

func (r resourceTemplate) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state Template

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.Value))

	res, err := r.provider.clients[state.Device.Value].GetData(state.Id.Value)
	if res.StatusCode == 404 {
		state = Template{Device: state.Device, Id: state.Id}
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

func (r resourceTemplate) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state Template

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

func (r resourceTemplate) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state Template

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

func (r resourceTemplate) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
