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

type dataSourceBGPAddressFamilyL2VPNType struct{}

func (t dataSourceBGPAddressFamilyL2VPNType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the BGP Address Family L2VPN configuration.",

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
			"asn": {
				MarkdownDescription: "",
				Type:                types.StringType,
				Computed:            true,
			},
			"af_name": {
				MarkdownDescription: "",
				Type:                types.StringType,
				Required:            true,
			},
		},
	}, nil
}

func (t dataSourceBGPAddressFamilyL2VPNType) NewDataSource(ctx context.Context, in tfsdk.Provider) (tfsdk.DataSource, diag.Diagnostics) {
	provider, diags := convertProviderType(in)

	return dataSourceBGPAddressFamilyL2VPN{
		provider: provider,
	}, diags
}

type dataSourceBGPAddressFamilyL2VPN struct {
	provider provider
}

func (d dataSourceBGPAddressFamilyL2VPN) Read(ctx context.Context, req tfsdk.ReadDataSourceRequest, resp *tfsdk.ReadDataSourceResponse) {
	var config, state BGPAddressFamilyL2VPN

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.getPath()))

	res, err := d.provider.clients[config.Device.Value].GetData(config.getPath())
	if res.StatusCode != 404 {
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
			return
		}

		state.fromBody(res.Res)
	}
	state.fromPlan(config)
	state.Id.Value = config.getPath()

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.getPath()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}