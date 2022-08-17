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

type dataSourceSystemType struct{}

func (t dataSourceSystemType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the System configuration.",

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
			"hostname": {
				MarkdownDescription: "Set system's network name",
				Type:                types.StringType,
				Computed:            true,
			},
			"ip_routing": {
				MarkdownDescription: "Enable or disable IP routing",
				Type:                types.BoolType,
				Computed:            true,
			},
			"ipv6_unicast_routing": {
				MarkdownDescription: "Enable unicast routing",
				Type:                types.BoolType,
				Computed:            true,
			},
			"mtu": {
				MarkdownDescription: "",
				Type:                types.Int64Type,
				Computed:            true,
			},
			"ip_source_route": {
				MarkdownDescription: "Process packets with source routing header options",
				Type:                types.BoolType,
				Computed:            true,
			},
			"ip_domain_lookup": {
				MarkdownDescription: "Enable IP Domain Name System hostname translation",
				Type:                types.BoolType,
				Computed:            true,
			},
			"ip_domain_name": {
				MarkdownDescription: "Define the default domain name",
				Type:                types.StringType,
				Computed:            true,
			},
			"login_delay": {
				MarkdownDescription: "Set delay between successive fail login",
				Type:                types.Int64Type,
				Computed:            true,
			},
			"login_on_failure": {
				MarkdownDescription: "Set options for failed login attempt",
				Type:                types.BoolType,
				Computed:            true,
			},
			"login_on_failure_log": {
				MarkdownDescription: "Generate syslogs on failure logins",
				Type:                types.BoolType,
				Computed:            true,
			},
			"login_on_success": {
				MarkdownDescription: "Set options for successful login attempt",
				Type:                types.BoolType,
				Computed:            true,
			},
			"login_on_success_log": {
				MarkdownDescription: "Generate syslogs on successful logins",
				Type:                types.BoolType,
				Computed:            true,
			},
			"multicast_routing": {
				MarkdownDescription: "Enable IP multicast forwarding",
				Type:                types.BoolType,
				Computed:            true,
			},
			"multicast_routing_switch": {
				MarkdownDescription: "Enable IP multicast forwarding, some XE devices use this option instead of `multicast_routing`.",
				Type:                types.BoolType,
				Computed:            true,
			},
			"multicast_routing_distributed": {
				MarkdownDescription: "Distributed multicast switching",
				Type:                types.BoolType,
				Computed:            true,
			},
			"multicast_routing_vrfs": {
				MarkdownDescription: "Select VPN Routing/Forwarding instance",
				Computed:            true,
				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
					"vrf": {
						MarkdownDescription: "",
						Type:                types.StringType,
						Computed:            true,
					},
					"distributed": {
						MarkdownDescription: "Distributed multicast switching",
						Type:                types.BoolType,
						Computed:            true,
					},
				}),
			},
		},
	}, nil
}

func (t dataSourceSystemType) NewDataSource(ctx context.Context, in tfsdk.Provider) (tfsdk.DataSource, diag.Diagnostics) {
	provider, diags := convertProviderType(in)

	return dataSourceSystem{
		provider: provider,
	}, diags
}

type dataSourceSystem struct {
	provider provider
}

func (d dataSourceSystem) Read(ctx context.Context, req tfsdk.ReadDataSourceRequest, resp *tfsdk.ReadDataSourceResponse) {
	var config System

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.getPath()))

	res, err := d.provider.clients[config.Device.Value].GetData(config.getPath())
	if res.StatusCode == 404 {
		config = System{Device: config.Device}
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
