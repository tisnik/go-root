package main

import (
	"fmt"
	"simd"
)

func main() {
	// konstrukce a inicializace řezu (odlišné typy)
	s1 := []float32{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0, 11.0, 12.0, 13.0, 14.0, 15.0, 16.0}
	s2 := []float32{0.5, 0.5, 0.5, 0.5, 0.5, 0.5, 0.5, 0.5}

	// výpis obsahu řezů
	fmt.Println("s1   ", s1)
	fmt.Println("s2   ", s2)

	// převod řezu na vektor
	v1 := simd.LoadFloat32s(s1)
	v2 := simd.LoadFloat32s(s2)

	// výpis délky vektorů
	fmt.Println("len(v1)=", v1.Len())
	fmt.Println("len(v2)=", v2.Len())

	// SIMD operace součtu prvků vektorů
	v3 := v1.Add(v2)

	// konstrukce řezu
	result := make([]float32, v3.Len())

	// zápis prvků vektoru do řezu
	v3.Store(result)

	// výpis obsahu řezu
	fmt.Println("s1+s2", result)
}
