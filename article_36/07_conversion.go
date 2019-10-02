package main

// #include <stdlib.h>
import "C"
import "fmt"

func main() {
	x := -10
	var y int = int(C.abs(C.int(x)))
	fmt.Printf("%v\n", y)
}
