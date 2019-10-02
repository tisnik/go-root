package main

// #include <stdio.h>
import "C"

func main() {
	cs := C.CString("Hello world!")
	C.puts(cs)
}
