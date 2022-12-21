// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-restconf"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &InterfaceSwitchportDataSource{}
	_ datasource.DataSourceWithConfigure = &InterfaceSwitchportDataSource{}
)

func NewInterfaceSwitchportDataSource() datasource.DataSource {
	return &InterfaceSwitchportDataSource{}
}

type InterfaceSwitchportDataSource struct {
	clients map[string]*restconf.Client
}

func (d *InterfaceSwitchportDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interface_switchport"
}

func (d *InterfaceSwitchportDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Interface Switchport configuration.",

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The path of the retrieved object.",
				Computed:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: "Interface type",
				Required:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "",
				Required:            true,
			},
			"mode_access": schema.BoolAttribute{
				MarkdownDescription: "Set trunking mode to ACCESS unconditionally",
				Computed:            true,
			},
			"mode_dot1q_tunnel": schema.BoolAttribute{
				MarkdownDescription: "set trunking mode to TUNNEL unconditionally",
				Computed:            true,
			},
			"mode_private_vlan_trunk": schema.BoolAttribute{
				MarkdownDescription: "Set the mode to private-vlan trunk",
				Computed:            true,
			},
			"mode_private_vlan_host": schema.BoolAttribute{
				MarkdownDescription: "Set the mode to private-vlan host",
				Computed:            true,
			},
			"mode_private_vlan_promiscuous": schema.BoolAttribute{
				MarkdownDescription: "Set the mode to private-vlan promiscuous",
				Computed:            true,
			},
			"mode_trunk": schema.BoolAttribute{
				MarkdownDescription: "Set trunking mode to TRUNK unconditionally",
				Computed:            true,
			},
			"nonegotiate": schema.BoolAttribute{
				MarkdownDescription: "Device will not engage in negotiation protocol on this interface",
				Computed:            true,
			},
			"access_vlan": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"trunk_allowed_vlans": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"trunk_native_vlan_tag": schema.BoolAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"trunk_native_vlan": schema.Int64Attribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"host": schema.BoolAttribute{
				MarkdownDescription: "Set port host",
				Computed:            true,
			},
		},
	}
}

func (d *InterfaceSwitchportDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.clients = req.ProviderData.(map[string]*restconf.Client)
}

func (d *InterfaceSwitchportDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config InterfaceSwitchport

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.getPath()))

	res, err := d.clients[config.Device.ValueString()].GetData(config.getPath())
	if res.StatusCode == 404 {
		config = InterfaceSwitchport{Device: config.Device}
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
