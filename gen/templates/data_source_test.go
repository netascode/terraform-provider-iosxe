//go:build ignore
// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceIosxe{{camelCase .Name}}(t *testing.T) {
	{{- if len .TestTags}}
	if {{range $i, $e := .TestTags}}{{if $i}} && {{end}}os.Getenv("{{$e}}") == ""{{end}} {
        t.Skip("skipping test, set environment variable {{range $i, $e := .TestTags}}{{if $i}} or {{end}}{{$e}}{{end}}")
    }
	{{- end}}
	var checks []resource.TestCheckFunc
	{{- $name := .Name }}
	{{- range .Attributes}}
	{{- if and (not .Id) (not .Reference) (not .WriteOnly) (not .ExcludeTest)}}
	{{- if eq .Type "List"}}
	{{- $list := .TfName }}
	{{- range .Attributes}}
	{{- if and (not .WriteOnly) (not .ExcludeTest)}}
	{{- if eq .Type "List"}}
	{{- $clist := .TfName }}
	{{- range .Attributes}}
	{{- if and (not .WriteOnly) (not .ExcludeTest)}}
	{{- if len .TestTags}}
	if {{range $i, $e := .TestTags}}{{if $i}} || {{end}}os.Getenv("{{$e}}") != ""{{end}} {
		checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_{{snakeCase $name}}.test", "{{$list}}.0.{{$clist}}.0.{{.TfName}}{{if or (eq .Type "StringList") (eq .Type "Int64List")}}.0{{end}}", "{{.Example}}"))
	}
	{{- else}}
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_{{snakeCase $name}}.test", "{{$list}}.0.{{$clist}}.0.{{.TfName}}{{if or (eq .Type "StringList") (eq .Type "Int64List")}}.0{{end}}", "{{.Example}}"))
	{{- end}}
	{{- end}}
	{{- end}}
	{{- else}}
	{{- if len .TestTags}}
	if {{range $i, $e := .TestTags}}{{if $i}} || {{end}}os.Getenv("{{$e}}") != ""{{end}} {
		checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_{{snakeCase $name}}.test", "{{$list}}.0.{{.TfName}}{{if or (eq .Type "StringList") (eq .Type "Int64List")}}.0{{end}}", "{{.Example}}"))
	}
	{{- else}}
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_{{snakeCase $name}}.test", "{{$list}}.0.{{.TfName}}{{if or (eq .Type "StringList") (eq .Type "Int64List")}}.0{{end}}", "{{.Example}}"))
	{{- end}}
	{{- end}}
	{{- end}}
	{{- end}}
	{{- else}}
	{{- if len .TestTags}}
	if {{range $i, $e := .TestTags}}{{if $i}} || {{end}}os.Getenv("{{$e}}") != ""{{end}} {
		checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_{{snakeCase $name}}.test", "{{.TfName}}{{if or (eq .Type "StringList") (eq .Type "Int64List")}}.0{{end}}", "{{.Example}}"))
	}
	{{- else}}
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_{{snakeCase $name}}.test", "{{.TfName}}{{if or (eq .Type "StringList") (eq .Type "Int64List")}}.0{{end}}", "{{.Example}}"))
	{{- end}}
	{{- end}}
	{{- end}}
	{{- end}}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: {{if .TestPrerequisites}}testAccDataSourceIosxe{{camelCase .Name}}PrerequisitesConfig+{{end}}testAccDataSourceIosxe{{camelCase .Name}}Config(),
				Check: resource.ComposeTestCheckFunc(checks...),
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
					{{- range .Attributes}}
					"{{.Name}}" = {{if .Reference}}{{.Reference}}{{else}}"{{.Value}}"{{end}}
					{{- end}}
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

func testAccDataSourceIosxe{{camelCase .Name}}Config() string {
	config := `resource "iosxe_{{snakeCase $name}}" "test" {` + "\n"
	{{- if and (not .NoDelete) (not .NoDeleteAttributes) (not .DefaultDeleteAttributes)}}
	config += `	delete_mode = "attributes"\n`
	{{- end}}
	{{- range  .Attributes}}
	{{- if not .ExcludeTest}}
	{{- if eq .Type "List"}}
	config += `	{{.TfName}} = [{` + "\n"
		{{- range  .Attributes}}
		{{- if not .ExcludeTest}}
		{{- if eq .Type "List"}}
	config += `		{{.TfName}} = [{` + "\n"
			{{- range  .Attributes}}
			{{- if not .ExcludeTest}}
			{{- if len .TestTags}}
	if {{range $i, $e := .TestTags}}{{if $i}} || {{end}}os.Getenv("{{$e}}") != ""{{end}} {
		config += `			{{.TfName}} = {{if eq .Type "String"}}"{{.Example}}"{{else if eq .Type "StringList"}}["{{.Example}}"]{{else if eq .Type "Int64List"}}[{{.Example}}]{{else}}{{.Example}}{{end}}` + "\n"
	}
			{{- else}}
	config += `			{{.TfName}} = {{if eq .Type "String"}}"{{.Example}}"{{else if eq .Type "StringList"}}["{{.Example}}"]{{else if eq .Type "Int64List"}}[{{.Example}}]{{else}}{{.Example}}{{end}}` + "\n"
			{{- end}}
			{{- end}}
			{{- end}}
	config += `		}]` + "\n"
		{{- else}}
		{{- if len .TestTags}}
	if {{range $i, $e := .TestTags}}{{if $i}} || {{end}}os.Getenv("{{$e}}") != ""{{end}} {
		config += `		{{.TfName}} = {{if eq .Type "String"}}"{{.Example}}"{{else if eq .Type "StringList"}}["{{.Example}}"]{{else if eq .Type "Int64List"}}[{{.Example}}]{{else}}{{.Example}}{{end}}` + "\n"
	}
		{{- else}}
	config += `		{{.TfName}} = {{if eq .Type "String"}}"{{.Example}}"{{else if eq .Type "StringList"}}["{{.Example}}"]{{else if eq .Type "Int64List"}}[{{.Example}}]{{else}}{{.Example}}{{end}}` + "\n"
		{{- end}}
		{{- end}}
		{{- end}}
		{{- end}}
		config += `	}]` + "\n"
	{{- else}}
	{{- if len .TestTags}}
	if {{range $i, $e := .TestTags}}{{if $i}} || {{end}}os.Getenv("{{$e}}") != ""{{end}} {
		config += `	{{.TfName}} = {{if eq .Type "String"}}"{{.Example}}"{{else if eq .Type "StringList"}}["{{.Example}}"]{{else if eq .Type "Int64List"}}[{{.Example}}]{{else}}{{.Example}}{{end}}` + "\n"
	}
	{{- else}}
	config += `	{{.TfName}} = {{if eq .Type "String"}}"{{.Example}}"{{else if eq .Type "StringList"}}["{{.Example}}"]{{else if eq .Type "Int64List"}}[{{.Example}}]{{else}}{{.Example}}{{end}}` + "\n"
	{{- end}}
	{{- end}}
	{{- end}}
	{{- end}}
	{{- if .TestPrerequisites}}
	config += `	depends_on = [{{range $index, $item := .TestPrerequisites}}iosxe_restconf.PreReq{{$index}}, {{end}}]` + "\n"
	{{- end}}
	config += `}` + "\n"
	
	config += `
		data "iosxe_{{snakeCase .Name}}" "test" {
			{{- range .Attributes}}
			{{- if or .Id .Reference}}
			{{.TfName}} = {{if eq .Type "String"}}"{{.Example}}"{{else if eq .Type "StringList"}}["{{.Example}}"]{{else if eq .Type "Int64List"}}[{{.Example}}]{{else}}{{.Example}}{{end}}
			{{- end}}
			{{- end}}
			depends_on = [iosxe_{{snakeCase $name}}.test]
		}
	`
	return config
}
