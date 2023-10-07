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

package testfixtures

// MessageTestTable is used by both the standard Go tests and also the
// Terraform acceptance tests.
// <https://github.com/golang/go/wiki/TableDrivenTests>
var MessageTestTable = map[string]struct {
	Message  string
	Expected string
}{
	"nocolor": {
		Message:  "This is a test.",
		Expected: "This is a test.\n",
	},
	"fg=yellow": {
		Message:  "<fg=yellow>This is a test.</>",
		Expected: "\x1b[33mThis is a test.\x1b[0m\n",
	},
	"fg=yellow;bg=blue": {
		Message:  "<fg=yellow;bg=blue>This is a test.</>",
		Expected: "\x1b[33;44mThis is a test.\x1b[0m\n",
	},
	"fg=yellow;bg=blue;op=bold": {
		Message:  "<fg=yellow;bg=blue;op=bold>This is a test.</>",
		Expected: "\x1b[33;44;1mThis is a test.\x1b[0m\n",
	},
	"fg=yellow;bg=blue;op=underline": {
		Message:  "<fg=yellow;bg=blue;op=underline>This is a test.</>",
		Expected: "\x1b[33;44mThis is a test.\x1b[0m\n",
	},
	"fg=yellow;bg=blue;op=fuzzy": {
		Message:  "<fg=yellow;bg=blue;op=fuzzy>This is a test.</>",
		Expected: "\x1b[33;44;2mThis is a test.\x1b[0m\n",
	},
	"fg=204,204,204": {
		Message:  "<fg=204,204,204>This is a test.</>",
		Expected: "\x1b[38;2;204;204;204mThis is a test.\x1b[0m\n",
	},
	"fg=6256CC": {
		Message:  "<fg=6256CC>This is a test.</>",
		Expected: "\x1b[38;2;98;86;204mThis is a test.\x1b[0m\n",
	},
	"<err>": {
		Message:  "<err>This is a test.</>",
		Expected: "\x1b[97;41mThis is a test.\x1b[0m\n",
	},
	"abc <err>err-text</> def <err>err-text</>": {
		Message:  "abc <err>err-text</> def <err>err-text</>",
		Expected: "abc \x1b[97;41merr-text\x1b[0m def \x1b[97;41merr-text\x1b[0m\n",
	},
	"<fg=e7b2a1;bg=176;op=bold>": {
		Message:  "<fg=e7b2a1;bg=176;op=bold>This is a test.</>",
		Expected: "\x1b[38;2;231;178;161;48;5;176;1mThis is a test.\x1b[0m\n",
	},
}
