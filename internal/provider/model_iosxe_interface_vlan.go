// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/terraform-provider-iosxe/internal/provider/helpers"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type InterfaceVLAN struct {
	Device          types.String `tfsdk:"device"`
	Id              types.String `tfsdk:"id"`
	Name            types.Int64  `tfsdk:"name"`
	Autostate       types.Bool   `tfsdk:"autostate"`
	Description     types.String `tfsdk:"description"`
	Shutdown        types.Bool   `tfsdk:"shutdown"`
	VrfForwarding   types.String `tfsdk:"vrf_forwarding"`
	Ipv4Address     types.String `tfsdk:"ipv4_address"`
	Ipv4AddressMask types.String `tfsdk:"ipv4_address_mask"`
	Unnumbered      types.String `tfsdk:"unnumbered"`
	PimSparseMode   types.Bool   `tfsdk:"pim_sparse_mode"`
}

func (data InterfaceVLAN) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/interface/Vlan=%v", data.Name.Value)
}

func (data InterfaceVLAN) toBody() string {
	body := `{"` + helpers.LastElement(data.getPath()) + `":{}}`
	if !data.Name.Null && !data.Name.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"name", strconv.FormatInt(data.Name.Value, 10))
	}
	if !data.Autostate.Null && !data.Autostate.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"autostate", data.Autostate.Value)
	}
	if !data.Description.Null && !data.Description.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"description", data.Description.Value)
	}
	if !data.Shutdown.Null && !data.Shutdown.Unknown {
		if data.Shutdown.Value {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"shutdown", map[string]string{})
		}
	}
	if !data.VrfForwarding.Null && !data.VrfForwarding.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"vrf.forwarding", data.VrfForwarding.Value)
	}
	if !data.Ipv4Address.Null && !data.Ipv4Address.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ip.address.primary.address", data.Ipv4Address.Value)
	}
	if !data.Ipv4AddressMask.Null && !data.Ipv4AddressMask.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ip.address.primary.mask", data.Ipv4AddressMask.Value)
	}
	if !data.Unnumbered.Null && !data.Unnumbered.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ip.unnumbered", data.Unnumbered.Value)
	}
	if !data.PimSparseMode.Null && !data.PimSparseMode.Unknown {
		if data.PimSparseMode.Value {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ip.pim.Cisco-IOS-XE-multicast:pim-mode-choice-cfg.sparse-mode", map[string]string{})
		}
	}
	return body
}

func (data *InterfaceVLAN) fromBody(res gjson.Result) {
	if value := res.Get(helpers.LastElement(data.getPath()) + "." + "autostate"); value.Exists() {
		data.Autostate.Value = value.Bool()
	}
	if value := res.Get(helpers.LastElement(data.getPath()) + "." + "description"); value.Exists() {
		data.Description.Value = value.String()
	}
	if value := res.Get(helpers.LastElement(data.getPath()) + "." + "shutdown"); value.Exists() {
		data.Shutdown.Value = true
	}
	if value := res.Get(helpers.LastElement(data.getPath()) + "." + "vrf.forwarding"); value.Exists() {
		data.VrfForwarding.Value = value.String()
	}
	if value := res.Get(helpers.LastElement(data.getPath()) + "." + "ip.address.primary.address"); value.Exists() {
		data.Ipv4Address.Value = value.String()
	}
	if value := res.Get(helpers.LastElement(data.getPath()) + "." + "ip.address.primary.mask"); value.Exists() {
		data.Ipv4AddressMask.Value = value.String()
	}
	if value := res.Get(helpers.LastElement(data.getPath()) + "." + "ip.unnumbered"); value.Exists() {
		data.Unnumbered.Value = value.String()
	}
	if value := res.Get(helpers.LastElement(data.getPath()) + "." + "ip.pim.Cisco-IOS-XE-multicast:pim-mode-choice-cfg.sparse-mode"); value.Exists() {
		data.PimSparseMode.Value = true
	}
}

func (data *InterfaceVLAN) fromPlan(plan InterfaceVLAN) {
	data.Device = plan.Device
	data.Name.Value = plan.Name.Value
}