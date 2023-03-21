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

type BGPNeighbor struct {
	Device               types.String `tfsdk:"device"`
	Id                   types.String `tfsdk:"id"`
	Asn                  types.String `tfsdk:"asn"`
	Ip                   types.String `tfsdk:"ip"`
	RemoteAs             types.String `tfsdk:"remote_as"`
	Description          types.String `tfsdk:"description"`
	Shutdown             types.Bool   `tfsdk:"shutdown"`
	UpdateSourceLoopback types.String `tfsdk:"update_source_loopback"`
}

func (data BGPNeighbor) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/router/Cisco-IOS-XE-bgp:bgp=%v/neighbor=%s", url.QueryEscape(fmt.Sprintf("%v", data.Asn.ValueString())), url.QueryEscape(fmt.Sprintf("%v", data.Ip.ValueString())))
}

// if last path element has a key -> remove it
func (data BGPNeighbor) getPathShort() string {
	path := data.getPath()
	re := regexp.MustCompile(`(.*)=[^\/]*$`)
	matches := re.FindStringSubmatch(path)
	if len(matches) <= 1 {
		return path
	}
	return matches[1]
}

func (data BGPNeighbor) toBody(ctx context.Context) string {
	body := `{"` + helpers.LastElement(data.getPath()) + `":{}}`
	if !data.Ip.IsNull() && !data.Ip.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"id", data.Ip.ValueString())
	}
	if !data.RemoteAs.IsNull() && !data.RemoteAs.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"remote-as", data.RemoteAs.ValueString())
	}
	if !data.Description.IsNull() && !data.Description.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"description", data.Description.ValueString())
	}
	if !data.Shutdown.IsNull() && !data.Shutdown.IsUnknown() {
		if data.Shutdown.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"shutdown", map[string]string{})
		}
	}
	if !data.UpdateSourceLoopback.IsNull() && !data.UpdateSourceLoopback.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"update-source.interface.Loopback", data.UpdateSourceLoopback.ValueString())
	}
	return body
}

func (data *BGPNeighbor) updateFromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "id"); value.Exists() && !data.Ip.IsNull() {
		data.Ip = types.StringValue(value.String())
	} else {
		data.Ip = types.StringNull()
	}
	if value := res.Get(prefix + "remote-as"); value.Exists() && !data.RemoteAs.IsNull() {
		data.RemoteAs = types.StringValue(value.String())
	} else {
		data.RemoteAs = types.StringNull()
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
	if value := res.Get(prefix + "update-source.interface.Loopback"); value.Exists() && !data.UpdateSourceLoopback.IsNull() {
		data.UpdateSourceLoopback = types.StringValue(value.String())
	} else {
		data.UpdateSourceLoopback = types.StringNull()
	}
}

func (data *BGPNeighbor) fromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "remote-as"); value.Exists() {
		data.RemoteAs = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "description"); value.Exists() {
		data.Description = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "shutdown"); value.Exists() {
		data.Shutdown = types.BoolValue(true)
	} else {
		data.Shutdown = types.BoolValue(false)
	}
	if value := res.Get(prefix + "update-source.interface.Loopback"); value.Exists() {
		data.UpdateSourceLoopback = types.StringValue(value.String())
	}
}

func (data *BGPNeighbor) getDeletedListItems(ctx context.Context, state BGPNeighbor) []string {
	deletedListItems := make([]string, 0)
	return deletedListItems
}

func (data *BGPNeighbor) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	if !data.Shutdown.IsNull() && !data.Shutdown.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/shutdown", data.getPath()))
	}
	return emptyLeafsDelete
}
