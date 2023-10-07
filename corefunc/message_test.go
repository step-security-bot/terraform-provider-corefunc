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

package corefunc

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/northwood-labs/terraform-provider-corefunc/testfixtures"
)

func TestMessage(t *testing.T) {
	for name, tc := range testfixtures.MessageTestTable {
		t.Run(name, func(t *testing.T) {
			actual := capture(func() {
				Message(tc.Message)
			})
			diff := cmp.Diff(tc.Expected, actual)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}

func FuzzMessage(f *testing.F) {
	for _, tc := range testfixtures.MessageTestTable {
		f.Add(tc.Message)
	}

	f.Fuzz(
		func(t *testing.T, input string) {
			Message(input)
		},
	)
}

func capture(f func()) string {
	r, w, err := os.Pipe()
	if err != nil {
		panic(err)
	}

	stdout := os.Stdout
	os.Stdout = w
	f()
	w.Close()
	os.Stdout = stdout

	var buf bytes.Buffer
	io.Copy(&buf, r)

	return buf.String()
}
