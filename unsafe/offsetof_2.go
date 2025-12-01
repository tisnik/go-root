package main

import (
	"fmt"
	"unsafe"
)

type Struct1 struct {
	x byte
	y byte
	z byte
}

type Struct2 struct {
	x byte
	y byte
	z uint32
}

type Struct3 struct {
	x byte
	y uint32
	z byte
}

type Struct4 struct {
	x uint32
	y byte
	z byte
}

type Struct5 struct {
	x byte
	y uint16
	z uint32
}

type Struct6 struct {
	z uint32
	y uint16
	x byte
}

type Struct7 struct {
	x byte
	z uint32
	y uint16
}

type Struct8 struct {
	y uint16
	z uint32
	x byte
}

func main() {
	fmt.Printf("offsetof Struct1.x   = %d byte(s)\n", unsafe.Offsetof(Struct1{}.x))
	fmt.Printf("offsetof Struct1.y   = %d byte(s)\n", unsafe.Offsetof(Struct1{}.y))
	fmt.Printf("offsetof Struct1.z   = %d byte(s)\n", unsafe.Offsetof(Struct1{}.z))
	fmt.Println()
	fmt.Printf("offsetof Struct2.x   = %d byte(s)\n", unsafe.Offsetof(Struct2{}.x))
	fmt.Printf("offsetof Struct2.y   = %d byte(s)\n", unsafe.Offsetof(Struct2{}.y))
	fmt.Printf("offsetof Struct2.z   = %d byte(s)\n", unsafe.Offsetof(Struct2{}.z))
	fmt.Println()
	fmt.Printf("offsetof Struct3.x   = %d byte(s)\n", unsafe.Offsetof(Struct3{}.x))
	fmt.Printf("offsetof Struct3.y   = %d byte(s)\n", unsafe.Offsetof(Struct3{}.y))
	fmt.Printf("offsetof Struct3.z   = %d byte(s)\n", unsafe.Offsetof(Struct3{}.z))
	fmt.Println()
	fmt.Printf("offsetof Struct4.x   = %d byte(s)\n", unsafe.Offsetof(Struct4{}.x))
	fmt.Printf("offsetof Struct4.y   = %d byte(s)\n", unsafe.Offsetof(Struct4{}.y))
	fmt.Printf("offsetof Struct4.z   = %d byte(s)\n", unsafe.Offsetof(Struct4{}.z))
	fmt.Println()
	fmt.Printf("offsetof Struct5.x   = %d byte(s)\n", unsafe.Offsetof(Struct5{}.x))
	fmt.Printf("offsetof Struct5.y   = %d byte(s)\n", unsafe.Offsetof(Struct5{}.y))
	fmt.Printf("offsetof Struct5.z   = %d byte(s)\n", unsafe.Offsetof(Struct5{}.z))
	fmt.Println()
	fmt.Printf("offsetof Struct6.x   = %d byte(s)\n", unsafe.Offsetof(Struct6{}.x))
	fmt.Printf("offsetof Struct6.y   = %d byte(s)\n", unsafe.Offsetof(Struct6{}.y))
	fmt.Printf("offsetof Struct6.z   = %d byte(s)\n", unsafe.Offsetof(Struct6{}.z))
	fmt.Println()
	fmt.Printf("offsetof Struct7.x   = %d byte(s)\n", unsafe.Offsetof(Struct7{}.x))
	fmt.Printf("offsetof Struct7.y   = %d byte(s)\n", unsafe.Offsetof(Struct7{}.y))
	fmt.Printf("offsetof Struct7.z   = %d byte(s)\n", unsafe.Offsetof(Struct7{}.z))
	fmt.Println()
	fmt.Printf("offsetof Struct8.x   = %d byte(s)\n", unsafe.Offsetof(Struct8{}.x))
	fmt.Printf("offsetof Struct8.y   = %d byte(s)\n", unsafe.Offsetof(Struct8{}.y))
	fmt.Printf("offsetof Struct8.z   = %d byte(s)\n", unsafe.Offsetof(Struct8{}.z))
	fmt.Println()
}
