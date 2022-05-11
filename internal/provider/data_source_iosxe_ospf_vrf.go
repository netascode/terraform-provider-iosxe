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

type dataSourceOSPFVRFType struct{}

func (t dataSourceOSPFVRFType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the OSPF VRF configuration.",

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
			"process_id": {
				MarkdownDescription: "Process ID",
				Type:                types.Int64Type,
				Required:            true,
			},
			"vrf": {
				MarkdownDescription: "VPN Routing/Forwarding Instance",
				Type:                types.StringType,
				Required:            true,
			},
			"bfd_all_interfaces": {
				MarkdownDescription: "Enable BFD on all interfaces",
				Type:                types.BoolType,
				Computed:            true,
			},
			"default_information_originate": {
				MarkdownDescription: "Distribute a default route",
				Type:                types.BoolType,
				Computed:            true,
			},
			"default_information_originate_always": {
				MarkdownDescription: "Always advertise default route",
				Type:                types.BoolType,
				Computed:            true,
			},
			"default_metric": {
				MarkdownDescription: "Set metric of redistributed routes",
				Type:                types.Int64Type,
				Computed:            true,
			},
			"distance": {
				MarkdownDescription: "Administrative distance",
				Type:                types.Int64Type,
				Computed:            true,
			},
			"domain_tag": {
				MarkdownDescription: "OSPF domain-tag",
				Type:                types.Int64Type,
				Computed:            true,
			},
			"mpls_ldp_autoconfig": {
				MarkdownDescription: "Configure LDP automatic configuration",
				Type:                types.BoolType,
				Computed:            true,
			},
			"mpls_ldp_sync": {
				MarkdownDescription: "Configure LDP-IGP Synchronization",
				Type:                types.BoolType,
				Computed:            true,
			},
			"neighbor": {
				MarkdownDescription: "Specify a neighbor router",
				Computed:            true,
				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
					"ip": {
						MarkdownDescription: "Neighbor address",
						Type:                types.StringType,
						Computed:            true,
					},
					"priority": {
						MarkdownDescription: "OSPF priority of non-broadcast neighbor",
						Type:                types.Int64Type,
						Computed:            true,
					},
					"cost": {
						MarkdownDescription: "OSPF cost for point-to-multipoint neighbor",
						Type:                types.Int64Type,
						Computed:            true,
					},
				}, tfsdk.ListNestedAttributesOptions{}),
			},
			"network": {
				MarkdownDescription: "Enable routing on an IP network",
				Computed:            true,
				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
					"ip": {
						MarkdownDescription: "Network number",
						Type:                types.StringType,
						Computed:            true,
					},
					"wildcard": {
						MarkdownDescription: "OSPF wild card bits",
						Type:                types.StringType,
						Computed:            true,
					},
					"area": {
						MarkdownDescription: "Set the OSPF area ID",
						Type:                types.StringType,
						Computed:            true,
					},
				}, tfsdk.ListNestedAttributesOptions{}),
			},
			"priority": {
				MarkdownDescription: "OSPF topology priority",
				Type:                types.Int64Type,
				Computed:            true,
			},
			"router_id": {
				MarkdownDescription: "Override configured router identifier (peers will reset)",
				Type:                types.StringType,
				Computed:            true,
			},
			"shutdown": {
				MarkdownDescription: "Shutdown the OSPF protocol under the current instance",
				Type:                types.BoolType,
				Computed:            true,
			},
			"summary_address": {
				MarkdownDescription: "Configure IP address summaries",
				Computed:            true,
				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
					"ip": {
						MarkdownDescription: "IP summary address",
						Type:                types.StringType,
						Computed:            true,
					},
					"mask": {
						MarkdownDescription: "Summary mask",
						Type:                types.StringType,
						Computed:            true,
					},
				}, tfsdk.ListNestedAttributesOptions{}),
			},
		},
	}, nil
}

func (t dataSourceOSPFVRFType) NewDataSource(ctx context.Context, in tfsdk.Provider) (tfsdk.DataSource, diag.Diagnostics) {
	provider, diags := convertProviderType(in)

	return dataSourceOSPFVRF{
		provider: provider,
	}, diags
}

type dataSourceOSPFVRF struct {
	provider provider
}

func (d dataSourceOSPFVRF) Read(ctx context.Context, req tfsdk.ReadDataSourceRequest, resp *tfsdk.ReadDataSourceResponse) {
	var config OSPFVRF

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.getPath()))

	res, err := d.provider.clients[config.Device.Value].GetData(config.getPath())
	if res.StatusCode == 404 {
		config = OSPFVRF{Device: config.Device}
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
