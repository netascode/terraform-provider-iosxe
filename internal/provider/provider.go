// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-restconf"
	"github.com/netascode/terraform-provider-iosxe/internal/provider/helpers"
)

// provider satisfies the tfsdk.Provider interface and usually is included
// with all Resource and DataSource implementations.
type provider struct {
	clients map[string]*restconf.Client

	// configured is set to true at the end of the Configure method.
	// This can be used in Resource and DataSource implementations to verify
	// that the provider was previously configured.
	configured bool

	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// providerData can be used to store data from the Terraform configuration.
type providerData struct {
	Username types.String         `tfsdk:"username"`
	Password types.String         `tfsdk:"password"`
	URL      types.String         `tfsdk:"url"`
	Insecure types.Bool           `tfsdk:"insecure"`
	Retries  types.Int64          `tfsdk:"retries"`
	Devices  []providerDataDevice `tfsdk:"devices"`
}

type providerDataDevice struct {
	Name types.String `tfsdk:"name"`
	URL  types.String `tfsdk:"url"`
}

func (p *provider) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"username": {
				MarkdownDescription: "Username for the IOS-XE device. This can also be set as the IOSXE_USERNAME environment variable.",
				Type:                types.StringType,
				Optional:            true,
			},
			"password": {
				MarkdownDescription: "Password for the IOS-XE device. This can also be set as the IOSXE_PASSWORD environment variable.",
				Type:                types.StringType,
				Optional:            true,
				Sensitive:           true,
			},
			"url": {
				MarkdownDescription: "URL of the Cisco IOS-XE device. Optionally a port can be added with `:12345`. The default port is `443`. This can also be set as the IOSXE_URL environment variable.",
				Type:                types.StringType,
				Optional:            true,
			},
			"insecure": {
				MarkdownDescription: "Allow insecure HTTPS client. This can also be set as the IOSXE_INSECURE environment variable. Defaults to `true`.",
				Type:                types.BoolType,
				Optional:            true,
			},
			"retries": {
				MarkdownDescription: "Number of retries for REST API calls. This can also be set as the IOSXE_RETRIES environment variable. Defaults to `2`.",
				Type:                types.Int64Type,
				Optional:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.IntegerRangeValidator(0, 9),
				},
			},
			"devices": {
				MarkdownDescription: "This can be used to manage a list of devices from a single provider. All devices must use the same credentials. Each resource and data source has an optional attribute named `device`, which can then select a device by its name from this list.",
				Optional:            true,
				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
					"name": {
						MarkdownDescription: "Device name.",
						Type:                types.StringType,
						Required:            true,
					},
					"url": {
						MarkdownDescription: "URL of the Cisco IOS-XE device.",
						Type:                types.StringType,
						Required:            true,
					},
				}, tfsdk.ListNestedAttributesOptions{}),
			},
		},
	}, nil
}

func (p *provider) Configure(ctx context.Context, req tfsdk.ConfigureProviderRequest, resp *tfsdk.ConfigureProviderResponse) {
	// Retrieve provider data from configuration
	var config providerData
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// User must provide a username to the provider
	var username string
	if config.Username.Unknown {
		// Cannot connect to client with an unknown value
		resp.Diagnostics.AddWarning(
			"Unable to create client",
			"Cannot use unknown value as username",
		)
		return
	}

	if config.Username.Null {
		username = os.Getenv("IOSXE_USERNAME")
	} else {
		username = config.Username.Value
	}

	if username == "" {
		// Error vs warning - empty value must stop execution
		resp.Diagnostics.AddError(
			"Unable to find username",
			"Username cannot be an empty string",
		)
		return
	}

	// User must provide a password to the provider
	var password string
	if config.Password.Unknown {
		// Cannot connect to client with an unknown value
		resp.Diagnostics.AddWarning(
			"Unable to create client",
			"Cannot use unknown value as password",
		)
		return
	}

	if config.Password.Null {
		password = os.Getenv("IOSXE_PASSWORD")
	} else {
		password = config.Password.Value
	}

	if password == "" {
		// Error vs warning - empty value must stop execution
		resp.Diagnostics.AddError(
			"Unable to find password",
			"Password cannot be an empty string",
		)
		return
	}

	// User must provide a username to the provider
	var url string
	if config.URL.Unknown {
		// Cannot connect to client with an unknown value
		resp.Diagnostics.AddWarning(
			"Unable to create client",
			"Cannot use unknown value as url",
		)
		return
	}

	if config.URL.Null {
		url = os.Getenv("IOSXE_URL")
	} else {
		url = config.URL.Value
		if url == "" && len(config.Devices) > 0 {
			url = config.Devices[0].URL.Value
		}
	}

	if url == "" {
		// Error vs warning - empty value must stop execution
		resp.Diagnostics.AddError(
			"Unable to find url",
			"URL cannot be an empty string",
		)
		return
	}

	var insecure bool
	if config.Insecure.Unknown {
		// Cannot connect to client with an unknown value
		resp.Diagnostics.AddWarning(
			"Unable to create client",
			"Cannot use unknown value as insecure",
		)
		return
	}

	if config.Insecure.Null {
		insecureStr := os.Getenv("IOSXE_INSECURE")
		if insecureStr == "" {
			insecure = true
		} else {
			insecure, _ = strconv.ParseBool(insecureStr)
		}
	} else {
		insecure = config.Insecure.Value
	}

	var retries int64
	if config.Retries.Unknown {
		// Cannot connect to client with an unknown value
		resp.Diagnostics.AddWarning(
			"Unable to create client",
			"Cannot use unknown value as retries",
		)
		return
	}

	if config.Retries.Null {
		retriesStr := os.Getenv("IOSXE_RETRIES")
		if retriesStr == "" {
			retries = 2
		} else {
			retries, _ = strconv.ParseInt(retriesStr, 0, 64)
		}
	} else {
		retries = config.Retries.Value
	}

	clients := make(map[string]*restconf.Client)
	c, err := restconf.NewClient(url, username, password, insecure, restconf.MaxRetries(int(retries)))
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to create client",
			"Unable to create restconf client:\n\n"+err.Error(),
		)
		return
	}
	clients[""] = &c

	for _, device := range config.Devices {
		c, err := restconf.NewClient(device.URL.Value, username, password, insecure, restconf.MaxRetries(int(retries)))
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to create client",
				"Unable to create restconf client:\n\n"+err.Error(),
			)
			return
		}
		clients[device.Name.Value] = &c
	}

	p.clients = clients
	p.configured = true
}

func (p *provider) GetResources(ctx context.Context) (map[string]tfsdk.ResourceType, diag.Diagnostics) {
	return map[string]tfsdk.ResourceType{
		"iosxe_restconf":                    resourceRestconfType{},
		"iosxe_banner":                      resourceBannerType{},
		"iosxe_bgp":                         resourceBGPType{},
		"iosxe_bgp_address_family_ipv4_vrf": resourceBGPAddressFamilyIPv4VRFType{},
		"iosxe_bgp_address_family_ipv6_vrf": resourceBGPAddressFamilyIPv6VRFType{},
		"iosxe_evpn":                        resourceEVPNType{},
		"iosxe_evpn_instance":               resourceEVPNInstanceType{},
		"iosxe_interface_loopback":          resourceInterfaceLoopbackType{},
		"iosxe_interface_nve":               resourceInterfaceNVEType{},
		"iosxe_interface_vlan":              resourceInterfaceVLANType{},
		"iosxe_ospf":                        resourceOSPFType{},
		"iosxe_ospf_vrf":                    resourceOSPFVRFType{},
		"iosxe_static_route":                resourceStaticRouteType{},
		"iosxe_system":                      resourceSystemType{},
		"iosxe_username":                    resourceUsernameType{},
		"iosxe_vlan":                        resourceVLANType{},
		"iosxe_vlan_configuration":          resourceVLANConfigurationType{},
		"iosxe_vrf":                         resourceVRFType{},
	}, nil
}

func (p *provider) GetDataSources(ctx context.Context) (map[string]tfsdk.DataSourceType, diag.Diagnostics) {
	return map[string]tfsdk.DataSourceType{
		"iosxe_restconf":                    dataSourceRestconfType{},
		"iosxe_banner":                      dataSourceBannerType{},
		"iosxe_bgp":                         dataSourceBGPType{},
		"iosxe_bgp_address_family_ipv4_vrf": dataSourceBGPAddressFamilyIPv4VRFType{},
		"iosxe_bgp_address_family_ipv6_vrf": dataSourceBGPAddressFamilyIPv6VRFType{},
		"iosxe_evpn":                        dataSourceEVPNType{},
		"iosxe_evpn_instance":               dataSourceEVPNInstanceType{},
		"iosxe_interface_loopback":          dataSourceInterfaceLoopbackType{},
		"iosxe_interface_nve":               dataSourceInterfaceNVEType{},
		"iosxe_interface_vlan":              dataSourceInterfaceVLANType{},
		"iosxe_ospf":                        dataSourceOSPFType{},
		"iosxe_ospf_vrf":                    dataSourceOSPFVRFType{},
		"iosxe_static_route":                dataSourceStaticRouteType{},
		"iosxe_system":                      dataSourceSystemType{},
		"iosxe_username":                    dataSourceUsernameType{},
		"iosxe_vlan":                        dataSourceVLANType{},
		"iosxe_vlan_configuration":          dataSourceVLANConfigurationType{},
		"iosxe_vrf":                         dataSourceVRFType{},
	}, nil
}

func New(version string) func() tfsdk.Provider {
	return func() tfsdk.Provider {
		return &provider{
			version: version,
		}
	}
}

// convertProviderType is a helper function for NewResource and NewDataSource
// implementations to associate the concrete provider type. Alternatively,
// this helper can be skipped and the provider type can be directly type
// asserted (e.g. provider: in.(*provider)), however using this can prevent
// potential panics.
func convertProviderType(in tfsdk.Provider) (provider, diag.Diagnostics) {
	var diags diag.Diagnostics

	p, ok := in.(*provider)

	if !ok {
		diags.AddError(
			"Unexpected Provider Instance Type",
			fmt.Sprintf("While creating the data source or resource, an unexpected provider type (%T) was received. This is always a bug in the provider code and should be reported to the provider developers.", p),
		)
		return provider{}, diags
	}

	if p == nil {
		diags.AddError(
			"Unexpected Provider Instance Type",
			"While creating the data source or resource, an unexpected empty provider instance was received. This is always a bug in the provider code and should be reported to the provider developers.",
		)
		return provider{}, diags
	}

	return *p, diags
}
