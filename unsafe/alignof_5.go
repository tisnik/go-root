package main

import (
	"fmt"
	"unsafe"
)

type Vector2D struct {
	x float32
	y float32
}

type Vector3D struct {
	x float32
	y float32
	z float32
}

type User struct {
	ID      uint32
	Name    string
	Surname string
}

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

type Struct5 struct {
	i int32
	x byte
}

type Struct6 struct {
	i int32
	x byte
	y byte
}

type Struct7 struct {
	i int32
	x byte
	y byte
	z byte
}

type Struct8 struct {
	i int32
	x byte
	y byte
	z byte
	w byte
}

func main() {
	const c01 = unsafe.Alignof(int8(0))
	const c02 = unsafe.Alignof(int16(0))
	const c03 = unsafe.Alignof(int32(0))
	const c04 = unsafe.Alignof(int64(0))
	const c05 = unsafe.Alignof(int(0))
	const c06 = unsafe.Alignof(float32(0))
	const c07 = unsafe.Alignof(float64(0))
	const c08 = unsafe.Alignof(complex64(0))
	const c09 = unsafe.Alignof(complex128(0))
	const c10 = unsafe.Alignof(uintptr(0))

	fmt.Printf("alignof int8       = %d byte(s)\n", c01)
	fmt.Printf("alignof int16      = %d byte(s)\n", c02)
	fmt.Printf("alignof int32      = %d byte(s)\n", c03)
	fmt.Printf("alignof int64      = %d byte(s)\n", c04)
	fmt.Printf("alignof int        = %d byte(s)\n", c05)
	fmt.Printf("alignof float32    = %d byte(s)\n", c06)
	fmt.Printf("alignof float64    = %d byte(s)\n", c07)
	fmt.Printf("alignof complex64  = %d byte(s)\n", c08)
	fmt.Printf("alignof complex128 = %d byte(s)\n", c09)
	fmt.Printf("alignof uintptr    = %d byte(s)\n", c10)
	fmt.Println()

	vector2d := Vector2D{
		x: 10,
		y: 20,
	}

	vector3d := Vector3D{
		x: 10,
		y: 20,
		z: 30,
	}

	user := User{
		ID:      42,
		Name:    "Pepa",
		Surname: "Vyskoč",
	}

	const c11 = unsafe.Alignof(vector2d)
	const c12 = unsafe.Alignof(vector3d)
	const c13 = unsafe.Alignof(user)

	fmt.Printf("alignof Vector2D   = %d byte(s)\n", c11)
	fmt.Printf("alignof Vector3D   = %d byte(s)\n", c12)
	fmt.Printf("alignof User       = %d byte(s)\n", c13)
	fmt.Println()

	const c14 = unsafe.Alignof(Struct1{})
	const c15 = unsafe.Alignof(Struct2{})
	const c16 = unsafe.Alignof(Struct3{})
	const c17 = unsafe.Alignof(Struct4{})

	fmt.Printf("alignof Struct1    = %d byte(s)\n", c14)
	fmt.Printf("alignof Struct2    = %d byte(s)\n", c15)
	fmt.Printf("alignof Struct3    = %d byte(s)\n", c16)
	fmt.Printf("alignof Struct4    = %d byte(s)\n", c17)
	fmt.Println()

	const c18 = unsafe.Alignof(Struct5{})
	const c19 = unsafe.Alignof(Struct6{})
	const c20 = unsafe.Alignof(Struct7{})
	const c21 = unsafe.Alignof(Struct8{})

	fmt.Printf("alignof Struct5    = %d byte(s)\n", c18)
	fmt.Printf("alignof Struct6    = %d byte(s)\n", c19)
	fmt.Printf("alignof Struct7    = %d byte(s)\n", c20)
	fmt.Printf("alignof Struct8    = %d byte(s)\n", c21)
}
