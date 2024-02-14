package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ resource.Resource = (*exampleResource)(nil)
var _ resource.ResourceWithImportState = (*exampleResource)(nil)

type exampleResource struct {
}

func NewResource() resource.Resource {
	return &exampleResource{}
}

func (e *exampleResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_resource"
}

func (e *exampleResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"configurable_attribute": schema.StringAttribute{
				CustomType:          CustomStringType{},
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Example configurable attribute",
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Example identifier",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

type exampleResourceData struct {
	ConfigurableAttribute CustomStringValue `tfsdk:"configurable_attribute"`
	Id                    types.String      `tfsdk:"id"`
}

func (e *exampleResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data exampleResourceData

	diags := req.Plan.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = types.StringValue("example-id")

	if data.ConfigurableAttribute.IsNull() || data.ConfigurableAttribute.IsUnknown() {
		data.ConfigurableAttribute = CustomStringValue{
			StringValue: types.StringValue("set in create"),
		}
	}

	tflog.Trace(ctx, "created a resource")

	diags = resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (e *exampleResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data exampleResourceData

	diags := req.State.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	if data.ConfigurableAttribute.IsNull() || data.ConfigurableAttribute.IsUnknown() {
		data.ConfigurableAttribute = CustomStringValue{
			StringValue: types.StringValue("set in read"),
		}
	}

	diags = resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

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

func (e *exampleResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data exampleResourceData

	diags := req.State.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}
}

func (e *exampleResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

/************ CustomStringType *************/
// Ensure the implementation satisfies the expected interfaces
var _ basetypes.StringTypable = CustomStringType{}

type CustomStringType struct {
	basetypes.StringType
	// ... potentially other fields ...
}

func (t CustomStringType) Equal(o attr.Type) bool {
	other, ok := o.(CustomStringType)

	if !ok {
		return false
	}

	return t.StringType.Equal(other.StringType)
}

func (t CustomStringType) String() string {
	return "CustomStringType"
}

// Validate CustomStringType defined in the schema type section
func (t CustomStringType) Validate(ctx context.Context, value tftypes.Value, valuePath path.Path) diag.Diagnostics {
	if value.IsNull() || !value.IsKnown() {
		return nil
	}

	var diags diag.Diagnostics
	var valueString string

	if err := value.As(&valueString); err != nil {
		diags.AddAttributeError(
			valuePath,
			"Invalid Terraform Value",
			"An unexpected error occurred while attempting to convert a Terraform value to a string. "+
				"This generally is an issue with the provider schema implementation. "+
				"Please contact the provider developers.\n\n"+
				"Path: "+valuePath.String()+"\n"+
				"Error: "+err.Error(),
		)

		return diags
	}

	if valueString != "some-value" {
		diags.AddAttributeError(
			valuePath,
			"Invalid String Value",
			"The supplied string does not equal \"some-value\"",
		)

		return diags
	}

	return diags
}

func (t CustomStringType) ValueFromString(ctx context.Context, in basetypes.StringValue) (basetypes.StringValuable, diag.Diagnostics) {
	// CustomStringValue defined in the value type section
	value := CustomStringValue{
		StringValue: in,
	}

	return value, nil
}

func (t CustomStringType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	attrValue, err := t.StringType.ValueFromTerraform(ctx, in)

	if err != nil {
		return nil, err
	}

	stringValue, ok := attrValue.(basetypes.StringValue)

	if !ok {
		return nil, fmt.Errorf("unexpected value type of %T", attrValue)
	}

	stringValuable, diags := t.ValueFromString(ctx, stringValue)

	if diags.HasError() {
		return nil, fmt.Errorf("unexpected error converting StringValue to StringValuable: %v", diags)
	}

	return stringValuable, nil
}

func (t CustomStringType) ValueType(ctx context.Context) attr.Value {
	// CustomStringValue defined in the value type section
	return CustomStringValue{}
}

/************ CustomStringValue *************/
// Ensure the implementation satisfies the expected interfaces
var _ basetypes.StringValuable = CustomStringValue{}

type CustomStringValue struct {
	basetypes.StringValue
	// ... potentially other fields ...
}

func (v CustomStringValue) Equal(o attr.Value) bool {
	other, ok := o.(CustomStringValue)

	if !ok {
		return false
	}

	return v.StringValue.Equal(other.StringValue)
}

func (v CustomStringValue) Type(ctx context.Context) attr.Type {
	// CustomStringType defined in the schema type section
	return CustomStringType{}
}
