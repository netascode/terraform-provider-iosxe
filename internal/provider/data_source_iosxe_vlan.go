// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-restconf"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &VLANDataSource{}
	_ datasource.DataSourceWithConfigure = &VLANDataSource{}
)

func NewVLANDataSource() datasource.DataSource {
	return &VLANDataSource{}
}

type VLANDataSource struct {
	clients map[string]*restconf.Client
}

func (d *VLANDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vlan"
}

func (d *VLANDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the VLAN configuration.",

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The path of the retrieved object.",
				Computed:            true,
			},
			"vlan_id": schema.Int64Attribute{
				MarkdownDescription: "a single VLAN id (allowed value range 1-4094)or Comma-separated VLAN id range.e.g. 99 or 1-30 or  1-20,30,40-50",
				Required:            true,
			},
			"remote_span": schema.BoolAttribute{
				MarkdownDescription: "Configure as Remote SPAN VLAN",
				Computed:            true,
			},
			"private_vlan_primary": schema.BoolAttribute{
				MarkdownDescription: "Configure the VLAN as a primary private VLAN",
				Computed:            true,
			},
			"private_vlan_association": schema.StringAttribute{
				MarkdownDescription: "Configure association between private VLANs",
				Computed:            true,
			},
			"private_vlan_community": schema.BoolAttribute{
				MarkdownDescription: "Configure the VLAN as a community private VLAN",
				Computed:            true,
			},
			"private_vlan_isolated": schema.BoolAttribute{
				MarkdownDescription: "Configure the VLAN as an isolated private VLAN",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Ascii name of the VLAN",
				Computed:            true,
			},
			"shutdown": schema.BoolAttribute{
				MarkdownDescription: "Shutdown VLAN switching",
				Computed:            true,
			},
		},
	}
}

func (d *VLANDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.clients = req.ProviderData.(map[string]*restconf.Client)
}

func (d *VLANDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config VLAN

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if _, ok := d.clients[config.Device.ValueString()]; !ok {
		resp.Diagnostics.AddAttributeError(path.Root("device"), "Invalid device", fmt.Sprintf("Device '%s' does not exist in provider configuration.", config.Device.ValueString()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.getPath()))

	res, err := d.clients[config.Device.ValueString()].GetData(config.getPath())
	if res.StatusCode == 404 {
		config = VLAN{Device: config.Device}
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
