// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
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

func (data Service) toBody() string {
	body := `{"` + helpers.LastElement(data.getPath()) + `":{}}`
	if !data.Pad.Null && !data.Pad.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"pad-conf.pad", data.Pad.Value)
	}
	if !data.PasswordEncryption.Null && !data.PasswordEncryption.Unknown {
		if data.PasswordEncryption.Value {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"password-encryption", map[string]string{})
		}
	}
	if !data.PasswordRecovery.Null && !data.PasswordRecovery.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"password-recovery", data.PasswordRecovery.Value)
	}
	if !data.Timestamps.Null && !data.Timestamps.Unknown {
		if data.Timestamps.Value {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"timestamps", map[string]string{})
		}
	}
	if !data.TimestampsDebug.Null && !data.TimestampsDebug.Unknown {
		if data.TimestampsDebug.Value {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"timestamps.debug-config", map[string]string{})
		}
	}
	if !data.TimestampsDebugDatetime.Null && !data.TimestampsDebugDatetime.Unknown {
		if data.TimestampsDebugDatetime.Value {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"timestamps.debug-config.datetime", map[string]string{})
		}
	}
	if !data.TimestampsDebugDatetimeMsec.Null && !data.TimestampsDebugDatetimeMsec.Unknown {
		if data.TimestampsDebugDatetimeMsec.Value {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"timestamps.debug-config.datetime.msec", map[string]string{})
		}
	}
	if !data.TimestampsDebugDatetimeLocaltime.Null && !data.TimestampsDebugDatetimeLocaltime.Unknown {
		if data.TimestampsDebugDatetimeLocaltime.Value {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"timestamps.debug-config.datetime.localtime", map[string]string{})
		}
	}
	if !data.TimestampsDebugDatetimeShowTimezone.Null && !data.TimestampsDebugDatetimeShowTimezone.Unknown {
		if data.TimestampsDebugDatetimeShowTimezone.Value {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"timestamps.debug-config.datetime.show-timezone", map[string]string{})
		}
	}
	if !data.TimestampsDebugDatetimeYear.Null && !data.TimestampsDebugDatetimeYear.Unknown {
		if data.TimestampsDebugDatetimeYear.Value {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"timestamps.debug-config.datetime.year", map[string]string{})
		}
	}
	if !data.TimestampsDebugUptime.Null && !data.TimestampsDebugUptime.Unknown {
		if data.TimestampsDebugUptime.Value {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"timestamps.debug-config.uptime", map[string]string{})
		}
	}
	if !data.TimestampsLog.Null && !data.TimestampsLog.Unknown {
		if data.TimestampsLog.Value {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"timestamps.log-config", map[string]string{})
		}
	}
	if !data.TimestampsLogDatetime.Null && !data.TimestampsLogDatetime.Unknown {
		if data.TimestampsLogDatetime.Value {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"timestamps.log-config.datetime", map[string]string{})
		}
	}
	if !data.TimestampsLogDatetimeMsec.Null && !data.TimestampsLogDatetimeMsec.Unknown {
		if data.TimestampsLogDatetimeMsec.Value {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"timestamps.log-config.datetime.msec", map[string]string{})
		}
	}
	if !data.TimestampsLogDatetimeLocaltime.Null && !data.TimestampsLogDatetimeLocaltime.Unknown {
		if data.TimestampsLogDatetimeLocaltime.Value {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"timestamps.log-config.datetime.localtime", map[string]string{})
		}
	}
	if !data.TimestampsLogDatetimeShowTimezone.Null && !data.TimestampsLogDatetimeShowTimezone.Unknown {
		if data.TimestampsLogDatetimeShowTimezone.Value {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"timestamps.log-config.datetime.show-timezone", map[string]string{})
		}
	}
	if !data.TimestampsLogDatetimeYear.Null && !data.TimestampsLogDatetimeYear.Unknown {
		if data.TimestampsLogDatetimeYear.Value {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"timestamps.log-config.datetime.year", map[string]string{})
		}
	}
	if !data.TimestampsLogUptime.Null && !data.TimestampsLogUptime.Unknown {
		if data.TimestampsLogUptime.Value {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"timestamps.log-config.uptime", map[string]string{})
		}
	}
	if !data.Dhcp.Null && !data.Dhcp.Unknown {
		if data.Dhcp.Value {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"dhcp", map[string]string{})
		}
	}
	if !data.TcpKeepalivesIn.Null && !data.TcpKeepalivesIn.Unknown {
		if data.TcpKeepalivesIn.Value {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"tcp-keepalives-in", map[string]string{})
		}
	}
	if !data.TcpKeepalivesOut.Null && !data.TcpKeepalivesOut.Unknown {
		if data.TcpKeepalivesOut.Value {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"tcp-keepalives-out", map[string]string{})
		}
	}
	return body
}

func (data *Service) updateFromBody(res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "pad-conf.pad"); value.Exists() {
		data.Pad.Value = value.Bool()
	} else {
		data.Pad.Value = false
	}
	if value := res.Get(prefix + "password-encryption"); value.Exists() {
		data.PasswordEncryption.Value = true
	} else {
		data.PasswordEncryption.Value = false
	}
	if value := res.Get(prefix + "password-recovery"); value.Exists() {
		data.PasswordRecovery.Value = value.Bool()
	} else {
		data.PasswordRecovery.Value = false
	}
	if value := res.Get(prefix + "timestamps"); value.Exists() {
		data.Timestamps.Value = true
	} else {
		data.Timestamps.Value = false
	}
	if value := res.Get(prefix + "timestamps.debug-config"); value.Exists() {
		data.TimestampsDebug.Value = true
	} else {
		data.TimestampsDebug.Value = false
	}
	if value := res.Get(prefix + "timestamps.debug-config.datetime"); value.Exists() {
		data.TimestampsDebugDatetime.Value = true
	} else {
		data.TimestampsDebugDatetime.Value = false
	}
	if value := res.Get(prefix + "timestamps.debug-config.datetime.msec"); value.Exists() {
		data.TimestampsDebugDatetimeMsec.Value = true
	} else {
		data.TimestampsDebugDatetimeMsec.Value = false
	}
	if value := res.Get(prefix + "timestamps.debug-config.datetime.localtime"); value.Exists() {
		data.TimestampsDebugDatetimeLocaltime.Value = true
	} else {
		data.TimestampsDebugDatetimeLocaltime.Value = false
	}
	if value := res.Get(prefix + "timestamps.debug-config.datetime.show-timezone"); value.Exists() {
		data.TimestampsDebugDatetimeShowTimezone.Value = true
	} else {
		data.TimestampsDebugDatetimeShowTimezone.Value = false
	}
	if value := res.Get(prefix + "timestamps.debug-config.datetime.year"); value.Exists() {
		data.TimestampsDebugDatetimeYear.Value = true
	} else {
		data.TimestampsDebugDatetimeYear.Value = false
	}
	if value := res.Get(prefix + "timestamps.debug-config.uptime"); value.Exists() {
		data.TimestampsDebugUptime.Value = true
	} else {
		data.TimestampsDebugUptime.Value = false
	}
	if value := res.Get(prefix + "timestamps.log-config"); value.Exists() {
		data.TimestampsLog.Value = true
	} else {
		data.TimestampsLog.Value = false
	}
	if value := res.Get(prefix + "timestamps.log-config.datetime"); value.Exists() {
		data.TimestampsLogDatetime.Value = true
	} else {
		data.TimestampsLogDatetime.Value = false
	}
	if value := res.Get(prefix + "timestamps.log-config.datetime.msec"); value.Exists() {
		data.TimestampsLogDatetimeMsec.Value = true
	} else {
		data.TimestampsLogDatetimeMsec.Value = false
	}
	if value := res.Get(prefix + "timestamps.log-config.datetime.localtime"); value.Exists() {
		data.TimestampsLogDatetimeLocaltime.Value = true
	} else {
		data.TimestampsLogDatetimeLocaltime.Value = false
	}
	if value := res.Get(prefix + "timestamps.log-config.datetime.show-timezone"); value.Exists() {
		data.TimestampsLogDatetimeShowTimezone.Value = true
	} else {
		data.TimestampsLogDatetimeShowTimezone.Value = false
	}
	if value := res.Get(prefix + "timestamps.log-config.datetime.year"); value.Exists() {
		data.TimestampsLogDatetimeYear.Value = true
	} else {
		data.TimestampsLogDatetimeYear.Value = false
	}
	if value := res.Get(prefix + "timestamps.log-config.uptime"); value.Exists() {
		data.TimestampsLogUptime.Value = true
	} else {
		data.TimestampsLogUptime.Value = false
	}
	if value := res.Get(prefix + "dhcp"); value.Exists() {
		data.Dhcp.Value = true
	} else {
		data.Dhcp.Value = false
	}
	if value := res.Get(prefix + "tcp-keepalives-in"); value.Exists() {
		data.TcpKeepalivesIn.Value = true
	} else {
		data.TcpKeepalivesIn.Value = false
	}
	if value := res.Get(prefix + "tcp-keepalives-out"); value.Exists() {
		data.TcpKeepalivesOut.Value = true
	} else {
		data.TcpKeepalivesOut.Value = false
	}
}

func (data *Service) fromBody(res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "pad-conf.pad"); value.Exists() {
		data.Pad.Value = value.Bool()
		data.Pad.Null = false
	} else {
		data.Pad.Value = false
		data.Pad.Null = false
	}
	if value := res.Get(prefix + "password-encryption"); value.Exists() {
		data.PasswordEncryption.Value = true
		data.PasswordEncryption.Null = false
	} else {
		data.PasswordEncryption.Value = false
		data.PasswordEncryption.Null = false
	}
	if value := res.Get(prefix + "password-recovery"); value.Exists() {
		data.PasswordRecovery.Value = value.Bool()
		data.PasswordRecovery.Null = false
	} else {
		data.PasswordRecovery.Value = false
		data.PasswordRecovery.Null = false
	}
	if value := res.Get(prefix + "timestamps"); value.Exists() {
		data.Timestamps.Value = true
		data.Timestamps.Null = false
	} else {
		data.Timestamps.Value = false
		data.Timestamps.Null = false
	}
	if value := res.Get(prefix + "timestamps.debug-config"); value.Exists() {
		data.TimestampsDebug.Value = true
		data.TimestampsDebug.Null = false
	} else {
		data.TimestampsDebug.Value = false
		data.TimestampsDebug.Null = false
	}
	if value := res.Get(prefix + "timestamps.debug-config.datetime"); value.Exists() {
		data.TimestampsDebugDatetime.Value = true
		data.TimestampsDebugDatetime.Null = false
	} else {
		data.TimestampsDebugDatetime.Value = false
		data.TimestampsDebugDatetime.Null = false
	}
	if value := res.Get(prefix + "timestamps.debug-config.datetime.msec"); value.Exists() {
		data.TimestampsDebugDatetimeMsec.Value = true
		data.TimestampsDebugDatetimeMsec.Null = false
	} else {
		data.TimestampsDebugDatetimeMsec.Value = false
		data.TimestampsDebugDatetimeMsec.Null = false
	}
	if value := res.Get(prefix + "timestamps.debug-config.datetime.localtime"); value.Exists() {
		data.TimestampsDebugDatetimeLocaltime.Value = true
		data.TimestampsDebugDatetimeLocaltime.Null = false
	} else {
		data.TimestampsDebugDatetimeLocaltime.Value = false
		data.TimestampsDebugDatetimeLocaltime.Null = false
	}
	if value := res.Get(prefix + "timestamps.debug-config.datetime.show-timezone"); value.Exists() {
		data.TimestampsDebugDatetimeShowTimezone.Value = true
		data.TimestampsDebugDatetimeShowTimezone.Null = false
	} else {
		data.TimestampsDebugDatetimeShowTimezone.Value = false
		data.TimestampsDebugDatetimeShowTimezone.Null = false
	}
	if value := res.Get(prefix + "timestamps.debug-config.datetime.year"); value.Exists() {
		data.TimestampsDebugDatetimeYear.Value = true
		data.TimestampsDebugDatetimeYear.Null = false
	} else {
		data.TimestampsDebugDatetimeYear.Value = false
		data.TimestampsDebugDatetimeYear.Null = false
	}
	if value := res.Get(prefix + "timestamps.debug-config.uptime"); value.Exists() {
		data.TimestampsDebugUptime.Value = true
		data.TimestampsDebugUptime.Null = false
	} else {
		data.TimestampsDebugUptime.Value = false
		data.TimestampsDebugUptime.Null = false
	}
	if value := res.Get(prefix + "timestamps.log-config"); value.Exists() {
		data.TimestampsLog.Value = true
		data.TimestampsLog.Null = false
	} else {
		data.TimestampsLog.Value = false
		data.TimestampsLog.Null = false
	}
	if value := res.Get(prefix + "timestamps.log-config.datetime"); value.Exists() {
		data.TimestampsLogDatetime.Value = true
		data.TimestampsLogDatetime.Null = false
	} else {
		data.TimestampsLogDatetime.Value = false
		data.TimestampsLogDatetime.Null = false
	}
	if value := res.Get(prefix + "timestamps.log-config.datetime.msec"); value.Exists() {
		data.TimestampsLogDatetimeMsec.Value = true
		data.TimestampsLogDatetimeMsec.Null = false
	} else {
		data.TimestampsLogDatetimeMsec.Value = false
		data.TimestampsLogDatetimeMsec.Null = false
	}
	if value := res.Get(prefix + "timestamps.log-config.datetime.localtime"); value.Exists() {
		data.TimestampsLogDatetimeLocaltime.Value = true
		data.TimestampsLogDatetimeLocaltime.Null = false
	} else {
		data.TimestampsLogDatetimeLocaltime.Value = false
		data.TimestampsLogDatetimeLocaltime.Null = false
	}
	if value := res.Get(prefix + "timestamps.log-config.datetime.show-timezone"); value.Exists() {
		data.TimestampsLogDatetimeShowTimezone.Value = true
		data.TimestampsLogDatetimeShowTimezone.Null = false
	} else {
		data.TimestampsLogDatetimeShowTimezone.Value = false
		data.TimestampsLogDatetimeShowTimezone.Null = false
	}
	if value := res.Get(prefix + "timestamps.log-config.datetime.year"); value.Exists() {
		data.TimestampsLogDatetimeYear.Value = true
		data.TimestampsLogDatetimeYear.Null = false
	} else {
		data.TimestampsLogDatetimeYear.Value = false
		data.TimestampsLogDatetimeYear.Null = false
	}
	if value := res.Get(prefix + "timestamps.log-config.uptime"); value.Exists() {
		data.TimestampsLogUptime.Value = true
		data.TimestampsLogUptime.Null = false
	} else {
		data.TimestampsLogUptime.Value = false
		data.TimestampsLogUptime.Null = false
	}
	if value := res.Get(prefix + "dhcp"); value.Exists() {
		data.Dhcp.Value = true
		data.Dhcp.Null = false
	} else {
		data.Dhcp.Value = false
		data.Dhcp.Null = false
	}
	if value := res.Get(prefix + "tcp-keepalives-in"); value.Exists() {
		data.TcpKeepalivesIn.Value = true
		data.TcpKeepalivesIn.Null = false
	} else {
		data.TcpKeepalivesIn.Value = false
		data.TcpKeepalivesIn.Null = false
	}
	if value := res.Get(prefix + "tcp-keepalives-out"); value.Exists() {
		data.TcpKeepalivesOut.Value = true
		data.TcpKeepalivesOut.Null = false
	} else {
		data.TcpKeepalivesOut.Value = false
		data.TcpKeepalivesOut.Null = false
	}
}

func (data *Service) setUnknownValues() {
	if data.Device.Unknown {
		data.Device.Unknown = false
		data.Device.Null = true
	}
	if data.Id.Unknown {
		data.Id.Unknown = false
		data.Id.Null = true
	}
	if data.Pad.Unknown {
		data.Pad.Unknown = false
		data.Pad.Null = true
	}
	if data.PasswordEncryption.Unknown {
		data.PasswordEncryption.Unknown = false
		data.PasswordEncryption.Null = true
	}
	if data.PasswordRecovery.Unknown {
		data.PasswordRecovery.Unknown = false
		data.PasswordRecovery.Null = true
	}
	if data.Timestamps.Unknown {
		data.Timestamps.Unknown = false
		data.Timestamps.Null = true
	}
	if data.TimestampsDebug.Unknown {
		data.TimestampsDebug.Unknown = false
		data.TimestampsDebug.Null = true
	}
	if data.TimestampsDebugDatetime.Unknown {
		data.TimestampsDebugDatetime.Unknown = false
		data.TimestampsDebugDatetime.Null = true
	}
	if data.TimestampsDebugDatetimeMsec.Unknown {
		data.TimestampsDebugDatetimeMsec.Unknown = false
		data.TimestampsDebugDatetimeMsec.Null = true
	}
	if data.TimestampsDebugDatetimeLocaltime.Unknown {
		data.TimestampsDebugDatetimeLocaltime.Unknown = false
		data.TimestampsDebugDatetimeLocaltime.Null = true
	}
	if data.TimestampsDebugDatetimeShowTimezone.Unknown {
		data.TimestampsDebugDatetimeShowTimezone.Unknown = false
		data.TimestampsDebugDatetimeShowTimezone.Null = true
	}
	if data.TimestampsDebugDatetimeYear.Unknown {
		data.TimestampsDebugDatetimeYear.Unknown = false
		data.TimestampsDebugDatetimeYear.Null = true
	}
	if data.TimestampsDebugUptime.Unknown {
		data.TimestampsDebugUptime.Unknown = false
		data.TimestampsDebugUptime.Null = true
	}
	if data.TimestampsLog.Unknown {
		data.TimestampsLog.Unknown = false
		data.TimestampsLog.Null = true
	}
	if data.TimestampsLogDatetime.Unknown {
		data.TimestampsLogDatetime.Unknown = false
		data.TimestampsLogDatetime.Null = true
	}
	if data.TimestampsLogDatetimeMsec.Unknown {
		data.TimestampsLogDatetimeMsec.Unknown = false
		data.TimestampsLogDatetimeMsec.Null = true
	}
	if data.TimestampsLogDatetimeLocaltime.Unknown {
		data.TimestampsLogDatetimeLocaltime.Unknown = false
		data.TimestampsLogDatetimeLocaltime.Null = true
	}
	if data.TimestampsLogDatetimeShowTimezone.Unknown {
		data.TimestampsLogDatetimeShowTimezone.Unknown = false
		data.TimestampsLogDatetimeShowTimezone.Null = true
	}
	if data.TimestampsLogDatetimeYear.Unknown {
		data.TimestampsLogDatetimeYear.Unknown = false
		data.TimestampsLogDatetimeYear.Null = true
	}
	if data.TimestampsLogUptime.Unknown {
		data.TimestampsLogUptime.Unknown = false
		data.TimestampsLogUptime.Null = true
	}
	if data.Dhcp.Unknown {
		data.Dhcp.Unknown = false
		data.Dhcp.Null = true
	}
	if data.TcpKeepalivesIn.Unknown {
		data.TcpKeepalivesIn.Unknown = false
		data.TcpKeepalivesIn.Null = true
	}
	if data.TcpKeepalivesOut.Unknown {
		data.TcpKeepalivesOut.Unknown = false
		data.TcpKeepalivesOut.Null = true
	}
}

func (data *Service) getDeletedListItems(state Service) []string {
	deletedListItems := make([]string, 0)
	return deletedListItems
}

func (data *Service) getEmptyLeafsDelete() []string {
	emptyLeafsDelete := make([]string, 0)
	if !data.PasswordEncryption.Value {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/password-encryption", data.getPath()))
	}
	if !data.TimestampsDebugDatetimeMsec.Value {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/timestamps/debug-config/datetime/msec", data.getPath()))
	}
	if !data.TimestampsDebugDatetimeLocaltime.Value {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/timestamps/debug-config/datetime/localtime", data.getPath()))
	}
	if !data.TimestampsDebugDatetimeShowTimezone.Value {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/timestamps/debug-config/datetime/show-timezone", data.getPath()))
	}
	if !data.TimestampsDebugDatetimeYear.Value {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/timestamps/debug-config/datetime/year", data.getPath()))
	}
	if !data.TimestampsDebugUptime.Value {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/timestamps/debug-config/uptime", data.getPath()))
	}
	if !data.TimestampsLogDatetimeMsec.Value {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/timestamps/log-config/datetime/msec", data.getPath()))
	}
	if !data.TimestampsLogDatetimeLocaltime.Value {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/timestamps/log-config/datetime/localtime", data.getPath()))
	}
	if !data.TimestampsLogDatetimeShowTimezone.Value {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/timestamps/log-config/datetime/show-timezone", data.getPath()))
	}
	if !data.TimestampsLogDatetimeYear.Value {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/timestamps/log-config/datetime/year", data.getPath()))
	}
	if !data.TimestampsLogUptime.Value {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/timestamps/log-config/uptime", data.getPath()))
	}
	if !data.Dhcp.Value {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/dhcp", data.getPath()))
	}
	if !data.TcpKeepalivesIn.Value {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/tcp-keepalives-in", data.getPath()))
	}
	if !data.TcpKeepalivesOut.Value {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/tcp-keepalives-out", data.getPath()))
	}
	return emptyLeafsDelete
}
