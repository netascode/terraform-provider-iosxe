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

type dataSourceDHCPType struct{}

func (t dataSourceDHCPType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the DHCP configuration.",

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
			"compatibility_suboption_link_selection": {
				MarkdownDescription: "",
				Type:                types.StringType,
				Computed:            true,
			},
			"compatibility_suboption_server_override": {
				MarkdownDescription: "",
				Type:                types.StringType,
				Computed:            true,
			},
			"relay_information_trust_all": {
				MarkdownDescription: "Received DHCP packets may contain relay info option with zero giaddr",
				Type:                types.BoolType,
				Computed:            true,
			},
			"relay_information_option_default": {
				MarkdownDescription: "Default option, no vpn",
				Type:                types.BoolType,
				Computed:            true,
			},
			"relay_information_option_vpn": {
				MarkdownDescription: "Insert VPN sub-options and change the giaddr to the outgoing interface",
				Type:                types.BoolType,
				Computed:            true,
			},
			"snooping": {
				MarkdownDescription: "DHCP Snooping",
				Type:                types.BoolType,
				Computed:            true,
			},
			"snooping_vlans": {
				MarkdownDescription: "DHCP Snooping vlan (Deprecated, use vlan-list)",
				Computed:            true,
				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
					"vlan_id": {
						MarkdownDescription: "",
						Type:                types.Int64Type,
						Computed:            true,
					},
				}, tfsdk.ListNestedAttributesOptions{}),
			},
		},
	}, nil
}

func (t dataSourceDHCPType) NewDataSource(ctx context.Context, in tfsdk.Provider) (tfsdk.DataSource, diag.Diagnostics) {
	provider, diags := convertProviderType(in)

	return dataSourceDHCP{
		provider: provider,
	}, diags
}

type dataSourceDHCP struct {
	provider provider
}

func (d dataSourceDHCP) Read(ctx context.Context, req tfsdk.ReadDataSourceRequest, resp *tfsdk.ReadDataSourceResponse) {
	var config DHCP

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.getPath()))

	res, err := d.provider.clients[config.Device.Value].GetData(config.getPath())
	if res.StatusCode == 404 {
		config = DHCP{Device: config.Device}
	} else {
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
			return
		}

		config.fromBody(res.Res)
	}

	config.Id = types.String{Value: config.getPath()}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.getPath()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}