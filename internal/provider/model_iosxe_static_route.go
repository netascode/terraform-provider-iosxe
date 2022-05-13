// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"
	"net/url"
	"reflect"
	"regexp"
	"strconv"

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
	return fmt.Sprintf("Cisco-IOS-XE-native:native/ip/route/ip-route-interface-forwarding-list=%s,%s", url.QueryEscape(fmt.Sprintf("%v", data.Prefix.Value)), url.QueryEscape(fmt.Sprintf("%v", data.Mask.Value)))
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

func (data StaticRoute) toBody() string {
	body := `{"` + helpers.LastElement(data.getPath()) + `":{}}`
	if !data.Prefix.Null && !data.Prefix.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"prefix", data.Prefix.Value)
	}
	if !data.Mask.Null && !data.Mask.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"mask", data.Mask.Value)
	}
	if len(data.NextHops) > 0 {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"fwd-list", []interface{}{})
		for index, item := range data.NextHops {
			if !item.NextHop.Null && !item.NextHop.Unknown {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"fwd-list"+"."+strconv.Itoa(index)+"."+"fwd", item.NextHop.Value)
			}
			if !item.Metric.Null && !item.Metric.Unknown {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"fwd-list"+"."+strconv.Itoa(index)+"."+"metric", strconv.FormatInt(item.Metric.Value, 10))
			}
			if !item.Global.Null && !item.Global.Unknown {
				if item.Global.Value {
					body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"fwd-list"+"."+strconv.Itoa(index)+"."+"global", map[string]string{})
				}
			}
			if !item.Name.Null && !item.Name.Unknown {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"fwd-list"+"."+strconv.Itoa(index)+"."+"name", item.Name.Value)
			}
			if !item.Permanent.Null && !item.Permanent.Unknown {
				if item.Permanent.Value {
					body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"fwd-list"+"."+strconv.Itoa(index)+"."+"permanent", map[string]string{})
				}
			}
			if !item.Tag.Null && !item.Tag.Unknown {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"fwd-list"+"."+strconv.Itoa(index)+"."+"tag", strconv.FormatInt(item.Tag.Value, 10))
			}
		}
	}
	return body
}

func (data *StaticRoute) updateFromBody(res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "prefix"); value.Exists() {
		data.Prefix.Value = value.String()
	} else {
		data.Prefix.Null = true
	}
	if value := res.Get(prefix + "mask"); value.Exists() {
		data.Mask.Value = value.String()
	} else {
		data.Mask.Null = true
	}
	for i := range data.NextHops {
		key := data.NextHops[i].NextHop.Value
		if value := res.Get(fmt.Sprintf("%vfwd-list.#(fwd==\"%v\").fwd", prefix, key)); value.Exists() {
			data.NextHops[i].NextHop.Value = value.String()
		} else {
			data.NextHops[i].NextHop.Null = true
		}
		if value := res.Get(fmt.Sprintf("%vfwd-list.#(fwd==\"%v\").metric", prefix, key)); value.Exists() {
			data.NextHops[i].Metric.Value = value.Int()
		} else {
			data.NextHops[i].Metric.Null = true
		}
		if value := res.Get(fmt.Sprintf("%vfwd-list.#(fwd==\"%v\").global", prefix, key)); value.Exists() {
			data.NextHops[i].Global.Value = true
		} else {
			data.NextHops[i].Global.Value = false
		}
		if value := res.Get(fmt.Sprintf("%vfwd-list.#(fwd==\"%v\").name", prefix, key)); value.Exists() {
			data.NextHops[i].Name.Value = value.String()
		} else {
			data.NextHops[i].Name.Null = true
		}
		if value := res.Get(fmt.Sprintf("%vfwd-list.#(fwd==\"%v\").permanent", prefix, key)); value.Exists() {
			data.NextHops[i].Permanent.Value = true
		} else {
			data.NextHops[i].Permanent.Value = false
		}
		if value := res.Get(fmt.Sprintf("%vfwd-list.#(fwd==\"%v\").tag", prefix, key)); value.Exists() {
			data.NextHops[i].Tag.Value = value.Int()
		} else {
			data.NextHops[i].Tag.Null = true
		}
	}
}

func (data *StaticRoute) fromBody(res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "fwd-list"); value.Exists() {
		data.NextHops = make([]StaticRouteNextHops, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := StaticRouteNextHops{}
			if cValue := v.Get("fwd"); cValue.Exists() {
				item.NextHop.Value = cValue.String()
				item.NextHop.Null = false
			}
			if cValue := v.Get("metric"); cValue.Exists() {
				item.Metric.Value = cValue.Int()
				item.Metric.Null = false
			}
			if cValue := v.Get("global"); cValue.Exists() {
				item.Global.Value = true
				item.Global.Null = false
			}
			if cValue := v.Get("name"); cValue.Exists() {
				item.Name.Value = cValue.String()
				item.Name.Null = false
			}
			if cValue := v.Get("permanent"); cValue.Exists() {
				item.Permanent.Value = true
				item.Permanent.Null = false
			}
			if cValue := v.Get("tag"); cValue.Exists() {
				item.Tag.Value = cValue.Int()
				item.Tag.Null = false
			}
			data.NextHops = append(data.NextHops, item)
			return true
		})
	}
}

func (data *StaticRoute) setUnknownValues() {
	if data.Device.Unknown {
		data.Device.Unknown = false
		data.Device.Null = true
	}
	if data.Id.Unknown {
		data.Id.Unknown = false
		data.Id.Null = true
	}
	if data.Prefix.Unknown {
		data.Prefix.Unknown = false
		data.Prefix.Null = true
	}
	if data.Mask.Unknown {
		data.Mask.Unknown = false
		data.Mask.Null = true
	}
	for i := range data.NextHops {
		if data.NextHops[i].NextHop.Unknown {
			data.NextHops[i].NextHop.Unknown = false
			data.NextHops[i].NextHop.Null = true
		}
		if data.NextHops[i].Metric.Unknown {
			data.NextHops[i].Metric.Unknown = false
			data.NextHops[i].Metric.Null = true
		}
		if data.NextHops[i].Global.Unknown {
			data.NextHops[i].Global.Unknown = false
			data.NextHops[i].Global.Null = true
		}
		if data.NextHops[i].Name.Unknown {
			data.NextHops[i].Name.Unknown = false
			data.NextHops[i].Name.Null = true
		}
		if data.NextHops[i].Permanent.Unknown {
			data.NextHops[i].Permanent.Unknown = false
			data.NextHops[i].Permanent.Null = true
		}
		if data.NextHops[i].Tag.Unknown {
			data.NextHops[i].Tag.Unknown = false
			data.NextHops[i].Tag.Null = true
		}
	}
}

func (data *StaticRoute) getDeletedListItems(state StaticRoute) []string {
	deletedListItems := make([]string, 0)
	for _, i := range state.NextHops {
		if reflect.ValueOf(i.NextHop.Value).IsZero() {
			continue
		}
		found := false
		for _, j := range data.NextHops {
			if i.NextHop.Value == j.NextHop.Value {
				found = true
			}
		}
		if !found {
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/fwd-list=%v", state.getPath(), i.NextHop.Value))
		}
	}
	return deletedListItems
}

func (data *StaticRoute) getEmptyLeafsDelete() []string {
	emptyLeafsDelete := make([]string, 0)
	return emptyLeafsDelete
}
