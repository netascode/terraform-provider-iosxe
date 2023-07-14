// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

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
	_ datasource.DataSource              = &LoggingIPv6HostVRFTransportDataSource{}
	_ datasource.DataSourceWithConfigure = &LoggingIPv6HostVRFTransportDataSource{}
)

func NewLoggingIPv6HostVRFTransportDataSource() datasource.DataSource {
	return &LoggingIPv6HostVRFTransportDataSource{}
}

type LoggingIPv6HostVRFTransportDataSource struct {
	clients map[string]*restconf.Client
}

func (d *LoggingIPv6HostVRFTransportDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_logging_ipv6_host_vrf_transport"
}

func (d *LoggingIPv6HostVRFTransportDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Logging IPv6 Host VRF Transport configuration.",

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The path of the retrieved object.",
				Computed:            true,
			},
			"ipv6_host": schema.StringAttribute{
				MarkdownDescription: "",
				Required:            true,
			},
			"vrf": schema.StringAttribute{
				MarkdownDescription: "Set VRF option",
				Required:            true,
			},
			"transport_udp_ports": schema.ListNestedAttribute{
				MarkdownDescription: "Port Number List",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"port_number": schema.Int64Attribute{
							MarkdownDescription: "Specify the UDP port number (default=514)",
							Computed:            true,
						},
					},
				},
			},
			"transport_tcp_ports": schema.ListNestedAttribute{
				MarkdownDescription: "Port Number List",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"port_number": schema.Int64Attribute{
							MarkdownDescription: "Specify the TCP port number (default=601)",
							Computed:            true,
						},
					},
				},
			},
			"transport_tls_ports": schema.ListNestedAttribute{
				MarkdownDescription: "Port Number List",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"port_number": schema.Int64Attribute{
							MarkdownDescription: "Specify the TLS port number (default=6514)",
							Computed:            true,
						},
						"profile": schema.StringAttribute{
							MarkdownDescription: "Specify the TLS profile",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *LoggingIPv6HostVRFTransportDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.clients = req.ProviderData.(map[string]*restconf.Client)
}

func (d *LoggingIPv6HostVRFTransportDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config LoggingIPv6HostVRFTransportData

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
		config = LoggingIPv6HostVRFTransportData{Device: config.Device}
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
