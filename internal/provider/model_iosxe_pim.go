// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

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

type PIM struct {
	Device                        types.String      `tfsdk:"device"`
	Id                            types.String      `tfsdk:"id"`
	DeleteMode                    types.String      `tfsdk:"delete_mode"`
	Autorp                        types.Bool        `tfsdk:"autorp"`
	AutorpListener                types.Bool        `tfsdk:"autorp_listener"`
	BsrCandidateLoopback          types.Int64       `tfsdk:"bsr_candidate_loopback"`
	BsrCandidateMask              types.Int64       `tfsdk:"bsr_candidate_mask"`
	BsrCandidatePriority          types.Int64       `tfsdk:"bsr_candidate_priority"`
	BsrCandidateAcceptRpCandidate types.String      `tfsdk:"bsr_candidate_accept_rp_candidate"`
	SsmRange                      types.String      `tfsdk:"ssm_range"`
	SsmDefault                    types.Bool        `tfsdk:"ssm_default"`
	RpAddress                     types.String      `tfsdk:"rp_address"`
	RpAddressOverride             types.Bool        `tfsdk:"rp_address_override"`
	RpAddressBidir                types.Bool        `tfsdk:"rp_address_bidir"`
	RpAddresses                   []PIMRpAddresses  `tfsdk:"rp_addresses"`
	RpCandidates                  []PIMRpCandidates `tfsdk:"rp_candidates"`
}

type PIMData struct {
	Device                        types.String      `tfsdk:"device"`
	Id                            types.String      `tfsdk:"id"`
	Autorp                        types.Bool        `tfsdk:"autorp"`
	AutorpListener                types.Bool        `tfsdk:"autorp_listener"`
	BsrCandidateLoopback          types.Int64       `tfsdk:"bsr_candidate_loopback"`
	BsrCandidateMask              types.Int64       `tfsdk:"bsr_candidate_mask"`
	BsrCandidatePriority          types.Int64       `tfsdk:"bsr_candidate_priority"`
	BsrCandidateAcceptRpCandidate types.String      `tfsdk:"bsr_candidate_accept_rp_candidate"`
	SsmRange                      types.String      `tfsdk:"ssm_range"`
	SsmDefault                    types.Bool        `tfsdk:"ssm_default"`
	RpAddress                     types.String      `tfsdk:"rp_address"`
	RpAddressOverride             types.Bool        `tfsdk:"rp_address_override"`
	RpAddressBidir                types.Bool        `tfsdk:"rp_address_bidir"`
	RpAddresses                   []PIMRpAddresses  `tfsdk:"rp_addresses"`
	RpCandidates                  []PIMRpCandidates `tfsdk:"rp_candidates"`
}
type PIMRpAddresses struct {
	AccessList types.String `tfsdk:"access_list"`
	RpAddress  types.String `tfsdk:"rp_address"`
	Override   types.Bool   `tfsdk:"override"`
	Bidir      types.Bool   `tfsdk:"bidir"`
}
type PIMRpCandidates struct {
	Interface types.String `tfsdk:"interface"`
	GroupList types.String `tfsdk:"group_list"`
	Interval  types.Int64  `tfsdk:"interval"`
	Priority  types.Int64  `tfsdk:"priority"`
	Bidir     types.Bool   `tfsdk:"bidir"`
}

func (data PIM) getPath() string {
	return "Cisco-IOS-XE-native:native/ip/pim"
}

func (data PIMData) getPath() string {
	return "Cisco-IOS-XE-native:native/ip/pim"
}

// if last path element has a key -> remove it
func (data PIM) getPathShort() string {
	path := data.getPath()
	re := regexp.MustCompile(`(.*)=[^\/]*$`)
	matches := re.FindStringSubmatch(path)
	if len(matches) <= 1 {
		return path
	}
	return matches[1]
}

func (data PIM) toBody(ctx context.Context) string {
	body := `{"` + helpers.LastElement(data.getPath()) + `":{}}`
	if !data.Autorp.IsNull() && !data.Autorp.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-multicast:autorp-container.autorp", data.Autorp.ValueBool())
	}
	if !data.AutorpListener.IsNull() && !data.AutorpListener.IsUnknown() {
		if data.AutorpListener.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-multicast:autorp-container.listener", map[string]string{})
		}
	}
	if !data.BsrCandidateLoopback.IsNull() && !data.BsrCandidateLoopback.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-multicast:bsr-candidate.Loopback", strconv.FormatInt(data.BsrCandidateLoopback.ValueInt64(), 10))
	}
	if !data.BsrCandidateMask.IsNull() && !data.BsrCandidateMask.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-multicast:bsr-candidate.mask", strconv.FormatInt(data.BsrCandidateMask.ValueInt64(), 10))
	}
	if !data.BsrCandidatePriority.IsNull() && !data.BsrCandidatePriority.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-multicast:bsr-candidate.priority", strconv.FormatInt(data.BsrCandidatePriority.ValueInt64(), 10))
	}
	if !data.BsrCandidateAcceptRpCandidate.IsNull() && !data.BsrCandidateAcceptRpCandidate.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-multicast:bsr-candidate.accept-rp-candidate", data.BsrCandidateAcceptRpCandidate.ValueString())
	}
	if !data.SsmRange.IsNull() && !data.SsmRange.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-multicast:ssm.range", data.SsmRange.ValueString())
	}
	if !data.SsmDefault.IsNull() && !data.SsmDefault.IsUnknown() {
		if data.SsmDefault.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-multicast:ssm.default", map[string]string{})
		}
	}
	if !data.RpAddress.IsNull() && !data.RpAddress.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-multicast:rp-address-conf.address", data.RpAddress.ValueString())
	}
	if !data.RpAddressOverride.IsNull() && !data.RpAddressOverride.IsUnknown() {
		if data.RpAddressOverride.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-multicast:rp-address-conf.override", map[string]string{})
		}
	}
	if !data.RpAddressBidir.IsNull() && !data.RpAddressBidir.IsUnknown() {
		if data.RpAddressBidir.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-multicast:rp-address-conf.bidir", map[string]string{})
		}
	}
	if len(data.RpAddresses) > 0 {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-multicast:rp-address-list", []interface{}{})
		for index, item := range data.RpAddresses {
			if !item.AccessList.IsNull() && !item.AccessList.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-multicast:rp-address-list"+"."+strconv.Itoa(index)+"."+"access-list", item.AccessList.ValueString())
			}
			if !item.RpAddress.IsNull() && !item.RpAddress.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-multicast:rp-address-list"+"."+strconv.Itoa(index)+"."+"rp-address", item.RpAddress.ValueString())
			}
			if !item.Override.IsNull() && !item.Override.IsUnknown() {
				if item.Override.ValueBool() {
					body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-multicast:rp-address-list"+"."+strconv.Itoa(index)+"."+"override", map[string]string{})
				}
			}
			if !item.Bidir.IsNull() && !item.Bidir.IsUnknown() {
				if item.Bidir.ValueBool() {
					body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-multicast:rp-address-list"+"."+strconv.Itoa(index)+"."+"bidir", map[string]string{})
				}
			}
		}
	}
	if len(data.RpCandidates) > 0 {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-multicast:rp-candidate", []interface{}{})
		for index, item := range data.RpCandidates {
			if !item.Interface.IsNull() && !item.Interface.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-multicast:rp-candidate"+"."+strconv.Itoa(index)+"."+"interface", item.Interface.ValueString())
			}
			if !item.GroupList.IsNull() && !item.GroupList.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-multicast:rp-candidate"+"."+strconv.Itoa(index)+"."+"group-list", item.GroupList.ValueString())
			}
			if !item.Interval.IsNull() && !item.Interval.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-multicast:rp-candidate"+"."+strconv.Itoa(index)+"."+"interval", strconv.FormatInt(item.Interval.ValueInt64(), 10))
			}
			if !item.Priority.IsNull() && !item.Priority.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-multicast:rp-candidate"+"."+strconv.Itoa(index)+"."+"priority", strconv.FormatInt(item.Priority.ValueInt64(), 10))
			}
			if !item.Bidir.IsNull() && !item.Bidir.IsUnknown() {
				if item.Bidir.ValueBool() {
					body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-multicast:rp-candidate"+"."+strconv.Itoa(index)+"."+"bidir", map[string]string{})
				}
			}
		}
	}
	return body
}

func (data *PIM) updateFromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-multicast:autorp-container.autorp"); !data.Autorp.IsNull() {
		if value.Exists() {
			data.Autorp = types.BoolValue(value.Bool())
		}
	} else {
		data.Autorp = types.BoolNull()
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-multicast:autorp-container.listener"); !data.AutorpListener.IsNull() {
		if value.Exists() {
			data.AutorpListener = types.BoolValue(true)
		} else {
			data.AutorpListener = types.BoolValue(false)
		}
	} else {
		data.AutorpListener = types.BoolNull()
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-multicast:bsr-candidate.Loopback"); value.Exists() && !data.BsrCandidateLoopback.IsNull() {
		data.BsrCandidateLoopback = types.Int64Value(value.Int())
	} else {
		data.BsrCandidateLoopback = types.Int64Null()
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-multicast:bsr-candidate.mask"); value.Exists() && !data.BsrCandidateMask.IsNull() {
		data.BsrCandidateMask = types.Int64Value(value.Int())
	} else {
		data.BsrCandidateMask = types.Int64Null()
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-multicast:bsr-candidate.priority"); value.Exists() && !data.BsrCandidatePriority.IsNull() {
		data.BsrCandidatePriority = types.Int64Value(value.Int())
	} else {
		data.BsrCandidatePriority = types.Int64Null()
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-multicast:bsr-candidate.accept-rp-candidate"); value.Exists() && !data.BsrCandidateAcceptRpCandidate.IsNull() {
		data.BsrCandidateAcceptRpCandidate = types.StringValue(value.String())
	} else {
		data.BsrCandidateAcceptRpCandidate = types.StringNull()
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-multicast:ssm.range"); value.Exists() && !data.SsmRange.IsNull() {
		data.SsmRange = types.StringValue(value.String())
	} else {
		data.SsmRange = types.StringNull()
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-multicast:ssm.default"); !data.SsmDefault.IsNull() {
		if value.Exists() {
			data.SsmDefault = types.BoolValue(true)
		} else {
			data.SsmDefault = types.BoolValue(false)
		}
	} else {
		data.SsmDefault = types.BoolNull()
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-multicast:rp-address-conf.address"); value.Exists() && !data.RpAddress.IsNull() {
		data.RpAddress = types.StringValue(value.String())
	} else {
		data.RpAddress = types.StringNull()
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-multicast:rp-address-conf.override"); !data.RpAddressOverride.IsNull() {
		if value.Exists() {
			data.RpAddressOverride = types.BoolValue(true)
		} else {
			data.RpAddressOverride = types.BoolValue(false)
		}
	} else {
		data.RpAddressOverride = types.BoolNull()
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-multicast:rp-address-conf.bidir"); !data.RpAddressBidir.IsNull() {
		if value.Exists() {
			data.RpAddressBidir = types.BoolValue(true)
		} else {
			data.RpAddressBidir = types.BoolValue(false)
		}
	} else {
		data.RpAddressBidir = types.BoolNull()
	}
	for i := range data.RpAddresses {
		keys := [...]string{"access-list"}
		keyValues := [...]string{data.RpAddresses[i].AccessList.ValueString()}

		var r gjson.Result
		res.Get(prefix + "Cisco-IOS-XE-multicast:rp-address-list").ForEach(
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
		if value := r.Get("access-list"); value.Exists() && !data.RpAddresses[i].AccessList.IsNull() {
			data.RpAddresses[i].AccessList = types.StringValue(value.String())
		} else {
			data.RpAddresses[i].AccessList = types.StringNull()
		}
		if value := r.Get("rp-address"); value.Exists() && !data.RpAddresses[i].RpAddress.IsNull() {
			data.RpAddresses[i].RpAddress = types.StringValue(value.String())
		} else {
			data.RpAddresses[i].RpAddress = types.StringNull()
		}
		if value := r.Get("override"); !data.RpAddresses[i].Override.IsNull() {
			if value.Exists() {
				data.RpAddresses[i].Override = types.BoolValue(true)
			} else {
				data.RpAddresses[i].Override = types.BoolValue(false)
			}
		} else {
			data.RpAddresses[i].Override = types.BoolNull()
		}
		if value := r.Get("bidir"); !data.RpAddresses[i].Bidir.IsNull() {
			if value.Exists() {
				data.RpAddresses[i].Bidir = types.BoolValue(true)
			} else {
				data.RpAddresses[i].Bidir = types.BoolValue(false)
			}
		} else {
			data.RpAddresses[i].Bidir = types.BoolNull()
		}
	}
	for i := range data.RpCandidates {
		keys := [...]string{"interface"}
		keyValues := [...]string{data.RpCandidates[i].Interface.ValueString()}

		var r gjson.Result
		res.Get(prefix + "Cisco-IOS-XE-multicast:rp-candidate").ForEach(
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
		if value := r.Get("interface"); value.Exists() && !data.RpCandidates[i].Interface.IsNull() {
			data.RpCandidates[i].Interface = types.StringValue(value.String())
		} else {
			data.RpCandidates[i].Interface = types.StringNull()
		}
		if value := r.Get("group-list"); value.Exists() && !data.RpCandidates[i].GroupList.IsNull() {
			data.RpCandidates[i].GroupList = types.StringValue(value.String())
		} else {
			data.RpCandidates[i].GroupList = types.StringNull()
		}
		if value := r.Get("interval"); value.Exists() && !data.RpCandidates[i].Interval.IsNull() {
			data.RpCandidates[i].Interval = types.Int64Value(value.Int())
		} else {
			data.RpCandidates[i].Interval = types.Int64Null()
		}
		if value := r.Get("priority"); value.Exists() && !data.RpCandidates[i].Priority.IsNull() {
			data.RpCandidates[i].Priority = types.Int64Value(value.Int())
		} else {
			data.RpCandidates[i].Priority = types.Int64Null()
		}
		if value := r.Get("bidir"); !data.RpCandidates[i].Bidir.IsNull() {
			if value.Exists() {
				data.RpCandidates[i].Bidir = types.BoolValue(true)
			} else {
				data.RpCandidates[i].Bidir = types.BoolValue(false)
			}
		} else {
			data.RpCandidates[i].Bidir = types.BoolNull()
		}
	}
}

func (data *PIMData) fromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-multicast:autorp-container.autorp"); value.Exists() {
		data.Autorp = types.BoolValue(value.Bool())
	} else {
		data.Autorp = types.BoolValue(false)
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-multicast:autorp-container.listener"); value.Exists() {
		data.AutorpListener = types.BoolValue(true)
	} else {
		data.AutorpListener = types.BoolValue(false)
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-multicast:bsr-candidate.Loopback"); value.Exists() {
		data.BsrCandidateLoopback = types.Int64Value(value.Int())
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-multicast:bsr-candidate.mask"); value.Exists() {
		data.BsrCandidateMask = types.Int64Value(value.Int())
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-multicast:bsr-candidate.priority"); value.Exists() {
		data.BsrCandidatePriority = types.Int64Value(value.Int())
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-multicast:bsr-candidate.accept-rp-candidate"); value.Exists() {
		data.BsrCandidateAcceptRpCandidate = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-multicast:ssm.range"); value.Exists() {
		data.SsmRange = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-multicast:ssm.default"); value.Exists() {
		data.SsmDefault = types.BoolValue(true)
	} else {
		data.SsmDefault = types.BoolValue(false)
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-multicast:rp-address-conf.address"); value.Exists() {
		data.RpAddress = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-multicast:rp-address-conf.override"); value.Exists() {
		data.RpAddressOverride = types.BoolValue(true)
	} else {
		data.RpAddressOverride = types.BoolValue(false)
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-multicast:rp-address-conf.bidir"); value.Exists() {
		data.RpAddressBidir = types.BoolValue(true)
	} else {
		data.RpAddressBidir = types.BoolValue(false)
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-multicast:rp-address-list"); value.Exists() {
		data.RpAddresses = make([]PIMRpAddresses, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := PIMRpAddresses{}
			if cValue := v.Get("access-list"); cValue.Exists() {
				item.AccessList = types.StringValue(cValue.String())
			}
			if cValue := v.Get("rp-address"); cValue.Exists() {
				item.RpAddress = types.StringValue(cValue.String())
			}
			if cValue := v.Get("override"); cValue.Exists() {
				item.Override = types.BoolValue(true)
			} else {
				item.Override = types.BoolValue(false)
			}
			if cValue := v.Get("bidir"); cValue.Exists() {
				item.Bidir = types.BoolValue(true)
			} else {
				item.Bidir = types.BoolValue(false)
			}
			data.RpAddresses = append(data.RpAddresses, item)
			return true
		})
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-multicast:rp-candidate"); value.Exists() {
		data.RpCandidates = make([]PIMRpCandidates, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := PIMRpCandidates{}
			if cValue := v.Get("interface"); cValue.Exists() {
				item.Interface = types.StringValue(cValue.String())
			}
			if cValue := v.Get("group-list"); cValue.Exists() {
				item.GroupList = types.StringValue(cValue.String())
			}
			if cValue := v.Get("interval"); cValue.Exists() {
				item.Interval = types.Int64Value(cValue.Int())
			}
			if cValue := v.Get("priority"); cValue.Exists() {
				item.Priority = types.Int64Value(cValue.Int())
			}
			if cValue := v.Get("bidir"); cValue.Exists() {
				item.Bidir = types.BoolValue(true)
			} else {
				item.Bidir = types.BoolValue(false)
			}
			data.RpCandidates = append(data.RpCandidates, item)
			return true
		})
	}
}

func (data *PIM) getDeletedListItems(ctx context.Context, state PIM) []string {
	deletedListItems := make([]string, 0)
	for i := range state.RpAddresses {
		stateKeyValues := [...]string{state.RpAddresses[i].AccessList.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.RpAddresses[i].AccessList.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.RpAddresses {
			found = true
			if state.RpAddresses[i].AccessList.ValueString() != data.RpAddresses[j].AccessList.ValueString() {
				found = false
			}
			if found {
				break
			}
		}
		if !found {
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/Cisco-IOS-XE-multicast:rp-address-list=%v", state.getPath(), strings.Join(stateKeyValues[:], ",")))
		}
	}
	for i := range state.RpCandidates {
		stateKeyValues := [...]string{state.RpCandidates[i].Interface.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.RpCandidates[i].Interface.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.RpCandidates {
			found = true
			if state.RpCandidates[i].Interface.ValueString() != data.RpCandidates[j].Interface.ValueString() {
				found = false
			}
			if found {
				break
			}
		}
		if !found {
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/Cisco-IOS-XE-multicast:rp-candidate=%v", state.getPath(), strings.Join(stateKeyValues[:], ",")))
		}
	}
	return deletedListItems
}

func (data *PIM) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	if !data.AutorpListener.IsNull() && !data.AutorpListener.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/Cisco-IOS-XE-multicast:autorp-container/listener", data.getPath()))
	}
	if !data.SsmDefault.IsNull() && !data.SsmDefault.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/Cisco-IOS-XE-multicast:ssm/default", data.getPath()))
	}
	if !data.RpAddressOverride.IsNull() && !data.RpAddressOverride.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/Cisco-IOS-XE-multicast:rp-address-conf/override", data.getPath()))
	}
	if !data.RpAddressBidir.IsNull() && !data.RpAddressBidir.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/Cisco-IOS-XE-multicast:rp-address-conf/bidir", data.getPath()))
	}

	for i := range data.RpAddresses {
		keyValues := [...]string{data.RpAddresses[i].AccessList.ValueString()}
		if !data.RpAddresses[i].Override.IsNull() && !data.RpAddresses[i].Override.ValueBool() {
			emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/Cisco-IOS-XE-multicast:rp-address-list=%v/override", data.getPath(), strings.Join(keyValues[:], ",")))
		}
		if !data.RpAddresses[i].Bidir.IsNull() && !data.RpAddresses[i].Bidir.ValueBool() {
			emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/Cisco-IOS-XE-multicast:rp-address-list=%v/bidir", data.getPath(), strings.Join(keyValues[:], ",")))
		}
	}

	for i := range data.RpCandidates {
		keyValues := [...]string{data.RpCandidates[i].Interface.ValueString()}
		if !data.RpCandidates[i].Bidir.IsNull() && !data.RpCandidates[i].Bidir.ValueBool() {
			emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/Cisco-IOS-XE-multicast:rp-candidate=%v/bidir", data.getPath(), strings.Join(keyValues[:], ",")))
		}
	}
	return emptyLeafsDelete
}

func (data *PIM) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	if !data.Autorp.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/Cisco-IOS-XE-multicast:autorp-container/autorp", data.getPath()))
	}
	if !data.AutorpListener.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/Cisco-IOS-XE-multicast:autorp-container/listener", data.getPath()))
	}
	if !data.BsrCandidateLoopback.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/Cisco-IOS-XE-multicast:bsr-candidate", data.getPath()))
	}
	if !data.BsrCandidateMask.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/Cisco-IOS-XE-multicast:bsr-candidate", data.getPath()))
	}
	if !data.BsrCandidatePriority.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/Cisco-IOS-XE-multicast:bsr-candidate", data.getPath()))
	}
	if !data.BsrCandidateAcceptRpCandidate.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/Cisco-IOS-XE-multicast:bsr-candidate", data.getPath()))
	}
	if !data.SsmRange.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/Cisco-IOS-XE-multicast:ssm/range", data.getPath()))
	}
	if !data.SsmDefault.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/Cisco-IOS-XE-multicast:ssm/default", data.getPath()))
	}
	if !data.RpAddress.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/Cisco-IOS-XE-multicast:rp-address-conf/address", data.getPath()))
	}
	if !data.RpAddressOverride.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/Cisco-IOS-XE-multicast:rp-address-conf/override", data.getPath()))
	}
	if !data.RpAddressBidir.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/Cisco-IOS-XE-multicast:rp-address-conf/bidir", data.getPath()))
	}
	for i := range data.RpAddresses {
		keyValues := [...]string{data.RpAddresses[i].AccessList.ValueString()}

		deletePaths = append(deletePaths, fmt.Sprintf("%v/Cisco-IOS-XE-multicast:rp-address-list=%v", data.getPath(), strings.Join(keyValues[:], ",")))
	}
	for i := range data.RpCandidates {
		keyValues := [...]string{data.RpCandidates[i].Interface.ValueString()}

		deletePaths = append(deletePaths, fmt.Sprintf("%v/Cisco-IOS-XE-multicast:rp-candidate=%v", data.getPath(), strings.Join(keyValues[:], ",")))
	}
	return deletePaths
}
