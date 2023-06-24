//go:build ignore
{{if .ExcludeTest}}//go:build testAll{{end}}
// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceIosxe{{camelCase .Name}}(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: {{if .TestPrerequisites}}testAccDataSourceIosxe{{camelCase .Name}}PrerequisitesConfig+{{end}}testAccDataSourceIosxe{{camelCase .Name}}Config,
				Check: resource.ComposeTestCheckFunc(
					{{- $name := .Name }}
					{{- range  .Attributes}}
					{{- if and (not .Id) (not .Reference) (not .WriteOnly) (not .ExcludeTest)}}
					{{- if eq .Type "List"}}
					{{- $list := .TfName }}
					{{- range  .Attributes}}
					{{- if and (not .WriteOnly) (not .ExcludeTest)}}
					resource.TestCheckResourceAttr("data.iosxe_{{snakeCase $name}}.test", "{{$list}}.0.{{.TfName}}{{if or (eq .Type "StringList") (eq .Type "Int64List")}}.0{{end}}", "{{.Example}}"),
					{{- end}}
					{{- end}}
					{{- else}}
					resource.TestCheckResourceAttr("data.iosxe_{{snakeCase $name}}.test", "{{.TfName}}{{if or (eq .Type "StringList") (eq .Type "Int64List")}}.0{{end}}", "{{.Example}}"),
					{{- end}}
					{{- end}}
					{{- end}}
				),
			},
		},
	})
}

{{- if .TestPrerequisites}}
const testAccDataSourceIosxe{{camelCase .Name}}PrerequisitesConfig = `
{{- range $index, $item := .TestPrerequisites}}
resource "iosxe_restconf" "PreReq{{$index}}" {
  path = "{{.Path}}"
  {{- if .NoDelete}}
  delete = false
  {{- end}}
  attributes = {
    {{- range  .Attributes}}
      "{{.Name}}" = {{if .Reference}}{{.Reference}}{{else}}"{{.Value}}"{{end}}
    {{- end}}
  }
  {{- if .Lists}}
  lists = [
  {{- range .Lists}}
    {
      name = "{{.Name}}"
      key = "{{.Key}}"
      items = [
        {{- range .Items}}
          {
            attributes = {
            {{- range .Attributes}}
              "{{.Name}}" = {{if .Reference}}{{.Reference}}{{else}}"{{.Value}}"{{end}}
            {{- end}}
            }
          },
        {{- end}}
      ] 
    },
  {{- end}}
  ]
  {{- end}}
  {{- if .Dependencies}}
  depends_on = [{{range .Dependencies}}iosxe_restconf.PreReq{{.}}, {{end}}]
  {{- end}}
}
{{ end}}
`
{{- end}}

const testAccDataSourceIosxe{{camelCase .Name}}Config = `

resource "iosxe_{{snakeCase $name}}" "test" {
{{- range  .Attributes}}
{{- if not .ExcludeTest}}
{{- if eq .Type "List"}}
  {{.TfName}} = [{
    {{- range  .Attributes}}
    {{- if not .ExcludeTest}}
    {{.TfName}} = {{if eq .Type "String"}}"{{.Example}}"{{else if eq .Type "StringList"}}["{{.Example}}"]{{else if eq .Type "Int64List"}}[{{.Example}}]{{else}}{{.Example}}{{end}}
    {{- end}}
    {{- end}}
  }]
{{- else}}
  {{.TfName}} = {{if eq .Type "String"}}"{{.Example}}"{{else if eq .Type "StringList"}}["{{.Example}}"]{{else if eq .Type "Int64List"}}[{{.Example}}]{{else}}{{.Example}}{{end}}
{{- end}}
{{- end}}
{{- end}}
{{- if .TestPrerequisites}}
  depends_on = [{{range $index, $item := .TestPrerequisites}}iosxe_restconf.PreReq{{$index}}, {{end}}]
{{- end}}
}

data "iosxe_{{snakeCase .Name}}" "test" {
{{- range  .Attributes}}
{{- if or .Id .Reference}}
  {{.TfName}} = {{if eq .Type "String"}}"{{.Example}}"{{else if eq .Type "StringList"}}["{{.Example}}"]{{else if eq .Type "Int64List"}}[{{.Example}}]{{else}}{{.Example}}{{end}}
{{- end}}
{{- end}}
  depends_on = [iosxe_{{snakeCase $name}}.test]
}
`
