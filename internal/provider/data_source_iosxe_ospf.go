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
	_ datasource.DataSource              = &OSPFDataSource{}
	_ datasource.DataSourceWithConfigure = &OSPFDataSource{}
)

func NewOSPFDataSource() datasource.DataSource {
	return &OSPFDataSource{}
}

type OSPFDataSource struct {
	clients map[string]*restconf.Client
}

func (d *OSPFDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ospf"
}

func (d *OSPFDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the OSPF configuration.",

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The path of the retrieved object.",
				Computed:            true,
			},
			"process_id": schema.Int64Attribute{
				MarkdownDescription: "Process ID",
				Required:            true,
			},
			"bfd_all_interfaces": schema.BoolAttribute{
				MarkdownDescription: "Enable BFD on all interfaces",
				Computed:            true,
			},
			"default_information_originate": schema.BoolAttribute{
				MarkdownDescription: "Distribute a default route",
				Computed:            true,
			},
			"default_information_originate_always": schema.BoolAttribute{
				MarkdownDescription: "Always advertise default route",
				Computed:            true,
			},
			"default_metric": schema.Int64Attribute{
				MarkdownDescription: "Set metric of redistributed routes",
				Computed:            true,
			},
			"distance": schema.Int64Attribute{
				MarkdownDescription: "Administrative distance",
				Computed:            true,
			},
			"domain_tag": schema.Int64Attribute{
				MarkdownDescription: "OSPF domain-tag",
				Computed:            true,
			},
			"mpls_ldp_autoconfig": schema.BoolAttribute{
				MarkdownDescription: "Configure LDP automatic configuration",
				Computed:            true,
			},
			"mpls_ldp_sync": schema.BoolAttribute{
				MarkdownDescription: "Configure LDP-IGP Synchronization",
				Computed:            true,
			},
			"neighbor": schema.ListNestedAttribute{
				MarkdownDescription: "Specify a neighbor router",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"ip": schema.StringAttribute{
							MarkdownDescription: "Neighbor address",
							Computed:            true,
						},
						"priority": schema.Int64Attribute{
							MarkdownDescription: "OSPF priority of non-broadcast neighbor",
							Computed:            true,
						},
						"cost": schema.Int64Attribute{
							MarkdownDescription: "OSPF cost for point-to-multipoint neighbor",
							Computed:            true,
						},
					},
				},
			},
			"network": schema.ListNestedAttribute{
				MarkdownDescription: "Enable routing on an IP network",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"ip": schema.StringAttribute{
							MarkdownDescription: "Network number",
							Computed:            true,
						},
						"wildcard": schema.StringAttribute{
							MarkdownDescription: "OSPF wild card bits",
							Computed:            true,
						},
						"area": schema.StringAttribute{
							MarkdownDescription: "Set the OSPF area ID",
							Computed:            true,
						},
					},
				},
			},
			"priority": schema.Int64Attribute{
				MarkdownDescription: "OSPF topology priority",
				Computed:            true,
			},
			"router_id": schema.StringAttribute{
				MarkdownDescription: "Override configured router identifier (peers will reset)",
				Computed:            true,
			},
			"shutdown": schema.BoolAttribute{
				MarkdownDescription: "Shutdown the OSPF protocol under the current instance",
				Computed:            true,
			},
			"summary_address": schema.ListNestedAttribute{
				MarkdownDescription: "Configure IP address summaries",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"ip": schema.StringAttribute{
							MarkdownDescription: "IP summary address",
							Computed:            true,
						},
						"mask": schema.StringAttribute{
							MarkdownDescription: "Summary mask",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *OSPFDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.clients = req.ProviderData.(map[string]*restconf.Client)
}

func (d *OSPFDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config OSPF

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
		config = OSPF{Device: config.Device}
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
