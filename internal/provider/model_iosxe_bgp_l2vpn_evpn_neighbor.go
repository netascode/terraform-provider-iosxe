// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/terraform-provider-iosxe/internal/provider/helpers"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type BGPL2VPNEVPNNeighbor struct {
	Device               types.String `tfsdk:"device"`
	Id                   types.String `tfsdk:"id"`
	Asn                  types.String `tfsdk:"asn"`
	Ip                   types.String `tfsdk:"ip"`
	Activate             types.Bool   `tfsdk:"activate"`
	SendCommunity        types.String `tfsdk:"send_community"`
	RouteReflectorClient types.Bool   `tfsdk:"route_reflector_client"`
}

func (data BGPL2VPNEVPNNeighbor) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/router/Cisco-IOS-XE-bgp:bgp=%v/address-family/no-vrf/l2vpn=evpn/l2vpn-evpn/neighbor=%s", url.QueryEscape(fmt.Sprintf("%v", data.Asn.ValueString())), url.QueryEscape(fmt.Sprintf("%v", data.Ip.ValueString())))
}

// if last path element has a key -> remove it
func (data BGPL2VPNEVPNNeighbor) getPathShort() string {
	path := data.getPath()
	re := regexp.MustCompile(`(.*)=[^\/]*$`)
	matches := re.FindStringSubmatch(path)
	if len(matches) <= 1 {
		return path
	}
	return matches[1]
}

func (data BGPL2VPNEVPNNeighbor) toBody(ctx context.Context) string {
	body := `{"` + helpers.LastElement(data.getPath()) + `":{}}`
	if !data.Ip.IsNull() && !data.Ip.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"id", data.Ip.ValueString())
	}
	if !data.Activate.IsNull() && !data.Activate.IsUnknown() {
		if data.Activate.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"activate", map[string]string{})
		}
	}
	if !data.SendCommunity.IsNull() && !data.SendCommunity.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"send-community.send-community-where", data.SendCommunity.ValueString())
	}
	if !data.RouteReflectorClient.IsNull() && !data.RouteReflectorClient.IsUnknown() {
		if data.RouteReflectorClient.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"route-reflector-client", map[string]string{})
		}
	}
	return body
}

func (data *BGPL2VPNEVPNNeighbor) updateFromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "id"); value.Exists() && !data.Ip.IsNull() {
		data.Ip = types.StringValue(value.String())
	} else {
		data.Ip = types.StringNull()
	}
	if value := res.Get(prefix + "activate"); !data.Activate.IsNull() {
		if value.Exists() {
			data.Activate = types.BoolValue(true)
		} else {
			data.Activate = types.BoolValue(false)
		}
	} else {
		data.Activate = types.BoolNull()
	}
	if value := res.Get(prefix + "send-community.send-community-where"); value.Exists() && !data.SendCommunity.IsNull() {
		data.SendCommunity = types.StringValue(value.String())
	} else {
		data.SendCommunity = types.StringNull()
	}
	if value := res.Get(prefix + "route-reflector-client"); !data.RouteReflectorClient.IsNull() {
		if value.Exists() {
			data.RouteReflectorClient = types.BoolValue(true)
		} else {
			data.RouteReflectorClient = types.BoolValue(false)
		}
	} else {
		data.RouteReflectorClient = types.BoolNull()
	}
}

func (data *BGPL2VPNEVPNNeighbor) fromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "activate"); value.Exists() {
		data.Activate = types.BoolValue(true)
	} else {
		data.Activate = types.BoolValue(false)
	}
	if value := res.Get(prefix + "send-community.send-community-where"); value.Exists() {
		data.SendCommunity = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "route-reflector-client"); value.Exists() {
		data.RouteReflectorClient = types.BoolValue(true)
	} else {
		data.RouteReflectorClient = types.BoolValue(false)
	}
}

func (data *BGPL2VPNEVPNNeighbor) getDeletedListItems(ctx context.Context, state BGPL2VPNEVPNNeighbor) []string {
	deletedListItems := make([]string, 0)
	return deletedListItems
}

func (data *BGPL2VPNEVPNNeighbor) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	if !data.Activate.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/activate", data.getPath()))
	}
	if !data.RouteReflectorClient.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/route-reflector-client", data.getPath()))
	}
	return emptyLeafsDelete
}
