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

type dataSourceVLANConfigurationType struct{}

func (t dataSourceVLANConfigurationType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the VLAN Configuration configuration.",

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
			"vlan_id": {
				MarkdownDescription: "VLAN ID List Eg. 1-10,15",
				Type:                types.Int64Type,
				Required:            true,
			},
			"vni": {
				MarkdownDescription: "VxLAN VNI value",
				Type:                types.Int64Type,
				Computed:            true,
			},
			"access_vfi": {
				MarkdownDescription: "Enter VFI name",
				Type:                types.StringType,
				Computed:            true,
			},
			"evpn_instance": {
				MarkdownDescription: "",
				Type:                types.Int64Type,
				Computed:            true,
			},
			"evpn_instance_vni": {
				MarkdownDescription: "VxLAN VNI value",
				Type:                types.Int64Type,
				Computed:            true,
			},
		},
	}, nil
}

func (t dataSourceVLANConfigurationType) NewDataSource(ctx context.Context, in tfsdk.Provider) (tfsdk.DataSource, diag.Diagnostics) {
	provider, diags := convertProviderType(in)

	return dataSourceVLANConfiguration{
		provider: provider,
	}, diags
}

type dataSourceVLANConfiguration struct {
	provider provider
}

func (d dataSourceVLANConfiguration) Read(ctx context.Context, req tfsdk.ReadDataSourceRequest, resp *tfsdk.ReadDataSourceResponse) {
	var config VLANConfiguration

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.getPath()))

	res, err := d.provider.clients[config.Device.Value].GetData(config.getPath())
	if res.StatusCode == 404 {
		config = VLANConfiguration{Device: config.Device}
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
