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

// Settings manages the behaviour of a single WebView.
type Settings struct {
	settings *C.WebKitSettings
}

// NewSettings creates a new Settings with default values.
func NewSettings() *Settings {
	return &Settings{C.webkit_settings_new()}
}

// ----------------------------------------------------------------------------------------------------

// GetAutoLoadImages gets whether images should be automatically loaded or not. On devices
// where network bandwidth is of concern, it might be useful to turn this property off.
func (s *Settings) GetAutoLoadImages() bool {
	return gobool(C.webkit_settings_get_auto_load_images(s.settings))
}

// SetAutoLoadImages sets whether images should be automatically loaded or not. On devices
// where network bandwidth is of concern, it might be useful to turn this property off.
func (s *Settings) SetAutoLoadImages(enabled bool) {
	C.webkit_settings_set_auto_load_images(s.settings, cgobool(enabled))
}

// ----------------------------------------------------------------------------------------------------

// GetFrameFlattening gets whether to enable the frame flattening. With this setting each
// subframe is expanded to its contents, which will flatten all the frames to become one
// scrollable page.
func (s *Settings) GetFrameFlattening() bool {
	return gobool(C.webkit_settings_get_enable_frame_flattening(s.settings))
}

// SetFrameFlattening sets whether to enable the frame flattening. With this setting each
// subframe is expanded to its contents, which will flatten all the frames to become one
// scrollable page.
func (s *Settings) SetFrameFlattening(enabled bool) {
	C.webkit_settings_set_enable_frame_flattening(s.settings, cgobool(enabled))
}

// ----------------------------------------------------------------------------------------------------

// GetHTML5Database gets whether to enable HTML5 client-side SQL database support.
// Client-side SQL database allows web pages to store structured data and be able to use
// SQL to manipulate that data asynchronously.
func (s *Settings) GetHTML5Database() bool {
	return gobool(C.webkit_settings_get_enable_html5_database(s.settings))
}

// SetHTML5Database sets whether to enable HTML5 client-side SQL database support.
// Client-side SQL database allows web pages to store structured data and be able to use
// SQL to manipulate that data asynchronously.
func (s *Settings) SetHTML5Database(enabled bool) {
	C.webkit_settings_set_enable_html5_database(s.settings, cgobool(enabled))
}

// ----------------------------------------------------------------------------------------------------

// GetJavaAllowed gets whether or not Java is enabled on the page.
func (s *Settings) GetJavaAllowed() bool {
	return gobool(C.webkit_settings_get_enable_java(s.settings))
}

// SetJavaAllowed sets whether or not Java is enabled on the page.
func (s *Settings) SetJavaAllowed(enabled bool) {
	C.webkit_settings_set_enable_java(s.settings, cgobool(enabled))
}

// ----------------------------------------------------------------------------------------------------

// GetJavascriptAllowed gets whether or not JavaScript executes within the page.
func (s *Settings) GetJavascriptAllowed() bool {
	return gobool(C.webkit_settings_get_enable_javascript(s.settings))
}

// SetJavascriptAllowed sets whether or not JavaScript executes within the page.
func (s *Settings) SetJavascriptAllowed(enabled bool) {
	C.webkit_settings_set_enable_javascript(s.settings, cgobool(enabled))
}

// ----------------------------------------------------------------------------------------------------

// GetOfflineCache gets whether HTML5 offline web application cache support is enabled or
// not. Offline web application cache allows web applications to run even when the user
// is not connected to the network.
func (s *Settings) GetOfflineCache() bool {
	return gobool(C.webkit_settings_get_enable_offline_web_application_cache(s.settings))
}

// SetOfflineCache sets whether HTML5 offline web application cache support is enabled or
// not. Offline web application cache allows web applications to run even when the user
// is not connected to the network.
func (s *Settings) SetOfflineCache(enabled bool) {
	C.webkit_settings_set_enable_offline_web_application_cache(s.settings, cgobool(enabled))
}

// ----------------------------------------------------------------------------------------------------

// GetUserAgent gets the user-agent string used by WebKit.
func (s *Settings) GetUserAgent() string {
	return C.GoString((*C.char)(C.webkit_settings_get_user_agent(s.settings)))
}

// SetUserAgent sets the user-agent string used by WebKit. Unusual user-agent strings may
// cause web content to render incorrectly or fail to run, as many web pages are written
// to parse the user-agent strings of only the most popular browsers. Therefore, it's
// typically better to not completely override the standard user-agent, but to use
// SetUserAgentWithApplicationDetails() instead.
func (s *Settings) SetUserAgent(userAgent string) {
	C.webkit_settings_set_user_agent(s.settings, (*C.gchar)(C.CString(userAgent)))
}

// ----------------------------------------------------------------------------------------------------

// SetUserAgentWithApplicationDetails sets the “user-agent” property by appending the
// application details to the default user agent. If no application name or version is
// given, the default user agent used will be used. If only the version is given, the
// default engine version is used with the given application name.
func (s *Settings) SetUserAgentWithApplicationDetails(appName, appVersion string) {
	C.webkit_settings_set_user_agent_with_application_details(s.settings, (*C.gchar)(C.CString(appName)), (*C.gchar)(C.CString(appVersion)))
}
