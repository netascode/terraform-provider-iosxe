//go:build ignore
{{if .ExcludeTest}}//go:build testAll{{end}}
// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIosxe{{camelCase .Name}}(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: {{if .TestPrerequisites}}testAccIosxe{{camelCase .Name}}PrerequisitesConfig+{{end}}testAccIosxe{{camelCase .Name}}Config_all(),
				Check: resource.ComposeTestCheckFunc(
					{{- $name := .Name }}
					{{- range  .Attributes}}
					{{- if and (ne .Reference true) (ne .WriteOnly true) (ne .ExcludeTest true)}}
					{{- if eq .Type "List"}}
					{{- $list := .TfName }}
					{{- range  .Attributes}}
					{{- if and (ne .WriteOnly true) (ne .ExcludeTest true)}}
					resource.TestCheckResourceAttr("iosxe_{{snakeCase $name}}.test", "{{$list}}.0.{{.TfName}}{{if or (eq .Type "StringList") (eq .Type "Int64List")}}.0{{end}}", "{{.Example}}"),
					{{- end}}
					{{- end}}
					{{- else}}
					resource.TestCheckResourceAttr("iosxe_{{snakeCase $name}}.test", "{{.TfName}}{{if or (eq .Type "StringList") (eq .Type "Int64List")}}.0{{end}}", "{{.Example}}"),
					{{- end}}
					{{- end}}
					{{- end}}
				),
			},
			{
				ResourceName:  "iosxe_{{snakeCase $name}}.test",
				ImportState:   true,
				ImportStateId: "{{getExamplePath .Path .Attributes}}",
			},
		},
	})
}

{{- if .TestPrerequisites}}
const testAccIosxe{{camelCase .Name}}PrerequisitesConfig = `
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

func testAccIosxe{{camelCase .Name}}Config_minimum() string {
	return `
	resource "iosxe_{{snakeCase $name}}" "test" {
	{{- range  .Attributes}}
	{{- if or (eq .Reference true) (eq .Id true) (eq .Mandatory true)}}
	{{- if eq .Type "List"}}
		{{.TfName}} = [{
		{{- range  .Attributes}}
		{{- if ne .ExcludeTest true}}
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
	`
}

func testAccIosxe{{camelCase .Name}}Config_all() string {
	return `
	resource "iosxe_{{snakeCase $name}}" "test" {
	{{- range  .Attributes}}
	{{- if ne .ExcludeTest true}}
	{{- if eq .Type "List"}}
		{{.TfName}} = [{
		{{- range  .Attributes}}
		{{- if ne .ExcludeTest true}}
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
	`
}
