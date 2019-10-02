package main

// #include <stdio.h>
// #include <stdlib.h>
import "C"
import "unsafe"

func main() {
	cs := C.CString("Hello world!")
	defer C.free(unsafe.Pointer(cs))
	C.puts(cs)
}
