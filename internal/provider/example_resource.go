package provider

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ resource.Resource = (*exampleResource)(nil)
var _ resource.ResourceWithImportState = (*exampleResource)(nil)

func (t *exampleResource) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Example resource",

		Attributes: map[string]tfsdk.Attribute{
			"configurable_attribute": {
				MarkdownDescription: "Example configurable attribute",
				Optional:            true,
				Type:                types.StringType,
			},
			"id": {
				Computed:            true,
				MarkdownDescription: "Example identifier",
				PlanModifiers: tfsdk.AttributePlanModifiers{
					resource.UseStateForUnknown(),
				},
				Type: types.StringType,
			},
			"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
				Create: true,
			}),
		},

		Blocks: map[string]tfsdk.Block{
			"timeouts": timeouts.Block(ctx, timeouts.Opts{
				Create: true,
				Read:   true,
			}),
		},

		//Blocks: map[string]tfsdk.Block{
		//	"timeouts": libraryCall()
		// only support Read
		// support All / CRUD timeouts block
		// type that defines how timeouts data is read out into model
		// helper(s) - defaults - can we do this?
		//
		// in Read function, try to read from single location, if not found provide default
		//},
	}, nil
}

func NewResource() resource.Resource {
	return &exampleResource{}
}

type TimeoutsType struct {
	types.Object
}

type exampleResourceData struct {
	ConfigurableAttribute types.String `tfsdk:"configurable_attribute"`
	Id                    types.String `tfsdk:"id"`
	Timeouts              types.Object `tfsdk:"timeouts"`
	//Timeouts              TimeoutsType `tfsdk:"timeouts"`
}

type Duration interface {
	GetDuration(ctx context.Context, op string) (*time.Duration, diag.Diagnostics)
}

type TimeoutsCreateReadStruct struct {
	Create types.String `tfsdk:"create"`
	Read   types.String `tfsdk:"read"`
}

func (t TimeoutsCreateReadStruct) GetDuration(ctx context.Context, op string) (*time.Duration, diag.Diagnostics) {
	var durationToParse string
	var diags diag.Diagnostics

	switch op {
	case "create":
		if t.Create.IsNull() || t.Create.IsUnknown() {
			diags.AddError("", "")
			return nil, diags
		}

		if t.Create.Value == "" {
			diags.AddError("", "")
			return nil, diags
		}

		durationToParse = t.Create.Value
	}

	dur, err := time.ParseDuration(durationToParse)
	if err != nil {
		diags.AddError("", "")
		return nil, diags
	}

	return &dur, nil
}

func (t TimeoutsCreateReadStruct) GetDurationDefault(ctx context.Context, op string, def time.Duration) *time.Duration {
	dur, diags := t.GetDuration(ctx, op)

	if diags.HasError() {
		return dur
	}

	return dur
}

type exampleResource struct {
	provider playgroundProvider
}

func (r *exampleResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_example"
}

func (r exampleResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data exampleResourceData

	diags := req.Config.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	d, diags := timeouts.Create(ctx, data.Timeouts)

	fmt.Println(d)

	context.WithTimeout(ctx, *d)
	//dur, diags := data.TimeoutsCreateReadStruct.GetDuration(ctx, "create")
	//fmt.Println(dur)

	//createTimeout := CreateTimeout(ctx, data.Timeouts, 20*time.Minute)
	//fmt.Println(createTimeout)

	data.Id = types.String{Value: "example-id"}

	tflog.Trace(ctx, "created a resource")

	diags = resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func CreateTimeout(ctx context.Context, timeouts types.Object, def time.Duration) time.Duration {
	if _, ok := timeouts.Attrs["create"]; !ok {
		return def
	}

	readTimeout := timeouts.Attrs["create"]

	if _, ok := readTimeout.(types.String); !ok {
		return def
	}

	duration, err := time.ParseDuration(readTimeout.(types.String).Value)
	if err != nil {
		return def
	}

	return duration
}

func (r exampleResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data exampleResourceData

	diags := req.State.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	//_ = resource.ReadTimeout(ctx, req, 20*time.Minute)

	diags = resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r exampleResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data exampleResourceData

	diags := req.Plan.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	diags = resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r exampleResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data exampleResourceData

	diags := req.State.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}
}

func (r exampleResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
