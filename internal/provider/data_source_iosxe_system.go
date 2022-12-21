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
	_ datasource.DataSource              = &SystemDataSource{}
	_ datasource.DataSourceWithConfigure = &SystemDataSource{}
)

func NewSystemDataSource() datasource.DataSource {
	return &SystemDataSource{}
}

type SystemDataSource struct {
	clients map[string]*restconf.Client
}

func (d *SystemDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_system"
}

func (d *SystemDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the System configuration.",

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The path of the retrieved object.",
				Computed:            true,
			},
			"hostname": schema.StringAttribute{
				MarkdownDescription: "Set system's network name",
				Computed:            true,
			},
			"ip_routing": schema.BoolAttribute{
				MarkdownDescription: "Enable or disable IP routing",
				Computed:            true,
			},
			"ipv6_unicast_routing": schema.BoolAttribute{
				MarkdownDescription: "Enable unicast routing",
				Computed:            true,
			},
			"mtu": schema.Int64Attribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"ip_source_route": schema.BoolAttribute{
				MarkdownDescription: "Process packets with source routing header options",
				Computed:            true,
			},
			"ip_domain_lookup": schema.BoolAttribute{
				MarkdownDescription: "Enable IP Domain Name System hostname translation",
				Computed:            true,
			},
			"ip_domain_name": schema.StringAttribute{
				MarkdownDescription: "Define the default domain name",
				Computed:            true,
			},
			"login_delay": schema.Int64Attribute{
				MarkdownDescription: "Set delay between successive fail login",
				Computed:            true,
			},
			"login_on_failure": schema.BoolAttribute{
				MarkdownDescription: "Set options for failed login attempt",
				Computed:            true,
			},
			"login_on_failure_log": schema.BoolAttribute{
				MarkdownDescription: "Generate syslogs on failure logins",
				Computed:            true,
			},
			"login_on_success": schema.BoolAttribute{
				MarkdownDescription: "Set options for successful login attempt",
				Computed:            true,
			},
			"login_on_success_log": schema.BoolAttribute{
				MarkdownDescription: "Generate syslogs on successful logins",
				Computed:            true,
			},
			"multicast_routing": schema.BoolAttribute{
				MarkdownDescription: "Enable IP multicast forwarding",
				Computed:            true,
			},
			"multicast_routing_switch": schema.BoolAttribute{
				MarkdownDescription: "Enable IP multicast forwarding, some XE devices use this option instead of `multicast_routing`.",
				Computed:            true,
			},
			"multicast_routing_distributed": schema.BoolAttribute{
				MarkdownDescription: "Distributed multicast switching",
				Computed:            true,
			},
			"multicast_routing_vrfs": schema.ListNestedAttribute{
				MarkdownDescription: "Select VPN Routing/Forwarding instance",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"vrf": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"distributed": schema.BoolAttribute{
							MarkdownDescription: "Distributed multicast switching",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *SystemDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.clients = req.ProviderData.(map[string]*restconf.Client)
}

func (d *SystemDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config System

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.getPath()))

	res, err := d.clients[config.Device.ValueString()].GetData(config.getPath())
	if res.StatusCode == 404 {
		config = System{Device: config.Device}
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
