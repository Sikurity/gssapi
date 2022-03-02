//go:build windows
// +build windows

package gssapi

/*
#include <Windows.h>

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

	if handle = unsafe.Pointer(C.LoadLibrary(cpath)); handle == nil {
		err = fmt.Errorf("cannot open library[%s]", path)
	}

	return
}
func (loader LoaderImpl) Close(handle unsafe.Pointer) (err error) {
	if int(C.FreeLibrary(C.HINSTANCE(handle))) == 0 {
		err = fmt.Errorf("failed free library[%d]", handle)
	}

	return
}
func (loader LoaderImpl) FuncHandle(handle unsafe.Pointer, funcName string) (symbol unsafe.Pointer, err error) {
	cfname := C.CString(funcName)
	defer C.free(unsafe.Pointer(cfname))

	if symbol = unsafe.Pointer(C.GetProcAddress(C.HINSTANCE(handle), cfname)); symbol == nil {
		err = fmt.Errorf("cannot get function[%s] from library handle[%d]", funcName, handle)
	}

	return
}
