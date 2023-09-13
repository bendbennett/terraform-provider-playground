package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ resource.Resource = (*exampleResource)(nil)
var _ resource.ResourceWithImportState = (*exampleResource)(nil)

type exampleResource struct {
	provider exampleProvider
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
	ConfigurableAttribute types.String `tfsdk:"configurable_attribute"`
	Id                    types.String `tfsdk:"id"`
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
		data.ConfigurableAttribute = types.StringValue("configurable attribute")
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
		data.ConfigurableAttribute = types.StringValue("set in read")
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
