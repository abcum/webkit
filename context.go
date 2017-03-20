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

package webkit

// #include <webkit2/webkit2.h>
import "C"

// WebContext manages all aspects common to all WebViews.
type Context struct {
	context *C.WebKitWebContext
}

// NewWebContext creates a new WebContext.
func NewContext() *Context {
	return &Context{C.webkit_web_context_new()}
}

// DefaultWebContext returns the default WebContext.
func NewDefaultContext() *Context {
	return &Context{C.webkit_web_context_get_default()}
}

// DefaultWebContext returns the default WebContext.
func NewEphemeralContext() *Context {
	return &Context{C.webkit_web_context_new_ephemeral()}
}

// ----------------------------------------------------------------------------------------------------

// ClearCache clears all resources currently cached.
func (c *Context) ClearCache() {
	C.webkit_web_context_clear_cache(c.context)
}

// IsEphemeral returns whether the specified Context is ephemeral.
func (c *Context) IsEphemeral() bool {
	return gobool(C.webkit_web_context_is_ephemeral(c.context))
}

// ----------------------------------------------------------------------------------------------------

// CacheModel describes the caching behavior.
type CacheModel int

const (
	// Disable the cache completely, which substantially reduces memory usage. Useful for applications that only
	// access a single local file, with no navigation to other pages. No remote resources will be cached.
	DocumentViewerCacheModel CacheModel = iota
	// Improve document load speed substantially by caching a very large number of resources and previously
	// viewed content.
	WebBrowserCacheModel
	// A cache model optimized for viewing a series of local files -- for example, a documentation viewer or a
	// website designer. WebKit will cache a moderate number of resources.
	DocumentBrowserCacheModel
)

// CacheModel returns the current cache model.
func (c *Context) GetCacheModel() CacheModel {
	return CacheModel(C.int(C.webkit_web_context_get_cache_model(c.context)))
}

// SetCacheModel sets the current cache model.
func (c *Context) SetCacheModel(model CacheModel) {
	C.webkit_web_context_set_cache_model(c.context, C.WebKitCacheModel(model))
}

// ----------------------------------------------------------------------------------------------------

// ProcessModel describes the process behavior.
type ProcessModel int

const (
	// Use a single process to perform content rendering. The process is shared among all the WebKitWebView
	// instances created by the application: if the process hangs or crashes all the web views in the
	// application will be affected. This is the default process model, and it should suffice for most cases.
	SharedProcessModel ProcessModel = iota
	// Use one process for each WebKitWebView, while still allowing for some of them to share a process
	// in certain situations. The main advantage of this process model is that the rendering process for
	// a web view can crash while the rest of the views keep working normally. This process model is
	// indicated for applications which may use a number of web views and the content of in each must
	// not interfere with the rest — for example a full-fledged web browser with support for multiple tabs.
	MultipleProcessModel
)

// ProcessModel returns the current cache model.
func (c *Context) GetProcessModel() ProcessModel {
	return ProcessModel(C.int(C.webkit_web_context_get_process_model(c.context)))
}

// SetCacheModel sets the current cache model.
func (c *Context) SetProcessModel(model ProcessModel) {
	C.webkit_web_context_set_process_model(c.context, C.WebKitProcessModel(model))
}

// ----------------------------------------------------------------------------------------------------

// TLSErrorsPolicy describes how tls errors are handled.
type TLSErrorsPolicy int

const (
	// Ignore TLS errors
	IgnoreTLSErrorsPolicy TLSErrorsPolicy = iota
	// TLS errors will emit “load-failed-with-tls-errors” and, if the signal is handled, finish the
	// load. In case the signal is not handled, “load-failed” is emitted before the load finishes.
	FailTLSErrorsPolicy
)

// CacheModel returns the current cache model.
func (c *Context) GetTLSErrorsPolicy() TLSErrorsPolicy {
	return TLSErrorsPolicy(C.int(C.webkit_web_context_get_tls_errors_policy(c.context)))
}

// SetCacheModel sets the current cache model.
func (c *Context) SetTLSErrorsPolicy(policy TLSErrorsPolicy) {
	C.webkit_web_context_set_tls_errors_policy(c.context, C.WebKitTLSErrorsPolicy(policy))
}
