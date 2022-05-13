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

type InterfaceNVE struct {
	Device                      types.String          `tfsdk:"device"`
	Id                          types.String          `tfsdk:"id"`
	Name                        types.Int64           `tfsdk:"name"`
	Description                 types.String          `tfsdk:"description"`
	Shutdown                    types.Bool            `tfsdk:"shutdown"`
	HostReachabilityProtocolBgp types.Bool            `tfsdk:"host_reachability_protocol_bgp"`
	SourceInterfaceLoopback     types.Int64           `tfsdk:"source_interface_loopback"`
	VniVrfs                     []InterfaceNVEVniVrfs `tfsdk:"vni_vrfs"`
	Vnis                        []InterfaceNVEVnis    `tfsdk:"vnis"`
}
type InterfaceNVEVniVrfs struct {
	VniRange types.String `tfsdk:"vni_range"`
	Vrf      types.String `tfsdk:"vrf"`
}
type InterfaceNVEVnis struct {
	VniRange           types.String `tfsdk:"vni_range"`
	Ipv4MulticastGroup types.String `tfsdk:"ipv4_multicast_group"`
	IngressReplication types.Bool   `tfsdk:"ingress_replication"`
}

func (data InterfaceNVE) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/interface/nve=%v", url.QueryEscape(fmt.Sprintf("%v", data.Name.Value)))
}

// if last path element has a key -> remove it
func (data InterfaceNVE) getPathShort() string {
	path := data.getPath()
	re := regexp.MustCompile(`(.*)=[^\/]*$`)
	matches := re.FindStringSubmatch(path)
	if len(matches) <= 1 {
		return path
	}
	return matches[1]
}

func (data InterfaceNVE) toBody() string {
	body := `{"` + helpers.LastElement(data.getPath()) + `":{}}`
	if !data.Name.Null && !data.Name.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"name", strconv.FormatInt(data.Name.Value, 10))
	}
	if !data.Description.Null && !data.Description.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"description", data.Description.Value)
	}
	if !data.Shutdown.Null && !data.Shutdown.Unknown {
		if data.Shutdown.Value {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"shutdown", map[string]string{})
		}
	}
	if !data.HostReachabilityProtocolBgp.Null && !data.HostReachabilityProtocolBgp.Unknown {
		if data.HostReachabilityProtocolBgp.Value {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"host-reachability.protocol.bgp", map[string]string{})
		}
	}
	if !data.SourceInterfaceLoopback.Null && !data.SourceInterfaceLoopback.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"source-interface.Loopback", strconv.FormatInt(data.SourceInterfaceLoopback.Value, 10))
	}
	if len(data.VniVrfs) > 0 {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"member-in-one-line.member.vni", []interface{}{})
		for index, item := range data.VniVrfs {
			if !item.VniRange.Null && !item.VniRange.Unknown {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"member-in-one-line.member.vni"+"."+strconv.Itoa(index)+"."+"vni-range", item.VniRange.Value)
			}
			if !item.Vrf.Null && !item.Vrf.Unknown {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"member-in-one-line.member.vni"+"."+strconv.Itoa(index)+"."+"vrf", item.Vrf.Value)
			}
		}
	}
	if len(data.Vnis) > 0 {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"member.vni", []interface{}{})
		for index, item := range data.Vnis {
			if !item.VniRange.Null && !item.VniRange.Unknown {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"member.vni"+"."+strconv.Itoa(index)+"."+"vni-range", item.VniRange.Value)
			}
			if !item.Ipv4MulticastGroup.Null && !item.Ipv4MulticastGroup.Unknown {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"member.vni"+"."+strconv.Itoa(index)+"."+"mcast-group.multicast-group-min", item.Ipv4MulticastGroup.Value)
			}
			if !item.IngressReplication.Null && !item.IngressReplication.Unknown {
				if item.IngressReplication.Value {
					body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"member.vni"+"."+strconv.Itoa(index)+"."+"ir-cp-config.ingress-replication", map[string]string{})
				}
			}
		}
	}
	return body
}

func (data *InterfaceNVE) updateFromBody(res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "name"); value.Exists() {
		data.Name.Value = value.Int()
	} else {
		data.Name.Null = true
	}
	if value := res.Get(prefix + "description"); value.Exists() {
		data.Description.Value = value.String()
	} else {
		data.Description.Null = true
	}
	if value := res.Get(prefix + "shutdown"); value.Exists() {
		data.Shutdown.Value = true
	} else {
		data.Shutdown.Value = false
	}
	if value := res.Get(prefix + "host-reachability.protocol.bgp"); value.Exists() {
		data.HostReachabilityProtocolBgp.Value = true
	} else {
		data.HostReachabilityProtocolBgp.Value = false
	}
	if value := res.Get(prefix + "source-interface.Loopback"); value.Exists() {
		data.SourceInterfaceLoopback.Value = value.Int()
	} else {
		data.SourceInterfaceLoopback.Null = true
	}
	for i := range data.VniVrfs {
		key := data.VniVrfs[i].VniRange.Value
		if value := res.Get(fmt.Sprintf("%vmember-in-one-line.member.vni.#(vni-range==\"%v\").vni-range", prefix, key)); value.Exists() {
			data.VniVrfs[i].VniRange.Value = value.String()
		} else {
			data.VniVrfs[i].VniRange.Null = true
		}
		if value := res.Get(fmt.Sprintf("%vmember-in-one-line.member.vni.#(vni-range==\"%v\").vrf", prefix, key)); value.Exists() {
			data.VniVrfs[i].Vrf.Value = value.String()
		} else {
			data.VniVrfs[i].Vrf.Null = true
		}
	}
	for i := range data.Vnis {
		key := data.Vnis[i].VniRange.Value
		if value := res.Get(fmt.Sprintf("%vmember.vni.#(vni-range==\"%v\").vni-range", prefix, key)); value.Exists() {
			data.Vnis[i].VniRange.Value = value.String()
		} else {
			data.Vnis[i].VniRange.Null = true
		}
		if value := res.Get(fmt.Sprintf("%vmember.vni.#(vni-range==\"%v\").mcast-group.multicast-group-min", prefix, key)); value.Exists() {
			data.Vnis[i].Ipv4MulticastGroup.Value = value.String()
		} else {
			data.Vnis[i].Ipv4MulticastGroup.Null = true
		}
		if value := res.Get(fmt.Sprintf("%vmember.vni.#(vni-range==\"%v\").ir-cp-config.ingress-replication", prefix, key)); value.Exists() {
			data.Vnis[i].IngressReplication.Value = true
		} else {
			data.Vnis[i].IngressReplication.Value = false
		}
	}
}

func (data *InterfaceNVE) fromBody(res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "description"); value.Exists() {
		data.Description.Value = value.String()
		data.Description.Null = false
	}
	if value := res.Get(prefix + "shutdown"); value.Exists() {
		data.Shutdown.Value = true
		data.Shutdown.Null = false
	} else {
		data.Shutdown.Value = false
		data.Shutdown.Null = false
	}
	if value := res.Get(prefix + "host-reachability.protocol.bgp"); value.Exists() {
		data.HostReachabilityProtocolBgp.Value = true
		data.HostReachabilityProtocolBgp.Null = false
	} else {
		data.HostReachabilityProtocolBgp.Value = false
		data.HostReachabilityProtocolBgp.Null = false
	}
	if value := res.Get(prefix + "source-interface.Loopback"); value.Exists() {
		data.SourceInterfaceLoopback.Value = value.Int()
		data.SourceInterfaceLoopback.Null = false
	}
	if value := res.Get(prefix + "member-in-one-line.member.vni"); value.Exists() {
		data.VniVrfs = make([]InterfaceNVEVniVrfs, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := InterfaceNVEVniVrfs{}
			if cValue := v.Get("vni-range"); cValue.Exists() {
				item.VniRange.Value = cValue.String()
				item.VniRange.Null = false
			}
			if cValue := v.Get("vrf"); cValue.Exists() {
				item.Vrf.Value = cValue.String()
				item.Vrf.Null = false
			}
			data.VniVrfs = append(data.VniVrfs, item)
			return true
		})
	}
	if value := res.Get(prefix + "member.vni"); value.Exists() {
		data.Vnis = make([]InterfaceNVEVnis, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := InterfaceNVEVnis{}
			if cValue := v.Get("vni-range"); cValue.Exists() {
				item.VniRange.Value = cValue.String()
				item.VniRange.Null = false
			}
			if cValue := v.Get("mcast-group.multicast-group-min"); cValue.Exists() {
				item.Ipv4MulticastGroup.Value = cValue.String()
				item.Ipv4MulticastGroup.Null = false
			}
			if cValue := v.Get("ir-cp-config.ingress-replication"); cValue.Exists() {
				item.IngressReplication.Value = true
				item.IngressReplication.Null = false
			}
			data.Vnis = append(data.Vnis, item)
			return true
		})
	}
}

func (data *InterfaceNVE) setUnknownValues() {
	if data.Device.Unknown {
		data.Device.Unknown = false
		data.Device.Null = true
	}
	if data.Id.Unknown {
		data.Id.Unknown = false
		data.Id.Null = true
	}
	if data.Name.Unknown {
		data.Name.Unknown = false
		data.Name.Null = true
	}
	if data.Description.Unknown {
		data.Description.Unknown = false
		data.Description.Null = true
	}
	if data.Shutdown.Unknown {
		data.Shutdown.Unknown = false
		data.Shutdown.Null = true
	}
	if data.HostReachabilityProtocolBgp.Unknown {
		data.HostReachabilityProtocolBgp.Unknown = false
		data.HostReachabilityProtocolBgp.Null = true
	}
	if data.SourceInterfaceLoopback.Unknown {
		data.SourceInterfaceLoopback.Unknown = false
		data.SourceInterfaceLoopback.Null = true
	}
	for i := range data.VniVrfs {
		if data.VniVrfs[i].VniRange.Unknown {
			data.VniVrfs[i].VniRange.Unknown = false
			data.VniVrfs[i].VniRange.Null = true
		}
		if data.VniVrfs[i].Vrf.Unknown {
			data.VniVrfs[i].Vrf.Unknown = false
			data.VniVrfs[i].Vrf.Null = true
		}
	}
	for i := range data.Vnis {
		if data.Vnis[i].VniRange.Unknown {
			data.Vnis[i].VniRange.Unknown = false
			data.Vnis[i].VniRange.Null = true
		}
		if data.Vnis[i].Ipv4MulticastGroup.Unknown {
			data.Vnis[i].Ipv4MulticastGroup.Unknown = false
			data.Vnis[i].Ipv4MulticastGroup.Null = true
		}
		if data.Vnis[i].IngressReplication.Unknown {
			data.Vnis[i].IngressReplication.Unknown = false
			data.Vnis[i].IngressReplication.Null = true
		}
	}
}

func (data *InterfaceNVE) getDeletedListItems(state InterfaceNVE) []string {
	deletedListItems := make([]string, 0)
	for _, i := range state.VniVrfs {
		if reflect.ValueOf(i.VniRange.Value).IsZero() {
			continue
		}
		found := false
		for _, j := range data.VniVrfs {
			if i.VniRange.Value == j.VniRange.Value {
				found = true
			}
		}
		if !found {
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/member-in-one-line/member/vni=%v", state.getPath(), i.VniRange.Value))
		}
	}
	for _, i := range state.Vnis {
		if reflect.ValueOf(i.VniRange.Value).IsZero() {
			continue
		}
		found := false
		for _, j := range data.Vnis {
			if i.VniRange.Value == j.VniRange.Value {
				found = true
			}
		}
		if !found {
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/member/vni=%v", state.getPath(), i.VniRange.Value))
		}
	}
	return deletedListItems
}

func (data *InterfaceNVE) getEmptyLeafsDelete() []string {
	emptyLeafsDelete := make([]string, 0)
	if !data.Shutdown.Value {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/shutdown", data.getPath()))
	}
	if !data.HostReachabilityProtocolBgp.Value {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/host-reachability/protocol/bgp", data.getPath()))
	}
	return emptyLeafsDelete
}
