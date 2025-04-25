package provider

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/code-of-kpp/terraform-provider-playground/internal/client"
)

// Hints what's still missing
var _ provider.Provider = &PlaygroundProvider{}
var _ provider.ProviderWithFunctions = &PlaygroundProvider{}

// Main provider struct
type PlaygroundProvider struct {
	version string
}

// Optional builder shortcut
func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &PlaygroundProvider{version: version}
	}
}

func (p *PlaygroundProvider) Metadata(
	ctx context.Context, req provider.MetadataRequest,
	resp *provider.MetadataResponse,
) {
	resp.TypeName = "playground" // Prefix in terraform data & resource blocks
	resp.Version = p.version
}

type PlaygroundProviderModel struct {
	Folder types.String `tfsdk:"folder"`
}

func (p *PlaygroundProvider) Schema(
	ctx context.Context, req provider.SchemaRequest,
	resp *provider.SchemaResponse,
) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"folder": schema.StringAttribute{
				MarkdownDescription: "Playground working folder",
				Description:         "Playground working folder",
				Optional:            true,
			},
		},
	}
}

func (p *PlaygroundProvider) Configure(
	ctx context.Context, req provider.ConfigureRequest,
	resp *provider.ConfigureResponse,
) {
	var data PlaygroundProviderModel
	var folder string

	resp.Diagnostics.Append(
		req.Config.Get(ctx, &data)...,
	)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.Folder.IsNull() {
		tflog.Info(ctx, "Value not provided, trying environment variable")
		folder = os.Getenv("PLAYGROUND_FOLDER")
		if folder == "" {
			resp.Diagnostics.AddAttributeError(
				path.Root("folder"),
				"Folder not provided",
				"Provide Folder value via environment variable "+
					"PLAYGROUND_FOLDER or configuration.",
			)
			return
		}
	} else {
		folder = data.Folder.ValueString()
	}

	play_client := client.New(folder)
	resp.DataSourceData = play_client
	resp.ResourceData = play_client
}

func (p *PlaygroundProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{}
}

func (p *PlaygroundProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewShoppingListDataSource,
	}
}

func (p *PlaygroundProvider) Functions(
	ctx context.Context,
) []func() function.Function {
	return []func() function.Function{
		NewMeaningOfLifeFunction,
		NewValidPathFunction,
	}
}
