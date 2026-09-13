package main

import (
	"fmt"
	"simd"
)

func main() {
	// konstrukce a inicializace řezu
	s := []float32{-4.0, -3.0, -2.0, -1.0, 1.0, 2.0, 3.0, 4.0}

	// výpis obsahu řezu
	fmt.Println("s      ", s)

	// převod řezu na vektor
	v := simd.LoadFloat32s(s)

	// konstrukce řezu pro uložení výsledků
	result := make([]float32, v.Len())

	// SIMD operace výpočtu absolutní hodnoty
	v2 := v.Abs()

	// zápis prvků vektoru do řezu
	v2.Store(result)

	// výpis obsahu řezu
	fmt.Println("abs(s) ", result)

	// SIMD operace výpočtu hodnoty s opačným znaménkem
	v3 := v.Neg()

	// zápis prvků vektoru do řezu
	v3.Store(result)

	// výpis obsahu řezu
	fmt.Println("neg(s) ", result)

	// SIMD operace výpočtu druhé odmocniny
	v4 := v.Sqrt()

	// zápis prvků vektoru do řezu
	v4.Store(result)

	// výpis obsahu řezu
	fmt.Println("sqrt(s)", result)
}
