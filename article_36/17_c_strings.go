package main

// #include <stdio.h>
// #include <stdlib.h>
import "C"
import "fmt"

func main() {
	fmt.Print("? ")
	s := C.calloc(20, 1)
	char_ptr := (*C.char)(s)
	C.fgets(char_ptr, 20, C.stdin)
	fmt.Printf("result=%v\n", C.GoString(char_ptr))
	C.free(s)
}
