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
	_ datasource.DataSource              = &SNMPServerGroupDataSource{}
	_ datasource.DataSourceWithConfigure = &SNMPServerGroupDataSource{}
)

func NewSNMPServerGroupDataSource() datasource.DataSource {
	return &SNMPServerGroupDataSource{}
}

type SNMPServerGroupDataSource struct {
	clients map[string]*restconf.Client
}

func (d *SNMPServerGroupDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_snmp_server_group"
}

func (d *SNMPServerGroupDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the SNMP Server Group configuration.",

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The path of the retrieved object.",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "",
				Required:            true,
			},
			"v3_security": schema.ListNestedAttribute{
				MarkdownDescription: "group using security Level",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"security_level": schema.StringAttribute{
							MarkdownDescription: "security level type",
							Computed:            true,
						},
						"context_node": schema.StringAttribute{
							MarkdownDescription: "specify a context to associate these views for the group",
							Computed:            true,
						},
						"match_node": schema.StringAttribute{
							MarkdownDescription: "context name match criteria",
							Computed:            true,
						},
						"read_node": schema.StringAttribute{
							MarkdownDescription: "specify a read view for the group",
							Computed:            true,
						},
						"write_node": schema.StringAttribute{
							MarkdownDescription: "specify a write view for the group",
							Computed:            true,
						},
						"notify_node": schema.StringAttribute{
							MarkdownDescription: "specify a notify view for the group",
							Computed:            true,
						},
						"access_ipv6_acl": schema.StringAttribute{
							MarkdownDescription: "Specify IPv6 Named Access-List",
							Computed:            true,
						},
						"access_standard_acl": schema.Int64Attribute{
							MarkdownDescription: "Standard IP Access-list allowing access with this community string",
							Computed:            true,
						},
						"access_acl_name": schema.StringAttribute{
							MarkdownDescription: "Access-list name",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *SNMPServerGroupDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.clients = req.ProviderData.(map[string]*restconf.Client)
}

func (d *SNMPServerGroupDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config SNMPServerGroup

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.getPath()))

	res, err := d.clients[config.Device.ValueString()].GetData(config.getPath())
	if res.StatusCode == 404 {
		config = SNMPServerGroup{Device: config.Device}
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
