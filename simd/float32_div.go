package main

import (
	"fmt"
	"simd"
)

func main() {
	// konstrukce a inicializace řezu
	s1 := []float32{0.0, 1.0, -1.0, 1.0, 0.0, 1.0, -1.0, 0.0}
	s2 := []float32{1.0, 1.0, 1.0, 1.0, 0.0, 0.0, 0.0, 0.0}

	// výpis obsahu řezů
	fmt.Println("s1   ", s1)
	fmt.Println("s2   ", s2)

	// převod řezu na vektor
	v1 := simd.LoadFloat32s(s1)
	v2 := simd.LoadFloat32s(s2)

	// SIMD operace podílu prvků vektorů
	v3 := v1.Div(v2)

	// konstrukce řezu
	result := make([]float32, v3.Len())

	// zápis prvků vektoru do řezu
	v3.Store(result)

	// výpis obsahu řezu
	fmt.Println("s1+s2", result)
}
