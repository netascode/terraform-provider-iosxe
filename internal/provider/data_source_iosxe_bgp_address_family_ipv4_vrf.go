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

type dataSourceBGPAddressFamilyIPv4VRFType struct{}

func (t dataSourceBGPAddressFamilyIPv4VRFType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the BGP Address Family IPv4 VRF configuration.",

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
				Required:            true,
			},
			"af_name": {
				MarkdownDescription: "",
				Type:                types.StringType,
				Required:            true,
			},
			"vrfs": {
				MarkdownDescription: "",
				Computed:            true,
				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
					"name": {
						MarkdownDescription: "",
						Type:                types.StringType,
						Computed:            true,
					},
					"advertise_l2vpn_evpn": {
						MarkdownDescription: "Advertise/export prefixes to l2vpn evpn table",
						Type:                types.BoolType,
						Computed:            true,
					},
					"redistribute_connected": {
						MarkdownDescription: "Connected",
						Type:                types.BoolType,
						Computed:            true,
					},
					"redistribute_static": {
						MarkdownDescription: "Static routes",
						Type:                types.BoolType,
						Computed:            true,
					},
				}, tfsdk.ListNestedAttributesOptions{}),
			},
		},
	}, nil
}

func (t dataSourceBGPAddressFamilyIPv4VRFType) NewDataSource(ctx context.Context, in tfsdk.Provider) (tfsdk.DataSource, diag.Diagnostics) {
	provider, diags := convertProviderType(in)

	return dataSourceBGPAddressFamilyIPv4VRF{
		provider: provider,
	}, diags
}

type dataSourceBGPAddressFamilyIPv4VRF struct {
	provider provider
}

func (d dataSourceBGPAddressFamilyIPv4VRF) Read(ctx context.Context, req tfsdk.ReadDataSourceRequest, resp *tfsdk.ReadDataSourceResponse) {
	var config BGPAddressFamilyIPv4VRF

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.getPath()))

	res, err := d.provider.clients[config.Device.Value].GetData(config.getPath())
	if res.StatusCode == 404 {
		config = BGPAddressFamilyIPv4VRF{Device: config.Device}
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
