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

func (data VLANConfiguration) getPath() string {
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
	if value := res.Get(prefix + "vlan-id"); value.Exists() {
		data.VlanId = types.Int64Value(value.Int())
	} else {
		data.VlanId = types.Int64Null()
	}
	if value := res.Get(prefix + "member.vni"); value.Exists() {
		data.Vni = types.Int64Value(value.Int())
	} else {
		data.Vni = types.Int64Null()
	}
	if value := res.Get(prefix + "member.access-vfi"); value.Exists() {
		data.AccessVfi = types.StringValue(value.String())
	} else {
		data.AccessVfi = types.StringNull()
	}
	if value := res.Get(prefix + "member.evpn-instance.evpn-instance"); value.Exists() {
		data.EvpnInstance = types.Int64Value(value.Int())
	} else {
		data.EvpnInstance = types.Int64Null()
	}
	if value := res.Get(prefix + "member.evpn-instance.vni"); value.Exists() {
		data.EvpnInstanceVni = types.Int64Value(value.Int())
	} else {
		data.EvpnInstanceVni = types.Int64Null()
	}
}

func (data *VLANConfiguration) fromBody(ctx context.Context, res gjson.Result) {
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

func (data *VLANConfiguration) setUnknownValues(ctx context.Context) {
	if data.Device.IsUnknown() {
		data.Device = types.StringNull()
	}
	if data.Id.IsUnknown() {
		data.Id = types.StringNull()
	}
	if data.VlanId.IsUnknown() {
		data.VlanId = types.Int64Null()
	}
	if data.Vni.IsUnknown() {
		data.Vni = types.Int64Null()
	}
	if data.AccessVfi.IsUnknown() {
		data.AccessVfi = types.StringNull()
	}
	if data.EvpnInstance.IsUnknown() {
		data.EvpnInstance = types.Int64Null()
	}
	if data.EvpnInstanceVni.IsUnknown() {
		data.EvpnInstanceVni = types.Int64Null()
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
