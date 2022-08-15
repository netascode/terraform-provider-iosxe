package provider

import (
	"context"
	"regexp"
	"strings"

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
	Name   types.String       `tfsdk:"name"`
	Key    types.String       `tfsdk:"key"`
	Items  []RestconfListItem `tfsdk:"items"`
	Values types.List         `tfsdk:"values"`
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
		attr = strings.ReplaceAll(attr, "/", ".")
		body, _ = sjson.Set(body, helpers.LastElement(data.Path.Value)+"."+attr, value)
	}
	for i := range data.Lists {
		listName := strings.ReplaceAll(data.Lists[i].Name.Value, "/", ".")
		if len(data.Lists[i].Items) > 0 {
			body, _ = sjson.Set(body, helpers.LastElement(data.Path.Value)+"."+listName, []interface{}{})
			for ii := range data.Lists[i].Items {
				var listAttributes map[string]string
				data.Lists[i].Items[ii].Attributes.ElementsAs(ctx, &listAttributes, false)
				attrs := restconf.Body{}
				for attr, value := range listAttributes {
					attrs = attrs.Set(attr, value)
				}
				body, _ = sjson.SetRaw(body, helpers.LastElement(data.Path.Value)+"."+listName+".-1", attrs.Str)
			}
		} else if len(data.Lists[i].Values.Elems) > 0 {
			var values []string
			data.Lists[i].Values.ElementsAs(ctx, &values, false)
			body, _ = sjson.Set(body, helpers.LastElement(data.Path.Value)+"."+listName, values)
		}
	}

	return body
}

func (data *Restconf) fromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	for attr := range data.Attributes.Elems {
		attrPath := strings.ReplaceAll(attr, "/", ".")
		value := res.Get(prefix + attrPath)
		if !value.Exists() ||
			(value.IsObject() && len(value.Map()) == 0) ||
			value.Raw == "[null]" {

			data.Attributes.Elems[attr] = types.String{Value: ""}
		} else {
			data.Attributes.Elems[attr] = types.String{Value: value.String()}
		}
	}

	for i := range data.Lists {
		keys := strings.Split(data.Lists[i].Key.Value, ",")
		namePath := strings.ReplaceAll(data.Lists[i].Name.Value, "/", ".")
		if len(data.Lists[i].Items) > 0 {
			for ii := range data.Lists[i].Items {
				var keyValues []string
				for _, key := range keys {
					v, _ := data.Lists[i].Items[ii].Attributes.Elems[key].ToTerraformValue(ctx)
					var keyValue string
					v.As(&keyValue)
					keyValues = append(keyValues, keyValue)
				}

				// find item by key(s)
				var r gjson.Result
				res.Get(prefix + namePath).ForEach(
					func(_, v gjson.Result) bool {
						found := false
						for ik := range keys {
							if v.Get(keys[ik]).String() == keyValues[ik] {
								found = true
								continue
							}
							found = false
							break
						}
						if found {
							r = v
							return false
						}
						return true
					},
				)

				for attr := range data.Lists[i].Items[ii].Attributes.Elems {
					attrPath := strings.ReplaceAll(attr, "/", ".")
					value := r.Get(attrPath)
					if !value.Exists() ||
						(value.IsObject() && len(value.Map()) == 0) ||
						value.Raw == "[null]" {

						data.Lists[i].Items[ii].Attributes.Elems[attr] = types.String{Value: ""}
					} else {
						data.Lists[i].Items[ii].Attributes.Elems[attr] = types.String{Value: value.String()}
					}
				}
			}
		} else if len(data.Lists[i].Values.Elems) > 0 {
			values := res.Get(prefix + namePath)
			if values.IsArray() {
				data.Lists[i].Values.Elems = helpers.GetValueSlice(values.Array())
			}
		}
	}
}

func (data *Restconf) getDeletedListItems(ctx context.Context, state Restconf) []string {
	deletedListItems := make([]string, 0)
	for l := range state.Lists {
		name := state.Lists[l].Name.Value
		namePath := strings.ReplaceAll(name, "/", ".")
		keys := strings.Split(state.Lists[l].Key.Value, ",")
		var dataList RestconfList
		for _, dl := range data.Lists {
			if dl.Name.Value == name {
				dataList = dl
			}
		}
		if len(state.Lists[l].Items) > 0 {
			// check if state item is also included in plan, if not delete item
			for i := range state.Lists[l].Items {
				var slia map[string]string
				state.Lists[l].Items[i].Attributes.ElementsAs(ctx, &slia, false)

				// if state key values are empty move on to next item
				emptyKey := false
				for _, key := range keys {
					if slia[key] == "" {
						emptyKey = true
						break
					}
				}
				if emptyKey {
					continue
				}

				// find data (plan) item with matching key values
				found := false
				for dli := range dataList.Items {
					var dlia map[string]string
					dataList.Items[dli].Attributes.ElementsAs(ctx, &dlia, false)
					for _, key := range keys {
						if dlia[key] == slia[key] {
							found = true
							continue
						}
						found = false
						break
					}
					if found {
						break
					}
				}

				// if no matching item in plan found -> delete
				if !found {
					keyValues := make([]string, len(keys))
					for k, key := range keys {
						keyValues[k] = slia[key]
					}
					deletedListItems = append(deletedListItems, state.getPath()+"/"+namePath+"="+strings.Join(keyValues, ","))
				}
			}
		} else if len(state.Lists[l].Values.Elems) > 0 {
			var slv []string
			state.Lists[l].Values.ElementsAs(ctx, &slv, false)
			// check if state value is also included in plan, if not delete value from list
			for _, stateValue := range slv {
				found := false
				var dlv []string
				dataList.Values.ElementsAs(ctx, &dlv, false)
				for _, dataValue := range dlv {
					if stateValue == dataValue {
						found = true
						break
					}
				}
				if !found {
					deletedListItems = append(deletedListItems, state.getPath()+"/"+namePath+"="+stateValue)
				}
			}
		}
	}
	return deletedListItems
}
