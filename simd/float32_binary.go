package main

import (
	"fmt"
	"simd"
)

func main() {
	// konstrukce a inicializace řezu
	s1 := []float32{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0}
	s2 := []float32{0.5, 0.5, 0.5, 0.5, 0.5, 0.5, 0.5, 0.5}

	// výpis obsahu řezů
	fmt.Println("s1        ", s1)
	fmt.Println("s2        ", s2)

	// převod řezu na vektor
	v1 := simd.LoadFloat32s(s1)
	v2 := simd.LoadFloat32s(s2)

	// konstrukce řezu pro uložení výsledků
	result := make([]float32, v1.Len())

	// SIMD operace součtu prvků vektorů
	v3 := v1.Add(v2)

	// zápis prvků vektoru do řezu
	v3.Store(result)

	// výpis obsahu řezu
	fmt.Println("s1+s2     ", result)

	// SIMD operace rozdílu prvků vektorů
	v4 := v1.Sub(v2)

	// zápis prvků vektoru do řezu
	v4.Store(result)

	// výpis obsahu řezu
	fmt.Println("s1-s2     ", result)

	// SIMD operace součinu prvků vektorů
	v5 := v1.Mul(v2)

	// zápis prvků vektoru do řezu
	v5.Store(result)

	// výpis obsahu řezu
	fmt.Println("s1*s2     ", result)

	// SIMD operace podílu prvků vektorů
	v6 := v1.Div(v2)

	// zápis prvků vektoru do řezu
	v6.Store(result)

	// výpis obsahu řezu
	fmt.Println("s1/s2     ", result)

	// SIMD operace výběru větších prvků vektorů
	v7 := v1.Max(v2)

	// zápis prvků vektoru do řezu
	v7.Store(result)

	// výpis obsahu řezu
	fmt.Println("max(s1/s2)", result)

	// SIMD operace výběru menších prvků vektorů
	v8 := v1.Min(v2)

	// zápis prvků vektoru do řezu
	v8.Store(result)

	// výpis obsahu řezu
	fmt.Println("min(s1/s2)", result)
}
