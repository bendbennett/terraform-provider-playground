package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.Resource = (*exampleResource)(nil)

type exampleResource struct {
}

func NewResource() resource.Resource {
	return &exampleResource{}
}

func (e *exampleResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_resource"
}

func (e *exampleResource) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Computed: true,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					resource.UseStateForUnknown(),
				},
				Type: types.StringType,
			},

			// Simple/primitive attributes
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

			"string_attribute": {
				Optional: true,
				Type:     types.StringType,
			},

			// Nested Attributes
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

			"single_nested_attribute": {
				Optional: true,
				Computed: true,
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
		},

		// Nested Blocks
		Blocks: map[string]tfsdk.Block{
			"list_nested_block": {
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
					"list_nested_nested_block": {
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

			"set_nested_block": {
				NestingMode: tfsdk.BlockNestingModeSet,
				Attributes: map[string]tfsdk.Attribute{
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
				},
				Blocks: map[string]tfsdk.Block{
					"set_nested_nested_block": {
						NestingMode: tfsdk.BlockNestingModeSet,
						Attributes: map[string]tfsdk.Attribute{
							"bool_attribute": {
								Optional: true,
								Type:     types.BoolType,
							},
						},
					},
				},
			},

			"single_nested_block": {
				NestingMode: tfsdk.BlockNestingModeSingle,
				Attributes: map[string]tfsdk.Attribute{
					"bool_attribute": {
						Optional: true,
						Type:     types.BoolType,
					},
					"float64_attribute": {
						Optional: true,
						Type:     types.Float64Type,
					},
				},
			},
		},
	}, nil
}

type exampleResourceData struct {
	Id types.String `tfsdk:"id"`

	// Simple/primitive attributes
	BoolAttribute    types.Bool    `tfsdk:"bool_attribute"`
	Float64Attribute types.Float64 `tfsdk:"float64_attribute"`
	Int64Attribute   types.Int64   `tfsdk:"int64_attribute"`
	ListAttribute    types.List    `tfsdk:"list_attribute"`
	MapAttribute     types.Map     `tfsdk:"map_attribute"`
	NumberAttribute  types.Number  `tfsdk:"number_attribute"`
	ObjectAttribute  types.Object  `tfsdk:"object_attribute"`
	SetAttribute     types.Set     `tfsdk:"set_attribute"`
	StringAttribute  types.String  `tfsdk:"string_attribute"`

	// Nested Attributes
	ListNestedAttribute   []types.Object          `tfsdk:"list_nested_attribute"`
	MapNestedAttribute    map[string]types.Object `tfsdk:"map_nested_attribute"`
	SetNestedAttribute    []types.Object          `tfsdk:"set_nested_attribute"`
	SingleNestedAttribute types.Object            `tfsdk:"single_nested_attribute"`

	// Nested Blocks
	ListNestedBlock   []types.Object `tfsdk:"list_nested_block"`
	SetNestedBlock    []types.Object `tfsdk:"set_nested_block"`
	SingleNestedBlock types.Object   `tfsdk:"single_nested_block"`
}

// Create is unmarshalling the config onto exampleResourceData and persisting to the state.
func (e *exampleResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data exampleResourceData

	diags := req.Config.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = types.String{Value: "example-id"}

	diags = resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

// Read is returning the contents of the state for this resource and the State field within resource.ReadResponse
// is pre-populated so no action is required in this function.
func (e *exampleResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
}

// Update overwrites the state with the plan. This is required in order that optional attributes in the config
// are updated in place.
func (e *exampleResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data exampleResourceData

	diags := req.Plan.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	diags = resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

// Delete is automatically handled by the Framework which sets an empty state. This function does not need to
// be populated unless additional handling needs to be implemented over and above setting an empty state.
func (e *exampleResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}
