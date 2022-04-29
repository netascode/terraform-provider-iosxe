package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/terraform-provider-iosxe/internal/provider/helpers"
	"github.com/tidwall/sjson"
)

type Restconf struct {
	Device     types.String `tfsdk:"device"`
	Id         types.String `tfsdk:"id"`
	Path       types.String `tfsdk:"path"`
	Delete     types.Bool   `tfsdk:"delete"`
	Attributes types.Map    `tfsdk:"attributes"`
}

type RestconfDataSource struct {
	Device     types.String `tfsdk:"device"`
	Id         types.String `tfsdk:"id"`
	Path       types.String `tfsdk:"path"`
	Attributes types.Map    `tfsdk:"attributes"`
}

func (data Restconf) toBody(ctx context.Context) string {
	body := `{"` + helpers.LastElement(data.Path.Value) + `":{}}`

	var attributes map[string]string
	data.Attributes.ElementsAs(ctx, &attributes, false)

	for attr, value := range attributes {
		body, _ = sjson.Set(body, helpers.LastElement(data.Path.Value)+"."+attr, value)
	}

	return body
}
