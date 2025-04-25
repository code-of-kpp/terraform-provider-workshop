package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Hints what's still missing
var _ provider.Provider = &PlaygroundProvider{}

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
	ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse,
) {
	resp.TypeName = "playground" // Prefix in terraform data & resource blocks
	resp.Version = p.version
}

func (p *PlaygroundProvider) Schema(
	ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse,
) {
	resp.Schema = schema.Schema{}
}
func (p *PlaygroundProvider) Configure(
	ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse,
) {
}

func (p *PlaygroundProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{}
}

func (p *PlaygroundProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}
