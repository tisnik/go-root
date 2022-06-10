// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Třicátá čtvrtá část
//    Programovací jazyk Go pro skalní céčkaře
//    https://www.root.cz/clanky/programovaci-jazyk-go-pro-skalni-ceckare/
//
// Seznam příkladů ze třicáté čtvrté části:
//    https://github.com/tisnik/go-root/blob/master/article_34/README.md
//
// Demonstrační příklad číslo 3:
//    Zjištění velikostí základních datových typů jazyka Go.

package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Printf("sizeof int8       = %d byte(s)\n", unsafe.Sizeof(int8(0)))
	fmt.Printf("sizeof int16      = %d byte(s)\n", unsafe.Sizeof(int16(0)))
	fmt.Printf("sizeof int32      = %d byte(s)\n", unsafe.Sizeof(int32(0)))
	fmt.Printf("sizeof int64      = %d byte(s)\n", unsafe.Sizeof(int64(0)))
	fmt.Printf("sizeof int        = %d byte(s)\n", unsafe.Sizeof(int(0)))
	fmt.Printf("sizeof float32    = %d byte(s)\n", unsafe.Sizeof(float32(0)))
	fmt.Printf("sizeof float64    = %d byte(s)\n", unsafe.Sizeof(float64(0)))
	fmt.Printf("sizeof complex64  = %d byte(s)\n", unsafe.Sizeof(complex64(0)))
	fmt.Printf("sizeof complex128 = %d byte(s)\n", unsafe.Sizeof(complex128(0)))
	fmt.Printf("sizeof uintptr    = %d byte(s)\n", unsafe.Sizeof(uintptr(0)))
}
