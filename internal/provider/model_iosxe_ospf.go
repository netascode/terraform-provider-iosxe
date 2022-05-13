// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"
	"net/url"
	"reflect"
	"regexp"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/terraform-provider-iosxe/internal/provider/helpers"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type OSPF struct {
	Device                            types.String         `tfsdk:"device"`
	Id                                types.String         `tfsdk:"id"`
	ProcessId                         types.Int64          `tfsdk:"process_id"`
	BfdAllInterfaces                  types.Bool           `tfsdk:"bfd_all_interfaces"`
	DefaultInformationOriginate       types.Bool           `tfsdk:"default_information_originate"`
	DefaultInformationOriginateAlways types.Bool           `tfsdk:"default_information_originate_always"`
	DefaultMetric                     types.Int64          `tfsdk:"default_metric"`
	Distance                          types.Int64          `tfsdk:"distance"`
	DomainTag                         types.Int64          `tfsdk:"domain_tag"`
	MplsLdpAutoconfig                 types.Bool           `tfsdk:"mpls_ldp_autoconfig"`
	MplsLdpSync                       types.Bool           `tfsdk:"mpls_ldp_sync"`
	Neighbor                          []OSPFNeighbor       `tfsdk:"neighbor"`
	Network                           []OSPFNetwork        `tfsdk:"network"`
	Priority                          types.Int64          `tfsdk:"priority"`
	RouterId                          types.String         `tfsdk:"router_id"`
	Shutdown                          types.Bool           `tfsdk:"shutdown"`
	SummaryAddress                    []OSPFSummaryAddress `tfsdk:"summary_address"`
}
type OSPFNeighbor struct {
	Ip       types.String `tfsdk:"ip"`
	Priority types.Int64  `tfsdk:"priority"`
	Cost     types.Int64  `tfsdk:"cost"`
}
type OSPFNetwork struct {
	Ip       types.String `tfsdk:"ip"`
	Wildcard types.String `tfsdk:"wildcard"`
	Area     types.String `tfsdk:"area"`
}
type OSPFSummaryAddress struct {
	Ip   types.String `tfsdk:"ip"`
	Mask types.String `tfsdk:"mask"`
}

func (data OSPF) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/router/Cisco-IOS-XE-ospf:router-ospf/ospf/process-id=%v", url.QueryEscape(fmt.Sprintf("%v", data.ProcessId.Value)))
}

// if last path element has a key -> remove it
func (data OSPF) getPathShort() string {
	path := data.getPath()
	re := regexp.MustCompile(`(.*)=[^\/]*$`)
	matches := re.FindStringSubmatch(path)
	if len(matches) <= 1 {
		return path
	}
	return matches[1]
}

func (data OSPF) toBody() string {
	body := `{"` + helpers.LastElement(data.getPath()) + `":{}}`
	if !data.ProcessId.Null && !data.ProcessId.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"id", strconv.FormatInt(data.ProcessId.Value, 10))
	}
	if !data.BfdAllInterfaces.Null && !data.BfdAllInterfaces.Unknown {
		if data.BfdAllInterfaces.Value {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"bfd.all-interfaces", map[string]string{})
		}
	}
	if !data.DefaultInformationOriginate.Null && !data.DefaultInformationOriginate.Unknown {
		if data.DefaultInformationOriginate.Value {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"default-information.originate", map[string]string{})
		}
	}
	if !data.DefaultInformationOriginateAlways.Null && !data.DefaultInformationOriginateAlways.Unknown {
		if data.DefaultInformationOriginateAlways.Value {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"default-information.originate.always", map[string]string{})
		}
	}
	if !data.DefaultMetric.Null && !data.DefaultMetric.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"default-metric", strconv.FormatInt(data.DefaultMetric.Value, 10))
	}
	if !data.Distance.Null && !data.Distance.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"distance.distance", strconv.FormatInt(data.Distance.Value, 10))
	}
	if !data.DomainTag.Null && !data.DomainTag.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"domain-tag", strconv.FormatInt(data.DomainTag.Value, 10))
	}
	if !data.MplsLdpAutoconfig.Null && !data.MplsLdpAutoconfig.Unknown {
		if data.MplsLdpAutoconfig.Value {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"mpls.ldp.autoconfig", map[string]string{})
		}
	}
	if !data.MplsLdpSync.Null && !data.MplsLdpSync.Unknown {
		if data.MplsLdpSync.Value {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"mpls.ldp.sync", map[string]string{})
		}
	}
	if !data.Priority.Null && !data.Priority.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"priority", strconv.FormatInt(data.Priority.Value, 10))
	}
	if !data.RouterId.Null && !data.RouterId.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"router-id", data.RouterId.Value)
	}
	if !data.Shutdown.Null && !data.Shutdown.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"shutdown", data.Shutdown.Value)
	}
	if len(data.Neighbor) > 0 {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"neighbor", []interface{}{})
		for index, item := range data.Neighbor {
			if !item.Ip.Null && !item.Ip.Unknown {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"neighbor"+"."+strconv.Itoa(index)+"."+"ip", item.Ip.Value)
			}
			if !item.Priority.Null && !item.Priority.Unknown {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"neighbor"+"."+strconv.Itoa(index)+"."+"priority", strconv.FormatInt(item.Priority.Value, 10))
			}
			if !item.Cost.Null && !item.Cost.Unknown {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"neighbor"+"."+strconv.Itoa(index)+"."+"cost", strconv.FormatInt(item.Cost.Value, 10))
			}
		}
	}
	if len(data.Network) > 0 {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"network", []interface{}{})
		for index, item := range data.Network {
			if !item.Ip.Null && !item.Ip.Unknown {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"network"+"."+strconv.Itoa(index)+"."+"ip", item.Ip.Value)
			}
			if !item.Wildcard.Null && !item.Wildcard.Unknown {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"network"+"."+strconv.Itoa(index)+"."+"wildcard", item.Wildcard.Value)
			}
			if !item.Area.Null && !item.Area.Unknown {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"network"+"."+strconv.Itoa(index)+"."+"area", item.Area.Value)
			}
		}
	}
	if len(data.SummaryAddress) > 0 {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"summary-address", []interface{}{})
		for index, item := range data.SummaryAddress {
			if !item.Ip.Null && !item.Ip.Unknown {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"summary-address"+"."+strconv.Itoa(index)+"."+"ip", item.Ip.Value)
			}
			if !item.Mask.Null && !item.Mask.Unknown {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"summary-address"+"."+strconv.Itoa(index)+"."+"mask", item.Mask.Value)
			}
		}
	}
	return body
}

func (data *OSPF) updateFromBody(res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "id"); value.Exists() {
		data.ProcessId.Value = value.Int()
	} else {
		data.ProcessId.Null = true
	}
	if value := res.Get(prefix + "bfd.all-interfaces"); value.Exists() {
		data.BfdAllInterfaces.Value = true
	} else {
		data.BfdAllInterfaces.Value = false
	}
	if value := res.Get(prefix + "default-information.originate"); value.Exists() {
		data.DefaultInformationOriginate.Value = true
	} else {
		data.DefaultInformationOriginate.Value = false
	}
	if value := res.Get(prefix + "default-information.originate.always"); value.Exists() {
		data.DefaultInformationOriginateAlways.Value = true
	} else {
		data.DefaultInformationOriginateAlways.Value = false
	}
	if value := res.Get(prefix + "default-metric"); value.Exists() {
		data.DefaultMetric.Value = value.Int()
	} else {
		data.DefaultMetric.Null = true
	}
	if value := res.Get(prefix + "distance.distance"); value.Exists() {
		data.Distance.Value = value.Int()
	} else {
		data.Distance.Null = true
	}
	if value := res.Get(prefix + "domain-tag"); value.Exists() {
		data.DomainTag.Value = value.Int()
	} else {
		data.DomainTag.Null = true
	}
	if value := res.Get(prefix + "mpls.ldp.autoconfig"); value.Exists() {
		data.MplsLdpAutoconfig.Value = true
	} else {
		data.MplsLdpAutoconfig.Value = false
	}
	if value := res.Get(prefix + "mpls.ldp.sync"); value.Exists() {
		data.MplsLdpSync.Value = true
	} else {
		data.MplsLdpSync.Value = false
	}
	for i := range data.Neighbor {
		key := data.Neighbor[i].Ip.Value
		if value := res.Get(fmt.Sprintf("%vneighbor.#(ip==\"%v\").ip", prefix, key)); value.Exists() {
			data.Neighbor[i].Ip.Value = value.String()
		} else {
			data.Neighbor[i].Ip.Null = true
		}
		if value := res.Get(fmt.Sprintf("%vneighbor.#(ip==\"%v\").priority", prefix, key)); value.Exists() {
			data.Neighbor[i].Priority.Value = value.Int()
		} else {
			data.Neighbor[i].Priority.Null = true
		}
		if value := res.Get(fmt.Sprintf("%vneighbor.#(ip==\"%v\").cost", prefix, key)); value.Exists() {
			data.Neighbor[i].Cost.Value = value.Int()
		} else {
			data.Neighbor[i].Cost.Null = true
		}
	}
	for i := range data.Network {
		key := data.Network[i].Ip.Value
		if value := res.Get(fmt.Sprintf("%vnetwork.#(ip==\"%v\").ip", prefix, key)); value.Exists() {
			data.Network[i].Ip.Value = value.String()
		} else {
			data.Network[i].Ip.Null = true
		}
		if value := res.Get(fmt.Sprintf("%vnetwork.#(ip==\"%v\").wildcard", prefix, key)); value.Exists() {
			data.Network[i].Wildcard.Value = value.String()
		} else {
			data.Network[i].Wildcard.Null = true
		}
		if value := res.Get(fmt.Sprintf("%vnetwork.#(ip==\"%v\").area", prefix, key)); value.Exists() {
			data.Network[i].Area.Value = value.String()
		} else {
			data.Network[i].Area.Null = true
		}
	}
	if value := res.Get(prefix + "priority"); value.Exists() {
		data.Priority.Value = value.Int()
	} else {
		data.Priority.Null = true
	}
	if value := res.Get(prefix + "router-id"); value.Exists() {
		data.RouterId.Value = value.String()
	} else {
		data.RouterId.Null = true
	}
	if value := res.Get(prefix + "shutdown"); value.Exists() {
		data.Shutdown.Value = value.Bool()
	} else {
		data.Shutdown.Value = false
	}
	for i := range data.SummaryAddress {
		key := data.SummaryAddress[i].Ip.Value
		if value := res.Get(fmt.Sprintf("%vsummary-address.#(ip==\"%v\").ip", prefix, key)); value.Exists() {
			data.SummaryAddress[i].Ip.Value = value.String()
		} else {
			data.SummaryAddress[i].Ip.Null = true
		}
		if value := res.Get(fmt.Sprintf("%vsummary-address.#(ip==\"%v\").mask", prefix, key)); value.Exists() {
			data.SummaryAddress[i].Mask.Value = value.String()
		} else {
			data.SummaryAddress[i].Mask.Null = true
		}
	}
}

func (data *OSPF) fromBody(res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "bfd.all-interfaces"); value.Exists() {
		data.BfdAllInterfaces.Value = true
		data.BfdAllInterfaces.Null = false
	} else {
		data.BfdAllInterfaces.Value = false
		data.BfdAllInterfaces.Null = false
	}
	if value := res.Get(prefix + "default-information.originate"); value.Exists() {
		data.DefaultInformationOriginate.Value = true
		data.DefaultInformationOriginate.Null = false
	} else {
		data.DefaultInformationOriginate.Value = false
		data.DefaultInformationOriginate.Null = false
	}
	if value := res.Get(prefix + "default-information.originate.always"); value.Exists() {
		data.DefaultInformationOriginateAlways.Value = true
		data.DefaultInformationOriginateAlways.Null = false
	} else {
		data.DefaultInformationOriginateAlways.Value = false
		data.DefaultInformationOriginateAlways.Null = false
	}
	if value := res.Get(prefix + "default-metric"); value.Exists() {
		data.DefaultMetric.Value = value.Int()
		data.DefaultMetric.Null = false
	}
	if value := res.Get(prefix + "distance.distance"); value.Exists() {
		data.Distance.Value = value.Int()
		data.Distance.Null = false
	}
	if value := res.Get(prefix + "domain-tag"); value.Exists() {
		data.DomainTag.Value = value.Int()
		data.DomainTag.Null = false
	}
	if value := res.Get(prefix + "mpls.ldp.autoconfig"); value.Exists() {
		data.MplsLdpAutoconfig.Value = true
		data.MplsLdpAutoconfig.Null = false
	} else {
		data.MplsLdpAutoconfig.Value = false
		data.MplsLdpAutoconfig.Null = false
	}
	if value := res.Get(prefix + "mpls.ldp.sync"); value.Exists() {
		data.MplsLdpSync.Value = true
		data.MplsLdpSync.Null = false
	} else {
		data.MplsLdpSync.Value = false
		data.MplsLdpSync.Null = false
	}
	if value := res.Get(prefix + "neighbor"); value.Exists() {
		data.Neighbor = make([]OSPFNeighbor, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := OSPFNeighbor{}
			if cValue := v.Get("ip"); cValue.Exists() {
				item.Ip.Value = cValue.String()
				item.Ip.Null = false
			}
			if cValue := v.Get("priority"); cValue.Exists() {
				item.Priority.Value = cValue.Int()
				item.Priority.Null = false
			}
			if cValue := v.Get("cost"); cValue.Exists() {
				item.Cost.Value = cValue.Int()
				item.Cost.Null = false
			}
			data.Neighbor = append(data.Neighbor, item)
			return true
		})
	}
	if value := res.Get(prefix + "network"); value.Exists() {
		data.Network = make([]OSPFNetwork, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := OSPFNetwork{}
			if cValue := v.Get("ip"); cValue.Exists() {
				item.Ip.Value = cValue.String()
				item.Ip.Null = false
			}
			if cValue := v.Get("wildcard"); cValue.Exists() {
				item.Wildcard.Value = cValue.String()
				item.Wildcard.Null = false
			}
			if cValue := v.Get("area"); cValue.Exists() {
				item.Area.Value = cValue.String()
				item.Area.Null = false
			}
			data.Network = append(data.Network, item)
			return true
		})
	}
	if value := res.Get(prefix + "priority"); value.Exists() {
		data.Priority.Value = value.Int()
		data.Priority.Null = false
	}
	if value := res.Get(prefix + "router-id"); value.Exists() {
		data.RouterId.Value = value.String()
		data.RouterId.Null = false
	}
	if value := res.Get(prefix + "shutdown"); value.Exists() {
		data.Shutdown.Value = value.Bool()
		data.Shutdown.Null = false
	} else {
		data.Shutdown.Value = false
		data.Shutdown.Null = false
	}
	if value := res.Get(prefix + "summary-address"); value.Exists() {
		data.SummaryAddress = make([]OSPFSummaryAddress, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := OSPFSummaryAddress{}
			if cValue := v.Get("ip"); cValue.Exists() {
				item.Ip.Value = cValue.String()
				item.Ip.Null = false
			}
			if cValue := v.Get("mask"); cValue.Exists() {
				item.Mask.Value = cValue.String()
				item.Mask.Null = false
			}
			data.SummaryAddress = append(data.SummaryAddress, item)
			return true
		})
	}
}

func (data *OSPF) setUnknownValues() {
	if data.Device.Unknown {
		data.Device.Unknown = false
		data.Device.Null = true
	}
	if data.Id.Unknown {
		data.Id.Unknown = false
		data.Id.Null = true
	}
	if data.ProcessId.Unknown {
		data.ProcessId.Unknown = false
		data.ProcessId.Null = true
	}
	if data.BfdAllInterfaces.Unknown {
		data.BfdAllInterfaces.Unknown = false
		data.BfdAllInterfaces.Null = true
	}
	if data.DefaultInformationOriginate.Unknown {
		data.DefaultInformationOriginate.Unknown = false
		data.DefaultInformationOriginate.Null = true
	}
	if data.DefaultInformationOriginateAlways.Unknown {
		data.DefaultInformationOriginateAlways.Unknown = false
		data.DefaultInformationOriginateAlways.Null = true
	}
	if data.DefaultMetric.Unknown {
		data.DefaultMetric.Unknown = false
		data.DefaultMetric.Null = true
	}
	if data.Distance.Unknown {
		data.Distance.Unknown = false
		data.Distance.Null = true
	}
	if data.DomainTag.Unknown {
		data.DomainTag.Unknown = false
		data.DomainTag.Null = true
	}
	if data.MplsLdpAutoconfig.Unknown {
		data.MplsLdpAutoconfig.Unknown = false
		data.MplsLdpAutoconfig.Null = true
	}
	if data.MplsLdpSync.Unknown {
		data.MplsLdpSync.Unknown = false
		data.MplsLdpSync.Null = true
	}
	for i := range data.Neighbor {
		if data.Neighbor[i].Ip.Unknown {
			data.Neighbor[i].Ip.Unknown = false
			data.Neighbor[i].Ip.Null = true
		}
		if data.Neighbor[i].Priority.Unknown {
			data.Neighbor[i].Priority.Unknown = false
			data.Neighbor[i].Priority.Null = true
		}
		if data.Neighbor[i].Cost.Unknown {
			data.Neighbor[i].Cost.Unknown = false
			data.Neighbor[i].Cost.Null = true
		}
	}
	for i := range data.Network {
		if data.Network[i].Ip.Unknown {
			data.Network[i].Ip.Unknown = false
			data.Network[i].Ip.Null = true
		}
		if data.Network[i].Wildcard.Unknown {
			data.Network[i].Wildcard.Unknown = false
			data.Network[i].Wildcard.Null = true
		}
		if data.Network[i].Area.Unknown {
			data.Network[i].Area.Unknown = false
			data.Network[i].Area.Null = true
		}
	}
	if data.Priority.Unknown {
		data.Priority.Unknown = false
		data.Priority.Null = true
	}
	if data.RouterId.Unknown {
		data.RouterId.Unknown = false
		data.RouterId.Null = true
	}
	if data.Shutdown.Unknown {
		data.Shutdown.Unknown = false
		data.Shutdown.Null = true
	}
	for i := range data.SummaryAddress {
		if data.SummaryAddress[i].Ip.Unknown {
			data.SummaryAddress[i].Ip.Unknown = false
			data.SummaryAddress[i].Ip.Null = true
		}
		if data.SummaryAddress[i].Mask.Unknown {
			data.SummaryAddress[i].Mask.Unknown = false
			data.SummaryAddress[i].Mask.Null = true
		}
	}
}

func (data *OSPF) getDeletedListItems(state OSPF) []string {
	deletedListItems := make([]string, 0)
	for _, i := range state.Neighbor {
		if reflect.ValueOf(i.Ip.Value).IsZero() {
			continue
		}
		found := false
		for _, j := range data.Neighbor {
			if i.Ip.Value == j.Ip.Value {
				found = true
			}
		}
		if !found {
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/neighbor=%v", state.getPath(), i.Ip.Value))
		}
	}
	for _, i := range state.Network {
		if reflect.ValueOf(i.Ip.Value).IsZero() {
			continue
		}
		found := false
		for _, j := range data.Network {
			if i.Ip.Value == j.Ip.Value {
				found = true
			}
		}
		if !found {
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/network=%v", state.getPath(), i.Ip.Value))
		}
	}
	for _, i := range state.SummaryAddress {
		if reflect.ValueOf(i.Ip.Value).IsZero() {
			continue
		}
		found := false
		for _, j := range data.SummaryAddress {
			if i.Ip.Value == j.Ip.Value {
				found = true
			}
		}
		if !found {
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/summary-address=%v", state.getPath(), i.Ip.Value))
		}
	}
	return deletedListItems
}

func (data *OSPF) getEmptyLeafsDelete() []string {
	emptyLeafsDelete := make([]string, 0)
	if !data.DefaultInformationOriginateAlways.Value {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/default-information/originate/always", data.getPath()))
	}
	return emptyLeafsDelete
}
