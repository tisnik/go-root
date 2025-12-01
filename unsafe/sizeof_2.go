package main

import (
	"fmt"
	"unsafe"
)

func main() {
	array1 := [...]int32{1}
	array2 := [...]int32{1, 2}
	array3 := [...]int32{1, 2, 3}

	slice1 := []int32{1}
	slice2 := []int32{1, 2}
	slice3 := []int32{1, 2, 3}

	fmt.Printf("sizeof int    = %d byte(s)\n", unsafe.Sizeof(int32(42)))
	fmt.Println()

	fmt.Printf("sizeof array1 = %d byte(s)\n", unsafe.Sizeof(array1))
	fmt.Printf("sizeof array2 = %d byte(s)\n", unsafe.Sizeof(array2))
	fmt.Printf("sizeof array3 = %d byte(s)\n", unsafe.Sizeof(array3))
	fmt.Println()

	fmt.Printf("sizeof slice1 = %d byte(s)\n", unsafe.Sizeof(slice1))
	fmt.Printf("sizeof slice2 = %d byte(s)\n", unsafe.Sizeof(slice2))
	fmt.Printf("sizeof slice3 = %d byte(s)\n", unsafe.Sizeof(slice3))
}
