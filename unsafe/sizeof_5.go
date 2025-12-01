package main

import (
	"fmt"
	"unsafe"
)

type Struct1 struct {
	x byte
	y uint16
	z uint32
}

type Struct2 struct {
	z uint32
	y uint16
	x byte
}

type Struct3 struct {
	x byte
	z uint32
	y uint16
}

type Struct4 struct {
	y uint16
	z uint32
	x byte
}

func main() {
	fmt.Printf("sizeof Struct1   = %d byte(s)\n", unsafe.Sizeof(Struct1{}))
	fmt.Printf("sizeof Struct2   = %d byte(s)\n", unsafe.Sizeof(Struct2{}))
	fmt.Printf("sizeof Struct3   = %d byte(s)\n", unsafe.Sizeof(Struct3{}))
	fmt.Printf("sizeof Struct4   = %d byte(s)\n", unsafe.Sizeof(Struct4{}))
}
