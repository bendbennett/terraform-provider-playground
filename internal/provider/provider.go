package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
)

var _ provider.Provider = &playgroundProvider{}

type playgroundProvider struct{}

func (p *playgroundProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "timeouts"
}

func (p *playgroundProvider) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{}, nil
}

func (p *playgroundProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
}

func (p *playgroundProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewResource,
	}
}

func (p *playgroundProvider) DataSources(context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
	//return []func() datasource.DataSource{
	//	"scaffolding_example": exampleDataSourceType{},
	//}, nil
}

func New() provider.Provider {
	return &playgroundProvider{}
}
