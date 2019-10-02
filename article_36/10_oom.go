package main

// #include <stdio.h>
import "C"

func main() {
	for {
		cs := C.CString("Hello world!")
		C.puts(cs)
	}
}
