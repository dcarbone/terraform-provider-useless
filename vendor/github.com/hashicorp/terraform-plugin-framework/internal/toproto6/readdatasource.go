package toproto6

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/internal/fwserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

// ReadDataSourceResponse returns the *tfprotov6.ReadDataSourceResponse
// equivalent of a *fwserver.ReadDataSourceResponse.
func ReadDataSourceResponse(ctx context.Context, fw *fwserver.ReadDataSourceResponse) *tfprotov6.ReadDataSourceResponse {
	if fw == nil {
		return nil
	}

	proto6 := &tfprotov6.ReadDataSourceResponse{
		Diagnostics: Diagnostics(fw.Diagnostics),
	}

	state, diags := State(ctx, fw.State)

	proto6.Diagnostics = append(proto6.Diagnostics, Diagnostics(diags)...)
	proto6.State = state

	return proto6
}
