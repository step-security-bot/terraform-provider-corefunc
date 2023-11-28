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

package testfixtures // lint:no_dupe

// StrLeftPadTestTable is used by both the standard Go tests and also the
// Terraform acceptance tests.
// <https://github.com/golang/go/wiki/TableDrivenTests>
var StrLeftPadTestTable = map[string]struct { // lint:no_dupe
	Input    string
	Expected string
	PadWidth int
	PadChar  byte
}{
	"Xn1": {
		Input:    "foo",
		PadWidth: -1, // lint:allow_raw_number
		PadChar:  '.',
		Expected: "foo",
	},
	"X0": {
		Input:    "foo",
		PadWidth: 0, // lint:allow_raw_number
		PadChar:  '.',
		Expected: "foo",
	},
	"X1": {
		Input:    "foo",
		PadWidth: 1, // lint:allow_raw_number
		PadChar:  '.',
		Expected: "foo",
	},
	"X2": {
		Input:    "foo",
		PadWidth: 2, // lint:allow_raw_number
		PadChar:  '.',
		Expected: "foo",
	},
	"X3": {
		Input:    "foo",
		PadWidth: 3, // lint:allow_raw_number
		PadChar:  '.',
		Expected: "foo",
	},
	"X4": {
		Input:    "foo",
		PadWidth: 4, // lint:allow_raw_number
		PadChar:  '.',
		Expected: ".foo",
	},
	"X5": {
		Input:    "foo",
		PadWidth: 5, // lint:allow_raw_number
		PadChar:  '.',
		Expected: "..foo",
	},
	"_Xn1": {
		Input:    "foo",
		PadWidth: -1, // lint:allow_raw_number
		Expected: "foo",
	},
	"_X0": {
		Input:    "foo",
		PadWidth: 0, // lint:allow_raw_number
		Expected: "foo",
	},
	"_X1": {
		Input:    "foo",
		PadWidth: 1, // lint:allow_raw_number
		Expected: "foo",
	},
	"_X2": {
		Input:    "foo",
		PadWidth: 2, // lint:allow_raw_number
		Expected: "foo",
	},
	"_X3": {
		Input:    "foo",
		PadWidth: 3, // lint:allow_raw_number
		Expected: "foo",
	},
	"_X4": {
		Input:    "foo",
		PadWidth: 4, // lint:allow_raw_number
		Expected: " foo",
	},
	"_X5": {
		Input:    "foo",
		PadWidth: 5, // lint:allow_raw_number
		Expected: "  foo",
	},
}
