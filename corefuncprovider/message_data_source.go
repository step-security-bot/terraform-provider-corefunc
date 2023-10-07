// Copyright 2023, Ryan Parman
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

	"github.com/gookit/color"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/lithammer/dedent"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &messageDataSource{}
	_ datasource.DataSourceWithConfigure = &messageDataSource{}
)

// messageDataSource is the data source implementation.
type (
	messageDataSource struct{}

	// messageDataSourceModel maps the data source schema data.
	messageDataSourceModel struct {
		ID      types.Int64  `tfsdk:"id"`
		Message types.String `tfsdk:"message"`
		Value   types.String `tfsdk:"value"`
	}
)

// MessageDataSource is a method that exposes its paired Go function as a
// Terraform Data Source.
func MessageDataSource() datasource.DataSource { // lint:allow_return_interface
	return &messageDataSource{}
}

// Metadata returns the data source type name.
func (d *messageDataSource) Metadata(
	ctx context.Context,
	req datasource.MetadataRequest,
	resp *datasource.MetadataResponse,
) {
	tflog.Info(ctx, "Starting Message DataSource Metadata method.")

	resp.TypeName = req.ProviderTypeName + "_message"

	tflog.Debug(ctx, fmt.Sprintf("req.ProviderTypeName = %s", req.ProviderTypeName))
	tflog.Debug(ctx, fmt.Sprintf("resp.TypeName = %s", resp.TypeName))

	tflog.Info(ctx, "Ending Message DataSource Metadata method.")
}

// Schema defines the schema for the data source.
func (d *messageDataSource) Schema(
	ctx context.Context,
	_ datasource.SchemaRequest,
	resp *datasource.SchemaResponse,
) {
	tflog.Info(ctx, "Starting Message DataSource Schema method.")

	resp.Schema = schema.Schema{
		MarkdownDescription: strings.TrimSpace(dedent.Dedent(`
		Message allows you to print a custom message to the screen while the code is
		running. This function wraps [` + "`" + `github.com/gookit/color.Message()` + "`" +
			`](https://pkg.go.dev/github.com/gookit/color#Message), which supports
		adding color to printed messages using an HTML-like syntax.

		The [HTML-like syntax](https://gookit.github.io/color/#/?id=html-like-tag-usage)
		can be a single (foreground) color like ` + "`" + "<green>" + "`" + `; it can have foreground and
		background colors like ` + "`" + "<fg=yellow;bg=black;op=underscore>" + "`" + `; it supports
		16-color mode, 256-color mode, and 24-bit “True Color” (16.7 million) color
		mode; and it supports macOS, Linux, and Windows.

		* [Foreground color labels](https://github.com/gookit/color/blob/v1.5.4/color_tag.go#L166-L215)
		* [Background color labels](https://github.com/gookit/color/blob/v1.5.4/color_tag.go#L217-L266)
		* [Option labels](https://github.com/gookit/color/blob/v1.5.4/color_tag.go#L268-L285)
		* [Predefined labels](https://github.com/gookit/color/blob/v1.5.4/color_tag.go#L145-L157)

		Maps to the ` + linkPackage("Message") + ` Go method, which can be used in
		` + Terratest + `.
		`)),
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Description: "Not used. Required by the " + TPF + ".",
				Computed:    true,
			},
			"message": schema.StringAttribute{
				Description: "The message that should be written to stdout during a plan/apply.",
				Required:    true,
			},
			"value": schema.StringAttribute{
				Description: "The value of the truncated label.",
				Computed:    true,
			},
		},
	}

	tflog.Info(ctx, "Ending Message DataSource Schema method.")
}

// Configure adds the provider configured client to the data source.
func (d *messageDataSource) Configure(
	ctx context.Context,
	req datasource.ConfigureRequest,
	_ *datasource.ConfigureResponse,
) {
	tflog.Info(ctx, "Starting Message DataSource Configure method.")

	if req.ProviderData == nil {
		return
	}

	tflog.Info(ctx, "Ending Message DataSource Configure method.")
}

func (d messageDataSource) Create(
	ctx context.Context,
	req resource.CreateRequest, // lint:allow_large_memory
	resp *resource.CreateResponse,
) {
	tflog.Info(ctx, "Starting Message DataSource Create method.")

	var plan messageDataSourceModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Ending Message DataSource Create method.")
}

// Read refreshes the Terraform state with the latest data.
func (d *messageDataSource) Read( // lint:no_dupe
	ctx context.Context,
	_ datasource.ReadRequest, // lint:allow_large_memory
	resp *datasource.ReadResponse,
) {
	tflog.Info(ctx, "Starting Message DataSource Read method.")

	var state messageDataSourceModel
	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)

	state.ID = types.Int64Value(1)

	// Return the attribute
	state.Value = types.StringValue(
		color.Sprintf(
			state.Message.ValueString(),
		),
	)

	// Print to output
	// corefunc.Message(state.Message.ValueString())

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Ending Message DataSource Read method.")
}
