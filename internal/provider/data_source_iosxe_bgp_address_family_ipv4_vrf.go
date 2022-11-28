// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-restconf"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &BGPAddressFamilyIPv4VRFDataSource{}
	_ datasource.DataSourceWithConfigure = &BGPAddressFamilyIPv4VRFDataSource{}
)

func NewBGPAddressFamilyIPv4VRFDataSource() datasource.DataSource {
	return &BGPAddressFamilyIPv4VRFDataSource{}
}

type BGPAddressFamilyIPv4VRFDataSource struct {
	clients map[string]*restconf.Client
}

func (d *BGPAddressFamilyIPv4VRFDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_bgp_address_family_ipv4_vrf"
}

func (d *BGPAddressFamilyIPv4VRFDataSource) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
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
				}),
			},
		},
	}, nil
}

func (d *BGPAddressFamilyIPv4VRFDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.clients = req.ProviderData.(map[string]*restconf.Client)
}

func (d *BGPAddressFamilyIPv4VRFDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config BGPAddressFamilyIPv4VRF

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.getPath()))

	res, err := d.clients[config.Device.ValueString()].GetData(config.getPath())
	if res.StatusCode == 404 {
		config = BGPAddressFamilyIPv4VRF{Device: config.Device}
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
