// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

type dataSourceAccessListExtendedType struct{}

func (t dataSourceAccessListExtendedType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Access List Extended configuration.",

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
			"name": {
				MarkdownDescription: "",
				Type:                types.StringType,
				Required:            true,
			},
			"entries": {
				MarkdownDescription: "",
				Computed:            true,
				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
					"sequence": {
						MarkdownDescription: "",
						Type:                types.Int64Type,
						Computed:            true,
					},
					"remark": {
						MarkdownDescription: "Access list entry comment",
						Type:                types.StringType,
						Computed:            true,
					},
					"ace_rule_action": {
						MarkdownDescription: "",
						Type:                types.StringType,
						Computed:            true,
					},
					"ace_rule_protocol": {
						MarkdownDescription: "",
						Type:                types.StringType,
						Computed:            true,
					},
					"service_object_group": {
						MarkdownDescription: "Service object group name",
						Type:                types.StringType,
						Computed:            true,
					},
					"source_prefix": {
						MarkdownDescription: "",
						Type:                types.StringType,
						Computed:            true,
					},
					"source_prefix_mask": {
						MarkdownDescription: "",
						Type:                types.StringType,
						Computed:            true,
					},
					"source_any": {
						MarkdownDescription: "Any source host",
						Type:                types.BoolType,
						Computed:            true,
					},
					"source_host": {
						MarkdownDescription: "A single source host(DEPRECATED - use host-address)",
						Type:                types.StringType,
						Computed:            true,
					},
					"source_object_group": {
						MarkdownDescription: "Source network object group",
						Type:                types.StringType,
						Computed:            true,
					},
					"source_port_equal": {
						MarkdownDescription: "Match only packets on a given port number up to 10 ports",
						Type:                types.StringType,
						Computed:            true,
					},
					"source_port_greater_than": {
						MarkdownDescription: "Match only packets with a greater port number",
						Type:                types.StringType,
						Computed:            true,
					},
					"source_port_lesser_than": {
						MarkdownDescription: "Match only packets with a lower port number",
						Type:                types.StringType,
						Computed:            true,
					},
					"source_port_range_from": {
						MarkdownDescription: "Match only packets in the range of port numbers",
						Type:                types.StringType,
						Computed:            true,
					},
					"source_port_range_to": {
						MarkdownDescription: "Match only packets in the range of port numbers",
						Type:                types.StringType,
						Computed:            true,
					},
					"destination_prefix": {
						MarkdownDescription: "",
						Type:                types.StringType,
						Computed:            true,
					},
					"destination_prefix_mask": {
						MarkdownDescription: "",
						Type:                types.StringType,
						Computed:            true,
					},
					"destination_any": {
						MarkdownDescription: "Any destination host",
						Type:                types.BoolType,
						Computed:            true,
					},
					"destination_host": {
						MarkdownDescription: "A single destination host(DEPRECATED - use dst-host-address)",
						Type:                types.StringType,
						Computed:            true,
					},
					"destination_object_group": {
						MarkdownDescription: "Destination network object group",
						Type:                types.StringType,
						Computed:            true,
					},
					"destination_port_equal": {
						MarkdownDescription: "Match only packets on a given port number up to 10 ports",
						Type:                types.StringType,
						Computed:            true,
					},
					"destination_port_greater_than": {
						MarkdownDescription: "Match only packets with a greater port number",
						Type:                types.StringType,
						Computed:            true,
					},
					"destination_port_lesser_than": {
						MarkdownDescription: "Match only packets with a lower port number",
						Type:                types.StringType,
						Computed:            true,
					},
					"destination_port_range_from": {
						MarkdownDescription: "Match only packets in the range of port numbers",
						Type:                types.StringType,
						Computed:            true,
					},
					"destination_port_range_to": {
						MarkdownDescription: "Match only packets in the range of port numbers",
						Type:                types.StringType,
						Computed:            true,
					},
					"ack": {
						MarkdownDescription: "Match on the ACK bit",
						Type:                types.BoolType,
						Computed:            true,
					},
					"fin": {
						MarkdownDescription: "Match on the FIN bit",
						Type:                types.BoolType,
						Computed:            true,
					},
					"psh": {
						MarkdownDescription: "Match on the PSH bit",
						Type:                types.BoolType,
						Computed:            true,
					},
					"rst": {
						MarkdownDescription: "Match on the RST bit",
						Type:                types.BoolType,
						Computed:            true,
					},
					"syn": {
						MarkdownDescription: "Match on the SYN bit",
						Type:                types.BoolType,
						Computed:            true,
					},
					"urg": {
						MarkdownDescription: "Match on the URG bit",
						Type:                types.BoolType,
						Computed:            true,
					},
					"established": {
						MarkdownDescription: "Match established connections",
						Type:                types.BoolType,
						Computed:            true,
					},
					"dscp": {
						MarkdownDescription: "Match packets with given dscp value",
						Type:                types.StringType,
						Computed:            true,
					},
					"fragments": {
						MarkdownDescription: "Check non-initial fragments",
						Type:                types.BoolType,
						Computed:            true,
					},
					"precedence": {
						MarkdownDescription: "Match packets with given precedence value",
						Type:                types.StringType,
						Computed:            true,
					},
					"tos": {
						MarkdownDescription: "Match packets with given TOS value",
						Type:                types.StringType,
						Computed:            true,
					},
				}),
			},
		},
	}, nil
}

func (t dataSourceAccessListExtendedType) NewDataSource(ctx context.Context, in provider.Provider) (datasource.DataSource, diag.Diagnostics) {
	provider, diags := convertProviderType(in)

	return dataSourceAccessListExtended{
		provider: provider,
	}, diags
}

type dataSourceAccessListExtended struct {
	provider iosxeProvider
}

func (d dataSourceAccessListExtended) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config AccessListExtended

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.getPath()))

	res, err := d.provider.clients[config.Device.Value].GetData(config.getPath())
	if res.StatusCode == 404 {
		config = AccessListExtended{Device: config.Device}
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
