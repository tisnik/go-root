package main

// #include <stdlib.h>
import "C"

func main() {
	C.srand(C.uint(42))

	for i := 1; i <= 10; i++ {
		x := C.random()
		println(x)
	}
}
