package helpers

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const (
	floatRangeValidatorErr = "Value must be between [%v] and [%v]"
)

type floatRangeValidator struct {
	min float64
	max float64
}

func (v floatRangeValidator) Description(context.Context) string {
	return fmt.Sprintf(floatRangeValidatorErr, v.min, v.max)
}

func (v floatRangeValidator) MarkdownDescription(context.Context) string {
	return fmt.Sprintf(floatRangeValidatorErr, v.min, v.max)
}

func (v floatRangeValidator) Validate(ctx context.Context, req tfsdk.ValidateAttributeRequest, resp *tfsdk.ValidateAttributeResponse) {
	var float types.Float64
	{
		diags := tfsdk.ValueAs(ctx, req.AttributeConfig, &float)
		resp.Diagnostics.Append(diags...)
		if diags.HasError() {
			return
		}
	}

	if float.Unknown || float.Null {
		return
	}

	if float.Value >= v.min && float.Value <= v.max {
		return
	}

	resp.Diagnostics.AddAttributeError(
		req.AttributePath,
		"Invalid value",
		fmt.Sprintf(floatRangeValidatorErr, v.min, v.max),
	)
}

func FloatRangeValidator(min, max float64) tfsdk.AttributeValidator {
	return floatRangeValidator{
		min: min,
		max: max,
	}
}
