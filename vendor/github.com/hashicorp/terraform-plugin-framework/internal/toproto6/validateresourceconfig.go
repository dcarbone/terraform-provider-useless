package toproto6

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/internal/fwserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

// ValidateResourceConfigResponse returns the *tfprotov6.ValidateResourceConfigResponse
// equivalent of a *fwserver.ValidateResourceConfigResponse.
func ValidateResourceConfigResponse(ctx context.Context, fw *fwserver.ValidateResourceConfigResponse) *tfprotov6.ValidateResourceConfigResponse {
	if fw == nil {
		return nil
	}

	proto6 := &tfprotov6.ValidateResourceConfigResponse{
		Diagnostics: Diagnostics(fw.Diagnostics),
	}

	return proto6
}
