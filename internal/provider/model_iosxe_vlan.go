// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"
	"net/url"
	"regexp"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/terraform-provider-iosxe/internal/provider/helpers"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type VLAN struct {
	Device                 types.String `tfsdk:"device"`
	Id                     types.String `tfsdk:"id"`
	VlanId                 types.Int64  `tfsdk:"vlan_id"`
	RemoteSpan             types.Bool   `tfsdk:"remote_span"`
	PrivateVlanPrimary     types.Bool   `tfsdk:"private_vlan_primary"`
	PrivateVlanAssociation types.String `tfsdk:"private_vlan_association"`
	PrivateVlanCommunity   types.Bool   `tfsdk:"private_vlan_community"`
	PrivateVlanIsolated    types.Bool   `tfsdk:"private_vlan_isolated"`
	Name                   types.String `tfsdk:"name"`
	Shutdown               types.Bool   `tfsdk:"shutdown"`
}

func (data VLAN) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/vlan/Cisco-IOS-XE-vlan:vlan-list=%v", url.QueryEscape(fmt.Sprintf("%v", data.VlanId.Value)))
}

// if last path element has a key -> remove it
func (data VLAN) getPathShort() string {
	path := data.getPath()
	re := regexp.MustCompile(`(.*)=[^\/]*$`)
	matches := re.FindStringSubmatch(path)
	if len(matches) <= 1 {
		return path
	}
	return matches[1]
}

func (data VLAN) toBody() string {
	body := `{"` + helpers.LastElement(data.getPath()) + `":{}}`
	if !data.VlanId.Null && !data.VlanId.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"id", strconv.FormatInt(data.VlanId.Value, 10))
	}
	if !data.RemoteSpan.Null && !data.RemoteSpan.Unknown {
		if data.RemoteSpan.Value {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"remote-span", map[string]string{})
		}
	}
	if !data.PrivateVlanPrimary.Null && !data.PrivateVlanPrimary.Unknown {
		if data.PrivateVlanPrimary.Value {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"private-vlan.primary", map[string]string{})
		}
	}
	if !data.PrivateVlanAssociation.Null && !data.PrivateVlanAssociation.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"private-vlan.association", data.PrivateVlanAssociation.Value)
	}
	if !data.PrivateVlanCommunity.Null && !data.PrivateVlanCommunity.Unknown {
		if data.PrivateVlanCommunity.Value {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"private-vlan.community", map[string]string{})
		}
	}
	if !data.PrivateVlanIsolated.Null && !data.PrivateVlanIsolated.Unknown {
		if data.PrivateVlanIsolated.Value {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"private-vlan.isolated", map[string]string{})
		}
	}
	if !data.Name.Null && !data.Name.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"name", data.Name.Value)
	}
	if !data.Shutdown.Null && !data.Shutdown.Unknown {
		if data.Shutdown.Value {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"shutdown", map[string]string{})
		}
	}
	return body
}

func (data *VLAN) updateFromBody(res gjson.Result) {
	if value := res.Get(helpers.LastElement(data.getPath()) + "." + "id"); value.Exists() {
		data.VlanId.Value = value.Int()
	} else {
		data.VlanId.Null = true
	}
	if value := res.Get(helpers.LastElement(data.getPath()) + "." + "remote-span"); value.Exists() {
		data.RemoteSpan.Value = true
	} else {
		data.RemoteSpan.Value = false
	}
	if value := res.Get(helpers.LastElement(data.getPath()) + "." + "private-vlan.primary"); value.Exists() {
		data.PrivateVlanPrimary.Value = true
	} else {
		data.PrivateVlanPrimary.Value = false
	}
	if value := res.Get(helpers.LastElement(data.getPath()) + "." + "private-vlan.association"); value.Exists() {
		data.PrivateVlanAssociation.Value = value.String()
	} else {
		data.PrivateVlanAssociation.Null = true
	}
	if value := res.Get(helpers.LastElement(data.getPath()) + "." + "private-vlan.community"); value.Exists() {
		data.PrivateVlanCommunity.Value = true
	} else {
		data.PrivateVlanCommunity.Value = false
	}
	if value := res.Get(helpers.LastElement(data.getPath()) + "." + "private-vlan.isolated"); value.Exists() {
		data.PrivateVlanIsolated.Value = true
	} else {
		data.PrivateVlanIsolated.Value = false
	}
	if value := res.Get(helpers.LastElement(data.getPath()) + "." + "name"); value.Exists() {
		data.Name.Value = value.String()
	} else {
		data.Name.Null = true
	}
	if value := res.Get(helpers.LastElement(data.getPath()) + "." + "shutdown"); value.Exists() {
		data.Shutdown.Value = true
	} else {
		data.Shutdown.Value = false
	}
}

func (data *VLAN) fromBody(res gjson.Result) {
	if value := res.Get(helpers.LastElement(data.getPath()) + "." + "remote-span"); value.Exists() {
		data.RemoteSpan.Value = true
		data.RemoteSpan.Null = false
	} else {
		data.RemoteSpan.Value = false
		data.RemoteSpan.Null = false
	}
	if value := res.Get(helpers.LastElement(data.getPath()) + "." + "private-vlan.primary"); value.Exists() {
		data.PrivateVlanPrimary.Value = true
		data.PrivateVlanPrimary.Null = false
	} else {
		data.PrivateVlanPrimary.Value = false
		data.PrivateVlanPrimary.Null = false
	}
	if value := res.Get(helpers.LastElement(data.getPath()) + "." + "private-vlan.association"); value.Exists() {
		data.PrivateVlanAssociation.Value = value.String()
		data.PrivateVlanAssociation.Null = false
	}
	if value := res.Get(helpers.LastElement(data.getPath()) + "." + "private-vlan.community"); value.Exists() {
		data.PrivateVlanCommunity.Value = true
		data.PrivateVlanCommunity.Null = false
	} else {
		data.PrivateVlanCommunity.Value = false
		data.PrivateVlanCommunity.Null = false
	}
	if value := res.Get(helpers.LastElement(data.getPath()) + "." + "private-vlan.isolated"); value.Exists() {
		data.PrivateVlanIsolated.Value = true
		data.PrivateVlanIsolated.Null = false
	} else {
		data.PrivateVlanIsolated.Value = false
		data.PrivateVlanIsolated.Null = false
	}
	if value := res.Get(helpers.LastElement(data.getPath()) + "." + "name"); value.Exists() {
		data.Name.Value = value.String()
		data.Name.Null = false
	}
	if value := res.Get(helpers.LastElement(data.getPath()) + "." + "shutdown"); value.Exists() {
		data.Shutdown.Value = true
		data.Shutdown.Null = false
	} else {
		data.Shutdown.Value = false
		data.Shutdown.Null = false
	}
}

func (data *VLAN) setUnknownValues() {
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
	if data.RemoteSpan.Unknown {
		data.RemoteSpan.Unknown = false
		data.RemoteSpan.Null = true
	}
	if data.PrivateVlanPrimary.Unknown {
		data.PrivateVlanPrimary.Unknown = false
		data.PrivateVlanPrimary.Null = true
	}
	if data.PrivateVlanAssociation.Unknown {
		data.PrivateVlanAssociation.Unknown = false
		data.PrivateVlanAssociation.Null = true
	}
	if data.PrivateVlanCommunity.Unknown {
		data.PrivateVlanCommunity.Unknown = false
		data.PrivateVlanCommunity.Null = true
	}
	if data.PrivateVlanIsolated.Unknown {
		data.PrivateVlanIsolated.Unknown = false
		data.PrivateVlanIsolated.Null = true
	}
	if data.Name.Unknown {
		data.Name.Unknown = false
		data.Name.Null = true
	}
	if data.Shutdown.Unknown {
		data.Shutdown.Unknown = false
		data.Shutdown.Null = true
	}
}

func (data *VLAN) getDeletedListItems(state VLAN) []string {
	deletedListItems := make([]string, 0)
	return deletedListItems
}
