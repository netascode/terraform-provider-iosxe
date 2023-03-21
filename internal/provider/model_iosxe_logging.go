// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/terraform-provider-iosxe/internal/provider/helpers"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type Logging struct {
	Device              types.String                 `tfsdk:"device"`
	Id                  types.String                 `tfsdk:"id"`
	MonitorSeverity     types.String                 `tfsdk:"monitor_severity"`
	BufferedSize        types.Int64                  `tfsdk:"buffered_size"`
	BufferedSeverity    types.String                 `tfsdk:"buffered_severity"`
	ConsoleSeverity     types.String                 `tfsdk:"console_severity"`
	Facility            types.String                 `tfsdk:"facility"`
	HistorySize         types.Int64                  `tfsdk:"history_size"`
	HistorySeverity     types.String                 `tfsdk:"history_severity"`
	Trap                types.Bool                   `tfsdk:"trap"`
	TrapSeverity        types.String                 `tfsdk:"trap_severity"`
	OriginIdType        types.String                 `tfsdk:"origin_id_type"`
	OriginIdName        types.String                 `tfsdk:"origin_id_name"`
	FileName            types.String                 `tfsdk:"file_name"`
	FileMaxSize         types.Int64                  `tfsdk:"file_max_size"`
	FileMinSize         types.Int64                  `tfsdk:"file_min_size"`
	FileSeverity        types.String                 `tfsdk:"file_severity"`
	SourceInterface     types.String                 `tfsdk:"source_interface"`
	SourceInterfacesVrf []LoggingSourceInterfacesVrf `tfsdk:"source_interfaces_vrf"`
	Ipv4Hosts           []LoggingIpv4Hosts           `tfsdk:"ipv4_hosts"`
	Ipv4VrfHosts        []LoggingIpv4VrfHosts        `tfsdk:"ipv4_vrf_hosts"`
	Ipv6Hosts           []LoggingIpv6Hosts           `tfsdk:"ipv6_hosts"`
	Ipv6VrfHosts        []LoggingIpv6VrfHosts        `tfsdk:"ipv6_vrf_hosts"`
}
type LoggingSourceInterfacesVrf struct {
	Vrf           types.String `tfsdk:"vrf"`
	InterfaceName types.String `tfsdk:"interface_name"`
}
type LoggingIpv4Hosts struct {
	Ipv4Host types.String `tfsdk:"ipv4_host"`
}
type LoggingIpv4VrfHosts struct {
	Ipv4Host types.String `tfsdk:"ipv4_host"`
	Vrf      types.String `tfsdk:"vrf"`
}
type LoggingIpv6Hosts struct {
	Ipv6Host types.String `tfsdk:"ipv6_host"`
}
type LoggingIpv6VrfHosts struct {
	Ipv6Host types.String `tfsdk:"ipv6_host"`
	Vrf      types.String `tfsdk:"vrf"`
}

func (data Logging) getPath() string {
	return "Cisco-IOS-XE-native:native/logging"
}

// if last path element has a key -> remove it
func (data Logging) getPathShort() string {
	path := data.getPath()
	re := regexp.MustCompile(`(.*)=[^\/]*$`)
	matches := re.FindStringSubmatch(path)
	if len(matches) <= 1 {
		return path
	}
	return matches[1]
}

func (data Logging) toBody(ctx context.Context) string {
	body := `{"` + helpers.LastElement(data.getPath()) + `":{}}`
	if !data.MonitorSeverity.IsNull() && !data.MonitorSeverity.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"monitor-config.common-config.monitor.severity", data.MonitorSeverity.ValueString())
	}
	if !data.BufferedSize.IsNull() && !data.BufferedSize.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"buffered.size-value", strconv.FormatInt(data.BufferedSize.ValueInt64(), 10))
	}
	if !data.BufferedSeverity.IsNull() && !data.BufferedSeverity.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"buffered.severity-level", data.BufferedSeverity.ValueString())
	}
	if !data.ConsoleSeverity.IsNull() && !data.ConsoleSeverity.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"console-config.common-config.console.severity", data.ConsoleSeverity.ValueString())
	}
	if !data.Facility.IsNull() && !data.Facility.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"facility", data.Facility.ValueString())
	}
	if !data.HistorySize.IsNull() && !data.HistorySize.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"history.size", strconv.FormatInt(data.HistorySize.ValueInt64(), 10))
	}
	if !data.HistorySeverity.IsNull() && !data.HistorySeverity.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"history.severity-level", data.HistorySeverity.ValueString())
	}
	if !data.Trap.IsNull() && !data.Trap.IsUnknown() {
		if data.Trap.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"trap", map[string]string{})
		}
	}
	if !data.TrapSeverity.IsNull() && !data.TrapSeverity.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"trap.severity", data.TrapSeverity.ValueString())
	}
	if !data.OriginIdType.IsNull() && !data.OriginIdType.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"origin-id.type-value", data.OriginIdType.ValueString())
	}
	if !data.OriginIdName.IsNull() && !data.OriginIdName.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"origin-id.string", data.OriginIdName.ValueString())
	}
	if !data.FileName.IsNull() && !data.FileName.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"file.name", data.FileName.ValueString())
	}
	if !data.FileMaxSize.IsNull() && !data.FileMaxSize.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"file.max-size", strconv.FormatInt(data.FileMaxSize.ValueInt64(), 10))
	}
	if !data.FileMinSize.IsNull() && !data.FileMinSize.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"file.min-size", strconv.FormatInt(data.FileMinSize.ValueInt64(), 10))
	}
	if !data.FileSeverity.IsNull() && !data.FileSeverity.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"file.severity", data.FileSeverity.ValueString())
	}
	if !data.SourceInterface.IsNull() && !data.SourceInterface.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"source-interface-conf.interface-name-non-vrf", data.SourceInterface.ValueString())
	}
	if len(data.SourceInterfacesVrf) > 0 {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"source-interface-conf.source-interface-vrf", []interface{}{})
		for index, item := range data.SourceInterfacesVrf {
			if !item.Vrf.IsNull() && !item.Vrf.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"source-interface-conf.source-interface-vrf"+"."+strconv.Itoa(index)+"."+"vrf", item.Vrf.ValueString())
			}
			if !item.InterfaceName.IsNull() && !item.InterfaceName.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"source-interface-conf.source-interface-vrf"+"."+strconv.Itoa(index)+"."+"interface-name", item.InterfaceName.ValueString())
			}
		}
	}
	if len(data.Ipv4Hosts) > 0 {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"host.ipv4-host-list", []interface{}{})
		for index, item := range data.Ipv4Hosts {
			if !item.Ipv4Host.IsNull() && !item.Ipv4Host.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"host.ipv4-host-list"+"."+strconv.Itoa(index)+"."+"ipv4-host", item.Ipv4Host.ValueString())
			}
		}
	}
	if len(data.Ipv4VrfHosts) > 0 {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"host.ipv4-host-vrf-list", []interface{}{})
		for index, item := range data.Ipv4VrfHosts {
			if !item.Ipv4Host.IsNull() && !item.Ipv4Host.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"host.ipv4-host-vrf-list"+"."+strconv.Itoa(index)+"."+"ipv4-host", item.Ipv4Host.ValueString())
			}
			if !item.Vrf.IsNull() && !item.Vrf.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"host.ipv4-host-vrf-list"+"."+strconv.Itoa(index)+"."+"vrf", item.Vrf.ValueString())
			}
		}
	}
	if len(data.Ipv6Hosts) > 0 {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"host.ipv6.ipv6-host-list", []interface{}{})
		for index, item := range data.Ipv6Hosts {
			if !item.Ipv6Host.IsNull() && !item.Ipv6Host.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"host.ipv6.ipv6-host-list"+"."+strconv.Itoa(index)+"."+"ipv6-host", item.Ipv6Host.ValueString())
			}
		}
	}
	if len(data.Ipv6VrfHosts) > 0 {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"host.ipv6.ipv6-host-vrf-list", []interface{}{})
		for index, item := range data.Ipv6VrfHosts {
			if !item.Ipv6Host.IsNull() && !item.Ipv6Host.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"host.ipv6.ipv6-host-vrf-list"+"."+strconv.Itoa(index)+"."+"ipv6-host", item.Ipv6Host.ValueString())
			}
			if !item.Vrf.IsNull() && !item.Vrf.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"host.ipv6.ipv6-host-vrf-list"+"."+strconv.Itoa(index)+"."+"vrf", item.Vrf.ValueString())
			}
		}
	}
	return body
}

func (data *Logging) updateFromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "monitor-config.common-config.monitor.severity"); value.Exists() && !data.MonitorSeverity.IsNull() {
		data.MonitorSeverity = types.StringValue(value.String())
	} else {
		data.MonitorSeverity = types.StringNull()
	}
	if value := res.Get(prefix + "buffered.size-value"); value.Exists() && !data.BufferedSize.IsNull() {
		data.BufferedSize = types.Int64Value(value.Int())
	} else {
		data.BufferedSize = types.Int64Null()
	}
	if value := res.Get(prefix + "buffered.severity-level"); value.Exists() && !data.BufferedSeverity.IsNull() {
		data.BufferedSeverity = types.StringValue(value.String())
	} else {
		data.BufferedSeverity = types.StringNull()
	}
	if value := res.Get(prefix + "console-config.common-config.console.severity"); value.Exists() && !data.ConsoleSeverity.IsNull() {
		data.ConsoleSeverity = types.StringValue(value.String())
	} else {
		data.ConsoleSeverity = types.StringNull()
	}
	if value := res.Get(prefix + "facility"); value.Exists() && !data.Facility.IsNull() {
		data.Facility = types.StringValue(value.String())
	} else {
		data.Facility = types.StringNull()
	}
	if value := res.Get(prefix + "history.size"); value.Exists() && !data.HistorySize.IsNull() {
		data.HistorySize = types.Int64Value(value.Int())
	} else {
		data.HistorySize = types.Int64Null()
	}
	if value := res.Get(prefix + "history.severity-level"); value.Exists() && !data.HistorySeverity.IsNull() {
		data.HistorySeverity = types.StringValue(value.String())
	} else {
		data.HistorySeverity = types.StringNull()
	}
	if value := res.Get(prefix + "trap"); !data.Trap.IsNull() {
		if value.Exists() {
			data.Trap = types.BoolValue(true)
		} else {
			data.Trap = types.BoolValue(false)
		}
	} else {
		data.Trap = types.BoolNull()
	}
	if value := res.Get(prefix + "trap.severity"); value.Exists() && !data.TrapSeverity.IsNull() {
		data.TrapSeverity = types.StringValue(value.String())
	} else {
		data.TrapSeverity = types.StringNull()
	}
	if value := res.Get(prefix + "origin-id.type-value"); value.Exists() && !data.OriginIdType.IsNull() {
		data.OriginIdType = types.StringValue(value.String())
	} else {
		data.OriginIdType = types.StringNull()
	}
	if value := res.Get(prefix + "origin-id.string"); value.Exists() && !data.OriginIdName.IsNull() {
		data.OriginIdName = types.StringValue(value.String())
	} else {
		data.OriginIdName = types.StringNull()
	}
	if value := res.Get(prefix + "file.name"); value.Exists() && !data.FileName.IsNull() {
		data.FileName = types.StringValue(value.String())
	} else {
		data.FileName = types.StringNull()
	}
	if value := res.Get(prefix + "file.max-size"); value.Exists() && !data.FileMaxSize.IsNull() {
		data.FileMaxSize = types.Int64Value(value.Int())
	} else {
		data.FileMaxSize = types.Int64Null()
	}
	if value := res.Get(prefix + "file.min-size"); value.Exists() && !data.FileMinSize.IsNull() {
		data.FileMinSize = types.Int64Value(value.Int())
	} else {
		data.FileMinSize = types.Int64Null()
	}
	if value := res.Get(prefix + "file.severity"); value.Exists() && !data.FileSeverity.IsNull() {
		data.FileSeverity = types.StringValue(value.String())
	} else {
		data.FileSeverity = types.StringNull()
	}
	if value := res.Get(prefix + "source-interface-conf.interface-name-non-vrf"); value.Exists() && !data.SourceInterface.IsNull() {
		data.SourceInterface = types.StringValue(value.String())
	} else {
		data.SourceInterface = types.StringNull()
	}
	for i := range data.SourceInterfacesVrf {
		keys := [...]string{"vrf"}
		keyValues := [...]string{data.SourceInterfacesVrf[i].Vrf.ValueString()}

		var r gjson.Result
		res.Get(prefix + "source-interface-conf.source-interface-vrf").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() == keyValues[ik] {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		if value := r.Get("vrf"); value.Exists() && !data.SourceInterfacesVrf[i].Vrf.IsNull() {
			data.SourceInterfacesVrf[i].Vrf = types.StringValue(value.String())
		} else {
			data.SourceInterfacesVrf[i].Vrf = types.StringNull()
		}
		if value := r.Get("interface-name"); value.Exists() && !data.SourceInterfacesVrf[i].InterfaceName.IsNull() {
			data.SourceInterfacesVrf[i].InterfaceName = types.StringValue(value.String())
		} else {
			data.SourceInterfacesVrf[i].InterfaceName = types.StringNull()
		}
	}
	for i := range data.Ipv4Hosts {
		keys := [...]string{"ipv4-host"}
		keyValues := [...]string{data.Ipv4Hosts[i].Ipv4Host.ValueString()}

		var r gjson.Result
		res.Get(prefix + "host.ipv4-host-list").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() == keyValues[ik] {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		if value := r.Get("ipv4-host"); value.Exists() && !data.Ipv4Hosts[i].Ipv4Host.IsNull() {
			data.Ipv4Hosts[i].Ipv4Host = types.StringValue(value.String())
		} else {
			data.Ipv4Hosts[i].Ipv4Host = types.StringNull()
		}
	}
	for i := range data.Ipv4VrfHosts {
		keys := [...]string{"ipv4-host", "vrf"}
		keyValues := [...]string{data.Ipv4VrfHosts[i].Ipv4Host.ValueString(), data.Ipv4VrfHosts[i].Vrf.ValueString()}

		var r gjson.Result
		res.Get(prefix + "host.ipv4-host-vrf-list").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() == keyValues[ik] {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		if value := r.Get("ipv4-host"); value.Exists() && !data.Ipv4VrfHosts[i].Ipv4Host.IsNull() {
			data.Ipv4VrfHosts[i].Ipv4Host = types.StringValue(value.String())
		} else {
			data.Ipv4VrfHosts[i].Ipv4Host = types.StringNull()
		}
		if value := r.Get("vrf"); value.Exists() && !data.Ipv4VrfHosts[i].Vrf.IsNull() {
			data.Ipv4VrfHosts[i].Vrf = types.StringValue(value.String())
		} else {
			data.Ipv4VrfHosts[i].Vrf = types.StringNull()
		}
	}
	for i := range data.Ipv6Hosts {
		keys := [...]string{"ipv6-host"}
		keyValues := [...]string{data.Ipv6Hosts[i].Ipv6Host.ValueString()}

		var r gjson.Result
		res.Get(prefix + "host.ipv6.ipv6-host-list").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() == keyValues[ik] {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		if value := r.Get("ipv6-host"); value.Exists() && !data.Ipv6Hosts[i].Ipv6Host.IsNull() {
			data.Ipv6Hosts[i].Ipv6Host = types.StringValue(value.String())
		} else {
			data.Ipv6Hosts[i].Ipv6Host = types.StringNull()
		}
	}
	for i := range data.Ipv6VrfHosts {
		keys := [...]string{"ipv6-host", "vrf"}
		keyValues := [...]string{data.Ipv6VrfHosts[i].Ipv6Host.ValueString(), data.Ipv6VrfHosts[i].Vrf.ValueString()}

		var r gjson.Result
		res.Get(prefix + "host.ipv6.ipv6-host-vrf-list").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() == keyValues[ik] {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		if value := r.Get("ipv6-host"); value.Exists() && !data.Ipv6VrfHosts[i].Ipv6Host.IsNull() {
			data.Ipv6VrfHosts[i].Ipv6Host = types.StringValue(value.String())
		} else {
			data.Ipv6VrfHosts[i].Ipv6Host = types.StringNull()
		}
		if value := r.Get("vrf"); value.Exists() && !data.Ipv6VrfHosts[i].Vrf.IsNull() {
			data.Ipv6VrfHosts[i].Vrf = types.StringValue(value.String())
		} else {
			data.Ipv6VrfHosts[i].Vrf = types.StringNull()
		}
	}
}

func (data *Logging) fromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "monitor-config.common-config.monitor.severity"); value.Exists() {
		data.MonitorSeverity = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "buffered.size-value"); value.Exists() {
		data.BufferedSize = types.Int64Value(value.Int())
	}
	if value := res.Get(prefix + "buffered.severity-level"); value.Exists() {
		data.BufferedSeverity = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "console-config.common-config.console.severity"); value.Exists() {
		data.ConsoleSeverity = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "facility"); value.Exists() {
		data.Facility = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "history.size"); value.Exists() {
		data.HistorySize = types.Int64Value(value.Int())
	}
	if value := res.Get(prefix + "history.severity-level"); value.Exists() {
		data.HistorySeverity = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "trap"); value.Exists() {
		data.Trap = types.BoolValue(true)
	} else {
		data.Trap = types.BoolValue(false)
	}
	if value := res.Get(prefix + "trap.severity"); value.Exists() {
		data.TrapSeverity = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "origin-id.type-value"); value.Exists() {
		data.OriginIdType = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "origin-id.string"); value.Exists() {
		data.OriginIdName = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "file.name"); value.Exists() {
		data.FileName = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "file.max-size"); value.Exists() {
		data.FileMaxSize = types.Int64Value(value.Int())
	}
	if value := res.Get(prefix + "file.min-size"); value.Exists() {
		data.FileMinSize = types.Int64Value(value.Int())
	}
	if value := res.Get(prefix + "file.severity"); value.Exists() {
		data.FileSeverity = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "source-interface-conf.interface-name-non-vrf"); value.Exists() {
		data.SourceInterface = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "source-interface-conf.source-interface-vrf"); value.Exists() {
		data.SourceInterfacesVrf = make([]LoggingSourceInterfacesVrf, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := LoggingSourceInterfacesVrf{}
			if cValue := v.Get("vrf"); cValue.Exists() {
				item.Vrf = types.StringValue(cValue.String())
			}
			if cValue := v.Get("interface-name"); cValue.Exists() {
				item.InterfaceName = types.StringValue(cValue.String())
			}
			data.SourceInterfacesVrf = append(data.SourceInterfacesVrf, item)
			return true
		})
	}
	if value := res.Get(prefix + "host.ipv4-host-list"); value.Exists() {
		data.Ipv4Hosts = make([]LoggingIpv4Hosts, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := LoggingIpv4Hosts{}
			if cValue := v.Get("ipv4-host"); cValue.Exists() {
				item.Ipv4Host = types.StringValue(cValue.String())
			}
			data.Ipv4Hosts = append(data.Ipv4Hosts, item)
			return true
		})
	}
	if value := res.Get(prefix + "host.ipv4-host-vrf-list"); value.Exists() {
		data.Ipv4VrfHosts = make([]LoggingIpv4VrfHosts, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := LoggingIpv4VrfHosts{}
			if cValue := v.Get("ipv4-host"); cValue.Exists() {
				item.Ipv4Host = types.StringValue(cValue.String())
			}
			if cValue := v.Get("vrf"); cValue.Exists() {
				item.Vrf = types.StringValue(cValue.String())
			}
			data.Ipv4VrfHosts = append(data.Ipv4VrfHosts, item)
			return true
		})
	}
	if value := res.Get(prefix + "host.ipv6.ipv6-host-list"); value.Exists() {
		data.Ipv6Hosts = make([]LoggingIpv6Hosts, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := LoggingIpv6Hosts{}
			if cValue := v.Get("ipv6-host"); cValue.Exists() {
				item.Ipv6Host = types.StringValue(cValue.String())
			}
			data.Ipv6Hosts = append(data.Ipv6Hosts, item)
			return true
		})
	}
	if value := res.Get(prefix + "host.ipv6.ipv6-host-vrf-list"); value.Exists() {
		data.Ipv6VrfHosts = make([]LoggingIpv6VrfHosts, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := LoggingIpv6VrfHosts{}
			if cValue := v.Get("ipv6-host"); cValue.Exists() {
				item.Ipv6Host = types.StringValue(cValue.String())
			}
			if cValue := v.Get("vrf"); cValue.Exists() {
				item.Vrf = types.StringValue(cValue.String())
			}
			data.Ipv6VrfHosts = append(data.Ipv6VrfHosts, item)
			return true
		})
	}
}

func (data *Logging) getDeletedListItems(ctx context.Context, state Logging) []string {
	deletedListItems := make([]string, 0)
	for i := range state.SourceInterfacesVrf {
		stateKeyValues := [...]string{state.SourceInterfacesVrf[i].Vrf.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.SourceInterfacesVrf[i].Vrf.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.SourceInterfacesVrf {
			found = true
			if state.SourceInterfacesVrf[i].Vrf.ValueString() != data.SourceInterfacesVrf[j].Vrf.ValueString() {
				found = false
			}
			if found {
				break
			}
		}
		if !found {
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/source-interface-conf/source-interface-vrf=%v", state.getPath(), strings.Join(stateKeyValues[:], ",")))
		}
	}
	for i := range state.Ipv4Hosts {
		stateKeyValues := [...]string{state.Ipv4Hosts[i].Ipv4Host.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.Ipv4Hosts[i].Ipv4Host.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.Ipv4Hosts {
			found = true
			if state.Ipv4Hosts[i].Ipv4Host.ValueString() != data.Ipv4Hosts[j].Ipv4Host.ValueString() {
				found = false
			}
			if found {
				break
			}
		}
		if !found {
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/host/ipv4-host-list=%v", state.getPath(), strings.Join(stateKeyValues[:], ",")))
		}
	}
	for i := range state.Ipv4VrfHosts {
		stateKeyValues := [...]string{state.Ipv4VrfHosts[i].Ipv4Host.ValueString(), state.Ipv4VrfHosts[i].Vrf.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.Ipv4VrfHosts[i].Ipv4Host.ValueString()).IsZero() {
			emptyKeys = false
		}
		if !reflect.ValueOf(state.Ipv4VrfHosts[i].Vrf.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.Ipv4VrfHosts {
			found = true
			if state.Ipv4VrfHosts[i].Ipv4Host.ValueString() != data.Ipv4VrfHosts[j].Ipv4Host.ValueString() {
				found = false
			}
			if state.Ipv4VrfHosts[i].Vrf.ValueString() != data.Ipv4VrfHosts[j].Vrf.ValueString() {
				found = false
			}
			if found {
				break
			}
		}
		if !found {
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/host/ipv4-host-vrf-list=%v", state.getPath(), strings.Join(stateKeyValues[:], ",")))
		}
	}
	for i := range state.Ipv6Hosts {
		stateKeyValues := [...]string{state.Ipv6Hosts[i].Ipv6Host.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.Ipv6Hosts[i].Ipv6Host.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.Ipv6Hosts {
			found = true
			if state.Ipv6Hosts[i].Ipv6Host.ValueString() != data.Ipv6Hosts[j].Ipv6Host.ValueString() {
				found = false
			}
			if found {
				break
			}
		}
		if !found {
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/host/ipv6/ipv6-host-list=%v", state.getPath(), strings.Join(stateKeyValues[:], ",")))
		}
	}
	for i := range state.Ipv6VrfHosts {
		stateKeyValues := [...]string{state.Ipv6VrfHosts[i].Ipv6Host.ValueString(), state.Ipv6VrfHosts[i].Vrf.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.Ipv6VrfHosts[i].Ipv6Host.ValueString()).IsZero() {
			emptyKeys = false
		}
		if !reflect.ValueOf(state.Ipv6VrfHosts[i].Vrf.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.Ipv6VrfHosts {
			found = true
			if state.Ipv6VrfHosts[i].Ipv6Host.ValueString() != data.Ipv6VrfHosts[j].Ipv6Host.ValueString() {
				found = false
			}
			if state.Ipv6VrfHosts[i].Vrf.ValueString() != data.Ipv6VrfHosts[j].Vrf.ValueString() {
				found = false
			}
			if found {
				break
			}
		}
		if !found {
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/host/ipv6/ipv6-host-vrf-list=%v", state.getPath(), strings.Join(stateKeyValues[:], ",")))
		}
	}
	return deletedListItems
}

func (data *Logging) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	if !data.Trap.IsNull() && !data.Trap.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/trap", data.getPath()))
	}

	return emptyLeafsDelete
}
