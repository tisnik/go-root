package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	ID      uint32
	Name    string
	Surname string
}

func main() {
	var a *int32 = nil
	var b *int64 = nil
	var c []int32 = nil
	var d interface{} = nil
	var e map[string]string = nil
	var f chan int = nil

	fmt.Printf("sizeof *int32      = %d byte(s)\n", unsafe.Sizeof(a))
	fmt.Printf("sizeof *int64      = %d byte(s)\n", unsafe.Sizeof(b))
	fmt.Printf("sizeof slice       = %d byte(s)\n", unsafe.Sizeof(c))
	fmt.Printf("sizeof {}interface = %d byte(s)\n", unsafe.Sizeof(d))
	fmt.Printf("sizeof map         = %d byte(s)\n", unsafe.Sizeof(e))
	fmt.Printf("sizeof chan        = %d byte(s)\n", unsafe.Sizeof(f))
}
