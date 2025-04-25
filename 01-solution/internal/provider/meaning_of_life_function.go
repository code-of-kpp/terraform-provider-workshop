package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

var _ function.Function = MeaningOfLifeFunction{}

type MeaningOfLifeFunction struct{}

func NewMeaningOfLifeFunction() function.Function { return MeaningOfLifeFunction{} }
func (f MeaningOfLifeFunction) Metadata(
	ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse,
) {
	resp.Name = "meaning_of_life"
}

func (f MeaningOfLifeFunction) Run(
	ctx context.Context, req function.RunRequest, resp *function.RunResponse,
) {
	resp.Error = function.ConcatFuncErrors(
		resp.Result.Set(ctx, 42),
	)
}

func (f MeaningOfLifeFunction) Definition(
	ctx context.Context, req function.DefinitionRequest,
	resp *function.DefinitionResponse,
) {
	resp.Definition = function.Definition{
		Parameters: []function.Parameter{},
		Return:     function.Int64Return{},
		Summary:    "Meaning of Life function",
		Description: "Provides the Ultimate answer " +
			"to Life, the Universe and Everything",
		MarkdownDescription: "Provides the Ultimate answer " +
			"to Life, the Universe and **Everything**",
	}
}
