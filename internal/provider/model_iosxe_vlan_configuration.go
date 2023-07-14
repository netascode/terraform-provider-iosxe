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
	"regexp"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/terraform-provider-iosxe/internal/provider/helpers"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type VLANConfiguration struct {
	Device          types.String `tfsdk:"device"`
	Id              types.String `tfsdk:"id"`
	VlanId          types.Int64  `tfsdk:"vlan_id"`
	Vni             types.Int64  `tfsdk:"vni"`
	AccessVfi       types.String `tfsdk:"access_vfi"`
	EvpnInstance    types.Int64  `tfsdk:"evpn_instance"`
	EvpnInstanceVni types.Int64  `tfsdk:"evpn_instance_vni"`
}

type VLANConfigurationData struct {
	Device          types.String `tfsdk:"device"`
	Id              types.String `tfsdk:"id"`
	VlanId          types.Int64  `tfsdk:"vlan_id"`
	Vni             types.Int64  `tfsdk:"vni"`
	AccessVfi       types.String `tfsdk:"access_vfi"`
	EvpnInstance    types.Int64  `tfsdk:"evpn_instance"`
	EvpnInstanceVni types.Int64  `tfsdk:"evpn_instance_vni"`
}

func (data VLANConfiguration) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/vlan/Cisco-IOS-XE-vlan:configuration=%v", url.QueryEscape(fmt.Sprintf("%v", data.VlanId.ValueInt64())))
}

func (data VLANConfigurationData) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/vlan/Cisco-IOS-XE-vlan:configuration=%v", url.QueryEscape(fmt.Sprintf("%v", data.VlanId.ValueInt64())))
}

// if last path element has a key -> remove it
func (data VLANConfiguration) getPathShort() string {
	path := data.getPath()
	re := regexp.MustCompile(`(.*)=[^\/]*$`)
	matches := re.FindStringSubmatch(path)
	if len(matches) <= 1 {
		return path
	}
	return matches[1]
}

func (data VLANConfiguration) toBody(ctx context.Context) string {
	body := `{"` + helpers.LastElement(data.getPath()) + `":{}}`
	if !data.VlanId.IsNull() && !data.VlanId.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"vlan-id", strconv.FormatInt(data.VlanId.ValueInt64(), 10))
	}
	if !data.Vni.IsNull() && !data.Vni.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"member.vni", strconv.FormatInt(data.Vni.ValueInt64(), 10))
	}
	if !data.AccessVfi.IsNull() && !data.AccessVfi.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"member.access-vfi", data.AccessVfi.ValueString())
	}
	if !data.EvpnInstance.IsNull() && !data.EvpnInstance.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"member.evpn-instance.evpn-instance", strconv.FormatInt(data.EvpnInstance.ValueInt64(), 10))
	}
	if !data.EvpnInstanceVni.IsNull() && !data.EvpnInstanceVni.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"member.evpn-instance.vni", strconv.FormatInt(data.EvpnInstanceVni.ValueInt64(), 10))
	}
	return body
}

func (data *VLANConfiguration) updateFromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "vlan-id"); value.Exists() && !data.VlanId.IsNull() {
		data.VlanId = types.Int64Value(value.Int())
	} else {
		data.VlanId = types.Int64Null()
	}
	if value := res.Get(prefix + "member.vni"); value.Exists() && !data.Vni.IsNull() {
		data.Vni = types.Int64Value(value.Int())
	} else {
		data.Vni = types.Int64Null()
	}
	if value := res.Get(prefix + "member.access-vfi"); value.Exists() && !data.AccessVfi.IsNull() {
		data.AccessVfi = types.StringValue(value.String())
	} else {
		data.AccessVfi = types.StringNull()
	}
	if value := res.Get(prefix + "member.evpn-instance.evpn-instance"); value.Exists() && !data.EvpnInstance.IsNull() {
		data.EvpnInstance = types.Int64Value(value.Int())
	} else {
		data.EvpnInstance = types.Int64Null()
	}
	if value := res.Get(prefix + "member.evpn-instance.vni"); value.Exists() && !data.EvpnInstanceVni.IsNull() {
		data.EvpnInstanceVni = types.Int64Value(value.Int())
	} else {
		data.EvpnInstanceVni = types.Int64Null()
	}
}

func (data *VLANConfigurationData) fromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "member.vni"); value.Exists() {
		data.Vni = types.Int64Value(value.Int())
	}
	if value := res.Get(prefix + "member.access-vfi"); value.Exists() {
		data.AccessVfi = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "member.evpn-instance.evpn-instance"); value.Exists() {
		data.EvpnInstance = types.Int64Value(value.Int())
	}
	if value := res.Get(prefix + "member.evpn-instance.vni"); value.Exists() {
		data.EvpnInstanceVni = types.Int64Value(value.Int())
	}
}

func (data *VLANConfiguration) getDeletedListItems(ctx context.Context, state VLANConfiguration) []string {
	deletedListItems := make([]string, 0)
	return deletedListItems
}

func (data *VLANConfiguration) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	return emptyLeafsDelete
}

func (data *VLANConfiguration) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	if !data.Vni.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/member/vni", data.getPath()))
	}
	if !data.AccessVfi.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/member/access-vfi", data.getPath()))
	}
	if !data.EvpnInstance.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/member/evpn-instance/evpn-instance", data.getPath()))
	}
	if !data.EvpnInstanceVni.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/member/evpn-instance/vni", data.getPath()))
	}
	return deletePaths
}
