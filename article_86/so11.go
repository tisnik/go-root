// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam příkladů z osmdesáté šesté části:
//    https://github.com/tisnik/go-root/blob/master/article_86/README.md

package main

import (
	"fmt"
	"math"
)

/*
struct Vector {
    double X;
    double Y;
};
*/
import "C"

//export length
func length(v C.struct_Vector) C.double {
	r := C.double(math.Sqrt(float64(v.X*v.X) + float64(v.Y*v.Y)))
	fmt.Printf("%f %f %f\n", float64(v.X), float64(v.Y), r)
	return r
}

func main() {}
