// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-restconf"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &ClockDataSource{}
	_ datasource.DataSourceWithConfigure = &ClockDataSource{}
)

func NewClockDataSource() datasource.DataSource {
	return &ClockDataSource{}
}

type ClockDataSource struct {
	clients map[string]*restconf.Client
}

func (d *ClockDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_clock"
}

func (d *ClockDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Clock configuration.",

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The path of the retrieved object.",
				Computed:            true,
			},
			"calendar_valid": schema.BoolAttribute{
				MarkdownDescription: "Calendar time is authoritative",
				Computed:            true,
			},
			"summer_time_zone": schema.StringAttribute{
				MarkdownDescription: "Name of time zone in summer",
				Computed:            true,
			},
			"summer_time_date": schema.BoolAttribute{
				MarkdownDescription: "Configure absolute summer time",
				Computed:            true,
			},
			"summer_time_date_start_day": schema.Int64Attribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"summer_time_date_start_month": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"summer_time_date_start_year": schema.Int64Attribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"summer_time_date_start_time": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"summer_time_date_end_day": schema.Int64Attribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"summer_time_date_end_month": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"summer_time_date_end_year": schema.Int64Attribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"summer_time_date_end_time": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"summer_time_date_offset": schema.Int64Attribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"summer_time_recurring": schema.BoolAttribute{
				MarkdownDescription: "Configure recurring summer time",
				Computed:            true,
			},
			"summer_time_recurring_start_week": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"summer_time_recurring_start_weekday": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"summer_time_recurring_start_month": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"summer_time_recurring_start_time": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"summer_time_recurring_end_week": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"summer_time_recurring_end_weekday": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"summer_time_recurring_end_month": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"summer_time_recurring_end_time": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"summer_time_recurring_offset": schema.Int64Attribute{
				MarkdownDescription: "",
				Computed:            true,
			},
		},
	}
}

func (d *ClockDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.clients = req.ProviderData.(map[string]*restconf.Client)
}

func (d *ClockDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config ClockData

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if _, ok := d.clients[config.Device.ValueString()]; !ok {
		resp.Diagnostics.AddAttributeError(path.Root("device"), "Invalid device", fmt.Sprintf("Device '%s' does not exist in provider configuration.", config.Device.ValueString()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.getPath()))

	res, err := d.clients[config.Device.ValueString()].GetData(config.getPath())
	if res.StatusCode == 404 {
		config = ClockData{Device: config.Device}
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
