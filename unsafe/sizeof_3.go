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

func main() {
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

	fmt.Printf("sizeof str      = %d byte(s)\n", unsafe.Sizeof("foo"))
	fmt.Println()

	fmt.Printf("sizeof Vector2D = %d byte(s)\n", unsafe.Sizeof(vector2d))
	fmt.Printf("sizeof Vector3D = %d byte(s)\n", unsafe.Sizeof(vector3d))
	fmt.Printf("sizeof User     = %d byte(s)\n", unsafe.Sizeof(user))
	fmt.Println()

	fmt.Printf("sizeof &Vector2D = %d byte(s)\n", unsafe.Sizeof(&vector2d))
	fmt.Printf("sizeof &Vector3D = %d byte(s)\n", unsafe.Sizeof(&vector3d))
	fmt.Printf("sizeof &User     = %d byte(s)\n", unsafe.Sizeof(&user))
}
