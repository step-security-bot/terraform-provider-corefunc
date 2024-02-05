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

package corefunc

import (
	"github.com/nlnwa/whatwg-url/canonicalizer"
	"github.com/nlnwa/whatwg-url/url"
	"github.com/northwood-labs/terraform-provider-corefunc/corefunc/types"
)

/*
URLParse is a WHATWG spec-compliant <https://url.spec.whatwg.org/#url-parsing>
URL parser returns a struct representing a parsed absolute URL. Will not resolve
relative URLs.

The parser is up to date as of 2023-05-24
<https://url.spec.whatwg.org/commit-snapshots/eee49fdf4f99d59f717cbeb0bce29fda930196d4/>
and passes all relevant tests from web-platform-tests
<https://github.com/web-platform-tests/wpt/tree/master/url>.

Its API is similar to Chapter 6 in WHATWG URL Standard
<https://url.spec.whatwg.org/#api>.

----

  - rawURL (string): An absolute URL.

  - canon (types.URLCanonicalizer): An optional parameter that allows you to
    specify which canonicalizer to use. If not specified, the standard
    canonicalizer will be used.
*/
func URLParse(rawURL string, canon ...types.URLCanonicalizer) (*url.Url, error) {
	if len(canon) > 0 {
		switch canon[0] {
		case types.Standard:
			return url.Parse(rawURL) // lint:allow_unwrapped_errors
		case types.GoogleSafeBrowsing:
			return canonicalizer.GoogleSafeBrowsing.Parse(rawURL) // lint:allow_unwrapped_errors
		}
	}

	// Default
	return url.Parse(rawURL) // lint:allow_unwrapped_errors
}
