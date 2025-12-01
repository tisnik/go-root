package main

import (
	"fmt"
	"unsafe"
)

type Vector3D struct {
	x float32
	y float32
	z float32
}

func main() {
	vector3d := Vector3D{
		x: 10,
		y: 20,
		z: 30,
	}

	fmt.Printf("offsetof Vector3D.x = %d byte(s)\n", unsafe.Offsetof(vector3d.x))
	fmt.Printf("offsetof Vector3D.y = %d byte(s)\n", unsafe.Offsetof(vector3d.y))
	fmt.Printf("offsetof Vector3D.a = %d byte(s)\n", unsafe.Offsetof(vector3d.z))
	fmt.Println()
}
