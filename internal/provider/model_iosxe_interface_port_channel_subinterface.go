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
	HelperAddresses          []InterfacePortChannelSubinterfaceHelperAddresses `tfsdk:"helper_addresses"`
}
type InterfacePortChannelSubinterfaceHelperAddresses struct {
	Address types.String `tfsdk:"address"`
	Global  types.Bool   `tfsdk:"global"`
	Vrf     types.String `tfsdk:"vrf"`
}

func (data InterfacePortChannelSubinterface) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/interface/Port-channel-subinterface/Port-channel=%v", url.QueryEscape(fmt.Sprintf("%v", data.Name.Value)))
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

func (data InterfacePortChannelSubinterface) toBody() string {
	body := `{"` + helpers.LastElement(data.getPath()) + `":{}}`
	if !data.Name.Null && !data.Name.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"name", data.Name.Value)
	}
	if !data.Description.Null && !data.Description.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"description", data.Description.Value)
	}
	if !data.Shutdown.Null && !data.Shutdown.Unknown {
		if data.Shutdown.Value {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"shutdown", map[string]string{})
		}
	}
	if !data.VrfForwarding.Null && !data.VrfForwarding.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"vrf.forwarding", data.VrfForwarding.Value)
	}
	if !data.Ipv4Address.Null && !data.Ipv4Address.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ip.address.primary.address", data.Ipv4Address.Value)
	}
	if !data.Ipv4AddressMask.Null && !data.Ipv4AddressMask.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ip.address.primary.mask", data.Ipv4AddressMask.Value)
	}
	if !data.EncapsulationDot1qVlanId.Null && !data.EncapsulationDot1qVlanId.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"encapsulation.dot1Q.vlan-id", strconv.FormatInt(data.EncapsulationDot1qVlanId.Value, 10))
	}
	if len(data.HelperAddresses) > 0 {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ip.helper-address", []interface{}{})
		for index, item := range data.HelperAddresses {
			if !item.Address.Null && !item.Address.Unknown {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ip.helper-address"+"."+strconv.Itoa(index)+"."+"address", item.Address.Value)
			}
			if !item.Global.Null && !item.Global.Unknown {
				if item.Global.Value {
					body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ip.helper-address"+"."+strconv.Itoa(index)+"."+"global", map[string]string{})
				}
			}
			if !item.Vrf.Null && !item.Vrf.Unknown {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ip.helper-address"+"."+strconv.Itoa(index)+"."+"vrf", item.Vrf.Value)
			}
		}
	}
	return body
}

func (data *InterfacePortChannelSubinterface) updateFromBody(res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "name"); value.Exists() {
		data.Name.Value = value.String()
	} else {
		data.Name.Null = true
	}
	if value := res.Get(prefix + "description"); value.Exists() {
		data.Description.Value = value.String()
	} else {
		data.Description.Null = true
	}
	if value := res.Get(prefix + "shutdown"); value.Exists() {
		data.Shutdown.Value = true
	} else {
		data.Shutdown.Value = false
	}
	if value := res.Get(prefix + "vrf.forwarding"); value.Exists() {
		data.VrfForwarding.Value = value.String()
	} else {
		data.VrfForwarding.Null = true
	}
	if value := res.Get(prefix + "ip.address.primary.address"); value.Exists() {
		data.Ipv4Address.Value = value.String()
	} else {
		data.Ipv4Address.Null = true
	}
	if value := res.Get(prefix + "ip.address.primary.mask"); value.Exists() {
		data.Ipv4AddressMask.Value = value.String()
	} else {
		data.Ipv4AddressMask.Null = true
	}
	if value := res.Get(prefix + "encapsulation.dot1Q.vlan-id"); value.Exists() {
		data.EncapsulationDot1qVlanId.Value = value.Int()
	} else {
		data.EncapsulationDot1qVlanId.Null = true
	}
	for i := range data.HelperAddresses {
		key := data.HelperAddresses[i].Address.Value
		if value := res.Get(fmt.Sprintf("%vip.helper-address.#(address==\"%v\").address", prefix, key)); value.Exists() {
			data.HelperAddresses[i].Address.Value = value.String()
		} else {
			data.HelperAddresses[i].Address.Null = true
		}
		if value := res.Get(fmt.Sprintf("%vip.helper-address.#(address==\"%v\").global", prefix, key)); value.Exists() {
			data.HelperAddresses[i].Global.Value = true
		} else {
			data.HelperAddresses[i].Global.Value = false
		}
		if value := res.Get(fmt.Sprintf("%vip.helper-address.#(address==\"%v\").vrf", prefix, key)); value.Exists() {
			data.HelperAddresses[i].Vrf.Value = value.String()
		} else {
			data.HelperAddresses[i].Vrf.Null = true
		}
	}
}

func (data *InterfacePortChannelSubinterface) fromBody(res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "description"); value.Exists() {
		data.Description.Value = value.String()
		data.Description.Null = false
	}
	if value := res.Get(prefix + "shutdown"); value.Exists() {
		data.Shutdown.Value = true
		data.Shutdown.Null = false
	} else {
		data.Shutdown.Value = false
		data.Shutdown.Null = false
	}
	if value := res.Get(prefix + "vrf.forwarding"); value.Exists() {
		data.VrfForwarding.Value = value.String()
		data.VrfForwarding.Null = false
	}
	if value := res.Get(prefix + "ip.address.primary.address"); value.Exists() {
		data.Ipv4Address.Value = value.String()
		data.Ipv4Address.Null = false
	}
	if value := res.Get(prefix + "ip.address.primary.mask"); value.Exists() {
		data.Ipv4AddressMask.Value = value.String()
		data.Ipv4AddressMask.Null = false
	}
	if value := res.Get(prefix + "encapsulation.dot1Q.vlan-id"); value.Exists() {
		data.EncapsulationDot1qVlanId.Value = value.Int()
		data.EncapsulationDot1qVlanId.Null = false
	}
	if value := res.Get(prefix + "ip.helper-address"); value.Exists() {
		data.HelperAddresses = make([]InterfacePortChannelSubinterfaceHelperAddresses, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := InterfacePortChannelSubinterfaceHelperAddresses{}
			if cValue := v.Get("address"); cValue.Exists() {
				item.Address.Value = cValue.String()
				item.Address.Null = false
			}
			if cValue := v.Get("global"); cValue.Exists() {
				item.Global.Value = true
				item.Global.Null = false
			}
			if cValue := v.Get("vrf"); cValue.Exists() {
				item.Vrf.Value = cValue.String()
				item.Vrf.Null = false
			}
			data.HelperAddresses = append(data.HelperAddresses, item)
			return true
		})
	}
}

func (data *InterfacePortChannelSubinterface) setUnknownValues() {
	if data.Device.Unknown {
		data.Device.Unknown = false
		data.Device.Null = true
	}
	if data.Id.Unknown {
		data.Id.Unknown = false
		data.Id.Null = true
	}
	if data.Name.Unknown {
		data.Name.Unknown = false
		data.Name.Null = true
	}
	if data.Description.Unknown {
		data.Description.Unknown = false
		data.Description.Null = true
	}
	if data.Shutdown.Unknown {
		data.Shutdown.Unknown = false
		data.Shutdown.Null = true
	}
	if data.VrfForwarding.Unknown {
		data.VrfForwarding.Unknown = false
		data.VrfForwarding.Null = true
	}
	if data.Ipv4Address.Unknown {
		data.Ipv4Address.Unknown = false
		data.Ipv4Address.Null = true
	}
	if data.Ipv4AddressMask.Unknown {
		data.Ipv4AddressMask.Unknown = false
		data.Ipv4AddressMask.Null = true
	}
	if data.EncapsulationDot1qVlanId.Unknown {
		data.EncapsulationDot1qVlanId.Unknown = false
		data.EncapsulationDot1qVlanId.Null = true
	}
	for i := range data.HelperAddresses {
		if data.HelperAddresses[i].Address.Unknown {
			data.HelperAddresses[i].Address.Unknown = false
			data.HelperAddresses[i].Address.Null = true
		}
		if data.HelperAddresses[i].Global.Unknown {
			data.HelperAddresses[i].Global.Unknown = false
			data.HelperAddresses[i].Global.Null = true
		}
		if data.HelperAddresses[i].Vrf.Unknown {
			data.HelperAddresses[i].Vrf.Unknown = false
			data.HelperAddresses[i].Vrf.Null = true
		}
	}
}

func (data *InterfacePortChannelSubinterface) getDeletedListItems(state InterfacePortChannelSubinterface) []string {
	deletedListItems := make([]string, 0)
	for _, i := range state.HelperAddresses {
		if reflect.ValueOf(i.Address.Value).IsZero() {
			continue
		}
		found := false
		for _, j := range data.HelperAddresses {
			if i.Address.Value == j.Address.Value {
				found = true
			}
		}
		if !found {
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/ip/helper-address=%v", state.getPath(), i.Address.Value))
		}
	}
	return deletedListItems
}
