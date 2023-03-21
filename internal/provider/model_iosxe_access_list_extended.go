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

type AccessListExtended struct {
	Device  types.String                `tfsdk:"device"`
	Id      types.String                `tfsdk:"id"`
	Name    types.String                `tfsdk:"name"`
	Entries []AccessListExtendedEntries `tfsdk:"entries"`
}
type AccessListExtendedEntries struct {
	Sequence                   types.Int64  `tfsdk:"sequence"`
	Remark                     types.String `tfsdk:"remark"`
	AceRuleAction              types.String `tfsdk:"ace_rule_action"`
	AceRuleProtocol            types.String `tfsdk:"ace_rule_protocol"`
	ServiceObjectGroup         types.String `tfsdk:"service_object_group"`
	SourcePrefix               types.String `tfsdk:"source_prefix"`
	SourcePrefixMask           types.String `tfsdk:"source_prefix_mask"`
	SourceAny                  types.Bool   `tfsdk:"source_any"`
	SourceHost                 types.String `tfsdk:"source_host"`
	SourceObjectGroup          types.String `tfsdk:"source_object_group"`
	SourcePortEqual            types.String `tfsdk:"source_port_equal"`
	SourcePortGreaterThan      types.String `tfsdk:"source_port_greater_than"`
	SourcePortLesserThan       types.String `tfsdk:"source_port_lesser_than"`
	SourcePortRangeFrom        types.String `tfsdk:"source_port_range_from"`
	SourcePortRangeTo          types.String `tfsdk:"source_port_range_to"`
	DestinationPrefix          types.String `tfsdk:"destination_prefix"`
	DestinationPrefixMask      types.String `tfsdk:"destination_prefix_mask"`
	DestinationAny             types.Bool   `tfsdk:"destination_any"`
	DestinationHost            types.String `tfsdk:"destination_host"`
	DestinationObjectGroup     types.String `tfsdk:"destination_object_group"`
	DestinationPortEqual       types.String `tfsdk:"destination_port_equal"`
	DestinationPortGreaterThan types.String `tfsdk:"destination_port_greater_than"`
	DestinationPortLesserThan  types.String `tfsdk:"destination_port_lesser_than"`
	DestinationPortRangeFrom   types.String `tfsdk:"destination_port_range_from"`
	DestinationPortRangeTo     types.String `tfsdk:"destination_port_range_to"`
	Ack                        types.Bool   `tfsdk:"ack"`
	Fin                        types.Bool   `tfsdk:"fin"`
	Psh                        types.Bool   `tfsdk:"psh"`
	Rst                        types.Bool   `tfsdk:"rst"`
	Syn                        types.Bool   `tfsdk:"syn"`
	Urg                        types.Bool   `tfsdk:"urg"`
	Established                types.Bool   `tfsdk:"established"`
	Dscp                       types.String `tfsdk:"dscp"`
	Fragments                  types.Bool   `tfsdk:"fragments"`
	Precedence                 types.String `tfsdk:"precedence"`
	Tos                        types.String `tfsdk:"tos"`
}

func (data AccessListExtended) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/ip/access-list/Cisco-IOS-XE-acl:extended=%v", url.QueryEscape(fmt.Sprintf("%v", data.Name.ValueString())))
}

// if last path element has a key -> remove it
func (data AccessListExtended) getPathShort() string {
	path := data.getPath()
	re := regexp.MustCompile(`(.*)=[^\/]*$`)
	matches := re.FindStringSubmatch(path)
	if len(matches) <= 1 {
		return path
	}
	return matches[1]
}

func (data AccessListExtended) toBody(ctx context.Context) string {
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
			if !item.AceRuleAction.IsNull() && !item.AceRuleAction.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"access-list-seq-rule"+"."+strconv.Itoa(index)+"."+"ace-rule.action", item.AceRuleAction.ValueString())
			}
			if !item.AceRuleProtocol.IsNull() && !item.AceRuleProtocol.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"access-list-seq-rule"+"."+strconv.Itoa(index)+"."+"ace-rule.protocol", item.AceRuleProtocol.ValueString())
			}
			if !item.ServiceObjectGroup.IsNull() && !item.ServiceObjectGroup.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"access-list-seq-rule"+"."+strconv.Itoa(index)+"."+"ace-rule.object-group-str", item.ServiceObjectGroup.ValueString())
			}
			if !item.SourcePrefix.IsNull() && !item.SourcePrefix.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"access-list-seq-rule"+"."+strconv.Itoa(index)+"."+"ace-rule.ipv4-address", item.SourcePrefix.ValueString())
			}
			if !item.SourcePrefixMask.IsNull() && !item.SourcePrefixMask.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"access-list-seq-rule"+"."+strconv.Itoa(index)+"."+"ace-rule.mask", item.SourcePrefixMask.ValueString())
			}
			if !item.SourceAny.IsNull() && !item.SourceAny.IsUnknown() {
				if item.SourceAny.ValueBool() {
					body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"access-list-seq-rule"+"."+strconv.Itoa(index)+"."+"ace-rule.any", map[string]string{})
				}
			}
			if !item.SourceHost.IsNull() && !item.SourceHost.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"access-list-seq-rule"+"."+strconv.Itoa(index)+"."+"ace-rule.host", item.SourceHost.ValueString())
			}
			if !item.SourceObjectGroup.IsNull() && !item.SourceObjectGroup.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"access-list-seq-rule"+"."+strconv.Itoa(index)+"."+"ace-rule.object-group", item.SourceObjectGroup.ValueString())
			}
			if !item.SourcePortEqual.IsNull() && !item.SourcePortEqual.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"access-list-seq-rule"+"."+strconv.Itoa(index)+"."+"ace-rule.src-eq", item.SourcePortEqual.ValueString())
			}
			if !item.SourcePortGreaterThan.IsNull() && !item.SourcePortGreaterThan.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"access-list-seq-rule"+"."+strconv.Itoa(index)+"."+"ace-rule.src-gt", item.SourcePortGreaterThan.ValueString())
			}
			if !item.SourcePortLesserThan.IsNull() && !item.SourcePortLesserThan.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"access-list-seq-rule"+"."+strconv.Itoa(index)+"."+"ace-rule.src-lt", item.SourcePortLesserThan.ValueString())
			}
			if !item.SourcePortRangeFrom.IsNull() && !item.SourcePortRangeFrom.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"access-list-seq-rule"+"."+strconv.Itoa(index)+"."+"ace-rule.src-range1", item.SourcePortRangeFrom.ValueString())
			}
			if !item.SourcePortRangeTo.IsNull() && !item.SourcePortRangeTo.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"access-list-seq-rule"+"."+strconv.Itoa(index)+"."+"ace-rule.src-range2", item.SourcePortRangeTo.ValueString())
			}
			if !item.DestinationPrefix.IsNull() && !item.DestinationPrefix.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"access-list-seq-rule"+"."+strconv.Itoa(index)+"."+"ace-rule.dest-ipv4-address", item.DestinationPrefix.ValueString())
			}
			if !item.DestinationPrefixMask.IsNull() && !item.DestinationPrefixMask.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"access-list-seq-rule"+"."+strconv.Itoa(index)+"."+"ace-rule.dest-mask", item.DestinationPrefixMask.ValueString())
			}
			if !item.DestinationAny.IsNull() && !item.DestinationAny.IsUnknown() {
				if item.DestinationAny.ValueBool() {
					body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"access-list-seq-rule"+"."+strconv.Itoa(index)+"."+"ace-rule.dst-any", map[string]string{})
				}
			}
			if !item.DestinationHost.IsNull() && !item.DestinationHost.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"access-list-seq-rule"+"."+strconv.Itoa(index)+"."+"ace-rule.dst-host", item.DestinationHost.ValueString())
			}
			if !item.DestinationObjectGroup.IsNull() && !item.DestinationObjectGroup.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"access-list-seq-rule"+"."+strconv.Itoa(index)+"."+"ace-rule.dst-object-group", item.DestinationObjectGroup.ValueString())
			}
			if !item.DestinationPortEqual.IsNull() && !item.DestinationPortEqual.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"access-list-seq-rule"+"."+strconv.Itoa(index)+"."+"ace-rule.dst-eq", item.DestinationPortEqual.ValueString())
			}
			if !item.DestinationPortGreaterThan.IsNull() && !item.DestinationPortGreaterThan.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"access-list-seq-rule"+"."+strconv.Itoa(index)+"."+"ace-rule.dst-gt", item.DestinationPortGreaterThan.ValueString())
			}
			if !item.DestinationPortLesserThan.IsNull() && !item.DestinationPortLesserThan.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"access-list-seq-rule"+"."+strconv.Itoa(index)+"."+"ace-rule.dst-lt", item.DestinationPortLesserThan.ValueString())
			}
			if !item.DestinationPortRangeFrom.IsNull() && !item.DestinationPortRangeFrom.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"access-list-seq-rule"+"."+strconv.Itoa(index)+"."+"ace-rule.dst-range1", item.DestinationPortRangeFrom.ValueString())
			}
			if !item.DestinationPortRangeTo.IsNull() && !item.DestinationPortRangeTo.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"access-list-seq-rule"+"."+strconv.Itoa(index)+"."+"ace-rule.dst-range2", item.DestinationPortRangeTo.ValueString())
			}
			if !item.Ack.IsNull() && !item.Ack.IsUnknown() {
				if item.Ack.ValueBool() {
					body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"access-list-seq-rule"+"."+strconv.Itoa(index)+"."+"ace-rule.ack", map[string]string{})
				}
			}
			if !item.Fin.IsNull() && !item.Fin.IsUnknown() {
				if item.Fin.ValueBool() {
					body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"access-list-seq-rule"+"."+strconv.Itoa(index)+"."+"ace-rule.fin", map[string]string{})
				}
			}
			if !item.Psh.IsNull() && !item.Psh.IsUnknown() {
				if item.Psh.ValueBool() {
					body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"access-list-seq-rule"+"."+strconv.Itoa(index)+"."+"ace-rule.psh", map[string]string{})
				}
			}
			if !item.Rst.IsNull() && !item.Rst.IsUnknown() {
				if item.Rst.ValueBool() {
					body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"access-list-seq-rule"+"."+strconv.Itoa(index)+"."+"ace-rule.rst", map[string]string{})
				}
			}
			if !item.Syn.IsNull() && !item.Syn.IsUnknown() {
				if item.Syn.ValueBool() {
					body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"access-list-seq-rule"+"."+strconv.Itoa(index)+"."+"ace-rule.syn", map[string]string{})
				}
			}
			if !item.Urg.IsNull() && !item.Urg.IsUnknown() {
				if item.Urg.ValueBool() {
					body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"access-list-seq-rule"+"."+strconv.Itoa(index)+"."+"ace-rule.urg", map[string]string{})
				}
			}
			if !item.Established.IsNull() && !item.Established.IsUnknown() {
				if item.Established.ValueBool() {
					body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"access-list-seq-rule"+"."+strconv.Itoa(index)+"."+"ace-rule.established", map[string]string{})
				}
			}
			if !item.Dscp.IsNull() && !item.Dscp.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"access-list-seq-rule"+"."+strconv.Itoa(index)+"."+"ace-rule.dscp", item.Dscp.ValueString())
			}
			if !item.Fragments.IsNull() && !item.Fragments.IsUnknown() {
				if item.Fragments.ValueBool() {
					body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"access-list-seq-rule"+"."+strconv.Itoa(index)+"."+"ace-rule.fragments", map[string]string{})
				}
			}
			if !item.Precedence.IsNull() && !item.Precedence.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"access-list-seq-rule"+"."+strconv.Itoa(index)+"."+"ace-rule.precedence", item.Precedence.ValueString())
			}
			if !item.Tos.IsNull() && !item.Tos.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"access-list-seq-rule"+"."+strconv.Itoa(index)+"."+"ace-rule.tos", item.Tos.ValueString())
			}
		}
	}
	return body
}

func (data *AccessListExtended) updateFromBody(ctx context.Context, res gjson.Result) {
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
		if value := r.Get("ace-rule.action"); value.Exists() && !data.Entries[i].AceRuleAction.IsNull() {
			data.Entries[i].AceRuleAction = types.StringValue(value.String())
		} else {
			data.Entries[i].AceRuleAction = types.StringNull()
		}
		if value := r.Get("ace-rule.protocol"); value.Exists() && !data.Entries[i].AceRuleProtocol.IsNull() {
			data.Entries[i].AceRuleProtocol = types.StringValue(value.String())
		} else {
			data.Entries[i].AceRuleProtocol = types.StringNull()
		}
		if value := r.Get("ace-rule.object-group-str"); value.Exists() && !data.Entries[i].ServiceObjectGroup.IsNull() {
			data.Entries[i].ServiceObjectGroup = types.StringValue(value.String())
		} else {
			data.Entries[i].ServiceObjectGroup = types.StringNull()
		}
		if value := r.Get("ace-rule.ipv4-address"); value.Exists() && !data.Entries[i].SourcePrefix.IsNull() {
			data.Entries[i].SourcePrefix = types.StringValue(value.String())
		} else {
			data.Entries[i].SourcePrefix = types.StringNull()
		}
		if value := r.Get("ace-rule.mask"); value.Exists() && !data.Entries[i].SourcePrefixMask.IsNull() {
			data.Entries[i].SourcePrefixMask = types.StringValue(value.String())
		} else {
			data.Entries[i].SourcePrefixMask = types.StringNull()
		}
		if value := r.Get("ace-rule.any"); !data.Entries[i].SourceAny.IsNull() {
			if value.Exists() {
				data.Entries[i].SourceAny = types.BoolValue(true)
			} else {
				data.Entries[i].SourceAny = types.BoolValue(false)
			}
		} else {
			data.Entries[i].SourceAny = types.BoolNull()
		}
		if value := r.Get("ace-rule.host"); value.Exists() && !data.Entries[i].SourceHost.IsNull() {
			data.Entries[i].SourceHost = types.StringValue(value.String())
		} else {
			data.Entries[i].SourceHost = types.StringNull()
		}
		if value := r.Get("ace-rule.object-group"); value.Exists() && !data.Entries[i].SourceObjectGroup.IsNull() {
			data.Entries[i].SourceObjectGroup = types.StringValue(value.String())
		} else {
			data.Entries[i].SourceObjectGroup = types.StringNull()
		}
		if value := r.Get("ace-rule.src-eq"); value.Exists() && !data.Entries[i].SourcePortEqual.IsNull() {
			data.Entries[i].SourcePortEqual = types.StringValue(value.String())
		} else {
			data.Entries[i].SourcePortEqual = types.StringNull()
		}
		if value := r.Get("ace-rule.src-gt"); value.Exists() && !data.Entries[i].SourcePortGreaterThan.IsNull() {
			data.Entries[i].SourcePortGreaterThan = types.StringValue(value.String())
		} else {
			data.Entries[i].SourcePortGreaterThan = types.StringNull()
		}
		if value := r.Get("ace-rule.src-lt"); value.Exists() && !data.Entries[i].SourcePortLesserThan.IsNull() {
			data.Entries[i].SourcePortLesserThan = types.StringValue(value.String())
		} else {
			data.Entries[i].SourcePortLesserThan = types.StringNull()
		}
		if value := r.Get("ace-rule.src-range1"); value.Exists() && !data.Entries[i].SourcePortRangeFrom.IsNull() {
			data.Entries[i].SourcePortRangeFrom = types.StringValue(value.String())
		} else {
			data.Entries[i].SourcePortRangeFrom = types.StringNull()
		}
		if value := r.Get("ace-rule.src-range2"); value.Exists() && !data.Entries[i].SourcePortRangeTo.IsNull() {
			data.Entries[i].SourcePortRangeTo = types.StringValue(value.String())
		} else {
			data.Entries[i].SourcePortRangeTo = types.StringNull()
		}
		if value := r.Get("ace-rule.dest-ipv4-address"); value.Exists() && !data.Entries[i].DestinationPrefix.IsNull() {
			data.Entries[i].DestinationPrefix = types.StringValue(value.String())
		} else {
			data.Entries[i].DestinationPrefix = types.StringNull()
		}
		if value := r.Get("ace-rule.dest-mask"); value.Exists() && !data.Entries[i].DestinationPrefixMask.IsNull() {
			data.Entries[i].DestinationPrefixMask = types.StringValue(value.String())
		} else {
			data.Entries[i].DestinationPrefixMask = types.StringNull()
		}
		if value := r.Get("ace-rule.dst-any"); !data.Entries[i].DestinationAny.IsNull() {
			if value.Exists() {
				data.Entries[i].DestinationAny = types.BoolValue(true)
			} else {
				data.Entries[i].DestinationAny = types.BoolValue(false)
			}
		} else {
			data.Entries[i].DestinationAny = types.BoolNull()
		}
		if value := r.Get("ace-rule.dst-host"); value.Exists() && !data.Entries[i].DestinationHost.IsNull() {
			data.Entries[i].DestinationHost = types.StringValue(value.String())
		} else {
			data.Entries[i].DestinationHost = types.StringNull()
		}
		if value := r.Get("ace-rule.dst-object-group"); value.Exists() && !data.Entries[i].DestinationObjectGroup.IsNull() {
			data.Entries[i].DestinationObjectGroup = types.StringValue(value.String())
		} else {
			data.Entries[i].DestinationObjectGroup = types.StringNull()
		}
		if value := r.Get("ace-rule.dst-eq"); value.Exists() && !data.Entries[i].DestinationPortEqual.IsNull() {
			data.Entries[i].DestinationPortEqual = types.StringValue(value.String())
		} else {
			data.Entries[i].DestinationPortEqual = types.StringNull()
		}
		if value := r.Get("ace-rule.dst-gt"); value.Exists() && !data.Entries[i].DestinationPortGreaterThan.IsNull() {
			data.Entries[i].DestinationPortGreaterThan = types.StringValue(value.String())
		} else {
			data.Entries[i].DestinationPortGreaterThan = types.StringNull()
		}
		if value := r.Get("ace-rule.dst-lt"); value.Exists() && !data.Entries[i].DestinationPortLesserThan.IsNull() {
			data.Entries[i].DestinationPortLesserThan = types.StringValue(value.String())
		} else {
			data.Entries[i].DestinationPortLesserThan = types.StringNull()
		}
		if value := r.Get("ace-rule.dst-range1"); value.Exists() && !data.Entries[i].DestinationPortRangeFrom.IsNull() {
			data.Entries[i].DestinationPortRangeFrom = types.StringValue(value.String())
		} else {
			data.Entries[i].DestinationPortRangeFrom = types.StringNull()
		}
		if value := r.Get("ace-rule.dst-range2"); value.Exists() && !data.Entries[i].DestinationPortRangeTo.IsNull() {
			data.Entries[i].DestinationPortRangeTo = types.StringValue(value.String())
		} else {
			data.Entries[i].DestinationPortRangeTo = types.StringNull()
		}
		if value := r.Get("ace-rule.ack"); !data.Entries[i].Ack.IsNull() {
			if value.Exists() {
				data.Entries[i].Ack = types.BoolValue(true)
			} else {
				data.Entries[i].Ack = types.BoolValue(false)
			}
		} else {
			data.Entries[i].Ack = types.BoolNull()
		}
		if value := r.Get("ace-rule.fin"); !data.Entries[i].Fin.IsNull() {
			if value.Exists() {
				data.Entries[i].Fin = types.BoolValue(true)
			} else {
				data.Entries[i].Fin = types.BoolValue(false)
			}
		} else {
			data.Entries[i].Fin = types.BoolNull()
		}
		if value := r.Get("ace-rule.psh"); !data.Entries[i].Psh.IsNull() {
			if value.Exists() {
				data.Entries[i].Psh = types.BoolValue(true)
			} else {
				data.Entries[i].Psh = types.BoolValue(false)
			}
		} else {
			data.Entries[i].Psh = types.BoolNull()
		}
		if value := r.Get("ace-rule.rst"); !data.Entries[i].Rst.IsNull() {
			if value.Exists() {
				data.Entries[i].Rst = types.BoolValue(true)
			} else {
				data.Entries[i].Rst = types.BoolValue(false)
			}
		} else {
			data.Entries[i].Rst = types.BoolNull()
		}
		if value := r.Get("ace-rule.syn"); !data.Entries[i].Syn.IsNull() {
			if value.Exists() {
				data.Entries[i].Syn = types.BoolValue(true)
			} else {
				data.Entries[i].Syn = types.BoolValue(false)
			}
		} else {
			data.Entries[i].Syn = types.BoolNull()
		}
		if value := r.Get("ace-rule.urg"); !data.Entries[i].Urg.IsNull() {
			if value.Exists() {
				data.Entries[i].Urg = types.BoolValue(true)
			} else {
				data.Entries[i].Urg = types.BoolValue(false)
			}
		} else {
			data.Entries[i].Urg = types.BoolNull()
		}
		if value := r.Get("ace-rule.established"); !data.Entries[i].Established.IsNull() {
			if value.Exists() {
				data.Entries[i].Established = types.BoolValue(true)
			} else {
				data.Entries[i].Established = types.BoolValue(false)
			}
		} else {
			data.Entries[i].Established = types.BoolNull()
		}
		if value := r.Get("ace-rule.dscp"); value.Exists() && !data.Entries[i].Dscp.IsNull() {
			data.Entries[i].Dscp = types.StringValue(value.String())
		} else {
			data.Entries[i].Dscp = types.StringNull()
		}
		if value := r.Get("ace-rule.fragments"); !data.Entries[i].Fragments.IsNull() {
			if value.Exists() {
				data.Entries[i].Fragments = types.BoolValue(true)
			} else {
				data.Entries[i].Fragments = types.BoolValue(false)
			}
		} else {
			data.Entries[i].Fragments = types.BoolNull()
		}
		if value := r.Get("ace-rule.precedence"); value.Exists() && !data.Entries[i].Precedence.IsNull() {
			data.Entries[i].Precedence = types.StringValue(value.String())
		} else {
			data.Entries[i].Precedence = types.StringNull()
		}
		if value := r.Get("ace-rule.tos"); value.Exists() && !data.Entries[i].Tos.IsNull() {
			data.Entries[i].Tos = types.StringValue(value.String())
		} else {
			data.Entries[i].Tos = types.StringNull()
		}
	}
}

func (data *AccessListExtended) fromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "access-list-seq-rule"); value.Exists() {
		data.Entries = make([]AccessListExtendedEntries, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := AccessListExtendedEntries{}
			if cValue := v.Get("sequence"); cValue.Exists() {
				item.Sequence = types.Int64Value(cValue.Int())
			}
			if cValue := v.Get("remark"); cValue.Exists() {
				item.Remark = types.StringValue(cValue.String())
			}
			if cValue := v.Get("ace-rule.action"); cValue.Exists() {
				item.AceRuleAction = types.StringValue(cValue.String())
			}
			if cValue := v.Get("ace-rule.protocol"); cValue.Exists() {
				item.AceRuleProtocol = types.StringValue(cValue.String())
			}
			if cValue := v.Get("ace-rule.object-group-str"); cValue.Exists() {
				item.ServiceObjectGroup = types.StringValue(cValue.String())
			}
			if cValue := v.Get("ace-rule.ipv4-address"); cValue.Exists() {
				item.SourcePrefix = types.StringValue(cValue.String())
			}
			if cValue := v.Get("ace-rule.mask"); cValue.Exists() {
				item.SourcePrefixMask = types.StringValue(cValue.String())
			}
			if cValue := v.Get("ace-rule.any"); cValue.Exists() {
				item.SourceAny = types.BoolValue(true)
			} else {
				item.SourceAny = types.BoolValue(false)
			}
			if cValue := v.Get("ace-rule.host"); cValue.Exists() {
				item.SourceHost = types.StringValue(cValue.String())
			}
			if cValue := v.Get("ace-rule.object-group"); cValue.Exists() {
				item.SourceObjectGroup = types.StringValue(cValue.String())
			}
			if cValue := v.Get("ace-rule.src-eq"); cValue.Exists() {
				item.SourcePortEqual = types.StringValue(cValue.String())
			}
			if cValue := v.Get("ace-rule.src-gt"); cValue.Exists() {
				item.SourcePortGreaterThan = types.StringValue(cValue.String())
			}
			if cValue := v.Get("ace-rule.src-lt"); cValue.Exists() {
				item.SourcePortLesserThan = types.StringValue(cValue.String())
			}
			if cValue := v.Get("ace-rule.src-range1"); cValue.Exists() {
				item.SourcePortRangeFrom = types.StringValue(cValue.String())
			}
			if cValue := v.Get("ace-rule.src-range2"); cValue.Exists() {
				item.SourcePortRangeTo = types.StringValue(cValue.String())
			}
			if cValue := v.Get("ace-rule.dest-ipv4-address"); cValue.Exists() {
				item.DestinationPrefix = types.StringValue(cValue.String())
			}
			if cValue := v.Get("ace-rule.dest-mask"); cValue.Exists() {
				item.DestinationPrefixMask = types.StringValue(cValue.String())
			}
			if cValue := v.Get("ace-rule.dst-any"); cValue.Exists() {
				item.DestinationAny = types.BoolValue(true)
			} else {
				item.DestinationAny = types.BoolValue(false)
			}
			if cValue := v.Get("ace-rule.dst-host"); cValue.Exists() {
				item.DestinationHost = types.StringValue(cValue.String())
			}
			if cValue := v.Get("ace-rule.dst-object-group"); cValue.Exists() {
				item.DestinationObjectGroup = types.StringValue(cValue.String())
			}
			if cValue := v.Get("ace-rule.dst-eq"); cValue.Exists() {
				item.DestinationPortEqual = types.StringValue(cValue.String())
			}
			if cValue := v.Get("ace-rule.dst-gt"); cValue.Exists() {
				item.DestinationPortGreaterThan = types.StringValue(cValue.String())
			}
			if cValue := v.Get("ace-rule.dst-lt"); cValue.Exists() {
				item.DestinationPortLesserThan = types.StringValue(cValue.String())
			}
			if cValue := v.Get("ace-rule.dst-range1"); cValue.Exists() {
				item.DestinationPortRangeFrom = types.StringValue(cValue.String())
			}
			if cValue := v.Get("ace-rule.dst-range2"); cValue.Exists() {
				item.DestinationPortRangeTo = types.StringValue(cValue.String())
			}
			if cValue := v.Get("ace-rule.ack"); cValue.Exists() {
				item.Ack = types.BoolValue(true)
			} else {
				item.Ack = types.BoolValue(false)
			}
			if cValue := v.Get("ace-rule.fin"); cValue.Exists() {
				item.Fin = types.BoolValue(true)
			} else {
				item.Fin = types.BoolValue(false)
			}
			if cValue := v.Get("ace-rule.psh"); cValue.Exists() {
				item.Psh = types.BoolValue(true)
			} else {
				item.Psh = types.BoolValue(false)
			}
			if cValue := v.Get("ace-rule.rst"); cValue.Exists() {
				item.Rst = types.BoolValue(true)
			} else {
				item.Rst = types.BoolValue(false)
			}
			if cValue := v.Get("ace-rule.syn"); cValue.Exists() {
				item.Syn = types.BoolValue(true)
			} else {
				item.Syn = types.BoolValue(false)
			}
			if cValue := v.Get("ace-rule.urg"); cValue.Exists() {
				item.Urg = types.BoolValue(true)
			} else {
				item.Urg = types.BoolValue(false)
			}
			if cValue := v.Get("ace-rule.established"); cValue.Exists() {
				item.Established = types.BoolValue(true)
			} else {
				item.Established = types.BoolValue(false)
			}
			if cValue := v.Get("ace-rule.dscp"); cValue.Exists() {
				item.Dscp = types.StringValue(cValue.String())
			}
			if cValue := v.Get("ace-rule.fragments"); cValue.Exists() {
				item.Fragments = types.BoolValue(true)
			} else {
				item.Fragments = types.BoolValue(false)
			}
			if cValue := v.Get("ace-rule.precedence"); cValue.Exists() {
				item.Precedence = types.StringValue(cValue.String())
			}
			if cValue := v.Get("ace-rule.tos"); cValue.Exists() {
				item.Tos = types.StringValue(cValue.String())
			}
			data.Entries = append(data.Entries, item)
			return true
		})
	}
}

func (data *AccessListExtended) getDeletedListItems(ctx context.Context, state AccessListExtended) []string {
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

func (data *AccessListExtended) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)

	for i := range data.Entries {
		keyValues := [...]string{strconv.FormatInt(data.Entries[i].Sequence.ValueInt64(), 10)}
		if !data.Entries[i].SourceAny.ValueBool() {
			emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/access-list-seq-rule=%v/ace-rule/source-choice/any-case/any", data.getPath(), strings.Join(keyValues[:], ",")))
		}
		if !data.Entries[i].DestinationAny.ValueBool() {
			emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/access-list-seq-rule=%v/ace-rule/destination-choice/any-case/dst-any", data.getPath(), strings.Join(keyValues[:], ",")))
		}
		if !data.Entries[i].Ack.ValueBool() {
			emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/access-list-seq-rule=%v/ace-rule/ack", data.getPath(), strings.Join(keyValues[:], ",")))
		}
		if !data.Entries[i].Fin.ValueBool() {
			emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/access-list-seq-rule=%v/ace-rule/fin", data.getPath(), strings.Join(keyValues[:], ",")))
		}
		if !data.Entries[i].Psh.ValueBool() {
			emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/access-list-seq-rule=%v/ace-rule/psh", data.getPath(), strings.Join(keyValues[:], ",")))
		}
		if !data.Entries[i].Rst.ValueBool() {
			emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/access-list-seq-rule=%v/ace-rule/rst", data.getPath(), strings.Join(keyValues[:], ",")))
		}
		if !data.Entries[i].Syn.ValueBool() {
			emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/access-list-seq-rule=%v/ace-rule/syn", data.getPath(), strings.Join(keyValues[:], ",")))
		}
		if !data.Entries[i].Urg.ValueBool() {
			emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/access-list-seq-rule=%v/ace-rule/urg", data.getPath(), strings.Join(keyValues[:], ",")))
		}
		if !data.Entries[i].Established.ValueBool() {
			emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/access-list-seq-rule=%v/ace-rule/established", data.getPath(), strings.Join(keyValues[:], ",")))
		}
		if !data.Entries[i].Fragments.ValueBool() {
			emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/access-list-seq-rule=%v/ace-rule/fragments", data.getPath(), strings.Join(keyValues[:], ",")))
		}
	}
	return emptyLeafsDelete
}
