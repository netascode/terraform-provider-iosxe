//go:build ignore
// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

type dataSource{{camelCase .Name}}Type struct{}

func (t dataSource{{camelCase .Name}}Type) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "{{.DsDescription}}",

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
			{{- range  .Attributes}}
			"{{.TfName}}": {
				MarkdownDescription: "{{.Description}}",
				{{- if ne .Type "List"}}
				Type:                types.{{.Type}}Type,
				{{- end}}
				{{- if or (eq .Id true) (eq .Reference true)}}
				Required:            true,
				{{- else}}
				Computed:            true,
				{{- end}}
				{{- if eq .Type "List"}}
				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
					{{- range  .Attributes}}
					"{{.TfName}}": {
						MarkdownDescription: "{{.Description}}",
						Type:                types.{{.Type}}Type,
						Computed:            true,
					},
					{{- end}}
				}),
				{{- end}}
			},
			{{- end}}
		},
	}, nil
}

func (t dataSource{{camelCase .Name}}Type) NewDataSource(ctx context.Context, in provider.Provider) (datasource.DataSource, diag.Diagnostics) {
	provider, diags := convertProviderType(in)

	return dataSource{{camelCase .Name}}{
		provider: provider,
	}, diags
}

type dataSource{{camelCase .Name}} struct {
	provider iosxeProvider
}

func (d dataSource{{camelCase .Name}}) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config {{camelCase .Name}}

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.getPath()))

	res, err := d.provider.clients[config.Device.Value].GetData(config.getPath())
	if res.StatusCode == 404 {
		config = {{camelCase .Name}}{Device: config.Device}
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
