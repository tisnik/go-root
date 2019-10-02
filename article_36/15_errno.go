package main

// #include <math.h>
// #cgo LDFLAGS: -lm
import "C"
import "fmt"

func main() {
	y, err := C.sqrt(C.double(100.0))
	fmt.Printf("result=%v err=%v\n", y, err)

	y, err = C.sqrt(C.double(-1.0))
	fmt.Printf("result=%v err=%v\n", y, err)

	y, err = C.sqrt(C.double(100.0))
	fmt.Printf("result=%v err=%v\n", y, err)
}
