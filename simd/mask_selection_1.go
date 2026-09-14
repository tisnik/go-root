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
	return v1.GreaterEqual(v2)
}

func main() {
	mask := createMask()

	// konstrukce a inicializace řezu
	s := []float32{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0}

	// výpis obsahu řezu
	fmt.Println("original", s)

	// převod řezu na vektor
	v1 := simd.LoadFloat32s(s)

	// výběr prvků na základě masky
	v2 := v1.Masked(mask)

	// konstrukce řezu
	result := make([]float32, v2.Len())

	// zápis prvků vektoru do řezu
	v2.Store(result)

	// výpis obsahu řezu
	fmt.Println("selected", result)
}
