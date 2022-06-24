package fromproto6

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

// ConfigureProviderRequest returns the *fwserver.ConfigureProviderRequest
// equivalent of a *tfprotov6.ConfigureProviderRequest.
func ConfigureProviderRequest(ctx context.Context, proto6 *tfprotov6.ConfigureProviderRequest, providerSchema *tfsdk.Schema) (*tfsdk.ConfigureProviderRequest, diag.Diagnostics) {
	if proto6 == nil {
		return nil, nil
	}

	fw := &tfsdk.ConfigureProviderRequest{
		TerraformVersion: proto6.TerraformVersion,
	}

	config, diags := Config(ctx, proto6.Config, providerSchema)

	if config != nil {
		fw.Config = *config
	}

	return fw, diags
}
