package main

import (
	"fmt"
	"simd"
)

func main() {
	// výpis bitové šířky vektorů
	fmt.Println("vector bit size=", simd.VectorBitSize())
}
