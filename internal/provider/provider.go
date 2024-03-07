package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ provider.Provider = (*playgroundProvider)(nil)

type playgroundProvider struct{}

func New() func() provider.Provider {
	return func() provider.Provider {
		return &playgroundProvider{}
	}
}

type providerData struct {
	ConfigurableAttribute types.String `tfsdk:"configurable_attribute"`
}

func (p *playgroundProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data providerData
	diags := req.Config.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}
}

func (p *playgroundProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "playground"
}

func (p *playgroundProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewDataSource,
	}
}

func (p *playgroundProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewResource,
	}
}

func (p *playgroundProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"configurable_attribute": schema.StringAttribute{
				Optional: true,
			},
		},
	}
}
