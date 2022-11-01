package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
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

			"list_nested_attribute_custom": {
				Optional: true,
				Attributes: ListNestedAttributesCustomType{
					tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
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
			},

			"map_nested_attribute": {
				Optional: true,
				Attributes: tfsdk.MapNestedAttributes(map[string]tfsdk.Attribute{
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

			"map_nested_attribute_custom": {
				Optional: true,
				Attributes: MapNestedAttributesCustomType{
					tfsdk.MapNestedAttributes(map[string]tfsdk.Attribute{
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
			},

			"set_nested_attribute": {
				Optional: true,
				Attributes: tfsdk.SetNestedAttributes(map[string]tfsdk.Attribute{
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

			"set_nested_attribute_custom": {
				Optional: true,
				Attributes: SetNestedAttributesCustomType{
					tfsdk.SetNestedAttributes(map[string]tfsdk.Attribute{
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
			},

			"single_nested_attribute": {
				Optional: true,
				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
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

			"single_nested_attribute_custom": {
				Optional: true,
				Attributes: SingleNestedAttributesCustomType{
					tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
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
			},
		},
		Blocks: map[string]tfsdk.Block{
			"single_nested_block": {
				Attributes: map[string]tfsdk.Attribute{
					"first": {
						Optional: true,
						Type:     types.StringType,
					},
					"second": {
						Optional: true,
						Type:     types.StringType,
					},
					"third": {
						Optional: true,
						Type:     types.StringType,
					},
				},
				NestingMode: tfsdk.BlockNestingModeSingle,
			},
		},
	}, nil
}

type exampleResourceData struct {
	Id types.String `tfsdk:"id"`

	// Nested Attributes
	ListNestedAttribute types.List `tfsdk:"list_nested_attribute"`

	ListNestedAttributeCustom ListNestedAttributesCustomValue `tfsdk:"list_nested_attribute_custom"`

	MapNestedAttribute types.Map `tfsdk:"map_nested_attribute"`

	MapNestedAttributeCustom MapNestedAttributesCustomValue `tfsdk:"map_nested_attribute_custom"`

	SetNestedAttribute types.Set `tfsdk:"set_nested_attribute"`

	SetNestedAttributeCustom SetNestedAttributesCustomValue `tfsdk:"set_nested_attribute_custom"`

	SingleNestedAttribute types.Object `tfsdk:"single_nested_attribute"`

	SingleNestedAttributeCustom SingleNestedAttributesCustomValue `tfsdk:"single_nested_attribute_custom"`

	// Nested Blocks
	SingleNestedBlock types.Object `tfsdk:"single_nested_block"`
}

// Create is unmarshalling the config onto exampleResourceData and persisting to the state.
func (e *exampleResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data exampleResourceData

	diags := req.Plan.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	data.ListNestedAttributeCustom.CustomFunc(ctx)
	data.MapNestedAttributeCustom.CustomFunc(ctx)
	data.SetNestedAttributeCustom.CustomFunc(ctx)
	data.SingleNestedAttributeCustom.CustomFunc(ctx)

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

// ListNestedAttributesCustomType for experimentation
type ListNestedAttributesCustomType struct {
	types.NestedAttributes
}

func (lnact ListNestedAttributesCustomType) Type() attr.Type {
	return ListNestedAttributesCustomTypeType{
		lnact.NestedAttributes.Type(),
	}
}

type ListNestedAttributesCustomTypeType struct {
	attr.Type
}

func (l ListNestedAttributesCustomTypeType) ValueFromTerraform(ctx context.Context, value tftypes.Value) (attr.Value, error) {
	val, err := l.Type.ValueFromTerraform(ctx, value)
	if err != nil {
		return nil, err
	}

	return ListNestedAttributesCustomValue{
		val.(types.List),
	}, nil
}

type ListNestedAttributesCustomValue struct {
	types.List
}

func (l ListNestedAttributesCustomValue) ToFrameworkValue() attr.Value {
	return l.List
}

func (l ListNestedAttributesCustomValue) CustomFunc(ctx context.Context) {
	tflog.Info(ctx, "calling CustomFunc on custom list nested attribute")
}

// MapNestedAttributesCustomType for experimentation
type MapNestedAttributesCustomType struct {
	types.NestedAttributes
}

func (lnact MapNestedAttributesCustomType) Type() attr.Type {
	return MapNestedAttributesCustomTypeType{
		lnact.NestedAttributes.Type(),
	}
}

type MapNestedAttributesCustomTypeType struct {
	attr.Type
}

func (l MapNestedAttributesCustomTypeType) ValueFromTerraform(ctx context.Context, value tftypes.Value) (attr.Value, error) {
	val, err := l.Type.ValueFromTerraform(ctx, value)
	if err != nil {
		return nil, err
	}

	return MapNestedAttributesCustomValue{
		val.(types.Map),
	}, nil
}

type MapNestedAttributesCustomValue struct {
	types.Map
}

func (l MapNestedAttributesCustomValue) ToFrameworkValue() attr.Value {
	return l.Map
}

func (l MapNestedAttributesCustomValue) CustomFunc(ctx context.Context) {
	tflog.Info(ctx, "calling CustomFunc on custom map nested attribute")
}

// SetNestedAttributesCustomType for experimentation
type SetNestedAttributesCustomType struct {
	types.NestedAttributes
}

func (lnact SetNestedAttributesCustomType) Type() attr.Type {
	return SetNestedAttributesCustomTypeType{
		lnact.NestedAttributes.Type(),
	}
}

type SetNestedAttributesCustomTypeType struct {
	attr.Type
}

func (l SetNestedAttributesCustomTypeType) ValueFromTerraform(ctx context.Context, value tftypes.Value) (attr.Value, error) {
	val, err := l.Type.ValueFromTerraform(ctx, value)
	if err != nil {
		return nil, err
	}

	return SetNestedAttributesCustomValue{
		val.(types.Set),
	}, nil
}

type SetNestedAttributesCustomValue struct {
	types.Set
}

func (l SetNestedAttributesCustomValue) ToFrameworkValue() attr.Value {
	return l.Set
}

func (l SetNestedAttributesCustomValue) CustomFunc(ctx context.Context) {
	tflog.Info(ctx, "calling CustomFunc on custom set nested attribute")
}

// SingleNestedAttributesCustomType for experimentation
type SingleNestedAttributesCustomType struct {
	types.NestedAttributes
}

func (lnact SingleNestedAttributesCustomType) Type() attr.Type {
	return SingleNestedAttributesCustomTypeType{
		lnact.NestedAttributes.Type(),
	}
}

type SingleNestedAttributesCustomTypeType struct {
	attr.Type
}

func (l SingleNestedAttributesCustomTypeType) ValueFromTerraform(ctx context.Context, value tftypes.Value) (attr.Value, error) {
	val, err := l.Type.ValueFromTerraform(ctx, value)
	if err != nil {
		return nil, err
	}

	return SingleNestedAttributesCustomValue{
		val.(types.Object),
	}, nil
}

type SingleNestedAttributesCustomValue struct {
	types.Object
}

func (l SingleNestedAttributesCustomValue) ToFrameworkValue() attr.Value {
	return l.Object
}

func (l SingleNestedAttributesCustomValue) CustomFunc(ctx context.Context) {
	tflog.Info(ctx, "calling CustomFunc on custom single nested attribute")
}
