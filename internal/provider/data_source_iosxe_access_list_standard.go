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

type dataSourceAccessListStandardType struct{}

func (t dataSourceAccessListStandardType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Access List Standard configuration.",

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
					"deny_prefix": {
						MarkdownDescription: "Network address prefix (DEPRECATED - use ipv4-address-prefix)",
						Type:                types.StringType,
						Computed:            true,
					},
					"deny_prefix_mask": {
						MarkdownDescription: "Wildcard bits",
						Type:                types.StringType,
						Computed:            true,
					},
					"deny_any": {
						MarkdownDescription: "Any source prefix",
						Type:                types.BoolType,
						Computed:            true,
					},
					"deny_host": {
						MarkdownDescription: "A single source host (DEPRECATED - use host-address)",
						Type:                types.StringType,
						Computed:            true,
					},
					"permit_prefix": {
						MarkdownDescription: "Network address prefix (DEPRECATED - use ipv4-address-prefix)",
						Type:                types.StringType,
						Computed:            true,
					},
					"permit_prefix_mask": {
						MarkdownDescription: "Wildcard bits",
						Type:                types.StringType,
						Computed:            true,
					},
					"permit_any": {
						MarkdownDescription: "Any source prefix",
						Type:                types.BoolType,
						Computed:            true,
					},
					"permit_host": {
						MarkdownDescription: "A single source host (DEPRECATED - use host-address)",
						Type:                types.StringType,
						Computed:            true,
					},
				}),
			},
		},
	}, nil
}

func (t dataSourceAccessListStandardType) NewDataSource(ctx context.Context, in provider.Provider) (datasource.DataSource, diag.Diagnostics) {
	provider, diags := convertProviderType(in)

	return dataSourceAccessListStandard{
		provider: provider,
	}, diags
}

type dataSourceAccessListStandard struct {
	provider iosxeProvider
}

func (d dataSourceAccessListStandard) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config AccessListStandard

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.getPath()))

	res, err := d.provider.clients[config.Device.Value].GetData(config.getPath())
	if res.StatusCode == 404 {
		config = AccessListStandard{Device: config.Device}
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