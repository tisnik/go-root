package main

import (
	"fmt"
	"simd"
)

func main() {
	// konstrukce a inicializace řezu
	s1 := []float32{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0}
	s2 := []float32{2.0, 2.0, 2.0, 2.0, 2.0, 2.0, 2.0, 2.0}
	s3 := []float32{0.5, 0.5, 0.5, 0.5, 0.5, 0.5, 0.5, 0.5}

	// výpis obsahu řezů
	fmt.Println("s1      ", s1)
	fmt.Println("s2      ", s2)
	fmt.Println("s3      ", s3)

	// převod řezu na vektor
	v1 := simd.LoadFloat32s(s1)
	v2 := simd.LoadFloat32s(s2)
	v3 := simd.LoadFloat32s(s3)

	// SIMD operace multiply-add
	v4 := v1.MulAdd(v2, v3)

	// konstrukce řezu
	result := make([]float32, v4.Len())

	// zápis prvků vektoru do řezu
	v4.Store(result)

	// výpis obsahu řezu
	fmt.Println("s1*s2+s3", result)
}
