//go:build linux || freebsd || darwin
// +build linux freebsd darwin

package gssapi

/*
#cgo linux LDFLAGS: -ldl
#cgo freebsd pkg-config: heimdal-gssapi

#include <dlfcn.h>

*/
import "C"
import (
	"fmt"
	"unsafe"
)

type LoaderImpl struct {
}

func (loader LoaderImpl) Open(path string) (handle unsafe.Pointer, err error) {
	cpath := C.CString(path)
	defer C.free(unsafe.Pointer(cpath))

	if handle = unsafe.Pointer(C.dlopen(cpath, C.RTLD_NOW|C.RTLD_LOCAL)); handle == nil {
		errMsg := C.GoString(C.dlerror())
		defer C.free(unsafe.Pointer(errMsg))
		err = fmt.Errorf(errMsg)
	}

	return
}
func (loader LoaderImpl) Close(handle unsafe.Pointer) (err error) {
	if C.dlclose(handle) < 0 {
		errMsg := C.GoString(C.dlerror())
		defer C.free(unsafe.Pointer(errMsg))
		err = fmt.Errorf(errMsg)
	}

	return
}
func (loader LoaderImpl) FuncHandle(handle unsafe.Pointer, funcName string) (symbol unsafe.Pointer, err error) {
	cfname := C.CString(funcName)
	defer C.free(unsafe.Pointer(cfname))

	if symbol = unsafe.Pointer(C.dlsym(handle, C.CString(cfname))); symbol == nil {
		errMsg := C.GoString(C.dlerror())
		defer C.free(unsafe.Pointer(errMsg))
		err = fmt.Errorf(errMsg)
	}

	return
}
