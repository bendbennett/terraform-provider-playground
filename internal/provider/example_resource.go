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
	listNestedAttributes := tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
		"int64_attribute": {
			Optional: true,
			Type:     types.Int64Type,
		},
		"list_attribute": {
			Optional: true,
			Type:     types.ListType{ElemType: types.StringType},
		},
	})

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
					listNestedAttributes,
				},
			},
		},
	}, nil
}

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
	tflog.Info(ctx, "calling CustomFunc")
}

/*
func (l ListNestedAttributesCustomTypeType) TerraformType(ctx context.Context) tftypes.Type {
	//TODO implement me
	panic("TerraformType")
}

func (l ListNestedAttributesCustomTypeType) ValueType(ctx context.Context) attr.Value {
	//TODO implement me
	panic("ValueType")
}

func (l ListNestedAttributesCustomTypeType) Equal(t attr.Type) bool {
	//TODO implement me
	panic("Equal")
}

func (l ListNestedAttributesCustomTypeType) String() string {
	//TODO implement me
	panic("String")
}

func (l ListNestedAttributesCustomTypeType) ApplyTerraform5AttributePathStep(step tftypes.AttributePathStep) (interface{}, error) {
	//TODO implement me
	panic("ApplyTerraform5AttributePathStep")
}
*/

type exampleResourceData struct {
	Id types.String `tfsdk:"id"`

	ListNestedAttribute types.List `tfsdk:"list_nested_attribute"`

	ListNestedAttributeCustom ListNestedAttributesCustomValue `tfsdk:"list_nested_attribute_custom"`
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
