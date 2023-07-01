// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/terraform-provider-iosxe/internal/provider/helpers"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type Service struct {
	Device                              types.String `tfsdk:"device"`
	Id                                  types.String `tfsdk:"id"`
	DeleteMode                          types.String `tfsdk:"delete_mode"`
	Pad                                 types.Bool   `tfsdk:"pad"`
	PasswordEncryption                  types.Bool   `tfsdk:"password_encryption"`
	PasswordRecovery                    types.Bool   `tfsdk:"password_recovery"`
	Timestamps                          types.Bool   `tfsdk:"timestamps"`
	TimestampsDebug                     types.Bool   `tfsdk:"timestamps_debug"`
	TimestampsDebugDatetime             types.Bool   `tfsdk:"timestamps_debug_datetime"`
	TimestampsDebugDatetimeMsec         types.Bool   `tfsdk:"timestamps_debug_datetime_msec"`
	TimestampsDebugDatetimeLocaltime    types.Bool   `tfsdk:"timestamps_debug_datetime_localtime"`
	TimestampsDebugDatetimeShowTimezone types.Bool   `tfsdk:"timestamps_debug_datetime_show_timezone"`
	TimestampsDebugDatetimeYear         types.Bool   `tfsdk:"timestamps_debug_datetime_year"`
	TimestampsDebugUptime               types.Bool   `tfsdk:"timestamps_debug_uptime"`
	TimestampsLog                       types.Bool   `tfsdk:"timestamps_log"`
	TimestampsLogDatetime               types.Bool   `tfsdk:"timestamps_log_datetime"`
	TimestampsLogDatetimeMsec           types.Bool   `tfsdk:"timestamps_log_datetime_msec"`
	TimestampsLogDatetimeLocaltime      types.Bool   `tfsdk:"timestamps_log_datetime_localtime"`
	TimestampsLogDatetimeShowTimezone   types.Bool   `tfsdk:"timestamps_log_datetime_show_timezone"`
	TimestampsLogDatetimeYear           types.Bool   `tfsdk:"timestamps_log_datetime_year"`
	TimestampsLogUptime                 types.Bool   `tfsdk:"timestamps_log_uptime"`
	Dhcp                                types.Bool   `tfsdk:"dhcp"`
	TcpKeepalivesIn                     types.Bool   `tfsdk:"tcp_keepalives_in"`
	TcpKeepalivesOut                    types.Bool   `tfsdk:"tcp_keepalives_out"`
}

type ServiceData struct {
	Device                              types.String `tfsdk:"device"`
	Id                                  types.String `tfsdk:"id"`
	Pad                                 types.Bool   `tfsdk:"pad"`
	PasswordEncryption                  types.Bool   `tfsdk:"password_encryption"`
	PasswordRecovery                    types.Bool   `tfsdk:"password_recovery"`
	Timestamps                          types.Bool   `tfsdk:"timestamps"`
	TimestampsDebug                     types.Bool   `tfsdk:"timestamps_debug"`
	TimestampsDebugDatetime             types.Bool   `tfsdk:"timestamps_debug_datetime"`
	TimestampsDebugDatetimeMsec         types.Bool   `tfsdk:"timestamps_debug_datetime_msec"`
	TimestampsDebugDatetimeLocaltime    types.Bool   `tfsdk:"timestamps_debug_datetime_localtime"`
	TimestampsDebugDatetimeShowTimezone types.Bool   `tfsdk:"timestamps_debug_datetime_show_timezone"`
	TimestampsDebugDatetimeYear         types.Bool   `tfsdk:"timestamps_debug_datetime_year"`
	TimestampsDebugUptime               types.Bool   `tfsdk:"timestamps_debug_uptime"`
	TimestampsLog                       types.Bool   `tfsdk:"timestamps_log"`
	TimestampsLogDatetime               types.Bool   `tfsdk:"timestamps_log_datetime"`
	TimestampsLogDatetimeMsec           types.Bool   `tfsdk:"timestamps_log_datetime_msec"`
	TimestampsLogDatetimeLocaltime      types.Bool   `tfsdk:"timestamps_log_datetime_localtime"`
	TimestampsLogDatetimeShowTimezone   types.Bool   `tfsdk:"timestamps_log_datetime_show_timezone"`
	TimestampsLogDatetimeYear           types.Bool   `tfsdk:"timestamps_log_datetime_year"`
	TimestampsLogUptime                 types.Bool   `tfsdk:"timestamps_log_uptime"`
	Dhcp                                types.Bool   `tfsdk:"dhcp"`
	TcpKeepalivesIn                     types.Bool   `tfsdk:"tcp_keepalives_in"`
	TcpKeepalivesOut                    types.Bool   `tfsdk:"tcp_keepalives_out"`
}

func (data Service) getPath() string {
	return "Cisco-IOS-XE-native:native/service"
}

func (data ServiceData) getPath() string {
	return "Cisco-IOS-XE-native:native/service"
}

// if last path element has a key -> remove it
func (data Service) getPathShort() string {
	path := data.getPath()
	re := regexp.MustCompile(`(.*)=[^\/]*$`)
	matches := re.FindStringSubmatch(path)
	if len(matches) <= 1 {
		return path
	}
	return matches[1]
}

func (data Service) toBody(ctx context.Context) string {
	body := `{"` + helpers.LastElement(data.getPath()) + `":{}}`
	if !data.Pad.IsNull() && !data.Pad.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"pad-conf.pad", data.Pad.ValueBool())
	}
	if !data.PasswordEncryption.IsNull() && !data.PasswordEncryption.IsUnknown() {
		if data.PasswordEncryption.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"password-encryption", map[string]string{})
		}
	}
	if !data.PasswordRecovery.IsNull() && !data.PasswordRecovery.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"password-recovery", data.PasswordRecovery.ValueBool())
	}
	if !data.Timestamps.IsNull() && !data.Timestamps.IsUnknown() {
		if data.Timestamps.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"timestamps", map[string]string{})
		}
	}
	if !data.TimestampsDebug.IsNull() && !data.TimestampsDebug.IsUnknown() {
		if data.TimestampsDebug.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"timestamps.debug-config", map[string]string{})
		}
	}
	if !data.TimestampsDebugDatetime.IsNull() && !data.TimestampsDebugDatetime.IsUnknown() {
		if data.TimestampsDebugDatetime.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"timestamps.debug-config.datetime", map[string]string{})
		}
	}
	if !data.TimestampsDebugDatetimeMsec.IsNull() && !data.TimestampsDebugDatetimeMsec.IsUnknown() {
		if data.TimestampsDebugDatetimeMsec.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"timestamps.debug-config.datetime.msec", map[string]string{})
		}
	}
	if !data.TimestampsDebugDatetimeLocaltime.IsNull() && !data.TimestampsDebugDatetimeLocaltime.IsUnknown() {
		if data.TimestampsDebugDatetimeLocaltime.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"timestamps.debug-config.datetime.localtime", map[string]string{})
		}
	}
	if !data.TimestampsDebugDatetimeShowTimezone.IsNull() && !data.TimestampsDebugDatetimeShowTimezone.IsUnknown() {
		if data.TimestampsDebugDatetimeShowTimezone.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"timestamps.debug-config.datetime.show-timezone", map[string]string{})
		}
	}
	if !data.TimestampsDebugDatetimeYear.IsNull() && !data.TimestampsDebugDatetimeYear.IsUnknown() {
		if data.TimestampsDebugDatetimeYear.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"timestamps.debug-config.datetime.year", map[string]string{})
		}
	}
	if !data.TimestampsDebugUptime.IsNull() && !data.TimestampsDebugUptime.IsUnknown() {
		if data.TimestampsDebugUptime.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"timestamps.debug-config.uptime", map[string]string{})
		}
	}
	if !data.TimestampsLog.IsNull() && !data.TimestampsLog.IsUnknown() {
		if data.TimestampsLog.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"timestamps.log-config", map[string]string{})
		}
	}
	if !data.TimestampsLogDatetime.IsNull() && !data.TimestampsLogDatetime.IsUnknown() {
		if data.TimestampsLogDatetime.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"timestamps.log-config.datetime", map[string]string{})
		}
	}
	if !data.TimestampsLogDatetimeMsec.IsNull() && !data.TimestampsLogDatetimeMsec.IsUnknown() {
		if data.TimestampsLogDatetimeMsec.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"timestamps.log-config.datetime.msec", map[string]string{})
		}
	}
	if !data.TimestampsLogDatetimeLocaltime.IsNull() && !data.TimestampsLogDatetimeLocaltime.IsUnknown() {
		if data.TimestampsLogDatetimeLocaltime.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"timestamps.log-config.datetime.localtime", map[string]string{})
		}
	}
	if !data.TimestampsLogDatetimeShowTimezone.IsNull() && !data.TimestampsLogDatetimeShowTimezone.IsUnknown() {
		if data.TimestampsLogDatetimeShowTimezone.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"timestamps.log-config.datetime.show-timezone", map[string]string{})
		}
	}
	if !data.TimestampsLogDatetimeYear.IsNull() && !data.TimestampsLogDatetimeYear.IsUnknown() {
		if data.TimestampsLogDatetimeYear.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"timestamps.log-config.datetime.year", map[string]string{})
		}
	}
	if !data.TimestampsLogUptime.IsNull() && !data.TimestampsLogUptime.IsUnknown() {
		if data.TimestampsLogUptime.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"timestamps.log-config.uptime", map[string]string{})
		}
	}
	if !data.Dhcp.IsNull() && !data.Dhcp.IsUnknown() {
		if data.Dhcp.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"dhcp", map[string]string{})
		}
	}
	if !data.TcpKeepalivesIn.IsNull() && !data.TcpKeepalivesIn.IsUnknown() {
		if data.TcpKeepalivesIn.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"tcp-keepalives-in", map[string]string{})
		}
	}
	if !data.TcpKeepalivesOut.IsNull() && !data.TcpKeepalivesOut.IsUnknown() {
		if data.TcpKeepalivesOut.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"tcp-keepalives-out", map[string]string{})
		}
	}
	return body
}

func (data *Service) updateFromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "pad-conf.pad"); !data.Pad.IsNull() {
		if value.Exists() {
			data.Pad = types.BoolValue(value.Bool())
		}
	} else {
		data.Pad = types.BoolNull()
	}
	if value := res.Get(prefix + "password-encryption"); !data.PasswordEncryption.IsNull() {
		if value.Exists() {
			data.PasswordEncryption = types.BoolValue(true)
		} else {
			data.PasswordEncryption = types.BoolValue(false)
		}
	} else {
		data.PasswordEncryption = types.BoolNull()
	}
	if value := res.Get(prefix + "password-recovery"); !data.PasswordRecovery.IsNull() {
		if value.Exists() {
			data.PasswordRecovery = types.BoolValue(value.Bool())
		}
	} else {
		data.PasswordRecovery = types.BoolNull()
	}
	if value := res.Get(prefix + "timestamps"); !data.Timestamps.IsNull() {
		if value.Exists() {
			data.Timestamps = types.BoolValue(true)
		} else {
			data.Timestamps = types.BoolValue(false)
		}
	} else {
		data.Timestamps = types.BoolNull()
	}
	if value := res.Get(prefix + "timestamps.debug-config"); !data.TimestampsDebug.IsNull() {
		if value.Exists() {
			data.TimestampsDebug = types.BoolValue(true)
		} else {
			data.TimestampsDebug = types.BoolValue(false)
		}
	} else {
		data.TimestampsDebug = types.BoolNull()
	}
	if value := res.Get(prefix + "timestamps.debug-config.datetime"); !data.TimestampsDebugDatetime.IsNull() {
		if value.Exists() {
			data.TimestampsDebugDatetime = types.BoolValue(true)
		} else {
			data.TimestampsDebugDatetime = types.BoolValue(false)
		}
	} else {
		data.TimestampsDebugDatetime = types.BoolNull()
	}
	if value := res.Get(prefix + "timestamps.debug-config.datetime.msec"); !data.TimestampsDebugDatetimeMsec.IsNull() {
		if value.Exists() {
			data.TimestampsDebugDatetimeMsec = types.BoolValue(true)
		} else {
			data.TimestampsDebugDatetimeMsec = types.BoolValue(false)
		}
	} else {
		data.TimestampsDebugDatetimeMsec = types.BoolNull()
	}
	if value := res.Get(prefix + "timestamps.debug-config.datetime.localtime"); !data.TimestampsDebugDatetimeLocaltime.IsNull() {
		if value.Exists() {
			data.TimestampsDebugDatetimeLocaltime = types.BoolValue(true)
		} else {
			data.TimestampsDebugDatetimeLocaltime = types.BoolValue(false)
		}
	} else {
		data.TimestampsDebugDatetimeLocaltime = types.BoolNull()
	}
	if value := res.Get(prefix + "timestamps.debug-config.datetime.show-timezone"); !data.TimestampsDebugDatetimeShowTimezone.IsNull() {
		if value.Exists() {
			data.TimestampsDebugDatetimeShowTimezone = types.BoolValue(true)
		} else {
			data.TimestampsDebugDatetimeShowTimezone = types.BoolValue(false)
		}
	} else {
		data.TimestampsDebugDatetimeShowTimezone = types.BoolNull()
	}
	if value := res.Get(prefix + "timestamps.debug-config.datetime.year"); !data.TimestampsDebugDatetimeYear.IsNull() {
		if value.Exists() {
			data.TimestampsDebugDatetimeYear = types.BoolValue(true)
		} else {
			data.TimestampsDebugDatetimeYear = types.BoolValue(false)
		}
	} else {
		data.TimestampsDebugDatetimeYear = types.BoolNull()
	}
	if value := res.Get(prefix + "timestamps.debug-config.uptime"); !data.TimestampsDebugUptime.IsNull() {
		if value.Exists() {
			data.TimestampsDebugUptime = types.BoolValue(true)
		} else {
			data.TimestampsDebugUptime = types.BoolValue(false)
		}
	} else {
		data.TimestampsDebugUptime = types.BoolNull()
	}
	if value := res.Get(prefix + "timestamps.log-config"); !data.TimestampsLog.IsNull() {
		if value.Exists() {
			data.TimestampsLog = types.BoolValue(true)
		} else {
			data.TimestampsLog = types.BoolValue(false)
		}
	} else {
		data.TimestampsLog = types.BoolNull()
	}
	if value := res.Get(prefix + "timestamps.log-config.datetime"); !data.TimestampsLogDatetime.IsNull() {
		if value.Exists() {
			data.TimestampsLogDatetime = types.BoolValue(true)
		} else {
			data.TimestampsLogDatetime = types.BoolValue(false)
		}
	} else {
		data.TimestampsLogDatetime = types.BoolNull()
	}
	if value := res.Get(prefix + "timestamps.log-config.datetime.msec"); !data.TimestampsLogDatetimeMsec.IsNull() {
		if value.Exists() {
			data.TimestampsLogDatetimeMsec = types.BoolValue(true)
		} else {
			data.TimestampsLogDatetimeMsec = types.BoolValue(false)
		}
	} else {
		data.TimestampsLogDatetimeMsec = types.BoolNull()
	}
	if value := res.Get(prefix + "timestamps.log-config.datetime.localtime"); !data.TimestampsLogDatetimeLocaltime.IsNull() {
		if value.Exists() {
			data.TimestampsLogDatetimeLocaltime = types.BoolValue(true)
		} else {
			data.TimestampsLogDatetimeLocaltime = types.BoolValue(false)
		}
	} else {
		data.TimestampsLogDatetimeLocaltime = types.BoolNull()
	}
	if value := res.Get(prefix + "timestamps.log-config.datetime.show-timezone"); !data.TimestampsLogDatetimeShowTimezone.IsNull() {
		if value.Exists() {
			data.TimestampsLogDatetimeShowTimezone = types.BoolValue(true)
		} else {
			data.TimestampsLogDatetimeShowTimezone = types.BoolValue(false)
		}
	} else {
		data.TimestampsLogDatetimeShowTimezone = types.BoolNull()
	}
	if value := res.Get(prefix + "timestamps.log-config.datetime.year"); !data.TimestampsLogDatetimeYear.IsNull() {
		if value.Exists() {
			data.TimestampsLogDatetimeYear = types.BoolValue(true)
		} else {
			data.TimestampsLogDatetimeYear = types.BoolValue(false)
		}
	} else {
		data.TimestampsLogDatetimeYear = types.BoolNull()
	}
	if value := res.Get(prefix + "timestamps.log-config.uptime"); !data.TimestampsLogUptime.IsNull() {
		if value.Exists() {
			data.TimestampsLogUptime = types.BoolValue(true)
		} else {
			data.TimestampsLogUptime = types.BoolValue(false)
		}
	} else {
		data.TimestampsLogUptime = types.BoolNull()
	}
	if value := res.Get(prefix + "dhcp"); !data.Dhcp.IsNull() {
		if value.Exists() {
			data.Dhcp = types.BoolValue(true)
		} else {
			data.Dhcp = types.BoolValue(false)
		}
	} else {
		data.Dhcp = types.BoolNull()
	}
	if value := res.Get(prefix + "tcp-keepalives-in"); !data.TcpKeepalivesIn.IsNull() {
		if value.Exists() {
			data.TcpKeepalivesIn = types.BoolValue(true)
		} else {
			data.TcpKeepalivesIn = types.BoolValue(false)
		}
	} else {
		data.TcpKeepalivesIn = types.BoolNull()
	}
	if value := res.Get(prefix + "tcp-keepalives-out"); !data.TcpKeepalivesOut.IsNull() {
		if value.Exists() {
			data.TcpKeepalivesOut = types.BoolValue(true)
		} else {
			data.TcpKeepalivesOut = types.BoolValue(false)
		}
	} else {
		data.TcpKeepalivesOut = types.BoolNull()
	}
}

func (data *ServiceData) fromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "pad-conf.pad"); value.Exists() {
		data.Pad = types.BoolValue(value.Bool())
	} else {
		data.Pad = types.BoolValue(false)
	}
	if value := res.Get(prefix + "password-encryption"); value.Exists() {
		data.PasswordEncryption = types.BoolValue(true)
	} else {
		data.PasswordEncryption = types.BoolValue(false)
	}
	if value := res.Get(prefix + "password-recovery"); value.Exists() {
		data.PasswordRecovery = types.BoolValue(value.Bool())
	} else {
		data.PasswordRecovery = types.BoolValue(false)
	}
	if value := res.Get(prefix + "timestamps"); value.Exists() {
		data.Timestamps = types.BoolValue(true)
	} else {
		data.Timestamps = types.BoolValue(false)
	}
	if value := res.Get(prefix + "timestamps.debug-config"); value.Exists() {
		data.TimestampsDebug = types.BoolValue(true)
	} else {
		data.TimestampsDebug = types.BoolValue(false)
	}
	if value := res.Get(prefix + "timestamps.debug-config.datetime"); value.Exists() {
		data.TimestampsDebugDatetime = types.BoolValue(true)
	} else {
		data.TimestampsDebugDatetime = types.BoolValue(false)
	}
	if value := res.Get(prefix + "timestamps.debug-config.datetime.msec"); value.Exists() {
		data.TimestampsDebugDatetimeMsec = types.BoolValue(true)
	} else {
		data.TimestampsDebugDatetimeMsec = types.BoolValue(false)
	}
	if value := res.Get(prefix + "timestamps.debug-config.datetime.localtime"); value.Exists() {
		data.TimestampsDebugDatetimeLocaltime = types.BoolValue(true)
	} else {
		data.TimestampsDebugDatetimeLocaltime = types.BoolValue(false)
	}
	if value := res.Get(prefix + "timestamps.debug-config.datetime.show-timezone"); value.Exists() {
		data.TimestampsDebugDatetimeShowTimezone = types.BoolValue(true)
	} else {
		data.TimestampsDebugDatetimeShowTimezone = types.BoolValue(false)
	}
	if value := res.Get(prefix + "timestamps.debug-config.datetime.year"); value.Exists() {
		data.TimestampsDebugDatetimeYear = types.BoolValue(true)
	} else {
		data.TimestampsDebugDatetimeYear = types.BoolValue(false)
	}
	if value := res.Get(prefix + "timestamps.debug-config.uptime"); value.Exists() {
		data.TimestampsDebugUptime = types.BoolValue(true)
	} else {
		data.TimestampsDebugUptime = types.BoolValue(false)
	}
	if value := res.Get(prefix + "timestamps.log-config"); value.Exists() {
		data.TimestampsLog = types.BoolValue(true)
	} else {
		data.TimestampsLog = types.BoolValue(false)
	}
	if value := res.Get(prefix + "timestamps.log-config.datetime"); value.Exists() {
		data.TimestampsLogDatetime = types.BoolValue(true)
	} else {
		data.TimestampsLogDatetime = types.BoolValue(false)
	}
	if value := res.Get(prefix + "timestamps.log-config.datetime.msec"); value.Exists() {
		data.TimestampsLogDatetimeMsec = types.BoolValue(true)
	} else {
		data.TimestampsLogDatetimeMsec = types.BoolValue(false)
	}
	if value := res.Get(prefix + "timestamps.log-config.datetime.localtime"); value.Exists() {
		data.TimestampsLogDatetimeLocaltime = types.BoolValue(true)
	} else {
		data.TimestampsLogDatetimeLocaltime = types.BoolValue(false)
	}
	if value := res.Get(prefix + "timestamps.log-config.datetime.show-timezone"); value.Exists() {
		data.TimestampsLogDatetimeShowTimezone = types.BoolValue(true)
	} else {
		data.TimestampsLogDatetimeShowTimezone = types.BoolValue(false)
	}
	if value := res.Get(prefix + "timestamps.log-config.datetime.year"); value.Exists() {
		data.TimestampsLogDatetimeYear = types.BoolValue(true)
	} else {
		data.TimestampsLogDatetimeYear = types.BoolValue(false)
	}
	if value := res.Get(prefix + "timestamps.log-config.uptime"); value.Exists() {
		data.TimestampsLogUptime = types.BoolValue(true)
	} else {
		data.TimestampsLogUptime = types.BoolValue(false)
	}
	if value := res.Get(prefix + "dhcp"); value.Exists() {
		data.Dhcp = types.BoolValue(true)
	} else {
		data.Dhcp = types.BoolValue(false)
	}
	if value := res.Get(prefix + "tcp-keepalives-in"); value.Exists() {
		data.TcpKeepalivesIn = types.BoolValue(true)
	} else {
		data.TcpKeepalivesIn = types.BoolValue(false)
	}
	if value := res.Get(prefix + "tcp-keepalives-out"); value.Exists() {
		data.TcpKeepalivesOut = types.BoolValue(true)
	} else {
		data.TcpKeepalivesOut = types.BoolValue(false)
	}
}

func (data *Service) getDeletedListItems(ctx context.Context, state Service) []string {
	deletedListItems := make([]string, 0)
	return deletedListItems
}

func (data *Service) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	if !data.PasswordEncryption.IsNull() && !data.PasswordEncryption.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/password-encryption", data.getPath()))
	}
	if !data.Timestamps.IsNull() && !data.Timestamps.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/timestamps", data.getPath()))
	}
	if !data.TimestampsDebug.IsNull() && !data.TimestampsDebug.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/timestamps/debug-config", data.getPath()))
	}
	if !data.TimestampsDebugDatetime.IsNull() && !data.TimestampsDebugDatetime.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/timestamps/debug-config/datetime", data.getPath()))
	}
	if !data.TimestampsDebugDatetimeMsec.IsNull() && !data.TimestampsDebugDatetimeMsec.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/timestamps/debug-config/datetime/msec", data.getPath()))
	}
	if !data.TimestampsDebugDatetimeLocaltime.IsNull() && !data.TimestampsDebugDatetimeLocaltime.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/timestamps/debug-config/datetime/localtime", data.getPath()))
	}
	if !data.TimestampsDebugDatetimeShowTimezone.IsNull() && !data.TimestampsDebugDatetimeShowTimezone.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/timestamps/debug-config/datetime/show-timezone", data.getPath()))
	}
	if !data.TimestampsDebugDatetimeYear.IsNull() && !data.TimestampsDebugDatetimeYear.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/timestamps/debug-config/datetime/year", data.getPath()))
	}
	if !data.TimestampsDebugUptime.IsNull() && !data.TimestampsDebugUptime.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/timestamps/debug-config/uptime", data.getPath()))
	}
	if !data.TimestampsLog.IsNull() && !data.TimestampsLog.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/timestamps/log-config", data.getPath()))
	}
	if !data.TimestampsLogDatetime.IsNull() && !data.TimestampsLogDatetime.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/timestamps/log-config/datetime", data.getPath()))
	}
	if !data.TimestampsLogDatetimeMsec.IsNull() && !data.TimestampsLogDatetimeMsec.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/timestamps/log-config/datetime/msec", data.getPath()))
	}
	if !data.TimestampsLogDatetimeLocaltime.IsNull() && !data.TimestampsLogDatetimeLocaltime.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/timestamps/log-config/datetime/localtime", data.getPath()))
	}
	if !data.TimestampsLogDatetimeShowTimezone.IsNull() && !data.TimestampsLogDatetimeShowTimezone.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/timestamps/log-config/datetime/show-timezone", data.getPath()))
	}
	if !data.TimestampsLogDatetimeYear.IsNull() && !data.TimestampsLogDatetimeYear.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/timestamps/log-config/datetime/year", data.getPath()))
	}
	if !data.TimestampsLogUptime.IsNull() && !data.TimestampsLogUptime.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/timestamps/log-config/uptime", data.getPath()))
	}
	if !data.Dhcp.IsNull() && !data.Dhcp.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/dhcp", data.getPath()))
	}
	if !data.TcpKeepalivesIn.IsNull() && !data.TcpKeepalivesIn.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/tcp-keepalives-in", data.getPath()))
	}
	if !data.TcpKeepalivesOut.IsNull() && !data.TcpKeepalivesOut.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/tcp-keepalives-out", data.getPath()))
	}
	return emptyLeafsDelete
}

func (data *Service) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	if !data.Pad.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/pad-conf/pad", data.getPath()))
	}
	if !data.PasswordEncryption.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/password-encryption", data.getPath()))
	}
	if !data.Timestamps.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/timestamps", data.getPath()))
	}
	if !data.TimestampsDebug.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/timestamps/debug-config", data.getPath()))
	}
	if !data.TimestampsDebugDatetime.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/timestamps/debug-config/datetime", data.getPath()))
	}
	if !data.TimestampsDebugDatetimeMsec.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/timestamps/debug-config/datetime/msec", data.getPath()))
	}
	if !data.TimestampsDebugDatetimeLocaltime.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/timestamps/debug-config/datetime/localtime", data.getPath()))
	}
	if !data.TimestampsDebugDatetimeShowTimezone.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/timestamps/debug-config/datetime/show-timezone", data.getPath()))
	}
	if !data.TimestampsDebugDatetimeYear.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/timestamps/debug-config/datetime/year", data.getPath()))
	}
	if !data.TimestampsDebugUptime.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/timestamps/debug-config/uptime", data.getPath()))
	}
	if !data.TimestampsLog.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/timestamps/log-config", data.getPath()))
	}
	if !data.TimestampsLogDatetime.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/timestamps/log-config/datetime", data.getPath()))
	}
	if !data.TimestampsLogDatetimeMsec.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/timestamps/log-config/datetime/msec", data.getPath()))
	}
	if !data.TimestampsLogDatetimeLocaltime.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/timestamps/log-config/datetime/localtime", data.getPath()))
	}
	if !data.TimestampsLogDatetimeShowTimezone.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/timestamps/log-config/datetime/show-timezone", data.getPath()))
	}
	if !data.TimestampsLogDatetimeYear.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/timestamps/log-config/datetime/year", data.getPath()))
	}
	if !data.TimestampsLogUptime.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/timestamps/log-config/uptime", data.getPath()))
	}
	if !data.Dhcp.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/dhcp", data.getPath()))
	}
	if !data.TcpKeepalivesIn.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/tcp-keepalives-in", data.getPath()))
	}
	if !data.TcpKeepalivesOut.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/tcp-keepalives-out", data.getPath()))
	}
	return deletePaths
}
