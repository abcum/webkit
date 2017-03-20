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
//
// See also: WebKitWebContext at
// http://webkitgtk.org/reference/webkit2gtk/stable/WebKitWebContext.html.
type WebContext struct {
	webContext *C.WebKitWebContext
}

// DefaultWebContext returns the default WebContext.
//
// See also: webkit_web_context_get_default at
// http://webkitgtk.org/reference/webkit2gtk/stable/WebKitWebContext.html#webkit-web-context-get-default.
func DefaultWebContext() *WebContext {
	return &WebContext{C.webkit_web_context_get_default()}
}

// CacheModel describes the caching behavior.
//
// See also: WebKitCacheModel at
// http://webkitgtk.org/reference/webkit2gtk/stable/WebKitWebContext.html#WebKitCacheModel.
type CacheModel int

// CacheModel enum values are described at
// http://webkitgtk.org/reference/webkit2gtk/stable/WebKitWebContext.html#WebKitCacheModel.
const (
	DocumentViewerCacheModel CacheModel = iota
	WebBrowserCacheModel
	DocumentBrowserCacheModel
)

// CacheModel returns the current cache model.
//
// See also: webkit_web_context_get_cache_model at
// http://webkitgtk.org/reference/webkit2gtk/stable/WebKitWebContext.html#webkit-web-context-get-cache-model.
func (wc *WebContext) CacheModel() CacheModel {
	return CacheModel(C.int(C.webkit_web_context_get_cache_model(wc.webContext)))
}

// SetCacheModel sets the current cache model.
//
// See also: webkit_web_context_set_cache_model at
// http://webkitgtk.org/reference/webkit2gtk/stable/WebKitWebContext.html#webkit-web-context-set-cache-model.
func (wc *WebContext) SetCacheModel(model CacheModel) {
	C.webkit_web_context_set_cache_model(wc.webContext, C.WebKitCacheModel(model))
}

// ClearCache clears all resources currently cached.
//
// See also: webkit_web_context_clear_cache at
// http://webkitgtk.org/reference/webkit2gtk/stable/WebKitWebContext.html#webkit-web-context-clear-cache.
func (wc *WebContext) ClearCache() {
	C.webkit_web_context_clear_cache(wc.webContext)
}
