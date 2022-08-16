package fromproto6

import (
	"context"

	"github.com/hashicorp/terraform-plugin-go/tfprotov6"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/internal/fwschema"
	"github.com/hashicorp/terraform-plugin-framework/internal/fwserver"
	"github.com/hashicorp/terraform-plugin-framework/internal/privatestate"
	"github.com/hashicorp/terraform-plugin-framework/provider"
)

// PlanResourceChangeRequest returns the *fwserver.PlanResourceChangeRequest
// equivalent of a *tfprotov6.PlanResourceChangeRequest.
func PlanResourceChangeRequest(ctx context.Context, proto6 *tfprotov6.PlanResourceChangeRequest, resourceType provider.ResourceType, resourceSchema fwschema.Schema, providerMetaSchema fwschema.Schema) (*fwserver.PlanResourceChangeRequest, diag.Diagnostics) {
	if proto6 == nil {
		return nil, nil
	}

	var diags diag.Diagnostics

	// Panic prevention here to simplify the calling implementations.
	// This should not happen, but just in case.
	if resourceSchema == nil {
		diags.AddError(
			"Missing Resource Schema",
			"An unexpected error was encountered when handling the request. "+
				"This is always an issue in terraform-plugin-framework used to implement the provider and should be reported to the provider developers.\n\n"+
				"Please report this to the provider developer:\n\n"+
				"Missing schema.",
		)

		return nil, diags
	}

	fw := &fwserver.PlanResourceChangeRequest{
		ResourceSchema: resourceSchema,
		ResourceType:   resourceType,
	}

	config, configDiags := Config(ctx, proto6.Config, resourceSchema)

	diags.Append(configDiags...)

	fw.Config = config

	priorState, priorStateDiags := State(ctx, proto6.PriorState, resourceSchema)

	diags.Append(priorStateDiags...)

	fw.PriorState = priorState

	proposedNewState, proposedNewStateDiags := Plan(ctx, proto6.ProposedNewState, resourceSchema)

	diags.Append(proposedNewStateDiags...)

	fw.ProposedNewState = proposedNewState

	providerMeta, providerMetaDiags := ProviderMeta(ctx, proto6.ProviderMeta, providerMetaSchema)

	diags.Append(providerMetaDiags...)

	fw.ProviderMeta = providerMeta

	privateData, privateDataDiags := privatestate.NewData(ctx, proto6.PriorPrivate)

	diags.Append(privateDataDiags...)

	fw.PriorPrivate = privateData

	return fw, diags
}
