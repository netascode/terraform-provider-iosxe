package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/terraform-provider-iosxe/internal/provider/helpers"
)

type dataSourceRestconfType struct{}

func (t dataSourceRestconfType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can retrieve one or more attributes via RESTCONF.",

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
			"path": {
				MarkdownDescription: "A RESTCONF path, e.g. `openconfig-interfaces:interfaces`.",
				Type:                types.StringType,
				Required:            true,
			},
			"attributes": {
				MarkdownDescription: "Map of key-value pairs which represents the attributes and its values.",
				Type:                types.MapType{ElemType: types.StringType},
				Computed:            true,
			},
		},
	}, nil
}

func (t dataSourceRestconfType) NewDataSource(ctx context.Context, in tfsdk.Provider) (tfsdk.DataSource, diag.Diagnostics) {
	provider, diags := convertProviderType(in)

	return dataSourceRestconf{
		provider: provider,
	}, diags
}

type dataSourceRestconf struct {
	provider provider
}

func (d dataSourceRestconf) Read(ctx context.Context, req tfsdk.ReadDataSourceRequest, resp *tfsdk.ReadDataSourceResponse) {
	var config, state RestconfDataSource

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.Value))

	res, err := d.provider.clients[config.Device.Value].GetData(config.Path.Value)
	if res.StatusCode == 404 {
		state.Attributes.Elems = map[string]attr.Value{}
	} else {
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
			return
		}

		state.Path.Value = config.Path.Value
		state.Id.Value = config.Path.Value

		attributes := make(map[string]attr.Value)

		for attr, value := range res.Res.Get(helpers.LastElement(config.Path.Value)).Map() {
			// handle empty maps
			if value.IsObject() && len(value.Map()) == 0 {
				attributes[attr] = types.String{Value: ""}
			} else if value.Raw == "[null]" {
				attributes[attr] = types.String{Value: ""}
			} else {
				attributes[attr] = types.String{Value: value.String()}
			}
		}
		state.Attributes.Elems = attributes
		state.Attributes.ElemType = types.StringType
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.Value))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}
