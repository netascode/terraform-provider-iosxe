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

type SNMPServerUser struct {
	Device                          types.String `tfsdk:"device"`
	Id                              types.String `tfsdk:"id"`
	Username                        types.String `tfsdk:"username"`
	Grpname                         types.String `tfsdk:"grpname"`
	V3AuthAlgorithm                 types.String `tfsdk:"v3_auth_algorithm"`
	V3AuthPassword                  types.String `tfsdk:"v3_auth_password"`
	V3AuthPrivAesAlgorithm          types.String `tfsdk:"v3_auth_priv_aes_algorithm"`
	V3AuthPrivAesPassword           types.String `tfsdk:"v3_auth_priv_aes_password"`
	V3AuthPrivAesAccessIpv6Acl      types.String `tfsdk:"v3_auth_priv_aes_access_ipv6_acl"`
	V3AuthPrivAesAccessStandardAcl  types.Int64  `tfsdk:"v3_auth_priv_aes_access_standard_acl"`
	V3AuthPrivAesAccessAclName      types.String `tfsdk:"v3_auth_priv_aes_access_acl_name"`
	V3AuthPrivDesPassword           types.String `tfsdk:"v3_auth_priv_des_password"`
	V3AuthPrivDesAccessIpv6Acl      types.String `tfsdk:"v3_auth_priv_des_access_ipv6_acl"`
	V3AuthPrivDesAccessStandardAcl  types.Int64  `tfsdk:"v3_auth_priv_des_access_standard_acl"`
	V3AuthPrivDesAccessAclName      types.String `tfsdk:"v3_auth_priv_des_access_acl_name"`
	V3AuthPrivDes3Password          types.String `tfsdk:"v3_auth_priv_des3_password"`
	V3AuthPrivDes3AccessIpv6Acl     types.String `tfsdk:"v3_auth_priv_des3_access_ipv6_acl"`
	V3AuthPrivDes3AccessStandardAcl types.Int64  `tfsdk:"v3_auth_priv_des3_access_standard_acl"`
	V3AuthPrivDes3AccessAclName     types.String `tfsdk:"v3_auth_priv_des3_access_acl_name"`
	V3AuthAccessIpv6Acl             types.String `tfsdk:"v3_auth_access_ipv6_acl"`
	V3AuthAccessStandardAcl         types.Int64  `tfsdk:"v3_auth_access_standard_acl"`
	V3AuthAccessAclName             types.String `tfsdk:"v3_auth_access_acl_name"`
}

func (data SNMPServerUser) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/snmp-server/Cisco-IOS-XE-snmp:user/names=%s,%s", url.QueryEscape(fmt.Sprintf("%v", data.Username.Value)), url.QueryEscape(fmt.Sprintf("%v", data.Grpname.Value)))
}

// if last path element has a key -> remove it
func (data SNMPServerUser) getPathShort() string {
	path := data.getPath()
	re := regexp.MustCompile(`(.*)=[^\/]*$`)
	matches := re.FindStringSubmatch(path)
	if len(matches) <= 1 {
		return path
	}
	return matches[1]
}

func (data SNMPServerUser) toBody(ctx context.Context) string {
	body := `{"` + helpers.LastElement(data.getPath()) + `":{}}`
	if !data.Username.Null && !data.Username.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"username", data.Username.Value)
	}
	if !data.Grpname.Null && !data.Grpname.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"grpname", data.Grpname.Value)
	}
	if !data.V3AuthAlgorithm.Null && !data.V3AuthAlgorithm.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"v3.auth-config.algorithm", data.V3AuthAlgorithm.Value)
	}
	if !data.V3AuthPassword.Null && !data.V3AuthPassword.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"v3.auth-config.password", data.V3AuthPassword.Value)
	}
	if !data.V3AuthPrivAesAlgorithm.Null && !data.V3AuthPrivAesAlgorithm.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"v3.auth-config.priv-config.aes.algorithm", data.V3AuthPrivAesAlgorithm.Value)
	}
	if !data.V3AuthPrivAesPassword.Null && !data.V3AuthPrivAesPassword.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"v3.auth-config.priv-config.aes.password", data.V3AuthPrivAesPassword.Value)
	}
	if !data.V3AuthPrivAesAccessIpv6Acl.Null && !data.V3AuthPrivAesAccessIpv6Acl.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"v3.auth-config.priv-config.aes.access-config.ipv6-acl", data.V3AuthPrivAesAccessIpv6Acl.Value)
	}
	if !data.V3AuthPrivAesAccessStandardAcl.Null && !data.V3AuthPrivAesAccessStandardAcl.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"v3.auth-config.priv-config.aes.access-config.standard-acl", strconv.FormatInt(data.V3AuthPrivAesAccessStandardAcl.Value, 10))
	}
	if !data.V3AuthPrivAesAccessAclName.Null && !data.V3AuthPrivAesAccessAclName.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"v3.auth-config.priv-config.aes.access-config.acl-name", data.V3AuthPrivAesAccessAclName.Value)
	}
	if !data.V3AuthPrivDesPassword.Null && !data.V3AuthPrivDesPassword.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"v3.auth-config.priv-config.des.password", data.V3AuthPrivDesPassword.Value)
	}
	if !data.V3AuthPrivDesAccessIpv6Acl.Null && !data.V3AuthPrivDesAccessIpv6Acl.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"v3.auth-config.priv-config.des.access-config.ipv6-acl", data.V3AuthPrivDesAccessIpv6Acl.Value)
	}
	if !data.V3AuthPrivDesAccessStandardAcl.Null && !data.V3AuthPrivDesAccessStandardAcl.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"v3.auth-config.priv-config.des.access-config.standard-acl", strconv.FormatInt(data.V3AuthPrivDesAccessStandardAcl.Value, 10))
	}
	if !data.V3AuthPrivDesAccessAclName.Null && !data.V3AuthPrivDesAccessAclName.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"v3.auth-config.priv-config.des.access-config.acl-name", data.V3AuthPrivDesAccessAclName.Value)
	}
	if !data.V3AuthPrivDes3Password.Null && !data.V3AuthPrivDes3Password.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"v3.auth-config.priv-config.des3.password", data.V3AuthPrivDes3Password.Value)
	}
	if !data.V3AuthPrivDes3AccessIpv6Acl.Null && !data.V3AuthPrivDes3AccessIpv6Acl.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"v3.auth-config.priv-config.des3.access-config.ipv6-acl", data.V3AuthPrivDes3AccessIpv6Acl.Value)
	}
	if !data.V3AuthPrivDes3AccessStandardAcl.Null && !data.V3AuthPrivDes3AccessStandardAcl.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"v3.auth-config.priv-config.des3.access-config.standard-acl", strconv.FormatInt(data.V3AuthPrivDes3AccessStandardAcl.Value, 10))
	}
	if !data.V3AuthPrivDes3AccessAclName.Null && !data.V3AuthPrivDes3AccessAclName.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"v3.auth-config.priv-config.des3.access-config.acl-name", data.V3AuthPrivDes3AccessAclName.Value)
	}
	if !data.V3AuthAccessIpv6Acl.Null && !data.V3AuthAccessIpv6Acl.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"v3.auth-config.access-config.ipv6-acl", data.V3AuthAccessIpv6Acl.Value)
	}
	if !data.V3AuthAccessStandardAcl.Null && !data.V3AuthAccessStandardAcl.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"v3.auth-config.access-config.standard-acl", strconv.FormatInt(data.V3AuthAccessStandardAcl.Value, 10))
	}
	if !data.V3AuthAccessAclName.Null && !data.V3AuthAccessAclName.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"v3.auth-config.access-config.acl-name", data.V3AuthAccessAclName.Value)
	}
	return body
}

func (data *SNMPServerUser) updateFromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "username"); value.Exists() {
		data.Username.Value = value.String()
	} else {
		data.Username.Null = true
	}
	if value := res.Get(prefix + "grpname"); value.Exists() {
		data.Grpname.Value = value.String()
	} else {
		data.Grpname.Null = true
	}
	if value := res.Get(prefix + "v3.auth-config.algorithm"); value.Exists() {
		data.V3AuthAlgorithm.Value = value.String()
	} else {
		data.V3AuthAlgorithm.Null = true
	}
	if value := res.Get(prefix + "v3.auth-config.password"); value.Exists() {
		data.V3AuthPassword.Value = value.String()
	} else {
		data.V3AuthPassword.Null = true
	}
	if value := res.Get(prefix + "v3.auth-config.priv-config.aes.algorithm"); value.Exists() {
		data.V3AuthPrivAesAlgorithm.Value = value.String()
	} else {
		data.V3AuthPrivAesAlgorithm.Null = true
	}
	if value := res.Get(prefix + "v3.auth-config.priv-config.aes.password"); value.Exists() {
		data.V3AuthPrivAesPassword.Value = value.String()
	} else {
		data.V3AuthPrivAesPassword.Null = true
	}
	if value := res.Get(prefix + "v3.auth-config.priv-config.aes.access-config.ipv6-acl"); value.Exists() {
		data.V3AuthPrivAesAccessIpv6Acl.Value = value.String()
	} else {
		data.V3AuthPrivAesAccessIpv6Acl.Null = true
	}
	if value := res.Get(prefix + "v3.auth-config.priv-config.aes.access-config.standard-acl"); value.Exists() {
		data.V3AuthPrivAesAccessStandardAcl.Value = value.Int()
	} else {
		data.V3AuthPrivAesAccessStandardAcl.Null = true
	}
	if value := res.Get(prefix + "v3.auth-config.priv-config.aes.access-config.acl-name"); value.Exists() {
		data.V3AuthPrivAesAccessAclName.Value = value.String()
	} else {
		data.V3AuthPrivAesAccessAclName.Null = true
	}
	if value := res.Get(prefix + "v3.auth-config.priv-config.des.password"); value.Exists() {
		data.V3AuthPrivDesPassword.Value = value.String()
	} else {
		data.V3AuthPrivDesPassword.Null = true
	}
	if value := res.Get(prefix + "v3.auth-config.priv-config.des.access-config.ipv6-acl"); value.Exists() {
		data.V3AuthPrivDesAccessIpv6Acl.Value = value.String()
	} else {
		data.V3AuthPrivDesAccessIpv6Acl.Null = true
	}
	if value := res.Get(prefix + "v3.auth-config.priv-config.des.access-config.standard-acl"); value.Exists() {
		data.V3AuthPrivDesAccessStandardAcl.Value = value.Int()
	} else {
		data.V3AuthPrivDesAccessStandardAcl.Null = true
	}
	if value := res.Get(prefix + "v3.auth-config.priv-config.des.access-config.acl-name"); value.Exists() {
		data.V3AuthPrivDesAccessAclName.Value = value.String()
	} else {
		data.V3AuthPrivDesAccessAclName.Null = true
	}
	if value := res.Get(prefix + "v3.auth-config.priv-config.des3.password"); value.Exists() {
		data.V3AuthPrivDes3Password.Value = value.String()
	} else {
		data.V3AuthPrivDes3Password.Null = true
	}
	if value := res.Get(prefix + "v3.auth-config.priv-config.des3.access-config.ipv6-acl"); value.Exists() {
		data.V3AuthPrivDes3AccessIpv6Acl.Value = value.String()
	} else {
		data.V3AuthPrivDes3AccessIpv6Acl.Null = true
	}
	if value := res.Get(prefix + "v3.auth-config.priv-config.des3.access-config.standard-acl"); value.Exists() {
		data.V3AuthPrivDes3AccessStandardAcl.Value = value.Int()
	} else {
		data.V3AuthPrivDes3AccessStandardAcl.Null = true
	}
	if value := res.Get(prefix + "v3.auth-config.priv-config.des3.access-config.acl-name"); value.Exists() {
		data.V3AuthPrivDes3AccessAclName.Value = value.String()
	} else {
		data.V3AuthPrivDes3AccessAclName.Null = true
	}
	if value := res.Get(prefix + "v3.auth-config.access-config.ipv6-acl"); value.Exists() {
		data.V3AuthAccessIpv6Acl.Value = value.String()
	} else {
		data.V3AuthAccessIpv6Acl.Null = true
	}
	if value := res.Get(prefix + "v3.auth-config.access-config.standard-acl"); value.Exists() {
		data.V3AuthAccessStandardAcl.Value = value.Int()
	} else {
		data.V3AuthAccessStandardAcl.Null = true
	}
	if value := res.Get(prefix + "v3.auth-config.access-config.acl-name"); value.Exists() {
		data.V3AuthAccessAclName.Value = value.String()
	} else {
		data.V3AuthAccessAclName.Null = true
	}
}

func (data *SNMPServerUser) fromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "v3.auth-config.algorithm"); value.Exists() {
		data.V3AuthAlgorithm.Value = value.String()
		data.V3AuthAlgorithm.Null = false
	}
	if value := res.Get(prefix + "v3.auth-config.password"); value.Exists() {
		data.V3AuthPassword.Value = value.String()
		data.V3AuthPassword.Null = false
	}
	if value := res.Get(prefix + "v3.auth-config.priv-config.aes.algorithm"); value.Exists() {
		data.V3AuthPrivAesAlgorithm.Value = value.String()
		data.V3AuthPrivAesAlgorithm.Null = false
	}
	if value := res.Get(prefix + "v3.auth-config.priv-config.aes.password"); value.Exists() {
		data.V3AuthPrivAesPassword.Value = value.String()
		data.V3AuthPrivAesPassword.Null = false
	}
	if value := res.Get(prefix + "v3.auth-config.priv-config.aes.access-config.ipv6-acl"); value.Exists() {
		data.V3AuthPrivAesAccessIpv6Acl.Value = value.String()
		data.V3AuthPrivAesAccessIpv6Acl.Null = false
	}
	if value := res.Get(prefix + "v3.auth-config.priv-config.aes.access-config.standard-acl"); value.Exists() {
		data.V3AuthPrivAesAccessStandardAcl.Value = value.Int()
		data.V3AuthPrivAesAccessStandardAcl.Null = false
	}
	if value := res.Get(prefix + "v3.auth-config.priv-config.aes.access-config.acl-name"); value.Exists() {
		data.V3AuthPrivAesAccessAclName.Value = value.String()
		data.V3AuthPrivAesAccessAclName.Null = false
	}
	if value := res.Get(prefix + "v3.auth-config.priv-config.des.password"); value.Exists() {
		data.V3AuthPrivDesPassword.Value = value.String()
		data.V3AuthPrivDesPassword.Null = false
	}
	if value := res.Get(prefix + "v3.auth-config.priv-config.des.access-config.ipv6-acl"); value.Exists() {
		data.V3AuthPrivDesAccessIpv6Acl.Value = value.String()
		data.V3AuthPrivDesAccessIpv6Acl.Null = false
	}
	if value := res.Get(prefix + "v3.auth-config.priv-config.des.access-config.standard-acl"); value.Exists() {
		data.V3AuthPrivDesAccessStandardAcl.Value = value.Int()
		data.V3AuthPrivDesAccessStandardAcl.Null = false
	}
	if value := res.Get(prefix + "v3.auth-config.priv-config.des.access-config.acl-name"); value.Exists() {
		data.V3AuthPrivDesAccessAclName.Value = value.String()
		data.V3AuthPrivDesAccessAclName.Null = false
	}
	if value := res.Get(prefix + "v3.auth-config.priv-config.des3.password"); value.Exists() {
		data.V3AuthPrivDes3Password.Value = value.String()
		data.V3AuthPrivDes3Password.Null = false
	}
	if value := res.Get(prefix + "v3.auth-config.priv-config.des3.access-config.ipv6-acl"); value.Exists() {
		data.V3AuthPrivDes3AccessIpv6Acl.Value = value.String()
		data.V3AuthPrivDes3AccessIpv6Acl.Null = false
	}
	if value := res.Get(prefix + "v3.auth-config.priv-config.des3.access-config.standard-acl"); value.Exists() {
		data.V3AuthPrivDes3AccessStandardAcl.Value = value.Int()
		data.V3AuthPrivDes3AccessStandardAcl.Null = false
	}
	if value := res.Get(prefix + "v3.auth-config.priv-config.des3.access-config.acl-name"); value.Exists() {
		data.V3AuthPrivDes3AccessAclName.Value = value.String()
		data.V3AuthPrivDes3AccessAclName.Null = false
	}
	if value := res.Get(prefix + "v3.auth-config.access-config.ipv6-acl"); value.Exists() {
		data.V3AuthAccessIpv6Acl.Value = value.String()
		data.V3AuthAccessIpv6Acl.Null = false
	}
	if value := res.Get(prefix + "v3.auth-config.access-config.standard-acl"); value.Exists() {
		data.V3AuthAccessStandardAcl.Value = value.Int()
		data.V3AuthAccessStandardAcl.Null = false
	}
	if value := res.Get(prefix + "v3.auth-config.access-config.acl-name"); value.Exists() {
		data.V3AuthAccessAclName.Value = value.String()
		data.V3AuthAccessAclName.Null = false
	}
}

func (data *SNMPServerUser) setUnknownValues(ctx context.Context) {
	if data.Device.Unknown {
		data.Device.Unknown = false
		data.Device.Null = true
	}
	if data.Id.Unknown {
		data.Id.Unknown = false
		data.Id.Null = true
	}
	if data.Username.Unknown {
		data.Username.Unknown = false
		data.Username.Null = true
	}
	if data.Grpname.Unknown {
		data.Grpname.Unknown = false
		data.Grpname.Null = true
	}
	if data.V3AuthAlgorithm.Unknown {
		data.V3AuthAlgorithm.Unknown = false
		data.V3AuthAlgorithm.Null = true
	}
	if data.V3AuthPassword.Unknown {
		data.V3AuthPassword.Unknown = false
		data.V3AuthPassword.Null = true
	}
	if data.V3AuthPrivAesAlgorithm.Unknown {
		data.V3AuthPrivAesAlgorithm.Unknown = false
		data.V3AuthPrivAesAlgorithm.Null = true
	}
	if data.V3AuthPrivAesPassword.Unknown {
		data.V3AuthPrivAesPassword.Unknown = false
		data.V3AuthPrivAesPassword.Null = true
	}
	if data.V3AuthPrivAesAccessIpv6Acl.Unknown {
		data.V3AuthPrivAesAccessIpv6Acl.Unknown = false
		data.V3AuthPrivAesAccessIpv6Acl.Null = true
	}
	if data.V3AuthPrivAesAccessStandardAcl.Unknown {
		data.V3AuthPrivAesAccessStandardAcl.Unknown = false
		data.V3AuthPrivAesAccessStandardAcl.Null = true
	}
	if data.V3AuthPrivAesAccessAclName.Unknown {
		data.V3AuthPrivAesAccessAclName.Unknown = false
		data.V3AuthPrivAesAccessAclName.Null = true
	}
	if data.V3AuthPrivDesPassword.Unknown {
		data.V3AuthPrivDesPassword.Unknown = false
		data.V3AuthPrivDesPassword.Null = true
	}
	if data.V3AuthPrivDesAccessIpv6Acl.Unknown {
		data.V3AuthPrivDesAccessIpv6Acl.Unknown = false
		data.V3AuthPrivDesAccessIpv6Acl.Null = true
	}
	if data.V3AuthPrivDesAccessStandardAcl.Unknown {
		data.V3AuthPrivDesAccessStandardAcl.Unknown = false
		data.V3AuthPrivDesAccessStandardAcl.Null = true
	}
	if data.V3AuthPrivDesAccessAclName.Unknown {
		data.V3AuthPrivDesAccessAclName.Unknown = false
		data.V3AuthPrivDesAccessAclName.Null = true
	}
	if data.V3AuthPrivDes3Password.Unknown {
		data.V3AuthPrivDes3Password.Unknown = false
		data.V3AuthPrivDes3Password.Null = true
	}
	if data.V3AuthPrivDes3AccessIpv6Acl.Unknown {
		data.V3AuthPrivDes3AccessIpv6Acl.Unknown = false
		data.V3AuthPrivDes3AccessIpv6Acl.Null = true
	}
	if data.V3AuthPrivDes3AccessStandardAcl.Unknown {
		data.V3AuthPrivDes3AccessStandardAcl.Unknown = false
		data.V3AuthPrivDes3AccessStandardAcl.Null = true
	}
	if data.V3AuthPrivDes3AccessAclName.Unknown {
		data.V3AuthPrivDes3AccessAclName.Unknown = false
		data.V3AuthPrivDes3AccessAclName.Null = true
	}
	if data.V3AuthAccessIpv6Acl.Unknown {
		data.V3AuthAccessIpv6Acl.Unknown = false
		data.V3AuthAccessIpv6Acl.Null = true
	}
	if data.V3AuthAccessStandardAcl.Unknown {
		data.V3AuthAccessStandardAcl.Unknown = false
		data.V3AuthAccessStandardAcl.Null = true
	}
	if data.V3AuthAccessAclName.Unknown {
		data.V3AuthAccessAclName.Unknown = false
		data.V3AuthAccessAclName.Null = true
	}
}

func (data *SNMPServerUser) getDeletedListItems(ctx context.Context, state SNMPServerUser) []string {
	deletedListItems := make([]string, 0)
	return deletedListItems
}

func (data *SNMPServerUser) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	return emptyLeafsDelete
}