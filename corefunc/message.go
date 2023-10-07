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
	"os"

	"github.com/gookit/color"
)

/*
Message allows you to print a custom message to the screen while the code is
running. This function wraps `github.com/gookit/color.Message()`, which supports
adding color to printed messages using an HTML-like syntax.

The [HTML-like syntax](https://gookit.github.io/color/#/?id=html-like-tag-usage)
can be a single (foreground) color like `<green>`; it can have foreground and
background colors like `<fg=yellow;bg=black;op=underscore>`; it supports
16-color mode, 256-color mode, and 24-bit “True Color” (16.7 million) color
mode; and it supports macOS, Linux, and Windows.

If you need more flexibility than this function provides, consider using
<https://github.com/gookit/color> directly.

* Foreground color labels: https://github.com/gookit/color/blob/v1.5.4/color_tag.go#L166-L215

* Background color labels: https://github.com/gookit/color/blob/v1.5.4/color_tag.go#L217-L266

* Option labels: https://github.com/gookit/color/blob/v1.5.4/color_tag.go#L268-L285

* Predefined labels: https://github.com/gookit/color/blob/v1.5.4/color_tag.go#L145-L157

----

* message (string): The message to print to the screen with a trailing newline.
*/
func Message(message string) {
	color.Fprintln(os.Stdout, message)
}
