// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

type dataSourceInterfaceVLANType struct{}

func (t dataSourceInterfaceVLANType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Interface VLAN configuration.",

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
			"name": {
				MarkdownDescription: "",
				Type:                types.Int64Type,
				Required:            true,
			},
			"autostate": {
				MarkdownDescription: "Enable auto-state determination for VLAN",
				Type:                types.BoolType,
				Computed:            true,
			},
			"description": {
				MarkdownDescription: "Interface specific description",
				Type:                types.StringType,
				Computed:            true,
			},
			"shutdown": {
				MarkdownDescription: "Shutdown the selected interface",
				Type:                types.BoolType,
				Computed:            true,
			},
			"vrf_forwarding": {
				MarkdownDescription: "Configure forwarding table",
				Type:                types.StringType,
				Computed:            true,
			},
			"ipv4_address": {
				MarkdownDescription: "",
				Type:                types.StringType,
				Computed:            true,
			},
			"ipv4_address_mask": {
				MarkdownDescription: "",
				Type:                types.StringType,
				Computed:            true,
			},
			"unnumbered": {
				MarkdownDescription: "Enable IP processing without an explicit address",
				Type:                types.StringType,
				Computed:            true,
			},
			"ip_dhcp_relay_source_interface": {
				MarkdownDescription: "Set source interface for relayed messages",
				Type:                types.StringType,
				Computed:            true,
			},
			"ip_access_group_in": {
				MarkdownDescription: "",
				Type:                types.StringType,
				Computed:            true,
			},
			"ip_access_group_in_enable": {
				MarkdownDescription: "inbound packets",
				Type:                types.BoolType,
				Computed:            true,
			},
			"ip_access_group_out": {
				MarkdownDescription: "",
				Type:                types.StringType,
				Computed:            true,
			},
			"ip_access_group_out_enable": {
				MarkdownDescription: "outbound packets",
				Type:                types.BoolType,
				Computed:            true,
			},
			"helper_addresses": {
				MarkdownDescription: "Specify a destination address for UDP broadcasts",
				Computed:            true,
				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
					"address": {
						MarkdownDescription: "",
						Type:                types.StringType,
						Computed:            true,
					},
					"global": {
						MarkdownDescription: "Helper-address is global",
						Type:                types.BoolType,
						Computed:            true,
					},
					"vrf": {
						MarkdownDescription: "VRF name for helper-address (if different from interface VRF)",
						Type:                types.StringType,
						Computed:            true,
					},
				}),
			},
		},
	}, nil
}

func (t dataSourceInterfaceVLANType) NewDataSource(ctx context.Context, in tfsdk.Provider) (tfsdk.DataSource, diag.Diagnostics) {
	provider, diags := convertProviderType(in)

	return dataSourceInterfaceVLAN{
		provider: provider,
	}, diags
}

type dataSourceInterfaceVLAN struct {
	provider provider
}

func (d dataSourceInterfaceVLAN) Read(ctx context.Context, req tfsdk.ReadDataSourceRequest, resp *tfsdk.ReadDataSourceResponse) {
	var config InterfaceVLAN

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.getPath()))

	res, err := d.provider.clients[config.Device.Value].GetData(config.getPath())
	if res.StatusCode == 404 {
		config = InterfaceVLAN{Device: config.Device}
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
