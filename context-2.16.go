// Copyright © 2016 Abcum Ltd
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

// +build webkit_2_1_6

package webkit

// #include <webkit2/webkit2.h>
import "C"

// DefaultWebContext returns the default WebContext.
func NewEphemeralContext() *Context {
	return &Context{C.webkit_web_context_new_ephemeral()}
}

// IsEphemeral returns whether the specified Context is ephemeral.
func (c *Context) IsEphemeral() bool {
	return gobool(C.webkit_web_context_is_ephemeral(c.context))
}
