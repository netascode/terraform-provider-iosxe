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

type dataSourceLoggingIPv4HostVRFTransportType struct{}

func (t dataSourceLoggingIPv4HostVRFTransportType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Logging IPv4 Host VRF Transport configuration.",

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
			"ipv4_host": {
				MarkdownDescription: "",
				Type:                types.StringType,
				Required:            true,
			},
			"vrf": {
				MarkdownDescription: "Set VRF option",
				Type:                types.StringType,
				Required:            true,
			},
			"transport_udp_ports": {
				MarkdownDescription: "Port Number List",
				Computed:            true,
				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
					"port_number": {
						MarkdownDescription: "Specify the UDP port number (default=514)",
						Type:                types.Int64Type,
						Computed:            true,
					},
				}),
			},
			"transport_tcp_ports": {
				MarkdownDescription: "Port Number List",
				Computed:            true,
				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
					"port_number": {
						MarkdownDescription: "Specify the TCP port number (default=601)",
						Type:                types.Int64Type,
						Computed:            true,
					},
				}),
			},
			"transport_tls_ports": {
				MarkdownDescription: "Port Number List",
				Computed:            true,
				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
					"port_number": {
						MarkdownDescription: "Specify the TLS port number (default=6514)",
						Type:                types.Int64Type,
						Computed:            true,
					},
					"profile": {
						MarkdownDescription: "Specify the TLS profile",
						Type:                types.StringType,
						Computed:            true,
					},
				}),
			},
		},
	}, nil
}

func (t dataSourceLoggingIPv4HostVRFTransportType) NewDataSource(ctx context.Context, in tfsdk.Provider) (tfsdk.DataSource, diag.Diagnostics) {
	provider, diags := convertProviderType(in)

	return dataSourceLoggingIPv4HostVRFTransport{
		provider: provider,
	}, diags
}

type dataSourceLoggingIPv4HostVRFTransport struct {
	provider provider
}

func (d dataSourceLoggingIPv4HostVRFTransport) Read(ctx context.Context, req tfsdk.ReadDataSourceRequest, resp *tfsdk.ReadDataSourceResponse) {
	var config LoggingIPv4HostVRFTransport

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.getPath()))

	res, err := d.provider.clients[config.Device.Value].GetData(config.getPath())
	if res.StatusCode == 404 {
		config = LoggingIPv4HostVRFTransport{Device: config.Device}
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
