// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-restconf"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &SNMPServerUserDataSource{}
	_ datasource.DataSourceWithConfigure = &SNMPServerUserDataSource{}
)

func NewSNMPServerUserDataSource() datasource.DataSource {
	return &SNMPServerUserDataSource{}
}

type SNMPServerUserDataSource struct {
	clients map[string]*restconf.Client
}

func (d *SNMPServerUserDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_snmp_server_user"
}

func (d *SNMPServerUserDataSource) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the SNMP Server User configuration.",

		Attributes: map[string]tfsdk.Attribute{
			"device": {
				MarkdownDescription: "A device name from the provider configuration.",
				Type:                types.StringType,
				Optional:            true,
			},
			"id": {
				MarkdownDescription: "The path of the retrieved object.",
				Type:                types.StringType,
				Computed:            true,
			},
			"username": {
				MarkdownDescription: "Name of the user",
				Type:                types.StringType,
				Required:            true,
			},
			"grpname": {
				MarkdownDescription: "Group to which the user belongs",
				Type:                types.StringType,
				Required:            true,
			},
			"v3_auth_algorithm": {
				MarkdownDescription: "Use HMAC SHA/MD5 algorithm for authentication",
				Type:                types.StringType,
				Computed:            true,
			},
			"v3_auth_password": {
				MarkdownDescription: "Authentication password for user",
				Type:                types.StringType,
				Computed:            true,
			},
			"v3_auth_priv_aes_algorithm": {
				MarkdownDescription: "",
				Type:                types.StringType,
				Computed:            true,
			},
			"v3_auth_priv_aes_password": {
				MarkdownDescription: "Authentication password for user",
				Type:                types.StringType,
				Computed:            true,
			},
			"v3_auth_priv_aes_access_ipv6_acl": {
				MarkdownDescription: "Specify IPv6 Named Access-List",
				Type:                types.StringType,
				Computed:            true,
			},
			"v3_auth_priv_aes_access_standard_acl": {
				MarkdownDescription: "Standard IP Access-list allowing access with this community string",
				Type:                types.Int64Type,
				Computed:            true,
			},
			"v3_auth_priv_aes_access_acl_name": {
				MarkdownDescription: "Access-list name",
				Type:                types.StringType,
				Computed:            true,
			},
			"v3_auth_priv_des_password": {
				MarkdownDescription: "Authentication password for user",
				Type:                types.StringType,
				Computed:            true,
			},
			"v3_auth_priv_des_access_ipv6_acl": {
				MarkdownDescription: "Specify IPv6 Named Access-List",
				Type:                types.StringType,
				Computed:            true,
			},
			"v3_auth_priv_des_access_standard_acl": {
				MarkdownDescription: "Standard IP Access-list allowing access with this community string",
				Type:                types.Int64Type,
				Computed:            true,
			},
			"v3_auth_priv_des_access_acl_name": {
				MarkdownDescription: "Access-list name",
				Type:                types.StringType,
				Computed:            true,
			},
			"v3_auth_priv_des3_password": {
				MarkdownDescription: "Authentication password for user",
				Type:                types.StringType,
				Computed:            true,
			},
			"v3_auth_priv_des3_access_ipv6_acl": {
				MarkdownDescription: "Specify IPv6 Named Access-List",
				Type:                types.StringType,
				Computed:            true,
			},
			"v3_auth_priv_des3_access_standard_acl": {
				MarkdownDescription: "Standard IP Access-list allowing access with this community string",
				Type:                types.Int64Type,
				Computed:            true,
			},
			"v3_auth_priv_des3_access_acl_name": {
				MarkdownDescription: "Access-list name",
				Type:                types.StringType,
				Computed:            true,
			},
			"v3_auth_access_ipv6_acl": {
				MarkdownDescription: "Specify IPv6 Named Access-List",
				Type:                types.StringType,
				Computed:            true,
			},
			"v3_auth_access_standard_acl": {
				MarkdownDescription: "Standard IP Access-list allowing access with this community string",
				Type:                types.Int64Type,
				Computed:            true,
			},
			"v3_auth_access_acl_name": {
				MarkdownDescription: "Access-list name",
				Type:                types.StringType,
				Computed:            true,
			},
		},
	}, nil
}

func (d *SNMPServerUserDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.clients = req.ProviderData.(map[string]*restconf.Client)
}

func (d *SNMPServerUserDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config SNMPServerUser

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.getPath()))

	res, err := d.clients[config.Device.ValueString()].GetData(config.getPath())
	if res.StatusCode == 404 {
		config = SNMPServerUser{Device: config.Device}
	} else {
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
			return
		}

		config.fromBody(ctx, res.Res)
	}

	config.Id = types.StringValue(config.getPath())

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.getPath()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}
