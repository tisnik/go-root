// Seriál "Programovací jazyk Go"
//
// Druhá část
//
// Demonstrační příklad číslo 8
//    Konverze mezi základními datovými typy

package main

import "fmt"

func main() {
	var a int8 = -10
	var signed_int int32 = -100000
	var unsigned_int uint32 = 100000
	var e float32 = 1e4
	var f float64 = 1.5e30

	var x int32 = int32(a)
	var y int32 = int32(e)
	var z float32 = float32(f)

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	var b2 uint8 = uint8(signed_int)
	var b3 uint8 = uint8(unsigned_int)

	fmt.Println(b2)
	fmt.Println(b3)
}
