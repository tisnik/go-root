package main

import (
	"fmt"
	"simd"
)

func createMask() simd.Mask32s {
	// konstrukce a inicializace řezu
	s1 := []float32{1, 2, 3, 4, 5, 6, 7, 8}
	s2 := []float32{4, 4, 4, 4, 4, 4, 4, 4}

	// převod řezu na vektor
	v1 := simd.LoadFloat32s(s1)
	v2 := simd.LoadFloat32s(s2)

	// SIMD operace porovnání prvků vektorů
	return v1.LessEqual(v2)
}

func main() {
	mask := createMask()
	vector := mask.ToInt32s()

	// konstrukce řezu
	result := make([]int32, vector.Len())

	// zápis prvků vektoru do řezu
	vector.Store(result)

	// výpis obsahu řezu
	fmt.Println("mask as vector", result)
}
