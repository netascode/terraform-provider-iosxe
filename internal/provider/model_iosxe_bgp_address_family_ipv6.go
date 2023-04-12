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

type BGPAddressFamilyIPv6 struct {
	Device types.String `tfsdk:"device"`
	Id     types.String `tfsdk:"id"`
	Asn    types.String `tfsdk:"asn"`
	AfName types.String `tfsdk:"af_name"`
}

func (data BGPAddressFamilyIPv6) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/router/Cisco-IOS-XE-bgp:bgp=%v/address-family/no-vrf/ipv6=%s", url.QueryEscape(fmt.Sprintf("%v", data.Asn.ValueString())), url.QueryEscape(fmt.Sprintf("%v", data.AfName.ValueString())))
}

// if last path element has a key -> remove it
func (data BGPAddressFamilyIPv6) getPathShort() string {
	path := data.getPath()
	re := regexp.MustCompile(`(.*)=[^\/]*$`)
	matches := re.FindStringSubmatch(path)
	if len(matches) <= 1 {
		return path
	}
	return matches[1]
}

func (data BGPAddressFamilyIPv6) toBody(ctx context.Context) string {
	body := `{"` + helpers.LastElement(data.getPath()) + `":{}}`
	if !data.AfName.IsNull() && !data.AfName.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"af-name", data.AfName.ValueString())
	}
	return body
}

func (data *BGPAddressFamilyIPv6) updateFromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "af-name"); value.Exists() && !data.AfName.IsNull() {
		data.AfName = types.StringValue(value.String())
	} else {
		data.AfName = types.StringNull()
	}
}

func (data *BGPAddressFamilyIPv6) fromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
}

func (data *BGPAddressFamilyIPv6) getDeletedListItems(ctx context.Context, state BGPAddressFamilyIPv6) []string {
	deletedListItems := make([]string, 0)
	return deletedListItems
}

func (data *BGPAddressFamilyIPv6) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	return emptyLeafsDelete
}
