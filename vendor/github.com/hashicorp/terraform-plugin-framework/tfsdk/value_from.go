package tfsdk

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/internal/reflect"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

// ValueFrom takes the Go value `val` and populates `target` with an attr.Value,
// based on the type definition provided in `targetType`.
//
// This is achieved using reflection rules provided by the internal/reflect package.
func ValueFrom(ctx context.Context, val interface{}, targetType attr.Type, target interface{}) diag.Diagnostics {
	v, diags := reflect.FromValue(ctx, targetType, val, tftypes.NewAttributePath())
	if diags.HasError() {
		return diags
	}

	return ValueAs(ctx, v, target)
}
