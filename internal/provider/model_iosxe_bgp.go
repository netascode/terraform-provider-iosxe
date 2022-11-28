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

type BGP struct {
	Device             types.String `tfsdk:"device"`
	Id                 types.String `tfsdk:"id"`
	Asn                types.String `tfsdk:"asn"`
	DefaultIpv4Unicast types.Bool   `tfsdk:"default_ipv4_unicast"`
	LogNeighborChanges types.Bool   `tfsdk:"log_neighbor_changes"`
	RouterIdLoopback   types.Int64  `tfsdk:"router_id_loopback"`
}

func (data BGP) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/router/Cisco-IOS-XE-bgp:bgp=%v", url.QueryEscape(fmt.Sprintf("%v", data.Asn.ValueString())))
}

// if last path element has a key -> remove it
func (data BGP) getPathShort() string {
	path := data.getPath()
	re := regexp.MustCompile(`(.*)=[^\/]*$`)
	matches := re.FindStringSubmatch(path)
	if len(matches) <= 1 {
		return path
	}
	return matches[1]
}

func (data BGP) toBody(ctx context.Context) string {
	body := `{"` + helpers.LastElement(data.getPath()) + `":{}}`
	if !data.Asn.IsNull() && !data.Asn.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"id", data.Asn.ValueString())
	}
	if !data.DefaultIpv4Unicast.IsNull() && !data.DefaultIpv4Unicast.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"bgp.default.ipv4-unicast", data.DefaultIpv4Unicast.ValueBool())
	}
	if !data.LogNeighborChanges.IsNull() && !data.LogNeighborChanges.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"bgp.log-neighbor-changes", data.LogNeighborChanges.ValueBool())
	}
	if !data.RouterIdLoopback.IsNull() && !data.RouterIdLoopback.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"bgp.router-id.interface.Loopback", strconv.FormatInt(data.RouterIdLoopback.ValueInt64(), 10))
	}
	return body
}

func (data *BGP) updateFromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "id"); value.Exists() {
		data.Asn = types.StringValue(value.String())
	} else {
		data.Asn = types.StringNull()
	}
	if value := res.Get(prefix + "bgp.default.ipv4-unicast"); value.Exists() {
		data.DefaultIpv4Unicast = types.BoolValue(value.Bool())
	} else {
		data.DefaultIpv4Unicast = types.BoolValue(false)
	}
	if value := res.Get(prefix + "bgp.log-neighbor-changes"); value.Exists() {
		data.LogNeighborChanges = types.BoolValue(value.Bool())
	} else {
		data.LogNeighborChanges = types.BoolValue(false)
	}
	if value := res.Get(prefix + "bgp.router-id.interface.Loopback"); value.Exists() {
		data.RouterIdLoopback = types.Int64Value(value.Int())
	} else {
		data.RouterIdLoopback = types.Int64Null()
	}
}

func (data *BGP) fromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "bgp.default.ipv4-unicast"); value.Exists() {
		data.DefaultIpv4Unicast = types.BoolValue(value.Bool())
	} else {
		data.DefaultIpv4Unicast = types.BoolValue(false)
	}
	if value := res.Get(prefix + "bgp.log-neighbor-changes"); value.Exists() {
		data.LogNeighborChanges = types.BoolValue(value.Bool())
	} else {
		data.LogNeighborChanges = types.BoolValue(false)
	}
	if value := res.Get(prefix + "bgp.router-id.interface.Loopback"); value.Exists() {
		data.RouterIdLoopback = types.Int64Value(value.Int())
	}
}

func (data *BGP) setUnknownValues(ctx context.Context) {
	if data.Device.IsUnknown() {
		data.Device = types.StringNull()
	}
	if data.Id.IsUnknown() {
		data.Id = types.StringNull()
	}
	if data.Asn.IsUnknown() {
		data.Asn = types.StringNull()
	}
	if data.DefaultIpv4Unicast.IsUnknown() {
		data.DefaultIpv4Unicast = types.BoolNull()
	}
	if data.LogNeighborChanges.IsUnknown() {
		data.LogNeighborChanges = types.BoolNull()
	}
	if data.RouterIdLoopback.IsUnknown() {
		data.RouterIdLoopback = types.Int64Null()
	}
}

func (data *BGP) getDeletedListItems(ctx context.Context, state BGP) []string {
	deletedListItems := make([]string, 0)
	return deletedListItems
}

func (data *BGP) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	return emptyLeafsDelete
}
