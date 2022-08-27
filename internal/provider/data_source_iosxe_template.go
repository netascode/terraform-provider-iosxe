// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

type dataSourceTemplateType struct{}

func (t dataSourceTemplateType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Template configuration.",

		Attributes: map[string]tfsdk.Attribute{
			"device": {
				MarkdownDescription: "A device name from the provider configuration.",
				Type:                types.StringType,
				Optional:            true,
			},
			"id": {
				MarkdownDescription: "The path of the retrieved object.",
				Type:                types.StringType,
				Computed:            true,
			},
			"template_name": {
				MarkdownDescription: "Template name",
				Type:                types.StringType,
				Required:            true,
			},
			"dot1x_pae": {
				MarkdownDescription: "Set 802.1x interface pae type",
				Type:                types.StringType,
				Computed:            true,
			},
			"dot1x_max_reauth_req": {
				MarkdownDescription: "Max No. of Reauthentication Attempts",
				Type:                types.Int64Type,
				Computed:            true,
			},
			"dot1x_max_req": {
				MarkdownDescription: "Max No. of Retries",
				Type:                types.Int64Type,
				Computed:            true,
			},
			"service_policy_input": {
				MarkdownDescription: "policy-map name",
				Type:                types.StringType,
				Computed:            true,
			},
			"service_policy_output": {
				MarkdownDescription: "policy-map name",
				Type:                types.StringType,
				Computed:            true,
			},
			"source_template": {
				MarkdownDescription: "Get config from a template",
				Type:                types.StringType,
				Computed:            true,
			},
			"switchport_mode_trunk": {
				MarkdownDescription: "Set trunking mode to TRUNK unconditionally",
				Type:                types.BoolType,
				Computed:            true,
			},
			"switchport_mode_access": {
				MarkdownDescription: "Set trunking mode to ACCESS unconditionally",
				Type:                types.BoolType,
				Computed:            true,
			},
			"switchport_nonegotiate": {
				MarkdownDescription: "Device will not engage in negotiation protocol on this interface",
				Type:                types.BoolType,
				Computed:            true,
			},
			"switchport_block_unicast": {
				MarkdownDescription: "Block unknown unicast addresses",
				Type:                types.BoolType,
				Computed:            true,
			},
			"switchport_port_security": {
				MarkdownDescription: "Security related command",
				Type:                types.BoolType,
				Computed:            true,
			},
			"switchport_port_security_aging_static": {
				MarkdownDescription: "Enable aging for configured secure addresses",
				Type:                types.BoolType,
				Computed:            true,
			},
			"switchport_port_security_aging_time": {
				MarkdownDescription: "Port-security aging time",
				Type:                types.Int64Type,
				Computed:            true,
			},
			"switchport_port_security_aging_type": {
				MarkdownDescription: "Port-security aging type",
				Type:                types.BoolType,
				Computed:            true,
			},
			"switchport_port_security_aging_type_inactivity": {
				MarkdownDescription: "Aging based on inactivity time period",
				Type:                types.BoolType,
				Computed:            true,
			},
			"switchport_port_security_maximum_range": {
				MarkdownDescription: "",
				Computed:            true,
				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
					"range": {
						MarkdownDescription: "Maximum addresses",
						Type:                types.Int64Type,
						Computed:            true,
					},
					"vlan": {
						MarkdownDescription: "Max secure addresses per vlan",
						Type:                types.BoolType,
						Computed:            true,
					},
					"vlan_access": {
						MarkdownDescription: "access vlan",
						Type:                types.BoolType,
						Computed:            true,
					},
				}),
			},
			"switchport_port_security_violation_protect": {
				MarkdownDescription: "Security violation protect mode",
				Type:                types.BoolType,
				Computed:            true,
			},
			"switchport_port_security_violation_restrict": {
				MarkdownDescription: "Security violation restrict mode",
				Type:                types.BoolType,
				Computed:            true,
			},
			"switchport_port_security_violation_shutdown": {
				MarkdownDescription: "Security violation shutdown mode",
				Type:                types.BoolType,
				Computed:            true,
			},
			"switchport_access_vlan": {
				MarkdownDescription: "VLAN ID of the VLAN when this port is in access mode",
				Type:                types.Int64Type,
				Computed:            true,
			},
			"switchport_voice_vlan": {
				MarkdownDescription: "Vlan for voice traffic",
				Type:                types.Int64Type,
				Computed:            true,
			},
			"switchport_private_vlan_host_association_primary_range": {
				MarkdownDescription: "Primary normal range VLAN ID of the private VLAN port association",
				Type:                types.Int64Type,
				Computed:            true,
			},
			"switchport_private_vlan_host_association_secondary_range": {
				MarkdownDescription: "Secondary normal range VLAN ID of the private VLAN host port association",
				Type:                types.Int64Type,
				Computed:            true,
			},
			"switchport_trunk_allowed_vlans": {
				MarkdownDescription: "VLAN IDs of the allowed VLANs when this port is in trunking mode",
				Type:                types.StringType,
				Computed:            true,
			},
			"switchport_trunk_native_vlan_tag": {
				MarkdownDescription: "Set native VLAN tagging state",
				Type:                types.BoolType,
				Computed:            true,
			},
			"switchport_trunk_native_vlan_vlan_id": {
				MarkdownDescription: "VLAN ID of the native VLAN when this port is in trunking mode",
				Type:                types.Int64Type,
				Computed:            true,
			},
			"mab": {
				MarkdownDescription: "MAC Authentication Bypass Interface Config Commands",
				Type:                types.BoolType,
				Computed:            true,
			},
			"mab_eap": {
				MarkdownDescription: "Use EAP authentication for MAC Auth Bypass",
				Type:                types.BoolType,
				Computed:            true,
			},
			"access_session_closed": {
				MarkdownDescription: "Enable closed access on port (disabled by default, i.e. open access)",
				Type:                types.BoolType,
				Computed:            true,
			},
			"access_session_monitor": {
				MarkdownDescription: "Apply template to monitor access sessions on the port",
				Type:                types.BoolType,
				Computed:            true,
			},
			"access_session_port_control": {
				MarkdownDescription: "Set the port-control value",
				Type:                types.StringType,
				Computed:            true,
			},
			"access_session_control_direction": {
				MarkdownDescription: "Set the control-direction on the interface",
				Type:                types.StringType,
				Computed:            true,
			},
			"access_session_host_mode": {
				MarkdownDescription: "Set the Host mode for authentication on this interface",
				Type:                types.StringType,
				Computed:            true,
			},
			"access_session_interface_template_sticky": {
				MarkdownDescription: "Interface templates set to sticky",
				Type:                types.BoolType,
				Computed:            true,
			},
			"access_session_interface_template_sticky_timer": {
				MarkdownDescription: "Sticky timer value",
				Type:                types.Int64Type,
				Computed:            true,
			},
			"authentication_periodic": {
				MarkdownDescription: "Enable or Disable Reauthentication for this port",
				Type:                types.BoolType,
				Computed:            true,
			},
			"authentication_timer_reauthenticate_server": {
				MarkdownDescription: "Obtain re-authentication timeout value from the server",
				Type:                types.BoolType,
				Computed:            true,
			},
			"authentication_timer_reauthenticate_range": {
				MarkdownDescription: "Enter a value between 1 and 65535",
				Type:                types.Int64Type,
				Computed:            true,
			},
			"spanning_tree_bpduguard_enable": {
				MarkdownDescription: "Enable BPDU guard for this interface",
				Type:                types.BoolType,
				Computed:            true,
			},
			"spanning_tree_service_policy": {
				MarkdownDescription: "help",
				Type:                types.BoolType,
				Computed:            true,
			},
			"spanning_tree_portfast": {
				MarkdownDescription: "Portfast options for the interface",
				Type:                types.BoolType,
				Computed:            true,
			},
			"spanning_tree_portfast_disable": {
				MarkdownDescription: "Disable portfast for this interface",
				Type:                types.BoolType,
				Computed:            true,
			},
			"spanning_tree_portfast_edge": {
				MarkdownDescription: "Enable portfast edge on the interface",
				Type:                types.BoolType,
				Computed:            true,
			},
			"spanning_tree_portfast_network": {
				MarkdownDescription: "Enable portfast network on the interface",
				Type:                types.BoolType,
				Computed:            true,
			},
			"storm_control_broadcast_level_pps_threshold": {
				MarkdownDescription: "Enter threshold",
				Type:                types.StringType,
				Computed:            true,
			},
			"storm_control_broadcast_level_bps_threshold": {
				MarkdownDescription: "Enter threshold",
				Type:                types.Float64Type,
				Computed:            true,
			},
			"storm_control_broadcast_level_threshold": {
				MarkdownDescription: "Enter threshold",
				Type:                types.Float64Type,
				Computed:            true,
			},
			"storm_control_multicast_level_pps_threshold": {
				MarkdownDescription: "Enter threshold",
				Type:                types.StringType,
				Computed:            true,
			},
			"storm_control_multicast_level_bps_threshold": {
				MarkdownDescription: "Enter threshold",
				Type:                types.Float64Type,
				Computed:            true,
			},
			"storm_control_multicast_level_threshold": {
				MarkdownDescription: "Enter threshold",
				Type:                types.Float64Type,
				Computed:            true,
			},
			"storm_control_action_shutdown": {
				MarkdownDescription: "Shutdown this interface if a storm occurs",
				Type:                types.BoolType,
				Computed:            true,
			},
			"storm_control_action_trap": {
				MarkdownDescription: "Send SNMP trap if a storm occurs",
				Type:                types.BoolType,
				Computed:            true,
			},
			"load_interval": {
				MarkdownDescription: "Load interval delay in seconds",
				Type:                types.Int64Type,
				Computed:            true,
			},
			"ip_dhcp_snooping_limit_rate": {
				MarkdownDescription: "DHCP snooping rate limit",
				Type:                types.Int64Type,
				Computed:            true,
			},
			"ip_dhcp_snooping_trust": {
				MarkdownDescription: "DHCP Snooping trust config",
				Type:                types.BoolType,
				Computed:            true,
			},
			"ip_access_group": {
				MarkdownDescription: "Access control list for IP packets",
				Computed:            true,
				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
					"direction": {
						MarkdownDescription: "packet flow direction",
						Type:                types.StringType,
						Computed:            true,
					},
					"access_list": {
						MarkdownDescription: "Access control list name",
						Type:                types.StringType,
						Computed:            true,
					},
				}),
			},
			"subscriber_aging_inactivity_timer_value": {
				MarkdownDescription: "Enter a value between 1 and 65535 in seconds",
				Type:                types.Int64Type,
				Computed:            true,
			},
			"subscriber_aging_inactivity_timer_probe": {
				MarkdownDescription: "ARP probe",
				Type:                types.BoolType,
				Computed:            true,
			},
			"subscriber_aging_probe": {
				MarkdownDescription: "ARP probe",
				Type:                types.BoolType,
				Computed:            true,
			},
			"device_tracking": {
				MarkdownDescription: "Configure device-tracking on the interface",
				Type:                types.BoolType,
				Computed:            true,
			},
			"device_tracking_attach_policy": {
				MarkdownDescription: "policy name for device tracking",
				Computed:            true,
				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
					"policy_name": {
						MarkdownDescription: "",
						Type:                types.StringType,
						Computed:            true,
					},
					"vlan_range": {
						MarkdownDescription: "VLAN IDs of the VLANs for which this policy applies",
						Type:                types.StringType,
						Computed:            true,
					},
				}),
			},
			"device_tracking_vlan_range": {
				MarkdownDescription: "VLAN IDs of the VLANs for which this policy applies",
				Type:                types.StringType,
				Computed:            true,
			},
			"cts_manual": {
				MarkdownDescription: "Supply local configuration for CTS parameters",
				Type:                types.BoolType,
				Computed:            true,
			},
			"cts_manual_policy_static_sgt": {
				MarkdownDescription: "Source Security Group Tag to apply to untagged or non-trusted incoming traffic",
				Type:                types.Int64Type,
				Computed:            true,
			},
			"cts_manual_policy_static_trusted": {
				MarkdownDescription: "Trust the Source Group Tags (SGT) that the peer uses for sending",
				Type:                types.BoolType,
				Computed:            true,
			},
			"cts_manual_propagate_sgt": {
				MarkdownDescription: "CTS SGT Propagation configuration",
				Type:                types.BoolType,
				Computed:            true,
			},
			"cts_role_based_enforcement": {
				MarkdownDescription: "Enable Role-based Access Control enforcement",
				Type:                types.BoolType,
				Computed:            true,
			},
		},
	}, nil
}

func (t dataSourceTemplateType) NewDataSource(ctx context.Context, in provider.Provider) (datasource.DataSource, diag.Diagnostics) {
	provider, diags := convertProviderType(in)

	return dataSourceTemplate{
		provider: provider,
	}, diags
}

type dataSourceTemplate struct {
	provider iosxeProvider
}

func (d dataSourceTemplate) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config Template

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.getPath()))

	res, err := d.provider.clients[config.Device.Value].GetData(config.getPath())
	if res.StatusCode == 404 {
		config = Template{Device: config.Device}
	} else {
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
			return
		}

		config.fromBody(ctx, res.Res)
	}

	config.Id = types.String{Value: config.getPath()}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.getPath()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}
