package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces
var _ provider.Provider = &scaffoldingProvider{}

// provider satisfies the tfsdk.Provider interface and usually is included
// with all Resource and DataSource implementations.
type scaffoldingProvider struct {
	// client can contain the upstream provider SDK or HTTP client used to
	// communicate with the upstream service. Resource and DataSource
	// implementations can then make calls using this client.
	//
	// TODO: If appropriate, implement upstream provider SDK or HTTP client.
	// client vendorsdk.ExampleClient

	// configured is set to true at the end of the Configure method.
	// This can be used in Resource and DataSource implementations to verify
	// that the provider was previously configured.
	configured bool

	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// providerData can be used to store data from the Terraform configuration.
type providerData struct {
	BoolAttribute    types.Bool    `tfsdk:"bool_attribute"`
	Float64Attribute types.Float64 `tfsdk:"float64_attribute"`
	Int64Attribute   types.Int64   `tfsdk:"int64_attribute"`
	ListAttribute    types.List    `tfsdk:"list_attribute"`
	MapAttribute     types.Map     `tfsdk:"map_attribute"`
	NumberAttribute  types.Number  `tfsdk:"number_attribute"`
	ObjectAttribute  types.Object  `tfsdk:"object_attribute"`
	SetAttribute     types.Set     `tfsdk:"set_attribute"`
	StringAttribute  types.String  `tfsdk:"string_attribute"`

	SingleNestedAttribute SingleNestedAttribute         `tfsdk:"single_nested_attribute"`
	ListNestedAttribute   []ListNestedAttribute         `tfsdk:"list_nested_attribute"`
	MapNestedAttribute    map[string]MapNestedAttribute `tfsdk:"map_nested_attribute"`
	SetNestedAttribute    []SetNestedAttribute          `tfsdk:"set_nested_attribute"`

	Block types.List `tfsdk:"block"`
}

type SingleNestedAttribute struct {
	BoolAttribute    types.Bool    `tfsdk:"bool_attribute"`
	Float64Attribute types.Float64 `tfsdk:"float64_attribute"`
}

type ListNestedAttribute struct {
	Int64Attribute types.Int64 `tfsdk:"int64_attribute"`
	ListAttribute  types.List  `tfsdk:"list_attribute"`
}

type MapNestedAttribute struct {
	MapAttribute    types.Map    `tfsdk:"map_attribute"`
	NumberAttribute types.Number `tfsdk:"number_attribute"`
}

type SetNestedAttribute struct {
	ObjectAttribute types.Object `tfsdk:"object_attribute"`
	SetAttribute    types.Set    `tfsdk:"set_attribute"`
	StringAttribute types.String `tfsdk:"string_attribute"`
}

func (p *scaffoldingProvider) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"bool_attribute": {
				Optional: true,
				Type:     types.BoolType,
			},

			"float64_attribute": {
				Optional: true,
				Type:     types.Float64Type,
			},

			"int64_attribute": {
				Optional: true,
				Type:     types.Int64Type,
			},

			"list_attribute": {
				Optional: true,
				Type:     types.ListType{ElemType: types.StringType},
			},

			"list_nested_attribute": {
				Optional: true,
				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
					"int64_attribute": {
						Optional: true,
						Type:     types.Int64Type,
					},
					"list_attribute": {
						Optional: true,
						Type:     types.ListType{ElemType: types.StringType},
					},
				}),
			},

			"map_attribute": {
				Optional: true,
				Type:     types.MapType{ElemType: types.StringType},
			},

			"number_attribute": {
				Optional: true,
				Type:     types.NumberType,
			},

			"object_attribute": {
				Optional: true,
				Type: types.ObjectType{
					AttrTypes: map[string]attr.Type{
						"bool_attribute":    types.BoolType,
						"float64_attribute": types.Float64Type,
						"int64_attribute":   types.Int64Type,
						"list_attribute":    types.ListType{ElemType: types.StringType},
						"map_attribute":     types.MapType{ElemType: types.StringType},
						"number_attribute":  types.NumberType,
						"set_attribute":     types.ListType{ElemType: types.StringType},
						"string_attribute":  types.StringType,
					},
				},
			},

			"set_attribute": {
				Optional: true,
				Type:     types.SetType{ElemType: types.StringType},
			},

			"single_nested_attribute": {
				Optional: true,
				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
					"bool_attribute": {
						Optional: true,
						Type:     types.BoolType,
					},
					"float64_attribute": {
						Optional: true,
						Type:     types.Float64Type,
					},
				}),
			},

			"string_attribute": {
				Optional: true,
				Type:     types.StringType,
			},

			"map_nested_attribute": {
				Optional: true,
				Attributes: tfsdk.MapNestedAttributes(map[string]tfsdk.Attribute{
					"map_attribute": {
						Optional: true,
						Type:     types.MapType{ElemType: types.StringType},
					},
					"number_attribute": {
						Optional: true,
						Type:     types.NumberType,
					},
				}),
			},

			"set_nested_attribute": {
				Optional: true,
				Attributes: tfsdk.SetNestedAttributes(map[string]tfsdk.Attribute{
					"object_attribute": {
						Optional: true,
						Type: types.ObjectType{
							AttrTypes: map[string]attr.Type{
								"bool_attribute":    types.BoolType,
								"float64_attribute": types.Float64Type,
								"int64_attribute":   types.Int64Type,
								"list_attribute":    types.ListType{ElemType: types.StringType},
								"map_attribute":     types.MapType{ElemType: types.StringType},
								"number_attribute":  types.NumberType,
								"set_attribute":     types.ListType{ElemType: types.StringType},
								"string_attribute":  types.StringType,
							},
						},
					},
					"set_attribute": {
						Optional: true,
						Type:     types.SetType{ElemType: types.StringType},
					},
					"string_attribute": {
						Optional: true,
						Type:     types.StringType,
					},
				}),
			},
		},

		Blocks: map[string]tfsdk.Block{
			"block": {
				NestingMode: tfsdk.BlockNestingModeList,
				Attributes: map[string]tfsdk.Attribute{
					"bool_attribute": {
						Optional: true,
						Type:     types.BoolType,
					},
					"float64_attribute": {
						Optional: true,
						Type:     types.Float64Type,
					},

					"int64_attribute": {
						Optional: true,
						Type:     types.Int64Type,
					},
					"list_attribute": {
						Optional: true,
						Type:     types.ListType{ElemType: types.StringType},
					},
				},
				Blocks: map[string]tfsdk.Block{
					"nested_block": {
						NestingMode: tfsdk.BlockNestingModeList,
						Attributes: map[string]tfsdk.Attribute{
							"bool_attribute": {
								Optional: true,
								Type:     types.BoolType,
							},
						},
					},
				},
			},
		},
	}, nil
}

func (p *scaffoldingProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data providerData
	diags := req.Config.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Configuration values are now available.
	// if data.Example.Null { /* ... */ }

	// If the upstream provider SDK or HTTP client requires configuration, such
	// as authentication or logging, this is a great opportunity to do so.

	p.configured = true
}

func (p *scaffoldingProvider) GetResources(ctx context.Context) (map[string]provider.ResourceType, diag.Diagnostics) {
	return map[string]provider.ResourceType{
		"attributes_example": exampleResourceType{},
	}, nil
}

func (p *scaffoldingProvider) GetDataSources(ctx context.Context) (map[string]provider.DataSourceType, diag.Diagnostics) {
	return map[string]provider.DataSourceType{
		"scaffolding_example": exampleDataSourceType{},
	}, nil
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &scaffoldingProvider{
			version: version,
		}
	}
}

// convertProviderType is a helper function for NewResource and NewDataSource
// implementations to associate the concrete provider type. Alternatively,
// this helper can be skipped and the provider type can be directly type
// asserted (e.g. provider: in.(*scaffoldingProvider)), however using this can prevent
// potential panics.
func convertProviderType(in provider.Provider) (scaffoldingProvider, diag.Diagnostics) {
	var diags diag.Diagnostics

	p, ok := in.(*scaffoldingProvider)

	if !ok {
		diags.AddError(
			"Unexpected Provider Instance Type",
			fmt.Sprintf("While creating the data source or resource, an unexpected provider type (%T) was received. This is always a bug in the provider code and should be reported to the provider developers.", p),
		)
		return scaffoldingProvider{}, diags
	}

	if p == nil {
		diags.AddError(
			"Unexpected Provider Instance Type",
			"While creating the data source or resource, an unexpected empty provider instance was received. This is always a bug in the provider code and should be reported to the provider developers.",
		)
		return scaffoldingProvider{}, diags
	}

	return *p, diags
}
