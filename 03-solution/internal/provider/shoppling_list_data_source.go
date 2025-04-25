package provider

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/code-of-kpp/terraform-provider-playground/internal/client"
)

var _ datasource.DataSourceWithConfigure = (*ShoppingListDataSource)(nil)

func NewShoppingListDataSource() datasource.DataSource {
	return &ShoppingListDataSource{}
}

type ShoppingListDataSource struct {
	client *client.Client
}

type ShoppingListDataSourceModel struct {
	FileName      types.String `tfsdk:"file_name"`
	NumberOfItems types.Int64  `tfsdk:"number_of_items"`
}

func (d *ShoppingListDataSource) Metadata(
	ctx context.Context, req datasource.MetadataRequest,
	resp *datasource.MetadataResponse,
) {
	resp.TypeName = req.ProviderTypeName + "_shopping_list"
}

func (d *ShoppingListDataSource) Schema(
	ctx context.Context, req datasource.SchemaRequest,
	resp *datasource.SchemaResponse,
) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Shopping List Data Source",

		Attributes: map[string]schema.Attribute{
			"file_name": schema.StringAttribute{
				MarkdownDescription: "Name of the file with the shopping list",
				Required:            true,
			},
			"number_of_items": schema.Int64Attribute{
				MarkdownDescription: "Number of items in the list",
				Computed:            true,
			},
		},
	}
}

func (d *ShoppingListDataSource) Configure(
	ctx context.Context, req datasource.ConfigureRequest,
	resp *datasource.ConfigureResponse,
) {
	if req.ProviderData == nil {
		return
	}

	config, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type of ProviderData",
			fmt.Sprintf(
				"Expected *client.Client, got: %T.",
				req.ProviderData,
			),
		)

		return
	}

	d.client = config
}

func (d *ShoppingListDataSource) Read(
	ctx context.Context, req datasource.ReadRequest,
	resp *datasource.ReadResponse,
) {
	var data ShoppingListDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	fullPath := filepath.Join(d.client.Folder, data.FileName.ValueString())

	var items int
	file, err := os.Open(fullPath)
	if err != nil {
		resp.Diagnostics.AddError("Error opening file", err.Error())
		return
	}
	defer func() {
		if err := file.Close(); err != nil {
			resp.Diagnostics.AddError("Error closing file", err.Error())
		}
	}()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		items++
	}

	data.NumberOfItems = types.Int64Value(int64(items))

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
