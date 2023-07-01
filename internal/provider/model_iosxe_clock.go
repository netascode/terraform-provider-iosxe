// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"regexp"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/terraform-provider-iosxe/internal/provider/helpers"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type Clock struct {
	Device                          types.String `tfsdk:"device"`
	Id                              types.String `tfsdk:"id"`
	DeleteMode                      types.String `tfsdk:"delete_mode"`
	CalendarValid                   types.Bool   `tfsdk:"calendar_valid"`
	SummerTimeZone                  types.String `tfsdk:"summer_time_zone"`
	SummerTimeDate                  types.Bool   `tfsdk:"summer_time_date"`
	SummerTimeDateStartDay          types.Int64  `tfsdk:"summer_time_date_start_day"`
	SummerTimeDateStartMonth        types.String `tfsdk:"summer_time_date_start_month"`
	SummerTimeDateStartYear         types.Int64  `tfsdk:"summer_time_date_start_year"`
	SummerTimeDateStartTime         types.String `tfsdk:"summer_time_date_start_time"`
	SummerTimeDateEndDay            types.Int64  `tfsdk:"summer_time_date_end_day"`
	SummerTimeDateEndMonth          types.String `tfsdk:"summer_time_date_end_month"`
	SummerTimeDateEndYear           types.Int64  `tfsdk:"summer_time_date_end_year"`
	SummerTimeDateEndTime           types.String `tfsdk:"summer_time_date_end_time"`
	SummerTimeDateOffset            types.Int64  `tfsdk:"summer_time_date_offset"`
	SummerTimeRecurring             types.Bool   `tfsdk:"summer_time_recurring"`
	SummerTimeRecurringStartWeek    types.String `tfsdk:"summer_time_recurring_start_week"`
	SummerTimeRecurringStartWeekday types.String `tfsdk:"summer_time_recurring_start_weekday"`
	SummerTimeRecurringStartMonth   types.String `tfsdk:"summer_time_recurring_start_month"`
	SummerTimeRecurringStartTime    types.String `tfsdk:"summer_time_recurring_start_time"`
	SummerTimeRecurringEndWeek      types.String `tfsdk:"summer_time_recurring_end_week"`
	SummerTimeRecurringEndWeekday   types.String `tfsdk:"summer_time_recurring_end_weekday"`
	SummerTimeRecurringEndMonth     types.String `tfsdk:"summer_time_recurring_end_month"`
	SummerTimeRecurringEndTime      types.String `tfsdk:"summer_time_recurring_end_time"`
	SummerTimeRecurringOffset       types.Int64  `tfsdk:"summer_time_recurring_offset"`
}

type ClockData struct {
	Device                          types.String `tfsdk:"device"`
	Id                              types.String `tfsdk:"id"`
	CalendarValid                   types.Bool   `tfsdk:"calendar_valid"`
	SummerTimeZone                  types.String `tfsdk:"summer_time_zone"`
	SummerTimeDate                  types.Bool   `tfsdk:"summer_time_date"`
	SummerTimeDateStartDay          types.Int64  `tfsdk:"summer_time_date_start_day"`
	SummerTimeDateStartMonth        types.String `tfsdk:"summer_time_date_start_month"`
	SummerTimeDateStartYear         types.Int64  `tfsdk:"summer_time_date_start_year"`
	SummerTimeDateStartTime         types.String `tfsdk:"summer_time_date_start_time"`
	SummerTimeDateEndDay            types.Int64  `tfsdk:"summer_time_date_end_day"`
	SummerTimeDateEndMonth          types.String `tfsdk:"summer_time_date_end_month"`
	SummerTimeDateEndYear           types.Int64  `tfsdk:"summer_time_date_end_year"`
	SummerTimeDateEndTime           types.String `tfsdk:"summer_time_date_end_time"`
	SummerTimeDateOffset            types.Int64  `tfsdk:"summer_time_date_offset"`
	SummerTimeRecurring             types.Bool   `tfsdk:"summer_time_recurring"`
	SummerTimeRecurringStartWeek    types.String `tfsdk:"summer_time_recurring_start_week"`
	SummerTimeRecurringStartWeekday types.String `tfsdk:"summer_time_recurring_start_weekday"`
	SummerTimeRecurringStartMonth   types.String `tfsdk:"summer_time_recurring_start_month"`
	SummerTimeRecurringStartTime    types.String `tfsdk:"summer_time_recurring_start_time"`
	SummerTimeRecurringEndWeek      types.String `tfsdk:"summer_time_recurring_end_week"`
	SummerTimeRecurringEndWeekday   types.String `tfsdk:"summer_time_recurring_end_weekday"`
	SummerTimeRecurringEndMonth     types.String `tfsdk:"summer_time_recurring_end_month"`
	SummerTimeRecurringEndTime      types.String `tfsdk:"summer_time_recurring_end_time"`
	SummerTimeRecurringOffset       types.Int64  `tfsdk:"summer_time_recurring_offset"`
}

func (data Clock) getPath() string {
	return "Cisco-IOS-XE-native:native/clock"
}

func (data ClockData) getPath() string {
	return "Cisco-IOS-XE-native:native/clock"
}

// if last path element has a key -> remove it
func (data Clock) getPathShort() string {
	path := data.getPath()
	re := regexp.MustCompile(`(.*)=[^\/]*$`)
	matches := re.FindStringSubmatch(path)
	if len(matches) <= 1 {
		return path
	}
	return matches[1]
}

func (data Clock) toBody(ctx context.Context) string {
	body := `{"` + helpers.LastElement(data.getPath()) + `":{}}`
	if !data.CalendarValid.IsNull() && !data.CalendarValid.IsUnknown() {
		if data.CalendarValid.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"calendar-valid", map[string]string{})
		}
	}
	if !data.SummerTimeZone.IsNull() && !data.SummerTimeZone.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"summer-time.zone", data.SummerTimeZone.ValueString())
	}
	if !data.SummerTimeDate.IsNull() && !data.SummerTimeDate.IsUnknown() {
		if data.SummerTimeDate.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"summer-time.date", map[string]string{})
		}
	}
	if !data.SummerTimeDateStartDay.IsNull() && !data.SummerTimeDateStartDay.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"summer-time.start-day", strconv.FormatInt(data.SummerTimeDateStartDay.ValueInt64(), 10))
	}
	if !data.SummerTimeDateStartMonth.IsNull() && !data.SummerTimeDateStartMonth.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"summer-time.start-month", data.SummerTimeDateStartMonth.ValueString())
	}
	if !data.SummerTimeDateStartYear.IsNull() && !data.SummerTimeDateStartYear.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"summer-time.start-year", strconv.FormatInt(data.SummerTimeDateStartYear.ValueInt64(), 10))
	}
	if !data.SummerTimeDateStartTime.IsNull() && !data.SummerTimeDateStartTime.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"summer-time.start-time", data.SummerTimeDateStartTime.ValueString())
	}
	if !data.SummerTimeDateEndDay.IsNull() && !data.SummerTimeDateEndDay.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"summer-time.date-end-day", strconv.FormatInt(data.SummerTimeDateEndDay.ValueInt64(), 10))
	}
	if !data.SummerTimeDateEndMonth.IsNull() && !data.SummerTimeDateEndMonth.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"summer-time.date-end-month", data.SummerTimeDateEndMonth.ValueString())
	}
	if !data.SummerTimeDateEndYear.IsNull() && !data.SummerTimeDateEndYear.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"summer-time.date-end-year", strconv.FormatInt(data.SummerTimeDateEndYear.ValueInt64(), 10))
	}
	if !data.SummerTimeDateEndTime.IsNull() && !data.SummerTimeDateEndTime.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"summer-time.date-end-time", data.SummerTimeDateEndTime.ValueString())
	}
	if !data.SummerTimeDateOffset.IsNull() && !data.SummerTimeDateOffset.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"summer-time.offset", strconv.FormatInt(data.SummerTimeDateOffset.ValueInt64(), 10))
	}
	if !data.SummerTimeRecurring.IsNull() && !data.SummerTimeRecurring.IsUnknown() {
		if data.SummerTimeRecurring.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"summer-time.recurring", map[string]string{})
		}
	}
	if !data.SummerTimeRecurringStartWeek.IsNull() && !data.SummerTimeRecurringStartWeek.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"summer-time.recurring-start", data.SummerTimeRecurringStartWeek.ValueString())
	}
	if !data.SummerTimeRecurringStartWeekday.IsNull() && !data.SummerTimeRecurringStartWeekday.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"summer-time.recurring-start-day", data.SummerTimeRecurringStartWeekday.ValueString())
	}
	if !data.SummerTimeRecurringStartMonth.IsNull() && !data.SummerTimeRecurringStartMonth.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"summer-time.recurring-start-month", data.SummerTimeRecurringStartMonth.ValueString())
	}
	if !data.SummerTimeRecurringStartTime.IsNull() && !data.SummerTimeRecurringStartTime.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"summer-time.recurring-start-time", data.SummerTimeRecurringStartTime.ValueString())
	}
	if !data.SummerTimeRecurringEndWeek.IsNull() && !data.SummerTimeRecurringEndWeek.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"summer-time.recurring-end", data.SummerTimeRecurringEndWeek.ValueString())
	}
	if !data.SummerTimeRecurringEndWeekday.IsNull() && !data.SummerTimeRecurringEndWeekday.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"summer-time.recurring-end-day", data.SummerTimeRecurringEndWeekday.ValueString())
	}
	if !data.SummerTimeRecurringEndMonth.IsNull() && !data.SummerTimeRecurringEndMonth.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"summer-time.recurring-end-month", data.SummerTimeRecurringEndMonth.ValueString())
	}
	if !data.SummerTimeRecurringEndTime.IsNull() && !data.SummerTimeRecurringEndTime.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"summer-time.recurring-end-time", data.SummerTimeRecurringEndTime.ValueString())
	}
	if !data.SummerTimeRecurringOffset.IsNull() && !data.SummerTimeRecurringOffset.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"summer-time.recurring-offset", strconv.FormatInt(data.SummerTimeRecurringOffset.ValueInt64(), 10))
	}
	return body
}

func (data *Clock) updateFromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "calendar-valid"); !data.CalendarValid.IsNull() {
		if value.Exists() {
			data.CalendarValid = types.BoolValue(true)
		} else {
			data.CalendarValid = types.BoolValue(false)
		}
	} else {
		data.CalendarValid = types.BoolNull()
	}
	if value := res.Get(prefix + "summer-time.zone"); value.Exists() && !data.SummerTimeZone.IsNull() {
		data.SummerTimeZone = types.StringValue(value.String())
	} else {
		data.SummerTimeZone = types.StringNull()
	}
	if value := res.Get(prefix + "summer-time.date"); !data.SummerTimeDate.IsNull() {
		if value.Exists() {
			data.SummerTimeDate = types.BoolValue(true)
		} else {
			data.SummerTimeDate = types.BoolValue(false)
		}
	} else {
		data.SummerTimeDate = types.BoolNull()
	}
	if value := res.Get(prefix + "summer-time.start-day"); value.Exists() && !data.SummerTimeDateStartDay.IsNull() {
		data.SummerTimeDateStartDay = types.Int64Value(value.Int())
	} else {
		data.SummerTimeDateStartDay = types.Int64Null()
	}
	if value := res.Get(prefix + "summer-time.start-month"); value.Exists() && !data.SummerTimeDateStartMonth.IsNull() {
		data.SummerTimeDateStartMonth = types.StringValue(value.String())
	} else {
		data.SummerTimeDateStartMonth = types.StringNull()
	}
	if value := res.Get(prefix + "summer-time.start-year"); value.Exists() && !data.SummerTimeDateStartYear.IsNull() {
		data.SummerTimeDateStartYear = types.Int64Value(value.Int())
	} else {
		data.SummerTimeDateStartYear = types.Int64Null()
	}
	if value := res.Get(prefix + "summer-time.start-time"); value.Exists() && !data.SummerTimeDateStartTime.IsNull() {
		data.SummerTimeDateStartTime = types.StringValue(value.String())
	} else {
		data.SummerTimeDateStartTime = types.StringNull()
	}
	if value := res.Get(prefix + "summer-time.date-end-day"); value.Exists() && !data.SummerTimeDateEndDay.IsNull() {
		data.SummerTimeDateEndDay = types.Int64Value(value.Int())
	} else {
		data.SummerTimeDateEndDay = types.Int64Null()
	}
	if value := res.Get(prefix + "summer-time.date-end-month"); value.Exists() && !data.SummerTimeDateEndMonth.IsNull() {
		data.SummerTimeDateEndMonth = types.StringValue(value.String())
	} else {
		data.SummerTimeDateEndMonth = types.StringNull()
	}
	if value := res.Get(prefix + "summer-time.date-end-year"); value.Exists() && !data.SummerTimeDateEndYear.IsNull() {
		data.SummerTimeDateEndYear = types.Int64Value(value.Int())
	} else {
		data.SummerTimeDateEndYear = types.Int64Null()
	}
	if value := res.Get(prefix + "summer-time.date-end-time"); value.Exists() && !data.SummerTimeDateEndTime.IsNull() {
		data.SummerTimeDateEndTime = types.StringValue(value.String())
	} else {
		data.SummerTimeDateEndTime = types.StringNull()
	}
	if value := res.Get(prefix + "summer-time.offset"); value.Exists() && !data.SummerTimeDateOffset.IsNull() {
		data.SummerTimeDateOffset = types.Int64Value(value.Int())
	} else {
		data.SummerTimeDateOffset = types.Int64Null()
	}
	if value := res.Get(prefix + "summer-time.recurring"); !data.SummerTimeRecurring.IsNull() {
		if value.Exists() {
			data.SummerTimeRecurring = types.BoolValue(true)
		} else {
			data.SummerTimeRecurring = types.BoolValue(false)
		}
	} else {
		data.SummerTimeRecurring = types.BoolNull()
	}
	if value := res.Get(prefix + "summer-time.recurring-start"); value.Exists() && !data.SummerTimeRecurringStartWeek.IsNull() {
		data.SummerTimeRecurringStartWeek = types.StringValue(value.String())
	} else {
		data.SummerTimeRecurringStartWeek = types.StringNull()
	}
	if value := res.Get(prefix + "summer-time.recurring-start-day"); value.Exists() && !data.SummerTimeRecurringStartWeekday.IsNull() {
		data.SummerTimeRecurringStartWeekday = types.StringValue(value.String())
	} else {
		data.SummerTimeRecurringStartWeekday = types.StringNull()
	}
	if value := res.Get(prefix + "summer-time.recurring-start-month"); value.Exists() && !data.SummerTimeRecurringStartMonth.IsNull() {
		data.SummerTimeRecurringStartMonth = types.StringValue(value.String())
	} else {
		data.SummerTimeRecurringStartMonth = types.StringNull()
	}
	if value := res.Get(prefix + "summer-time.recurring-start-time"); value.Exists() && !data.SummerTimeRecurringStartTime.IsNull() {
		data.SummerTimeRecurringStartTime = types.StringValue(value.String())
	} else {
		data.SummerTimeRecurringStartTime = types.StringNull()
	}
	if value := res.Get(prefix + "summer-time.recurring-end"); value.Exists() && !data.SummerTimeRecurringEndWeek.IsNull() {
		data.SummerTimeRecurringEndWeek = types.StringValue(value.String())
	} else {
		data.SummerTimeRecurringEndWeek = types.StringNull()
	}
	if value := res.Get(prefix + "summer-time.recurring-end-day"); value.Exists() && !data.SummerTimeRecurringEndWeekday.IsNull() {
		data.SummerTimeRecurringEndWeekday = types.StringValue(value.String())
	} else {
		data.SummerTimeRecurringEndWeekday = types.StringNull()
	}
	if value := res.Get(prefix + "summer-time.recurring-end-month"); value.Exists() && !data.SummerTimeRecurringEndMonth.IsNull() {
		data.SummerTimeRecurringEndMonth = types.StringValue(value.String())
	} else {
		data.SummerTimeRecurringEndMonth = types.StringNull()
	}
	if value := res.Get(prefix + "summer-time.recurring-end-time"); value.Exists() && !data.SummerTimeRecurringEndTime.IsNull() {
		data.SummerTimeRecurringEndTime = types.StringValue(value.String())
	} else {
		data.SummerTimeRecurringEndTime = types.StringNull()
	}
	if value := res.Get(prefix + "summer-time.recurring-offset"); value.Exists() && !data.SummerTimeRecurringOffset.IsNull() {
		data.SummerTimeRecurringOffset = types.Int64Value(value.Int())
	} else {
		data.SummerTimeRecurringOffset = types.Int64Null()
	}
}

func (data *ClockData) fromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "calendar-valid"); value.Exists() {
		data.CalendarValid = types.BoolValue(true)
	} else {
		data.CalendarValid = types.BoolValue(false)
	}
	if value := res.Get(prefix + "summer-time.zone"); value.Exists() {
		data.SummerTimeZone = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "summer-time.date"); value.Exists() {
		data.SummerTimeDate = types.BoolValue(true)
	} else {
		data.SummerTimeDate = types.BoolValue(false)
	}
	if value := res.Get(prefix + "summer-time.start-day"); value.Exists() {
		data.SummerTimeDateStartDay = types.Int64Value(value.Int())
	}
	if value := res.Get(prefix + "summer-time.start-month"); value.Exists() {
		data.SummerTimeDateStartMonth = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "summer-time.start-year"); value.Exists() {
		data.SummerTimeDateStartYear = types.Int64Value(value.Int())
	}
	if value := res.Get(prefix + "summer-time.start-time"); value.Exists() {
		data.SummerTimeDateStartTime = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "summer-time.date-end-day"); value.Exists() {
		data.SummerTimeDateEndDay = types.Int64Value(value.Int())
	}
	if value := res.Get(prefix + "summer-time.date-end-month"); value.Exists() {
		data.SummerTimeDateEndMonth = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "summer-time.date-end-year"); value.Exists() {
		data.SummerTimeDateEndYear = types.Int64Value(value.Int())
	}
	if value := res.Get(prefix + "summer-time.date-end-time"); value.Exists() {
		data.SummerTimeDateEndTime = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "summer-time.offset"); value.Exists() {
		data.SummerTimeDateOffset = types.Int64Value(value.Int())
	}
	if value := res.Get(prefix + "summer-time.recurring"); value.Exists() {
		data.SummerTimeRecurring = types.BoolValue(true)
	} else {
		data.SummerTimeRecurring = types.BoolValue(false)
	}
	if value := res.Get(prefix + "summer-time.recurring-start"); value.Exists() {
		data.SummerTimeRecurringStartWeek = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "summer-time.recurring-start-day"); value.Exists() {
		data.SummerTimeRecurringStartWeekday = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "summer-time.recurring-start-month"); value.Exists() {
		data.SummerTimeRecurringStartMonth = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "summer-time.recurring-start-time"); value.Exists() {
		data.SummerTimeRecurringStartTime = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "summer-time.recurring-end"); value.Exists() {
		data.SummerTimeRecurringEndWeek = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "summer-time.recurring-end-day"); value.Exists() {
		data.SummerTimeRecurringEndWeekday = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "summer-time.recurring-end-month"); value.Exists() {
		data.SummerTimeRecurringEndMonth = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "summer-time.recurring-end-time"); value.Exists() {
		data.SummerTimeRecurringEndTime = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "summer-time.recurring-offset"); value.Exists() {
		data.SummerTimeRecurringOffset = types.Int64Value(value.Int())
	}
}

func (data *Clock) getDeletedListItems(ctx context.Context, state Clock) []string {
	deletedListItems := make([]string, 0)
	return deletedListItems
}

func (data *Clock) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	if !data.CalendarValid.IsNull() && !data.CalendarValid.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/calendar-valid", data.getPath()))
	}
	if !data.SummerTimeDate.IsNull() && !data.SummerTimeDate.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/summer-time/date", data.getPath()))
	}
	if !data.SummerTimeRecurring.IsNull() && !data.SummerTimeRecurring.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/summer-time/recurring", data.getPath()))
	}
	return emptyLeafsDelete
}

func (data *Clock) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	if !data.CalendarValid.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/calendar-valid", data.getPath()))
	}
	if !data.SummerTimeZone.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/summer-time", data.getPath()))
	}
	if !data.SummerTimeDate.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/summer-time/date", data.getPath()))
	}
	if !data.SummerTimeDateStartDay.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/summer-time/start-day", data.getPath()))
	}
	if !data.SummerTimeDateStartMonth.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/summer-time/start-month", data.getPath()))
	}
	if !data.SummerTimeDateStartYear.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/summer-time/start-year", data.getPath()))
	}
	if !data.SummerTimeDateStartTime.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/summer-time/start-time", data.getPath()))
	}
	if !data.SummerTimeDateEndDay.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/summer-time/date-end-day", data.getPath()))
	}
	if !data.SummerTimeDateEndMonth.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/summer-time/date-end-month", data.getPath()))
	}
	if !data.SummerTimeDateEndYear.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/summer-time/date-end-year", data.getPath()))
	}
	if !data.SummerTimeDateEndTime.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/summer-time/date-end-time", data.getPath()))
	}
	if !data.SummerTimeDateOffset.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/summer-time/offset", data.getPath()))
	}
	if !data.SummerTimeRecurring.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/summer-time/recurring", data.getPath()))
	}
	if !data.SummerTimeRecurringStartWeek.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/summer-time/recurring-start", data.getPath()))
	}
	if !data.SummerTimeRecurringStartWeekday.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/summer-time/recurring-start-day", data.getPath()))
	}
	if !data.SummerTimeRecurringStartMonth.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/summer-time/recurring-start-month", data.getPath()))
	}
	if !data.SummerTimeRecurringStartTime.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/summer-time/recurring-start-time", data.getPath()))
	}
	if !data.SummerTimeRecurringEndWeek.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/summer-time/recurring-end", data.getPath()))
	}
	if !data.SummerTimeRecurringEndWeekday.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/summer-time/recurring-end-day", data.getPath()))
	}
	if !data.SummerTimeRecurringEndMonth.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/summer-time/recurring-end-month", data.getPath()))
	}
	if !data.SummerTimeRecurringEndTime.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/summer-time/recurring-end-time", data.getPath()))
	}
	if !data.SummerTimeRecurringOffset.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/summer-time/recurring-offset", data.getPath()))
	}
	return deletePaths
}