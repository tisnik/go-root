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

	fmt.Printf("alignof Vector2D = %d byte(s)\n", unsafe.Alignof(vector2d))
	fmt.Printf("alignof Vector3D = %d byte(s)\n", unsafe.Alignof(vector3d))
	fmt.Printf("alignof User     = %d byte(s)\n", unsafe.Alignof(user))
	fmt.Println()
}
