// Copyright 2023-2024, Ryan Parman
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package corefuncprovider // lint:no_dupe

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/lithammer/dedent"
	"github.com/northwood-labs/terraform-provider-corefunc/corefunc"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &homedirGetDataSource{}
	_ datasource.DataSourceWithConfigure = &homedirGetDataSource{}
)

// homedirGetDataSource is the data source implementation.
type (
	homedirGetDataSource struct{}

	// homedirGetDataSourceModel maps the data source schema data.
	homedirGetDataSourceModel struct {
		Value types.String `tfsdk:"value"`
	}
)

// HomedirGetDataSource is a method that exposes its paired Go function as a
// Terraform Data Source.
func HomedirGetDataSource() datasource.DataSource { // lint:allow_return_interface
	return &homedirGetDataSource{}
}

// Metadata returns the data source type name.
func (d *homedirGetDataSource) Metadata(
	ctx context.Context,
	req datasource.MetadataRequest,
	resp *datasource.MetadataResponse,
) {
	tflog.Debug(ctx, "Starting HomedirGet DataSource Metadata method.")

	resp.TypeName = req.ProviderTypeName + "_homedir_get"

	tflog.Debug(ctx, fmt.Sprintf("req.ProviderTypeName = %s", req.ProviderTypeName))
	tflog.Debug(ctx, fmt.Sprintf("resp.TypeName = %s", resp.TypeName))

	tflog.Debug(ctx, "Ending HomedirGet DataSource Metadata method.")
}

// Schema defines the schema for the data source.
func (d *homedirGetDataSource) Schema(
	ctx context.Context,
	_ datasource.SchemaRequest,
	resp *datasource.SchemaResponse,
) {
	tflog.Debug(ctx, "Starting HomedirGet DataSource Schema method.")

	resp.Schema = schema.Schema{
		MarkdownDescription: strings.TrimSpace(dedent.Dedent(`
		Returns the path to the home directory of the current user.

		Maps to the ` + linkPackage("Homedir") + ` Go method, which can be used in ` + Terratest + `.
		`)),
		Attributes: map[string]schema.Attribute{
			"value": schema.StringAttribute{
				Description: "The value of the home directory.",
				Computed:    true,
			},
		},
	}

	tflog.Debug(ctx, "Ending HomedirGet DataSource Schema method.")
}

// Configure adds the provider configured client to the data source.
func (d *homedirGetDataSource) Configure(
	ctx context.Context,
	req datasource.ConfigureRequest,
	_ *datasource.ConfigureResponse,
) {
	tflog.Debug(ctx, "Starting HomedirGet DataSource Configure method.")

	if req.ProviderData == nil {
		return
	}

	tflog.Debug(ctx, "Ending HomedirGet DataSource Configure method.")
}

func (d *homedirGetDataSource) Create(
	ctx context.Context,
	req resource.CreateRequest, // lint:allow_large_memory
	resp *resource.CreateResponse,
) {
	tflog.Debug(ctx, "Starting HomedirGet DataSource Create method.")

	var plan homedirGetDataSourceModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Ending HomedirGet DataSource Create method.")
}

// Read refreshes the Terraform state with the latest data.
func (d *homedirGetDataSource) Read( // lint:no_dupe
	ctx context.Context,
	_ datasource.ReadRequest, // lint:allow_large_memory
	resp *datasource.ReadResponse,
) {
	tflog.Debug(ctx, "Starting HomedirGet DataSource Read method.")

	var state homedirGetDataSourceModel
	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)

	homedir, err := corefunc.Homedir()
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to get the user's home directory",
			err.Error(),
		)

		return
	}

	state.Value = types.StringValue(homedir)

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Ending HomedirGet DataSource Read method.")
}
