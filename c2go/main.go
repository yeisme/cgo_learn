package main

/*
#include <stdio.h>
#include <stdlib.h>
*/
import "C"
import "unsafe"

func main() {
	cstr := C.CString("Hello, World!")
	defer C.free(unsafe.Pointer(cstr))
	C.puts(cstr)
}
