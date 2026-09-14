package main

import (
	"fmt"
	"simd"
)

func main() {
	// konstrukce a inicializace řezu (odlišné typy)
	s1 := []float32{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0}
	s2 := []float64{0.5, 0.5, 0.5, 0.5, 0.5, 0.5, 0.5, 0.5}

	// výpis obsahu řezů
	fmt.Println("s1   ", s1)
	fmt.Println("s2   ", s2)

	// převod řezu na vektor
	v1 := simd.LoadFloat32s(s1)
	v2 := simd.LoadFloat64s(s2)

	// SIMD operace součtu prvků vektorů
	v3 := v1.Add(v2)

	// konstrukce řezu
	result := make([]float32, v3.Len())

	// zápis prvků vektoru do řezu
	v3.Store(result)

	// výpis obsahu řezu
	fmt.Println("s1+s2", result)
}
