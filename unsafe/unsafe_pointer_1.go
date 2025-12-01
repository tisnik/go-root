package main

// #include <stdio.h>
// #include <stdlib.h>
import "C"
import "unsafe"

func main() {
	cs := C.CString("Hello world!")
	C.puts(cs)
	C.free(unsafe.Pointer(cs))
}
