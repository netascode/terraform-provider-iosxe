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

type InterfaceOSPFProcess struct {
	Device    types.String               `tfsdk:"device"`
	Id        types.String               `tfsdk:"id"`
	Type      types.String               `tfsdk:"type"`
	Name      types.String               `tfsdk:"name"`
	ProcessId types.Int64                `tfsdk:"process_id"`
	Area      []InterfaceOSPFProcessArea `tfsdk:"area"`
}
type InterfaceOSPFProcessArea struct {
	AreaId types.String `tfsdk:"area_id"`
}

func (data InterfaceOSPFProcess) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/interface/%s=%v/ip/Cisco-IOS-XE-ospf:router-ospf/ospf/process-id=%v", url.QueryEscape(fmt.Sprintf("%v", data.Type.ValueString())), url.QueryEscape(fmt.Sprintf("%v", data.Name.ValueString())), url.QueryEscape(fmt.Sprintf("%v", data.ProcessId.ValueInt64())))
}

// if last path element has a key -> remove it
func (data InterfaceOSPFProcess) getPathShort() string {
	path := data.getPath()
	re := regexp.MustCompile(`(.*)=[^\/]*$`)
	matches := re.FindStringSubmatch(path)
	if len(matches) <= 1 {
		return path
	}
	return matches[1]
}

func (data InterfaceOSPFProcess) toBody(ctx context.Context) string {
	body := `{"` + helpers.LastElement(data.getPath()) + `":{}}`
	if !data.ProcessId.IsNull() && !data.ProcessId.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"id", strconv.FormatInt(data.ProcessId.ValueInt64(), 10))
	}
	if len(data.Area) > 0 {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"area", []interface{}{})
		for index, item := range data.Area {
			if !item.AreaId.IsNull() && !item.AreaId.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"area"+"."+strconv.Itoa(index)+"."+"area-id", item.AreaId.ValueString())
			}
		}
	}
	return body
}

func (data *InterfaceOSPFProcess) updateFromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "id"); value.Exists() {
		data.ProcessId = types.Int64Value(value.Int())
	} else {
		data.ProcessId = types.Int64Null()
	}
	for i := range data.Area {
		keys := [...]string{"area-id"}
		keyValues := [...]string{data.Area[i].AreaId.ValueString()}

		var r gjson.Result
		res.Get(prefix + "area").ForEach(
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
		if value := r.Get("area-id"); value.Exists() {
			data.Area[i].AreaId = types.StringValue(value.String())
		} else {
			data.Area[i].AreaId = types.StringNull()
		}
	}
}

func (data *InterfaceOSPFProcess) fromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "area"); value.Exists() {
		data.Area = make([]InterfaceOSPFProcessArea, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := InterfaceOSPFProcessArea{}
			if cValue := v.Get("area-id"); cValue.Exists() {
				item.AreaId = types.StringValue(cValue.String())
			}
			data.Area = append(data.Area, item)
			return true
		})
	}
}

func (data *InterfaceOSPFProcess) setUnknownValues(ctx context.Context) {
	if data.Device.IsUnknown() {
		data.Device = types.StringNull()
	}
	if data.Id.IsUnknown() {
		data.Id = types.StringNull()
	}
	if data.Type.IsUnknown() {
		data.Type = types.StringNull()
	}
	if data.Name.IsUnknown() {
		data.Name = types.StringNull()
	}
	if data.ProcessId.IsUnknown() {
		data.ProcessId = types.Int64Null()
	}
	for i := range data.Area {
		if data.Area[i].AreaId.IsUnknown() {
			data.Area[i].AreaId = types.StringNull()
		}
	}
}

func (data *InterfaceOSPFProcess) getDeletedListItems(ctx context.Context, state InterfaceOSPFProcess) []string {
	deletedListItems := make([]string, 0)
	for i := range state.Area {
		stateKeyValues := [...]string{state.Area[i].AreaId.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.Area[i].AreaId.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.Area {
			found = true
			if state.Area[i].AreaId.ValueString() != data.Area[j].AreaId.ValueString() {
				found = false
			}
			if found {
				break
			}
		}
		if !found {
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/area=%v", state.getPath(), strings.Join(stateKeyValues[:], ",")))
		}
	}
	return deletedListItems
}

func (data *InterfaceOSPFProcess) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)

	return emptyLeafsDelete
}
