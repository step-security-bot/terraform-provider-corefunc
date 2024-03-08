// Copyright 2023-2024, Northwood Labs
// Copyright 2023-2024, Ryan Parman <rparman@northwood-labs.com>
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

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/lithammer/dedent"
	"github.com/northwood-labs/terraform-provider-corefunc/corefunc"
)

// Ensure the implementation satisfies the expected interfaces.
var _ function.Function = &strSnakeFunction{}

type (
	// strSnakeFunction is the function implementation.
	strSnakeFunction struct{}
)

// StrSnakeFunction is a method that exposes its paired Go function as a
// Terraform Function.
func StrSnakeFunction() function.Function { // lint:allow_return_interface
	return &strSnakeFunction{}
}

func (f *strSnakeFunction) Metadata(
	ctx context.Context,
	req function.MetadataRequest,
	resp *function.MetadataResponse,
) {
	tflog.Debug(ctx, "Starting StrSnake Function Metadata method.")

	resp.Name = "str_snake"

	tflog.Debug(ctx, fmt.Sprintf("resp.Name = %s", resp.Name))

	tflog.Debug(ctx, "Ending StrSnake Function Metadata method.")
}

// Definition defines the parameters and return type for the function.
func (f *strSnakeFunction) Definition(
	ctx context.Context,
	req function.DefinitionRequest,
	resp *function.DefinitionResponse,
) {
	tflog.Debug(ctx, "Starting StrSnake Function Definition method.")

	resp.Definition = function.Definition{
		Summary: "Echo a string",
		Description: strings.TrimSpace(dedent.Dedent(`
		Converts a string to ` + "`" + `snake_case` + "`" + `, removing any non-alphanumeric characters.

		Maps to the StrSnake Go method, which can be used in Terratest.
		`)),
		MarkdownDescription: strings.TrimSpace(dedent.Dedent(`
		Converts a string to ` + "`" + `snake_case` + "`" + `, removing any non-alphanumeric characters.

		Maps to the ` + linkPackage("StrSnake") + ` Go method, which can be used in ` + Terratest + `.
        `)),
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:                "str",
				Description:         "The string to convert to `snake_case`.",
				MarkdownDescription: "The string to convert to `snake_case`.",
			},
		},
		Return: function.StringReturn{},
	}

	tflog.Debug(ctx, "Ending StrSnake Function Definition method.")
}

func (f *strSnakeFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	tflog.Debug(ctx, "Starting StrSnake Function Run method.")

	var str string
	err := req.Arguments.Get(ctx, &str)

	resp.Error = function.ConcatFuncErrors(err)
	if resp.Error != nil {
		return
	}

	value := corefunc.StrSnake(
		str,
		// opts,
	)

	// Function errors need to be wrapped and passed to resp.Error.
	// function.NewArgumentFuncError(1, "")
	// https://github.com/Yantrio/terraform-provider-tfutils/blob/main/internal/provider/cidrcontains.go

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, value))

	tflog.Debug(ctx, "Ending StrSnake Function Run method.")
}
