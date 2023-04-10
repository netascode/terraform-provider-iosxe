package helpers

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
)

func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func LastElement(path string) string {
	pes := strings.Split(path, "/")
	var prefix, element string
	for _, pe := range pes {
		// remove key
		if strings.Contains(pe, "=") {
			pe = pe[:strings.Index(pe, "=")]
		}
		if strings.Contains(pe, ":") {
			prefix = strings.Split(pe, ":")[0]
			element = strings.Split(pe, ":")[1]
		} else {
			element = pe
		}
	}
	return prefix + ":" + element
}

func GetValueSlice(result []gjson.Result) []attr.Value {
	v := make([]attr.Value, len(result))
	for r := range result {
		v[r] = types.StringValue(result[r].String())
	}
	return v
}

func GetStringList(result []gjson.Result) types.List {
	v := make([]attr.Value, len(result))
	for r := range result {
		v[r] = types.StringValue(result[r].String())
	}
	return types.ListValueMust(types.StringType, v)
}

func GetInt64List(result []gjson.Result) types.List {
	v := make([]attr.Value, len(result))
	for r := range result {
		v[r] = types.Int64Value(result[r].Int())
	}
	return types.ListValueMust(types.Int64Type, v)
}
