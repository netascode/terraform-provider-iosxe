// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"
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
	return fmt.Sprintf("Cisco-IOS-XE-native:native/vlan/Cisco-IOS-XE-vlan:configuration=%v", data.VlanId.Value)
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

func (data VLANConfiguration) toBody() string {
	body := `{"` + helpers.LastElement(data.getPath()) + `":{}}`
	if !data.VlanId.Null && !data.VlanId.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"vlan-id", strconv.FormatInt(data.VlanId.Value, 10))
	}
	if !data.Vni.Null && !data.Vni.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"member.vni", strconv.FormatInt(data.Vni.Value, 10))
	}
	if !data.AccessVfi.Null && !data.AccessVfi.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"member.access-vfi", data.AccessVfi.Value)
	}
	if !data.EvpnInstance.Null && !data.EvpnInstance.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"member.evpn-instance.evpn-instance", strconv.FormatInt(data.EvpnInstance.Value, 10))
	}
	if !data.EvpnInstanceVni.Null && !data.EvpnInstanceVni.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"member.evpn-instance.vni", strconv.FormatInt(data.EvpnInstanceVni.Value, 10))
	}
	return body
}

func (data *VLANConfiguration) updateFromBody(res gjson.Result) {
	if value := res.Get(helpers.LastElement(data.getPath()) + "." + "vlan-id"); value.Exists() {
		data.VlanId.Value = value.Int()
	} else {
		data.VlanId.Null = true
	}
	if value := res.Get(helpers.LastElement(data.getPath()) + "." + "member.vni"); value.Exists() {
		data.Vni.Value = value.Int()
	} else {
		data.Vni.Null = true
	}
	if value := res.Get(helpers.LastElement(data.getPath()) + "." + "member.access-vfi"); value.Exists() {
		data.AccessVfi.Value = value.String()
	} else {
		data.AccessVfi.Null = true
	}
	if value := res.Get(helpers.LastElement(data.getPath()) + "." + "member.evpn-instance.evpn-instance"); value.Exists() {
		data.EvpnInstance.Value = value.Int()
	} else {
		data.EvpnInstance.Null = true
	}
	if value := res.Get(helpers.LastElement(data.getPath()) + "." + "member.evpn-instance.vni"); value.Exists() {
		data.EvpnInstanceVni.Value = value.Int()
	} else {
		data.EvpnInstanceVni.Null = true
	}
}

func (data *VLANConfiguration) fromBody(res gjson.Result) {
	if value := res.Get(helpers.LastElement(data.getPath()) + "." + "member.vni"); value.Exists() {
		data.Vni.Value = value.Int()
		data.Vni.Null = false
	}
	if value := res.Get(helpers.LastElement(data.getPath()) + "." + "member.access-vfi"); value.Exists() {
		data.AccessVfi.Value = value.String()
		data.AccessVfi.Null = false
	}
	if value := res.Get(helpers.LastElement(data.getPath()) + "." + "member.evpn-instance.evpn-instance"); value.Exists() {
		data.EvpnInstance.Value = value.Int()
		data.EvpnInstance.Null = false
	}
	if value := res.Get(helpers.LastElement(data.getPath()) + "." + "member.evpn-instance.vni"); value.Exists() {
		data.EvpnInstanceVni.Value = value.Int()
		data.EvpnInstanceVni.Null = false
	}
}

func (data *VLANConfiguration) setUnknownValues() {
	if data.Device.Unknown {
		data.Device.Unknown = false
		data.Device.Null = true
	}
	if data.Id.Unknown {
		data.Id.Unknown = false
		data.Id.Null = true
	}
	if data.VlanId.Unknown {
		data.VlanId.Unknown = false
		data.VlanId.Null = true
	}
	if data.Vni.Unknown {
		data.Vni.Unknown = false
		data.Vni.Null = true
	}
	if data.AccessVfi.Unknown {
		data.AccessVfi.Unknown = false
		data.AccessVfi.Null = true
	}
	if data.EvpnInstance.Unknown {
		data.EvpnInstance.Unknown = false
		data.EvpnInstance.Null = true
	}
	if data.EvpnInstanceVni.Unknown {
		data.EvpnInstanceVni.Unknown = false
		data.EvpnInstanceVni.Null = true
	}
}

func (data *VLANConfiguration) getDeletedListItems(state VLANConfiguration) []string {
	deletedListItems := make([]string, 0)
	return deletedListItems
}
