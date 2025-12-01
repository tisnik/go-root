package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Printf("alignof int8       = %d byte(s)\n", unsafe.Alignof(int8(0)))
	fmt.Printf("alignof int16      = %d byte(s)\n", unsafe.Alignof(int16(0)))
	fmt.Printf("alignof int32      = %d byte(s)\n", unsafe.Alignof(int32(0)))
	fmt.Printf("alignof int64      = %d byte(s)\n", unsafe.Alignof(int64(0)))
	fmt.Printf("alignof int        = %d byte(s)\n", unsafe.Alignof(int(0)))
	fmt.Printf("alignof float32    = %d byte(s)\n", unsafe.Alignof(float32(0)))
	fmt.Printf("alignof float64    = %d byte(s)\n", unsafe.Alignof(float64(0)))
	fmt.Printf("alignof complex64  = %d byte(s)\n", unsafe.Alignof(complex64(0)))
	fmt.Printf("alignof complex128 = %d byte(s)\n", unsafe.Alignof(complex128(0)))
	fmt.Printf("alignof uintptr    = %d byte(s)\n", unsafe.Alignof(uintptr(0)))
}
