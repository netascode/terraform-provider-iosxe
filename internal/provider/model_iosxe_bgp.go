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
	return fmt.Sprintf("Cisco-IOS-XE-native:native/router/Cisco-IOS-XE-bgp:bgp=%v", url.QueryEscape(fmt.Sprintf("%v", data.Asn.Value)))
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
	if !data.Asn.Null && !data.Asn.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"id", data.Asn.Value)
	}
	if !data.DefaultIpv4Unicast.Null && !data.DefaultIpv4Unicast.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"bgp.default.ipv4-unicast", data.DefaultIpv4Unicast.Value)
	}
	if !data.LogNeighborChanges.Null && !data.LogNeighborChanges.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"bgp.log-neighbor-changes", data.LogNeighborChanges.Value)
	}
	if !data.RouterIdLoopback.Null && !data.RouterIdLoopback.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"bgp.router-id.interface.Loopback", strconv.FormatInt(data.RouterIdLoopback.Value, 10))
	}
	return body
}

func (data *BGP) updateFromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "id"); value.Exists() {
		data.Asn.Value = value.String()
	} else {
		data.Asn.Null = true
	}
	if value := res.Get(prefix + "bgp.default.ipv4-unicast"); value.Exists() {
		data.DefaultIpv4Unicast.Value = value.Bool()
	} else {
		data.DefaultIpv4Unicast.Value = false
	}
	if value := res.Get(prefix + "bgp.log-neighbor-changes"); value.Exists() {
		data.LogNeighborChanges.Value = value.Bool()
	} else {
		data.LogNeighborChanges.Value = false
	}
	if value := res.Get(prefix + "bgp.router-id.interface.Loopback"); value.Exists() {
		data.RouterIdLoopback.Value = value.Int()
	} else {
		data.RouterIdLoopback.Null = true
	}
}

func (data *BGP) fromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "bgp.default.ipv4-unicast"); value.Exists() {
		data.DefaultIpv4Unicast.Value = value.Bool()
		data.DefaultIpv4Unicast.Null = false
	} else {
		data.DefaultIpv4Unicast.Value = false
		data.DefaultIpv4Unicast.Null = false
	}
	if value := res.Get(prefix + "bgp.log-neighbor-changes"); value.Exists() {
		data.LogNeighborChanges.Value = value.Bool()
		data.LogNeighborChanges.Null = false
	} else {
		data.LogNeighborChanges.Value = false
		data.LogNeighborChanges.Null = false
	}
	if value := res.Get(prefix + "bgp.router-id.interface.Loopback"); value.Exists() {
		data.RouterIdLoopback.Value = value.Int()
		data.RouterIdLoopback.Null = false
	}
}

func (data *BGP) setUnknownValues(ctx context.Context) {
	if data.Device.Unknown {
		data.Device.Unknown = false
		data.Device.Null = true
	}
	if data.Id.Unknown {
		data.Id.Unknown = false
		data.Id.Null = true
	}
	if data.Asn.Unknown {
		data.Asn.Unknown = false
		data.Asn.Null = true
	}
	if data.DefaultIpv4Unicast.Unknown {
		data.DefaultIpv4Unicast.Unknown = false
		data.DefaultIpv4Unicast.Null = true
	}
	if data.LogNeighborChanges.Unknown {
		data.LogNeighborChanges.Unknown = false
		data.LogNeighborChanges.Null = true
	}
	if data.RouterIdLoopback.Unknown {
		data.RouterIdLoopback.Unknown = false
		data.RouterIdLoopback.Null = true
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
