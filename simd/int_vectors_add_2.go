// ----------------------------------------------------------------------
// Podpora nativních SIMD operací v experimentálním balíčku simd/archsimd
// ----------------------------------------------------------------------

package main

import (
	"fmt"
	"simd/archsimd"
)

func main() {
	// konstrukce a inicializace pole s odvozením počtu prvků
	a1 := [...]int8{1, 2, 3, 4, 5, 6, 7, 8, 1, 2, 3, 4, 5, 6, 7, 8}
	a2 := [...]int8{1, 2, 3, 4, 5, 6, 7, 8, 1, 2, 3, 4, 5, 6, 7, 8}

	// výpis obsahu polí
	fmt.Println("a1   ", a1)
	fmt.Println("a2   ", a2)

	// převod pole na vektor se šestnácti prvky
	v1 := archsimd.LoadInt8x16Array(&a1)
	v2 := archsimd.LoadInt8x16Array(&a2)

	// SIMD operace součtu prvků vektorů
	v3 := v1.Add(v2)

	// konstrukce řezu
	result := make([]int8, 16)

	// zápis prvků vektoru do řezu
	v3.Store(result)

	// výpis obsahu řezu
	fmt.Println("a1+a2", result)
}
