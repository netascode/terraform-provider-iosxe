// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-restconf"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &TemplateDataSource{}
	_ datasource.DataSourceWithConfigure = &TemplateDataSource{}
)

func NewTemplateDataSource() datasource.DataSource {
	return &TemplateDataSource{}
}

type TemplateDataSource struct {
	clients map[string]*restconf.Client
}

func (d *TemplateDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_template"
}

func (d *TemplateDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Template configuration.",

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The path of the retrieved object.",
				Computed:            true,
			},
			"template_name": schema.StringAttribute{
				MarkdownDescription: "Template name",
				Required:            true,
			},
			"dot1x_pae": schema.StringAttribute{
				MarkdownDescription: "Set 802.1x interface pae type",
				Computed:            true,
			},
			"dot1x_max_reauth_req": schema.Int64Attribute{
				MarkdownDescription: "Max No. of Reauthentication Attempts",
				Computed:            true,
			},
			"dot1x_max_req": schema.Int64Attribute{
				MarkdownDescription: "Max No. of Retries",
				Computed:            true,
			},
			"service_policy_input": schema.StringAttribute{
				MarkdownDescription: "policy-map name",
				Computed:            true,
			},
			"service_policy_output": schema.StringAttribute{
				MarkdownDescription: "policy-map name",
				Computed:            true,
			},
			"source_template": schema.StringAttribute{
				MarkdownDescription: "Get config from a template",
				Computed:            true,
			},
			"switchport_mode_trunk": schema.BoolAttribute{
				MarkdownDescription: "Set trunking mode to TRUNK unconditionally",
				Computed:            true,
			},
			"switchport_mode_access": schema.BoolAttribute{
				MarkdownDescription: "Set trunking mode to ACCESS unconditionally",
				Computed:            true,
			},
			"switchport_nonegotiate": schema.BoolAttribute{
				MarkdownDescription: "Device will not engage in negotiation protocol on this interface",
				Computed:            true,
			},
			"switchport_block_unicast": schema.BoolAttribute{
				MarkdownDescription: "Block unknown unicast addresses",
				Computed:            true,
			},
			"switchport_port_security": schema.BoolAttribute{
				MarkdownDescription: "Security related command",
				Computed:            true,
			},
			"switchport_port_security_aging_static": schema.BoolAttribute{
				MarkdownDescription: "Enable aging for configured secure addresses",
				Computed:            true,
			},
			"switchport_port_security_aging_time": schema.Int64Attribute{
				MarkdownDescription: "Port-security aging time",
				Computed:            true,
			},
			"switchport_port_security_aging_type": schema.BoolAttribute{
				MarkdownDescription: "Port-security aging type",
				Computed:            true,
			},
			"switchport_port_security_aging_type_inactivity": schema.BoolAttribute{
				MarkdownDescription: "Aging based on inactivity time period",
				Computed:            true,
			},
			"switchport_port_security_maximum_range": schema.ListNestedAttribute{
				MarkdownDescription: "",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"range": schema.Int64Attribute{
							MarkdownDescription: "Maximum addresses",
							Computed:            true,
						},
						"vlan": schema.BoolAttribute{
							MarkdownDescription: "Max secure addresses per vlan",
							Computed:            true,
						},
						"vlan_access": schema.BoolAttribute{
							MarkdownDescription: "access vlan",
							Computed:            true,
						},
					},
				},
			},
			"switchport_port_security_violation_protect": schema.BoolAttribute{
				MarkdownDescription: "Security violation protect mode",
				Computed:            true,
			},
			"switchport_port_security_violation_restrict": schema.BoolAttribute{
				MarkdownDescription: "Security violation restrict mode",
				Computed:            true,
			},
			"switchport_port_security_violation_shutdown": schema.BoolAttribute{
				MarkdownDescription: "Security violation shutdown mode",
				Computed:            true,
			},
			"switchport_access_vlan": schema.Int64Attribute{
				MarkdownDescription: "VLAN ID of the VLAN when this port is in access mode",
				Computed:            true,
			},
			"switchport_voice_vlan": schema.Int64Attribute{
				MarkdownDescription: "Vlan for voice traffic",
				Computed:            true,
			},
			"switchport_private_vlan_host_association_primary_range": schema.Int64Attribute{
				MarkdownDescription: "Primary normal range VLAN ID of the private VLAN port association",
				Computed:            true,
			},
			"switchport_private_vlan_host_association_secondary_range": schema.Int64Attribute{
				MarkdownDescription: "Secondary normal range VLAN ID of the private VLAN host port association",
				Computed:            true,
			},
			"switchport_trunk_allowed_vlans": schema.StringAttribute{
				MarkdownDescription: "VLAN IDs of the allowed VLANs when this port is in trunking mode",
				Computed:            true,
			},
			"switchport_trunk_native_vlan_tag": schema.BoolAttribute{
				MarkdownDescription: "Set native VLAN tagging state",
				Computed:            true,
			},
			"switchport_trunk_native_vlan_vlan_id": schema.Int64Attribute{
				MarkdownDescription: "VLAN ID of the native VLAN when this port is in trunking mode",
				Computed:            true,
			},
			"mab": schema.BoolAttribute{
				MarkdownDescription: "MAC Authentication Bypass Interface Config Commands",
				Computed:            true,
			},
			"mab_eap": schema.BoolAttribute{
				MarkdownDescription: "Use EAP authentication for MAC Auth Bypass",
				Computed:            true,
			},
			"access_session_closed": schema.BoolAttribute{
				MarkdownDescription: "Enable closed access on port (disabled by default, i.e. open access)",
				Computed:            true,
			},
			"access_session_monitor": schema.BoolAttribute{
				MarkdownDescription: "Apply template to monitor access sessions on the port",
				Computed:            true,
			},
			"access_session_port_control": schema.StringAttribute{
				MarkdownDescription: "Set the port-control value",
				Computed:            true,
			},
			"access_session_control_direction": schema.StringAttribute{
				MarkdownDescription: "Set the control-direction on the interface",
				Computed:            true,
			},
			"access_session_host_mode": schema.StringAttribute{
				MarkdownDescription: "Set the Host mode for authentication on this interface",
				Computed:            true,
			},
			"access_session_interface_template_sticky": schema.BoolAttribute{
				MarkdownDescription: "Interface templates set to sticky",
				Computed:            true,
			},
			"access_session_interface_template_sticky_timer": schema.Int64Attribute{
				MarkdownDescription: "Sticky timer value",
				Computed:            true,
			},
			"authentication_periodic": schema.BoolAttribute{
				MarkdownDescription: "Enable or Disable Reauthentication for this port",
				Computed:            true,
			},
			"authentication_timer_reauthenticate_server": schema.BoolAttribute{
				MarkdownDescription: "Obtain re-authentication timeout value from the server",
				Computed:            true,
			},
			"authentication_timer_reauthenticate_range": schema.Int64Attribute{
				MarkdownDescription: "Enter a value between 1 and 65535",
				Computed:            true,
			},
			"spanning_tree_bpduguard_enable": schema.BoolAttribute{
				MarkdownDescription: "Enable BPDU guard for this interface",
				Computed:            true,
			},
			"spanning_tree_service_policy": schema.BoolAttribute{
				MarkdownDescription: "help",
				Computed:            true,
			},
			"spanning_tree_portfast": schema.BoolAttribute{
				MarkdownDescription: "Portfast options for the interface",
				Computed:            true,
			},
			"spanning_tree_portfast_disable": schema.BoolAttribute{
				MarkdownDescription: "Disable portfast for this interface",
				Computed:            true,
			},
			"spanning_tree_portfast_edge": schema.BoolAttribute{
				MarkdownDescription: "Enable portfast edge on the interface",
				Computed:            true,
			},
			"spanning_tree_portfast_network": schema.BoolAttribute{
				MarkdownDescription: "Enable portfast network on the interface",
				Computed:            true,
			},
			"storm_control_broadcast_level_pps_threshold": schema.StringAttribute{
				MarkdownDescription: "Enter threshold",
				Computed:            true,
			},
			"storm_control_broadcast_level_bps_threshold": schema.Float64Attribute{
				MarkdownDescription: "Enter threshold",
				Computed:            true,
			},
			"storm_control_broadcast_level_threshold": schema.Float64Attribute{
				MarkdownDescription: "Enter threshold",
				Computed:            true,
			},
			"storm_control_multicast_level_pps_threshold": schema.StringAttribute{
				MarkdownDescription: "Enter threshold",
				Computed:            true,
			},
			"storm_control_multicast_level_bps_threshold": schema.Float64Attribute{
				MarkdownDescription: "Enter threshold",
				Computed:            true,
			},
			"storm_control_multicast_level_threshold": schema.Float64Attribute{
				MarkdownDescription: "Enter threshold",
				Computed:            true,
			},
			"storm_control_action_shutdown": schema.BoolAttribute{
				MarkdownDescription: "Shutdown this interface if a storm occurs",
				Computed:            true,
			},
			"storm_control_action_trap": schema.BoolAttribute{
				MarkdownDescription: "Send SNMP trap if a storm occurs",
				Computed:            true,
			},
			"load_interval": schema.Int64Attribute{
				MarkdownDescription: "Load interval delay in seconds",
				Computed:            true,
			},
			"ip_dhcp_snooping_limit_rate": schema.Int64Attribute{
				MarkdownDescription: "DHCP snooping rate limit",
				Computed:            true,
			},
			"ip_dhcp_snooping_trust": schema.BoolAttribute{
				MarkdownDescription: "DHCP Snooping trust config",
				Computed:            true,
			},
			"ip_access_group": schema.ListNestedAttribute{
				MarkdownDescription: "Access control list for IP packets",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"direction": schema.StringAttribute{
							MarkdownDescription: "packet flow direction",
							Computed:            true,
						},
						"access_list": schema.StringAttribute{
							MarkdownDescription: "Access control list name",
							Computed:            true,
						},
					},
				},
			},
			"subscriber_aging_inactivity_timer_value": schema.Int64Attribute{
				MarkdownDescription: "Enter a value between 1 and 65535 in seconds",
				Computed:            true,
			},
			"subscriber_aging_inactivity_timer_probe": schema.BoolAttribute{
				MarkdownDescription: "ARP probe",
				Computed:            true,
			},
			"subscriber_aging_probe": schema.BoolAttribute{
				MarkdownDescription: "ARP probe",
				Computed:            true,
			},
			"device_tracking": schema.BoolAttribute{
				MarkdownDescription: "Configure device-tracking on the interface",
				Computed:            true,
			},
			"device_tracking_attach_policy": schema.ListNestedAttribute{
				MarkdownDescription: "policy name for device tracking",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"policy_name": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"vlan_range": schema.StringAttribute{
							MarkdownDescription: "VLAN IDs of the VLANs for which this policy applies",
							Computed:            true,
						},
					},
				},
			},
			"device_tracking_vlan_range": schema.StringAttribute{
				MarkdownDescription: "VLAN IDs of the VLANs for which this policy applies",
				Computed:            true,
			},
			"cts_manual": schema.BoolAttribute{
				MarkdownDescription: "Supply local configuration for CTS parameters",
				Computed:            true,
			},
			"cts_manual_policy_static_sgt": schema.Int64Attribute{
				MarkdownDescription: "Source Security Group Tag to apply to untagged or non-trusted incoming traffic",
				Computed:            true,
			},
			"cts_manual_policy_static_trusted": schema.BoolAttribute{
				MarkdownDescription: "Trust the Source Group Tags (SGT) that the peer uses for sending",
				Computed:            true,
			},
			"cts_manual_propagate_sgt": schema.BoolAttribute{
				MarkdownDescription: "CTS SGT Propagation configuration",
				Computed:            true,
			},
			"cts_role_based_enforcement": schema.BoolAttribute{
				MarkdownDescription: "Enable Role-based Access Control enforcement",
				Computed:            true,
			},
		},
	}
}

func (d *TemplateDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.clients = req.ProviderData.(map[string]*restconf.Client)
}

func (d *TemplateDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config Template

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if _, ok := d.clients[config.Device.ValueString()]; !ok {
		resp.Diagnostics.AddAttributeError(path.Root("device"), "Invalid device", fmt.Sprintf("Device '%s' does not exist in provider configuration.", config.Device.ValueString()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.getPath()))

	res, err := d.clients[config.Device.ValueString()].GetData(config.getPath())
	if res.StatusCode == 404 {
		config = Template{Device: config.Device}
	} else {
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
			return
		}

		config.fromBody(ctx, res.Res)
	}

	config.Id = types.StringValue(config.getPath())

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.getPath()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}
