package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

var _ function.Function = ExampleFunction{}

func NewExampleFunction() function.Function {
	return &ExampleFunction{}
}

type ExampleFunction struct{}

func (f ExampleFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "function"
}

func (f ExampleFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Parameters: []function.Parameter{
			function.StringParameter{
				CustomType: CustomStringType{},
			},
		},
		Return: function.StringReturn{},
	}
}

func (f ExampleFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var arg string

	resp.Diagnostics.Append(req.Arguments.Get(ctx, &arg)...)

	resp.Diagnostics.Append(resp.Result.Set(ctx, arg)...)
}
