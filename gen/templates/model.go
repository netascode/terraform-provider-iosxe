//go:build ignore
// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"regexp"
	{{- if hasId .Attributes }}
	"fmt"
	{{- end}}
	{{- $neturl := false }}{{ range .Attributes}}{{ if or (eq .Id true) (eq .Reference true) }}{{ $neturl = true }}{{ end}}{{ end}}
	{{- if $neturl }}
	"net/url"
	{{- end}}
	{{- $strconv := false }}{{ range .Attributes}}{{ if or (and (eq .Type "Int64") (ne .Reference true)) (eq .Type "List") }}{{ $strconv = true }}{{ end}}{{ end}}
	{{- if $strconv }}
	"strconv"
	{{- end}}
	{{- $reflect := false }}{{ range .Attributes}}{{ if eq .Type "List" }}{{ $reflect = true }}{{ end}}{{ end}}
	{{- if $reflect }}
	"reflect"
	{{- end}}

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/terraform-provider-iosxe/internal/provider/helpers"
	{{- $sjson := false }}{{ range .Attributes}}{{ if ne .Reference true}}{{ $sjson = true }}{{ end}}{{ end}}
	{{- if $sjson }}
	"github.com/tidwall/sjson"
	{{- end}}
	{{- $gjson := false }}{{ range .Attributes}}{{ if and (ne .Reference true) (ne .WriteOnly true)}}{{ $gjson = true }}{{ end}}{{ end}}
	{{- if $gjson }}
	"github.com/tidwall/gjson"
	{{- end}}
)

{{- $name := camelCase .Name}}
type {{camelCase .Name}} struct {
	Device types.String `tfsdk:"device"`
	Id     types.String `tfsdk:"id"`
{{- range .Attributes}}
{{- if eq .Type "List"}}
	{{toGoName .TfName}} []{{$name}}{{toGoName .TfName}} `tfsdk:"{{.TfName}}"`
{{- else}}
	{{toGoName .TfName}} types.{{.Type}} `tfsdk:"{{.TfName}}"`
{{- end}}
{{- end}}
}

{{- range .Attributes}}
{{- if eq .Type "List"}}
type {{$name}}{{toGoName .TfName}} struct {
{{- range .Attributes}}
	{{toGoName .TfName}} types.{{.Type}} `tfsdk:"{{.TfName}}"`
{{- end}}
}
{{- end}}
{{- end}}

func (data {{camelCase .Name}}) getPath() string {
{{- if hasId .Attributes}}
	return fmt.Sprintf("{{.Path}}"{{range .Attributes}}{{if or (eq .Id true) (eq .Reference true)}}, url.QueryEscape(fmt.Sprintf("%v", data.{{toGoName .TfName}}.Value)){{end}}{{end}})
{{- else}}
	return "{{.Path}}"
{{- end}}
}

// if last path element has a key -> remove it
func (data {{camelCase .Name}}) getPathShort() string {
	path := data.getPath()
	re := regexp.MustCompile(`(.*)=[^\/]*$`)
	matches := re.FindStringSubmatch(path)
	if len(matches) <= 1 {
		return path
	}
	return matches[1]
}

func (data {{camelCase .Name}}) toBody() string {
	body := `{"` + helpers.LastElement(data.getPath()) + `":{}}`
	{{- range .Attributes}}
	{{- if and (ne .Reference true) (ne .Type "List")}}
	if !data.{{toGoName .TfName}}.Null && !data.{{toGoName .TfName}}.Unknown {
		{{- if eq .Type "Int64"}}
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"{{toJsonPath .YangName .XPath}}", strconv.FormatInt(data.{{toGoName .TfName}}.Value, 10))
		{{- else if and (eq .Type "Bool") (ne .TypeYangBool "boolean")}}
		if data.{{toGoName .TfName}}.Value {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"{{toJsonPath .YangName .XPath}}", map[string]string{})
		}
		{{- else if and (eq .Type "Bool") (eq .TypeYangBool "boolean")}}
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"{{toJsonPath .YangName .XPath}}", data.{{toGoName .TfName}}.Value)
		{{- else if eq .Type "String"}}
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"{{toJsonPath .YangName .XPath}}", data.{{toGoName .TfName}}.Value)
		{{- end}}
	}
	{{- end}}
	{{- end}}
	{{- range .Attributes}}
	{{- if eq .Type "List"}}
	{{- $list := toJsonPath .YangName .XPath }}
	if len(data.{{toGoName .TfName}}) > 0 {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"{{toJsonPath .YangName .XPath}}", []interface{}{})
		for index, item := range data.{{toGoName .TfName}} {
			{{- range .Attributes}}
			if !item.{{toGoName .TfName}}.Null && !item.{{toGoName .TfName}}.Unknown {
				{{- if eq .Type "Int64"}}
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"{{$list}}"+"."+strconv.Itoa(index)+"."+"{{toJsonPath .YangName .XPath}}", strconv.FormatInt(item.{{toGoName .TfName}}.Value, 10))
				{{- else if and (eq .Type "Bool") (ne .TypeYangBool "boolean")}}
				if item.{{toGoName .TfName}}.Value {
					body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"{{$list}}"+"."+strconv.Itoa(index)+"."+"{{toJsonPath .YangName .XPath}}", map[string]string{})
				}
				{{- else if and (eq .Type "Bool") (eq .TypeYangBool "boolean")}}
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"{{$list}}"+"."+strconv.Itoa(index)+"."+"{{toJsonPath .YangName .XPath}}", (data.{{toGoName .TfName}}.Value)
				{{- else if eq .Type "String"}}
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"{{$list}}"+"."+strconv.Itoa(index)+"."+"{{toJsonPath .YangName .XPath}}", item.{{toGoName .TfName}}.Value)
				{{- end}}
			}
			{{- end}}
		}
	}
	{{- end}}
	{{- end}}
	return body
}

func (data *{{camelCase .Name}}) updateFromBody(res gjson.Result) {
	{{- range .Attributes}}
	{{- if and (ne .Reference true) (ne .WriteOnly true)}}
	{{- if eq .Type "Int64"}}
	if value := res.Get(helpers.LastElement(data.getPath())+"."+"{{toJsonPath .YangName .XPath}}"); value.Exists() {
		data.{{toGoName .TfName}}.Value = value.Int()
	} else {
		data.{{toGoName .TfName}}.Null = true
	}
	{{- else if eq .Type "Bool"}}
	if value := res.Get(helpers.LastElement(data.getPath())+"."+"{{toJsonPath .YangName .XPath}}"); value.Exists() {
		{{- if eq .TypeYangBool "boolean"}}
		data.{{toGoName .TfName}}.Value = value.Bool()
		{{- else}}
		data.{{toGoName .TfName}}.Value = true
		{{- end}}
	} else {
		data.{{toGoName .TfName}}.Value = false
	}
	{{- else if eq .Type "String"}}
	if value := res.Get(helpers.LastElement(data.getPath())+"."+"{{toJsonPath .YangName .XPath}}"); value.Exists() {
		data.{{toGoName .TfName}}.Value = value.String()
	} else {
		data.{{toGoName .TfName}}.Null = true
	}
	{{- else if eq .Type "List"}}
	{{- $list := (toGoName .TfName)}}
	{{- $listPath := (toJsonPath .YangName .XPath)}}
	{{- $yangKey := ""}}
	for i := range data.{{$list}}{
		{{- range .Attributes}}
		{{- if eq .Id true}}
		{{- $yangKey = .YangName}}
		key := data.{{$list}}[i].{{toGoName .TfName}}.Value
		{{- end}}
		{{- end}}
		{{- range .Attributes}}
		{{- if ne .WriteOnly true}}
		{{- if eq .Type "Int64"}}
		if value := res.Get(helpers.LastElement(data.getPath())+"."+"{{$listPath}}.#({{$yangKey}}==\""+ key +"\")." + "{{toJsonPath .YangName .XPath}}"); value.Exists() {
			data.{{$list}}[i].{{toGoName .TfName}}.Value = value.Int()
		} else {
			data.{{$list}}[i].{{toGoName .TfName}}.Null = true
		}
		{{- else if eq .Type "Bool"}}
		if value := res.Get(helpers.LastElement(data.getPath())+"."+"{{$listPath}}.#({{$yangKey}}==\""+ key +"\")." + "{{toJsonPath .YangName .XPath}}"); value.Exists() {
			{{- if eq .TypeYangBool "boolean"}}
			data.{{$list}}[i].{{toGoName .TfName}}.Value = value.Bool()
			{{- else}}
			data.{{$list}}[i].{{toGoName .TfName}}.Value = true
			{{- end}}
		} else {
			data.{{$list}}[i].{{toGoName .TfName}}.Value = false
		}
		{{- else if eq .Type "String"}}
		if value := res.Get(helpers.LastElement(data.getPath())+"."+"{{$listPath}}.#({{$yangKey}}==\""+ key +"\")." + "{{toJsonPath .YangName .XPath}}"); value.Exists() {
			data.{{$list}}[i].{{toGoName .TfName}}.Value = value.String()
		} else {
			data.{{$list}}[i].{{toGoName .TfName}}.Null = true
		}
		{{- end}}
		{{- end}}
		{{- end}}
	}
	{{- end}}
	{{- end}}
	{{- end}}
}

func (data *{{camelCase .Name}}) fromBody(res gjson.Result) {
	{{- range .Attributes}}
	{{- if and (ne .Reference true) (ne .Id true) (ne .WriteOnly true)}}
	{{- if eq .Type "Int64"}}
	if value := res.Get(helpers.LastElement(data.getPath())+"."+"{{toJsonPath .YangName .XPath}}"); value.Exists() {
		data.{{toGoName .TfName}}.Value = value.Int()
		data.{{toGoName .TfName}}.Null = false
	}
	{{- else if eq .Type "Bool"}}
	if value := res.Get(helpers.LastElement(data.getPath())+"."+"{{toJsonPath .YangName .XPath}}"); value.Exists() {
		{{- if eq .TypeYangBool "boolean"}}
		data.{{toGoName .TfName}}.Value = value.Bool()
		{{- else}}
		data.{{toGoName .TfName}}.Value = true
		{{- end}}
		data.{{toGoName .TfName}}.Null = false
	} else {
		data.{{toGoName .TfName}}.Value = false
		data.{{toGoName .TfName}}.Null = false
	}
	{{- else if eq .Type "String"}}
	if value := res.Get(helpers.LastElement(data.getPath())+"."+"{{toJsonPath .YangName .XPath}}"); value.Exists() {
		data.{{toGoName .TfName}}.Value = value.String()
		data.{{toGoName .TfName}}.Null = false
	}
	{{- else if eq .Type "List"}}
	if value := res.Get(helpers.LastElement(data.getPath())+"."+"{{toJsonPath .YangName .XPath}}"); value.Exists() {
		data.{{toGoName .TfName}} = make([]{{$name}}{{toGoName .TfName}}, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := {{$name}}{{toGoName .TfName}}{}
			{{- range .Attributes}}
			{{- if ne .WriteOnly true}}
			if cValue := v.Get("{{toJsonPath .YangName .XPath}}"); cValue.Exists() {
				{{- if eq .Type "Int64"}}
				item.{{toGoName .TfName}}.Value = cValue.Int()
				{{- else if and (eq .Type "Bool") (eq .TypeYangBool "boolean")}}
				item.{{toGoName .TfName}}.Value = value.Bool()
				{{- else if and (eq .Type "Bool") (ne .TypeYangBool "boolean")}}
				item.{{toGoName .TfName}}.Value = true
				{{- else if eq .Type "String"}}
				item.{{toGoName .TfName}}.Value = cValue.String()
				{{- end}}
				item.{{toGoName .TfName}}.Null = false
			}
			{{- end}}
			{{- end}}
			data.{{toGoName .TfName}} = append(data.{{toGoName .TfName}}, item)
			return true
		})
	}
	{{- end}}
	{{- end}}
	{{- end}}
}

func (data *{{camelCase .Name}}) setUnknownValues() {
	if data.Device.Unknown {
		data.Device.Unknown = false
		data.Device.Null = true
	}
	if data.Id.Unknown {
		data.Id.Unknown = false
		data.Id.Null = true
	}
	{{- range .Attributes}}
	{{- if ne .Type "List"}}
	if data.{{toGoName .TfName}}.Unknown {
		data.{{toGoName .TfName}}.Unknown = false
		data.{{toGoName .TfName}}.Null = true
	}
	{{- else}}
	{{- $list := (toGoName .TfName)}}
	for i := range data.{{$list}} {
		{{- range .Attributes}}
		if data.{{$list}}[i].{{toGoName .TfName}}.Unknown {
			data.{{$list}}[i].{{toGoName .TfName}}.Unknown = false
			data.{{$list}}[i].{{toGoName .TfName}}.Null = true
		}
		{{- end}}
	}
	{{- end}}
	{{- end}}
}

func (data *{{camelCase .Name}}) getDeletedListItems(state {{camelCase .Name}}) []string {
	deletedListItems := make([]string, 0)
	{{- range .Attributes}}
	{{- if eq .Type "List"}}
	{{- $goKey := ""}}
	{{- range .Attributes}}
	{{- if eq .Id true}}
	{{- $goKey = (toGoName .TfName)}}
	{{- end}}
	{{- end}}
	for _, i := range state.{{toGoName .TfName}} {
		if reflect.ValueOf(i.{{$goKey}}.Value).IsZero() {
			continue
		}
		found := false
		for _, j := range data.{{toGoName .TfName}} {
			if i.{{$goKey}}.Value == j.{{$goKey}}.Value {
				found = true
			}
		}
		if !found {
			deletedListItems = append(deletedListItems, state.getPath()+"/{{.YangName}}="+i.{{$goKey}}.Value)
		}
	}
	{{- end}}
	{{- end}}
	return deletedListItems
}
