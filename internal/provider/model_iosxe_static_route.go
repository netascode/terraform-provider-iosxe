// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/terraform-provider-iosxe/internal/provider/helpers"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type StaticRoute struct {
	Device   types.String          `tfsdk:"device"`
	Id       types.String          `tfsdk:"id"`
	Prefix   types.String          `tfsdk:"prefix"`
	Mask     types.String          `tfsdk:"mask"`
	NextHops []StaticRouteNextHops `tfsdk:"next_hops"`
}
type StaticRouteNextHops struct {
	NextHop   types.String `tfsdk:"next_hop"`
	Metric    types.Int64  `tfsdk:"metric"`
	Global    types.Bool   `tfsdk:"global"`
	Name      types.String `tfsdk:"name"`
	Permanent types.Bool   `tfsdk:"permanent"`
	Tag       types.Int64  `tfsdk:"tag"`
}

func (data StaticRoute) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/ip/route/ip-route-interface-forwarding-list=%s,%s", url.QueryEscape(fmt.Sprintf("%v", data.Prefix.ValueString())), url.QueryEscape(fmt.Sprintf("%v", data.Mask.ValueString())))
}

// if last path element has a key -> remove it
func (data StaticRoute) getPathShort() string {
	path := data.getPath()
	re := regexp.MustCompile(`(.*)=[^\/]*$`)
	matches := re.FindStringSubmatch(path)
	if len(matches) <= 1 {
		return path
	}
	return matches[1]
}

func (data StaticRoute) toBody(ctx context.Context) string {
	body := `{"` + helpers.LastElement(data.getPath()) + `":{}}`
	if !data.Prefix.IsNull() && !data.Prefix.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"prefix", data.Prefix.ValueString())
	}
	if !data.Mask.IsNull() && !data.Mask.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"mask", data.Mask.ValueString())
	}
	if len(data.NextHops) > 0 {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"fwd-list", []interface{}{})
		for index, item := range data.NextHops {
			if !item.NextHop.IsNull() && !item.NextHop.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"fwd-list"+"."+strconv.Itoa(index)+"."+"fwd", item.NextHop.ValueString())
			}
			if !item.Metric.IsNull() && !item.Metric.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"fwd-list"+"."+strconv.Itoa(index)+"."+"metric", strconv.FormatInt(item.Metric.ValueInt64(), 10))
			}
			if !item.Global.IsNull() && !item.Global.IsUnknown() {
				if item.Global.ValueBool() {
					body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"fwd-list"+"."+strconv.Itoa(index)+"."+"global", map[string]string{})
				}
			}
			if !item.Name.IsNull() && !item.Name.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"fwd-list"+"."+strconv.Itoa(index)+"."+"name", item.Name.ValueString())
			}
			if !item.Permanent.IsNull() && !item.Permanent.IsUnknown() {
				if item.Permanent.ValueBool() {
					body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"fwd-list"+"."+strconv.Itoa(index)+"."+"permanent", map[string]string{})
				}
			}
			if !item.Tag.IsNull() && !item.Tag.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"fwd-list"+"."+strconv.Itoa(index)+"."+"tag", strconv.FormatInt(item.Tag.ValueInt64(), 10))
			}
		}
	}
	return body
}

func (data *StaticRoute) updateFromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "prefix"); value.Exists() && !data.Prefix.IsNull() {
		data.Prefix = types.StringValue(value.String())
	} else {
		data.Prefix = types.StringNull()
	}
	if value := res.Get(prefix + "mask"); value.Exists() && !data.Mask.IsNull() {
		data.Mask = types.StringValue(value.String())
	} else {
		data.Mask = types.StringNull()
	}
	for i := range data.NextHops {
		keys := [...]string{"fwd"}
		keyValues := [...]string{data.NextHops[i].NextHop.ValueString()}

		var r gjson.Result
		res.Get(prefix + "fwd-list").ForEach(
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
		if value := r.Get("fwd"); value.Exists() && !data.NextHops[i].NextHop.IsNull() {
			data.NextHops[i].NextHop = types.StringValue(value.String())
		} else {
			data.NextHops[i].NextHop = types.StringNull()
		}
		if value := r.Get("metric"); value.Exists() && !data.NextHops[i].Metric.IsNull() {
			data.NextHops[i].Metric = types.Int64Value(value.Int())
		} else {
			data.NextHops[i].Metric = types.Int64Null()
		}
		if value := r.Get("global"); !data.NextHops[i].Global.IsNull() {
			if value.Exists() {
				data.NextHops[i].Global = types.BoolValue(true)
			} else {
				data.NextHops[i].Global = types.BoolValue(false)
			}
		} else {
			data.NextHops[i].Global = types.BoolNull()
		}
		if value := r.Get("name"); value.Exists() && !data.NextHops[i].Name.IsNull() {
			data.NextHops[i].Name = types.StringValue(value.String())
		} else {
			data.NextHops[i].Name = types.StringNull()
		}
		if value := r.Get("permanent"); !data.NextHops[i].Permanent.IsNull() {
			if value.Exists() {
				data.NextHops[i].Permanent = types.BoolValue(true)
			} else {
				data.NextHops[i].Permanent = types.BoolValue(false)
			}
		} else {
			data.NextHops[i].Permanent = types.BoolNull()
		}
		if value := r.Get("tag"); value.Exists() && !data.NextHops[i].Tag.IsNull() {
			data.NextHops[i].Tag = types.Int64Value(value.Int())
		} else {
			data.NextHops[i].Tag = types.Int64Null()
		}
	}
}

func (data *StaticRoute) fromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "fwd-list"); value.Exists() {
		data.NextHops = make([]StaticRouteNextHops, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := StaticRouteNextHops{}
			if cValue := v.Get("fwd"); cValue.Exists() {
				item.NextHop = types.StringValue(cValue.String())
			}
			if cValue := v.Get("metric"); cValue.Exists() {
				item.Metric = types.Int64Value(cValue.Int())
			}
			if cValue := v.Get("global"); cValue.Exists() {
				item.Global = types.BoolValue(true)
			} else {
				item.Global = types.BoolValue(false)
			}
			if cValue := v.Get("name"); cValue.Exists() {
				item.Name = types.StringValue(cValue.String())
			}
			if cValue := v.Get("permanent"); cValue.Exists() {
				item.Permanent = types.BoolValue(true)
			} else {
				item.Permanent = types.BoolValue(false)
			}
			if cValue := v.Get("tag"); cValue.Exists() {
				item.Tag = types.Int64Value(cValue.Int())
			}
			data.NextHops = append(data.NextHops, item)
			return true
		})
	}
}

func (data *StaticRoute) getDeletedListItems(ctx context.Context, state StaticRoute) []string {
	deletedListItems := make([]string, 0)
	for i := range state.NextHops {
		stateKeyValues := [...]string{state.NextHops[i].NextHop.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.NextHops[i].NextHop.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.NextHops {
			found = true
			if state.NextHops[i].NextHop.ValueString() != data.NextHops[j].NextHop.ValueString() {
				found = false
			}
			if found {
				break
			}
		}
		if !found {
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/fwd-list=%v", state.getPath(), strings.Join(stateKeyValues[:], ",")))
		}
	}
	return deletedListItems
}

func (data *StaticRoute) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)

	for i := range data.NextHops {
		keyValues := [...]string{data.NextHops[i].NextHop.ValueString()}
		if !data.NextHops[i].Global.IsNull() && !data.NextHops[i].Global.ValueBool() {
			emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/fwd-list=%v/global", data.getPath(), strings.Join(keyValues[:], ",")))
		}
		if !data.NextHops[i].Permanent.IsNull() && !data.NextHops[i].Permanent.ValueBool() {
			emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/fwd-list=%v/permanent", data.getPath(), strings.Join(keyValues[:], ",")))
		}
	}
	return emptyLeafsDelete
}
