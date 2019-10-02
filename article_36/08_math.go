package main

// #include <math.h>
import "C"
import "fmt"

func main() {
	x := 3.1415927 / 6.0
	y := float32(C.sinf(C.float(x)))
	fmt.Printf("%v\n", y)
}
