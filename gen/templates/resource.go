//go:build ignore
// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-restconf"
	"github.com/netascode/terraform-provider-iosxe/internal/provider/helpers"
)

type resource{{camelCase .Name}}Type struct{}

func (t resource{{camelCase .Name}}Type) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "{{.ResDescription}}",

		Attributes: map[string]tfsdk.Attribute{
			"device": {
				MarkdownDescription: "A device name from the provider configuration.",
				Type:                types.StringType,
				Optional:            true,
			},
			"id": {
				MarkdownDescription: "The path of the object.",
				Type:                types.StringType,
				Computed:            true,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					tfsdk.UseStateForUnknown(),
				},
			},
			{{- range  .Attributes}}
			"{{.TfName}}": {
				MarkdownDescription: helpers.NewAttributeDescription("{{.Description}}")
					{{- if len .EnumValues -}}
					.AddStringEnumDescription({{range .EnumValues}}"{{.}}", {{end}})
					{{- end -}}
					{{- if or (ne .MinInt 0) (ne .MaxInt 0) -}}
					.AddIntegerRangeDescription({{.MinInt}}, {{.MaxInt}})
					{{- end -}}
					{{- if len .DefaultValue -}}
					.AddDefaultValueDescription("{{.DefaultValue}}")
					{{- end -}}
					.String,
				{{- if ne .Type "List"}}
				Type:                types.{{.Type}}Type,
				{{- end}}
				{{- if or (eq .Id true) (eq .Reference true) (eq .Mandatory true)}}
				Required:            true,
				{{- else}}
				Optional:            true,
				Computed:            true,
				{{- end}}
				{{- if len .EnumValues}}
				Validators: []tfsdk.AttributeValidator{
					helpers.StringEnumValidator({{range .EnumValues}}"{{.}}", {{end}}),
				},
				{{- else if len .StringPatterns}}
				Validators: []tfsdk.AttributeValidator{
					helpers.StringPatternValidator({{.StringMinLength}}, {{.StringMaxLength}}, {{range .StringPatterns}}`{{.}}`, {{end}}),
				},
				{{- else if or (ne .MinInt 0) (ne .MaxInt 0)}}
				Validators: []tfsdk.AttributeValidator{
					helpers.IntegerRangeValidator({{.MinInt}}, {{.MaxInt}}),
				},
				{{- end}}
				{{- if or (len .DefaultValue) (eq .Id true) (eq .Reference true)}}
				PlanModifiers: tfsdk.AttributePlanModifiers{
					{{- if or (eq .Id true) (eq .Reference true)}}
					tfsdk.RequiresReplace(),
					{{- else if eq .Type "Int64"}}
					helpers.IntegerDefaultModifier({{.DefaultValue}}),
					{{- else if eq .Type "Bool"}}
					helpers.BooleanDefaultModifier({{.DefaultValue}}),
					{{- else if eq .Type "String"}}
					helpers.StringDefaultModifier("{{.DefaultValue}}"),
					{{- end}}
				},
				{{- end}}
				{{- if eq .Type "List"}}
				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
					{{- range  .Attributes}}
					"{{.TfName}}": {
						MarkdownDescription: helpers.NewAttributeDescription("{{.Description}}")
							{{- if len .EnumValues -}}
							.AddStringEnumDescription({{range .EnumValues}}"{{.}}", {{end}})
							{{- end -}}
							{{- if or (ne .MinInt 0) (ne .MaxInt 0) -}}
							.AddIntegerRangeDescription({{.MinInt}}, {{.MaxInt}})
							{{- end -}}
							{{- if len .DefaultValue -}}
							.AddDefaultValueDescription("{{.DefaultValue}}")
							{{- end -}}
							.String,
						Type:                types.{{.Type}}Type,
						{{- if eq .Mandatory true}}
						Required:            true,
						{{- else}}
						Optional:            true,
						Computed:            true,
						{{- end}}
						{{- if len .EnumValues}}
						Validators: []tfsdk.AttributeValidator{
							helpers.StringEnumValidator({{range .EnumValues}}"{{.}}", {{end}}),
						},
						{{- else if len .StringPatterns}}
						Validators: []tfsdk.AttributeValidator{
							helpers.StringPatternValidator({{.StringMinLength}}, {{.StringMaxLength}}, {{range .StringPatterns}}`{{.}}`, {{end}}),
						},
						{{- else if or (ne .MinInt 0) (ne .MaxInt 0)}}
						Validators: []tfsdk.AttributeValidator{
							helpers.IntegerRangeValidator({{.MinInt}}, {{.MaxInt}}),
						},
						{{- end}}
						{{- if or (len .DefaultValue) (eq .Id true) (eq .Reference true)}}
						PlanModifiers: tfsdk.AttributePlanModifiers{
							{{- if or (eq .Id true) (eq .Reference true)}}
							tfsdk.RequiresReplace(),
							{{- else if eq .Type "Int64"}}
							helpers.IntegerDefaultModifier({{.DefaultValue}}),
							{{- else if eq .Type "Bool"}}
							helpers.BooleanDefaultModifier({{.DefaultValue}}),
							{{- else if eq .Type "String"}}
							helpers.StringDefaultModifier("{{.DefaultValue}}"),
							{{- end}}
						},
						{{- end}}
					},
					{{- end}}
				}, tfsdk.ListNestedAttributesOptions{}),
				{{- end}}
			},
			{{- end}}
		},
	}, nil
}

func (t resource{{camelCase .Name}}Type) NewResource(ctx context.Context, in tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	provider, diags := convertProviderType(in)

	return resource{{camelCase .Name}}{
		provider: provider,
	}, diags
}

type resource{{camelCase .Name}} struct {
	provider provider
}

func (r resource{{camelCase .Name}}) Create(ctx context.Context, req tfsdk.CreateResourceRequest, resp *tfsdk.CreateResourceResponse) {
	var plan, state {{camelCase .Name}}

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.getPath()))

	// Create object
	body := plan.toBody()

	res, _ := r.provider.clients[plan.Device.Value].GetData(plan.getPath(), restconf.Query("depth", "1"))
	if res.StatusCode < 200 || res.StatusCode > 299 {
		_, err := r.provider.clients[plan.Device.Value].PutData(plan.getPath(), body)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s", err))
			return
		}
	} else {
		_, err := r.provider.clients[plan.Device.Value].PatchData(plan.getPath(), body)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PATCH), got error: %s", err))
			return
		}
	}

	// Read object
	res, err := r.provider.clients[plan.Device.Value].GetData(plan.getPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	state.fromBody(res.Res)
	state.fromPlan(plan)
	state.Id.Value = plan.getPath()

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.getPath()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r resource{{camelCase .Name}}) Read(ctx context.Context, req tfsdk.ReadResourceRequest, resp *tfsdk.ReadResourceResponse) {
	var state, newState {{camelCase .Name}}

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.Value))

	res, err := r.provider.clients[state.Device.Value].GetData(state.Id.Value)
	if res.StatusCode != 404 {
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
			return
		}

		newState.fromBody(res.Res)
	}
	newState.fromPlan(state)
	newState.Id = state.Id

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.Value))

	diags = resp.State.Set(ctx, &newState)
	resp.Diagnostics.Append(diags...)
}

func (r resource{{camelCase .Name}}) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {
	var plan, state {{camelCase .Name}}

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.Id.Value))

	// Update object
	body := plan.toBody()
	res, _ := r.provider.clients[plan.Device.Value].GetData(plan.getPath(), restconf.Query("depth", "1"))
	if res.StatusCode < 200 || res.StatusCode > 299 {
		_, err := r.provider.clients[plan.Device.Value].PutData(plan.getPath(), body)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object (PUT), got error: %s", err))
			return
		}
	} else {
		_, err := r.provider.clients[plan.Device.Value].PatchData(plan.getPath(), body)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object (PATCH), got error: %s", err))
			return
		}
	}

	// Read object
	res, err := r.provider.clients[plan.Device.Value].GetData(plan.getPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	state.fromBody(res.Res)
	state.fromPlan(plan)
	state.Id.Value = plan.getPath()

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.Value))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r resource{{camelCase .Name}}) Delete(ctx context.Context, req tfsdk.DeleteResourceRequest, resp *tfsdk.DeleteResourceResponse) {
	var state {{camelCase .Name}}

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.Value))
{{ if not .NoDelete}}
	_, err := r.provider.clients[state.Device.Value].DeleteData(state.getPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object, got error: %s", err))
		return
	}
{{ end}}
	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.Value))

	resp.State.RemoveResource(ctx)
}

func (r resource{{camelCase .Name}}) ImportState(ctx context.Context, req tfsdk.ImportResourceStateRequest, resp *tfsdk.ImportResourceStateResponse) {
	tfsdk.ResourceImportStatePassthroughID(ctx, tftypes.NewAttributePath().WithAttributeName("id"), req, resp)
}
