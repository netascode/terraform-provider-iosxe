// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"
	"regexp"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/terraform-provider-iosxe/internal/provider/helpers"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type InterfacePIM struct {
	Device          types.String `tfsdk:"device"`
	Id              types.String `tfsdk:"id"`
	Type            types.String `tfsdk:"type"`
	Name            types.String `tfsdk:"name"`
	Passive         types.Bool   `tfsdk:"passive"`
	DenseMode       types.Bool   `tfsdk:"dense_mode"`
	SparseMode      types.Bool   `tfsdk:"sparse_mode"`
	SparseDenseMode types.Bool   `tfsdk:"sparse_dense_mode"`
	Bfd             types.Bool   `tfsdk:"bfd"`
	Border          types.Bool   `tfsdk:"border"`
	BsrBorder       types.Bool   `tfsdk:"bsr_border"`
	DrPriority      types.Int64  `tfsdk:"dr_priority"`
}

type InterfacePIMData struct {
	Device          types.String `tfsdk:"device"`
	Id              types.String `tfsdk:"id"`
	Type            types.String `tfsdk:"type"`
	Name            types.String `tfsdk:"name"`
	Passive         types.Bool   `tfsdk:"passive"`
	DenseMode       types.Bool   `tfsdk:"dense_mode"`
	SparseMode      types.Bool   `tfsdk:"sparse_mode"`
	SparseDenseMode types.Bool   `tfsdk:"sparse_dense_mode"`
	Bfd             types.Bool   `tfsdk:"bfd"`
	Border          types.Bool   `tfsdk:"border"`
	BsrBorder       types.Bool   `tfsdk:"bsr_border"`
	DrPriority      types.Int64  `tfsdk:"dr_priority"`
}

func (data InterfacePIM) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/interface/%s=%v/ip/pim", url.QueryEscape(fmt.Sprintf("%v", data.Type.ValueString())), url.QueryEscape(fmt.Sprintf("%v", data.Name.ValueString())))
}

func (data InterfacePIMData) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/interface/%s=%v/ip/pim", url.QueryEscape(fmt.Sprintf("%v", data.Type.ValueString())), url.QueryEscape(fmt.Sprintf("%v", data.Name.ValueString())))
}

// if last path element has a key -> remove it
func (data InterfacePIM) getPathShort() string {
	path := data.getPath()
	re := regexp.MustCompile(`(.*)=[^\/]*$`)
	matches := re.FindStringSubmatch(path)
	if len(matches) <= 1 {
		return path
	}
	return matches[1]
}

func (data InterfacePIM) toBody(ctx context.Context) string {
	body := `{"` + helpers.LastElement(data.getPath()) + `":{}}`
	if !data.Passive.IsNull() && !data.Passive.IsUnknown() {
		if data.Passive.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-multicast:pim-mode-choice-cfg.passive", map[string]string{})
		}
	}
	if !data.DenseMode.IsNull() && !data.DenseMode.IsUnknown() {
		if data.DenseMode.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-multicast:pim-mode-choice-cfg.dense-mode", map[string]string{})
		}
	}
	if !data.SparseMode.IsNull() && !data.SparseMode.IsUnknown() {
		if data.SparseMode.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-multicast:pim-mode-choice-cfg.sparse-mode", map[string]string{})
		}
	}
	if !data.SparseDenseMode.IsNull() && !data.SparseDenseMode.IsUnknown() {
		if data.SparseDenseMode.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-multicast:pim-mode-choice-cfg.sparse-dense-mode", map[string]string{})
		}
	}
	if !data.Bfd.IsNull() && !data.Bfd.IsUnknown() {
		if data.Bfd.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-multicast:bfd", map[string]string{})
		}
	}
	if !data.Border.IsNull() && !data.Border.IsUnknown() {
		if data.Border.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-multicast:border", map[string]string{})
		}
	}
	if !data.BsrBorder.IsNull() && !data.BsrBorder.IsUnknown() {
		if data.BsrBorder.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-multicast:bsr-border", map[string]string{})
		}
	}
	if !data.DrPriority.IsNull() && !data.DrPriority.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-multicast:dr-priority", strconv.FormatInt(data.DrPriority.ValueInt64(), 10))
	}
	return body
}

func (data *InterfacePIM) updateFromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-multicast:pim-mode-choice-cfg.passive"); !data.Passive.IsNull() {
		if value.Exists() {
			data.Passive = types.BoolValue(true)
		} else {
			data.Passive = types.BoolValue(false)
		}
	} else {
		data.Passive = types.BoolNull()
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-multicast:pim-mode-choice-cfg.dense-mode"); !data.DenseMode.IsNull() {
		if value.Exists() {
			data.DenseMode = types.BoolValue(true)
		} else {
			data.DenseMode = types.BoolValue(false)
		}
	} else {
		data.DenseMode = types.BoolNull()
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-multicast:pim-mode-choice-cfg.sparse-mode"); !data.SparseMode.IsNull() {
		if value.Exists() {
			data.SparseMode = types.BoolValue(true)
		} else {
			data.SparseMode = types.BoolValue(false)
		}
	} else {
		data.SparseMode = types.BoolNull()
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-multicast:pim-mode-choice-cfg.sparse-dense-mode"); !data.SparseDenseMode.IsNull() {
		if value.Exists() {
			data.SparseDenseMode = types.BoolValue(true)
		} else {
			data.SparseDenseMode = types.BoolValue(false)
		}
	} else {
		data.SparseDenseMode = types.BoolNull()
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-multicast:bfd"); !data.Bfd.IsNull() {
		if value.Exists() {
			data.Bfd = types.BoolValue(true)
		} else {
			data.Bfd = types.BoolValue(false)
		}
	} else {
		data.Bfd = types.BoolNull()
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-multicast:border"); !data.Border.IsNull() {
		if value.Exists() {
			data.Border = types.BoolValue(true)
		} else {
			data.Border = types.BoolValue(false)
		}
	} else {
		data.Border = types.BoolNull()
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-multicast:bsr-border"); !data.BsrBorder.IsNull() {
		if value.Exists() {
			data.BsrBorder = types.BoolValue(true)
		} else {
			data.BsrBorder = types.BoolValue(false)
		}
	} else {
		data.BsrBorder = types.BoolNull()
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-multicast:dr-priority"); value.Exists() && !data.DrPriority.IsNull() {
		data.DrPriority = types.Int64Value(value.Int())
	} else {
		data.DrPriority = types.Int64Null()
	}
}

func (data *InterfacePIMData) fromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-multicast:pim-mode-choice-cfg.passive"); value.Exists() {
		data.Passive = types.BoolValue(true)
	} else {
		data.Passive = types.BoolValue(false)
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-multicast:pim-mode-choice-cfg.dense-mode"); value.Exists() {
		data.DenseMode = types.BoolValue(true)
	} else {
		data.DenseMode = types.BoolValue(false)
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-multicast:pim-mode-choice-cfg.sparse-mode"); value.Exists() {
		data.SparseMode = types.BoolValue(true)
	} else {
		data.SparseMode = types.BoolValue(false)
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-multicast:pim-mode-choice-cfg.sparse-dense-mode"); value.Exists() {
		data.SparseDenseMode = types.BoolValue(true)
	} else {
		data.SparseDenseMode = types.BoolValue(false)
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-multicast:bfd"); value.Exists() {
		data.Bfd = types.BoolValue(true)
	} else {
		data.Bfd = types.BoolValue(false)
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-multicast:border"); value.Exists() {
		data.Border = types.BoolValue(true)
	} else {
		data.Border = types.BoolValue(false)
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-multicast:bsr-border"); value.Exists() {
		data.BsrBorder = types.BoolValue(true)
	} else {
		data.BsrBorder = types.BoolValue(false)
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-multicast:dr-priority"); value.Exists() {
		data.DrPriority = types.Int64Value(value.Int())
	}
}

func (data *InterfacePIM) getDeletedListItems(ctx context.Context, state InterfacePIM) []string {
	deletedListItems := make([]string, 0)
	return deletedListItems
}

func (data *InterfacePIM) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	if !data.Passive.IsNull() && !data.Passive.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/Cisco-IOS-XE-multicast:pim-mode-choice-cfg/passive", data.getPath()))
	}
	if !data.DenseMode.IsNull() && !data.DenseMode.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/Cisco-IOS-XE-multicast:pim-mode-choice-cfg/dense-mode", data.getPath()))
	}
	if !data.SparseMode.IsNull() && !data.SparseMode.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/Cisco-IOS-XE-multicast:pim-mode-choice-cfg/sparse-mode", data.getPath()))
	}
	if !data.SparseDenseMode.IsNull() && !data.SparseDenseMode.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/Cisco-IOS-XE-multicast:pim-mode-choice-cfg/sparse-dense-mode", data.getPath()))
	}
	if !data.Bfd.IsNull() && !data.Bfd.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/Cisco-IOS-XE-multicast:bfd", data.getPath()))
	}
	if !data.Border.IsNull() && !data.Border.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/Cisco-IOS-XE-multicast:border", data.getPath()))
	}
	if !data.BsrBorder.IsNull() && !data.BsrBorder.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/Cisco-IOS-XE-multicast:bsr-border", data.getPath()))
	}
	return emptyLeafsDelete
}

func (data *InterfacePIM) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	if !data.Passive.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/Cisco-IOS-XE-multicast:pim-mode-choice-cfg/passive", data.getPath()))
	}
	if !data.DenseMode.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/Cisco-IOS-XE-multicast:pim-mode-choice-cfg/dense-mode", data.getPath()))
	}
	if !data.SparseMode.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/Cisco-IOS-XE-multicast:pim-mode-choice-cfg/sparse-mode", data.getPath()))
	}
	if !data.SparseDenseMode.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/Cisco-IOS-XE-multicast:pim-mode-choice-cfg/sparse-dense-mode", data.getPath()))
	}
	if !data.Bfd.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/Cisco-IOS-XE-multicast:bfd", data.getPath()))
	}
	if !data.Border.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/Cisco-IOS-XE-multicast:border", data.getPath()))
	}
	if !data.BsrBorder.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/Cisco-IOS-XE-multicast:bsr-border", data.getPath()))
	}
	if !data.DrPriority.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/Cisco-IOS-XE-multicast:dr-priority", data.getPath()))
	}
	return deletePaths
}
