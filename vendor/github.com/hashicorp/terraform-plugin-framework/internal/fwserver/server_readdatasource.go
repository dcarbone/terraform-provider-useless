package fwserver

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/internal/fwschema"
	"github.com/hashicorp/terraform-plugin-framework/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
)

// ReadDataSourceRequest is the framework server request for the
// ReadDataSource RPC.
type ReadDataSourceRequest struct {
	Config           *tfsdk.Config
	DataSourceSchema fwschema.Schema
	DataSourceType   provider.DataSourceType
	ProviderMeta     *tfsdk.Config
}

// ReadDataSourceResponse is the framework server response for the
// ReadDataSource RPC.
type ReadDataSourceResponse struct {
	Diagnostics diag.Diagnostics
	State       *tfsdk.State
}

// ReadDataSource implements the framework server ReadDataSource RPC.
func (s *Server) ReadDataSource(ctx context.Context, req *ReadDataSourceRequest, resp *ReadDataSourceResponse) {
	if req == nil {
		return
	}

	// Always instantiate new DataSource instances.
	logging.FrameworkDebug(ctx, "Calling provider defined DataSourceType NewDataSource")
	dataSource, diags := req.DataSourceType.NewDataSource(ctx, s.Provider)
	logging.FrameworkDebug(ctx, "Called provider defined DataSourceType NewDataSource")

	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	readReq := datasource.ReadRequest{
		Config: tfsdk.Config{
			Schema: schema(req.DataSourceSchema),
		},
	}
	readResp := datasource.ReadResponse{
		State: tfsdk.State{
			Schema: schema(req.DataSourceSchema),
		},
	}

	if req.Config != nil {
		readReq.Config = *req.Config
		readResp.State.Raw = req.Config.Raw.Copy()
	}

	if req.ProviderMeta != nil {
		readReq.ProviderMeta = *req.ProviderMeta
	}

	logging.FrameworkDebug(ctx, "Calling provider defined DataSource Read")
	dataSource.Read(ctx, readReq, &readResp)
	logging.FrameworkDebug(ctx, "Called provider defined DataSource Read")

	resp.Diagnostics = readResp.Diagnostics
	resp.State = &readResp.State
}
