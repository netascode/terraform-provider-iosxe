// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/terraform-provider-iosxe/internal/provider/helpers"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type AccessListStandard struct {
	Device  types.String                `tfsdk:"device"`
	Id      types.String                `tfsdk:"id"`
	Name    types.String                `tfsdk:"name"`
	Entries []AccessListStandardEntries `tfsdk:"entries"`
}
type AccessListStandardEntries struct {
	Sequence         types.Int64  `tfsdk:"sequence"`
	Remark           types.String `tfsdk:"remark"`
	DenyPrefix       types.String `tfsdk:"deny_prefix"`
	DenyPrefixMask   types.String `tfsdk:"deny_prefix_mask"`
	DenyAny          types.Bool   `tfsdk:"deny_any"`
	DenyHost         types.String `tfsdk:"deny_host"`
	PermitPrefix     types.String `tfsdk:"permit_prefix"`
	PermitPrefixMask types.String `tfsdk:"permit_prefix_mask"`
	PermitAny        types.Bool   `tfsdk:"permit_any"`
	PermitHost       types.String `tfsdk:"permit_host"`
}

func (data AccessListStandard) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/ip/access-list/Cisco-IOS-XE-acl:standard=%v", url.QueryEscape(fmt.Sprintf("%v", data.Name.ValueString())))
}

// if last path element has a key -> remove it
func (data AccessListStandard) getPathShort() string {
	path := data.getPath()
	re := regexp.MustCompile(`(.*)=[^\/]*$`)
	matches := re.FindStringSubmatch(path)
	if len(matches) <= 1 {
		return path
	}
	return matches[1]
}

func (data AccessListStandard) toBody(ctx context.Context) string {
	body := `{"` + helpers.LastElement(data.getPath()) + `":{}}`
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"name", data.Name.ValueString())
	}
	if len(data.Entries) > 0 {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"access-list-seq-rule", []interface{}{})
		for index, item := range data.Entries {
			if !item.Sequence.IsNull() && !item.Sequence.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"access-list-seq-rule"+"."+strconv.Itoa(index)+"."+"sequence", strconv.FormatInt(item.Sequence.ValueInt64(), 10))
			}
			if !item.Remark.IsNull() && !item.Remark.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"access-list-seq-rule"+"."+strconv.Itoa(index)+"."+"remark", item.Remark.ValueString())
			}
			if !item.DenyPrefix.IsNull() && !item.DenyPrefix.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"access-list-seq-rule"+"."+strconv.Itoa(index)+"."+"deny.std-ace.ipv4-prefix", item.DenyPrefix.ValueString())
			}
			if !item.DenyPrefixMask.IsNull() && !item.DenyPrefixMask.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"access-list-seq-rule"+"."+strconv.Itoa(index)+"."+"deny.std-ace.mask", item.DenyPrefixMask.ValueString())
			}
			if !item.DenyAny.IsNull() && !item.DenyAny.IsUnknown() {
				if item.DenyAny.ValueBool() {
					body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"access-list-seq-rule"+"."+strconv.Itoa(index)+"."+"deny.std-ace.any", map[string]string{})
				}
			}
			if !item.DenyHost.IsNull() && !item.DenyHost.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"access-list-seq-rule"+"."+strconv.Itoa(index)+"."+"deny.std-ace.host", item.DenyHost.ValueString())
			}
			if !item.PermitPrefix.IsNull() && !item.PermitPrefix.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"access-list-seq-rule"+"."+strconv.Itoa(index)+"."+"permit.std-ace.ipv4-prefix", item.PermitPrefix.ValueString())
			}
			if !item.PermitPrefixMask.IsNull() && !item.PermitPrefixMask.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"access-list-seq-rule"+"."+strconv.Itoa(index)+"."+"permit.std-ace.mask", item.PermitPrefixMask.ValueString())
			}
			if !item.PermitAny.IsNull() && !item.PermitAny.IsUnknown() {
				if item.PermitAny.ValueBool() {
					body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"access-list-seq-rule"+"."+strconv.Itoa(index)+"."+"permit.std-ace.any", map[string]string{})
				}
			}
			if !item.PermitHost.IsNull() && !item.PermitHost.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"access-list-seq-rule"+"."+strconv.Itoa(index)+"."+"permit.std-ace.host", item.PermitHost.ValueString())
			}
		}
	}
	return body
}

func (data *AccessListStandard) updateFromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	for i := range data.Entries {
		keys := [...]string{"sequence"}
		keyValues := [...]string{strconv.FormatInt(data.Entries[i].Sequence.ValueInt64(), 10)}

		var r gjson.Result
		res.Get(prefix + "access-list-seq-rule").ForEach(
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
		if value := r.Get("sequence"); value.Exists() && !data.Entries[i].Sequence.IsNull() {
			data.Entries[i].Sequence = types.Int64Value(value.Int())
		} else {
			data.Entries[i].Sequence = types.Int64Null()
		}
		if value := r.Get("remark"); value.Exists() && !data.Entries[i].Remark.IsNull() {
			data.Entries[i].Remark = types.StringValue(value.String())
		} else {
			data.Entries[i].Remark = types.StringNull()
		}
		if value := r.Get("deny.std-ace.ipv4-prefix"); value.Exists() && !data.Entries[i].DenyPrefix.IsNull() {
			data.Entries[i].DenyPrefix = types.StringValue(value.String())
		} else {
			data.Entries[i].DenyPrefix = types.StringNull()
		}
		if value := r.Get("deny.std-ace.mask"); value.Exists() && !data.Entries[i].DenyPrefixMask.IsNull() {
			data.Entries[i].DenyPrefixMask = types.StringValue(value.String())
		} else {
			data.Entries[i].DenyPrefixMask = types.StringNull()
		}
		if value := r.Get("deny.std-ace.any"); !data.Entries[i].DenyAny.IsNull() {
			if value.Exists() {
				data.Entries[i].DenyAny = types.BoolValue(true)
			} else {
				data.Entries[i].DenyAny = types.BoolValue(false)
			}
		} else {
			data.Entries[i].DenyAny = types.BoolNull()
		}
		if value := r.Get("deny.std-ace.host"); value.Exists() && !data.Entries[i].DenyHost.IsNull() {
			data.Entries[i].DenyHost = types.StringValue(value.String())
		} else {
			data.Entries[i].DenyHost = types.StringNull()
		}
		if value := r.Get("permit.std-ace.ipv4-prefix"); value.Exists() && !data.Entries[i].PermitPrefix.IsNull() {
			data.Entries[i].PermitPrefix = types.StringValue(value.String())
		} else {
			data.Entries[i].PermitPrefix = types.StringNull()
		}
		if value := r.Get("permit.std-ace.mask"); value.Exists() && !data.Entries[i].PermitPrefixMask.IsNull() {
			data.Entries[i].PermitPrefixMask = types.StringValue(value.String())
		} else {
			data.Entries[i].PermitPrefixMask = types.StringNull()
		}
		if value := r.Get("permit.std-ace.any"); !data.Entries[i].PermitAny.IsNull() {
			if value.Exists() {
				data.Entries[i].PermitAny = types.BoolValue(true)
			} else {
				data.Entries[i].PermitAny = types.BoolValue(false)
			}
		} else {
			data.Entries[i].PermitAny = types.BoolNull()
		}
		if value := r.Get("permit.std-ace.host"); value.Exists() && !data.Entries[i].PermitHost.IsNull() {
			data.Entries[i].PermitHost = types.StringValue(value.String())
		} else {
			data.Entries[i].PermitHost = types.StringNull()
		}
	}
}

func (data *AccessListStandard) fromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "access-list-seq-rule"); value.Exists() {
		data.Entries = make([]AccessListStandardEntries, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := AccessListStandardEntries{}
			if cValue := v.Get("sequence"); cValue.Exists() {
				item.Sequence = types.Int64Value(cValue.Int())
			}
			if cValue := v.Get("remark"); cValue.Exists() {
				item.Remark = types.StringValue(cValue.String())
			}
			if cValue := v.Get("deny.std-ace.ipv4-prefix"); cValue.Exists() {
				item.DenyPrefix = types.StringValue(cValue.String())
			}
			if cValue := v.Get("deny.std-ace.mask"); cValue.Exists() {
				item.DenyPrefixMask = types.StringValue(cValue.String())
			}
			if cValue := v.Get("deny.std-ace.any"); cValue.Exists() {
				item.DenyAny = types.BoolValue(true)
			} else {
				item.DenyAny = types.BoolValue(false)
			}
			if cValue := v.Get("deny.std-ace.host"); cValue.Exists() {
				item.DenyHost = types.StringValue(cValue.String())
			}
			if cValue := v.Get("permit.std-ace.ipv4-prefix"); cValue.Exists() {
				item.PermitPrefix = types.StringValue(cValue.String())
			}
			if cValue := v.Get("permit.std-ace.mask"); cValue.Exists() {
				item.PermitPrefixMask = types.StringValue(cValue.String())
			}
			if cValue := v.Get("permit.std-ace.any"); cValue.Exists() {
				item.PermitAny = types.BoolValue(true)
			} else {
				item.PermitAny = types.BoolValue(false)
			}
			if cValue := v.Get("permit.std-ace.host"); cValue.Exists() {
				item.PermitHost = types.StringValue(cValue.String())
			}
			data.Entries = append(data.Entries, item)
			return true
		})
	}
}

func (data *AccessListStandard) getDeletedListItems(ctx context.Context, state AccessListStandard) []string {
	deletedListItems := make([]string, 0)
	for i := range state.Entries {
		stateKeyValues := [...]string{strconv.FormatInt(state.Entries[i].Sequence.ValueInt64(), 10)}

		emptyKeys := true
		if !reflect.ValueOf(state.Entries[i].Sequence.ValueInt64()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.Entries {
			found = true
			if state.Entries[i].Sequence.ValueInt64() != data.Entries[j].Sequence.ValueInt64() {
				found = false
			}
			if found {
				break
			}
		}
		if !found {
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/access-list-seq-rule=%v", state.getPath(), strings.Join(stateKeyValues[:], ",")))
		}
	}
	return deletedListItems
}

func (data *AccessListStandard) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)

	for i := range data.Entries {
		keyValues := [...]string{strconv.FormatInt(data.Entries[i].Sequence.ValueInt64(), 10)}
		if !data.Entries[i].DenyAny.ValueBool() {
			emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/access-list-seq-rule=%v/deny-permit/deny/deny/std-ace/source-choice/any-case/any", data.getPath(), strings.Join(keyValues[:], ",")))
		}
		if !data.Entries[i].PermitAny.ValueBool() {
			emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/access-list-seq-rule=%v/deny-permit/permit/permit/std-ace/source-choice/any-case/any", data.getPath(), strings.Join(keyValues[:], ",")))
		}
	}
	return emptyLeafsDelete
}
