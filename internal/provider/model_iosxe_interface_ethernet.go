// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

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

type InterfaceEthernet struct {
	Device                     types.String                       `tfsdk:"device"`
	Id                         types.String                       `tfsdk:"id"`
	Type                       types.String                       `tfsdk:"type"`
	Name                       types.String                       `tfsdk:"name"`
	MediaType                  types.String                       `tfsdk:"media_type"`
	Switchport                 types.Bool                         `tfsdk:"switchport"`
	Description                types.String                       `tfsdk:"description"`
	Shutdown                   types.Bool                         `tfsdk:"shutdown"`
	IpProxyArp                 types.Bool                         `tfsdk:"ip_proxy_arp"`
	IpRedirects                types.Bool                         `tfsdk:"ip_redirects"`
	Unreachables               types.Bool                         `tfsdk:"unreachables"`
	VrfForwarding              types.String                       `tfsdk:"vrf_forwarding"`
	Ipv4Address                types.String                       `tfsdk:"ipv4_address"`
	Ipv4AddressMask            types.String                       `tfsdk:"ipv4_address_mask"`
	Unnumbered                 types.String                       `tfsdk:"unnumbered"`
	EncapsulationDot1qVlanId   types.Int64                        `tfsdk:"encapsulation_dot1q_vlan_id"`
	ChannelGroupNumber         types.Int64                        `tfsdk:"channel_group_number"`
	ChannelGroupMode           types.String                       `tfsdk:"channel_group_mode"`
	IpDhcpRelaySourceInterface types.String                       `tfsdk:"ip_dhcp_relay_source_interface"`
	IpAccessGroupIn            types.String                       `tfsdk:"ip_access_group_in"`
	IpAccessGroupInEnable      types.Bool                         `tfsdk:"ip_access_group_in_enable"`
	IpAccessGroupOut           types.String                       `tfsdk:"ip_access_group_out"`
	IpAccessGroupOutEnable     types.Bool                         `tfsdk:"ip_access_group_out_enable"`
	HelperAddresses            []InterfaceEthernetHelperAddresses `tfsdk:"helper_addresses"`
	SourceTemplate             []InterfaceEthernetSourceTemplate  `tfsdk:"source_template"`
}

type InterfaceEthernetData struct {
	Device                     types.String                       `tfsdk:"device"`
	Id                         types.String                       `tfsdk:"id"`
	Type                       types.String                       `tfsdk:"type"`
	Name                       types.String                       `tfsdk:"name"`
	MediaType                  types.String                       `tfsdk:"media_type"`
	Switchport                 types.Bool                         `tfsdk:"switchport"`
	Description                types.String                       `tfsdk:"description"`
	Shutdown                   types.Bool                         `tfsdk:"shutdown"`
	IpProxyArp                 types.Bool                         `tfsdk:"ip_proxy_arp"`
	IpRedirects                types.Bool                         `tfsdk:"ip_redirects"`
	Unreachables               types.Bool                         `tfsdk:"unreachables"`
	VrfForwarding              types.String                       `tfsdk:"vrf_forwarding"`
	Ipv4Address                types.String                       `tfsdk:"ipv4_address"`
	Ipv4AddressMask            types.String                       `tfsdk:"ipv4_address_mask"`
	Unnumbered                 types.String                       `tfsdk:"unnumbered"`
	EncapsulationDot1qVlanId   types.Int64                        `tfsdk:"encapsulation_dot1q_vlan_id"`
	ChannelGroupNumber         types.Int64                        `tfsdk:"channel_group_number"`
	ChannelGroupMode           types.String                       `tfsdk:"channel_group_mode"`
	IpDhcpRelaySourceInterface types.String                       `tfsdk:"ip_dhcp_relay_source_interface"`
	IpAccessGroupIn            types.String                       `tfsdk:"ip_access_group_in"`
	IpAccessGroupInEnable      types.Bool                         `tfsdk:"ip_access_group_in_enable"`
	IpAccessGroupOut           types.String                       `tfsdk:"ip_access_group_out"`
	IpAccessGroupOutEnable     types.Bool                         `tfsdk:"ip_access_group_out_enable"`
	HelperAddresses            []InterfaceEthernetHelperAddresses `tfsdk:"helper_addresses"`
	SourceTemplate             []InterfaceEthernetSourceTemplate  `tfsdk:"source_template"`
}
type InterfaceEthernetHelperAddresses struct {
	Address types.String `tfsdk:"address"`
	Global  types.Bool   `tfsdk:"global"`
	Vrf     types.String `tfsdk:"vrf"`
}
type InterfaceEthernetSourceTemplate struct {
	TemplateName types.String `tfsdk:"template_name"`
	Merge        types.Bool   `tfsdk:"merge"`
}

func (data InterfaceEthernet) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/interface/%s=%v", url.QueryEscape(fmt.Sprintf("%v", data.Type.ValueString())), url.QueryEscape(fmt.Sprintf("%v", data.Name.ValueString())))
}

func (data InterfaceEthernetData) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/interface/%s=%v", url.QueryEscape(fmt.Sprintf("%v", data.Type.ValueString())), url.QueryEscape(fmt.Sprintf("%v", data.Name.ValueString())))
}

// if last path element has a key -> remove it
func (data InterfaceEthernet) getPathShort() string {
	path := data.getPath()
	re := regexp.MustCompile(`(.*)=[^\/]*$`)
	matches := re.FindStringSubmatch(path)
	if len(matches) <= 1 {
		return path
	}
	return matches[1]
}

func (data InterfaceEthernet) toBody(ctx context.Context) string {
	body := `{"` + helpers.LastElement(data.getPath()) + `":{}}`
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"name", data.Name.ValueString())
	}
	if !data.MediaType.IsNull() && !data.MediaType.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"media-type", data.MediaType.ValueString())
	}
	if !data.Switchport.IsNull() && !data.Switchport.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"switchport-conf.switchport", data.Switchport.ValueBool())
	}
	if !data.Description.IsNull() && !data.Description.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"description", data.Description.ValueString())
	}
	if !data.Shutdown.IsNull() && !data.Shutdown.IsUnknown() {
		if data.Shutdown.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"shutdown", map[string]string{})
		}
	}
	if !data.IpProxyArp.IsNull() && !data.IpProxyArp.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ip.proxy-arp", data.IpProxyArp.ValueBool())
	}
	if !data.IpRedirects.IsNull() && !data.IpRedirects.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ip.redirects", data.IpRedirects.ValueBool())
	}
	if !data.Unreachables.IsNull() && !data.Unreachables.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ip.Cisco-IOS-XE-icmp:unreachables", data.Unreachables.ValueBool())
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
	if !data.Unnumbered.IsNull() && !data.Unnumbered.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ip.unnumbered", data.Unnumbered.ValueString())
	}
	if !data.EncapsulationDot1qVlanId.IsNull() && !data.EncapsulationDot1qVlanId.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"encapsulation.dot1Q.vlan-id", strconv.FormatInt(data.EncapsulationDot1qVlanId.ValueInt64(), 10))
	}
	if !data.ChannelGroupNumber.IsNull() && !data.ChannelGroupNumber.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-ethernet:channel-group.number", strconv.FormatInt(data.ChannelGroupNumber.ValueInt64(), 10))
	}
	if !data.ChannelGroupMode.IsNull() && !data.ChannelGroupMode.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-ethernet:channel-group.mode", data.ChannelGroupMode.ValueString())
	}
	if !data.IpDhcpRelaySourceInterface.IsNull() && !data.IpDhcpRelaySourceInterface.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ip.dhcp.Cisco-IOS-XE-dhcp:relay.source-interface", data.IpDhcpRelaySourceInterface.ValueString())
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
	if len(data.SourceTemplate) > 0 {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"source.template.template-name", []interface{}{})
		for index, item := range data.SourceTemplate {
			if !item.TemplateName.IsNull() && !item.TemplateName.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"source.template.template-name"+"."+strconv.Itoa(index)+"."+"template-name", item.TemplateName.ValueString())
			}
			if !item.Merge.IsNull() && !item.Merge.IsUnknown() {
				if item.Merge.ValueBool() {
					body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"source.template.template-name"+"."+strconv.Itoa(index)+"."+"merge", map[string]string{})
				}
			}
		}
	}
	return body
}

func (data *InterfaceEthernet) updateFromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get(prefix + "media-type"); value.Exists() && !data.MediaType.IsNull() {
		data.MediaType = types.StringValue(value.String())
	} else {
		data.MediaType = types.StringNull()
	}
	if value := res.Get(prefix + "switchport-conf.switchport"); !data.Switchport.IsNull() {
		if value.Exists() {
			data.Switchport = types.BoolValue(value.Bool())
		}
	} else {
		data.Switchport = types.BoolNull()
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
	if value := res.Get(prefix + "ip.proxy-arp"); !data.IpProxyArp.IsNull() {
		if value.Exists() {
			data.IpProxyArp = types.BoolValue(value.Bool())
		}
	} else {
		data.IpProxyArp = types.BoolNull()
	}
	if value := res.Get(prefix + "ip.redirects"); !data.IpRedirects.IsNull() {
		if value.Exists() {
			data.IpRedirects = types.BoolValue(value.Bool())
		}
	} else {
		data.IpRedirects = types.BoolNull()
	}
	if value := res.Get(prefix + "ip.Cisco-IOS-XE-icmp:unreachables"); !data.Unreachables.IsNull() {
		if value.Exists() {
			data.Unreachables = types.BoolValue(value.Bool())
		}
	} else {
		data.Unreachables = types.BoolNull()
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
	if value := res.Get(prefix + "ip.unnumbered"); value.Exists() && !data.Unnumbered.IsNull() {
		data.Unnumbered = types.StringValue(value.String())
	} else {
		data.Unnumbered = types.StringNull()
	}
	if value := res.Get(prefix + "encapsulation.dot1Q.vlan-id"); value.Exists() && !data.EncapsulationDot1qVlanId.IsNull() {
		data.EncapsulationDot1qVlanId = types.Int64Value(value.Int())
	} else {
		data.EncapsulationDot1qVlanId = types.Int64Null()
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-ethernet:channel-group.number"); value.Exists() && !data.ChannelGroupNumber.IsNull() {
		data.ChannelGroupNumber = types.Int64Value(value.Int())
	} else {
		data.ChannelGroupNumber = types.Int64Null()
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-ethernet:channel-group.mode"); value.Exists() && !data.ChannelGroupMode.IsNull() {
		data.ChannelGroupMode = types.StringValue(value.String())
	} else {
		data.ChannelGroupMode = types.StringNull()
	}
	if value := res.Get(prefix + "ip.dhcp.Cisco-IOS-XE-dhcp:relay.source-interface"); value.Exists() && !data.IpDhcpRelaySourceInterface.IsNull() {
		data.IpDhcpRelaySourceInterface = types.StringValue(value.String())
	} else {
		data.IpDhcpRelaySourceInterface = types.StringNull()
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
	for i := range data.SourceTemplate {
		keys := [...]string{"template-name"}
		keyValues := [...]string{data.SourceTemplate[i].TemplateName.ValueString()}

		var r gjson.Result
		res.Get(prefix + "source.template.template-name").ForEach(
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
		if value := r.Get("template-name"); value.Exists() && !data.SourceTemplate[i].TemplateName.IsNull() {
			data.SourceTemplate[i].TemplateName = types.StringValue(value.String())
		} else {
			data.SourceTemplate[i].TemplateName = types.StringNull()
		}
		if value := r.Get("merge"); !data.SourceTemplate[i].Merge.IsNull() {
			if value.Exists() {
				data.SourceTemplate[i].Merge = types.BoolValue(true)
			} else {
				data.SourceTemplate[i].Merge = types.BoolValue(false)
			}
		} else {
			data.SourceTemplate[i].Merge = types.BoolNull()
		}
	}
}

func (data *InterfaceEthernetData) fromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "media-type"); value.Exists() {
		data.MediaType = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "switchport-conf.switchport"); value.Exists() {
		data.Switchport = types.BoolValue(value.Bool())
	} else {
		data.Switchport = types.BoolValue(false)
	}
	if value := res.Get(prefix + "description"); value.Exists() {
		data.Description = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "shutdown"); value.Exists() {
		data.Shutdown = types.BoolValue(true)
	} else {
		data.Shutdown = types.BoolValue(false)
	}
	if value := res.Get(prefix + "ip.proxy-arp"); value.Exists() {
		data.IpProxyArp = types.BoolValue(value.Bool())
	} else {
		data.IpProxyArp = types.BoolValue(false)
	}
	if value := res.Get(prefix + "ip.redirects"); value.Exists() {
		data.IpRedirects = types.BoolValue(value.Bool())
	} else {
		data.IpRedirects = types.BoolValue(false)
	}
	if value := res.Get(prefix + "ip.Cisco-IOS-XE-icmp:unreachables"); value.Exists() {
		data.Unreachables = types.BoolValue(value.Bool())
	} else {
		data.Unreachables = types.BoolValue(false)
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
	if value := res.Get(prefix + "ip.unnumbered"); value.Exists() {
		data.Unnumbered = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "encapsulation.dot1Q.vlan-id"); value.Exists() {
		data.EncapsulationDot1qVlanId = types.Int64Value(value.Int())
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-ethernet:channel-group.number"); value.Exists() {
		data.ChannelGroupNumber = types.Int64Value(value.Int())
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-ethernet:channel-group.mode"); value.Exists() {
		data.ChannelGroupMode = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "ip.dhcp.Cisco-IOS-XE-dhcp:relay.source-interface"); value.Exists() {
		data.IpDhcpRelaySourceInterface = types.StringValue(value.String())
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
		data.HelperAddresses = make([]InterfaceEthernetHelperAddresses, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := InterfaceEthernetHelperAddresses{}
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
	if value := res.Get(prefix + "source.template.template-name"); value.Exists() {
		data.SourceTemplate = make([]InterfaceEthernetSourceTemplate, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := InterfaceEthernetSourceTemplate{}
			if cValue := v.Get("template-name"); cValue.Exists() {
				item.TemplateName = types.StringValue(cValue.String())
			}
			if cValue := v.Get("merge"); cValue.Exists() {
				item.Merge = types.BoolValue(true)
			} else {
				item.Merge = types.BoolValue(false)
			}
			data.SourceTemplate = append(data.SourceTemplate, item)
			return true
		})
	}
}

func (data *InterfaceEthernet) getDeletedListItems(ctx context.Context, state InterfaceEthernet) []string {
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
	for i := range state.SourceTemplate {
		stateKeyValues := [...]string{state.SourceTemplate[i].TemplateName.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.SourceTemplate[i].TemplateName.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.SourceTemplate {
			found = true
			if state.SourceTemplate[i].TemplateName.ValueString() != data.SourceTemplate[j].TemplateName.ValueString() {
				found = false
			}
			if found {
				break
			}
		}
		if !found {
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/source/template/template-name=%v", state.getPath(), strings.Join(stateKeyValues[:], ",")))
		}
	}
	return deletedListItems
}

func (data *InterfaceEthernet) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	if !data.Shutdown.IsNull() && !data.Shutdown.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/shutdown", data.getPath()))
	}
	if !data.IpAccessGroupInEnable.IsNull() && !data.IpAccessGroupInEnable.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/ip/access-group/in/acl/in", data.getPath()))
	}
	if !data.IpAccessGroupOutEnable.IsNull() && !data.IpAccessGroupOutEnable.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/ip/access-group/out/acl/out", data.getPath()))
	}

	for i := range data.HelperAddresses {
		keyValues := [...]string{data.HelperAddresses[i].Address.ValueString()}
		if !data.HelperAddresses[i].Global.IsNull() && !data.HelperAddresses[i].Global.ValueBool() {
			emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/ip/helper-address=%v/global", data.getPath(), strings.Join(keyValues[:], ",")))
		}
	}

	for i := range data.SourceTemplate {
		keyValues := [...]string{data.SourceTemplate[i].TemplateName.ValueString()}
		if !data.SourceTemplate[i].Merge.IsNull() && !data.SourceTemplate[i].Merge.ValueBool() {
			emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/source/template/template-name=%v/merge", data.getPath(), strings.Join(keyValues[:], ",")))
		}
	}
	return emptyLeafsDelete
}

func (data *InterfaceEthernet) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	if !data.MediaType.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/media-type", data.getPath()))
	}
	if !data.Switchport.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/switchport-conf/switchport", data.getPath()))
	}
	if !data.Description.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/description", data.getPath()))
	}
	if !data.Shutdown.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/shutdown", data.getPath()))
	}
	if !data.IpProxyArp.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/ip/proxy-arp", data.getPath()))
	}
	if !data.IpRedirects.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/ip/redirects", data.getPath()))
	}
	if !data.Unreachables.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/ip/Cisco-IOS-XE-icmp:unreachables", data.getPath()))
	}
	if !data.VrfForwarding.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/vrf/forwarding", data.getPath()))
	}
	if !data.Ipv4Address.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/ip/address/primary", data.getPath()))
	}
	if !data.Ipv4AddressMask.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/ip/address/primary", data.getPath()))
	}
	if !data.Unnumbered.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/ip/unnumbered", data.getPath()))
	}
	if !data.EncapsulationDot1qVlanId.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/encapsulation/dot1Q/vlan-id", data.getPath()))
	}
	if !data.ChannelGroupNumber.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/Cisco-IOS-XE-ethernet:channel-group/number", data.getPath()))
	}
	if !data.ChannelGroupMode.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/Cisco-IOS-XE-ethernet:channel-group/mode", data.getPath()))
	}
	if !data.IpDhcpRelaySourceInterface.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/ip/dhcp/Cisco-IOS-XE-dhcp:relay/source-interface", data.getPath()))
	}
	if !data.IpAccessGroupIn.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/ip/access-group/in/acl", data.getPath()))
	}
	if !data.IpAccessGroupInEnable.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/ip/access-group/in/acl/in", data.getPath()))
	}
	if !data.IpAccessGroupOut.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/ip/access-group/out/acl", data.getPath()))
	}
	if !data.IpAccessGroupOutEnable.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/ip/access-group/out/acl/out", data.getPath()))
	}
	for i := range data.HelperAddresses {
		keyValues := [...]string{data.HelperAddresses[i].Address.ValueString()}

		deletePaths = append(deletePaths, fmt.Sprintf("%v/ip/helper-address=%v", data.getPath(), strings.Join(keyValues[:], ",")))
	}
	for i := range data.SourceTemplate {
		keyValues := [...]string{data.SourceTemplate[i].TemplateName.ValueString()}

		deletePaths = append(deletePaths, fmt.Sprintf("%v/source/template/template-name=%v", data.getPath(), strings.Join(keyValues[:], ",")))
	}
	return deletePaths
}
