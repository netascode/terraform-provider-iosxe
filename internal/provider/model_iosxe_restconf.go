package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-restconf"
	"github.com/netascode/terraform-provider-iosxe/internal/provider/helpers"
	"github.com/tidwall/sjson"
)

type Restconf struct {
	Device     types.String   `tfsdk:"device"`
	Id         types.String   `tfsdk:"id"`
	Path       types.String   `tfsdk:"path"`
	Delete     types.Bool     `tfsdk:"delete"`
	Attributes types.Map      `tfsdk:"attributes"`
	Lists      []RestconfList `tfsdk:"lists"`
}

type RestconfList struct {
	Name  types.String       `tfsdk:"name"`
	Items []RestconfListItem `tfsdk:"items"`
}

type RestconfListItem struct {
	Attributes types.Map `tfsdk:"attributes"`
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
	for i := range data.Lists {
		body, _ = sjson.Set(body, helpers.LastElement(data.Path.Value)+"."+data.Lists[i].Name.Value, []interface{}{})
		for ii := range data.Lists[i].Items {
			var listAttributes map[string]string
			data.Lists[i].Items[ii].Attributes.ElementsAs(ctx, &listAttributes, false)
			attrs := restconf.Body{}
			for attr, value := range listAttributes {
				attrs = attrs.Set(attr, value)
			}
			body, _ = sjson.SetRaw(body, helpers.LastElement(data.Path.Value)+"."+data.Lists[i].Name.Value+".-1", attrs.Str)
		}
	}

	return body
}
