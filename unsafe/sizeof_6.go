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
	const s01 = unsafe.Sizeof(int8(0))
	const s02 = unsafe.Sizeof(int16(0))
	const s03 = unsafe.Sizeof(int32(0))
	const s04 = unsafe.Sizeof(int64(0))
	const s05 = unsafe.Sizeof(int(0))
	const s06 = unsafe.Sizeof(float32(0))
	const s07 = unsafe.Sizeof(float64(0))
	const s08 = unsafe.Sizeof(complex64(0))
	const s09 = unsafe.Sizeof(complex128(0))
	const s10 = unsafe.Sizeof(uintptr(0))

	fmt.Printf("sizeof int8        = %d byte(s)\n", s01)
	fmt.Printf("sizeof int16       = %d byte(s)\n", s02)
	fmt.Printf("sizeof int32       = %d byte(s)\n", s03)
	fmt.Printf("sizeof int64       = %d byte(s)\n", s04)
	fmt.Printf("sizeof int         = %d byte(s)\n", s05)
	fmt.Printf("sizeof float32     = %d byte(s)\n", s06)
	fmt.Printf("sizeof float64     = %d byte(s)\n", s07)
	fmt.Printf("sizeof complex64   = %d byte(s)\n", s08)
	fmt.Printf("sizeof complex128  = %d byte(s)\n", s09)
	fmt.Printf("sizeof uintptr     = %d byte(s)\n", s10)
	fmt.Println()

	array1 := [...]int32{1}
	array2 := [...]int32{1, 2}
	array3 := [...]int32{1, 2, 3}

	const s11 = unsafe.Sizeof(array1)
	const s12 = unsafe.Sizeof(array2)
	const s13 = unsafe.Sizeof(array3)

	fmt.Printf("sizeof array1      = %d byte(s)\n", s11)
	fmt.Printf("sizeof array2      = %d byte(s)\n", s12)
	fmt.Printf("sizeof array3      = %d byte(s)\n", s13)
	fmt.Println()

	slice1 := []int32{1}
	slice2 := []int32{1, 2}
	slice3 := []int32{1, 2, 3}

	const s14 = unsafe.Sizeof(slice1)
	const s15 = unsafe.Sizeof(slice2)
	const s16 = unsafe.Sizeof(slice3)

	fmt.Printf("sizeof slice1      = %d byte(s)\n", s14)
	fmt.Printf("sizeof slice2      = %d byte(s)\n", s15)
	fmt.Printf("sizeof slice3      = %d byte(s)\n", s16)
	fmt.Println()

	fmt.Printf("sizeof str         = %d byte(s)\n", unsafe.Sizeof("foo"))
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

	const s17 = unsafe.Sizeof(vector2d)
	const s18 = unsafe.Sizeof(vector3d)
	const s19 = unsafe.Sizeof(user)
	const s20 = unsafe.Sizeof(&vector2d)
	const s21 = unsafe.Sizeof(&vector3d)
	const s22 = unsafe.Sizeof(&user)

	fmt.Printf("sizeof Vector2D    = %d byte(s)\n", s17)
	fmt.Printf("sizeof Vector3D    = %d byte(s)\n", s18)
	fmt.Printf("sizeof User        = %d byte(s)\n", s19)
	fmt.Println()

	fmt.Printf("sizeof &Vector2D   = %d byte(s)\n", s20)
	fmt.Printf("sizeof &Vector3D   = %d byte(s)\n", s21)
	fmt.Printf("sizeof &User       = %d byte(s)\n", s22)

	var a *int32 = nil
	var b *int64 = nil
	var c []int32 = nil
	var d interface{} = nil
	var e map[string]string = nil
	var f chan int = nil

	const s23 = unsafe.Sizeof(a)
	const s24 = unsafe.Sizeof(b)
	const s25 = unsafe.Sizeof(c)
	const s26 = unsafe.Sizeof(d)
	const s27 = unsafe.Sizeof(e)
	const s28 = unsafe.Sizeof(f)

	const s29 = unsafe.Sizeof(Struct1{})
	const s30 = unsafe.Sizeof(Struct2{})
	const s31 = unsafe.Sizeof(Struct3{})
	const s32 = unsafe.Sizeof(Struct4{})

	fmt.Printf("sizeof *int32      = %d byte(s)\n", s23)
	fmt.Printf("sizeof *int64      = %d byte(s)\n", s24)
	fmt.Printf("sizeof slice       = %d byte(s)\n", s25)
	fmt.Printf("sizeof {}interface = %d byte(s)\n", s26)
	fmt.Printf("sizeof map         = %d byte(s)\n", s27)
	fmt.Printf("sizeof chan        = %d byte(s)\n", s28)

	fmt.Printf("sizeof Struct1     = %d byte(s)\n", s29)
	fmt.Printf("sizeof Struct2     = %d byte(s)\n", s30)
	fmt.Printf("sizeof Struct3     = %d byte(s)\n", s31)
	fmt.Printf("sizeof Struct4     = %d byte(s)\n", s32)
}
