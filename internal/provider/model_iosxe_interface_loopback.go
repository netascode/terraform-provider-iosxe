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

type InterfaceLoopback struct {
	Device                 types.String `tfsdk:"device"`
	Id                     types.String `tfsdk:"id"`
	Name                   types.Int64  `tfsdk:"name"`
	Description            types.String `tfsdk:"description"`
	Shutdown               types.Bool   `tfsdk:"shutdown"`
	VrfForwarding          types.String `tfsdk:"vrf_forwarding"`
	Ipv4Address            types.String `tfsdk:"ipv4_address"`
	Ipv4AddressMask        types.String `tfsdk:"ipv4_address_mask"`
	IpAccessGroupIn        types.String `tfsdk:"ip_access_group_in"`
	IpAccessGroupInEnable  types.Bool   `tfsdk:"ip_access_group_in_enable"`
	IpAccessGroupOut       types.String `tfsdk:"ip_access_group_out"`
	IpAccessGroupOutEnable types.Bool   `tfsdk:"ip_access_group_out_enable"`
}

func (data InterfaceLoopback) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/interface/Loopback=%v", url.QueryEscape(fmt.Sprintf("%v", data.Name.ValueInt64())))
}

// if last path element has a key -> remove it
func (data InterfaceLoopback) getPathShort() string {
	path := data.getPath()
	re := regexp.MustCompile(`(.*)=[^\/]*$`)
	matches := re.FindStringSubmatch(path)
	if len(matches) <= 1 {
		return path
	}
	return matches[1]
}

func (data InterfaceLoopback) toBody(ctx context.Context) string {
	body := `{"` + helpers.LastElement(data.getPath()) + `":{}}`
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"name", strconv.FormatInt(data.Name.ValueInt64(), 10))
	}
	if !data.Description.IsNull() && !data.Description.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"description", data.Description.ValueString())
	}
	if !data.Shutdown.IsNull() && !data.Shutdown.IsUnknown() {
		if data.Shutdown.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"shutdown", map[string]string{})
		}
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
	return body
}

func (data *InterfaceLoopback) updateFromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.Int64Value(value.Int())
	} else {
		data.Name = types.Int64Null()
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
}

func (data *InterfaceLoopback) fromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "description"); value.Exists() {
		data.Description = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "shutdown"); value.Exists() {
		data.Shutdown = types.BoolValue(true)
	} else {
		data.Shutdown = types.BoolValue(false)
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
}

func (data *InterfaceLoopback) getDeletedListItems(ctx context.Context, state InterfaceLoopback) []string {
	deletedListItems := make([]string, 0)
	return deletedListItems
}

func (data *InterfaceLoopback) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	if !data.Shutdown.IsNull() && !data.Shutdown.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/shutdown", data.getPath()))
	}
	if !data.IpAccessGroupInEnable.IsNull() && !data.IpAccessGroupInEnable.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/ip/access-group/in/apply-type/apply-intf/acl/in", data.getPath()))
	}
	if !data.IpAccessGroupOutEnable.IsNull() && !data.IpAccessGroupOutEnable.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/ip/access-group/out/apply-type/apply-intf/acl/out", data.getPath()))
	}
	return emptyLeafsDelete
}
