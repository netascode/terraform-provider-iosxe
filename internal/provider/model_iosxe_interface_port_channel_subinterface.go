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

type InterfacePortChannelSubinterface struct {
	Device                   types.String                                      `tfsdk:"device"`
	Id                       types.String                                      `tfsdk:"id"`
	Name                     types.String                                      `tfsdk:"name"`
	Description              types.String                                      `tfsdk:"description"`
	Shutdown                 types.Bool                                        `tfsdk:"shutdown"`
	VrfForwarding            types.String                                      `tfsdk:"vrf_forwarding"`
	Ipv4Address              types.String                                      `tfsdk:"ipv4_address"`
	Ipv4AddressMask          types.String                                      `tfsdk:"ipv4_address_mask"`
	EncapsulationDot1qVlanId types.Int64                                       `tfsdk:"encapsulation_dot1q_vlan_id"`
	IpAccessGroupIn          types.String                                      `tfsdk:"ip_access_group_in"`
	IpAccessGroupInEnable    types.Bool                                        `tfsdk:"ip_access_group_in_enable"`
	IpAccessGroupOut         types.String                                      `tfsdk:"ip_access_group_out"`
	IpAccessGroupOutEnable   types.Bool                                        `tfsdk:"ip_access_group_out_enable"`
	HelperAddresses          []InterfacePortChannelSubinterfaceHelperAddresses `tfsdk:"helper_addresses"`
}
type InterfacePortChannelSubinterfaceHelperAddresses struct {
	Address types.String `tfsdk:"address"`
	Global  types.Bool   `tfsdk:"global"`
	Vrf     types.String `tfsdk:"vrf"`
}

func (data InterfacePortChannelSubinterface) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/interface/Port-channel-subinterface/Port-channel=%v", url.QueryEscape(fmt.Sprintf("%v", data.Name.ValueString())))
}

// if last path element has a key -> remove it
func (data InterfacePortChannelSubinterface) getPathShort() string {
	path := data.getPath()
	re := regexp.MustCompile(`(.*)=[^\/]*$`)
	matches := re.FindStringSubmatch(path)
	if len(matches) <= 1 {
		return path
	}
	return matches[1]
}

func (data InterfacePortChannelSubinterface) toBody(ctx context.Context) string {
	body := `{"` + helpers.LastElement(data.getPath()) + `":{}}`
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"name", data.Name.ValueString())
	}
	if !data.Description.IsNull() && !data.Description.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"description", data.Description.ValueString())
	}
	if !data.Shutdown.IsNull() && !data.Shutdown.IsUnknown() {
		if data.Shutdown.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"shutdown", map[string]string{})
		}
	}
	if !data.VrfForwarding.IsNull() && !data.VrfForwarding.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"vrf.forwarding", data.VrfForwarding.ValueString())
	}
	if !data.Ipv4Address.IsNull() && !data.Ipv4Address.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ip.address.primary.address", data.Ipv4Address.ValueString())
	}
	if !data.Ipv4AddressMask.IsNull() && !data.Ipv4AddressMask.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ip.address.primary.mask", data.Ipv4AddressMask.ValueString())
	}
	if !data.EncapsulationDot1qVlanId.IsNull() && !data.EncapsulationDot1qVlanId.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"encapsulation.dot1Q.vlan-id", strconv.FormatInt(data.EncapsulationDot1qVlanId.ValueInt64(), 10))
	}
	if !data.IpAccessGroupIn.IsNull() && !data.IpAccessGroupIn.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ip.access-group.in.acl.acl-name", data.IpAccessGroupIn.ValueString())
	}
	if !data.IpAccessGroupInEnable.IsNull() && !data.IpAccessGroupInEnable.IsUnknown() {
		if data.IpAccessGroupInEnable.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ip.access-group.in.acl.in", map[string]string{})
		}
	}
	if !data.IpAccessGroupOut.IsNull() && !data.IpAccessGroupOut.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ip.access-group.out.acl.acl-name", data.IpAccessGroupOut.ValueString())
	}
	if !data.IpAccessGroupOutEnable.IsNull() && !data.IpAccessGroupOutEnable.IsUnknown() {
		if data.IpAccessGroupOutEnable.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ip.access-group.out.acl.out", map[string]string{})
		}
	}
	if len(data.HelperAddresses) > 0 {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ip.helper-address", []interface{}{})
		for index, item := range data.HelperAddresses {
			if !item.Address.IsNull() && !item.Address.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ip.helper-address"+"."+strconv.Itoa(index)+"."+"address", item.Address.ValueString())
			}
			if !item.Global.IsNull() && !item.Global.IsUnknown() {
				if item.Global.ValueBool() {
					body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ip.helper-address"+"."+strconv.Itoa(index)+"."+"global", map[string]string{})
				}
			}
			if !item.Vrf.IsNull() && !item.Vrf.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ip.helper-address"+"."+strconv.Itoa(index)+"."+"vrf", item.Vrf.ValueString())
			}
		}
	}
	return body
}

func (data *InterfacePortChannelSubinterface) updateFromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get(prefix + "description"); value.Exists() && !data.Description.IsNull() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get(prefix + "shutdown"); !data.Shutdown.IsNull() {
		if value.Exists() {
			data.Shutdown = types.BoolValue(true)
		} else {
			data.Shutdown = types.BoolValue(false)
		}
	} else {
		data.Shutdown = types.BoolNull()
	}
	if value := res.Get(prefix + "vrf.forwarding"); value.Exists() && !data.VrfForwarding.IsNull() {
		data.VrfForwarding = types.StringValue(value.String())
	} else {
		data.VrfForwarding = types.StringNull()
	}
	if value := res.Get(prefix + "ip.address.primary.address"); value.Exists() && !data.Ipv4Address.IsNull() {
		data.Ipv4Address = types.StringValue(value.String())
	} else {
		data.Ipv4Address = types.StringNull()
	}
	if value := res.Get(prefix + "ip.address.primary.mask"); value.Exists() && !data.Ipv4AddressMask.IsNull() {
		data.Ipv4AddressMask = types.StringValue(value.String())
	} else {
		data.Ipv4AddressMask = types.StringNull()
	}
	if value := res.Get(prefix + "encapsulation.dot1Q.vlan-id"); value.Exists() && !data.EncapsulationDot1qVlanId.IsNull() {
		data.EncapsulationDot1qVlanId = types.Int64Value(value.Int())
	} else {
		data.EncapsulationDot1qVlanId = types.Int64Null()
	}
	if value := res.Get(prefix + "ip.access-group.in.acl.acl-name"); value.Exists() && !data.IpAccessGroupIn.IsNull() {
		data.IpAccessGroupIn = types.StringValue(value.String())
	} else {
		data.IpAccessGroupIn = types.StringNull()
	}
	if value := res.Get(prefix + "ip.access-group.in.acl.in"); !data.IpAccessGroupInEnable.IsNull() {
		if value.Exists() {
			data.IpAccessGroupInEnable = types.BoolValue(true)
		} else {
			data.IpAccessGroupInEnable = types.BoolValue(false)
		}
	} else {
		data.IpAccessGroupInEnable = types.BoolNull()
	}
	if value := res.Get(prefix + "ip.access-group.out.acl.acl-name"); value.Exists() && !data.IpAccessGroupOut.IsNull() {
		data.IpAccessGroupOut = types.StringValue(value.String())
	} else {
		data.IpAccessGroupOut = types.StringNull()
	}
	if value := res.Get(prefix + "ip.access-group.out.acl.out"); !data.IpAccessGroupOutEnable.IsNull() {
		if value.Exists() {
			data.IpAccessGroupOutEnable = types.BoolValue(true)
		} else {
			data.IpAccessGroupOutEnable = types.BoolValue(false)
		}
	} else {
		data.IpAccessGroupOutEnable = types.BoolNull()
	}
	for i := range data.HelperAddresses {
		keys := [...]string{"address"}
		keyValues := [...]string{data.HelperAddresses[i].Address.ValueString()}

		var r gjson.Result
		res.Get(prefix + "ip.helper-address").ForEach(
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
		if value := r.Get("address"); value.Exists() && !data.HelperAddresses[i].Address.IsNull() {
			data.HelperAddresses[i].Address = types.StringValue(value.String())
		} else {
			data.HelperAddresses[i].Address = types.StringNull()
		}
		if value := r.Get("global"); !data.HelperAddresses[i].Global.IsNull() {
			if value.Exists() {
				data.HelperAddresses[i].Global = types.BoolValue(true)
			} else {
				data.HelperAddresses[i].Global = types.BoolValue(false)
			}
		} else {
			data.HelperAddresses[i].Global = types.BoolNull()
		}
		if value := r.Get("vrf"); value.Exists() && !data.HelperAddresses[i].Vrf.IsNull() {
			data.HelperAddresses[i].Vrf = types.StringValue(value.String())
		} else {
			data.HelperAddresses[i].Vrf = types.StringNull()
		}
	}
}

func (data *InterfacePortChannelSubinterface) fromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "description"); value.Exists() {
		data.Description = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "shutdown"); value.Exists() {
		data.Shutdown = types.BoolValue(true)
	} else {
		data.Shutdown = types.BoolValue(false)
	}
	if value := res.Get(prefix + "vrf.forwarding"); value.Exists() {
		data.VrfForwarding = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "ip.address.primary.address"); value.Exists() {
		data.Ipv4Address = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "ip.address.primary.mask"); value.Exists() {
		data.Ipv4AddressMask = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "encapsulation.dot1Q.vlan-id"); value.Exists() {
		data.EncapsulationDot1qVlanId = types.Int64Value(value.Int())
	}
	if value := res.Get(prefix + "ip.access-group.in.acl.acl-name"); value.Exists() {
		data.IpAccessGroupIn = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "ip.access-group.in.acl.in"); value.Exists() {
		data.IpAccessGroupInEnable = types.BoolValue(true)
	} else {
		data.IpAccessGroupInEnable = types.BoolValue(false)
	}
	if value := res.Get(prefix + "ip.access-group.out.acl.acl-name"); value.Exists() {
		data.IpAccessGroupOut = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "ip.access-group.out.acl.out"); value.Exists() {
		data.IpAccessGroupOutEnable = types.BoolValue(true)
	} else {
		data.IpAccessGroupOutEnable = types.BoolValue(false)
	}
	if value := res.Get(prefix + "ip.helper-address"); value.Exists() {
		data.HelperAddresses = make([]InterfacePortChannelSubinterfaceHelperAddresses, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := InterfacePortChannelSubinterfaceHelperAddresses{}
			if cValue := v.Get("address"); cValue.Exists() {
				item.Address = types.StringValue(cValue.String())
			}
			if cValue := v.Get("global"); cValue.Exists() {
				item.Global = types.BoolValue(true)
			} else {
				item.Global = types.BoolValue(false)
			}
			if cValue := v.Get("vrf"); cValue.Exists() {
				item.Vrf = types.StringValue(cValue.String())
			}
			data.HelperAddresses = append(data.HelperAddresses, item)
			return true
		})
	}
}

func (data *InterfacePortChannelSubinterface) getDeletedListItems(ctx context.Context, state InterfacePortChannelSubinterface) []string {
	deletedListItems := make([]string, 0)
	for i := range state.HelperAddresses {
		stateKeyValues := [...]string{state.HelperAddresses[i].Address.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.HelperAddresses[i].Address.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.HelperAddresses {
			found = true
			if state.HelperAddresses[i].Address.ValueString() != data.HelperAddresses[j].Address.ValueString() {
				found = false
			}
			if found {
				break
			}
		}
		if !found {
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/ip/helper-address=%v", state.getPath(), strings.Join(stateKeyValues[:], ",")))
		}
	}
	return deletedListItems
}

func (data *InterfacePortChannelSubinterface) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	if !data.Shutdown.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/shutdown", data.getPath()))
	}
	if !data.IpAccessGroupInEnable.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/ip/access-group/in/apply-type/apply-intf/acl/in", data.getPath()))
	}
	if !data.IpAccessGroupOutEnable.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/ip/access-group/out/apply-type/apply-intf/acl/out", data.getPath()))
	}

	for i := range data.HelperAddresses {
		keyValues := [...]string{data.HelperAddresses[i].Address.ValueString()}
		if !data.HelperAddresses[i].Global.ValueBool() {
			emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/ip/helper-address=%v/helper-choice/global/global", data.getPath(), strings.Join(keyValues[:], ",")))
		}
	}
	return emptyLeafsDelete
}
