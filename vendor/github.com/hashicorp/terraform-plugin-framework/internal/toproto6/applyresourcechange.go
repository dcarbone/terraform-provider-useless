package toproto6

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/internal/fwserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

// ApplyResourceChangeResponse returns the *tfprotov6.ApplyResourceChangeResponse
// equivalent of a *fwserver.ApplyResourceChangeResponse.
func ApplyResourceChangeResponse(ctx context.Context, fw *fwserver.ApplyResourceChangeResponse) *tfprotov6.ApplyResourceChangeResponse {
	if fw == nil {
		return nil
	}

	proto6 := &tfprotov6.ApplyResourceChangeResponse{
		Diagnostics: Diagnostics(ctx, fw.Diagnostics),
		Private:     fw.Private,
	}

	newState, diags := State(ctx, fw.NewState)

	proto6.Diagnostics = append(proto6.Diagnostics, Diagnostics(ctx, diags)...)
	proto6.NewState = newState

	return proto6
}
