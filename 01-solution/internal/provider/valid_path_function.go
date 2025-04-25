package provider

import (
	"context"
	"os"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ function.Function = ValidPathFunction{}

type ValidPathFunction struct{}

func NewValidPathFunction() function.Function { return ValidPathFunction{} }
func (f ValidPathFunction) Metadata(
	ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse,
) {
	resp.Name = "valid_path"
}

func (f ValidPathFunction) Run(
	ctx context.Context, req function.RunRequest, resp *function.RunResponse,
) {
	var path types.String
	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &path))
	if resp.Error != nil {
		return
	}
	if _, err := os.Stat(path.ValueString()); err != nil {
		if os.IsNotExist(err) {
			resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, false))
			return
		}
		resp.Error = function.NewFuncError(
			"Error checking path: " + err.Error(),
		)
		return
	}
	if strings.Contains(path.ValueString(), "production") {
		resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, true))
	} else {
		resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, true))
	}
}

func (f ValidPathFunction) Definition(
	ctx context.Context, req function.DefinitionRequest,
	resp *function.DefinitionResponse,
) {
	resp.Definition = function.Definition{
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:                "path",
				Description:         "Path to check",
				MarkdownDescription: "Path to check",
			},
		},
		Return:              function.BoolReturn{},
		Summary:             "Check if path is valid for the playground",
		Description:         "Returns True if path is valid for the playground",
		MarkdownDescription: "Returns True if path is valid for the playground",
	}
}
