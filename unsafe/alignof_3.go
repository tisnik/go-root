package main

import (
	"fmt"
	"unsafe"
)

type Struct1 struct {
	x byte
}

type Struct2 struct {
	x byte
	y byte
}

type Struct3 struct {
	x byte
	y byte
	z byte
}

type Struct4 struct {
	x byte
	y byte
	z byte
	w byte
}

func main() {
	fmt.Printf("alignof Struct1   = %d byte(s)\n", unsafe.Alignof(Struct1{}))
	fmt.Printf("alignof Struct2   = %d byte(s)\n", unsafe.Alignof(Struct2{}))
	fmt.Printf("alignof Struct3   = %d byte(s)\n", unsafe.Alignof(Struct3{}))
	fmt.Printf("alignof Struct4   = %d byte(s)\n", unsafe.Alignof(Struct4{}))
}
