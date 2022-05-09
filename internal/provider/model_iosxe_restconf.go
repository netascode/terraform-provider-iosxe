package provider

import (
	"context"
	"regexp"
	"sort"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-restconf"
	"github.com/netascode/terraform-provider-iosxe/internal/provider/helpers"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type Restconf struct {
	Device     types.String   `tfsdk:"device"`
	Id         types.String   `tfsdk:"id"`
	Path       types.String   `tfsdk:"path"`
	Delete     types.Bool     `tfsdk:"delete"`
	Attributes types.Map      `tfsdk:"attributes"`
	Lists      []RestconfList `tfsdk:"lists"`
}

type RestconfList struct {
	Name  types.String       `tfsdk:"name"`
	Key   types.String       `tfsdk:"key"`
	Items []RestconfListItem `tfsdk:"items"`
}

type RestconfListItem struct {
	Attributes types.Map `tfsdk:"attributes"`
}

type RestconfDataSource struct {
	Device     types.String `tfsdk:"device"`
	Id         types.String `tfsdk:"id"`
	Path       types.String `tfsdk:"path"`
	Attributes types.Map    `tfsdk:"attributes"`
}

func (data Restconf) getPath() string {
	return data.Path.Value
}

// if last path element has a key -> remove it
func (data Restconf) getPathShort() string {
	path := data.Path.Value
	re := regexp.MustCompile(`(.*)=[^\/]*$`)
	matches := re.FindStringSubmatch(path)
	if len(matches) <= 1 {
		return path
	}
	return matches[1]
}

func (data *Restconf) fromBody(ctx context.Context, res gjson.Result) {
	var ea map[string]string
	data.Attributes.ElementsAs(ctx, &ea, false)
	existingAttr := make([]string, len(ea))
	for k := range ea {
		existingAttr = append(existingAttr, k)
	}

	attributes := make(map[string]attr.Value)

	for attr, value := range res.Get(helpers.LastElement(data.Path.Value)).Map() {
		if helpers.Contains(existingAttr, attr) {
			// handle empty maps
			if value.IsObject() && len(value.Map()) == 0 {
				attributes[attr] = types.String{Value: ""}
			} else if value.Raw == "[null]" {
				attributes[attr] = types.String{Value: ""}
			} else {
				attributes[attr] = types.String{Value: value.String()}
			}
		}
	}
	data.Attributes.Elems = attributes

	lists := make([]RestconfList, 0)
	for i := range data.Lists {
		list := RestconfList{Name: data.Lists[i].Name, Key: data.Lists[i].Key}
		if value := res.Get(helpers.LastElement(data.Path.Value) + "." + data.Lists[i].Name.Value); value.Exists() {
			value.ForEach(func(k, v gjson.Result) bool {
				ca := make(map[string]attr.Value)
				for attr, value := range v.Map() {
					if value.IsObject() && len(value.Map()) == 0 {
						ca[attr] = types.String{Value: ""}
					} else if value.Raw == "[null]" {
						ca[attr] = types.String{Value: ""}
					} else if !value.IsArray() && !value.IsObject() {
						ca[attr] = types.String{Value: value.String()}
					}
				}
				item := RestconfListItem{Attributes: types.Map{ElemType: types.StringType, Elems: ca}}
				list.Items = append(list.Items, item)
				return true
			})
		}
		// sort slice by existing order in data
		sort.SliceStable(list.Items, func(iii, jjj int) bool {
			key := data.Lists[i].Key.Value
			for ii := range data.Lists[i].Items {
				var dataAttrs map[string]string
				data.Lists[i].Items[ii].Attributes.ElementsAs(ctx, &dataAttrs, false)
				var newAttrs map[string]string
				list.Items[iii].Attributes.ElementsAs(ctx, &newAttrs, false)
				if dataAttrs[key] == newAttrs[key] {
					return true
				}
				list.Items[jjj].Attributes.ElementsAs(ctx, &newAttrs, false)
				if dataAttrs[key] == newAttrs[key] {
					return false
				}
			}
			return false
		})
		lists = append(lists, list)
	}

	if len(lists) > 0 {
		data.Lists = lists
	} else {
		data.Lists = nil
	}
}

func (data Restconf) toBody(ctx context.Context) string {
	body := `{"` + helpers.LastElement(data.Path.Value) + `":{}}`

	var attributes map[string]string
	data.Attributes.ElementsAs(ctx, &attributes, false)

	for attr, value := range attributes {
		body, _ = sjson.Set(body, helpers.LastElement(data.Path.Value)+"."+attr, value)
	}
	for i := range data.Lists {
		body, _ = sjson.Set(body, helpers.LastElement(data.Path.Value)+"."+data.Lists[i].Name.Value, []interface{}{})
		for ii := range data.Lists[i].Items {
			var listAttributes map[string]string
			data.Lists[i].Items[ii].Attributes.ElementsAs(ctx, &listAttributes, false)
			attrs := restconf.Body{}
			for attr, value := range listAttributes {
				attrs = attrs.Set(attr, value)
			}
			body, _ = sjson.SetRaw(body, helpers.LastElement(data.Path.Value)+"."+data.Lists[i].Name.Value+".-1", attrs.Str)
		}
	}

	return body
}
