package provider

import (
	"context"
	"fmt"
	"regexp"

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

func (data *Restconf) fromBody(ctx context.Context, res gjson.Result) {
	for attr := range data.Attributes.Elems {
		value := res.Get(helpers.LastElement(data.getPath()) + "." + attr)
		if !value.Exists() ||
			(value.IsObject() && len(value.Map()) == 0) ||
			value.Raw == "[null]" {

			data.Attributes.Elems[attr] = types.String{Value: ""}
		} else {
			data.Attributes.Elems[attr] = types.String{Value: value.String()}
		}
	}

	for i := range data.Lists {
		for ii := range data.Lists[i].Items {
			for attr := range data.Lists[i].Items[ii].Attributes.Elems {
				key := data.Lists[i].Key.Value
				v, _ := data.Lists[i].Items[ii].Attributes.Elems[attr].ToTerraformValue(ctx)
				var keyValue string
				v.As(&keyValue)
				jsonPath := fmt.Sprintf(`%s.%s.#(%s=="%s").%s`, helpers.LastElement(data.getPath()), data.Lists[i].Name.Value, key, keyValue, attr)
				value := res.Get(jsonPath)
				if !value.Exists() ||
					(value.IsObject() && len(value.Map()) == 0) ||
					value.Raw == "[null]" {

					data.Lists[i].Items[ii].Attributes.Elems[attr] = types.String{Value: ""}
				} else {
					data.Lists[i].Items[ii].Attributes.Elems[attr] = types.String{Value: value.String()}
				}
			}
		}
	}
}

func (data *Restconf) getDeletedListItems(ctx context.Context, state Restconf) []string {
	deletedListItems := make([]string, 0)
	for l := range state.Lists {
		name := state.Lists[l].Name.Value
		key := state.Lists[l].Key.Value
		var dataList RestconfList
		for _, dl := range data.Lists {
			if dl.Name.Value == name {
				dataList = dl
			}
		}
		// check if state item is also included in plan, if not delete item
		for i := range state.Lists[l].Items {
			var slia map[string]string
			state.Lists[l].Items[i].Attributes.ElementsAs(ctx, &slia, false)
			if slia[key] == "" {
				continue
			}
			found := false
			for dli := range dataList.Items {
				var dlia map[string]string
				dataList.Items[dli].Attributes.ElementsAs(ctx, &dlia, false)
				if dlia[key] == slia[key] {
					found = true
					break
				}
			}
			if !found {
				deletedListItems = append(deletedListItems, state.getPath()+"/"+name+"="+slia[key])
			}
		}
	}
	return deletedListItems
}
