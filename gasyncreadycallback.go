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

// #include <gio/gio.h>
// #include "gasyncreadycallback.go.h"
import "C"
import (
	"errors"
	"reflect"
	"unsafe"
)

type garCallback struct {
	f reflect.Value
}

//export _go_gasyncreadycallback_call
func _go_gasyncreadycallback_call(cbinfoRaw C.gpointer, cresult unsafe.Pointer) {
	result := (*C.GAsyncResult)(cresult)
	cbinfo := (*garCallback)(unsafe.Pointer(cbinfoRaw))
	cbinfo.f.Call([]reflect.Value{reflect.ValueOf(result)})
}

func newGAsyncReadyCallback(f interface{}) (cCallback C.GAsyncReadyCallback, userData C.gpointer, err error) {
	rf := reflect.ValueOf(f)
	if rf.Kind() != reflect.Func {
		return nil, nil, errors.New("f is not a function")
	}
	data := C.malloc(C.size_t(unsafe.Sizeof(garCallback{})))
	cbinfo := (*garCallback)(data)
	cbinfo.f = rf
	return C.GAsyncReadyCallback(C._gasyncreadycallback_call), C.gpointer(unsafe.Pointer(cbinfo)), nil
}
