package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/timeouts/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ datasource.DataSource = exampleDataSource{}

func (t exampleDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Example data source",

		Attributes: map[string]schema.Attribute{
			"configurable_attribute": schema.StringAttribute{
				MarkdownDescription: "Example configurable attribute",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "Example identifier",
				Computed:            true,
			},
			"timeouts": timeouts.Attributes(ctx),
		},

		//Blocks: map[string]schema.Block{
		//	"timeouts": timeouts.Block(ctx),
		//},
	}
}

func NewDatasource() datasource.DataSource {
	return &exampleDataSource{}
}

type exampleDataSourceData struct {
	ConfigurableAttribute types.String   `tfsdk:"configurable_attribute"`
	Id                    types.String   `tfsdk:"id"`
	Timeouts              timeouts.Value `tfsdk:"timeouts"`
}

type exampleDataSource struct {
	provider playgroundProvider
}

func (r exampleDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_example"
}

func (d exampleDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data exampleDataSourceData

	diags := req.Config.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	readTimeout, err := data.Timeouts.Read(ctx)
	if err != nil {
		// handle error
	}

	tflog.Info(ctx, fmt.Sprintf("%v", readTimeout))

	_, cancel := context.WithTimeout(ctx, readTimeout)
	defer cancel()

	// For the purposes of this example code, hardcoding a response value to
	// save into the Terraform state.
	data.Id = types.StringValue("example-id")

	// Write logs using the tflog package
	// Documentation: https://terraform.io/plugin/log
	tflog.Trace(ctx, "read a data source")

	diags = resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}
